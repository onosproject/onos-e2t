// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"testing"
	"time"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestTransactionTimeout tests that channels and subscriptions are removed if their subscription timeout period expires
func (s *TestSuite) TestTransactionTimeout(t *testing.T) {
	// spin up a ransim instance
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "transaction-timeout")
	assert.NotNil(t, sim)

	// create a KPM V2 subscription
	const baseTimeout = 10 * time.Second
	nodeID := utils.GetTestNodeID(t)
	subSpec := utils.CreateKpmV2Sub(t, nodeID)
	subName := "TestTransactionTimeout"
	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan v1beta1.Indication)
	ctx, cancel := e2utils.GetCtx()
	_, err := node.Subscribe(ctx, subName, subSpec, ch, sdkclient.WithTransactionTimeout(baseTimeout))
	assert.NoError(t, err)

	// make sure the subscription channel is working by reading an indication from it
	indication := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	assert.NotNil(t, indication)

	// check that the number of subscriptions and channels is now 1
	assert.Equal(t, 1, len(e2utils.GetSubscriptionList(t)))
	assert.Equal(t, 1, len(e2utils.GetChannelList(t)))

	// Cause the subscription to time out and wait for it to happen
	cancel()
	time.Sleep(baseTimeout + (2 * time.Second))

	// Make sure that the subscription and the channel were removed
	e2utils.CheckForEmptySubscriptionList(t)
	assert.Equal(t, 0, len(e2utils.GetChannelList(t)))

	// clean up the simulator
	utils.UninstallRanSimulatorOrDie(t, sim)
}
