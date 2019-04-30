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
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/open-match/internal/pb"

	shellTesting "github.com/GoogleCloudPlatform/open-match/internal/future/testing"
	netlistenerTesting "github.com/GoogleCloudPlatform/open-match/internal/util/netlistener/testing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestStartStopServer(t *testing.T) {
	assert := assert.New(t)
	grpcLh := netlistenerTesting.MustListen()
	httpLh := netlistenerTesting.MustListen()
	ff := shellTesting.NewFakeFrontend()

	params := NewParamsFromListeners(grpcLh, httpLh)
	params.AddHandleFunc(func(s *grpc.Server) {
		pb.RegisterFrontendServer(s, ff)
	}, pb.RegisterFrontendHandlerFromEndpoint)
	s := New()
	defer s.Stop()

	waitForStart, err := s.Start(params)
	assert.Nil(err)
	waitForStart()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", grpcLh.Number()), grpc.WithInsecure())
	assert.Nil(err)

	endpoint := fmt.Sprintf("http://localhost:%d", httpLh.Number())
	httpClient := &http.Client{
		Timeout: time.Second,
	}

	runGrpcWithProxyTests(assert, s.serverWithProxy, conn, httpClient, endpoint)
}

func TestMustServeForever(t *testing.T) {
	assert := assert.New(t)
	grpcLh := netlistenerTesting.MustListen()
	httpLh := netlistenerTesting.MustListen()
	ff := shellTesting.NewFakeFrontend()

	params := NewParamsFromListeners(grpcLh, httpLh)
	params.AddHandleFunc(func(s *grpc.Server) {
		pb.RegisterFrontendServer(s, ff)
	}, pb.RegisterFrontendHandlerFromEndpoint)
	serveUntilKilledFunc, stopServingFunc, err := startServingIndefinitely(params)
	assert.Nil(err)
	go func() {
		// Wait for 500ms before killing the server.
		// It really doesn't matter if it actually comes up.
		// We just care that the server can respect an unexpected shutdown quickly after starting.
		time.Sleep(time.Millisecond * 500)
		stopServingFunc()
	}()
	serveUntilKilledFunc()
	// This test will intentionally deadlock if the stop function is not respected.
}
