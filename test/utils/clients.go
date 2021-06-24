// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"testing"

	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/stretchr/testify/assert"
)

// getSubClient returns an SDK subscription client
func GetSubAdminClient(t *testing.T) subapi.SubscriptionAdminServiceClient {
	conn, err := ConnectE2tServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	return subapi.NewSubscriptionAdminServiceClient(conn)
}

// getE2Client gets an E2 client
func GetE2Client2(t *testing.T, serviceModelName string, serviceModelVersion string, encoding sdkclient.Encoding) sdkclient.Client {
	client := sdkclient.NewClient(sdkclient.WithE2TAddress(E2TServiceHost, E2TServicePort),
		sdkclient.WithServiceModel(sdkclient.ServiceModelName(serviceModelName),
			sdkclient.ServiceModelVersion(serviceModelVersion)),
		sdkclient.WithEncoding(encoding))
	assert.NotNil(t, client)
	return client
}
