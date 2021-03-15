// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	"time"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	epapi "github.com/onosproject/onos-api/go/onos/e2sub/endpoint"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "subscription")

// NewController returns a new network controller
func NewController(catalog *RequestJournal, subs subapi.E2SubscriptionServiceClient, tasks subtaskapi.E2SubscriptionTaskServiceClient, channels *e2server.ChannelManager, models *modelregistry.ModelRegistry) *controller.Controller {
	c := controller.NewController("SubscriptionTask")
	c.Watch(&Watcher{
		endpointID: epapi.ID(env.GetPodID()),
		tasks:      tasks,
	})
	c.Watch(&ChannelWatcher{
		endpointID: epapi.ID(env.GetPodID()),
		subs:       subs,
		tasks:      tasks,
		channels:   channels,
	})
	c.Reconcile(&Reconciler{
		catalog:  catalog,
		subs:     subs,
		tasks:    tasks,
		channels: channels,
		models:   models,
	})
	return c
}

// Reconciler is a device change reconciler
type Reconciler struct {
	catalog   *RequestJournal
	subs      subapi.E2SubscriptionServiceClient
	tasks     subtaskapi.E2SubscriptionTaskServiceClient
	channels  *e2server.ChannelManager
	models    *modelregistry.ModelRegistry
	requestID RequestID
}

// Reconcile reconciles the state of a device change
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	taskRequest := &subtaskapi.GetSubscriptionTaskRequest{
		ID: id.Value.(subtaskapi.ID),
	}
	taskResponse, err := r.tasks.GetSubscriptionTask(ctx, taskRequest)
	if err != nil {
		if stat, ok := status.FromError(err); ok && errors.IsNotFound(errors.FromStatus(stat)) {
			return controller.Result{}, nil
		}
		return controller.Result{}, err
	}
	task := taskResponse.Task

	log.Infof("Reconciling SubscriptionTask %+v", task)

	// If the task is COMPLETE or FAILED, ignore the request
	switch task.Lifecycle.Status {
	case subtaskapi.Status_COMPLETE, subtaskapi.Status_FAILED:
		return controller.Result{}, nil
	}

	// Process the request based on the lifecycle phase
	switch task.Lifecycle.Phase {
	case subtaskapi.Phase_OPEN:
		log.Infof("Opening SubscriptionTask %+v", task)
		return r.reconcileOpenSubscriptionTask(task)
	case subtaskapi.Phase_CLOSE:
		log.Infof("Closing SubscriptionTask %+v", task)
		return r.reconcileCloseSubscriptionTask(task)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileOpenSubscriptionTask(task *subtaskapi.SubscriptionTask) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Get the subscription
	subRequest := &subapi.GetSubscriptionRequest{
		ID: task.SubscriptionID,
	}
	subResponse, err := r.subs.GetSubscription(ctx, subRequest)
	if err != nil {
		if stat, ok := status.FromError(err); ok && errors.IsNotFound(errors.FromStatus(stat)) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}
	sub := subResponse.Subscription

	channel, err := r.channels.Get(ctx, e2server.ChannelID(sub.Details.E2NodeID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	serviceModelID := modelregistry.ModelFullName(sub.Details.ServiceModel.ID)
	serviceModelPlugin, err := r.models.GetPlugin(serviceModelID)
	if err != nil {
		log.Warn(err)
		task.Lifecycle.Status = subtaskapi.Status_FAILED
		task.Lifecycle.Failure = &subtaskapi.Failure{
			Cause:   subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID,
			Message: subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID.String(),
		}
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		updateResponse, updateError := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if updateError != nil {
			log.Errorf("Unable to update subscription task for unknown service model. resp %v err %v", updateResponse, updateError)
		}
		return controller.Result{}, errors.NewInvalid("Service Model Plugin cannot be loaded", serviceModelID)
	}
	a, b, c := serviceModelPlugin.ServiceModelData()
	log.Infof("Service model found %s %s %s", a, b, c)

	r.requestID++
	requestID := r.requestID

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(requestID),
		InstanceID:  config.InstanceID,
	}

	ranFuncID := channel.GetRANFunctionID(serviceModelID)

	var eventTriggerBytes []byte
	if sub.Details.EventTrigger.Payload.Encoding == subapi.Encoding_ENCODING_ASN1 {
		eventTriggerBytes = sub.Details.EventTrigger.Payload.Data
	} else if sub.Details.EventTrigger.Payload.Encoding == subapi.Encoding_ENCODING_PROTO {
		eventTriggerBytes = sub.Details.EventTrigger.Payload.Data
		bytes, err := serviceModelPlugin.EventTriggerDefinitionProtoToASN1(eventTriggerBytes)
		if err != nil {
			log.Warnf("Error transforming Proto bytes to ASN: %s", err.Error())
			cause := subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
			task.Lifecycle.Status = subtaskapi.Status_FAILED
			task.Lifecycle.Failure = &subtaskapi.Failure{
				Cause:   cause,
				Message: cause.String(),
			}
			updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
				Task: task,
			}
			_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
			if err != nil {
				log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
				return controller.Result{}, err
			}

			return controller.Result{}, nil
		}
		eventTriggerBytes = bytes
	} else {
		cause := subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
		task.Lifecycle.Status = subtaskapi.Status_FAILED
		task.Lifecycle.Failure = &subtaskapi.Failure{
			Cause:   cause,
			Message: cause.String(),
		}
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if err != nil {
			log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
			return controller.Result{}, err
		}
		return controller.Result{}, fmt.Errorf("failed to initialize subscription %+v", sub)
	}

	ricEventDef := types.RicEventDefintion(eventTriggerBytes)

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	for _, action := range sub.Details.Actions {
		var actionBytes []byte
		if action.Payload.Encoding == subapi.Encoding_ENCODING_ASN1 {
			actionBytes = action.Payload.Data
		} else if action.Payload.Encoding == subapi.Encoding_ENCODING_PROTO {
			actionBytes = action.Payload.Data
			bytes, err := serviceModelPlugin.ActionDefinitionProtoToASN1(actionBytes)
			if err != nil {
				log.Warnf("Error transforming Proto bytes to ASN: %s", err.Error())
				cause := subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
				task.Lifecycle.Status = subtaskapi.Status_FAILED
				task.Lifecycle.Failure = &subtaskapi.Failure{
					Cause:   cause,
					Message: cause.String(),
				}
				updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
					Task: task,
				}
				_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
				if err != nil {
					log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
					return controller.Result{}, err
				}

				return controller.Result{}, nil
			}
			actionBytes = bytes
		} else {
			cause := subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
			task.Lifecycle.Status = subtaskapi.Status_FAILED
			task.Lifecycle.Failure = &subtaskapi.Failure{
				Cause:   cause,
				Message: cause.String(),
			}
			updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
				Task: task,
			}
			_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
			if err != nil {
				log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
				return controller.Result{}, err
			}
			return controller.Result{}, fmt.Errorf("failed to initialize subscription %+v", sub)
		}

		ricActionsToBeSetup[types.RicActionID(action.ID)] = types.RicActionDef{
			RicActionID:         types.RicActionID(action.ID),
			RicActionType:       e2apies.RicactionType(action.Type),
			RicSubsequentAction: e2apies.RicsubsequentActionType(action.SubsequentAction.Type),
			Ricttw:              e2apies.RictimeToWait(action.SubsequentAction.TimeToWait),
			RicActionDefinition: types.RicActionDefinition(actionBytes),
		}
	}

	request, err := pdubuilder.NewRicSubscriptionRequest(ricRequest, ranFuncID, ricEventDef, ricActionsToBeSetup)
	if err != nil {
		log.Warnf("Failed to create E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	}

	// Validate the subscribe request
	if err := request.Validate(); err != nil {
		log.Warnf("Failed to validate E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		cause := subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
		task.Lifecycle.Status = subtaskapi.Status_FAILED
		task.Lifecycle.Failure = &subtaskapi.Failure{
			Cause:   cause,
			Message: cause.String(),
		}
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if err != nil {
			log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
			return controller.Result{}, err
		}
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, failure, err := channel.RICSubscription(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	} else if response != nil {
		record := RequestEntry{
			RequestID:    requestID,
			Subscription: *sub,
		}
		r.catalog.Add(sub.ID, record)

		task.Lifecycle.Status = subtaskapi.Status_COMPLETE
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if err != nil {
			log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
			return controller.Result{}, err
		}
	} else if failure != nil {
		log.Warnf("Failed to initialize SubscriptionTask %+v: %s", task, err)
		cause := getTaskFailureCause(failure)
		task.Lifecycle.Status = subtaskapi.Status_FAILED
		task.Lifecycle.Failure = &subtaskapi.Failure{
			Cause:   cause,
			Message: cause.String(),
		}
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if err != nil {
			log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
			return controller.Result{}, err
		}
		return controller.Result{}, fmt.Errorf("failed to initialize subscription %+v", sub)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileCloseSubscriptionTask(task *subtaskapi.SubscriptionTask) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Get the subscription
	subRequest := &subapi.GetSubscriptionRequest{
		ID: task.SubscriptionID,
	}
	subResponse, err := r.subs.GetSubscription(ctx, subRequest)
	if err != nil {
		if stat, ok := status.FromError(err); ok && errors.IsNotFound(errors.FromStatus(stat)) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}
	sub := subResponse.Subscription

	channel, err := r.channels.Get(ctx, e2server.ChannelID(sub.Details.E2NodeID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	record := r.catalog.Get(sub.ID)

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(record.RequestID),
		InstanceID:  config.InstanceID,
	}

	serviceModelID := modelregistry.ModelFullName(sub.Details.ServiceModel.ID)
	ranFuncID := channel.GetRANFunctionID(serviceModelID)

	request, err := pdubuilder.NewRicSubscriptionDeleteRequest(ricRequest, ranFuncID)
	if err != nil {
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, failure, err := channel.RICSubscriptionDelete(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	} else if response != nil {
		task.Lifecycle.Status = subtaskapi.Status_COMPLETE
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if err != nil {
			log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
			return controller.Result{}, err
		}
	} else if failure != nil {
		log.Warnf("Failed to initialize SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, fmt.Errorf("failed to delete subscription %+v", sub)
	}
	return controller.Result{}, nil
}

func getTaskFailureCause(failure *e2appducontents.RicsubscriptionFailure) subtaskapi.Cause {
	for _, item := range failure.ProtocolIes.E2ApProtocolIes18.Value.Value {
		switch c := item.Value.Cause.Cause.(type) {
		case *e2apies.Cause_RicRequest:
			switch c.RicRequest {
			case e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID:
				return subtaskapi.Cause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID
			case e2apies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED:
				return subtaskapi.Cause_CAUSE_RIC_ACTION_NOT_SUPPORTED
			case e2apies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS:
				return subtaskapi.Cause_CAUSE_RIC_EXCESSIVE_ACTIONS
			case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION:
				return subtaskapi.Cause_CAUSE_RIC_DUPLICATE_ACTION
			case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT:
				return subtaskapi.Cause_CAUSE_RIC_DUPLICATE_EVENT
			case e2apies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT:
				return subtaskapi.Cause_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT
			case e2apies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN:
				return subtaskapi.Cause_CAUSE_RIC_REQUEST_ID_UNKNOWN
			case e2apies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
				return subtaskapi.Cause_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
			case e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID:
				return subtaskapi.Cause_CAUSE_RIC_CONTROL_MESSAGE_INVALID
			case e2apies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID:
				return subtaskapi.Cause_CAUSE_RIC_CALL_PROCESS_ID_INVALID
			case e2apies.CauseRic_CAUSE_RIC_UNSPECIFIED:
				return subtaskapi.Cause_CAUSE_RIC_UNSPECIFIED
			}
		case *e2apies.Cause_RicService:
			switch c.RicService {
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED:
				return subtaskapi.Cause_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS:
				return subtaskapi.Cause_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT:
				return subtaskapi.Cause_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT
			}
		case *e2apies.Cause_Protocol:
			switch c.Protocol {
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
				return subtaskapi.Cause_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
				return subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
				return subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
				return subtaskapi.Cause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
				return subtaskapi.Cause_CAUSE_PROTOCOL_SEMANTIC_ERROR
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
				return subtaskapi.Cause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
				return subtaskapi.Cause_CAUSE_PROTOCOL_UNSPECIFIED
			}
		case *e2apies.Cause_Transport:
			switch c.Transport {
			case e2apies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED:
				return subtaskapi.Cause_CAUSE_TRANSPORT_UNSPECIFIED
			case e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE:
				return subtaskapi.Cause_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE
			}
		case *e2apies.Cause_Misc:
			switch c.Misc {
			case e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD:
				return subtaskapi.Cause_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD
			case e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE:
				return subtaskapi.Cause_CAUSE_MISC_HARDWARE_FAILURE
			case e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION:
				return subtaskapi.Cause_CAUSE_MISC_OM_INTERVENTION
			case e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED:
				return subtaskapi.Cause_CAUSE_MISC_UNSPECIFIED
			}
		}
	}
	return 0
}
