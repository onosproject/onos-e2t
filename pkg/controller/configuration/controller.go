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

func connectionExist(e2tConn *topoapi.Interface, mgmtConn *e2server.ManagementConn) bool {
	for _, e2NodeConn := range mgmtConn.E2NodeConfig.Connections {
		if e2NodeConn.IP == e2tConn.IP &&
			e2NodeConn.Port == e2tConn.Port && e2NodeConn.Type == e2tConn.Type {
			log.Debugf("Connection %+v already exists for e2node: %s", e2NodeConn, mgmtConn.E2NodeID)
			return true
		}
	}
	log.Debugf("Connection %+v does not exists", e2tConn)
	return false
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

	for _, e2tNode := range e2tNodes {
		e2tNodeInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tNodeInfo)
		if err != nil {
			log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
			return controller.Result{}, err
		}

		missingConnList := getMissingConnList(mgmtConn, e2tNodeInfo.Interfaces)
		log.Debugf("Missing connection list: %+v", missingConnList)
		connAddList := createConnectionAddListIE(missingConnList)

		// Creates a connection update request to include list of the connections
		// that should be added and list of connections that should be removed
		connUpdateReq := NewConnectionUpdate(
			WithConnectionAddList(connAddList)).
			Build()

		log.Infof("Sending connection update request: %+v", connUpdateReq)
		connUpdateAck, connUpdateFailure, err := mgmtConn.E2ConnectionUpdate(ctx, connUpdateReq)
		if err != nil {
			log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
			return controller.Result{}, err
		}

		if connUpdateAck != nil {
			log.Infof("Received connection update ack:%+v", connUpdateAck)
			for _, e2tConn := range missingConnList {
				mgmtConn.E2NodeConfig.Connections = append(mgmtConn.E2NodeConfig.Connections, e2tConn)
			}
			return controller.Result{}, nil
		}
		if connUpdateFailure != nil {
			log.Infof("Received connection update failure: %+v", connUpdateFailure)

		}
	}
	return controller.Result{}, nil
}
