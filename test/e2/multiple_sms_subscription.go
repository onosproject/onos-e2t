// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"time"

	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"github.com/onosproject/onos-e2t/test/e2utils"

	//"github.com/onosproject/onos-e2t/test/e2utils"
	"testing"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	//"time"

	//e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestMultiSmSubscription tests multiple subscription to different service models on different nodes
func (s *TestSuite) TestMultiSmSubscription(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "multi-sm-subscription")
	assert.NotNil(t, sim)

	nodeIDs := utils.GetTestNodeIDs(t, 2)
	assert.True(t, len(nodeIDs) > 0)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	kpmNodeID := nodeIDs[0]
	rcPreNodeID := nodeIDs[1]

	cells, err := topoSdkClient.GetCells(ctx, kpmNodeID)
	assert.NoError(t, err)

	reportPeriod := uint32(5000)
	granularity := uint32(500)
	KPMSubName := "TestSubscriptionKpmV2"
	RCSubName := "TestSubscriptionRCPreV2"

	// Kpm v2 interval is defined in ms
	KPMEventTriggerBytes, err := utils.CreateKpmV2EventTrigger(reportPeriod)
	assert.NoError(t, err)

	// Use one of the cell object IDs for action definition
	cellObjectID := cells[0].CellObjectID
	actionDefinitionBytes, err := utils.CreateKpmV2ActionDefinition(cellObjectID, granularity)
	assert.NoError(t, err)

	var KPMActions []e2api.Action
	KPMAction := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
		Payload: actionDefinitionBytes,
	}

	KPMActions = append(KPMActions, KPMAction)

	KPMSubRequest := utils.Subscription{
		NodeID:              string(kpmNodeID),
		EventTrigger:        KPMEventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
		Actions:             KPMActions,
	}

	KPMSubSpec, err := KPMSubRequest.Create()
	assert.NoError(t, err)

	KPMSdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	KPMNode := KPMSdkClient.Node(sdkclient.NodeID(kpmNodeID))
	KPMch := make(chan e2api.Indication)
	KPMCtx, KPMCancel := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = KPMNode.Subscribe(KPMCtx, KPMSubName, KPMSubSpec, KPMch)
	assert.NoError(t, err)

	// Subscribe to RC service model
	RCch := make(chan e2api.Indication)
	RCEventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)

	var RCActions []e2api.Action
	RCAction := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}

	RCActions = append(RCActions, RCAction)
	RCSubSpec := utils.Subscription{
		NodeID:              string(rcPreNodeID),
		Actions:             RCActions,
		EventTrigger:        RCEventTriggerBytes,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	RCSubReq, err := RCSubSpec.Create()
	assert.NoError(t, err)

	RCSdkClient := utils.GetE2Client(t, utils.RcServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	RCNode := RCSdkClient.Node(sdkclient.NodeID(rcPreNodeID))
	assert.NotNil(t, RCNode)
	RCCtx, RCCancel := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = RCNode.Subscribe(RCCtx, RCSubName, RCSubReq, RCch)
	assert.NoError(t, err)

	// Check that indications can be received
	KPMMsg := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, KPMch)
	RCMsg := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, RCch)

	kpmIndicationHeader := &e2sm_kpm_ies.E2SmKpmIndicationHeader{}
	rcIndicationHeader := &e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(KPMMsg.Header, kpmIndicationHeader)
	assert.NoError(t, err)

	err = proto.Unmarshal(RCMsg.Header, rcIndicationHeader)
	assert.NoError(t, err)

	// Clean up subscriptions
	err = KPMNode.Unsubscribe(context.Background(), KPMSubName)
	assert.NoError(t, err)

	err = RCNode.Unsubscribe(context.Background(), RCSubName)
	assert.NoError(t, err)

	KPMCancel()
	RCCancel()

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
