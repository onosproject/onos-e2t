// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"time"

	e2smkpmies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-e2t/test/e2utils"

	"testing"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

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

	kpmNodeID := nodeIDs[0]
	rcPreNodeID := nodeIDs[1]

	KPMSubName := "TestSubscriptionKpmV2"
	RCSubName := "TestSubscriptionRCPreV2"

	KPMCtx, KPMCancel := context.WithTimeout(context.Background(), 30*time.Second)

	nodeID := utils.GetTestNodeID(t)

	cellObjectID := e2utils.GetFirstCellObjectID(t, nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   KPMSubName,
			NodeID: kpmNodeID,
		},
		CellObjectID: cellObjectID,
	}
	kpmv2Sub.SubscribeOrFail(KPMCtx, t)

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
	KPMMsg := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub.Sub.Ch)
	RCMsg := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, RCch)

	kpmIndicationHeader := &e2smkpmies.E2SmKpmIndicationHeader{}
	rcIndicationHeader := &e2smrcpreies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(KPMMsg.Header, kpmIndicationHeader)
	assert.NoError(t, err)

	err = proto.Unmarshal(RCMsg.Header, rcIndicationHeader)
	assert.NoError(t, err)

	// Clean up subscriptions
	kpmv2Sub.UnsubscribeOrFail(context.Background(), t)

	err = RCNode.Unsubscribe(context.Background(), RCSubName)
	assert.NoError(t, err)

	KPMCancel()
	RCCancel()

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
