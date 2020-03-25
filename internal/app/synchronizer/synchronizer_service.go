// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package synchronizer

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/sirupsen/logrus"
	"open-match.dev/open-match/internal/config"
	"open-match.dev/open-match/internal/ipb"
	"open-match.dev/open-match/internal/statestore"
	"open-match.dev/open-match/pkg/pb"
)

var (
	logger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "app.synchronizer",
	})
)

// Matches flow through channels in the synchronizer.  Channel variable names
// are used to be consistent between function calls to help track everything.

// Streams from multiple GRPC calls of matches are combined on a single channel.
// These matches are sent to the evaluator, then the tickets are added to the
// ignore list.  Finally the matches are returned to the calling stream.

// receive from backend                  | Synchronize
//  -> m1c ->
// close incoming when done or timed out | newCutoffSender
//   -> m2c ->
// remember return channel m7c for match | fanInFanOut
//   -> m3c ->
// setmappings from matchIDs to ticketIDs| cacheMatchIDToTicketIDs
//   -> m4c -> (buffered)
// send to evaluator                     | wrapEvaluator
//   -> m5c -> (buffered)
// add tickets to ignore list            | addMatchesToIgnoreList
//   -> m6c ->
// fan out to origin synchronize call    | fanInFanOut
//   -> (Synchronize call specific ) m7c -> (buffered)
// return to backend                     | Synchronize

type synchronizerService struct {
	cfg   config.View
	store statestore.Service
	eval  evaluator

	synchronizeRegistration chan *registrationRequest

	// startCycle is a buffered channel for containing a single value.  The value
	// is present only when a cycle is not running.
	startCycle chan struct{}
}

func newSynchronizerService(cfg config.View, eval evaluator, store statestore.Service) *synchronizerService {
	s := &synchronizerService{
		cfg:   cfg,
		store: store,
		eval:  eval,

		synchronizeRegistration: make(chan *registrationRequest),
		startCycle:              make(chan struct{}, 1),
	}

	s.startCycle <- struct{}{}

	return s
}

func (s *synchronizerService) Synchronize(stream ipb.Synchronizer_SynchronizeServer) error {
	// Synchronize first registers against a cycle.  Then it creates two go
	// routines:
	// 1. Receive proposals from backend, send them to cycle.
	// 2. Receive matches and signals from cycle, send them to backend.

	eg, ctx := errgroup.WithContext(stream.Context())
	registration := s.register(ctx)
	m6cBuffer := bufferStringChannel(registration.m7c)
	defer func() {
		for range m6cBuffer {
		}
	}()

	matches := map[string]*pb.Match{}
	eg.Go(func() error {
		defer registration.allM1cSent.Done()

		duplicateIDs := []string{}
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("error streaming in synchronizer from backend, desc: %w", err)
			}
			if _, ok := matches[req.GetProposal().GetMatchId()]; ok {
				duplicateIDs = append(duplicateIDs, req.GetProposal().GetMatchId())
			}
			matches[req.GetProposal().GetMatchId()] = req.GetProposal()
		}

		if len(duplicateIDs) != 0 {
			return fmt.Errorf("found duplicate matchIDs %s from different FetchMatches calls", duplicateIDs)
		}

		for _, proposal := range matches {
			registration.m1c.send(mAndM6c{m: proposal, m7c: registration.m7c})
		}

		return nil
	})

	err := stream.Send(&ipb.SynchronizeResponse{StartMmfs: true})
	if err != nil {
		return err
	}

	eg.Go(func() error {
		for {
			select {
			case mIDs, ok := <-m6cBuffer:
				if !ok {
					return nil
				}
				for _, mID := range mIDs {
					err = stream.Send(&ipb.SynchronizeResponse{MatchId: mID})
					if err != nil {
						logger.WithFields(logrus.Fields{
							"error": err.Error(),
						}).Error("error streaming match in synchronizer to backend")
						return err
					}
				}
			case <-registration.cancelMmfs:
				err = stream.Send(&ipb.SynchronizeResponse{CancelMmfs: true})
				if err != nil {
					logger.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Error("error streaming mmf cancel in synchronizer to backend")
					return err
				}
			case <-ctx.Done():
				logger.WithFields(logrus.Fields{
					"error": ctx.Err().Error(),
				}).Error("error streaming in synchronizer to backend: context is done")
				return ctx.Err()
			case <-registration.cycleCtx.Done():
				return registration.cycleCtx.Err()
			}
		}
	})

	return eg.Wait()
}

///////////////////////////////////////
///////////////////////////////////////

// Registration of a Synchronize call for a cycle does the following:
// - Sends a registration request, starting a cycle if none is running.
// - The cycle creates the registration.
// - The registration is sent back to the origin synchronize call on channel as
//     part of the sychronize request.

type registrationRequest struct {
	resp chan *registration
	ctx  context.Context
}

type registration struct {
	m1c        *cutoffSender
	allM1cSent *sync.WaitGroup
	m7c        chan string
	cancelMmfs chan struct{}
	cycleCtx   context.Context
}

func (s synchronizerService) register(ctx context.Context) *registration {
	req := &registrationRequest{
		resp: make(chan *registration),
		ctx:  ctx,
	}
	for {
		select {
		case s.synchronizeRegistration <- req:
			return <-req.resp
		case <-s.startCycle:
			go func() {
				s.runCycle()
				s.startCycle <- struct{}{}
			}()
		}
	}
}

///////////////////////////////////////
///////////////////////////////////////

func (s *synchronizerService) runCycle() {
	/////////////////////////////////////// Initialize cycle
	ctx, cancel := withCancelCause(context.Background())

	m2c := make(chan mAndM6c)
	m3c := make(chan *pb.Match)
	m4c := make(chan *pb.Match)
	m5c := make(chan string)
	m6c := make(chan string)

	m1c := newCutoffSender(m2c)
	// m7c, unlike other channels, is specific to a synchronize call.  There are
	// multiple values in a given cycle.

	var allM1cSent sync.WaitGroup

	registrations := []*registration{}
	callingCtx := []context.Context{}
	closedOnCycleEnd := make(chan struct{})

	go func() {
		fanInFanOut(m2c, m3c, m6c)
		// Close response channels after all responses have been sent.
		for _, r := range registrations {
			close(r.m7c)
		}
	}()

	matchTickets := &sync.Map{}
	go s.cacheMatchIDToTicketIDs(matchTickets, m3c, m4c)
	go s.wrapEvaluator(ctx, cancel, bufferMatchChannel(m4c), m5c)
	go func() {
		s.addMatchesToIgnoreList(ctx, matchTickets, cancel, bufferStringChannel(m5c), m6c)
		// Wait for ignore list, but not all matches returned, the next cycle
		// can start now.
		close(closedOnCycleEnd)
	}()

	/////////////////////////////////////// Run Registration Period
	closeRegistration := time.After(s.registrationInterval())
Registration:
	for {
		select {
		case req := <-s.synchronizeRegistration:
			allM1cSent.Add(1)
			callingCtx = append(callingCtx, req.ctx)
			r := &registration{
				m1c:        m1c,
				m7c:        make(chan string),
				cancelMmfs: make(chan struct{}, 1),
				cycleCtx:   ctx,
				allM1cSent: &allM1cSent,
			}
			registrations = append(registrations, r)
			req.resp <- r
		case <-closeRegistration:
			break Registration
		}
	}
	/////////////////////////////////////// Wait for cycle completion.

	go func() {
		for _, ctx := range callingCtx {
			<-ctx.Done()
		}
		cancel(fmt.Errorf("canceled because all callers were done"))
	}()

	go func() {
		allM1cSent.Wait()
		m1c.cutoff()
	}()

	cancelProposalCollection := time.AfterFunc(s.proposalCollectionInterval(), func() {
		m1c.cutoff()
		for _, r := range registrations {
			r.cancelMmfs <- struct{}{}
		}
	})
	<-closedOnCycleEnd

	// Clean up in case it was never needed.
	cancelProposalCollection.Stop()
}

///////////////////////////////////////
///////////////////////////////////////

type mAndM6c struct {
	m   *pb.Match
	m7c chan string
}

// fanInFanOut routes evaluated matches back to it's source synchronize call.
// Each incoming match is passed along with it's synchronize call's m7c channel.
// This channel is remembered in a map, and the match is passed to be evaluated.
// When a match returns from evaluation, it's ID is looked up in the map and the
// match is returned on that channel.
func fanInFanOut(m2c <-chan mAndM6c, m3c chan<- *pb.Match, m6c <-chan string) {
	m6cMap := make(map[string]chan<- string)

	defer func(m2c <-chan mAndM6c) {
		for range m2c {
		}
	}(m2c)

	for {
		select {
		case m2, ok := <-m2c:
			if ok {
				m6cMap[m2.m.GetMatchId()] = m2.m7c
				m3c <- m2.m
			} else {
				close(m3c)
				// No longer select on m2c
				m2c = nil
			}

		case m5, ok := <-m6c:
			if !ok {
				return
			}

			m7c, ok := m6cMap[m5]
			if ok {
				m7c <- m5
			} else {
				logger.WithFields(logrus.Fields{
					"matchId": m5,
				}).Error("Match ID from evaluator does not match any id sent to it.")
			}
		}
	}
}

///////////////////////////////////////
///////////////////////////////////////

type cutoffSender struct {
	m1c       chan<- mAndM6c
	m2c       chan<- mAndM6c
	closed    chan struct{}
	closeOnce sync.Once
}

// cutoffSender allows values to be passed on the provided channel until cutoff
// has been called.  This closed the provided channel.  Calls to send after
// cutoff work, but values are ignored.
func newCutoffSender(m2c chan<- mAndM6c) *cutoffSender {
	m1c := make(chan mAndM6c)
	c := &cutoffSender{
		m1c:    m1c,
		m2c:    m2c,
		closed: make(chan struct{}),
	}

	go func() {
		defer close(m2c)

		for {
			select {
			case <-c.closed:
				return
			case match := <-m1c:
				m2c <- match
			}
		}
	}()
	return c
}

// send passes the value on the channel if still open, otherwise does nothing.
func (c *cutoffSender) send(match mAndM6c) {
	select {
	case <-c.closed:
	case c.m1c <- match:
	}
}

// cutoff closes m2c.  Safe to call from multiple go routines.
func (c *cutoffSender) cutoff() {
	c.closeOnce.Do(func() {
		close(c.closed)
	})
}

///////////////////////////////////////
///////////////////////////////////////

// Calls the evaluator with the matches.
func (s *synchronizerService) wrapEvaluator(ctx context.Context, cancel cancelErrFunc, m3c <-chan []*pb.Match, m5c chan<- string) {
	matchIDs, err := s.eval.evaluate(ctx, m3c)
	if err == nil {
		for _, mID := range matchIDs {
			m5c <- mID
		}
	} else {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("error calling evaluator, canceling cycle")
		cancel(fmt.Errorf("error calling evaluator: %w", err))
	}
	close(m5c)
}

///////////////////////////////////////
///////////////////////////////////////

func (s *synchronizerService) cacheMatchIDToTicketIDs(m *sync.Map, m3c <-chan *pb.Match, m4c chan<- *pb.Match) {
	for match := range m3c {
		m.Store(match.GetMatchId(), getTicketIds(match.GetTickets()))
		m4c <- match
	}
	close(m4c)
}

func getTicketIds(tickets []*pb.Ticket) []string {
	tids := []string{}
	for _, ticket := range tickets {
		tids = append(tids, ticket.GetId())
	}
	return tids
}

///////////////////////////////////////
///////////////////////////////////////

// Calls statestore to add all of the tickets returned by the evaluator to the
// ignorelist.  If it partially fails for whatever reason (not all tickets will
// nessisarily be in the same call), only the matches which can be safely
// returned to the Synchronize calls are.
func (s *synchronizerService) addMatchesToIgnoreList(ctx context.Context, m *sync.Map, cancel cancelErrFunc, m5c <-chan []string, m6c chan<- string) {
	totalMatches := 0
	successfulMatches := 0
	var lastErr error
	for mIDs := range m5c {
		ids := []string{}
		for _, mID := range mIDs {
			tids, ok := m.Load(mID)
			if ok {
				ids = append(ids, tids.([]string)...)
			} else {
				logger.Errorf("failed to get MatchId %s with its corresponding tickets from the cache", mID)
			}
		}

		err := s.store.AddTicketsToIgnoreList(ctx, ids)

		totalMatches += len(mIDs)
		if err == nil {
			successfulMatches += len(mIDs)
		} else {
			lastErr = err
		}

		for _, mID := range mIDs {
			m6c <- mID
		}
	}

	if lastErr != nil {
		logger.WithFields(logrus.Fields{
			"error":             lastErr.Error(),
			"totalMatches":      totalMatches,
			"successfulMatches": successfulMatches,
		}).Error("some or all matches were not successfully added to the ignore list, failed matches dropped")

		if successfulMatches == 0 {
			cancel(fmt.Errorf("no matches successfully added to the ignore list.  Last error: %w", lastErr))
		}
	}
	close(m6c)
}

///////////////////////////////////////
///////////////////////////////////////

func (s *synchronizerService) registrationInterval() time.Duration {
	const (
		name            = "synchronizer.registrationIntervalMs"
		defaultInterval = time.Second
	)

	if !s.cfg.IsSet(name) {
		return defaultInterval
	}

	return s.cfg.GetDuration(name)
}

func (s *synchronizerService) proposalCollectionInterval() time.Duration {
	const (
		name            = "synchronizer.proposalCollectionIntervalMs"
		defaultInterval = 10 * time.Second
	)

	if !s.cfg.IsSet(name) {
		return defaultInterval
	}

	return s.cfg.GetDuration(name)
}

///////////////////////////////////////
///////////////////////////////////////

// bufferMatchChannel collects matches from the input, and sends
// slice of matches on the output.  It never (for long) blocks
// the input channel, always appending to the slice which will
// next be used for output.  Used before external calls, so that
// network won't back up internal processing.
func bufferMatchChannel(in chan *pb.Match) chan []*pb.Match {
	out := make(chan []*pb.Match)
	go func() {
		var a []*pb.Match

	outerLoop:
		for {
			m, ok := <-in
			if !ok {
				break outerLoop
			}
			a = []*pb.Match{m}

			for len(a) > 0 {
				select {
				case m, ok := <-in:
					if !ok {
						break outerLoop
					}
					a = append(a, m)
				case out <- a:
					a = nil
				}
			}
		}
		if len(a) > 0 {
			out <- a
		}
		close(out)
	}()
	return out
}

// bufferStringChannel collects strings from the input, and sends
// slice of strings on the output.  It never (for long) blocks
// the input channel, always appending to the slice which will
// next be used for output.  Used before external calls, so that
// network won't back up internal processing.
func bufferStringChannel(in chan string) chan []string {
	out := make(chan []string)
	go func() {
		var a []string

	outerLoop:
		for {
			mID, ok := <-in
			if !ok {
				break outerLoop
			}
			a = []string{mID}

			for len(a) > 0 {
				select {
				case m, ok := <-in:
					if !ok {
						break outerLoop
					}
					a = append(a, m)
				case out <- a:
					a = nil
				}
			}
		}
		if len(a) > 0 {
			out <- a
		}
		close(out)
	}()
	return out
}

///////////////////////////////////////
///////////////////////////////////////

// withCancelCause returns a copy of parent with a new Done channel. The
// returned context's Done channel is closed when the returned cancel function
// is called or when the parent context's Done channel is closed, whichever
// happens first.  Unlike the conext package's WithCancel, the cancel func takes
// an error, and will return that error on subsequent calls to Err().
func withCancelCause(parent context.Context) (context.Context, cancelErrFunc) {
	parent, cancel := context.WithCancel(parent)

	ctx := &contextWithCancelCause{
		Context: parent,
	}

	return ctx, func(err error) {
		ctx.m.Lock()
		defer ctx.m.Unlock()

		if ctx.err == nil && parent.Err() == nil {
			ctx.err = err
		}
		cancel()
	}
}

type cancelErrFunc func(err error)

type contextWithCancelCause struct {
	context.Context
	m   sync.Mutex
	err error
}

func (ctx *contextWithCancelCause) Err() error {
	ctx.m.Lock()
	defer ctx.m.Unlock()
	if ctx.err == nil {
		return ctx.Context.Err()
	}
	return ctx.err
}
