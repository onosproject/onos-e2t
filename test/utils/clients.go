// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"fmt"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/stretchr/testify/assert"
)

// GetSubAdminClient returns an SDK subscription client
func GetSubAdminClient(t *testing.T) subapi.SubscriptionAdminServiceClient {
	conn, err := ConnectE2tServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return subapi.NewSubscriptionAdminServiceClient(conn)
}

// GetSubClientForIP returns an SDK subscription client
func GetSubClientForIP(t *testing.T, IP string, port uint32) subapi.SubscriptionServiceClient {
	conn, err := ConnectE2t(IP, port)
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return subapi.NewSubscriptionServiceClient(conn)
}

// GetE2Client gets an E2 client
func GetE2Client(t *testing.T, serviceModelName string, serviceModelVersion string, encoding sdkclient.Encoding) sdkclient.Client {
	client := sdkclient.NewClient(sdkclient.WithE2TAddress(E2TServiceHost, E2TServicePort),
		sdkclient.WithServiceModel(sdkclient.ServiceModelName(serviceModelName),
			sdkclient.ServiceModelVersion(serviceModelVersion)),
		sdkclient.WithEncoding(encoding))
	assert.NotNil(t, client)
	return client
}

// ConnectE2tServiceHost connects to subscription service via service name
func ConnectE2tServiceHost() (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	return grpc.DialContext(context.Background(), E2tServiceAddress, opts...)
}

// ConnectE2t connects to subscription service via IP/port
func ConnectE2t(IP string, port uint32) (*grpc.ClientConn, error) {
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	addr := fmt.Sprintf("%s:%d", IP, port)
	return grpc.DialContext(context.Background(), addr, opts...)
}

// ReadToEndOfChannel reads messages from a channel until an error occurs, clearing the
// channel of messages
func ReadToEndOfChannel(ch chan e2api.Indication) bool {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return true
			}
		case <-time.After(2 * time.Minute):
			return false
		}
	}
}
