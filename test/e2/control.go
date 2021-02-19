// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"

	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/pdubuilder"
	"github.com/stretchr/testify/assert"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/test/utils"
)

func createControlHeader() ([]byte, error) {
	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlHeader()
	if err != nil {
		return []byte{}, err
	}

	err = newE2SmRcPrePdu.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil

}

func createControlMessage() ([]byte, error) {
	var ranParameterName = "PCITEST"
	var ranParameterValue int32 = 15

	newE2SmRcPrePdu, err := pdubuilder.CreateE2SmRcPreControlMessage(ranParameterName, ranParameterValue)
	if err != nil {
		return []byte{}, err
	}
	err = newE2SmRcPrePdu.Validate()
	if err != nil {
		return []byte{}, err
	}

	protoBytes, err := proto.Marshal(newE2SmRcPrePdu)
	if err != nil {
		return []byte{}, err
	}

	return protoBytes, nil
}

// TestControl
func (s *TestSuite) TestControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithName(t, "ran-simulator")
	assert.NotNil(t, sim)

	clientConfig := e2client.Config{
		AppID: "control-test",
		E2TService: e2client.ServiceConfig{
			Host: utils.E2TServiceHost,
			Port: utils.E2TServicePort,
		},
	}

	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	controlMessageBytes, err := createControlMessage()
	controlHeaderBytes, err := createControlHeader()

	request := &e2tapi.ControlRequest{
		E2NodeID: e2tapi.E2NodeID(nodeIDs[1]),
		Header: &e2tapi.RequestHeader{
			EncodingType: e2tapi.EncodingType_PROTO,
			ServiceModel: &e2tapi.ServiceModel{
				ID: e2tapi.ServiceModelID("e2sm_rc_pre-v1"),
			},
		},
		ControlAckRequest: e2tapi.ControlAckRequest_ACK,
		ControlMessage:    controlMessageBytes,
		ControlHeader:     controlHeaderBytes,
	}

	start := time.Now()
	response, err := client.Control(ctx, request)
	duration := time.Since(start)
	t.Log(duration.Microseconds())
	assert.NoError(t, err)
	if response == nil {
		return
	}

	//t.Log(response)
	// Code to measure

	//_ = sim.Uninstall()

}
