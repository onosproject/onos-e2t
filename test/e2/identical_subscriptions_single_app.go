// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"google.golang.org/protobuf/proto"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

func verifyIndicationMessages(t *testing.T, ch chan e2api.Indication, cellObjectID string) {
	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err := proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	assert.Equal(t, indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(reportPeriod/granularity), len(indicationMessage.GetIndicationMessageFormat1().GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)
}

// TestIdenticalSubscriptionSingleApp tests identical subscriptions are absorbed by E2T in a single xApp
func (s *TestSuite) TestIdenticalSubscriptionSingleApp(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "identical-subscriptions-single-app")
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

	subName1 := "identical-sub1"
	subName2 := "identical-sub2"

	sdkClient := utils.GetE2Client(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch1 := make(chan v1beta1.Indication)
	channelID1, err := node.Subscribe(ctx, subName1, subSpec, ch1)
	assert.NoError(t, err)

	ch2 := make(chan v1beta1.Indication)
	channelID2, err := node.Subscribe(ctx, subName2, subSpec, ch2)
	assert.NoError(t, err)

	assert.True(t, channelID1 != channelID2)

	// Should be able to receive indication messages on both channels
	verifyIndicationMessages(t, ch1, cellObjectID)
	verifyIndicationMessages(t, ch2, cellObjectID)

	subList := e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))

	err = node.Unsubscribe(ctx, subName1)
	assert.NoError(t, err)

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName1, subList)

	err = node.Unsubscribe(ctx, subName2)
	assert.NoError(t, err)

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName2, subList)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
