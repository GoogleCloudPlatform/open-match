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
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"open-match.dev/open-match/internal/config"
	"open-match.dev/open-match/internal/util/netlistener"
)

var (
	serverLogger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "server",
	})
)

// GrpcHandler binds gRPC services.
type GrpcHandler func(*grpc.Server)

// GrpcProxyHandler binds HTTP handler to gRPC service.
type GrpcProxyHandler func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

// Params holds all the parameters required to start a gRPC server.
type Params struct {
	// ServeMux is the router for the HTTP server. You can use this to serve pages in addition to the HTTP proxy.
	// Do NOT register "/" handler because it's reserved for the proxy.
	ServeMux             *http.ServeMux
	handlersForGrpc      []GrpcHandler
	handlersForGrpcProxy []GrpcProxyHandler
	grpcListener         *netlistener.ListenerHolder
	grpcProxyListener    *netlistener.ListenerHolder

	// Root CA public certificate in PEM format.
	rootCaPublicCertificateFileData []byte
	// Public certificate in PEM format.
	// If this field is the same as rootCaPublicCertificateFileData then the certificate is not backed by a CA.
	publicCertificateFileData []byte
	// Private key in PEM format.
	privateKeyFileData []byte
}

// NewParamsFromConfig returns server Params initialized from the configuration file.
func NewParamsFromConfig(cfg config.View, prefix string) (*Params, error) {
	grpcLh, err := netlistener.NewFromPortNumber(cfg.GetInt(prefix + ".port"))
	if err != nil {
		serverLogger.Fatal(err)
		return nil, err
	}
	httpLh, err := netlistener.NewFromPortNumber(cfg.GetInt(prefix + ".proxyport"))
	if err != nil {
		closeErr := grpcLh.Close()
		if closeErr != nil {
			serverLogger.WithFields(logrus.Fields{
				"error": closeErr.Error(),
			}).Info("failed to gRPC close port")
		}
		serverLogger.Fatal(err)
		return nil, err
	}
	p := NewParamsFromListeners(grpcLh, httpLh)

	certFile := cfg.GetString("api.tls.certificatefile")
	privateKeyFile := cfg.GetString("api.tls.privatekey")
	if len(certFile) > 0 && len(privateKeyFile) > 0 {
		serverLogger.Printf("Loading TLS certificate from, %s", certFile)
		publicCertData, err := ioutil.ReadFile(certFile)
		if err != nil {
			p.invalidate()
			return nil, errors.WithStack(err)
		}
		privateKeyData, err := ioutil.ReadFile(privateKeyFile)
		if err != nil {
			p.invalidate()
			return nil, errors.WithStack(err)
		}
		p.SetTLSConfiguration(publicCertData, publicCertData, privateKeyData)
	}
	return p, nil
}

// NewParamsFromListeners returns server Params initialized with the ListenerHolder variables.
func NewParamsFromListeners(grpcLh *netlistener.ListenerHolder, proxyLh *netlistener.ListenerHolder) *Params {
	return &Params{
		ServeMux:             http.NewServeMux(),
		handlersForGrpc:      []GrpcHandler{},
		handlersForGrpcProxy: []GrpcProxyHandler{},
		grpcListener:         grpcLh,
		grpcProxyListener:    proxyLh,
	}
}

// SetTLSConfiguration configures the server to run in TLS mode.
func (p *Params) SetTLSConfiguration(rootCaPublicCertificateFileData []byte, publicCertificateFileData []byte, privateKeyFileData []byte) *Params {
	p.rootCaPublicCertificateFileData = rootCaPublicCertificateFileData
	if len(p.rootCaPublicCertificateFileData) == 0 {
		p.rootCaPublicCertificateFileData = publicCertificateFileData
	}
	p.publicCertificateFileData = publicCertificateFileData
	p.privateKeyFileData = privateKeyFileData
	return p
}

// usingTLS returns true if a certificate is set.
func (p *Params) usingTLS() bool {
	return len(p.publicCertificateFileData) > 0
}

// AddHandleFunc binds gRPC service handler and an associated HTTP proxy handler.
func (p *Params) AddHandleFunc(handlerFunc GrpcHandler, grpcProxyHandler GrpcProxyHandler) {
	if handlerFunc != nil {
		p.handlersForGrpc = append(p.handlersForGrpc, handlerFunc)
	}
	if grpcProxyHandler != nil {
		p.handlersForGrpcProxy = append(p.handlersForGrpcProxy, grpcProxyHandler)
	}
}

// invalidate closes all the TCP listeners that would otherwise leak if initialization fails.
func (p *Params) invalidate() {
	if err := p.grpcListener.Close(); err != nil {
		serverLogger.Errorf("error closing grpc handler, %s", err)
	}
	if err := p.grpcProxyListener.Close(); err != nil {
		serverLogger.Errorf("error closing grpc-proxy handler, %s", err)
	}
}

// Server hosts a gRPC and HTTP server.
// All HTTP traffic is served from a common http.ServeMux.
type Server struct {
	serverWithProxy grpcServerWithProxy
}

// grpcServerWithProxy this will go away when insecure.go and tls.go are merged into the same server.
type grpcServerWithProxy interface {
	start(*Params) (func(), error)
	stop()
}

// Start the gRPC+HTTP(s) REST server.
func (s *Server) Start(p *Params) (func(), error) {
	if p.usingTLS() {
		s.serverWithProxy = newTLSServer(p.grpcListener, p.grpcProxyListener)
	} else {
		s.serverWithProxy = newInsecureServer(p.grpcListener, p.grpcProxyListener)
	}
	return s.serverWithProxy.start(p)
}

// Stop the gRPC+HTTP(s) REST server.
func (s *Server) Stop() {
	s.serverWithProxy.stop()
}

// startServingIndefinitely creates a server based on the params and begins serving the gRPC and HTTP proxy.
// It returns waitUntilKilled() which will wait indefinitely until crash or Ctrl+C is pressed.
// forceStopServingFunc() is also returned which is used to force kill the server for tests.
func startServingIndefinitely(params *Params) (func(), func(), error) {
	s := &Server{}

	// Start serving traffic.
	waitForStart, err := s.Start(params)
	if err != nil {
		serverLogger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Failed to start gRPC and HTTP servers.")
		return func() {}, func() {}, err
	}
	serverLogger.Info("Server has started.")
	// Exit when we see a signal
	waitUntilKilled, forceStopServingFunc := signal.New()

	waitForStart()
	serveUntilKilledFunc := func() {
		waitUntilKilled()
		s.Stop()
		serverLogger.Info("Shutting down server")
	}
	return serveUntilKilledFunc, forceStopServingFunc, nil
}

// MustServeForever is a convenience method for starting a server and running it indefinitely.
func MustServeForever(params *Params) {
	serveUntilKilledFunc, _, err := startServingIndefinitely(params)
	if err != nil {
		return
	}
	serveUntilKilledFunc()
}
