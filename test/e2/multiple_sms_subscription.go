// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"testing"
	"time"

	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestMultipleSmSubscription tests multiple subscription to different service models
func (s *TestSuite) TestMultiSmSubscription(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "multi-sm-subscription")
	done := make(chan struct{})
	defer close(done)

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)

	// Subscribe to kpm service model
	ch1 := make(chan e2api.Indication)

	nodeIDs, err := utils.GetNodeIDs(t)
	assert.NoError(t, err)

	testNode1 := nodeIDs[0]
	testNode2 := nodeIDs[1]

	kpmEventTriggerBytes, err := utils.CreateKpmV2EventTrigger(12)
	assert.NoError(t, err)
	var actions []e2api.Action
	action := e2api.Action{
		ID:   100,
		Type: e2api.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subSpec := utils.Subscription2{
		NodeID:              string(testNode1),
		Actions:             actions,
		EventTrigger:        kpmEventTriggerBytes,
		ServiceModelName:    utils.KpmServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	subReq, err := subSpec.Create()
	assert.NoError(t, err)

	sdkClient := utils.GetE2Client2(t, utils.KpmServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node1 := sdkClient.Node(sdkclient.NodeID(testNode1))
	ctx1, cancel1 := context.WithTimeout(context.Background(), 15*time.Second)
	_, err = node1.Subscribe(ctx1, "TestSubscriptionKpmV1", subReq, ch1)
	assert.NoError(t, err)

	// Subscribe to RC service model
	ch2 := make(chan e2api.Indication)
	rcEventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)

	subSpec = utils.Subscription2{
		NodeID:              string(testNode2),
		Actions:             actions,
		EventTrigger:        rcEventTriggerBytes,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	subReq, err = subSpec.CreateWithActionDefinition2()
	assert.NoError(t, err)

	sdkClient2 := utils.GetE2Client2(t, utils.RcServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node2 := sdkClient2.Node(sdkclient.NodeID(testNode1))
	ctx2, cancel2 := context.WithTimeout(context.Background(), 15*time.Second)
	_, err = node2.Subscribe(ctx2, "TestSubscriptionKpmV2", subReq, ch2)
	assert.NoError(t, err)

	msg1 := e2utils.CheckIndicationMessage2(t, e2utils.DefaultIndicationTimeout, ch1)
	msg2 := e2utils.CheckIndicationMessage2(t, e2utils.DefaultIndicationTimeout, ch2)

	kpmIndicationHeader := &e2sm_kpm_ies.E2SmKpmIndicationHeader{}
	rcIndicationHeader := &e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(msg1.Header, kpmIndicationHeader)
	assert.NoError(t, err)

	err = proto.Unmarshal(msg2.Header, rcIndicationHeader)
	assert.NoError(t, err)

	err = sim.Uninstall()
	assert.NoError(t, err)

	cancel1()
	cancel2()
}
