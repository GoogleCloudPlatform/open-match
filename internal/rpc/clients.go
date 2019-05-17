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

package rpc

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"open-match.dev/open-match/internal/config"
)

var (
	clientLogger = logrus.WithFields(logrus.Fields{
		"app":       "openmatch",
		"component": "client",
	})
)

// ClientParams contains the connection parameters to connect to an Open Match service.
type ClientParams struct {
	Hostname           string
	Port               int
	TrustedCertificate []byte
}

func (p *ClientParams) usingTLS() bool {
	return len(p.TrustedCertificate) > 0
}

// GRPCClientFromConfig creates a gRPC client connection from a configuration.
func GRPCClientFromConfig(cfg config.View, prefix string) (*grpc.ClientConn, error) {
	clientParams := &ClientParams{
		Hostname: cfg.GetString(prefix + ".hostname"),
		Port:     cfg.GetInt(prefix + ".grpcport"),
	}

	// If TLS support is enabled in the config, fill in the trusted certificates for decrpting server certificate.
	if cfg.GetBool("tls.enabled") {
		_, err := os.Stat(cfg.GetString("tls.trustedCertificatesPath"))
		if err != nil {
			clientLogger.WithError(err).Error("trusted certificate file may not exists.")
			return nil, err
		}

		clientParams.TrustedCertificate, err = ioutil.ReadFile(cfg.GetString("tls.trustedCertificatesPath"))
		if err != nil {
			clientLogger.WithError(err).Error("failed to read tls trusted certificate to establish a secure grpc client.")
			return nil, err
		}
	}

	return GRPCClientFromParams(clientParams)
}

// GRPCClientFromParams creates a gRPC client connection from the parameters.
func GRPCClientFromParams(params *ClientParams) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", params.Hostname, params.Port)

	grpcOptions := []grpc.DialOption{}

	if params.usingTLS() {
		creds, err := clientCredentialsFromFileData(params.TrustedCertificate, "")
		if err != nil {
			clientLogger.WithError(err).Error("failed to get transport credentials from file.")
			return nil, errors.WithStack(err)
		}
		grpcOptions = append(grpcOptions, grpc.WithTransportCredentials(creds))
	} else {
		grpcOptions = append(grpcOptions, grpc.WithInsecure())
	}

	return grpc.Dial(address, grpcOptions...)
}

// HTTPClientFromConfig creates a HTTP client from from a configuration.
func HTTPClientFromConfig(cfg config.View, prefix string) (*http.Client, string, error) {
	clientParams := &ClientParams{
		Hostname: cfg.GetString(prefix + ".hostname"),
		Port:     cfg.GetInt(prefix + ".httpport"),
	}

	// If TLS support is enabled in the config, fill in the trusted certificates for decrpting server certificate.
	if cfg.GetBool("tls.enabled") {
		_, err := os.Stat(cfg.GetString("tls.trustedCertificatesPath"))
		if err != nil {
			clientLogger.WithError(err).Error("trusted certificate file may not exists.")
			return nil, "", err
		}

		clientParams.TrustedCertificate, err = ioutil.ReadFile(cfg.GetString("tls.trustedCertificatesPath"))
		if err != nil {
			clientLogger.WithError(err).Error("failed to read tls trusted certificate to establish a secure grpc client.")
			return nil, "", err
		}
	}

	return HTTPClientFromParams(clientParams)
}

// HTTPClientFromParams creates a HTTP client from the parameters.
func HTTPClientFromParams(params *ClientParams) (*http.Client, string, error) {
	address := fmt.Sprintf("%s:%d", params.Hostname, params.Port)
	// Make client Timeout configurable
	httpClient := &http.Client{Timeout: time.Second * 3}
	var baseURL string

	if params.usingTLS() {
		baseURL = "https://" + address

		pool, err := trustedCertificates(params.TrustedCertificate)
		if err != nil {
			clientLogger.WithError(err).Error("failed to get cert pool from file.")
			return nil, "", err
		}

		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				ServerName: address,
				RootCAs:    pool,
			},
		}
	} else {
		baseURL = "http://" + address
	}

	return httpClient, baseURL, nil
}
