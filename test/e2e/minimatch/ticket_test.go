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

package minimatch

import (
	"context"
	"io"
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"open-match.dev/open-match/pkg/pb"

	rpcTesting "open-match.dev/open-match/internal/rpc/testing"
)

func TestAssignTickets(t *testing.T) {
	tc := createBackendForTest(t)
	defer tc.Close()

	fe := pb.NewFrontendClient(tc.MustGRPC())
	be := pb.NewBackendClient(tc.MustGRPC())

	ctResp, err := fe.CreateTicket(tc.Context(), &pb.CreateTicketRequest{Ticket: &pb.Ticket{}})
	assert.Nil(t, err)
	fakeAssignment := &pb.Assignment{Connection: "localhost"}

	var tt = []struct {
		description    string
		req            *pb.AssignTicketsRequest
		wantResp       *pb.AssignTicketsResponse
		wantAssignment *pb.Assignment
		wantCode       codes.Code
	}{
		{
			"expects invalid argument code since request is empty",
			&pb.AssignTicketsRequest{},
			nil,
			nil,
			codes.InvalidArgument,
		},
		{
			"expects invalid argument code since assignment is nil",
			&pb.AssignTicketsRequest{
				TicketId: []string{"1"},
			},
			nil,
			nil,
			codes.InvalidArgument,
		},
		{
			"expects not found code since ticket id does not exist in the statestore",
			&pb.AssignTicketsRequest{
				TicketId:   []string{"2"},
				Assignment: fakeAssignment,
			},
			nil,
			nil,
			codes.NotFound,
		},
		{
			"expects not found code since ticket id 'unknown id' does not exist in the statestore",
			&pb.AssignTicketsRequest{
				TicketId:   []string{ctResp.GetTicket().GetId(), "unknown id"},
				Assignment: fakeAssignment,
			},
			nil,
			nil,
			codes.NotFound,
		},
		{
			"expects ok code",
			&pb.AssignTicketsRequest{
				TicketId:   []string{ctResp.GetTicket().GetId()},
				Assignment: fakeAssignment,
			},
			&pb.AssignTicketsResponse{},
			fakeAssignment,
			codes.OK,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			resp, err := be.AssignTickets(tc.Context(), test.req)
			assert.Equal(t, test.wantResp, resp)
			assert.Equal(t, test.wantCode, status.Convert(err).Code())
			gtResp, err := fe.GetTicket(tc.Context(), &pb.GetTicketRequest{TicketId: ctResp.GetTicket().GetId()})
			assert.Nil(t, err)

			// grpc will write something to the reserved fields of this protobuf object, so we have to do comparisions fields by fields.
			assert.Equal(t, test.wantAssignment.GetConnection(), gtResp.GetAssignment().GetConnection())
			assert.Equal(t, test.wantAssignment.GetProperties(), gtResp.GetAssignment().GetProperties())
			assert.Equal(t, test.wantAssignment.GetError(), gtResp.GetAssignment().GetError())

		})
	}
}

// TODO: rewrite with table testing format and add more corner test cases
// TestFrontendService tests creating, getting and deleting a ticket using Frontend service.
func TestTicketLifeCycle(t *testing.T) {
	assert := assert.New(t)

	tc := createStore(t)
	fe := pb.NewFrontendClient(tc.MustGRPC())
	assert.NotNil(fe)

	ticket := &pb.Ticket{
		Properties: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"test-property": {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
			},
		},
		Assignment: &pb.Assignment{
			Connection: "test-tbd",
		},
	}

	// Create a ticket, validate that it got an id and set its id in the expected ticket.
	resp, err := fe.CreateTicket(context.Background(), &pb.CreateTicketRequest{Ticket: ticket})
	assert.NotNil(resp)
	assert.Nil(err)
	want := resp.Ticket
	assert.NotNil(want)
	assert.NotNil(want.Id)
	ticket.Id = want.Id
	validateTicket(t, resp.Ticket, ticket)

	// Fetch the ticket and validate that it is identical to the expected ticket.
	gotTicket, err := fe.GetTicket(context.Background(), &pb.GetTicketRequest{TicketId: ticket.Id})
	assert.NotNil(gotTicket)
	assert.Nil(err)
	validateTicket(t, gotTicket, ticket)

	// Delete the ticket and validate that it was actually deleted.
	_, err = fe.DeleteTicket(context.Background(), &pb.DeleteTicketRequest{TicketId: ticket.Id})
	assert.Nil(err)
	validateDelete(t, fe, ticket.Id)
}

func TestQueryTickets(t *testing.T) {
	assert := assert.New(t)
	tc := createMmlogicForTest(t)
	defer tc.Close()

	queryTicketsLoop(t, tc, &pb.QueryTicketsRequest{}, func(_ *pb.QueryTicketsResponse, err error) {
		assert.Equal(codes.InvalidArgument, status.Convert(err).Code())
	})
}

func TestQueryTicketsForEmptyDatabase(t *testing.T) {
	assert := assert.New(t)
	tc := createMmlogicForTest(t)
	defer tc.Close()

	queryTicketsLoop(t, tc,
		&pb.QueryTicketsRequest{
			Pool: &pb.Pool{
				Filter: []*pb.Filter{{
					Attribute: "ok",
				}},
			},
		},
		func(resp *pb.QueryTicketsResponse, err error) {
			assert.NotNil(resp)
		})
}

func queryTicketsLoop(t *testing.T, tc *rpcTesting.TestContext, req *pb.QueryTicketsRequest, handleResponse func(*pb.QueryTicketsResponse, error)) {
	c := pb.NewMmLogicClient(tc.MustGRPC())
	stream, err := c.QueryTickets(tc.Context(), req)
	if err != nil {
		t.Fatalf("error querying tickets, %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		handleResponse(resp, err)
		if err != nil {
			return
		}
	}
}
