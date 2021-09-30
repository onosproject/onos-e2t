// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package controlrelation

import (
	"context"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("controller", "controlrelation")

const (
	defaultTimeout = 30 * time.Second
)

// NewController returns a new E2 connection update controller
func NewController(rnib rnib.Store, e2apConns e2server.E2APConnManager) *controller.Controller {
	c := controller.NewController("control_relation")

	c.Watch(&E2APConnWatcher{
		e2apConns: e2apConns,
	})

	c.Reconcile(&Reconciler{
		e2apConns: e2apConns,
		rnib:      rnib,
	})

	return c
}

// Reconciler reconciles rnib control relations
type Reconciler struct {
	e2apConns e2server.E2APConnManager
	rnib      rnib.Store
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	connID := id.Value.(e2server.ConnID)
	log.Infof("Reconciling E2 node Control relation for connection: %s", connID)
	_, err := r.e2apConns.Get(ctx, connID)
	if err != nil {
		if errors.IsNotFound(err) {
			return r.reconcileDeleteE2ControlRelation(connID)
		}
		log.Warnf("Failed to reconcile E2 node control relation for connection %s: %s", connID, err)
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) deleteE2ControlRelation(ctx context.Context, connID e2server.ConnID) error {
	relationID := e2server.GetE2ControlRelationID(connID)
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

func (r *Reconciler) updateE2NodeConfig(ctx context.Context, connID e2server.ConnID) error {
	relationID := e2server.GetE2ControlRelationID(connID)
	object, err := r.rnib.Get(ctx, relationID)
	if err == nil {
		relation := object.Obj.(*topoapi.Object_Relation)
		e2tID := relation.Relation.GetSrcEntityID()
		e2tObject, err := r.rnib.Get(ctx, e2tID)
		if err == nil {
			e2tInfoAspect := &topoapi.E2TInfo{}
			err = e2tObject.GetAspect(e2tInfoAspect)
			if err != nil {
				return err
			}
			e2tInterfaces := e2tInfoAspect.Interfaces
			e2NodeID := relation.Relation.GetTgtEntityID()
			e2NodeObject, err := r.rnib.Get(ctx, e2NodeID)
			if err == nil {
				e2NodeConfig := &topoapi.E2NodeConfig{}
				err = e2NodeObject.GetAspect(e2NodeConfig)
				if err != nil {
					return err
				}
				var e2NodeConns []topoapi.Interface
				log.Infof("Updating E2 node config for e2 Node: %s", e2NodeID)
				for _, e2tInterface := range e2tInterfaces {
					e2NodeConns := e2NodeConfig.Connections
					for _, e2NodeConn := range e2NodeConns {
						if e2NodeConn.IP == e2tInterface.IP && e2NodeConn.Type == e2tInterface.Type &&
							e2NodeConn.Port == e2tInterface.Port {
							continue
						}
						e2NodeConns = append(e2NodeConns, topoapi.Interface{
							IP:   e2NodeConn.IP,
							Port: e2NodeConn.Port,
							Type: e2NodeConn.Type,
						})
					}
				}
				e2NodeConfig.Connections = e2NodeConns
				err := e2NodeObject.SetAspect(e2NodeConfig)
				if err != nil {
					return err
				}
				err = r.rnib.Update(ctx, e2NodeObject)
				if err != nil {
					if errors.IsNotFound(err) {
						return nil
					}
					return err
				}
			}
		}
	}
	return nil

}

func (r *Reconciler) reconcileDeleteE2ControlRelation(connID e2server.ConnID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.updateE2NodeConfig(ctx, connID); err != nil {
		return controller.Result{}, err
	}

	if err := r.deleteE2ControlRelation(ctx, connID); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}
