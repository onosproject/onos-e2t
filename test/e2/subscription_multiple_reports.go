// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"google.golang.org/protobuf/proto"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionMultipleReports tests e2 subscription with multiple reports in one subscription
func (s *TestSuite) TestSubscriptionMultipleReports(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-multiple-reports")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	cells, err := topoSdkClient.GetCells(ctx, nodeID)
	assert.NoError(t, err)

	// Kpm v2 interval is defined in ms
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(5000)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	cellObjectID0 := cells[0].CellObjectID
	cellObjectID1 := cells[1].CellObjectID

	cellObjectIDList := make([]string, 2)
	cellObjectIDList[0] = cellObjectID0
	cellObjectIDList[1] = cellObjectID1

	actionDefinitionBytes0, err := utils.CreateKpmV2ActionDefinition(cellObjectIDList[0], 1000)
	assert.NoError(t, err)
	actionDefinitionBytes1, err := utils.CreateKpmV2ActionDefinition(cellObjectIDList[1], 1000)
	assert.NoError(t, err)

	var actions []e2api.Action
	action0 := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes0,
	}

	action1 := e2api.Action{
		ID:   101,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes1,
	}

	actions = append(actions, action0)
	actions = append(actions, action1)

	subRequest := utils.Subscription{
		NodeID:              string(nodeID),
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
		Actions:             actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	assert.NoError(t, err)

	subName := "TestSubscriptionMultipleReports-kpm"

	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan e2api.Indication)
	_, err = node.Subscribe(ctx, subName, subSpec, ch)
	assert.NoError(t, err)

	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	for i := 0; i < 2; i++ {
		indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
		err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
		assert.NoError(t, err)
		err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
		assert.NoError(t, err)
		cellObjectID := indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value
		assert.True(t, cellObjectID == cellObjectIDList[0] || cellObjectID == cellObjectIDList[1])
	}

	err = node.Unsubscribe(ctx, subName)
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
