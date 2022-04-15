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
	"github.com/stretchr/testify/assert"
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

func runTestCase(t *testing.T, testCase invalidSubscriptionTestCase) {
	if !testCase.enabled {
		t.Skip()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	cellObjectID := e2utils.GetFirstCellObjectID(t, nodeID)

	// Create a KPM V2 subscription
	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:                testCase.description,
			NodeID:              nodeID,
			ActionID:            testCase.actionID,
			ActionType:          testCase.actionType,
			EventTriggerBytes:   testCase.eventTrigger,
			ServiceModelName:    testCase.serviceModelName,
			ServiceModelVersion: testCase.serviceModelVersion,
			EncodingType:        testCase.encodingType,
		},
		CellObjectID: cellObjectID,
	}
	_, err := kpmv2Sub.Subscribe(ctx)

	assert.True(t, testCase.expectedError(err))

	kpmv2Sub.UnsubscribeOrFail(ctx, t)

}

// TestInvalidSubscriptions tests invalid inputs into the SDK
func (s *TestSuite) TestInvalidSubscriptions(t *testing.T) {
	const actionID = 11
	eventTriggerBytes, err := utils.CreateKpmV2EventTrigger(1000)
	assert.NoError(t, err)

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

	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "invalid-subscriptions")

	for _, testCase := range testCases {
		pinTestCase := testCase
		t.Run(pinTestCase.description, func(t *testing.T) {
			runTestCase(t, pinTestCase)
		})
	}
	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
