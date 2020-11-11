// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"time"

	"github.com/gogo/protobuf/proto"
	regapi "github.com/onosproject/onos-e2sub/api/e2/registry/v1beta1"
	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	subtaskapi "github.com/onosproject/onos-e2sub/api/e2/task/v1beta1"
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

var log = logging.GetLogger("subscription", "controller")

const defaultTimeout = 30 * time.Second

// NewController returns a new network controller
func NewController(catalog *Catalog, subs subapi.E2SubscriptionServiceClient, tasks subtaskapi.E2SubscriptionTaskServiceClient, channels *channel.Manager) *controller.Controller {
	c := controller.NewController("SubscriptionTask")
	c.Watch(&Watcher{
		endpointID: regapi.ID(env.GetPodID()),
		tasks:      tasks,
	})
	c.Watch(&ChannelWatcher{
		endpointID: regapi.ID(env.GetPodID()),
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
	catalog   *Catalog
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

	// If the task is COMPLETE, ignore the request
	if task.Lifecycle.Status == subtaskapi.Status_COMPLETE {
		return controller.Result{}, nil
	}

	// Process the request based on the lifecycle phase
	switch task.Lifecycle.Phase {
	case subtaskapi.Phase_OPEN:
		return r.reconcileOpenSubscriptionTask(task)
	case subtaskapi.Phase_CLOSE:
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
		return controller.Result{}, err
	}
	sub := subResponse.Subscription

	channel, err := r.channels.Get(ctx, channel.ID(sub.E2NodeID))
	if err != nil {
		return controller.Result{}, err
	}

	r.requestID++
	requestID := r.requestID

	request := &e2appdudescriptions.E2ApPdu{}
	err = proto.Unmarshal(sub.Payload.Bytes, request)
	if err != nil {
		return controller.Result{}, err
	}

	// Generate a request ID
	ricRequestID := &e2apies.RicrequestId{
		RicRequestorId: int32(requestID),
		RicInstanceId:  config.InstanceID,
	}

	// Update the subscription request with a request ID
	request.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
		Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ricRequestID,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	// Validate the subscribe request
	if err := request.Validate(); err != nil {
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, err := channel.SendRecv(request, channelfilter.RicSubscription(ricRequestID), codec.XER)
	if err != nil {
		return controller.Result{}, err
	}

	switch response.E2ApPdu.(type) {
	case *e2appdudescriptions.E2ApPdu_SuccessfulOutcome:
		record := CatalogRecord{
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
			return controller.Result{}, err
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
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
		return controller.Result{}, err
	}
	sub := subResponse.Subscription

	channel, err := r.channels.Get(ctx, channel.ID(sub.E2NodeID))
	if err != nil {
		return controller.Result{}, err
	}

	record := r.catalog.Get(sub.ID)

	subscriptionRequest := &e2appdudescriptions.E2ApPdu{}
	err = proto.Unmarshal(sub.Payload.Bytes, subscriptionRequest)
	if err != nil {
		return controller.Result{}, err
	}

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
		Value:       subscriptionRequest.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes5.Value,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
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
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	response, err := channel.SendRecv(request, channelfilter.RicSubscriptionDelete(ricRequestID.Value), codec.XER)
	if err != nil {
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
			return controller.Result{}, err
		}
	case *e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome:
		return controller.Result{}, fmt.Errorf("failed to delete subscription %+v", sub)
	}
	return controller.Result{}, nil
}
