// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
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
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "control-admin-api")
	assert.NotNil(t, sim)
	ch := make(chan indication.Indication)
	ctx := context.Background()

	e2Client := utils.GetE2Client(t, "control-pci-test-admin-api")

	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)

	nodeIDs, err := utils.NodeIDs()
	assert.NoError(t, err)
	testNodeID := nodeIDs[0]

	// Subscription
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)
	var actions []subapi.Action
	action := subapi.Action{
		ID:   100,
		Type: subapi.ActionType_ACTION_TYPE_REPORT,
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: subapi.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              testNodeID,
		EncodingType:        subapi.Encoding_ENCODING_PROTO,
		Actions:             actions,
		EventTrigger:        eventTriggerBytes,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := e2Client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)
	indMessage := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, ch)
	header := indMessage.Payload.Header
	ricIndicationHeader := e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}

	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	plmnID := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetPLmnIdentity().Value
	nrcid := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetNrCgi().GetNRcellIdentity().Value.Value

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
		NodeID:              testNodeID,
		EncodingType:        e2tapi.EncodingType_PROTO,
		ServiceModelName:    utils.RcServiceModelName,
		ServiceModelVersion: utils.Version2,
		ControlAckRequest:   e2tapi.ControlAckRequest_ACK,
		ControlMessage:      controlMessageBytes,
		ControlHeader:       controlHeaderBytes,
	}

	request, err := controlRequest.Create()
	assert.NoError(t, err)
	response, err := e2Client.Control(ctx, request)
	assert.NoError(t, err)
	if response == nil {
		t.Fail()
	}

	ack := response.GetControlAcknowledge()
	failure := response.GetControlFailure()
	if ack != nil {
		controlOutcome := &e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
		err = proto.Unmarshal(ack.GetControlOutcome(), controlOutcome)
		assert.NoError(t, err)

		outcomeRanParameterID := controlOutcome.
			GetControlOutcomeFormat1().
			GetOutcomeElementList()[0].
			RanParameterId.Value

		assert.Equal(t, ranParameterID, outcomeRanParameterID)
	}
	if failure != nil {
		t.Fail()
	}

	err = sub.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)

}
