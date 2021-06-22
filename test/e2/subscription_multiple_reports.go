// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"testing"

	"github.com/onosproject/onos-e2t/test/e2utils"

	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"google.golang.org/protobuf/proto"

	subapi "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionMultipleReports tests e2 subscription with multiple reports in one subscription
func (s *TestSuite) TestSubscriptionMultipleReports(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-multiple-reports")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs(t)
	assert.NoError(t, err)

	nodeID := nodeIDs[0]
	cells, err := utils.GetCellIDsPerNode(nodeID)
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

	var actions []subapi.Action
	action0 := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes0,
	}

	action1 := subapi.Action{
		ID:   101,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes1,
	}

	actions = append(actions, action0)
	actions = append(actions, action1)

	subRequest := utils.Subscription2{
		NodeID:              string(nodeID),
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
		Actions:             actions,
	}

	subSpec, err := subRequest.CreateWithActionDefinition2()
	assert.NoError(t, err)

	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan subapi.Indication)
	_, err = node.Subscribe(ctx, "TestSubscriptionKpmV2", subSpec, ch)
	assert.NoError(t, err)

	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	for i := 0; i < 2; i++ {
		indicationReport := e2utils.CheckIndicationMessage2(t, e2utils.DefaultIndicationTimeout, ch)
		err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
		assert.NoError(t, err)
		err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
		assert.NoError(t, err)
		cellObjectID := indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value
		assert.True(t, cellObjectID == cellObjectIDList[0] || cellObjectID == cellObjectIDList[1])
	}

	err = sim.Uninstall()
	assert.NoError(t, err)

}
