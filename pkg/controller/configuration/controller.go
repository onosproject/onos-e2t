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

func connectionExist(e2NodeID topoapi.ID, e2tConn *topoapi.Interface, e2NodeConns []topoapi.Interface) bool {
	for _, e2NodeConn := range e2NodeConns {
		if e2NodeConn.IP == e2tConn.IP &&
			e2NodeConn.Port == e2tConn.Port && e2NodeConn.Type == e2tConn.Type {
			log.Debugf("Connection %+v already exists for e2node: %s", e2NodeConn, e2NodeID)
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
	log.Infof("Test4 Reconciling  configuration using mgmt connection: %s", connID)
	mgmtConn, err := r.mgmtConns.Get(ctx, connID)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warn(err)
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
		return controller.Result{}, err
	}
	e2NodeID := mgmtConn.NodeID
	e2tNodes, err := r.rnib.List(ctx, utils.GetE2TFilter())
	if err != nil {
		log.Warn(err)
		return controller.Result{}, err
	}
	if len(e2tNodes) == 0 {
		return controller.Result{Requeue: id}, nil
	}
	e2Node, err := r.rnib.Get(ctx, topoapi.ID(e2NodeID))
	if err != nil {
		log.Warn(err)
		if !errors.IsNotFound(err) {
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	e2NodeConfig := &topoapi.E2NodeConfig{}
	_ = e2Node.GetAspect(e2NodeConfig)

	e2NodeConns := e2NodeConfig.Connections
	for _, e2tNode := range e2tNodes {
		log.Infof("Test4: len e2t: %d", len(e2tNodes))
		e2tNodeInfo := &topoapi.E2TInfo{}
		err := e2tNode.GetAspect(e2tNodeInfo)
		if err != nil {
			log.Warnf("Test4 Failed to reconcile configuration using management connection %s: %s", connID, err)
			return controller.Result{}, err
		}

		for _, e2tConn := range e2tNodeInfo.GetInterfaces() {
			log.Infof("Test4 connections::%+v", e2NodeConns)
			if e2tConn.Type == topoapi.Interface_INTERFACE_E2AP200 && !connectionExist(topoapi.ID(e2NodeID), e2tConn, e2NodeConns) {
				connUpdateReq := createConnectionUpdateReq(e2tConn.IP)
				log.Infof("Test4 e2tNode: %s", e2tNode.ID)
				log.Infof("Test4 new connection: %+v, %s", e2tConn, e2NodeID)
				log.Infof("Test4 Sending connection update using management connection: %s", mgmtConn.ID)
				mgmtConn, err := r.mgmtConns.Get(ctx, connID)
				if err != nil {
					if errors.IsNotFound(err) {
						log.Warn(err)
						return controller.Result{}, nil
					}
					log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
					return controller.Result{}, err
				}
				connUpdateAck, connUpdateFailure, err := mgmtConn.E2ConnectionUpdate(ctx, connUpdateReq)
				if err != nil {
					log.Warnf("Test4 Failed to reconcile configuration using management connection %s: %s", connID, err)
					return controller.Result{}, err
				}

				if connUpdateAck != nil {
					log.Infof("Received connection update ack:%+v", connUpdateAck)
					e2NodeConns = append(e2NodeConns, *e2tConn)
					e2NodeConfig.Connections = e2NodeConns
					err := e2Node.SetAspect(e2NodeConfig)
					if err != nil {
						log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
						return controller.Result{}, err
					}
					log.Infof("Test4 After updater:", e2NodeConns)

					err = r.rnib.Update(ctx, e2Node)
					if err != nil {
						log.Warnf("Test4 err", err)
						if !errors.IsNotFound(err) {
							log.Warnf("Failed to reconcile configuration using management connection %s: %s", connID, err)
							return controller.Result{}, err
						}
						return controller.Result{}, nil
					}
					return controller.Result{}, nil
				}
				if connUpdateFailure != nil {
					log.Infof("Received connection update failure: %+v", connUpdateFailure)
				}
			}
		}
	}
	log.Infof("Test4: End")

	return controller.Result{}, nil
}
