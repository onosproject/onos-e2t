// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

type invalidControlTestCase struct {
	description   string
	control       utils.Control
	enabled       bool
	expectedError errors.Type
}

func runControlTestCase(t *testing.T, testCase invalidControlTestCase) {
	ctx := context.Background()
	if !testCase.enabled {
		t.Skip()
		return
	}
	clientConfig := e2client.Config{
		AppID: "invalid-control-test",
		E2TService: e2client.ServiceConfig{
			Host: utils.E2TServiceHost,
			Port: utils.E2TServicePort,
		},
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}

	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	request, err := testCase.control.Create()
	assert.NoError(t, err)
	response, err := client.Control(ctx, request)
	assert.Nil(t, response)
	assert.Error(t, err, testCase.expectedError)

}

// TestInvalidControl tests invalid control requests
func (s *TestSuite) TestInvalidControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "invalid-control")
	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)
	nodeID := nodeIDs[0]

	testCases := []invalidControlTestCase{
		{
			control: utils.Control{
				NodeID:         nodeID,
				EncodingType:   10,
				ServiceModelID: utils.RcServiceModelID,
			},
			description:   "Invalid encoding type",
			enabled:       true,
			expectedError: errors.Invalid,
		},
		{
			control: utils.Control{
				NodeID:         nodeID,
				EncodingType:   e2tapi.EncodingType_PROTO,
				ServiceModelID: "no-such-service-model",
			},
			description:   "Invalid service model",
			enabled:       true,
			expectedError: errors.NotFound,
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
