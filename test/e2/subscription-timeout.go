// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/test/utils"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSubscription
func (s *TestSuite) TestSubscriptionTimeout(t *testing.T) {
	sim := utils.CreateRanSimulatorWithName(t, "ran-simulator")
	assert.NotNil(t, sim)

	clientConfig := e2client.Config{
		AppID: "subscription-test",
		SubscriptionService: e2client.ServiceConfig{
			Host: SubscriptionServiceHost,
			Port: SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	subReq, err := createSubscriptionRequest(nodeIDs[0])
	assert.NoError(t, err)

	err = sim.Uninstall()
	assert.NoError(t, err)

	_, err = client.Subscribe(ctx, subReq, ch)
	assert.Error(t, err)
	// TODO - check that the proper error is returned
}
