// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscription

import (
	"context"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	"time"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/oid"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/config"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "subscription")

// NewController returns a new network controller
func NewController(streams stream.Manager, subs substore.Store, topo rnib.Store, conns e2server.E2APConnManager,
	models modelregistry.ModelRegistry, oidRegistry oid.Registry) *controller.Controller {
	c := controller.NewController("Subscription")
	c.Watch(&Watcher{
		subs: subs,
	})
	c.Watch(&ConnWatcher{
		subs:  subs,
		conns: conns,
	})
	c.Watch(&TopoWatcher{
		subs: subs,
		topo: topo,
	})
	c.Watch(&StreamWatcher{
		streams: streams,
	})
	c.Reconcile(&Reconciler{
		streams:                   streams,
		subs:                      subs,
		topo:                      topo,
		conns:                     conns,
		models:                    models,
		oidRegistry:               oidRegistry,
		newRicSubscriptionRequest: pdubuilder.NewRicSubscriptionRequest,
	})
	return c
}

type RicSubscriptionRequestBuilder func(ricReq types.RicRequest,
	ranFuncID types.RanFunctionID, ricEventDef types.RicEventDefintion,
	ricActionsToBeSetup map[types.RicActionID]types.RicActionDef) (
	*e2appducontents.RicsubscriptionRequest, error)

// Reconciler is a device change reconciler
type Reconciler struct {
	streams                   stream.Manager
	subs                      substore.Store
	topo                      rnib.Store
	conns                     e2server.E2APConnManager
	models                    modelregistry.ModelRegistry
	oidRegistry               oid.Registry
	newRicSubscriptionRequest RicSubscriptionRequestBuilder
}

// Reconcile reconciles the state of a device change
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	subID := id.Value.(e2api.SubscriptionID)
	log.Infof("Reconciling Subscription '%s'", subID)
	sub, err := r.subs.Get(ctx, subID)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Debugf("Subscription '%s' not found", subID)
			return controller.Result{}, nil
		}
		log.Errorf("Failed to reconcile Subscription '%s'", subID, err)
		return controller.Result{}, err
	}

	log.Debug(sub)

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

	// If there are no channels associated with this subscription, close the subscription
	if len(sub.Status.Channels) == 0 {
		log.Infof("Closing Subscription '%s'", sub.ID)
		sub.Status.Phase = e2api.SubscriptionPhase_SUBSCRIPTION_CLOSED
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
		sub.Status.Error = nil
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error closing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	log.Debugf("Fetching mastership state for E2Node '%s'", sub.E2NodeID)

	// Get the E2 node entity from topo
	e2NodeEntity, err := r.topo.Get(ctx, topoapi.ID(sub.E2NodeID))
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warnf("Mastership state not found for E2Node '%s'", sub.E2NodeID, err)
			return controller.Result{}, nil
		}
		log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
		return controller.Result{}, err
	}

	// Decode the mastership state from the E2 node entity
	mastership := topoapi.MastershipState{}
	_ = e2NodeEntity.GetAspect(&mastership)
	e2NodeMastershipTerm := e2api.TermID(mastership.Term)

	// If no master is found, skip reconciliation
	if e2NodeMastershipTerm == 0 {
		log.Warnf("Mastership state not found for E2Node '%s'", sub.E2NodeID)
		return controller.Result{}, nil
	}

	e2NodeMasterRelation, err := r.topo.Get(ctx, topoapi.ID(mastership.NodeId))
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warnf("Master relation not found for E2Node '%s'", sub.E2NodeID)
			return controller.Result{}, nil
		}
		log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
		return controller.Result{}, err
	}

	e2NodeMasterID := e2api.MasterID(e2NodeMasterRelation.ID)
	e2NodeMasterInstanceID := e2api.E2TInstanceID(e2NodeMasterRelation.GetRelation().SrcEntityID)
	localInstanceID := e2api.E2TInstanceID(utils.GetE2TID())

	// If the subscription term has not been set, set the subscription master and term
	if sub.Status.Term == 0 {
		log.Infof("Updating Subscription %s mastership to term %d", sub.ID, e2NodeMastershipTerm)
		sub.Status.Term = e2NodeMastershipTerm
		sub.Status.Master = e2NodeMasterID
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error updating mastership for Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	// If the subscription term is greater than the E2 node mastership term, return
	// to wait for the E2 node state change
	if sub.Status.Term > e2NodeMastershipTerm {
		return controller.Result{}, nil
	}

	// If mastership has changed, delete the subscription and update the subscription status
	if sub.Status.Term < e2NodeMastershipTerm {
		// Check whether the master relation still exists
		subMasterRelation, err := r.topo.Get(ctx, topoapi.ID(sub.Status.Master))
		if err != nil {
			// Subscriptions are associated with specific connections. If the master relation
			// is not found the subscription cannot be deleted. Simply update the subscription term.
			if errors.IsNotFound(err) {
				log.Infof("Updating Subscription %s mastership to term %d", sub.ID, e2NodeMastershipTerm)
				sub.Status.Term = e2NodeMastershipTerm
				sub.Status.Master = e2NodeMasterID
				sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
				sub.Status.Error = nil
				log.Debug(sub)
				if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
					log.Errorf("Failed to reconcile Subscription '%s'", sub.ID, err)
					return controller.Result{}, err
				}
				return controller.Result{}, nil
			}
			log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
			return controller.Result{}, err
		}

		// If this node is still the subscription master, delete the subscription before
		// updating the subscription status to the new master/term
		subMasterInstanceID := e2api.E2TInstanceID(subMasterRelation.GetRelation().SrcEntityID)
		if subMasterInstanceID != localInstanceID {
			return controller.Result{}, nil
		}

		// Get the connection from the registry
		subMasterConn, err := r.conns.Get(ctx, e2server.ConnID(subMasterRelation.ID))
		if err != nil {
			if errors.IsNotFound(err) {
				log.Warnf("Connection not found for E2Node '%s'", sub.E2NodeID)
				return controller.Result{}, nil
			}
			log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
			return controller.Result{}, err
		}

		stream := r.streams.Open(sub.ID)

		ricRequest := types.RicRequest{
			RequestorID: types.RicRequestorID(stream.StreamID()),
			InstanceID:  config.InstanceID,
		}

		serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry, string(sub.ServiceModel.Name), string(sub.ServiceModel.Version))
		if err != nil {
			log.Warnf("Failed to reconcile Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}

		ranFunctionID, ok := subMasterConn.GetRANFunctionID(ctx, serviceModelOID)
		if !ok {
			log.Warn("RAN function not found for SM %s", serviceModelOID)
		}

		request, err := pdubuilder.NewRicSubscriptionDeleteRequest(ricRequest, ranFunctionID)
		if err != nil {
			log.Warnf("Failed to reconcile Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}

		// Send the subscription request and await a response
		log.Debugf("Sending RicsubscriptionDeleteRequest %+v", request)
		response, failure, err := subMasterConn.RICSubscriptionDelete(ctx, request)
		if err != nil {
			log.Warnf("Failed to send E2ApPdu %+v for Subscription '%s'", request, sub.ID, err)
			return controller.Result{}, err
		} else if response != nil {
			defer stream.Close()
			log.Debugf("Received RicsubscriptionDeleteResponse %+v", response)
			log.Infof("Updating Subscription '%s' mastership to term %d", sub.ID, e2NodeMastershipTerm)
			sub.Status.Term = e2NodeMastershipTerm
			sub.Status.Master = e2NodeMasterID
			sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
			sub.Status.Error = nil
			log.Debug(sub)
			if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Errorf("Error updating mastership for Subscription '%s'", sub.ID, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		} else if failure != nil {
			var cause *e2apies.Cause
			for _, v := range failure.GetProtocolIes() {
				if v.Id == int32(v2.ProtocolIeIDCause) {
					cause = v.GetValue().GetC()
					break
				}
			}
			switch cause.GetCause().(type) {
			case *e2apies.Cause_RicRequest:
				e2apErr := getSubscriptionDeleteError(failure)
				switch c := e2apErr.GetCause().GetCause().(type) {
				case *e2api.Error_Cause_Ric_:
					switch c.Ric.GetType() {
					case e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN:
						defer stream.Close()
						log.Infof("Updating Subscription %s mastership to term %d", sub.ID, e2NodeMastershipTerm)
						sub.Status.Term = e2NodeMastershipTerm
						sub.Status.Master = e2NodeMasterID
						sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_PENDING
						sub.Status.Error = nil
						log.Debug(sub)
						if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
							log.Errorf("Error updating mastership for Subscription '%s'", sub.ID, err)
							return controller.Result{}, err
						}
						return controller.Result{}, nil
					default:
						return controller.Result{}, err
					}
				}
			default:
				return controller.Result{}, nil
			}
			log.Warnf("RicsubscriptionDeleteRequest %+v failed: %+v", request, failure)
			return controller.Result{}, errors.NewUnknown("Failed to delete Subscription '%s'", sub.ID)
		}
		return controller.Result{}, nil
	}

	// If we've gotten this far, the E2 node's mastership term matches the subscription
	// mastership term. Reconcile the subscription.

	// If the subscription has already been created, skip reconciliation
	if sub.Status.State != e2api.SubscriptionState_SUBSCRIPTION_PENDING {
		return controller.Result{}, nil
	}

	// If this node is the master, create the subscription on the E2 node if hasn't already been created.
	if localInstanceID != e2NodeMasterInstanceID {
		log.Debugf("Not the master for E2Node '%s'", sub.E2NodeID)
		return controller.Result{}, nil
	}

	// Get the master connection from the connection registry
	e2NodeMasterConn, err := r.conns.Get(ctx, e2server.ConnID(e2NodeMasterRelation.ID))
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warnf("Connection not found for E2Node '%s'", sub.E2NodeID)
			log.Infof("Failing Subscription '%s'", sub.ID)
			sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_FAILED
			sub.Status.Error = &e2api.Error{
				Cause: &e2api.Error_Cause{
					Cause: &e2api.Error_Cause_Transport_{
						Transport: &e2api.Error_Cause_Transport{
							Type: e2api.Error_Cause_Transport_TRANSPORT_RESOURCE_UNAVAILABLE,
						},
					},
				},
			}
			log.Debug(sub)
			if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Errorf("Error failing Subscription '%s'", sub.ID, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}
		log.Warnf("Fetching mastership state for E2Node '%s' failed", sub.E2NodeID, err)
		return controller.Result{}, err
	}

	serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry,
		string(sub.ServiceModel.Name),
		string(sub.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		log.Infof("Failing Subscription '%s'", sub.ID)
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
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error failing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	serviceModelPlugin, err := r.models.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warn(err)
		log.Infof("Failing Subscription '%s'", sub.ID)
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
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error failing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	smData := serviceModelPlugin.ServiceModelData()
	log.Debugf("Service model found %s %s %s", smData.Name, smData.Version, smData.OID)

	stream := r.streams.Open(sub.ID)

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(stream.StreamID()),
		InstanceID:  config.InstanceID,
	}

	ranFunctionID, ok := e2NodeMasterConn.GetRANFunctionID(ctx, serviceModelOID)
	if !ok {
		log.Warn("RAN function not found for SM %s", serviceModelOID)
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

	request, err := r.newRicSubscriptionRequest(ricRequest, ranFunctionID, ricEventDef, ricActionsToBeSetup)
	if err != nil {
		log.Warnf("Failed to create E2ApPdu %+v for Subscription '%s'", request, sub.ID, err)
		return controller.Result{}, err
	}

	// Validate the subscribe request
	// TODO enable this when validation function is available
	/*if err := request.Validate(); err != nil {
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
		log.Debugf("Updating failed Subscription %+v", sub)
		err := r.subs.Update(ctx, sub)
		if err != nil {
			log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}*/

	// Send the subscription request and await a response
	log.Debugf("Sending RicsubscriptionRequest %+v", request)
	response, failure, err := e2NodeMasterConn.RICSubscription(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for Subscription '%s'", request, sub.ID, err)
		return controller.Result{}, err
	} else if response != nil {
		log.Debugf("Received RicsubscriptionResponse %+v", response)
		log.Infof("Completing Subscription '%s'", sub.ID)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error completing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	} else if failure != nil {
		log.Warnf("RicsubscriptionRequest %+v failed: %+v", request, failure)
		log.Infof("Failing Subscription '%s'", sub.ID)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_FAILED
		sub.Status.Error = getSubscriptionError(failure)
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error completing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileClosedSubscription(sub *e2api.Subscription) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// If the close has completed, delete the subscription
	if sub.Status.State == e2api.SubscriptionState_SUBSCRIPTION_COMPLETE {
		log.Infof("Deleting closed Subscription '%s'", sub.ID)
		err := r.subs.Delete(ctx, sub)
		if err != nil && !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	if sub.Status.Term == 0 {
		log.Warnf("No master for Subscription '%s'", sub.ID)
		log.Infof("Completing Subscription '%s'", sub.ID)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error completing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	log.Debugf("Fetching mastership state for E2Node '%s'", sub.E2NodeID)
	e2NodeRelation, err := r.topo.Get(ctx, topoapi.ID(sub.Status.Master))
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warnf("Mastership state not found for E2Node '%s'", sub.E2NodeID)
			log.Infof("Completing Subscription '%s'", sub.ID)
			sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
			log.Debug(sub)
			if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Errorf("Error completing Subscription '%s'", sub.ID, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}
		log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
		return controller.Result{}, err
	}

	if e2NodeRelation.GetRelation().SrcEntityID != utils.GetE2TID() {
		return controller.Result{}, nil
	}

	conn, err := r.conns.Get(ctx, e2server.ConnID(e2NodeRelation.ID))
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warnf("Connection %s not found for E2Node '%s'", e2NodeRelation.ID, sub.E2NodeID)
			log.Infof("Completing Subscription '%s'", sub.ID)
			sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
			log.Debug(sub)
			if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
				log.Errorf("Error completing Subscription '%s'", sub.ID, err)
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}
		log.Errorf("Error fetching mastership state for E2Node '%s'", sub.E2NodeID, err)
		return controller.Result{}, err
	}

	stream := r.streams.Open(sub.ID)

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(stream.StreamID()),
		InstanceID:  config.InstanceID,
	}

	serviceModelOID, err := oid.ModelIDToOid(r.oidRegistry, string(sub.ServiceModel.Name), string(sub.ServiceModel.Version))
	if err != nil {
		log.Errorf("Failed to reconcile Subscription '%s'", sub.ID, err)
		return controller.Result{}, err
	}

	ranFunctionID, ok := conn.GetRANFunctionID(ctx, serviceModelOID)
	if !ok {
		log.Warn("RAN function not found for SM %s", serviceModelOID)
	}

	request, err := pdubuilder.NewRicSubscriptionDeleteRequest(ricRequest, ranFunctionID)
	if err != nil {
		log.Warnf("Failed to reconcile Subscription '%s'", sub.ID, err)
		return controller.Result{}, err
	}

	// Send the subscription request and await a response
	log.Debugf("Sending RicsubscriptionDeleteRequest %+v", request)
	response, failure, err := conn.RICSubscriptionDelete(ctx, request)
	if err != nil {
		log.Warnf("Failed to send E2ApPdu %+v for Subscription '%s'", request, sub.ID, err)
		return controller.Result{}, err
	} else if response != nil {
		defer stream.Close()
		log.Debugf("Received RicsubscriptionDeleteResponse %+v", response)
		log.Infof("Completing Subscription '%s'", sub.ID)
		sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
		log.Debug(sub)
		if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
			log.Errorf("Error completing Subscription '%s'", sub.ID, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	} else if failure != nil {
		var cause *e2apies.Cause
		for _, v := range failure.GetProtocolIes() {
			if v.Id == int32(v2.ProtocolIeIDCause) {
				cause = v.GetValue().GetC()
				break
			}
		}
		switch cause.GetCause().(type) {
		case *e2apies.Cause_RicRequest:
			e2apErr := getSubscriptionDeleteError(failure)
			switch c := e2apErr.GetCause().GetCause().(type) {
			case *e2api.Error_Cause_Ric_:
				switch c.Ric.GetType() {
				case e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN:
					defer stream.Close()
					log.Infof("Completing Subscription '%s'", sub.ID)
					sub.Status.State = e2api.SubscriptionState_SUBSCRIPTION_COMPLETE
					log.Debug(sub)
					if err := r.subs.Update(ctx, sub); err != nil && !errors.IsNotFound(err) && !errors.IsConflict(err) {
						log.Errorf("Error completing Subscription '%s'", sub.ID, err)
						return controller.Result{}, err
					}
					return controller.Result{}, nil
				default:
					return controller.Result{}, err
				}
			}
		default:
			return controller.Result{}, nil
		}
		log.Warnf("RicsubscriptionDeleteRequest %+v failed: %+v", request, failure)
		return controller.Result{}, errors.NewUnknown("Failed to delete Subscription '%s'", sub.ID)
	}
	return controller.Result{}, nil
}

func getSubscriptionDeleteError(failure *e2appducontents.RicsubscriptionDeleteFailure) *e2api.Error {
	if failure == nil {
		return nil
	}

	var cause *e2apies.Cause
	for _, v := range failure.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDCause) {
			cause = v.GetValue().GetC()
			break
		}
	}
	switch c := cause.GetCause().(type) {
	case *e2apies.Cause_RicRequest:
		var errType e2api.Error_Cause_Ric_Type
		switch c.RicRequest {
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_RAN_FUNCTION_ID_INVALID:
			errType = e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_ACTION_NOT_SUPPORTED:
			errType = e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_EXCESSIVE_ACTIONS:
			errType = e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_ACTION:
			errType = e2api.Error_Cause_Ric_DUPLICATE_ACTION
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_EVENT_TRIGGER:
			errType = e2api.Error_Cause_Ric_DUPLICATE_EVENT
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_FUNCTION_RESOURCE_LIMIT:
			errType = e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN:
			errType = e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
			errType = e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_MESSAGE_INVALID:
			errType = e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_RIC_CALL_PROCESS_ID_INVALID:
			errType = e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID
		//ToDo - fill in missing part in onos-api
		//case e2apies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_TIMER_EXPIRED:
		//	errType = e2api.Error_Cause_Ric_CONTROL_TIMER_EXPIRED
		//case e2apies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_FAILED_TO_EXECUTE:
		//	errType = e2api.Error_Cause_Ric_CONTROL_FAILED_TO_EXECUTE
		//case e2apies.CauseRicrequest_CAUSE_RICREQUEST_SYSTEM_NOT_READY:
		//	errType = e2api.Error_Cause_Ric_CONTROL_SYSTEM_NOT_READY
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED:
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
		//ToDo - change naming in onos-api to Error_Cause_RicService_RAN_FUNCTION_NOT_REQUIRED
		case e2apies.CauseRicservice_CAUSE_RICSERVICE_RAN_FUNCTION_NOT_SUPPORTED:
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

func getSubscriptionError(failure *e2appducontents.RicsubscriptionFailure) *e2api.Error {
	if failure == nil {
		return nil
	}

	var cause *e2apies.Cause
	for _, v := range failure.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDCause) {
			cause = v.GetValue().GetC()
			break
		}
	}
	switch c := cause.GetCause().(type) {
	case *e2apies.Cause_RicRequest:
		var errType e2api.Error_Cause_Ric_Type
		switch c.RicRequest {
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_RAN_FUNCTION_ID_INVALID:
			errType = e2api.Error_Cause_Ric_RAN_FUNCTION_ID_INVALID
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_ACTION_NOT_SUPPORTED:
			errType = e2api.Error_Cause_Ric_ACTION_NOT_SUPPORTED
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_EXCESSIVE_ACTIONS:
			errType = e2api.Error_Cause_Ric_EXCESSIVE_ACTIONS
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_ACTION:
			errType = e2api.Error_Cause_Ric_DUPLICATE_ACTION
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_DUPLICATE_EVENT_TRIGGER:
			errType = e2api.Error_Cause_Ric_DUPLICATE_EVENT
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_FUNCTION_RESOURCE_LIMIT:
			errType = e2api.Error_Cause_Ric_FUNCTION_RESOURCE_LIMIT
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_REQUEST_ID_UNKNOWN:
			errType = e2api.Error_Cause_Ric_REQUEST_ID_UNKNOWN
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE:
			errType = e2api.Error_Cause_Ric_INCONSISTENT_ACTION_SUBSEQUENT_ACTION_SEQUENCE
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_MESSAGE_INVALID:
			errType = e2api.Error_Cause_Ric_CONTROL_MESSAGE_INVALID
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_RIC_CALL_PROCESS_ID_INVALID:
			errType = e2api.Error_Cause_Ric_CALL_PROCESS_ID_INVALID
		//ToDo - fill in missing part in onos-api
		//case e2apies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_TIMER_EXPIRED:
		//	errType = e2api.Error_Cause_Ric_CONTROL_TIMER_EXPIRED
		//case e2apies.CauseRicrequest_CAUSE_RICREQUEST_CONTROL_FAILED_TO_EXECUTE:
		//	errType = e2api.Error_Cause_Ric_CONTROL_FAILED_TO_EXECUTE
		//case e2apies.CauseRicrequest_CAUSE_RICREQUEST_SYSTEM_NOT_READY:
		//	errType = e2api.Error_Cause_Ric_CONTROL_SYSTEM_NOT_READY
		case e2apies.CauseRicrequest_CAUSE_RICREQUEST_UNSPECIFIED:
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
		//ToDo - change naming in onos-api to Error_Cause_RicService_RAN_FUNCTION_NOT_REQUIRED
		case e2apies.CauseRicservice_CAUSE_RICSERVICE_RAN_FUNCTION_NOT_SUPPORTED:
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
