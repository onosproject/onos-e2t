// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	subscriptionName    = "TestSubscriptionDelete-kpm"
	granularity         = uint32(500)
	reportPeriod        = uint32(5000)
	subscriptionTimeout = 2 * time.Minute
)

// createAndVerifySubscription creates a subscription to the given node and makes sure that
// at least one verification message can be received from it. The channel ID of the subscription
// is returned
func createAndVerifySubscription(ctx context.Context, t *testing.T, nodeID topo.ID, node sdkclient.Node) (e2api.ChannelID, chan e2api.Indication) {

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	// Use one of the cell object IDs for action definition
	cells, err := topoSdkClient.GetCells(ctx, nodeID)
	assert.NoError(t, err)
	cellObjectID := cells[0].CellObjectID
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	assert.NoError(t, err)
	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}
	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              string(nodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version1,
	}

	subSpec, err := subRequest.Create()
	assert.NoError(t, err)

	ch := make(chan v1beta1.Indication)
	channelID, err := node.Subscribe(ctx, subscriptionName, subSpec, ch)
	assert.NoError(t, err)

	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)
		assert.NotNil(t, indicationMsg)

	case <-time.After(10 * time.Second):
		assert.Equal(t, false, "test is failed because of timeout")

	}
	return channelID, ch
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
	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))

	// Add a subscription
	channelID, _ := createAndVerifySubscription(ctx, t, nodeID, node)
	subscriptionID := getSubscriptionID(t, channelID)

	// Check that the subscription list is correct
	subList := e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subscriptionID)

	// Close the subscription
	err := node.Unsubscribe(ctx, subscriptionName)
	assert.NoError(t, err)

	// Check number of subscriptions is correct after deleting the subscription
	e2utils.CheckForEmptySubscriptionList(t)

	//  Open the subscription again and make sure it is open
	channelID, ch := createAndVerifySubscription(ctx, t, nodeID, node)
	subscriptionID = getSubscriptionID(t, channelID)

	// Check that the number of subscriptions is correct after reopening
	subList = e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subscriptionID)

	// Close the subscription
	err = node.Unsubscribe(ctx, subscriptionName)
	assert.NoError(t, err)

	assert.True(t, utils.ReadToEndOfChannel(ch))

	e2utils.CheckForEmptySubscriptionList(t)

	// Clean up the ran-sim instance
	utils.UninstallRanSimulatorOrDie(t, sim)
}
