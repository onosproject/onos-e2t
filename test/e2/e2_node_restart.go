// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"sync"
	"testing"

	"github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2smkpmv2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestE2NodeRestart checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeRestart(t *testing.T) {
	// Create a simulator
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "e2node-restart-subscription")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	subClient := utils.GetSubAdminClient(t)
	res, err := subClient.WatchSubscriptions(ctx, &v1beta1.WatchSubscriptionsRequest{NoReplay: true})
	assert.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(1)

	var pause sync.WaitGroup
	pause.Add(1)
	go func() {
		e, err := res.Recv()
		assert.NoError(t, err)
		assert.Equal(t, v1beta1.SubscriptionEventType_SUBSCRIPTION_CREATED, e.Event.Type)
		assert.Equal(t, v1beta1.SubscriptionState_SUBSCRIPTION_PENDING, e.Event.Subscription.Status.State)
		pause.Done()

		e, err = res.Recv()
		assert.NoError(t, err)
		assert.Equal(t, v1beta1.SubscriptionEventType_SUBSCRIPTION_UPDATED, e.Event.Type)
		assert.Equal(t, v1beta1.SubscriptionState_SUBSCRIPTION_COMPLETE, e.Event.Subscription.Status.State)

		e, err = res.Recv()
		assert.NoError(t, err)
		assert.Equal(t, v1beta1.SubscriptionEventType_SUBSCRIPTION_UPDATED, e.Event.Type)
		assert.Equal(t, v1beta1.SubscriptionState_SUBSCRIPTION_PENDING, e.Event.Subscription.Status.State)

		e, err = res.Recv()
		assert.NoError(t, err)
		assert.Equal(t, v1beta1.SubscriptionEventType_SUBSCRIPTION_UPDATED, e.Event.Type)
		assert.Equal(t, v1beta1.SubscriptionState_SUBSCRIPTION_COMPLETE, e.Event.Subscription.Status.State)

		wg.Done()
	}()

	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	nodeID := utils.GetTestNodeID(t)
	cells, err := topoSdkClient.GetCells(context.Background(), nodeID)
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

	subSpec, err := subRequest.CreateWithActionDefinition2()
	assert.NoError(t, err)

	subName := "TestSubscriptionKpmV2"

	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	ch := make(chan v1beta1.Indication)
	_, err = node.Subscribe(ctx, subName, subSpec, ch)
	assert.NoError(t, err)

	indicationReport := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	indicationMessage := e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader := e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	assert.Equal(t, indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(reportPeriod/granularity), len(indicationMessage.GetIndicationMessageFormat1().GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	t.Log("Restart e2 node")
	_ = sim.Uninstall()
	_ = sim.Install(true)

	t.Log("Check indications")
	indicationReport = e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	indicationMessage = e2smkpmv2.E2SmKpmIndicationMessage{}
	indicationHeader = e2smkpmv2.E2SmKpmIndicationHeader{}

	err = proto.Unmarshal(indicationReport.Payload, &indicationMessage)
	assert.NoError(t, err)
	assert.Equal(t, indicationMessage.GetIndicationMessageFormat1().GetCellObjId().Value, cellObjectID)
	assert.Equal(t, int(reportPeriod/granularity), len(indicationMessage.GetIndicationMessageFormat1().GetMeasData().GetValue()))

	err = proto.Unmarshal(indicationReport.Header, &indicationHeader)
	assert.NoError(t, err)

	t.Log("Unsubscribe")
	err = node.Unsubscribe(context.Background(), subName)
	assert.NoError(t, err)

	err = sim.Uninstall()
	assert.NoError(t, err)

	e2utils.CheckForEmptySubscriptionList(t)
}
