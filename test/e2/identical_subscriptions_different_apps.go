// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestIdenticalSubscriptionMultiApps tests identical subscriptions are absorbed by E2T from different xApps
func (s *TestSuite) TestIdenticalSubscriptionMultiApps(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "identical-subscription-multi-app")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	cells, err := topoSdkClient.GetCells(ctx, nodeID)
	assert.NoError(t, err)

	reportPeriod := uint32(5000)
	granularity := uint32(500)

	// Kpm v2 interval is defined in ms
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	cellObjectID := cells[0].CellObjectID
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	assert.NoError(t, err)

	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}

	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              string(nodeID),
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
		Actions:             actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	assert.NoError(t, err)

	subName1 := "identical-sub-app1"
	subName2 := "identical-sub-app2"

	clientApp1 := sdkclient.NewClient(sdkclient.WithE2TAddress(utils.E2TServiceHost, utils.E2TServicePort),
		sdkclient.WithServiceModel(utils.KpmServiceModelName,
			utils.Version2),
		sdkclient.WithEncoding(sdkclient.ProtoEncoding),
		sdkclient.WithAppID("app1"))

	nodeApp1 := clientApp1.Node(sdkclient.NodeID(nodeID))

	clientApp2 := sdkclient.NewClient(sdkclient.WithE2TAddress(utils.E2TServiceHost, utils.E2TServicePort),
		sdkclient.WithServiceModel(utils.KpmServiceModelName, utils.Version2),
		sdkclient.WithEncoding(sdkclient.ProtoEncoding),
		sdkclient.WithAppID("app2"))

	nodeApp2 := clientApp2.Node(sdkclient.NodeID(nodeID))

	ch1 := make(chan v1beta1.Indication)
	channelIDApp1, err := nodeApp1.Subscribe(ctx, subName1, subSpec, ch1)
	assert.NoError(t, err)

	ch2 := make(chan v1beta1.Indication)
	channelIDApp2, err := nodeApp2.Subscribe(ctx, subName2, subSpec, ch2)
	assert.NoError(t, err)

	assert.True(t, channelIDApp1 != channelIDApp2)

	indicationReportApp1 := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch1)
	assert.NotNil(t, indicationReportApp1)
	indicationReportApp2 := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch2)
	assert.NotNil(t, indicationReportApp2)

	subList := e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))

	err = nodeApp1.Unsubscribe(ctx, subName1)
	assert.NoError(t, err)

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName1, subList)

	err = nodeApp2.Unsubscribe(ctx, subName2)
	assert.NoError(t, err)

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName2, subList)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
