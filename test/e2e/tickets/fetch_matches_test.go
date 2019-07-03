// +build !e2ecluster
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

package tickets

import (
	"testing"
	"time"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	statestoreTesting "open-match.dev/open-match/internal/statestore/testing"
	"open-match.dev/open-match/internal/testing/e2e"
	"open-match.dev/open-match/pkg/pb"
)

func TestGameMatchWorkFlow(t *testing.T) {
	var gotFmResp *pb.FetchMatchesResponse
	var gotAtResp *pb.AssignTicketsResponse
	var ctResp *pb.CreateTicketResponse
	var gotDtResp *pb.DeleteTicketResponse
	var err error
	/*
		This end to end test does the following things step by step
		1. Create a few tickets with delicate designs and hand crafted properties
		2. Call backend.FetchMatches and verify it returns expected matches.
		3. Call backend.FetchMatches within redis.ignoreLists.ttl seconds and expect FailedPrecondition error.
		4. Wait for redis.ignoreLists.ttl seconds and call backend.FetchMatches the third time, expect the same result as step 2.
		5. Call backend.AssignTickets to assign DGSs for the tickets in FetchMatches' response
		6. Call backend.FetchMatches and verify it no longer returns tickets got assigned in the previous step.
		7. Call frontend.DeleteTicket to delete part of the tickets that got assigned.
		8. Call backend.AssignTickets to the tickets got deleted and expect an error.
	*/

	om, closer := e2e.New(t)
	defer closer()
	fe := om.MustFrontendGRPC()
	be := om.MustBackendGRPC()
	mmfCfg := om.MustMmfConfigGRPC()

	ticket1 := &pb.Ticket{
		Properties: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"level":   {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
				"defense": {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
			},
		},
	}

	ticket2 := &pb.Ticket{
		Properties: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"level":   {Kind: &structpb.Value_NumberValue{NumberValue: 5}},
				"defense": {Kind: &structpb.Value_NumberValue{NumberValue: 5}},
			},
		},
	}

	ticket3 := &pb.Ticket{
		Properties: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"level":   {Kind: &structpb.Value_NumberValue{NumberValue: 10}},
				"defense": {Kind: &structpb.Value_NumberValue{NumberValue: 10}},
			},
		},
	}

	ticket4 := &pb.Ticket{
		Properties: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"level":   {Kind: &structpb.Value_NumberValue{NumberValue: 15}},
				"defense": {Kind: &structpb.Value_NumberValue{NumberValue: 15}},
			},
		},
	}

	ticket5 := &pb.Ticket{
		Properties: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"level":   {Kind: &structpb.Value_NumberValue{NumberValue: 20}},
				"defense": {Kind: &structpb.Value_NumberValue{NumberValue: 20}},
			},
		},
	}

	tickets := []*pb.Ticket{ticket1, ticket2, ticket3, ticket4, ticket5}

	// 1. Create a few tickets with delicate designs and hand crafted properties
	for i := 0; i < len(tickets); i++ {
		ctResp, err = fe.CreateTicket(om.Context(), &pb.CreateTicketRequest{Ticket: tickets[i]})
		assert.Nil(t, err)
		assert.NotNil(t, ctResp)
		// Assign Open Match ids back to the input tickets
		*tickets[i] = *ctResp.GetTicket()
	}

	fmReq := &pb.FetchMatchesRequest{
		Config: mmfCfg,
		Profile: []*pb.MatchProfile{
			{
				Name: "test-profile",
				Pool: []*pb.Pool{
					{
						Name:   "ticket12",
						Filter: []*pb.Filter{{Attribute: "level", Min: 0, Max: 6}, {Attribute: "defense", Min: 0, Max: 100}},
					},
					{
						Name:   "ticket23",
						Filter: []*pb.Filter{{Attribute: "level", Min: 3, Max: 14}, {Attribute: "defense", Min: 0, Max: 100}},
					},
					{
						Name:   "ticket5",
						Filter: []*pb.Filter{{Attribute: "level", Min: 0, Max: 100}, {Attribute: "defense", Min: 17, Max: 25}},
					},
					{
						Name:   "ticket234",
						Filter: []*pb.Filter{{Attribute: "level", Min: 3, Max: 17}, {Attribute: "defense", Min: 3, Max: 17}},
					},
				},
			},
		},
	}

	// 2. Call backend.FetchMatches and expects two matches with the following tickets
	gotFmResp, err = be.FetchMatches(om.Context(), fmReq)
	tmpMatches := gotFmResp.GetMatch()
	assert.Nil(t, err)
	validateFetchMatchesResponse(t, [][]*pb.Ticket{{ticket2, ticket3, ticket4}, {ticket5}}, gotFmResp)

	// 3. Call backend.FetchMatches within redis.ignoreLists.ttl seconds and expects it return a match with ticket1 .
	gotFmResp, err = be.FetchMatches(om.Context(), fmReq)
	assert.Nil(t, err)
	validateFetchMatchesResponse(t, [][]*pb.Ticket{{ticket1}}, gotFmResp)

	// 4. Wait for redis.ignoreLists.ttl seconds and call backend.FetchMatches the third time, expect the same result as step 2.
	time.Sleep(statestoreTesting.IgnoreListTTL)
	gotFmResp, err = be.FetchMatches(om.Context(), fmReq)
	assert.Nil(t, err)
	validateFetchMatchesResponse(t, [][]*pb.Ticket{{ticket2, ticket3, ticket4}, {ticket5}}, gotFmResp)

	// 5. Call backend.AssignTickets to assign DGSs for the tickets in FetchMatches' response
	for _, match := range tmpMatches {
		tids := []string{}
		for _, ticket := range match.GetTicket() {
			tids = append(tids, ticket.GetId())
		}
		gotAtResp, err = be.AssignTickets(om.Context(), &pb.AssignTicketsRequest{TicketId: tids, Assignment: &pb.Assignment{Connection: "agones-1"}})
		assert.Nil(t, err)
		assert.NotNil(t, gotAtResp)
	}

	// 6. Call backend.FetchMatches and verify it no longer returns tickets got assigned in the previous step.
	time.Sleep(statestoreTesting.IgnoreListTTL)
	gotFmResp, err = be.FetchMatches(om.Context(), fmReq)
	assert.Nil(t, err)
	validateFetchMatchesResponse(t, [][]*pb.Ticket{{ticket1}}, gotFmResp)

	// 7. Call frontend.DeleteTicket to delete part of the tickets that got assigned.
	gotDtResp, err = fe.DeleteTicket(om.Context(), &pb.DeleteTicketRequest{TicketId: ticket2.GetId()})
	assert.Nil(t, err)
	assert.NotNil(t, gotDtResp)

	// 8. Call backend.AssignTickets to the tickets got deleted and expect an error.
	gotAtResp, err = be.AssignTickets(om.Context(), &pb.AssignTicketsRequest{TicketId: []string{ticket2.GetId(), ticket3.GetId()}, Assignment: &pb.Assignment{Connection: "agones-2"}})
	assert.Equal(t, codes.NotFound, status.Convert(err).Code())
	assert.Nil(t, gotAtResp)
}

func validateFetchMatchesResponse(t *testing.T, wantTickets [][]*pb.Ticket, resp *pb.FetchMatchesResponse) {
	assert.NotNil(t, resp)
	assert.Equal(t, len(wantTickets), len(resp.GetMatch()))
	for _, match := range resp.GetMatch() {
		assert.Contains(t, wantTickets, match.GetTicket())
	}
}
