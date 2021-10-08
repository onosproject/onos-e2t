// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package configuration

import (
	"context"
	"time"

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
