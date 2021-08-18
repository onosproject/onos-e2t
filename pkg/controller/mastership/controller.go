// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mastership

import (
	"context"
	"math/rand"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "mastership")

// NewController returns a new mastership controller
func NewController(store rnib.Store, channels e2server.ChannelManager) *controller.Controller {
	c := controller.NewController("mastership")
	c.Watch(&Watcher{
		topo: store,
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

// Reconcile reconciles E2 nodes, E2 cells, and mastership election
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID := id.Value.(e2server.ChannelID)
	log.Infof("Reconciling mastership election for channel %s", channelID)
	channel, err := r.channels.Get(ctx, channelID)
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile mastership election for channel %s: %s", channelID, err)
		return controller.Result{}, err
	}
	return r.reconcileMastershipElection(channel)
}

func (r *Reconciler) reconcileMastershipElection(channel *e2server.E2Channel) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.updateE2NodeMaster(ctx, channel); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) updateE2NodeMaster(ctx context.Context, channel *e2server.E2Channel) error {
	log.Debugf("Verifying mastership for E2Node '%s'", channel.E2NodeID)
	e2NodeEntity, err := r.store.Get(ctx, channel.E2NodeID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", channel.E2NodeID, err)
			return err
		}
		log.Warnf("E2Node entity '%s' not found", channel.E2NodeID)
		return nil
	}

	// List the objects in the topo store
	objects, err := r.store.List(ctx, nil)
	if err != nil {
		log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", channel.E2NodeID, err)
		return err
	}

	// Filter the topo objects for relations
	e2NodeRelations := make(map[string]topoapi.Object)
	for _, object := range objects {
		if relation, ok := object.Obj.(*topoapi.Object_Relation); ok &&
			relation.Relation.KindID == topoapi.CONTROLS &&
			relation.Relation.TgtEntityID == channel.E2NodeID {
			e2NodeRelations[string(object.ID)] = object
		}
	}

	mastership := &topoapi.MastershipState{}
	mastershipValue := e2NodeEntity.GetAspect(mastership)
	if _, ok := e2NodeRelations[mastership.NodeId]; !ok || mastershipValue == nil {
		log.Debugf("Updating MastershipState for E2Node '%s'", channel.E2NodeID)
		if len(e2NodeRelations) == 0 {
			log.Warnf("No controls relations found for E2Node entity '%s'", channel.E2NodeID)
			return nil
		}

		// Select a random master to assign to the E2 node
		relations := make([]topoapi.Object, 0, len(e2NodeRelations))
		for _, e2nodeRelation := range e2NodeRelations {
			relations = append(relations, e2nodeRelation)
		}
		relation := relations[rand.Intn(len(relations))]

		// Increment the mastership term and assign the selected master
		mastership.Term++
		mastership.NodeId = string(relation.ID)
		err = e2NodeEntity.SetAspect(mastership)
		if err != nil {
			log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", channel.E2NodeID, err)
			return err
		}

		// Update the E2 node entity
		err = r.store.Update(ctx, e2NodeEntity)
		if err != nil {
			if !errors.IsNotFound(err) {
				log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", channel.E2NodeID, err)
				return err
			}
			return nil
		}
		return nil
	}
	return nil
}
