// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mastership

import (
	"context"
	"math/rand"
	"time"

	"github.com/onosproject/onos-e2t/pkg/controller/utils"

	gogotypes "github.com/gogo/protobuf/types"
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

	if err := r.createE2Node(ctx, channel); err != nil {
		return controller.Result{}, err
	}

	if err := r.createE2Cells(ctx, channel); err != nil {
		return controller.Result{}, err
	}

	if err := r.updateE2NodeMaster(ctx, channel); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) createE2Node(ctx context.Context, channel *e2server.E2Channel) error {
	_, err := r.store.Get(ctx, channel.E2NodeID)
	if err == nil {
		return nil
	} else if !errors.IsNotFound(err) {
		log.Warnf("Creating E2Node entity '%s' for Channel '%s': %v", channel.E2NodeID, channel.ID, err)
		return err
	}

	log.Debugf("Creating E2Node entity '%s' for Channel '%s'", channel.E2NodeID, channel.ID)
	object := &topoapi.Object{
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
	log.Warnf("Creating E2Node entity '%s' for Channel '%s': %v", channel.E2NodeID, channel.ID, err)
	if err != nil {
		return err
	}

	err = r.store.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			log.Warnf("Creating E2Node entity '%s' for Channel '%s': %v", channel.E2NodeID, channel.ID, err)
			return err
		}
		return nil
	}
	return nil
}

func (r *Reconciler) createE2Cells(ctx context.Context, channel *e2server.E2Channel) error {
	for _, e2Cell := range channel.E2Cells {
		if err := r.createE2Cell(ctx, channel, e2Cell); err != nil {
			return err
		}
		if err := r.createE2CellRelation(ctx, channel, e2Cell); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reconciler) createE2Cell(ctx context.Context, channel *e2server.E2Channel, cell *topoapi.E2Cell) error {
	cellID := utils.GetCellID(channel, cell)
	object, err := r.store.Get(ctx, cellID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
			return err
		}

		log.Debugf("Creating E2Cell entity '%s' for Channel '%s'", cell.CellGlobalID.Value, channel.ID)
		object := &topoapi.Object{
			ID:   cellID,
			Type: topoapi.Object_ENTITY,
			Obj: &topoapi.Object_Entity{
				Entity: &topoapi.Entity{
					KindID: topoapi.E2CELL,
				},
			},
			Aspects: make(map[string]*gogotypes.Any),
			Labels:  map[string]string{},
		}

		err = object.SetAspect(cell)
		if err != nil {
			log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
			return err
		}

		err = r.store.Create(ctx, object)
		if err != nil {
			log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
			return err
		}
		return nil
	}

	log.Debugf("Updating E2Cell entity '%s' for Channel '%s'", cell.CellGlobalID.Value, channel.ID)
	err = object.SetAspect(cell)
	if err != nil {
		log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
		return err
	}

	err = r.store.Update(ctx, object)
	if err != nil {
		log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
		return err
	}
	return nil
}

func (r *Reconciler) createE2CellRelation(ctx context.Context, channel *e2server.E2Channel, cell *topoapi.E2Cell) error {
	cellID := utils.GetCellID(channel, cell)
	relationID := utils.GetCellRelationID(channel, cell)
	_, err := r.store.Get(ctx, relationID)
	if err == nil {
		return nil
	} else if !errors.IsNotFound(err) {
		log.Warnf("Creating E2Cell '%s' relation '%s' for Channel '%s': %v", cellID, relationID, channel.ID, err)
		return err
	}

	log.Debugf("Creating E2Cell '%s' relation '%s' for Channel '%s'", cellID, relationID, channel.ID)
	object := &topoapi.Object{
		ID:   relationID,
		Type: topoapi.Object_RELATION,
		Obj: &topoapi.Object_Relation{
			Relation: &topoapi.Relation{
				KindID:      topoapi.CONTAINS,
				SrcEntityID: channel.E2NodeID,
				TgtEntityID: cellID,
			},
		},
	}

	err = r.store.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			log.Warnf("Creating E2Cell '%s' relation '%s' for Channel '%s': %v", cellID, relationID, channel.ID, err)
			return err
		}
		return nil
	}
	return nil
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
