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

package serving

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"open-match.dev/open-match/internal/future/pb"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	shellTesting "open-match.dev/open-match/internal/future/testing"
	netlistenerTesting "open-match.dev/open-match/internal/util/netlistener/testing"
)

func TestInsecureStartStop(t *testing.T) {
	assert := assert.New(t)
	grpcLh := netlistenerTesting.MustListen()
	httpLh := netlistenerTesting.MustListen()
	ff := &shellTesting.FakeFrontend{}

	params := NewParamsFromListeners(grpcLh, httpLh)
	params.AddHandleFunc(func(s *grpc.Server) {
		pb.RegisterFrontendServer(s, ff)
	}, pb.RegisterFrontendHandlerFromEndpoint)
	s := newInsecureServer(grpcLh, httpLh)
	defer s.stop()
	waitForStart, err := s.start(params)
	assert.Nil(err)
	waitForStart()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", grpcLh.Number()), grpc.WithInsecure())
	defer conn.Close()
	assert.Nil(err)

	endpoint := fmt.Sprintf("http://localhost:%d", httpLh.Number())
	httpClient := &http.Client{
		Timeout: time.Second,
	}
	runGrpcWithProxyTests(assert, s, conn, httpClient, endpoint)
}

func runGrpcWithProxyTests(assert *assert.Assertions, s grpcServerWithProxy, conn *grpc.ClientConn, httpClient *http.Client, endpoint string) {
	ctx := context.Background()
	feClient := pb.NewFrontendClient(conn)
	grpcResp, err := feClient.CreateTicket(ctx, &pb.CreateTicketRequest{})
	assert.Nil(err)
	assert.NotNil(grpcResp)

	httpReq, err := http.NewRequest(http.MethodPost, endpoint+"/v1/frontend/tickets", strings.NewReader("{}"))
	assert.Nil(err)
	assert.NotNil(httpReq)
	httpResp, err := httpClient.Do(httpReq)
	assert.Nil(err)
	assert.NotNil(httpResp)
	defer func() {
		if httpResp != nil {
			httpResp.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(httpResp.Body)
	assert.Nil(err)
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal("{}", string(body))

	httpReq, err = http.NewRequest(http.MethodGet, endpoint+"/healthz", nil)
	assert.Nil(err)

	httpResp, err = httpClient.Do(httpReq)
	assert.Nil(err)
	assert.NotNil(httpResp)
	body, err = ioutil.ReadAll(httpResp.Body)
	assert.Nil(err)
	assert.Equal(200, httpResp.StatusCode)
	assert.Equal("ok", string(body))

	s.stop()
}
