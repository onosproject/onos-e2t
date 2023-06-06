// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"google.golang.org/protobuf/proto"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscriptionMultipleReports tests e2 subscription with multiple reports in one subscription
func (s *TestSuite) TestSubscriptionMultipleReports() {
	sim := s.CreateRanSimulatorWithNameOrDie("subscription-multiple-reports")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(s.T())

	topoSdkClient, err := utils.NewTopoClient()
	s.NoError(err)

	cells, err := topoSdkClient.GetCells(ctx, nodeID)
	s.NoError(err)
	s.GreaterOrEqual(len(cells), 2)

	// Use one of the cell object IDs for action definition
	cellObjectID0 := cells[0].CellObjectID
	cellObjectID1 := cells[1].CellObjectID

	cellObjectIDList := make([]string, 2)
	cellObjectIDList[0] = cellObjectID0
	cellObjectIDList[1] = cellObjectID1

	subName := "TestSubscriptionMultipleReports-kpm"

	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		Granularity:  1000,
		CellObjectID: cellObjectID,
	}

	actionDefinitionBytes0, err := kpmv2Sub.CreateKpmV2ActionDefinition()
	s.NoError(err)
	actionDefinitionBytes1, err := kpmv2Sub.CreateKpmV2ActionDefinition()
	s.NoError(err)

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

	kpmv2Sub.Sub.Actions = actions
	kpmv2Sub.SubscribeOrFail(ctx, s.T())

	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	for i := 0; i < 2; i++ {
		indicationReport := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)
		err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
		s.NoError(err)
		err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
		s.NoError(err)
		indMsgFormat1 := indicationMessage.GetIndicationMessageFormats().GetIndicationMessageFormat1()
		cellObjectID := indMsgFormat1.GetCellObjId().Value
		s.True(cellObjectID == cellObjectIDList[0] || cellObjectID == cellObjectIDList[1])
	}

	kpmv2Sub.Sub.UnsubscribeOrFail(ctx, s.T())

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "subscription-multiple-reports")
}
