// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"google.golang.org/protobuf/proto"

	e2smrcpreies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"

	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	ranParameterValue int64 = 200
	ranParameterName        = "pci"
	ranParameterID    int32 = 1
	priority                = 10
)

// TestControl tests E2 control procedure using ransim and SDK
func (s *TestSuite) TestControl() {
	sim := s.CreateRanSimulatorWithNameOrDie("control-oran-e2sm-rc-pre-v2")
	s.NotNil(sim)

	// Get a test e2 node ID
	testNodeID := utils.GetTestNodeID(s.T())

	// Create E2 SDK Client
	sdkClient := utils.GetE2Client(s.T(), utils.RcServiceModelName, utils.Version2, sdkclient.ProtoEncoding)
	node := sdkClient.Node(sdkclient.NodeID(testNodeID))
	subName := "control-subscribe-oran-e2sm-rc-pre-v2"

	// Create an RC PRE subscription
	rcPreSub := e2utils.RCPreSub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: testNodeID,
		},
	}
	s.NoError(rcPreSub.UseDefaultReportAction())
	rcPreSub.SubscribeOrFail(s.Context(), s.T())

	// Receive and process the first indication message
	indMessage := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, rcPreSub.Sub.Ch)
	header := indMessage.Header
	ricIndicationHeader := e2smrcpreies.E2SmRcPreIndicationHeader{}

	err := proto.Unmarshal(header, &ricIndicationHeader)
	s.NoError(err)
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
	s.NoError(err)
	controlHeaderBytes, err := rcControlHeader.CreateRcControlHeader()
	s.NoError(err)

	controlRequest := utils.Control{
		Payload: controlMessageBytes,
		Header:  controlHeaderBytes,
	}

	// Create a control request to change PCI value
	request, err := controlRequest.Create()
	s.NoError(err)
	response, err := node.Control(s.Context(), request, nil)
	s.NoError(err)

	s.NotNil(response)
	s.NotNil(response.Payload)
	controlOutcome := &e2smrcpreies.E2SmRcPreControlOutcome{}
	err = proto.Unmarshal(response.Payload, controlOutcome)
	s.NoError(err)

	outcomeRanParameterID := controlOutcome.
		GetControlOutcomeFormat1().
		GetOutcomeElementList()[0].
		RanParameterId.Value

	s.Equal(ranParameterID, outcomeRanParameterID)

	// Delete subscription and ran simulator
	rcPreSub.Sub.UnsubscribeOrFail(s.Context(), s.T())
	e2utils.CheckForEmptySubscriptionList(s.T())

	s.UninstallRanSimulatorOrDie(sim, "control-oran-e2sm-rc-pre-v2")
}
