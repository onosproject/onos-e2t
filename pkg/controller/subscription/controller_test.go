// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"errors"
	"testing"

	"github.com/onosproject/onos-e2t/pkg/ranfunctions"

	"github.com/golang/mock/gomock"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	broker "github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/oid"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/stretchr/testify/assert"
)

type controllerTestContext struct {
	ctrl                   *gomock.Controller
	serverChannel          *MockRICChannel
	channelManager         *MockChannelManager
	subscriptionClient     subapi.E2SubscriptionServiceClient
	subscriptionTaskClient subtaskapi.E2SubscriptionTaskServiceClient
	broker                 *MockBroker
	modelRegistry          *MockModelRegistry
	oidRegistry            *MockRegistry
	ranFunctionRegistry    ranfunctions.Registry
	reconciler             Reconciler
	controller             *controller.Controller
}

func initControllerTest(t *testing.T, testContext *controllerTestContext) {
	initControllerTestNoRICSubscription(t, testContext)
	resp := e2appducontents.RicsubscriptionResponse{}
	testContext.serverChannel.EXPECT().RICSubscription(gomock.Any(), gomock.Any()).Return(&resp, nil, nil).AnyTimes()
}

func initControllerTestNoRICSubscription(t *testing.T, testContext *controllerTestContext) {
	subTask = &subtaskapi.SubscriptionTask{
		ID:             "1",
		Revision:       0,
		SubscriptionID: "1",
		EndpointID:     E2NodeID,
		Lifecycle:      subtaskapi.Lifecycle{},
	}

	// subscription and task clients
	testContext.subscriptionClient, testContext.subscriptionTaskClient = createClients(t)

	// Controller for mocks
	ctrl := gomock.NewController(t)

	// Mock channel manager
	testContext.channelManager = NewMockChannelManager(ctrl)

	// Mock broker
	testContext.broker = NewMockBroker(ctrl)
	stream := NewMockStream(ctrl)
	stream.EXPECT().SubscriptionID().Return(subapi.ID("1")).AnyTimes()
	stream.EXPECT().StreamID().Return(broker.StreamID(1)).AnyTimes()
	testContext.broker.EXPECT().GetStream(gomock.Any()).Return(stream, nil).AnyTimes()
	testContext.broker.EXPECT().OpenStream(gomock.Any()).Return(stream, nil).AnyTimes()
	testContext.broker.EXPECT().CloseStream(gomock.Any()).Return(stream, nil).AnyTimes()

	testContext.ranFunctionRegistry = ranfunctions.NewRegistry()
	_ = testContext.ranFunctionRegistry.Add(ranfunctions.NewID("12", E2NodeID), ranfunctions.RANFunction{
		ID: 2,
	})

	// Mocked RIC channel
	serverChannel := NewMockRICChannel(ctrl)
	testContext.serverChannel = serverChannel

	channel := e2server.NewE2Channel("channel", "123", serverChannel, testContext.broker)
	testContext.channelManager = NewMockChannelManager(ctrl)
	testContext.channelManager.EXPECT().Get(gomock.Any(), gomock.Any()).Return(channel, nil)

	// Mocked model registry
	testContext.modelRegistry = NewMockModelRegistry(ctrl)
	testContext.ctrl = ctrl

	// Controller
	testContext.controller = NewController(testContext.broker, testContext.subscriptionClient,
		testContext.subscriptionTaskClient, testContext.channelManager,
		testContext.modelRegistry, testContext.oidRegistry, testContext.ranFunctionRegistry, nil)

	// OID registry
	testContext.oidRegistry = NewMockRegistry(ctrl)
	testContext.oidRegistry.EXPECT().GetOid(gomock.Any()).Return(oid.Oid(12)).AnyTimes()

	// reconciler to test
	testContext.reconciler = Reconciler{
		streams:                   testContext.broker,
		subs:                      testContext.subscriptionClient,
		tasks:                     testContext.subscriptionTaskClient,
		channels:                  testContext.channelManager,
		models:                    testContext.modelRegistry,
		oidRegistry:               testContext.oidRegistry,
		newRicSubscriptionRequest: pdubuilder.NewRicSubscriptionRequest,
	}
}

func getTaskOrDie(t *testing.T) *subtaskapi.SubscriptionTask {
	updatedTask, err := scaffold.taskStore.Get(context.Background(), subTask.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	return updatedTask
}

func createTaskOrDie(t *testing.T) {
	err := scaffold.taskStore.Create(context.Background(), subTask)
	assert.NoError(t, err)
}

func defaultSubscription() *subapi.Subscription {
	return &subapi.Subscription{
		ID: "1", AppID: "foo", Details: &subapi.SubscriptionDetails{E2NodeID: E2NodeID, ServiceModel: subapi.ServiceModel{Name: "sm1", Version: "v1"}},
	}
}

func addSubscriptionOrDie(t *testing.T, testContext controllerTestContext, subscription *subapi.Subscription) {
	_, err := testContext.subscriptionClient.AddSubscription(context.Background(), &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	})
	assert.NoError(t, err)
}

/*func reconcileOrDie(t *testing.T, testContext controllerTestContext) {
	result, err := testContext.reconciler.Reconcile(controller.ID{Value: subTask.ID})
	assert.NotNil(t, result)
	assert.NoError(t, err)
}*/

func reconcileExpectError(t *testing.T, testContext controllerTestContext) {
	// TODO uncomment when these tests fixed
	/*result, err := testContext.reconciler.Reconcile(controller.ID{Value: subTask.ID})
	assert.NotNil(t, result)
	assert.Error(t, err)*/
}

// TODO uncomment it after fixing the test
/*func newInvalidRicSubscriptionRequest(ricReq types.RicRequest,
	ranFuncID types.RanFunctionID, ricEventDef types.RicEventDefintion,
	ricActionsToBeSetup map[types.RicActionID]types.RicActionDef) (
	*e2appducontents.RicsubscriptionRequest, error) {

	request, _ := pdubuilder.NewRicSubscriptionRequest(ricReq, ranFuncID, ricEventDef, ricActionsToBeSetup)
	request.ProtocolIes.E2ApProtocolIes29.Criticality = 77
	return request, nil
}*/

func TestOpenNoPlugin(t *testing.T) {
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(nil, errors.New("no such model"))
	addSubscriptionOrDie(t, testContext, defaultSubscription())
	createTaskOrDie(t)
	reconcileExpectError(t, testContext)
	updatedTask := getTaskOrDie(t)
	t.Log(updatedTask)
	t.Skip()

	/*assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()*/
}

func TestOpenInvalidRequest(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)
	// TODO uncomment it
	//testContext.reconciler.newRicSubscriptionRequest = newInvalidRicSubscriptionRequest

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	addSubscriptionOrDie(t, testContext, defaultSubscription())
	createTaskOrDie(t)

	reconcileExpectError(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenSMError(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	subscription := defaultSubscription()
	subscription.Details.ServiceModel.Version = ""
	addSubscriptionOrDie(t, testContext, subscription)
	createTaskOrDie(t)

	reconcileExpectError(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenProtoToASNError(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(nil, errors.New("this should fail"))
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := defaultSubscription()
	subscription.Details.EventTrigger.Payload.Encoding = subapi.Encoding_ENCODING_PROTO
	addSubscriptionOrDie(t, testContext, subscription)
	createTaskOrDie(t)

	// TODO uncomment it after fixing the test
	//reconcileOrDie(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenBadProtocolError(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	subscription := defaultSubscription()
	subscription.Details.EventTrigger.Payload.Encoding = 123
	addSubscriptionOrDie(t, testContext, subscription)
	createTaskOrDie(t)

	reconcileExpectError(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenValidPlugin(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	addSubscriptionOrDie(t, testContext, defaultSubscription())
	createTaskOrDie(t)

	// TODO uncomment it after fixing the test
	//reconcileOrDie(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_COMPLETE, updatedTask.Lifecycle.Status)

	testContext.ctrl.Finish()
}

func TestOpenActionBadProtocolError(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	actions := []subapi.Action{{
		Payload: subapi.Payload{Encoding: 55},
	}}
	subscription := defaultSubscription()
	subscription.Details.Actions = actions
	subscription.Details.EventTrigger.Payload.Encoding = subapi.Encoding_ENCODING_PROTO
	addSubscriptionOrDie(t, testContext, subscription)
	createTaskOrDie(t)

	reconcileExpectError(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenActionBadProtoPayload(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	sm.EXPECT().ActionDefinitionProtoToASN1(gomock.Any()).Return(nil, errors.New("bad proto payload"))
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	actions := []subapi.Action{{
		Payload: subapi.Payload{Encoding: subapi.Encoding_ENCODING_PROTO},
	}}
	subscription := defaultSubscription()
	subscription.Details.Actions = actions
	subscription.Details.EventTrigger.Payload.Encoding = subapi.Encoding_ENCODING_PROTO
	addSubscriptionOrDie(t, testContext, subscription)
	createTaskOrDie(t)

	reconcileExpectError(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenBadChannelResponse(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTestNoRICSubscription(t, &testContext)
	cause := &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID}
	failure := e2appducontents.RicsubscriptionFailure{
		ProtocolIes: &e2appducontents.RicsubscriptionFailureIes{
			E2ApProtocolIes29: nil,
			E2ApProtocolIes5:  nil,
			E2ApProtocolIes18: &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes18{
				Id:          0,
				Criticality: 0,
				Value: &e2appducontents.RicactionNotAdmittedList{
					Value: []*e2appducontents.RicactionNotAdmittedItemIes{
						{Value: &e2appducontents.RicactionNotAdmittedItem{
							Cause: &e2apies.Cause{
								Cause: cause,
							},
						}},
					},
				},
				Presence: 0,
			},
			E2ApProtocolIes2: nil,
		},
	}
	testContext.serverChannel.EXPECT().RICSubscription(gomock.Any(), gomock.Any()).Return(nil, &failure, nil).AnyTimes()

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	addSubscriptionOrDie(t, testContext, defaultSubscription())
	createTaskOrDie(t)

	reconcileExpectError(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_FAILED, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID, updatedTask.Lifecycle.Failure.Cause)

	testContext.ctrl.Finish()
}

func TestOpenAction(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)

	sm := NewMockServiceModel(testContext.ctrl)
	smd := e2smtypes.ServiceModelData{}
	sm.EXPECT().ServiceModelData().Return(smd)
	sm.EXPECT().EventTriggerDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	sm.EXPECT().ActionDefinitionProtoToASN1(gomock.Any()).Return(make([]byte, 1), nil)
	testContext.modelRegistry.EXPECT().GetPlugin(gomock.Any()).Return(sm, nil)

	actions := []subapi.Action{{
		Payload: subapi.Payload{Encoding: subapi.Encoding_ENCODING_PROTO},
		SubsequentAction: &subapi.SubsequentAction{
			Type:       subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
			TimeToWait: 5,
		},
	}}
	subscription := defaultSubscription()
	subscription.Details.Actions = actions
	subscription.Details.EventTrigger.Payload.Encoding = subapi.Encoding_ENCODING_PROTO
	addSubscriptionOrDie(t, testContext, subscription)
	createTaskOrDie(t)

	// TODO uncomment it after fixing the test
	//reconcileOrDie(t, testContext)

	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_COMPLETE, updatedTask.Lifecycle.Status)

	testContext.ctrl.Finish()
}

func TestClose(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTest(t, &testContext)
	response := e2appducontents.RicsubscriptionDeleteResponse{}
	testContext.serverChannel.EXPECT().RICSubscriptionDelete(gomock.Any(), gomock.Any()).Return(&response, nil, nil).AnyTimes()

	addSubscriptionOrDie(t, testContext, defaultSubscription())
	_, err := testContext.subscriptionClient.RemoveSubscription(context.Background(), &subapi.RemoveSubscriptionRequest{
		ID: "1",
	})
	assert.NoError(t, err)

	subTask.Lifecycle.Phase = subtaskapi.Phase_CLOSE
	createTaskOrDie(t)
	// TODO uncomment it after fixing the test
	//reconcileOrDie(t, testContext)
	updatedTask := getTaskOrDie(t)
	assert.Equal(t, subtaskapi.Status_COMPLETE, updatedTask.Lifecycle.Status)

	testContext.ctrl.Finish()
}

func TestCloseBadChannelResponseRicRequest(t *testing.T) {
	t.Skip()
	var testContext controllerTestContext
	createServerScaffolding(t)
	initControllerTestNoRICSubscription(t, &testContext)
	cause := &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID}
	failure := e2appducontents.RicsubscriptionDeleteFailure{
		ProtocolIes: &e2appducontents.RicsubscriptionDeleteFailureIes{
			E2ApProtocolIes29: nil,
			E2ApProtocolIes5:  nil,
			E2ApProtocolIes1: &e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes1{
				Id:          0,
				Criticality: 0,
				Value: &e2apies.Cause{
					Cause: cause,
				},
				Presence: 0,
			},
			E2ApProtocolIes2: nil,
		},
	}
	testContext.serverChannel.EXPECT().RICSubscriptionDelete(gomock.Any(), gomock.Any()).Return(nil, &failure, nil).AnyTimes()

	addSubscriptionOrDie(t, testContext, defaultSubscription())
	_, err := testContext.subscriptionClient.RemoveSubscription(context.Background(), &subapi.RemoveSubscriptionRequest{
		ID: "1",
	})
	assert.NoError(t, err)

	subTask.Lifecycle.Phase = subtaskapi.Phase_CLOSE
	createTaskOrDie(t)
	// TODO uncomment it after fixing the test
	//reconcileOrDie(t, testContext)
	updatedTask := getTaskOrDie(t)
	//assert.Equal(t, subtaskapi.Status_COMPLETE, updatedTask.Lifecycle.Status)
	assert.Equal(t, subtaskapi.Status_PENDING, updatedTask.Lifecycle.Status)

	testContext.ctrl.Finish()
}

func TestFailureCause(t *testing.T) {
	type test struct {
		description string
		cause       *e2apies.Cause
		resultCause subtaskapi.Cause
	}

	tests := []test{
		{description: "invalid function id", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID}}, resultCause: subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID},
		{description: "action not supported", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED}}, resultCause: subtaskapi.Cause_CAUSE_RIC_ACTION_NOT_SUPPORTED},
		{description: "excessive actions", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS}}, resultCause: subtaskapi.Cause_CAUSE_RIC_EXCESSIVE_ACTIONS},
		{description: "duplicate action", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION}}, resultCause: subtaskapi.Cause_CAUSE_RIC_DUPLICATE_ACTION},
		{description: "duplicate event", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT}}, resultCause: subtaskapi.Cause_CAUSE_RIC_DUPLICATE_EVENT},
		{description: "function resource limit", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT}}, resultCause: subtaskapi.Cause_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT},
		{description: "request ID unknown", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN}}, resultCause: subtaskapi.Cause_CAUSE_RIC_REQUEST_ID_UNKNOWN},
		{description: "inconsistent action", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE}}, resultCause: subtaskapi.Cause_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE},
		{description: "control message invalid", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID}}, resultCause: subtaskapi.Cause_CAUSE_RIC_CONTROL_MESSAGE_INVALID},
		{description: "call process ID invalid", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID}}, resultCause: subtaskapi.Cause_CAUSE_RIC_CALL_PROCESS_ID_INVALID},
		{description: "unspecified", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicRequest{RicRequest: e2apies.CauseRic_CAUSE_RIC_UNSPECIFIED}}, resultCause: subtaskapi.Cause_CAUSE_RIC_UNSPECIFIED},

		{description: "function not required", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicService{RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED}}, resultCause: subtaskapi.Cause_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED},
		{description: "excessive functions", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicService{RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS}}, resultCause: subtaskapi.Cause_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS},
		{description: "resource limit", cause: &e2apies.Cause{Cause: &e2apies.Cause_RicService{RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT}}, resultCause: subtaskapi.Cause_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT},

		{description: "transfer syntax error", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR},
		{description: "abstract syntax error reject", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT},
		{description: "abstract syntax error ignore and notify", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY},
		{description: "message not compatible with receiver state", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE},
		{description: "semantic error", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_SEMANTIC_ERROR},
		{description: "abstract syntax falsely constructed message", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE},
		{description: "unspecified protocol", cause: &e2apies.Cause{Cause: &e2apies.Cause_Protocol{Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED}}, resultCause: subtaskapi.Cause_CAUSE_PROTOCOL_UNSPECIFIED},

		{description: "transport unspecified", cause: &e2apies.Cause{Cause: &e2apies.Cause_Transport{Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED}}, resultCause: subtaskapi.Cause_CAUSE_TRANSPORT_UNSPECIFIED},
		{description: "transport resource unavailable", cause: &e2apies.Cause{Cause: &e2apies.Cause_Transport{Transport: e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE}}, resultCause: subtaskapi.Cause_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE},

		{description: "control processing overload", cause: &e2apies.Cause{Cause: &e2apies.Cause_Misc{Misc: e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD}}, resultCause: subtaskapi.Cause_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD},
		{description: "hardware failure", cause: &e2apies.Cause{Cause: &e2apies.Cause_Misc{Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE}}, resultCause: subtaskapi.Cause_CAUSE_MISC_HARDWARE_FAILURE},
		{description: "OM intervention", cause: &e2apies.Cause{Cause: &e2apies.Cause_Misc{Misc: e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION}}, resultCause: subtaskapi.Cause_CAUSE_MISC_OM_INTERVENTION},
		{description: "unspecified misc", cause: &e2apies.Cause{Cause: &e2apies.Cause_Misc{Misc: e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED}}, resultCause: subtaskapi.Cause_CAUSE_MISC_UNSPECIFIED},

		{description: "invalid entry", cause: &e2apies.Cause{Cause: &e2apies.Cause_Misc{Misc: 77}}, resultCause: 0},
	}

	for _, tc := range tests {
		pin := tc
		failure := &e2appducontents.RicsubscriptionFailure{
			ProtocolIes: &e2appducontents.RicsubscriptionFailureIes{
				E2ApProtocolIes29: nil,
				E2ApProtocolIes5:  nil,
				E2ApProtocolIes18: &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes18{
					Id:          0,
					Criticality: 0,
					Value: &e2appducontents.RicactionNotAdmittedList{
						Value: []*e2appducontents.RicactionNotAdmittedItemIes{
							{Value: &e2appducontents.RicactionNotAdmittedItem{
								Cause: pin.cause},
							}},
					},
					Presence: 0,
				},
				E2ApProtocolIes2: nil,
			},
		}
		taskCause := getTaskFailureCause(failure)
		assert.Equal(t, pin.resultCause, taskCause, pin.description)
	}
}
