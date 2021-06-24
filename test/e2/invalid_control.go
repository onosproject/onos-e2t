// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"testing"

	"github.com/onosproject/onos-e2t/test/utils"
)

type invalidControlTestCase struct {
	description   string
	control       utils.Control
	enabled       bool
	expectedError func(err error) bool
}

func runControlTestCase(t *testing.T, testCase invalidControlTestCase) {
	/*ctx := context.Background()
	if !testCase.enabled {
		t.Skip()
		return
	}

	e2Client := utils.GetE2Client(t, "invalid-control-test")
	request, err := testCase.control.Create()
	assert.NoError(t, err)
	response, err := e2Client.Control(ctx, request)
	assert.Nil(t, response)
	err = errors.FromGRPC(err)
	assert.Equal(t, true, testCase.expectedError(err))
	t.Log(err)*/

}

// TestInvalidControl tests invalid control requests
func (s *TestSuite) TestInvalidControl(t *testing.T) {
	/*sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "invalid-control")
	nodeIDs, err := utils.GetNodeIDs(t)

	assert.NoError(t, err)
	nodeID := string(nodeIDs[0])

	// The values in the header are for testing of error checking in the NB
	rcControlHeader := utils.RcControlHeader{
		Priority: priority,
		CellID:   123456,
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
		{
			control: utils.Control{
				NodeID:              nodeID,
				EncodingType:        10,
				ServiceModelName:    utils.RcServiceModelName,
				ServiceModelVersion: utils.Version2,
			},
			description:   "Invalid encoding type",
			enabled:       true,
			expectedError: errors.IsInvalid,
		},
		{
			control: utils.Control{
				NodeID:              nodeID,
				EncodingType:        e2tapi.EncodingType_PROTO,
				ServiceModelName:    "no-such-service-model",
				ServiceModelVersion: "v1",
			},
			description:   "Invalid service model",
			enabled:       true,
			expectedError: errors.IsNotFound,
		},
		{
			control: utils.Control{
				Header:  []byte("invalid-control-header"),
				Payload: controlMessageBytes,
			},
			description:   "Invalid control header",
			enabled:       true,
			expectedError: errors.IsInvalid,
		},
		{
			control: utils.Control{
				Header:  controlHeaderBytes,
				Payload: []byte("invalid-control-message"),
			},
			description:   "Invalid control message",
			enabled:       true,
			expectedError: errors.IsInvalid,
		},
	}
	for _, testCase := range testCases {
		pinTestCase := testCase
		t.Run(pinTestCase.description, func(t *testing.T) {
			runControlTestCase(t, pinTestCase)
		})
	}
	err = sim.Uninstall()
	assert.NoError(t, err)*/

}
