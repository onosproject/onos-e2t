// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
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

func runControlTestCase(t *testing.T, testCase invalidControlTestCase, testNodeID string) {
	ctx := context.Background()
	if !testCase.enabled {
		t.Skip()
		return
	}

	sdkClient := utils.GetE2Client(t, testCase.serviceModelName, testCase.serviceModelVersion, testCase.encodingType)
	node := sdkClient.Node(sdkclient.NodeID(testNodeID))
	request, err := testCase.control.Create()
	assert.NoError(t, err)
	response, err := node.Control(ctx, request)
	assert.Nil(t, response)
	assert.Equal(t, true, testCase.expectedError(err))

}

// TestInvalidControl tests invalid control requests
func (s *TestSuite) TestInvalidControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "invalid-control")
	nodeID := utils.GetTestNodeID(t)

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
	assert.NoError(t, err)
	controlHeaderBytes, err := rcControlHeader.CreateRcControlHeader()
	assert.NoError(t, err)

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
		t.Run(pinTestCase.description, func(t *testing.T) {
			runControlTestCase(t, pinTestCase, string(nodeID))
		})
	}
	utils.UninstallRanSimulatorOrDie(t, sim)
}
