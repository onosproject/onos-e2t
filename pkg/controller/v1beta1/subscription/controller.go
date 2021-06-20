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

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/config"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "subscription")

// NewController returns a new network controller
func NewController(streams subscription.Broker, subs substore.Store, channels e2server.ChannelManager,
	models modelregistry.ModelRegistry, oidRegistry oid.Registry, ranFunctionRegistry ranfunctions.Registry,
	topoManager topo.Manager) *controller.Controller {
	c := controller.NewController("Subscription")
	c.Watch(&Watcher{
		subs: subs,
	})
	c.Watch(&ChannelWatcher{
		subs:     subs,
		channels: channels,
	})
	c.Reconcile(&Reconciler{
		streams:                   streams,
		subs:                      subs,
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

	subID := id.Value.(e2api.SubscriptionID)
	sub, err := r.subs.Get(ctx, subID)
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
		return controller.Result{}, err
	}

	log.Infof("Reconciling Subscription %+v", sub)

	// If the sub is COMPLETE or FAILED, ignore the request
	switch sub.Status.State {
	case e2api.SubscriptionState_SUBSCRIPTION_COMPLETE, e2api.SubscriptionState_SUBSCRIPTION_FAILED:
		return controller.Result{}, nil
	}

	// Process the request based on the lifecycle phase
	switch sub.Status.Phase {
	case e2api.SubscriptionPhase_SUBSCRIPTION_OPEN:
		return r.reconcileOpenSubscription(sub)
	case e2api.SubscriptionPhase_SUBSCRIPTION_CLOSED:
		return r.reconcileClosedSubscription(sub)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileOpenSubscription(sub *e2api.Subscription) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// If no northbound channels are linked to the subscription, close the subscription
	if len(sub.Status.Channels) == 0 {
		sub.Status.Phase = e2api.SubscriptionPhase_SUBSCRIPTION_CLOSED
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
		sub.Status.Error = nil
		if err := r.subs.Update(ctx, sub); err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// Get the southbound channel for the E2 node
	channelID, err := r.topoManager.GetE2Relation(ctx, topoapi.ID(sub.SubscriptionMeta.E2NodeID))
	if err != nil || channelID == "" {
		return controller.Result{}, err
	}
	channel, err := r.channels.Get(ctx, e2server.ChannelID(channelID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
		return controller.Result{}, err
	}

	serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry,
		string(sub.ServiceModel.Name),
		string(sub.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_FAILED
		sub.Status.Error = &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Ric_{
					Ric: &e2api.Error_Cause_Ric{
						Type: e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID,
					},
				},
			},
		}
		err = r.subs.Update(ctx, sub)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	serviceModelPlugin, err := r.models.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warn(err)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_FAILED
		sub.Status.Error = &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Ric_{
					Ric: &e2api.Error_Cause_Ric{
						Type: e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID,
					},
				},
			},
		}
		err = r.subs.Update(ctx, sub)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	smData := serviceModelPlugin.ServiceModelData()
	log.Infof("Service model found %s %s %s", smData.Name, smData.Version, smData.OID)

	stream, ok := r.streams.GetReader(sub.ID)
	if !ok {
		return controller.Result{}, errors.NewNotFound("stream not found for sub %s", sub.ID)
	}

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(stream.ID()),
		InstanceID:  config.InstanceID,
	}

	ranFunction, err := r.ranFunctionRegistry.Get(ranfunctions.NewID(serviceModelOID, string(channelID)))
	if err != nil {
		log.Warn(err)
	}

	ricEventDef := types.RicEventDefintion(sub.Spec.EventTrigger.Payload)

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	for _, action := range sub.Spec.Actions {
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
		log.Warnf("Failed to create E2ApPdu %+v for Subscription %+v: %s", request, sub, err)
		return controller.Result{}, err
	}

	// Validate the subscribe request
	if err := request.Validate(); err != nil {
		log.Warnf("Failed to validate E2ApPdu %+v for Subscription %+v: %s", request, sub, err)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_FAILED
		sub.Status.Error = &e2api.Error{
			Cause: &e2api.Error_Cause{
				Cause: &e2api.Error_Cause_Protocol_{
					Protocol: &e2api.Error_Cause_Protocol{
						Type: e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
					},
				},
			},
		}
		err := r.subs.Update(ctx, sub)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// Send the subscription request and await a response
	response, failure, err := channel.RICSubscription(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for Subscription %+v: %s", request, sub, err)
		return controller.Result{}, err
	} else if response != nil {
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
		err := r.subs.Update(ctx, sub)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	} else if failure != nil {
		log.Warnf("Failed to initialize Subscription %+v: %s", sub, err)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_FAILED
		sub.Status.Error = getSubscriptionError(failure)
		err := r.subs.Update(ctx, sub)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedSubscription(sub *e2api.Subscription) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// If northbound channels are linked to the subscription, change the phase to OPEN
	if len(sub.Status.Channels) > 0 {
		sub.Status.Phase = e2api.SubscriptionPhase_SUBSCRIPTION_OPEN
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
		sub.Status.Error = nil
		if err := r.subs.Update(ctx, sub); err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// Get the southbound channel ID for the E2 node
	channelID, err := r.topoManager.GetE2Relation(ctx, topoapi.ID(sub.SubscriptionMeta.E2NodeID))
	if err != nil || channelID == "" {
		return controller.Result{}, err
	}

	// Get the southbound indications channel for the E2 node
	channel, err := r.channels.Get(ctx, e2server.ChannelID(channelID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
		return controller.Result{}, err
	}

	// Get the subscription stream reader
	stream, ok := r.streams.GetReader(sub.ID)
	if !ok {
		return controller.Result{}, errors.NewNotFound("stream not found for sub %s", sub.ID)
	}

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(stream.ID()),
		InstanceID:  config.InstanceID,
	}

	serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry, string(sub.ServiceModel.Name), string(sub.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		return controller.Result{}, err
	}

	ranFunction, err := r.ranFunctionRegistry.Get(ranfunctions.NewID(serviceModelOID, string(channelID)))
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
		log.Warnf("Failed to send E2ApPdu %+v for Subscription %+v: %s", request, sub, err)
		return controller.Result{}, err
	} else if response != nil {
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
		err := r.subs.Update(ctx, sub)
		if err != nil {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	} else if failure != nil {
		switch failure.ProtocolIes.E2ApProtocolIes1.Value.Cause.(type) {
		case *e2apies.Cause_RicRequest:
			return controller.Result{}, nil
		}
		log.Warnf("SubscriptionDeleteRequest %+v failed: %+v", request, failure)
		return controller.Result{}, fmt.Errorf("failed to delete sub %+v", sub)
	}
	return controller.Result{}, nil
}

func getSubscriptionError(failure *e2appducontents.RicsubscriptionFailure) *e2api.Error {
	if failure == nil {
		return nil
	}

	for _, item := range failure.ProtocolIes.E2ApProtocolIes18.Value.Value {
		switch c := item.Value.Cause.Cause.(type) {
		case *e2apies.Cause_RicRequest:
			var errType e2api.Error_Cause_Ric_Type
			switch c.RicRequest {
			case e2apies.CauseRic_CAUSE_RIC_RAN_FUNCTION_ID_INVALID:
				errType = e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID
			case e2apies.CauseRic_CAUSE_RIC_ACTION_NOT_SUPPORTED:
				errType = e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED
			case e2apies.CauseRic_CAUSE_RIC_EXCESSIVE_ACTIONS:
				errType = e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS
			case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_ACTION:
				errType = e2api.Error_Cause_Ric_DUPLICATE_ACTION
			case e2apies.CauseRic_CAUSE_RIC_DUPLICATE_EVENT:
				errType = e2api.Error_Cause_Ric_DUPLICATE_EVENT
			case e2apies.CauseRic_CAUSE_RIC_FUNCTION_RESOURCE_LIMIT:
				errType = e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT
			case e2apies.CauseRic_CAUSE_RIC_REQUEST_ID_UNKNOWN:
				errType = e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN
			case e2apies.CauseRic_CAUSE_RIC_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
				errType = e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
			case e2apies.CauseRic_CAUSE_RIC_CONTROL_MESSAGE_INVALID:
				errType = e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID
			case e2apies.CauseRic_CAUSE_RIC_CALL_PROCESS_ID_INVALID:
				errType = e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID
			case e2apies.CauseRic_CAUSE_RIC_UNSPECIFIED:
				errType = e2api.Error_Cause_Ric_UNSPECIFIED
			}
			return &e2api.Error{
				Cause: &e2api.Error_Cause{
					Cause: &e2api.Error_Cause_Ric_{
						Ric: &e2api.Error_Cause_Ric{
							Type: errType,
						},
					},
				},
			}
		case *e2apies.Cause_RicService:
			var errType e2api.Error_Cause_RicService_Type
			switch c.RicService {
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_FUNCTION_NOT_REQUIRED:
				errType = e2api.Error_Cause_RicService_FUNCTION_NOT_REQUIRED
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_EXCESSIVE_FUNCTIONS:
				errType = e2api.Error_Cause_RicService_EXCESSIVE_FUNCTIONS
			case e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT:
				errType = e2api.Error_Cause_RicService_RIC_RESOURCE_LIMIT
			}
			return &e2api.Error{
				Cause: &e2api.Error_Cause{
					Cause: &e2api.Error_Cause_RicService_{
						RicService: &e2api.Error_Cause_RicService{
							Type: errType,
						},
					},
				},
			}
		case *e2apies.Cause_Protocol:
			var errType e2api.Error_Cause_Protocol_Type
			switch c.Protocol {
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR:
				errType = e2api.Error_Cause_Protocol_TRANSFER_SYNTAX_ERROR
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_REJECT:
				errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_REJECT
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY:
				errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_IGNORE_AND_NOTIFY
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE:
				errType = e2api.Error_Cause_Protocol_MESSAGE_NOT_COMPATIBLE_WITH_RECEIVER_STATE
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR:
				errType = e2api.Error_Cause_Protocol_SEMANTIC_ERROR
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE:
				errType = e2api.Error_Cause_Protocol_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE
			case e2apies.CauseProtocol_CAUSE_PROTOCOL_UNSPECIFIED:
				errType = e2api.Error_Cause_Protocol_UNSPECIFIED
			}
			return &e2api.Error{
				Cause: &e2api.Error_Cause{
					Cause: &e2api.Error_Cause_Protocol_{
						Protocol: &e2api.Error_Cause_Protocol{
							Type: errType,
						},
					},
				},
			}
		case *e2apies.Cause_Transport:
			var errType e2api.Error_Cause_Transport_Type
			switch c.Transport {
			case e2apies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED:
				errType = e2api.Error_Cause_Transport_UNSPECIFIED
			case e2apies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE:
				errType = e2api.Error_Cause_Transport_TRANSPORT_RESOURCE_UNAVAILABLE
			}
			return &e2api.Error{
				Cause: &e2api.Error_Cause{
					Cause: &e2api.Error_Cause_Transport_{
						Transport: &e2api.Error_Cause_Transport{
							Type: errType,
						},
					},
				},
			}
		case *e2apies.Cause_Misc:
			var errType e2api.Error_Cause_Misc_Type
			switch c.Misc {
			case e2apies.CauseMisc_CAUSE_MISC_CONTROL_PROCESSING_OVERLOAD:
				errType = e2api.Error_Cause_Misc_CONTROL_PROCESSING_OVERLOAD
			case e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE:
				errType = e2api.Error_Cause_Misc_HARDWARE_FAILURE
			case e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION:
				errType = e2api.Error_Cause_Misc_OM_INTERVENTION
			case e2apies.CauseMisc_CAUSE_MISC_UNSPECIFIED:
				errType = e2api.Error_Cause_Misc_UNSPECIFIED
			}
			return &e2api.Error{
				Cause: &e2api.Error_Cause{
					Cause: &e2api.Error_Cause_Misc_{
						Misc: &e2api.Error_Cause_Misc{
							Type: errType,
						},
					},
				},
			}
		}
		return nil
	}
	return nil
}
