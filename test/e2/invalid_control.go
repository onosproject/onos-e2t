// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"testing"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

type invalidControlTestCase struct {
	description         string
	control             utils.Control
	enabled             bool
	expectedError       func(err error) bool
	serviceModelName    string
	serviceModelVersion string
	encodingType        sdkclient.Encoding
}

func (s *TestSuite) runControlTestCase(testCase invalidControlTestCase, testNodeID string) {
	if !testCase.enabled {
		s.T().Skip()
		return
	}

	sdkClient := utils.GetE2Client(s.T(), testCase.serviceModelName, testCase.serviceModelVersion, testCase.encodingType)
	node := sdkClient.Node(sdkclient.NodeID(testNodeID))
	request, err := testCase.control.Create()
	s.NoError(err)
	response, err := node.Control(s.Context(), request, nil)
	s.Nil(response)
	s.Equal(true, testCase.expectedError(err))

}

// TestInvalidControl tests invalid control requests
func (s *TestSuite) TestInvalidControl() {
	sim := s.CreateRanSimulatorWithNameOrDie("invalid-control")
	nodeID := utils.GetTestNodeID(s.T())

	// The values in the header are for testing of error checking in the NB
	rcControlHeader := utils.RcControlHeader{
		Priority: priority,
		CellID:   utils.Uint64ToBitString(123456, 36),
		PlmnID:   ransimtypes.NewUint24(654321).ToBytes(),
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

	testCases := []invalidControlTestCase{
		/*{
			control: utils.Control{
				NodeID:              nodeID,
				EncodingType:        10,
				ServiceModelName:    utils.RcServiceModelName,
				ServiceModelVersion: utils.Version2,
			},
			description:   "Invalid encoding type",
			enabled:       true,
			expectedError: errors.IsInvalid,
		},*/
		{
			control:             utils.Control{},
			description:         "Invalid service model",
			enabled:             true,
			expectedError:       errors.IsNotFound,
			serviceModelName:    "no-such-service-model",
			serviceModelVersion: utils.Version2,
			encodingType:        sdkclient.ProtoEncoding,
		},
		{
			control: utils.Control{
				Header:  []byte("invalid-control-header"),
				Payload: controlMessageBytes,
			},
			description:         "Invalid control header",
			enabled:             true,
			expectedError:       errors.IsInvalid,
			serviceModelName:    utils.RcServiceModelName,
			serviceModelVersion: utils.Version2,
			encodingType:        sdkclient.ProtoEncoding,
		},
		{
			control: utils.Control{
				Header:  controlHeaderBytes,
				Payload: []byte("invalid-control-message"),
			},
			description:         "Invalid control message",
			enabled:             true,
			expectedError:       errors.IsInvalid,
			serviceModelName:    utils.RcServiceModelName,
			serviceModelVersion: utils.Version2,
			encodingType:        sdkclient.ProtoEncoding,
		},
	}
	for _, testCase := range testCases {
		pinTestCase := testCase
		s.T().Run(pinTestCase.description, func(t *testing.T) {
			s.runControlTestCase(pinTestCase, string(nodeID))
		})
	}
	s.UninstallRanSimulatorOrDie(sim, "invalid-control")
}
