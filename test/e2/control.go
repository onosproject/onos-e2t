// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"

	"gotest.tools/assert"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestSubscription
func (s *TestSuite) TestControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithName(t, "ran-simulator")
	assert.Assert(t, sim != nil)

	clientConfig := e2client.Config{
		AppID: "control-test",
		E2TService: e2client.ServiceConfig{
			Host: E2TServiceHost,
			Port: E2tServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NilError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NilError(t, err)

	response, err := client.Control(ctx, &e2tapi.ControlRequest{
		E2NodeID: e2tapi.E2NodeID(nodeIDs[0]),
		Header: &e2tapi.RequestHeader{
			ServiceModel: &e2tapi.ServiceModel{
				ID: e2tapi.ServiceModelID("e2sm_rc_pre-v1"),
			},
		},
	})

	t.Log(response)

	//_ = sim.Uninstall()

}
