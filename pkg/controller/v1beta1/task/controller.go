// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/topo"

	subscription "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/ranfunctions"

	"github.com/onosproject/onos-e2t/pkg/oid"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	storeapi "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	"github.com/onosproject/onos-e2t/pkg/config"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	taskstore "github.com/onosproject/onos-e2t/pkg/store/task"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "subscription")

// NewController returns a new network controller
func NewController(streams subscription.Broker, subs substore.Store,
	tasks taskstore.Store, channels e2server.ChannelManager,
	models modelregistry.ModelRegistry,
	oidRegistry oid.Registry,
	ranFunctionRegistry ranfunctions.Registry, topoManager topo.Manager) *controller.Controller {
	c := controller.NewController("Task")
	c.Watch(&Watcher{
		tasks: tasks,
	})
	c.Watch(&ChannelWatcher{
		subs:     subs,
		tasks:    tasks,
		channels: channels,
	})
	c.Reconcile(&Reconciler{
		streams:                   streams,
		subs:                      subs,
		tasks:                     tasks,
		channels:                  channels,
		models:                    models,
		oidRegistry:               oidRegistry,
		newRicSubscriptionRequest: pdubuilder.NewRicSubscriptionRequest,
		ranFunctionRegistry:       ranFunctionRegistry,
		topoManager:               topoManager,
	})
	return c
}

type RicSubscriptionRequestBuilder func(ricReq types.RicRequest,
	ranFuncID types.RanFunctionID, ricEventDef types.RicEventDefintion,
	ricActionsToBeSetup map[types.RicActionID]types.RicActionDef) (
	*e2appducontents.RicsubscriptionRequest, error)

// Reconciler is a device change reconciler
type Reconciler struct {
	streams                   subscription.Broker
	subs                      substore.Store
	tasks                     taskstore.Store
	channels                  e2server.ChannelManager
	models                    modelregistry.ModelRegistry
	oidRegistry               oid.Registry
	newRicSubscriptionRequest RicSubscriptionRequestBuilder
	ranFunctionRegistry       ranfunctions.Registry
	topoManager               topo.Manager
}

// Reconcile reconciles the state of a device change
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	taskID := id.Value.(storeapi.TaskID)
	task, err := r.tasks.Get(ctx, taskID)
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	log.Infof("Reconciling SubscriptionTask %+v", task)

	// If the task is COMPLETE or FAILED, ignore the request
	switch task.Status.State {
	case storeapi.TaskState_TASK_COMPLETE, storeapi.TaskState_TASK_FAILED:
		return controller.Result{}, nil
	}

	// Process the request based on the lifecycle phase
	switch task.Status.Phase {
	case storeapi.TaskPhase_OPEN:
		log.Infof("Opening SubscriptionTask %+v", task)
		return r.reconcileOpenSubscriptionTask(task)
	case storeapi.TaskPhase_CLOSE:
		log.Infof("Closing SubscriptionTask %+v", task)
		return r.reconcileCloseSubscriptionTask(task)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileOpenSubscriptionTask(task *storeapi.Task) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID, err := r.topoManager.GetE2Relation(ctx, topoapi.ID(task.ID.NodeID))
	if err != nil || channelID == "" {
		return controller.Result{}, err
	}
	channel, err := r.channels.Get(ctx, e2server.ChannelID(channelID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry,
		string(task.ServiceModel.Name),
		string(task.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		task.Status.State = storeapi.TaskState_TASK_FAILED
		task.Status.Failure = &storeapi.TaskFailure{
			Cause:   storeapi.TaskFailureCause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID,
			Message: fmt.Sprintf("Service Model Plugin %s cannot be loaded", serviceModelOID),
		}
		err = r.tasks.Update(ctx, task)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	serviceModelPlugin, err := r.models.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warn(err)
		task.Status.State = storeapi.TaskState_TASK_FAILED
		task.Status.Failure = &storeapi.TaskFailure{
			Cause:   storeapi.TaskFailureCause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID,
			Message: fmt.Sprintf("Service Model Plugin %s cannot be loaded", serviceModelOID),
		}
		err = r.tasks.Update(ctx, task)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	smData := serviceModelPlugin.ServiceModelData()
	log.Infof("Service model found %s %s %s", smData.Name, smData.Version, smData.OID)

	stream, ok := r.streams.GetReader(task.ID)
	if !ok {
		return controller.Result{}, errors.NewNotFound("stream not found for task %s", task.ID)
	}

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(stream.ID()),
		InstanceID:  config.InstanceID,
	}

	ranFunction, err := r.ranFunctionRegistry.Get(ranfunctions.NewID(serviceModelOID, string(task.ID.NodeID)))
	if err != nil {
		log.Warn(err)
	}

	ricEventDef := types.RicEventDefintion(task.Spec.Subscription.EventTrigger.Payload)

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	for _, action := range task.Spec.Subscription.Actions {
		ricActionsToBeSetup[types.RicActionID(action.ID)] = types.RicActionDef{
			RicActionID:         types.RicActionID(action.ID),
			RicActionType:       e2apies.RicactionType(action.Type),
			RicSubsequentAction: e2apies.RicsubsequentActionType(action.SubsequentAction.Type),
			Ricttw:              e2apies.RictimeToWait(action.SubsequentAction.TimeToWait),
			RicActionDefinition: action.Payload,
		}
	}

	request, err := r.newRicSubscriptionRequest(ricRequest, ranFunction.ID, ricEventDef, ricActionsToBeSetup)
	if err != nil {
		log.Warnf("Failed to create E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	}

	// Validate the subscribe request
	if err := request.Validate(); err != nil {
		log.Warnf("Failed to validate E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		task.Status.State = storeapi.TaskState_TASK_FAILED
		task.Status.Failure = &storeapi.TaskFailure{
			Cause: storeapi.TaskFailureCause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
		}
		err := r.tasks.Update(ctx, task)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// Send the subscription request and await a response
	response, failure, err := channel.RICSubscription(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	} else if response != nil {
		task.Status.State = storeapi.TaskState_TASK_COMPLETE
		err := r.tasks.Update(ctx, task)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	} else if failure != nil {
		log.Warnf("Failed to initialize SubscriptionTask %+v: %s", task, err)
		task.Status.State = storeapi.TaskState_TASK_FAILED
		task.Status.Failure = &storeapi.TaskFailure{
			Cause: getTaskFailureCause(failure),
		}
		err := r.tasks.Update(ctx, task)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileCloseSubscriptionTask(task *storeapi.Task) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID, err := r.topoManager.GetE2Relation(ctx, topoapi.ID(task.ID.NodeID))
	if err != nil || channelID == "" {
		return controller.Result{}, err
	}
	channel, err := r.channels.Get(ctx, e2server.ChannelID(channelID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	stream, ok := r.streams.GetReader(task.ID)
	if !ok {
		return controller.Result{}, errors.NewNotFound("stream not found for task %s", task.ID)
	}

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(stream.ID()),
		InstanceID:  config.InstanceID,
	}

	serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry, string(task.ServiceModel.Name), string(task.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		return controller.Result{}, err
	}

	ranFunction, err := r.ranFunctionRegistry.Get(ranfunctions.NewID(serviceModelOID, string(task.ID.NodeID)))
	if err != nil {
		log.Warn(err)
	}

	request, err := pdubuilder.NewRicSubscriptionDeleteRequest(ricRequest, ranFunction.ID)
	if err != nil {
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, failure, err := channel.RICSubscriptionDelete(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	} else if response != nil {
		task.Status.State = storeapi.TaskState_TASK_COMPLETE
		err := r.tasks.Update(ctx, task)
		if err != nil {
			return controller.Result{}, err
		}
		_ = stream.Close()
		return controller.Result{}, nil
	} else if failure != nil {
		switch failure.ProtocolIes.E2ApProtocolIes1.Value.Cause.(type) {
		case *e2apies.Cause_RicRequest:
			return controller.Result{}, nil
		}
		log.Warnf("SubscriptionDeleteRequest %+v failed: %+v", request, failure)
		return controller.Result{}, fmt.Errorf("failed to delete task %+v", task)
	}
	return controller.Result{}, nil
}

func getTaskFailureCause(failure *e2appducontents.RicsubscriptionFailure) storeapi.TaskFailureCause {
	for _, item := range failure.ProtocolIes.E2ApProtocolIes18.Value.Value {
		switch c := item.Value.Cause.Cause.(type) {
		case *e2apies.Cause_RicRequest:
			switch c.RicRequest {
			case e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID:
				return storeapi.TaskFailureCause_CAUSE_RIC_RAN_FUNCTION_ID_INVALID
			case e2apies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED:
				return storeapi.TaskFailureCause_CAUSE_RIC_ACTION_NOT_SUPPORTED
			case e2apies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS:
				return storeapi.TaskFailureCause_CAUSE_RIC_EXCESSIVE_ACTIONS
			case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION:
				return storeapi.TaskFailureCause_CAUSE_RIC_DUPLICATE_ACTION
			case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT:
				return storeapi.TaskFailureCause_CAUSE_RIC_DUPLICATE_EVENT
			case e2apies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT:
				return storeapi.TaskFailureCause_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT
			case e2apies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN:
				return storeapi.TaskFailureCause_CAUSE_RIC_REQUEST_ID_UNKNOWN
			case e2apies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
				return storeapi.TaskFailureCause_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
			case e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID:
				return storeapi.TaskFailureCause_CAUSE_RIC_CONTROL_MESSAGE_INVALID
			case e2apies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID:
				return storeapi.TaskFailureCause_CAUSE_RIC_CALL_PROCESS_ID_INVALID
			case e2apies.CauseRic_CAUSE_RIC_UNSPECIFIED:
				return storeapi.TaskFailureCause_CAUSE_RIC_UNSPECIFIED
			}
		case *e2apies.Cause_RicService:
			switch c.RicService {
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED:
				return storeapi.TaskFailureCause_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS:
				return storeapi.TaskFailureCause_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT:
				return storeapi.TaskFailureCause_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT
			}
		case *e2apies.Cause_Protocol:
			switch c.Protocol {
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_SEMANTIC_ERROR
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
				return storeapi.TaskFailureCause_CAUSE_PROTOCOL_UNSPECIFIED
			}
		case *e2apies.Cause_Transport:
			switch c.Transport {
			case e2apies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED:
				return storeapi.TaskFailureCause_CAUSE_TRANSPORT_UNSPECIFIED
			case e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE:
				return storeapi.TaskFailureCause_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE
			}
		case *e2apies.Cause_Misc:
			switch c.Misc {
			case e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD:
				return storeapi.TaskFailureCause_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD
			case e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE:
				return storeapi.TaskFailureCause_CAUSE_MISC_HARDWARE_FAILURE
			case e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION:
				return storeapi.TaskFailureCause_CAUSE_MISC_OM_INTERVENTION
			case e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED:
				return storeapi.TaskFailureCause_CAUSE_MISC_UNSPECIFIED
			}
		}
	}
	return 0
}
