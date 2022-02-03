// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"google.golang.org/protobuf/proto"

	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	ranParameterValue uint32 = 200
	ranParameterName         = "pci"
	ranParameterID    int32  = 1
	priority                 = 10
)

// TestControl tests E2 control procedure using ransim and SDK
func (s *TestSuite) TestControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "control-oran-e2sm-rc-pre-v2")
	assert.NotNil(t, sim)
	ch := make(chan e2api.Indication)
	ctx := context.Background()

	// Get a test e2 node ID
	testNodeID := utils.GetTestNodeID(t)

	// Create E2 SDK Client
	sdkClient := utils.GetE2Client(t, utils.RcServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(testNodeID))

	// Create a subscription request
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
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

	subRequest := utils.Subscription{
		NodeID:              string(testNodeID),
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	subName := "control-subscribe-oran-e2sm-rc-pre-v2"

	// Subscribe to RC Pre service model
	_, err = node.Subscribe(ctx, subName, subReq, ch)
	assert.NoError(t, err)

	// Receive and process the first indication message
	indMessage := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	header := indMessage.Header
	ricIndicationHeader := e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().Value
	nrcid := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().Value.GetValue()

	rcControlHeader := utils.RcControlHeader{
		Priority: priority,
		CellID:   nrcid,
		PlmnID:   plmnID,
	}

	rcControlMessage := utils.RcControlMessage{
		RanParameterName:  ranParameterName,
		RanParameterID:    ranParameterID,
		RanParameterValue: ranParameterValue,
	}

	controlMessageBytes, err := rcControlMessage.CreateRcControlMessage()
	assert.NoError(t, err)
	controlHeaderBytes, err := rcControlHeader.CreateRcControlHeader()
	assert.NoError(t, err)

	controlRequest := utils.Control{
		Payload: controlMessageBytes,
		Header:  controlHeaderBytes,
	}

	// Create a control request to change PCI value
	request, err := controlRequest.Create()
	assert.NoError(t, err)
	response, err := node.Control(ctx, request)
	assert.NoError(t, err)

	assert.NotNil(t, response)
	assert.NotNil(t, response.Payload)
	controlOutcome := &e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(response.Payload, controlOutcome)
	assert.NoError(t, err)

	outcomeRanParameterID := controlOutcome.
		GetControlOutcomeFormat1().
		GetOutcomeElementList()[0].
		RanParameterId.Value

	assert.Equal(t, ranParameterID, outcomeRanParameterID)

	// Delete subscription and ran simulator
	err = node.Unsubscribe(ctx, subName)
	assert.NoError(t, err)
	e2utils.CheckForEmptySubscriptionList(t)

	err = sim.Uninstall()
	assert.NoError(t, err)
}
