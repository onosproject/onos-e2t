// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"context"
	"crypto/md5"
	"fmt"
	gogotypes "github.com/gogo/protobuf/types"
	uuid2 "github.com/google/uuid"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/uri"
	"math/rand"
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

	if err := r.createE2T(ctx); err != nil {
		return controller.Result{}, err
	}

	if err := r.createE2Node(ctx, channel); err != nil {
		return controller.Result{}, err
	}

	if err := r.createE2Cells(ctx, channel); err != nil {
		return controller.Result{}, err
	}

	if err := r.createE2NodeRelation(ctx, channel); err != nil {
		return controller.Result{}, err
	}

	if err := r.updateE2NodeMaster(ctx, channel); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) createE2T(ctx context.Context) error {
	_, err := r.store.Get(ctx, getE2TID())
	if err == nil {
		return nil
	} else if !errors.IsNotFound(err) {
		log.Infof("Creating E2T entity failed: %v", err)
		return err
	}

	log.Debugf("Creating E2T entity")
	object := &topoapi.Object{
		ID:   getE2TID(),
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
			log.Infof("Creating E2T entity failed: %v", err)
			return err
		}
		return nil
	}
	return nil
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
	cellID := getCellID(channel, cell)
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
	cellID := getCellID(channel, cell)
	relationID := getCellRelationID(channel, cell)
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

func (r *Reconciler) createE2NodeRelation(ctx context.Context, channel *e2server.E2Channel) error {
	relationID := getE2RelationID(channel.ID)
	_, err := r.store.Get(ctx, relationID)
	if err == nil {
		return nil
	} else if !errors.IsNotFound(err) {
		log.Warnf("Creating E2Node '%s' relation '%s' failed: %v", channel.E2NodeID, relationID, err)
		return err
	}

	log.Debugf("Creating E2Node '%s' relation '%s'", channel.E2NodeID, relationID)
	object := &topoapi.Object{
		ID:   relationID,
		Type: topoapi.Object_RELATION,
		Obj: &topoapi.Object_Relation{
			Relation: &topoapi.Relation{
				KindID:      topoapi.CONTROLS,
				SrcEntityID: getE2TID(),
				TgtEntityID: channel.E2NodeID,
			},
		},
	}

	err = r.store.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			log.Warnf("Creating E2Node '%s' relation '%s' failed: %v", channel.E2NodeID, relationID, err)
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

func (r *Reconciler) reconcileClosedChannel(channelID e2server.ChannelID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.deleteE2Relation(ctx, channelID); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) deleteE2Relation(ctx context.Context, channelID e2server.ChannelID) error {
	relationID := getE2RelationID(channelID)
	log.Debugf("Deleting E2Node relation '%s' for Channel '%s'", relationID, channelID)
	err := r.store.Delete(ctx, relationID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Deleting E2Node relation '%s' for Channel '%s' failed: %v", relationID, channelID, err)
			return err
		}
		return nil
	}
	return nil
}

func getE2TID() topoapi.ID {
	return topoapi.ID(uri.NewURI(
		uri.WithScheme("e2"),
		uri.WithOpaque(env.GetPodID())).String())
}

func getE2RelationID(channelID e2server.ChannelID) topoapi.ID {
	return topoapi.ID(channelID)
}

func getCellID(channel *e2server.E2Channel, cell *topoapi.E2Cell) topoapi.ID {
	return topoapi.ID(uri.NewURI(uri.WithOpaque(fmt.Sprintf("%s/%s", channel.E2NodeID, cell.CellGlobalID.Value))).String())
}

func getCellRelationID(channel *e2server.E2Channel, cell *topoapi.E2Cell) topoapi.ID {
	bytes := md5.Sum([]byte(fmt.Sprintf("%s/%s", channel.E2NodeID, cell.CellGlobalID.Value)))
	uuid, err := uuid2.FromBytes(bytes[:])
	if err != nil {
		panic(err)
	}
	return topoapi.ID(uri.NewURI(
		uri.WithScheme("uuid"),
		uri.WithOpaque(uuid.String())).String())
}
