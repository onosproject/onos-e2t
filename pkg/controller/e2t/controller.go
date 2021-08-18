// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2t

import (
	"context"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/env"

	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/controller/utils"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const (
	defaultTimeout            = 30 * time.Second
	defaultGRPCPort           = 5150
	defaultE2APPort           = 36421
	defaultExpirationDuration = 10
)

var log = logging.GetLogger("controller", "e2t")

// NewController returns a new E2T controller
func NewController(rnib rnib.Store) *controller.Controller {
	c := controller.NewController("E2T")
	c.Watch(&Watcher{
		rnib: rnib,
	})

	c.Reconcile(&Reconciler{
		rnib: rnib,
	})

	return c
}

// Reconciler is an E2T reconciler
type Reconciler struct {
	rnib rnib.Store
}

func (r *Reconciler) createE2T(ctx context.Context, e2tID topoapi.ID) error {
	_, err := r.rnib.Get(ctx, e2tID)
	if err == nil {
		return nil
	} else if !errors.IsNotFound(err) {
		log.Infof("Creating E2T entity failed: %v", err)
		return err
	}

	log.Debugf("Creating E2T entity")
	object := &topoapi.Object{
		ID:   utils.GetE2TID(),
		Type: topoapi.Object_ENTITY,
		Obj: &topoapi.Object_Entity{
			Entity: &topoapi.Entity{
				KindID: topoapi.E2T,
			},
		},
		Aspects: make(map[string]*gogotypes.Any),
		Labels:  map[string]string{},
	}
	interfaces := make([]*topoapi.Interface, 2)
	interfaces[0] = &topoapi.Interface{
		IP:   env.GetPodIP(),
		Port: defaultE2APPort,
		Type: topoapi.Interface_INTERFACE_E2AP101,
	}

	interfaces[1] = &topoapi.Interface{
		IP:   env.GetPodIP(),
		Port: defaultGRPCPort,
		Type: topoapi.Interface_INTERFACE_E2T,
	}

	e2tAspect := &topoapi.E2TInfo{
		Interfaces: interfaces,
	}

	var expiration *time.Time
	t := time.Now().Add(defaultExpirationDuration)
	expiration = &t

	leaseAspect := &topoapi.Lease{
		Expiration: expiration,
	}

	err = object.SetAspect(e2tAspect)
	if err != nil {
		return err
	}

	err = object.SetAspect(leaseAspect)
	if err != nil {
		return err
	}

	err = r.rnib.Create(ctx, object)
	if err != nil {
		if !errors.IsAlreadyExists(err) {
			log.Infof("Creating E2T entity failed: %v", err)
			return err
		}
		return nil
	}
	return nil
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	e2tID := id.Value.(topoapi.ID)
	log.Infof("Reconciling E2T entity with ID: %s", e2tID)
	if err := r.createE2T(ctx, e2tID); err != nil {
		return controller.Result{}, err
	}

	return controller.Result{}, nil
}
