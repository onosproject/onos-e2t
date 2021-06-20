// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestSubscriptionKpmV1 tests e2 subscription and subscription delete procedures
func (s *TestSuite) TestSubscriptionKpmV1(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v1")

	ctx, cancel := context.WithCancel(context.Background())

	nodeIDs, err := utils.GetNodeIDs(t)
	assert.NoError(t, err)
	nodeID := nodeIDs[0]

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

	subRequest := utils.Subscription2{
		NodeID:              string(nodeIDs[0]),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version1,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version1)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan v1beta1.Indication)
	_, err = node.Subscribe(ctx, "TestSubscriptionKpmV1", subReq, ch)
	assert.NoError(t, err)

	e2utils.CheckIndicationMessage2(t, e2utils.DefaultIndicationTimeout, ch)

	cancel()

	select {
	case <-ch:
		assert.Equal(t, false, "received an extraneous indication")

	case <-time.After(10 * time.Second):
		t.Log("Subscription test is PASSED")
	}
	err = sim.Uninstall()
	assert.NoError(t, err)

}
