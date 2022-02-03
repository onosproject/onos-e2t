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
	sdkClient := utils.GetE2Client(t, string(testCase.serviceModelName), string(testCase.serviceModelVersion), testCase.encodingType)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)
	var actions []e2api.Action
	action := e2api.Action{
		ID:   testCase.actionID,
		Type: testCase.actionType,
		SubsequentAction: &e2api.SubsequentAction{
			Type:       e2api.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: e2api.TimeToWait_TIME_TO_WAIT_ZERO,
		},
	}
	actions = append(actions, action)

	subRequest := utils.Subscription{
		NodeID:              string(nodeID),
		Actions:             actions,
		EventTrigger:        testCase.eventTrigger,
		ServiceModelName:    testCase.serviceModelName,
		ServiceModelVersion: testCase.serviceModelVersion,
	}

	subSpec, err := subRequest.CreateWithActionDefinition()
	assert.NoError(t, err)

	ch := make(chan e2api.Indication)
	node := sdkClient.Node(sdkclient.NodeID(nodeID))
	_, err = node.Subscribe(ctx, testCase.description, subSpec, ch)
	assert.Error(t, err)

	assert.True(t, testCase.expectedError(err))

	_ = node.Unsubscribe(ctx, testCase.description)

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
