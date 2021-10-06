// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package controlrelation

import (
	"context"
	"time"

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
	log.Debugf("Deleting  controls relation '%s' for connection '%s'", relationID, connID)
	object, err := r.rnib.Get(ctx, relationID)
	if err == nil {
		err := r.rnib.Delete(ctx, object)
		if err != nil {
			if !errors.IsNotFound(err) {
				log.Warnf("Deleting control relation '%s' for connection '%s' failed: %v", relationID, connID, err)
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
		log.Warnf("Failed to delete control relation for connection: %s, %s", connID, err)
		return controller.Result{}, err
	}

	return controller.Result{}, nil
}
