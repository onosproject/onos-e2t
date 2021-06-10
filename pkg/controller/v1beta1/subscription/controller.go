// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"time"

	api "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
	"github.com/onosproject/onos-e2t/pkg/store/task"
	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("controller", "subscription")

const defaultTimeout = 30 * time.Second

// NewController returns a new network controller
func NewController(subs subscription.Store, tasks task.Store) *controller.Controller {
	c := controller.NewController("Subscription")
	c.Watch(&Watcher{
		subs: subs,
	})
	c.Watch(&TaskWatcher{
		subs:  subs,
		tasks: tasks,
	})
	c.Reconcile(&Reconciler{
		subs:  subs,
		tasks: tasks,
	})
	return c
}

// Reconciler is a device change reconciler
type Reconciler struct {
	subs  subscription.Store
	tasks task.Store
}

// Reconcile reconciles the state of a device change
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	sub, err := r.subs.Get(ctx, id.Value.(api.SubscriptionID))
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		return controller.Result{}, err
	}

	log.Infof("Reconciling Subscription %+v", sub)

	switch sub.Status.State {
	case api.SubscriptionState_SUBSCRIPTION_ACTIVE:
		return r.reconcileActiveSubscription(sub)
	case api.SubscriptionState_SUBSCRIPTION_PENDING_DELETE:
		return r.reconcileDeletedSubscription(sub)
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileActiveSubscription(sub *api.Subscription) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	_, err := r.tasks.Get(ctx, sub.ID.TaskID())
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
			return controller.Result{}, err
		}

		log.Warnf("Creating Task for Subscription %+v: %s", sub, err)
		task := &api.Task{
			TaskMeta: api.TaskMeta{
				ID:           sub.ID.TaskID(),
				ServiceModel: sub.ServiceModel,
			},
			Spec: api.TaskSpec{
				Subscription: sub.Spec,
			},
		}
		err := r.tasks.Create(ctx, task)
		if err != nil {
			if !errors.IsAlreadyExists(err) {
				log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
				return controller.Result{}, err
			}
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}

func (r *Reconciler) reconcileDeletedSubscription(sub *api.Subscription) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	task, err := r.tasks.Get(ctx, sub.ID.TaskID())
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}

	err = r.tasks.Delete(ctx, task)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Failed to reconcile Subscription %+v: %s", sub, err)
			return controller.Result{}, err
		}
		return controller.Result{}, nil
	}
	return controller.Result{}, nil
}
