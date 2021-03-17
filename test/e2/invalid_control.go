// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"

	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

type invalidControlTestCase struct {
	description   string
	control       utils.Control
	enabled       bool
	expectedError func(err error) bool
}

func runControlTestCase(t *testing.T, testCase invalidControlTestCase) {
	ctx := context.Background()
	if !testCase.enabled {
		t.Skip()
		return
	}

	e2Client := getE2Client(t, "invalid-control-test")
	request, err := testCase.control.Create()
	assert.NoError(t, err)
	response, err := e2Client.Control(ctx, request)
	assert.Nil(t, response)
	err = errors.FromGRPC(err)
	assert.Equal(t, true, testCase.expectedError(err))
	t.Log(err)

}

// TestInvalidControl tests invalid control requests
func (s *TestSuite) TestInvalidControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "invalid-control")
	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)
	nodeID := nodeIDs[0]

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
				NodeID:         nodeID,
				EncodingType:   10,
				ServiceModelID: utils.RcServiceModelID,
			},
			description:   "Invalid encoding type",
			enabled:       true,
			expectedError: errors.IsInvalid,
		},
		{
			control: utils.Control{
				NodeID:         nodeID,
				EncodingType:   e2tapi.EncodingType_PROTO,
				ServiceModelID: "no-such-service-model",
			},
			description:   "Invalid service model",
			enabled:       true,
			expectedError: errors.IsNotFound,
		},
		{
			control: utils.Control{
				NodeID:         nodeID,
				EncodingType:   e2tapi.EncodingType_PROTO,
				ServiceModelID: utils.RcServiceModelID,
				ControlHeader:  []byte("invalid-control-header"),
				ControlMessage: controlMessageBytes,
			},
			description:   "Invalid control header",
			enabled:       true,
			expectedError: errors.IsInvalid,
		},
		{
			control: utils.Control{
				NodeID:         nodeID,
				EncodingType:   e2tapi.EncodingType_PROTO,
				ServiceModelID: utils.RcServiceModelID,
				ControlHeader:  controlHeaderBytes,
				ControlMessage: []byte("invalid-control-message"),
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
	assert.NoError(t, err)

}
