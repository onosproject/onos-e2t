// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2node

import (
	"context"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const (
	defaultTimeout = 30 * time.Second
)

var log = logging.GetLogger("controller", "channel")

// NewController returns a new E2 control relation controller
func NewController(rnib rnib.Store, channels e2server.ChannelManager) *controller.Controller {
	c := controller.NewController("channel")
	c.Watch(&Watcher{
		channels: channels,
	})

	c.Watch(&TopoWatcher{
		topo:     rnib,
		channels: channels,
	})

	c.Reconcile(&Reconciler{
		channels: channels,
		rnib:     rnib,
	})

	return c
}

// Reconciler is for reconciling RAN entities such as E2 node , E2 cell and their relations
type Reconciler struct {
	channels e2server.ChannelManager
	rnib     rnib.Store
}

func (r *Reconciler) createE2ControlRelation(ctx context.Context, channel *e2server.E2Channel) (bool, error) {
	relationID := utils.GetE2ControlRelationID(channel.ID)
	_, err := r.rnib.Get(ctx, relationID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Node '%s' control relation '%s' failed: %v", channel.E2NodeID, relationID, err)
			return false, err
		}
		log.Debugf("Creating E2Node '%s' control relation '%s'", channel.E2NodeID, relationID)
		object := &topoapi.Object{
			ID:   relationID,
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.CONTROLS,
					SrcEntityID: utils.GetE2TID(),
					TgtEntityID: channel.E2NodeID,
				},
			},
		}
		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Node '%s' control relation '%s' failed: %v", channel.E2NodeID, relationID, err)
				return false, err
			}
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	channelID := id.Value.(e2server.ChannelID)
	log.Infof("Reconciling E2 node Control relation for channel: %s", channelID)
	channel, err := r.channels.Get(ctx, channelID)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.reconcileDeleteE2ControlRelation(channelID)
		}
		log.Warnf("Failed to reconcile E2 node control relation for channel %s: %s", channelID, err)
		return controller.Result{}, err
	}

	if ok, err := r.createE2Node(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2Cells(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2CellRelations(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2ControlRelation(ctx, channel); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) createE2Node(ctx context.Context, channel *e2server.E2Channel) (bool, error) {
	log.Debug("Creating E2 node %s for channel %v", channel.E2NodeID, channel.ID)
	object, err := r.rnib.Get(ctx, channel.E2NodeID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Node entity '%s' for Channel '%s': %v", channel.E2NodeID, channel.ID, err)
			return false, err
		}
		log.Debugf("Creating E2Node entity '%s' for Channel '%s'", channel.E2NodeID, channel.ID)
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

		aspect := &topoapi.E2Node{
			ServiceModels: channel.ServiceModels,
		}

		err = object.SetAspect(aspect)
		if err != nil {
			log.Warnf("Creating E2Node entity '%s' for Channel failed '%s': %v", channel.E2NodeID, channel.ID, err)
			return false, err
		}

		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Node entity '%s' for Channel '%s': %v", channel.E2NodeID, channel.ID, err)
				return false, err
			}
			return false, nil
		}
		return true, nil
	}

	e2NodeAspect := &topoapi.E2Node{}
	err = object.GetAspect(e2NodeAspect)
	if err == nil {
		log.Debug("E2 node %s aspect is already set ", channel.E2NodeID)
		return false, nil
	}

	e2NodeAspect = &topoapi.E2Node{
		ServiceModels: channel.ServiceModels,
	}
	err = object.SetAspect(e2NodeAspect)
	if err != nil {
		return false, err
	}
	err = r.rnib.Update(ctx, object)
	if err != nil {
		if !errors.IsNotFound(err) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) createE2CellRelations(ctx context.Context, channel *e2server.E2Channel) (bool, error) {
	for _, e2Cell := range channel.E2Cells {
		if ok, err := r.createE2CellRelation(ctx, channel, e2Cell); err != nil {
			return false, err
		} else if ok {
			return true, nil
		}
	}
	return false, nil
}

func (r *Reconciler) createE2Cells(ctx context.Context, channel *e2server.E2Channel) (bool, error) {
	for _, e2Cell := range channel.E2Cells {
		if ok, err := r.createE2Cell(ctx, channel, e2Cell); err != nil {
			return false, err
		} else if ok {
			return true, nil
		}
	}
	return false, nil
}

func (r *Reconciler) createE2Cell(ctx context.Context, channel *e2server.E2Channel, cell *topoapi.E2Cell) (bool, error) {
	cellID := utils.GetCellID(channel, cell)
	object, err := r.rnib.Get(ctx, cellID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
			return false, err
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
			return false, err
		}

		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
				return false, err
			}
			return false, nil
		}
		return true, nil
	}

	e2CellAspect := &topoapi.E2Cell{}
	err = object.GetAspect(e2CellAspect)
	if err == nil {
		log.Debug("E2 cell %s aspect is already set", cellID)
		return false, nil
	}

	log.Debugf("Updating E2Cell entity '%s' for Channel '%s'", cell.CellGlobalID.Value, channel.ID)
	err = object.SetAspect(cell)
	if err != nil {
		log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
		return false, err
	}

	err = r.rnib.Update(ctx, object)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell entity '%s' for Channel '%s': %v", cell.CellGlobalID.Value, channel.ID, err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) createE2CellRelation(ctx context.Context, channel *e2server.E2Channel, cell *topoapi.E2Cell) (bool, error) {
	cellID := utils.GetCellID(channel, cell)
	relationID := utils.GetCellRelationID(channel, cell)
	_, err := r.rnib.Get(ctx, relationID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell '%s' relation '%s' for Channel '%s': %v", cellID, relationID, channel.ID, err)
			return false, err
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

		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Cell '%s' relation '%s' for Channel '%s': %v", cellID, relationID, channel.ID, err)
				return false, err
			}
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

func (r *Reconciler) deleteE2ControlRelation(ctx context.Context, channelID e2server.ChannelID) error {
	relationID := utils.GetE2ControlRelationID(channelID)
	log.Debugf("Deleting E2Node relation '%s' for Channel '%s'", relationID, channelID)
	object, err := r.rnib.Get(ctx, relationID)
	if err == nil {
		err := r.rnib.Delete(ctx, object)
		if err != nil {
			if !errors.IsNotFound(err) {
				log.Warnf("Deleting E2Node relation '%s' for Channel '%s' failed: %v", relationID, channelID, err)
				return err
			}
			return nil
		}
	}

	return nil
}

func (r *Reconciler) reconcileDeleteE2ControlRelation(channelID e2server.ChannelID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.deleteE2ControlRelation(ctx, channelID); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}
