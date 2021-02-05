// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	epapi "github.com/onosproject/onos-api/go/onos/e2sub/endpoint"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
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

	// If the task is COMPLETE, ignore the request
	if task.Lifecycle.Status == subtaskapi.Status_COMPLETE {
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
	serviceModelPlugin, ok := r.models.ModelPlugins[serviceModelID]
	if !ok {
		log.Errorf("Service Model Plugin cannot be loaded %s", serviceModelID)
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

	eventTriggerBytes := sub.Details.EventTrigger.Payload.Data
	if sub.Details.EventTrigger.Payload.Encoding == subapi.Encoding_ENCODING_PROTO {
		bytes, err := serviceModelPlugin.EventTriggerDefinitionProtoToASN1(eventTriggerBytes)
		if err != nil {
			log.Errorf("Error transforming Proto bytes to ASN: %s", err.Error())
			return controller.Result{}, nil
		}
		eventTriggerBytes = bytes
	}

	ricEventDef := types.RicEventDefintion(eventTriggerBytes)

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	for _, action := range sub.Details.Actions {
		actionBytes := action.Payload.Data
		if action.Payload.Encoding == subapi.Encoding_ENCODING_PROTO {
			bytes, err := serviceModelPlugin.ActionDefinitionProtoToASN1(actionBytes)
			if err != nil {
				log.Errorf("Error transforming Proto bytes to ASN: %s", err.Error())
				return controller.Result{}, nil
			}
			actionBytes = bytes
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
	ranFunctionID := types.RanFunctionID(1)

	request, err := pdubuilder.NewRicSubscriptionDeleteRequest(ricRequest, ranFunctionID)
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
