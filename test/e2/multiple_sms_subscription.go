// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// TestMultipleSmSubscription tests multiple subscription to different service models
func (s *TestSuite) TestMultiSmSubscription(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "multi-sm-subscription")
	done := make(chan struct{})
	defer close(done)
	// Create an E2 client
	e2Client := getE2Client(t, "multi-sm-subscription-test")

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)

	// Subscribe to kpm service model
	ch1 := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	testNode1 := nodeIDs[0]
	testNode2 := nodeIDs[1]

	kpmEventTriggerBytes, err := utils.CreateKpmEventTrigger(12)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               testNode1,
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         kpmEventTriggerBytes,
		ServiceModelName:     utils.KpmServiceModelName,
		ServiceModelVersion:  utils.KpmServiceModelVersion1,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub1, err := e2Client.Subscribe(ctx, subReq, ch1)
	assert.NoError(t, err)

	// Subscribe to RC service model
	ch2 := make(chan indication.Indication)
	rcEventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)

	subRequest = utils.Subscription{
		NodeID:               testNode2,
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         rcEventTriggerBytes,
		ServiceModelName:     utils.RcServiceModelName,
		ServiceModelVersion:  utils.RcServiceModelVersion1,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err = subRequest.Create()
	assert.NoError(t, err)

	sub2, err := e2Client.Subscribe(ctx, subReq, ch2)
	assert.NoError(t, err)

	msg1 := checkIndicationMessage(t, defaultIndicationTimeout, ch1)
	msg2 := checkIndicationMessage(t, defaultIndicationTimeout, ch2)

	kpmIndicationHeader := &e2sm_kpm_ies.E2SmKpmIndicationHeader{}
	rcIndicationHeader := &e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(msg1.Payload.Header, kpmIndicationHeader)
	assert.NoError(t, err)

	err = proto.Unmarshal(msg2.Payload.Header, rcIndicationHeader)
	assert.NoError(t, err)

	err = sub1.Close()
	assert.NoError(t, err)
	err = sub2.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)

}
