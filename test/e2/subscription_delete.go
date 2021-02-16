// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	sdksub "github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/test/utils"
)

func checkSubscriptionList(t *testing.T, expectedLen int) []subscription.Subscription {
	conn, err := utils.ConnectSubscriptionServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	subClient := sdksub.NewClient(conn)
	subList, err := subClient.List(context.Background())
	assert.NoError(t, err)
	assert.Len(t, subList, expectedLen)
	return subList
}

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
	sim := utils.CreateRanSimulatorWithName(t, "ran-simulator")
	assert.NotNil(t, sim)

	conn, err := utils.ConnectSubscriptionServiceHost()
	assert.NoError(t, err)
	assert.NotNil(t, conn)

	//  Initially the subscription list should be empty
	checkSubscriptionList(t, 0)

	// Add a subscription and that there is one item in the list
	sub := checkSubscription(t)
	assert.NotNil(t, sub)
	checkSubscriptionList(t, 1)

	//  Close the subscription and check that the list is empty again
	err = sub.Close()
	assert.NoError(t, err)
	checkSubscriptionList(t, 0)

	//  Open the subscription again and make sure it is open
	sub = checkSubscription(t)
	assert.NotNil(t, sub)
	checkSubscriptionList(t, 1)

	err = sim.Uninstall()
	assert.NoError(t, err)
}
