// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"testing"
	"time"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	"gotest.tools/assert"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	SubscriptionServiceHost = "onos-e2sub"
	SubscriptionServicePort = 5150
)

func createSubscriptionRequest(nodeID string) (subscription.SubscriptionDetails, error) {
	return subscription.SubscriptionDetails{
		E2NodeID: subscription.E2NodeID(nodeID),
		ServiceModel: subscription.ServiceModel{
			ID: subscription.ServiceModelID("test"),
		},
		EventTrigger: subscription.EventTrigger{
			Payload: subscription.Payload{
				Encoding: subscription.Encoding_ENCODING_PROTO,
				Data:     []byte{},
			},
		},
		Actions: []subscription.Action{
			{
				ID:   100,
				Type: subscription.ActionType_ACTION_TYPE_REPORT,
				SubsequentAction: &subscription.SubsequentAction{
					Type:       subscription.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
					TimeToWait: subscription.TimeToWait_TIME_TO_WAIT_ZERO,
				},
			},
		},
	}, nil
}

// TestSubscription
func (s *TestSuite) TestSubscription(t *testing.T) {
	utils.CreateE2SimulatorWithName(t, "e2-simulator")

	clientConfig := e2client.Config{
		AppID: "subscription-test",
		SubscriptionService: e2client.ServiceConfig{
			Host: SubscriptionServiceHost,
			Port: SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NilError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NilError(t, err)

	subReq, err := createSubscriptionRequest(nodeIDs[0])
	assert.NilError(t, err)

	err = client.Subscribe(ctx, subReq, ch)
	assert.NilError(t, err)

	select {
	case indicationMsg := <-ch:
		t.Log(indicationMsg)

	case <-time.After(20 * time.Second):
		t.Fatal("test is failed because of timeout")

	}

}
