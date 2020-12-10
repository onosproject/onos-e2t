// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	epapi "github.com/onosproject/onos-api/go/onos/e2sub/endpoint"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	channelfilter "github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "subscription")

// NewController returns a new network controller
func NewController(catalog *RequestJournal, subs subapi.E2SubscriptionServiceClient, tasks subtaskapi.E2SubscriptionTaskServiceClient, channels *channel.Manager) *controller.Controller {
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
	})
	return c
}

// Reconciler is a device change reconciler
type Reconciler struct {
	catalog   *RequestJournal
	subs      subapi.E2SubscriptionServiceClient
	tasks     subtaskapi.E2SubscriptionTaskServiceClient
	channels  *channel.Manager
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

	channel, err := r.channels.Get(ctx, channel.ID(sub.E2NodeID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	r.requestID++
	requestID := r.requestID

	ricRequestID := &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
		Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(requestID),
			RicInstanceId:  config.InstanceID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionID := &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes5{
		Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: 0, // TODO: Map service model to RAN function ID
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	actions := make([]*e2appducontents.RicactionToBeSetupItemIes, len(sub.Details.Actions))
	for i, action := range sub.Details.Actions {
		var subsequentAction *e2apies.RicsubsequentAction
		if action.SubsequentAction != nil {
			subsequentAction = &e2apies.RicsubsequentAction{
				RicSubsequentActionType: e2apies.RicsubsequentActionType(action.SubsequentAction.Type),
				RicTimeToWait:           e2apies.RictimeToWait(action.SubsequentAction.TimeToWait),
			}
		}

		actions[i] = &e2appducontents.RicactionToBeSetupItemIes{
			Id:          int32(v1beta1.ProtocolIeIDRicactionToBeSetupItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RicactionToBeSetupItem{
				RicActionId: &e2apies.RicactionId{
					Value: int32(action.ID),
				},
				RicActionType:       e2apies.RicactionType(action.Type),
				RicActionDefinition: &e2ap_commondatatypes.RicactionDefinition{},
				RicSubsequentAction: subsequentAction,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	}

	ricSubscriptionDetails := &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes30{
		Id:          int32(v1beta1.ProtocolIeIDRicsubscriptionDetails),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RicsubscriptionDetails{
			RicEventTriggerDefinition: &e2ap_commondatatypes.RiceventTriggerDefinition{
				Value: sub.Details.EventTriggerDefinition,
			},
			RicActionToBeSetupList: &e2appducontents.RicactionsToBeSetupList{
				Value: actions,
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	request := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicSubscription: &e2appdudescriptions.RicSubscription{
						InitiatingMessage: &e2appducontents.RicsubscriptionRequest{
							ProtocolIes: &e2appducontents.RicsubscriptionRequestIes{
								E2ApProtocolIes29: ricRequestID,
								E2ApProtocolIes5:  ranFunctionID,
								E2ApProtocolIes30: ricSubscriptionDetails,
							},
						},
					},
				},
			},
		},
	}

	// Validate the subscribe request
	if err := request.Validate(); err != nil {
		log.Warnf("Failed to validate E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, err := channel.SendRecv(ctx, request, channelfilter.RicSubscription(ricRequestID.Value), codec.PER)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	}

	switch response.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
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
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
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

	channel, err := r.channels.Get(ctx, channel.ID(sub.E2NodeID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, err
	}

	record := r.catalog.Get(sub.ID)

	// Generate a request ID
	ricRequestID := e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes29{
		Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(record.RequestID),
			RicInstanceId:  config.InstanceID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	// Create a RAN function ID from the requested function ID
	ranFunctionID := e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes5{
		Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: 0, // TODO: Map service model to RAN function ID
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	// Create a subscription delete request
	request := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicSubscriptionDelete: &e2appdudescriptions.RicSubscriptionDelete{
						InitiatingMessage: &e2appducontents.RicsubscriptionDeleteRequest{
							ProtocolIes: &e2appducontents.RicsubscriptionDeleteRequestIes{
								E2ApProtocolIes29: &ricRequestID,
								E2ApProtocolIes5:  &ranFunctionID,
							},
						},
					},
				},
			},
		},
	}

	// Validate the subscription delete request
	if err := request.Validate(); err != nil {
		log.Warnf("Failed to validate E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, err := channel.SendRecv(ctx, request, channelfilter.RicSubscriptionDelete(ricRequestID.Value), codec.PER)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for SubscriptionTask %+v: %s", request, task, err)
		return controller.Result{}, err
	}

	switch response.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		task.Lifecycle.Status = subtaskapi.Status_COMPLETE
		updateRequest := &subtaskapi.UpdateSubscriptionTaskRequest{
			Task: task,
		}
		_, err := r.tasks.UpdateSubscriptionTask(ctx, updateRequest)
		if err != nil {
			log.Warnf("Failed to update SubscriptionTask %+v: %s", task, err)
			return controller.Result{}, err
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		log.Warnf("Failed to initialize SubscriptionTask %+v: %s", task, err)
		return controller.Result{}, fmt.Errorf("failed to delete subscription %+v", sub)
	}
	return controller.Result{}, nil
}
