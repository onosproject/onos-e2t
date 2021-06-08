// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/app"
	"testing"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/stretchr/testify/assert"
)

// getE2Client gets an E2 client
func GetE2Client(t *testing.T, appID string) e2client.Client {
	clientConfig := e2client.Config{
		AppID: app.ID(appID),
		E2TService: e2client.ServiceConfig{
			Host: E2TServiceHost,
			Port: E2TServicePort,
		},
		SubscriptionService: e2client.ServiceConfig{
			Host: SubscriptionServiceHost,
			Port: SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	return client

}

// getSubClient returns an SDK subscription client
func GetSubClient(t *testing.T) sdksub.Client {
	conn, err := ConnectSubscriptionServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return sdksub.NewClient(conn)
}
