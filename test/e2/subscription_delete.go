// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	subscriptionName    = "TestSubscriptionDelete-kpm"
	subscriptionTimeout = 2 * time.Minute
)

// createAndVerifySubscription creates a subscription to the given node and makes sure that
// at least one verification message can be received from it. The channel ID of the subscription
// is returned
func createAndVerifySubscription(ctx context.Context, t *testing.T, nodeID topo.ID) (e2api.ChannelID, e2utils.KPMV2Sub) {
	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subscriptionName,
			NodeID: nodeID,
		},
		CellObjectID: e2utils.GetFirstCellObjectID(t, nodeID),
	}
	channelID, err := kpmv2Sub.Subscribe(ctx)
	assert.NoError(t, err)

	select {
	case indicationMsg := <-kpmv2Sub.Sub.Ch:
		t.Log(indicationMsg)
		assert.NotNil(t, indicationMsg)

	case <-time.After(10 * time.Second):
		assert.Equal(t, false, "test is failed because of timeout")

	}
	return channelID, kpmv2Sub
}

func getSubscriptionID(t *testing.T, channelID e2api.ChannelID) e2api.SubscriptionID {
	getChannelRequest := &e2api.GetChannelRequest{ChannelID: channelID}
	channelResponse, err := utils.GetSubAdminClient(t).GetChannel(context.Background(), getChannelRequest)
	assert.NoError(t, err)
	channel := channelResponse.Channel
	return channel.GetSubscriptionID()
}

// TestSubscriptionDelete tests subscription delete procedure
func (s *TestSuite) TestSubscriptionDelete(t *testing.T) {

	// Start up a ran-sim instance
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-delete")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()
	//  Initially the subscription list should be empty
	e2utils.CheckForEmptySubscriptionList(t)

	// Create a Node
	nodeID := utils.GetTestNodeID(t)

	// Add a subscription
	channelID, sub := createAndVerifySubscription(ctx, t, nodeID)
	subscriptionID := getSubscriptionID(t, channelID)

	// Check that the subscription list is correct
	subList := e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subscriptionID)

	// Close the subscription
	err := sub.Unsubscribe(ctx)
	assert.NoError(t, err)

	// Check number of subscriptions is correct after deleting the subscription
	e2utils.CheckForEmptySubscriptionList(t)

	//  Open the subscription again and make sure it is open
	channelID, sub = createAndVerifySubscription(ctx, t, nodeID)
	subscriptionID = getSubscriptionID(t, channelID)

	// Check that the number of subscriptions is correct after reopening
	subList = e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subscriptionID)

	// Close the subscription
	err = sub.Unsubscribe(ctx)
	assert.NoError(t, err)

	assert.True(t, utils.ReadToEndOfChannel(sub.Sub.Ch))

	e2utils.CheckForEmptySubscriptionList(t)

	// Clean up the ran-sim instance
	utils.UninstallRanSimulatorOrDie(t, sim)
}
