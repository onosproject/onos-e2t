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
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const (
	defaultTimeout = 30 * time.Second
)

var log = logging.GetLogger("controller", "e2node")

// NewController returns a new E2 control relation controller
func NewController(rnib rnib.Store, conns e2server.ConnManager) *controller.Controller {
	c := controller.NewController("e2node")
	c.Watch(&Watcher{
		conns: conns,
	})

	c.Watch(&TopoWatcher{
		topo:  rnib,
		conns: conns,
	})

	c.Reconcile(&Reconciler{
		conns: conns,
		rnib:  rnib,
	})

	return c
}

// Reconciler is for reconciling RAN entities such as E2 node , E2 cell and their relations
type Reconciler struct {
	conns e2server.ConnManager
	rnib  rnib.Store
}

func (r *Reconciler) createE2ControlRelation(ctx context.Context, conn *e2server.E2APConn) (bool, error) {
	relationID := utils.GetE2ControlRelationID(conn.ID)
	_, err := r.rnib.Get(ctx, relationID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Node '%s' control relation '%s' failed: %v", conn.E2NodeID, relationID, err)
			return false, err
		}
		log.Debugf("Creating E2Node '%s' control relation '%s'", conn.E2NodeID, relationID)
		object := &topoapi.Object{
			ID:   relationID,
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.CONTROLS,
					SrcEntityID: utils.GetE2TID(),
					TgtEntityID: conn.E2NodeID,
				},
			},
		}
		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Node '%s' control relation '%s' failed: %v", conn.E2NodeID, relationID, err)
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

	connID := id.Value.(e2server.ConnID)
	log.Infof("Reconciling E2 node Control relation for connection: %s", connID)
	conn, err := r.conns.Get(ctx, connID)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.reconcileDeleteE2ControlRelation(connID)
		}
		log.Warnf("Failed to reconcile E2 node control relation for connection %s: %s", connID, err)
		return controller.Result{}, err
	}

	if ok, err := r.createE2Node(ctx, conn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2Cells(ctx, conn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2CellRelations(ctx, conn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2ControlRelation(ctx, conn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) createE2Node(ctx context.Context, conn *e2server.E2APConn) (bool, error) {
	log.Debug("Creating E2 node %s for connection %v", conn.E2NodeID, conn.ID)
	object, err := r.rnib.Get(ctx, conn.E2NodeID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Node entity '%s' for connection '%s': %v", conn.E2NodeID, conn.ID, err)
			return false, err
		}
		log.Debugf("Creating E2Node entity '%s' for connection '%s'", conn.E2NodeID, conn.ID)
		object = &topoapi.Object{
			ID:   conn.E2NodeID,
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
			ServiceModels: conn.ServiceModels,
		}

		err = object.SetAspect(aspect)
		if err != nil {
			log.Warnf("Creating E2Node entity '%s' for connection failed '%s': %v", conn.E2NodeID, conn.ID, err)
			return false, err
		}

		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Node entity '%s' for connection '%s': %v", conn.E2NodeID, conn.ID, err)
				return false, err
			}
			return false, nil
		}
		return true, nil
	}

	e2NodeAspect := &topoapi.E2Node{}
	err = object.GetAspect(e2NodeAspect)
	if err == nil {
		log.Debug("E2 node %s aspect is already set ", conn.E2NodeID)
		return false, nil
	}

	e2NodeAspect = &topoapi.E2Node{
		ServiceModels: conn.ServiceModels,
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

func (r *Reconciler) createE2CellRelations(ctx context.Context, conn *e2server.E2APConn) (bool, error) {
	for _, e2Cell := range conn.E2Cells {
		if ok, err := r.createE2CellRelation(ctx, conn, e2Cell); err != nil {
			return false, err
		} else if ok {
			return true, nil
		}
	}
	return false, nil
}

func (r *Reconciler) createE2Cells(ctx context.Context, conn *e2server.E2APConn) (bool, error) {
	for _, e2Cell := range conn.E2Cells {
		if ok, err := r.createE2Cell(ctx, conn, e2Cell); err != nil {
			return false, err
		} else if ok {
			return true, nil
		}
	}
	return false, nil
}

func (r *Reconciler) createE2Cell(ctx context.Context, conn *e2server.E2APConn, cell *topoapi.E2Cell) (bool, error) {
	cellID := utils.GetCellID(conn, cell)
	object, err := r.rnib.Get(ctx, cellID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell entity '%s' for connection '%s': %v", cell.CellGlobalID.Value, conn.ID, err)
			return false, err
		}

		log.Debugf("Creating E2Cell entity '%s' for connection '%s'", cell.CellGlobalID.Value, conn.ID)
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
			log.Warnf("Creating E2Cell entity '%s' for connection '%s': %v", cell.CellGlobalID.Value, conn.ID, err)
			return false, err
		}

		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Cell entity '%s' for connection '%s': %v", cell.CellGlobalID.Value, conn.ID, err)
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

	log.Debugf("Updating E2Cell entity '%s' for connection '%s'", cell.CellGlobalID.Value, conn.ID)
	err = object.SetAspect(cell)
	if err != nil {
		log.Warnf("Creating E2Cell entity '%s' for connection '%s': %v", cell.CellGlobalID.Value, conn.ID, err)
		return false, err
	}

	err = r.rnib.Update(ctx, object)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell entity '%s' for connection '%s': %v", cell.CellGlobalID.Value, conn.ID, err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Reconciler) createE2CellRelation(ctx context.Context, conn *e2server.E2APConn, cell *topoapi.E2Cell) (bool, error) {
	cellID := utils.GetCellID(conn, cell)
	relationID := utils.GetCellRelationID(conn, cell)
	_, err := r.rnib.Get(ctx, relationID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Creating E2Cell '%s' relation '%s' for connection '%s': %v", cellID, relationID, conn.ID, err)
			return false, err
		}
		log.Debugf("Creating E2Cell '%s' relation '%s' for connection '%s'", cellID, relationID, conn.ID)
		object := &topoapi.Object{
			ID:   relationID,
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.CONTAINS,
					SrcEntityID: conn.E2NodeID,
					TgtEntityID: cellID,
				},
			},
		}

		err = r.rnib.Create(ctx, object)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Creating E2Cell '%s' relation '%s' for connection '%s': %v", cellID, relationID, conn.ID, err)
				return false, err
			}
			return false, nil
		}
		return true, nil
	}
	return false, nil
}

func (r *Reconciler) deleteE2ControlRelation(ctx context.Context, connID e2server.ConnID) error {
	relationID := utils.GetE2ControlRelationID(connID)
	log.Debugf("Deleting E2Node relation '%s' for connection '%s'", relationID, connID)
	object, err := r.rnib.Get(ctx, relationID)
	if err == nil {
		err := r.rnib.Delete(ctx, object)
		if err != nil {
			if !errors.IsNotFound(err) {
				log.Warnf("Deleting E2Node relation '%s' for connection '%s' failed: %v", relationID, connID, err)
				return err
			}
			return nil
		}
	}

	return nil
}

func (r *Reconciler) reconcileDeleteE2ControlRelation(connID e2server.ConnID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.deleteE2ControlRelation(ctx, connID); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}
