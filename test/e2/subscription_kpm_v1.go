// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"
)

// TestSubscriptionKpmV1 tests e2 subscription and subscription delete procedures
func (s *TestSuite) TestSubscriptionKpmV1(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v1")

	nodeID := utils.GetTestNodeID(t)

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
		NodeID:              string(nodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version1,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	subName := "TestSubscriptionKpmV1"

	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version1, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan v1beta1.Indication)
	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	_, err = node.Subscribe(ctx, subName, subReq, ch)
	assert.NoError(t, err)

	e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)

	err = node.Unsubscribe(ctx, subName)
	assert.NoError(t, err)

	err = sim.Uninstall()
	assert.NoError(t, err)

	cancel()
	e2utils.CheckForEmptySubscriptionList(t)
}
