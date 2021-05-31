// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/app"
	"testing"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	e2v1beta1client "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
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

// GetE2V1beta1Client gets an E2 v1beta1 client
func GetE2V1beta1Client(t *testing.T, appID string, instanceID string, serviceModelName, serviceModelVersion string) e2v1beta1client.Client {
	return e2v1beta1client.NewClient(
		e2v1beta1client.WithAppID(e2v1beta1client.AppID(appID)),
		e2v1beta1client.WithInstanceID(e2v1beta1client.InstanceID(instanceID)),
		e2v1beta1client.WithE2THost(E2TServiceHost),
		e2v1beta1client.WithE2TPort(E2TServicePort),
		e2v1beta1client.WithServiceModel(
			e2v1beta1client.ServiceModelName(serviceModelName),
			e2v1beta1client.ServiceModelVersion(serviceModelVersion)))
}

// getSubClient returns an SDK subscription client
func GetSubClient(t *testing.T) sdksub.Client {
	conn, err := ConnectSubscriptionServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return sdksub.NewClient(conn)
}
