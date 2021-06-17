// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	subscriptionName    = "sub1"
	granularity         = uint32(500)
	reportPeriod        = uint32(5000)
	subscriptionTimeout = 10 * time.Second
)

// createAndVerifySubscription creates a subscription to the given node and makes sure that
// at least one verification message can be received from it. The channel ID of the subscription
// is returned
func createAndVerifySubscription(ctx context.Context, t *testing.T, nodeID topo.ID, node sdkclient.Node) subapi.ChannelID {

	// Use one of the cell object IDs for action definition
	cells, err := utils.GetCellIDsPerNode(nodeID)
	assert.NoError(t, err)
	cellObjectID := cells[0].CellObjectID
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	assert.NoError(t, err)
	var actions []subapi.Action
	action := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}
	actions = append(actions, action)

	subRequest := utils.Subscription2{
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
	return channelID
}

func getSubscriptionID(t *testing.T, channelID subapi.ChannelID) subapi.SubscriptionID {
	getChannelRequest := &subapi.GetChannelRequest{ChannelID: channelID}
	channelResponse, err := utils.GetSubAdminClient(t).GetChannel(context.Background(), getChannelRequest)
	assert.NoError(t, err)
	channel := channelResponse.Channel
	return channel.GetSubscriptionID()
}

// TestSubscriptionDelete tests subscription delete procedure
func (s *TestSuite) TestSubscriptionDelete(t *testing.T) {
	var err error

	// Start up a ran-sim instance
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-delete")
	assert.NotNil(t, sim)

	//  Initially the subscription list should be empty
	subList := e2utils.GetSubscriptionList2(t)
	defaultNumSubs := len(subList)

	// Create a Node
	nodeID := utils.GetFirstNodeID(t)
	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version2)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)

	// Add a subscription
	channelID := createAndVerifySubscription(ctx, t, nodeID, node)
	subscriptionID := getSubscriptionID(t, channelID)

	// Check that the subscription list is correct
	subList = e2utils.GetSubscriptionList2(t)
	assert.Equal(t, defaultNumSubs+1, len(subList))
	e2utils.CheckSubscriptionIDInList2(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet2(t, subscriptionID)

	// Close the subscription
	err = node.Unsubscribe(ctx, subscriptionName)
	assert.NoError(t, err)
	cancel()

	// Create a context specifying a timeout
	ctx, cancel = context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	// Check number of subscriptions is correct after deleting the subscription
	subList = e2utils.GetSubscriptionList2(t)
	assert.Equal(t, defaultNumSubs, len(subList))

	//  Open the subscription again and make sure it is open
	channelID = createAndVerifySubscription(ctx, t, nodeID, node)
	subscriptionID = getSubscriptionID(t, channelID)

	// Check that the number of subscriptions is correct after reopening
	subList = e2utils.GetSubscriptionList2(t)
	assert.Equal(t, defaultNumSubs+1, len(subList))
	e2utils.CheckSubscriptionIDInList2(t, subscriptionID, subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet2(t, subscriptionID)

	// Clean up the ran-sim instance
	simErr := sim.Uninstall()
	assert.NoError(t, simErr)
}
