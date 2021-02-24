// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/test/utils"
)

func checkSubscription(t *testing.T) sdksub.Context {
	clientConfig := e2client.Config{
		AppID: "subscription-delete-test",
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmEventTrigger(12)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelID:       utils.KpmServiceModelID,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := client.Subscribe(ctx, subReq, ch)
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
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "ran-simulator")

	//  Initially the subscription list should be empty
	checkSubscriptionList(t, 0)

	// Add a subscription
	subBeforeDelete := checkSubscription(t)

	// Check that the subscription list is correct
	subListBeforeDelete := checkSubscriptionList(t, 1)
	checkSubscriptionIDInList(t, subBeforeDelete.ID(), subListBeforeDelete)

	// Check that querying the subscription is correct
	checkSubscriptionGet(t, subBeforeDelete.ID())

	//  Close the subscription
	err := subBeforeDelete.Close()
	assert.NoError(t, err)

	// List should be empty now
	checkSubscriptionList(t, 0)

	//  Open the subscription again and make sure it is open
	subAfterDelete := checkSubscription(t)

	// List should be one item now
	subListAfterDelete := checkSubscriptionList(t, 1)
	checkSubscriptionIDInList(t, subAfterDelete.ID(), subListAfterDelete)

	// Check that querying the subscription is correct
	checkSubscriptionGet(t, subAfterDelete.ID())

	// Clean up the ran-sim instance
	err = sim.Uninstall()
	assert.NoError(t, err)
}
