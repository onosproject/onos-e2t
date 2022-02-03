// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionCancel : tests cancelling the context for a subscription then subscribing again
func (s *TestSuite) TestSubscriptionCancel(t *testing.T) {
	const (
		simName    = "subscription-cancel"
		subName    = "TestSubscriptionCancelKpmV2"
		iterations = 10
	)
	var err error
	// make a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, simName)

	// make an SDK client and a subscription request
	nodeID := utils.GetTestNodeID(t)
	subSpec := utils.CreateKpmV2Sub(t, nodeID)
	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))

	for i := 1; i <= iterations; i++ {
		// Create the subscription
		subCtx, subCancel := context.WithTimeout(context.Background(), 15*time.Second)
		ch := make(chan v1beta1.Indication)
		_, err = node.Subscribe(subCtx, subName, subSpec, ch)
		assert.NoError(t, err)

		// Check that there is a message available
		_ = e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)

		// Cancel the subscription
		subCancel()
		_ = utils.ReadToEndOfChannel(ch)
	}

	unsubCtx, unsubCancel := context.WithTimeout(context.Background(), 15*time.Second)
	err = node.Unsubscribe(unsubCtx, subName)
	assert.NoError(t, err)
	unsubCancel()

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
