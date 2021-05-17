// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	"github.com/onosproject/onos-e2t/test/utils"
)

func checkSubscription(t *testing.T) sdksub.Context {

	e2Client := utils.GetE2Client(t, "subscription-delete-test")

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmV1EventTrigger(12)
	assert.NoError(t, err)
	var actions []subapi.Action
	action := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              nodeIDs[0],
		EncodingType:        subapi.Encoding_ENCODING_PROTO,
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version1,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)
	assert.NotNil(t, sub)

	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)
		assert.NotNil(t, indicationMsg)

	case <-time.After(10 * time.Second):
		t.Fatal("test is failed because of timeout")

	}
	return sub
}

// TestSubscriptionDelete tests subscription delete procedure
func (s *TestSuite) TestSubscriptionDelete(t *testing.T) {
	// Start up a ran-sim instance
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "subscription-delete")

	//  Initially the subscription list should be empty
	subList := e2utils.GetSubscriptionList(t)
	defaultNumSubs := len(subList)

	// Add a subscription
	subBeforeDelete := checkSubscription(t)

	// Check that the subscription list is correct
	subList = e2utils.GetSubscriptionList(t)
	assert.Equal(t, defaultNumSubs+1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subBeforeDelete.ID(), subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subBeforeDelete.ID())

	// Close the subscription
	err := subBeforeDelete.Close()
	assert.NoError(t, err)

	// Check number of subscriptions is correct after deleting the subscription
	subList = e2utils.GetSubscriptionList(t)
	assert.Equal(t, defaultNumSubs, len(subList))

	//  Open the subscription again and make sure it is open
	subAfterDelete := checkSubscription(t)
	subList = e2utils.GetSubscriptionList(t)
	assert.Equal(t, defaultNumSubs+1, len(subList))
	e2utils.CheckSubscriptionIDInList(t, subAfterDelete.ID(), subList)

	// Check that querying the subscription is correct
	e2utils.CheckSubscriptionGet(t, subAfterDelete.ID())

	// Clean up the ran-sim instance
	err = sim.Uninstall()
	assert.NoError(t, err)
}
