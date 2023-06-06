// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"time"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	subscriptionName    = "TestSubscriptionDelete-kpm"
	subscriptionTimeout = 2 * time.Minute
)

// createAndVerifySubscription creates a subscription to the given node and makes sure that
// at least one verification message can be received from it. The channel ID of the subscription
// is returned
func (s *TestSuite) createAndVerifySubscription(ctx context.Context, nodeID topo.ID) (e2api.ChannelID, e2utils.KPMV2Sub) {
	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subscriptionName,
			NodeID: nodeID,
		},
		CellObjectID: e2utils.GetFirstCellObjectID(s.T(), nodeID),
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())
	channelID, err := kpmv2Sub.Subscribe(ctx)
	s.NoError(err)

	select {
	case indicationMsg := <-kpmv2Sub.Sub.Ch:
		s.T().Log(indicationMsg)
		s.NotNil(indicationMsg)

	case <-time.After(10 * time.Second):
		s.Equal(false, "test is failed because of timeout")

	}
	return channelID, kpmv2Sub
}

func (s *TestSuite) getSubscriptionID(channelID e2api.ChannelID) e2api.SubscriptionID {
	getChannelRequest := &e2api.GetChannelRequest{ChannelID: channelID}
	channelResponse, err := utils.GetSubAdminClient(s.T()).GetChannel(s.Context(), getChannelRequest)
	s.NoError(err)
	channel := channelResponse.Channel
	return channel.GetSubscriptionID()
}

// TestSubscriptionDelete tests subscription delete procedure
func (s *TestSuite) TestSubscriptionDelete() {
	t := s.T()

	// Start up a ran-sim instance
	sim := s.CreateRanSimulatorWithNameOrDie("subscription-delete")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()
	//  Initially the subscription list should be empty
	e2utils.CheckForEmptySubscriptionList(t)

	// Create a Node
	nodeID := utils.GetTestNodeID(t)

	// Add a subscription
	channelID, sub := s.createAndVerifySubscription(ctx, nodeID)
	subscriptionID := s.getSubscriptionID(channelID)

	// Check that the subscription list is correct
	subList := e2utils.GetSubscriptionList(t)
	s.Equal(1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subscriptionID)

	// Close the subscription
	err := sub.Sub.Unsubscribe(ctx)
	s.NoError(err)

	// Check number of subscriptions is correct after deleting the subscription
	e2utils.CheckForEmptySubscriptionList(t)

	//  Open the subscription again and make sure it is open
	channelID, sub = s.createAndVerifySubscription(ctx, nodeID)
	subscriptionID = s.getSubscriptionID(channelID)

	// Check that the number of subscriptions is correct after reopening
	subList = e2utils.GetSubscriptionList(t)
	s.Equal(1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subscriptionID)

	// Close the subscription
	err = sub.Sub.Unsubscribe(ctx)
	s.NoError(err)

	s.True(utils.ReadToEndOfChannel(sub.Sub.Ch))

	e2utils.CheckForEmptySubscriptionList(t)

	// Clean up the ran-sim instance
	s.UninstallRanSimulatorOrDie(sim, "subscription-delete")
}
