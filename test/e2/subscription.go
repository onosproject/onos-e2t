// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"testing"
	"time"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/encoding"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	"gotest.tools/assert"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/node"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
)

const (
	SubscriptionServiceHost = "onos-e2sub"
	SubscriptionServicePort = 5150
)

func createSubscriptionRequest(nodeID string) (subscription.Subscription, error) {
	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[100] = types.RicActionDef{
		RicActionID:         100,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_REPORT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO,
		RicActionDefinition: []byte{0x11, 0x22},
	}

	E2apPdu, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(types.RicRequest{RequestorID: 0, InstanceID: 0},
		0, nil, ricActionsToBeSetup)

	if err != nil {
		return subscription.Subscription{}, err
	}

	subReq := subscription.Subscription{
		EncodingType: encoding.PROTO,
		NodeID:       node.ID(nodeID),
		Payload: subscription.Payload{
			Value: E2apPdu,
		},
	}

	return subReq, nil

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
