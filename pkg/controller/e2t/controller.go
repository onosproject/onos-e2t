// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

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
	defaultExpirationDuration = 30 * time.Second
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
	log.Debugf("Creating E2T entity %s", e2tID)
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
		Type: topoapi.Interface_INTERFACE_E2AP200,
	}

	interfaces[1] = &topoapi.Interface{
		IP:   env.GetPodIP(),
		Port: defaultGRPCPort,
		Type: topoapi.Interface_INTERFACE_E2T,
	}

	e2tAspect := &topoapi.E2TInfo{
		Interfaces: interfaces,
	}

	expiration := time.Now().Add(defaultExpirationDuration)
	leaseAspect := &topoapi.Lease{
		Expiration: &expiration,
	}

	err := object.SetAspect(e2tAspect)
	if err != nil {
		return err
	}

	err = object.SetAspect(leaseAspect)
	if err != nil {
		return err
	}

	err = r.rnib.Create(ctx, object)
	if err != nil && !errors.IsAlreadyExists(err) {
		log.Infof("Creating E2T entity %s failed: %v", e2tID, err)
		return err
	}

	return nil
}

func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	e2tID := id.Value.(topoapi.ID)
	log.Infof("Reconciling E2T entity with ID: %s", e2tID)

	object, err := r.rnib.Get(ctx, e2tID)
	if err == nil {
		//  Reconciles an E2T entity that’s not local so the controller should requeue
		//  it for the lease expiration time and delete the entity if the lease has not been renewed
		if e2tID != utils.GetE2TID() {
			lease := &topoapi.Lease{}
			_ = object.GetAspect(lease)

			// Check if the the lease is expired
			if lease.Expiration.Before(time.Now()) {
				log.Debugf("Deleting the expired lease for E2T with ID: %s", e2tID)
				err := r.rnib.Delete(ctx, object)
				if !errors.IsNotFound(err) {
					return controller.Result{}, err
				}
				return controller.Result{}, nil
			}

			// Requeue the object to be reconciled at the expiration time
			return controller.Result{
				RequeueAfter: time.Until(*lease.Expiration),
			}, nil
		}

		// Renew the lease If this is the E2T entity for the local node
		if e2tID == utils.GetE2TID() {
			lease := &topoapi.Lease{}

			err := object.GetAspect(lease)
			if err != nil {
				return controller.Result{}, err
			}

			remainingTime := time.Until(*lease.GetExpiration())
			// If the remaining time of lease is more than  half the lease duration, no need to renew the lease
			// schedule the next renewal
			if remainingTime > defaultExpirationDuration/2 {
				log.Debugf("No need to renew the lease for %s, the remaining lease time is %v seconds", e2tID, remainingTime)
				return controller.Result{
					RequeueAfter: time.Until(lease.Expiration.Add(defaultExpirationDuration / 2 * -1)),
				}, nil
			}

			// Renew the release to trigger the reconciler
			log.Debugf("Renew the lease for E2T with ID: %s", e2tID)
			expiration := time.Now().Add(defaultExpirationDuration)
			lease = &topoapi.Lease{
				Expiration: &expiration,
			}

			err = object.SetAspect(lease)
			if err != nil {
				return controller.Result{}, err
			}
			err = r.rnib.Update(ctx, object)
			if !errors.IsNotFound(err) {
				return controller.Result{}, err
			}
			return controller.Result{}, nil
		}

	} else if !errors.IsNotFound(err) {
		log.Infof("Renewing E2T entity lease failed for E2T with ID %s: %v", e2tID, err)
		return controller.Result{}, err
	}

	// Create the E2T entity
	if err := r.createE2T(ctx, e2tID); err != nil {
		log.Infof("Creating E2T entity with ID %s failed: %v", e2tID, err)
		return controller.Result{}, err
	}

	return controller.Result{}, nil
}
