// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1alpha1

import (
	"context"
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"time"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "subscription")

// NewController returns a new network controller
func NewController(store rnib.Store, channels e2server.ChannelManager) *controller.Controller {
	c := controller.NewController("Subscription")
	c.Watch(&Watcher{
		topo: store,
	})
	c.Watch(&ChannelWatcher{
		channels: channels,
	})
	c.Reconcile(&Reconciler{
		store:    store,
		channels: channels,
	})
	return c
}

type RicSubscriptionRequestBuilder func(ricReq types.RicRequest,
	ranFuncID types.RanFunctionID, ricEventDef types.RicEventDefintion,
	ricActionsToBeSetup map[types.RicActionID]types.RicActionDef) (
	*e2appducontents.RicsubscriptionRequest, error)

// Reconciler is a device change reconciler
type Reconciler struct {
	store    rnib.Store
	channels e2server.ChannelManager
}

// Reconcile reconciles the state of a device change
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID := id.Value.(e2server.ChannelID)
	log.Infof("Reconciling Channel %s", channelID)
	channel, err := r.channels.Get(ctx, channelID)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.reconcileClosedChannel(channelID)
		}
		log.Warnf("Failed to reconcile Channel %s: %s", channelID, err)
		return controller.Result{}, err
	}
	return r.reconcileOpenChannel(channel)
}

func (r *Reconciler) reconcileOpenChannel(channel *e2server.E2Channel) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if ok, err := r.addE2T(ctx); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.addE2Node(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.addE2Relation(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) addE2T(ctx context.Context) (bool, error) {
	object, err := r.store.Get(ctx, topoapi.ID(env.GetPodID()))
	if err != nil {
		if !errors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	}

	object = &topoapi.Object{
		ID:   topoapi.ID(env.GetPodID()),
		Type: topoapi.Object_ENTITY,
		Obj: &topoapi.Object_Entity{
			Entity: &topoapi.Entity{
				KindID: topoapi.E2T,
			},
		},
		Aspects: make(map[string]*gogotypes.Any),
		Labels:  map[string]string{},
	}
	err = r.store.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) addE2Node(ctx context.Context, channel *e2server.E2Channel) (bool, error) {
	object, err := r.store.Get(ctx, channel.E2NodeID)
	if err != nil {
		if !errors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	}

	object = &topoapi.Object{
		ID:   channel.E2NodeID,
		Type: topoapi.Object_ENTITY,
		Obj: &topoapi.Object_Entity{
			Entity: &topoapi.Entity{
				KindID: topoapi.E2NODE,
			},
		},
		Aspects: make(map[string]*gogotypes.Any),
		Labels:  map[string]string{},
	}

	aspects := &topoapi.E2Node{
		ServiceModels: channel.ServiceModels,
	}

	err = object.SetAspect(aspects)
	if err != nil {
		return false, nil
	}

	err = r.store.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) addE2Relation(ctx context.Context, channel *e2server.E2Channel) (bool, error) {
	object, err := r.store.Get(ctx, topoapi.ID(channel.ID))
	if err != nil {
		if !errors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	}

	object = &topoapi.Object{
		ID:   topoapi.ID(channel.ID),
		Type: topoapi.Object_RELATION,
		Obj: &topoapi.Object_Relation{
			Relation: &topoapi.Relation{
				KindID:      topoapi.CONTROLS,
				SrcEntityID: topoapi.ID(env.GetPodID()),
				TgtEntityID: channel.E2NodeID,
			},
		},
	}

	err = r.store.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) reconcileClosedChannel(channelID e2server.ChannelID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if ok, err := r.removeE2Relation(ctx, channelID); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	return controller.Result{}, nil
}

func (r *Reconciler) removeE2Relation(ctx context.Context, channelID e2server.ChannelID) (bool, error) {
	err := r.store.Delete(ctx, topoapi.ID(channelID))
	if err != nil {
		if !errors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
