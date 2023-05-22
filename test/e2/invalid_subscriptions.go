// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/test/e2utils"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1/e2errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"

	"github.com/onosproject/onos-e2t/test/utils"
)

type invalidSubscriptionTestCase struct {
	description         string
	encodingType        sdkclient.Encoding
	actionType          e2api.ActionType
	serviceModelName    e2api.ServiceModelName
	serviceModelVersion e2api.ServiceModelVersion
	eventTrigger        []byte
	actionID            int32
	expectedError       func(err error) bool
	enabled             bool
}

func (s *TestSuite) runTestCase(testCase invalidSubscriptionTestCase) {
	if !testCase.enabled {
		s.T().Skip()
		return
	}

	ctx, cancel := context.WithTimeout(s.Context(), 30*time.Second)
	defer cancel()

	nodeID := utils.GetTestNodeID(s.T())

	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:                testCase.description,
			NodeID:              nodeID,
			EventTriggerBytes:   testCase.eventTrigger,
			ServiceModelName:    testCase.serviceModelName,
			ServiceModelVersion: testCase.serviceModelVersion,
			EncodingType:        testCase.encodingType,
		},
		CellObjectID: cellObjectID,
	}

	action := e2api.Action{
		ID:   testCase.actionID,
		Type: testCase.actionType,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	kpmv2Sub.Sub.Actions = append(kpmv2Sub.Sub.Actions, action)

	_, err := kpmv2Sub.Subscribe(ctx)

	s.True(testCase.expectedError(err))

	kpmv2Sub.Sub.UnsubscribeOrFail(ctx, s.T())

}

// TestInvalidSubscriptions tests invalid inputs into the SDK
func (s *TestSuite) TestInvalidSubscriptions() {
	const actionID = 11
	eventTriggerBytes := e2utils.KPMV2Sub{}.Sub.EventTriggerBytes

	testCases := []invalidSubscriptionTestCase{
		{
			description:         "Non-existent Service Model ID",
			enabled:             true,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    "no-such-service-model",
			serviceModelVersion: "v1",
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			expectedError:       errors.IsNotFound,
		},
		{
			description:         "Invalid action type",
			enabled:             true,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_INSERT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version2,
			eventTrigger:        eventTriggerBytes,
			actionID:            actionID,
			expectedError:       e2errors.IsRICActionNotSupported,
		},
		{
			description:         "Invalid action ID",
			enabled:             false,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version2,
			eventTrigger:        eventTriggerBytes,
			actionID:            100000,
			expectedError:       e2errors.IsProtocolAbstractSyntaxErrorFalselyConstructedMessage,
		},
		{
			description:         "Invalid event trigger",
			enabled:             true,
			encodingType:        sdkclient.ProtoEncoding,
			actionType:          e2api.ActionType_ACTION_TYPE_REPORT,
			serviceModelName:    utils.KpmServiceModelName,
			serviceModelVersion: utils.Version2,
			eventTrigger:        make([]byte, 50),
			actionID:            actionID,
			expectedError:       errors.IsInvalid,
		},
	}

	sim := s.CreateRanSimulatorWithNameOrDie("invalid-subscriptions")

	for _, testCase := range testCases {
		pinTestCase := testCase
		s.T().Run(pinTestCase.description, func(t *testing.T) {
			s.runTestCase(pinTestCase)
		})
	}
	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "invalid-subscriptions")
}
