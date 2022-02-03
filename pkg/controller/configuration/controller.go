// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package configuration

import (
	"context"
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/controller/utils"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("controller", "configuration")

const (
	defaultTimeout = 30 * time.Second
)

// NewController returns a new E2 connection update controller
func NewController(rnib rnib.Store, mgmtConns e2server.MgmtConnManager, e2apConns e2server.E2APConnManager) *controller.Controller {
	c := controller.NewController("configuration")

	c.Watch(&MgmtConnWatcher{
		mgmtConns: mgmtConns,
	})

	c.Watch(&TopoWatcher{
		mgmtConns: mgmtConns,
		e2apConns: e2apConns,
		rnib:      rnib,
	})

	c.Reconcile(&Reconciler{
		mgmtConns: mgmtConns,
		e2apConns: e2apConns,
		rnib:      rnib,
	})

	return c
}

// Reconciler reconciles configuration of an E2 node
type Reconciler struct {
	mgmtConns e2server.MgmtConnManager
	e2apConns e2server.E2APConnManager
	rnib      rnib.Store
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	connID := id.Value.(e2server.ConnID)
	log.Infof("Reconciling  configuration using mgmt connection: %s", connID)
	mgmtConn, err := r.mgmtConns.Get(ctx, connID)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warn(err)
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
		return controller.Result{}, err
	}

	if ok, err := r.createE2Node(ctx, mgmtConn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2Cells(ctx, mgmtConn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	if ok, err := r.createE2CellRelations(ctx, mgmtConn); err != nil {
		return controller.Result{}, err
	} else if ok {
		return controller.Result{}, nil
	}

	e2tNodes, err := r.rnib.List(ctx, utils.GetE2TFilter())
	if err != nil {
		log.Warn(err)
		return controller.Result{}, err
	}

	var e2tNodesInterfaces []*topoapi.Interface
	for _, e2tNode := range e2tNodes {
		e2tNodeInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tNodeInfo)
		if err != nil {
			log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
			return controller.Result{}, err
		}
		for _, e2tIface := range e2tNodeInfo.Interfaces {
			if e2tIface.Type == topoapi.Interface_INTERFACE_E2AP200 {
				e2tNodesInterfaces = append(e2tNodesInterfaces, e2tIface)
			}
		}
	}

	// Creates list of connections that should be added by the E2 node
	connToAddList := getConnToAddList(mgmtConn, e2tNodesInterfaces)
	connToRemoveList := getConnToRemoveList(mgmtConn, e2tNodesInterfaces)

	if len(connToAddList) > 0 || len(connToRemoveList) > 0 {
		connAddList := createConnectionAddListIE(connToAddList)
		log.Debugf("Connection To Add List for e2 node %s, %+v", mgmtConn.E2NodeID, connAddList)

		connRemoveList := createConnectionRemoveList(connToRemoveList)
		log.Debugf("Connection To Remove List for e2 node %s, %+v", mgmtConn.E2NodeID, connRemoveList)

		// Creates a connection update request to include list of the connections
		// that should be added and list of connections that should be removed
		connUpdateReq := NewConnectionUpdate(
			WithConnectionAddList(connAddList),
			WithConnectionRemoveList(connRemoveList),
			WithTransactionID(3)).
			Build()

		log.Infof("Sending connection update request for e2Node: %s, %+v", mgmtConn.E2NodeID, connUpdateReq)
		connUpdateAck, connUpdateFailure, err := mgmtConn.E2ConnectionUpdate(ctx, connUpdateReq)
		if err != nil {
			log.Warnf("Failed to reconcile configuration for e2 node %s using management connection %s: %s", mgmtConn.E2NodeID, connID, err)
			return controller.Result{}, err
		}

		if connUpdateAck != nil {
			log.Infof("Received connection update ack for e2 node %s:%+v", mgmtConn.E2NodeID, connUpdateAck)
			mgmtConn.E2NodeConfig.Connections = append(mgmtConn.E2NodeConfig.Connections, connToAddList...)
			return controller.Result{}, nil
		}
		if connUpdateFailure != nil {
			// TODO returns an appropriate error to retry
			log.Infof("Received connection update failure: %+v", connUpdateFailure)

		}
	}
	return controller.Result{}, nil
}

func (r *Reconciler) createE2Node(ctx context.Context, conn *e2server.ManagementConn) (bool, error) {
	log.Debugf("Creating E2 node %s for connection %v", conn.E2NodeID, conn.ID)
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
		log.Debugf("E2 node %s aspect is already set ", conn.E2NodeID)
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

func (r *Reconciler) createE2CellRelations(ctx context.Context, conn *e2server.ManagementConn) (bool, error) {
	for _, e2Cell := range conn.E2Cells {
		if ok, err := r.createE2CellRelation(ctx, conn, e2Cell); err != nil {
			return false, err
		} else if ok {
			return true, nil
		}
	}
	return false, nil
}

func (r *Reconciler) createE2Cells(ctx context.Context, conn *e2server.ManagementConn) (bool, error) {
	for _, e2Cell := range conn.E2Cells {
		if ok, err := r.createE2Cell(ctx, conn, e2Cell); err != nil {
			return false, err
		} else if ok {
			return true, nil
		}
	}
	return false, nil
}

func (r *Reconciler) createE2Cell(ctx context.Context, conn *e2server.ManagementConn, cell *topoapi.E2Cell) (bool, error) {
	cellID := e2server.GetCellID(conn, cell)
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
		log.Debugf("E2 cell %s aspect is already set", cellID)
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

func (r *Reconciler) createE2CellRelation(ctx context.Context, conn *e2server.ManagementConn, cell *topoapi.E2Cell) (bool, error) {
	cellID := e2server.GetCellID(conn, cell)
	relationID := e2server.GetCellRelationID(conn, cell)
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
