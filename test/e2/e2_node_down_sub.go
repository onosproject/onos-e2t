// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"testing"
	"time"

	"github.com/onosproject/helmit/pkg/kubernetes"

	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
)

// TestE2NodeDownSubscription checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeDownSubscription(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2node-down-subscription")

	ctx, cancel := context.WithCancel(context.Background())

	nodeID := utils.GetFirstNodeID(t)

	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(12)
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
		NodeID:              string(nodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	// Create a subscription request to indication messages from the client
	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	kube, err := kubernetes.NewForRelease(sim)
	assert.NoError(t, err)

	// Cause the simulator to crash
	err = sim.Uninstall()
	assert.NoError(t, err)

	for {
		pods, err := kube.CoreV1().Pods().List()
		assert.NoError(t, err)
		if len(pods) > 0 {
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	//  Create the subscription
	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version2)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan subapi.Indication)
	err = node.Subscribe(ctx, "TestE2NodeDownSubscription", subReq, ch)
	assert.NoError(t, err)

	// Make sure that reads on the subscription channel time out. There should be no
	// indication messages available
	indicationFailed := false

	select {
	case indicationMsg := <-ch:
		// We got an indication. This is an error, as there is no E2 node to send one
		t.Log(indicationMsg)

	case <-time.After(10 * time.Second):
		// The read timed out. This is the expected behavior.
		indicationFailed = true

	}

	assert.True(t, indicationFailed, "Indication message was delivered for a node that is down")
	cancel()
}
