// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

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
	nodeID := utils.GetTestNodeID(t)

	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(5000)
	assert.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	cells, err := topoSdkClient.GetCells(ctx, nodeID)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	granularity := uint32(500)
	cellObjectID := cells[0].CellObjectID
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
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

	subRequest := utils.Subscription{
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
	utils.UninstallRanSimulatorOrDie(t, sim)

	for {
		pods, err := kube.CoreV1().Pods().List(context.Background())
		assert.NoError(t, err)
		if len(pods) > 0 {
			time.Sleep(time.Second)
		} else {
			t.Log("no ransim pod")
			break
		}
	}

	//  Create the subscription
	subName := "TestE2NodeDownSubscription"
	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan subapi.Indication)
	_, err = node.Subscribe(ctx, subName, subReq, ch)

	//  Subscribe should have failed because of a timeout
	assert.Error(t, err)
	cancel()

	// Delete the subscription and ran simulator
	sim = utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2node-down-subscription")
	node = sdkClient.Node(sdkclient.NodeID(nodeID))
	err = node.Unsubscribe(context.Background(), subName)
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)
	err = sim.Uninstall()
	assert.NoError(t, err)

}
