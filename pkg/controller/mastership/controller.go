// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package mastership

import (
	"context"
	"math/rand"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-lib-go/pkg/controller"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultTimeout = 30 * time.Second

var log = logging.GetLogger("controller", "mastership")

// NewController returns a new mastership controller
func NewController(rnib rnib.Store) *controller.Controller {
	c := controller.NewController("mastership")
	c.Watch(&TopoWatcher{
		topo: rnib,
	})

	c.Reconcile(&Reconciler{
		rnib: rnib,
	})
	return c
}

// Reconciler is a device change reconciler
type Reconciler struct {
	rnib rnib.Store
}

// Reconcile reconciles and mastership election
func (r *Reconciler) Reconcile(id controller.ID) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	e2NodeID := id.Value.(topoapi.ID)
	log.Infof("Reconciling mastership election for e2 Node  %s", e2NodeID)
	e2Node, err := r.rnib.Get(ctx, e2NodeID)
	if err != nil {
		if errors.IsNotFound(err) {
			return controller.Result{}, nil
		}
		log.Warnf("Failed to reconcile mastership election for e2 node  with ID %s: %s", e2Node, err)
		return controller.Result{}, err
	}
	return r.reconcileMastershipElection(e2Node)
}

func (r *Reconciler) reconcileMastershipElection(e2Node *topoapi.Object) (controller.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	if err := r.updateE2NodeMaster(ctx, e2Node); err != nil {
		return controller.Result{}, err
	}
	return controller.Result{}, nil
}

func (r *Reconciler) updateE2NodeMaster(ctx context.Context, e2Node *topoapi.Object) error {
	log.Debugf("Verifying mastership for E2Node '%s'", e2Node.GetID())
	e2NodeEntity, err := r.rnib.Get(ctx, e2Node.GetID())
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", e2Node.GetID(), err)
			return err
		}
		log.Warnf("E2Node entity '%s' not found", e2Node.GetID())
		return nil
	}

	// List the objects in the topo store
	objects, err := r.rnib.List(ctx, nil)
	if err != nil {
		log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", e2Node.GetID(), err)
		return err
	}

	// Filter the topo objects for relations
	e2NodeRelations := make(map[string]topoapi.Object)
	for _, object := range objects {
		if relation, ok := object.Obj.(*topoapi.Object_Relation); ok &&
			relation.Relation.KindID == topoapi.CONTROLS &&
			relation.Relation.TgtEntityID == e2Node.GetID() {
			e2NodeRelations[string(object.ID)] = object
		}
	}

	mastership := &topoapi.MastershipState{}
	_ = e2NodeEntity.GetAspect(mastership)
	if _, ok := e2NodeRelations[mastership.NodeId]; !ok {
		log.Debugf("Updating MastershipState for E2Node '%s'", e2Node.GetID())
		if len(e2NodeRelations) == 0 {
			log.Warnf("No controls relations found for E2Node entity '%s'", e2Node.GetID())
			return nil
		}

		// Select a random master to assign to the E2 node
		relations := make([]topoapi.Object, 0, len(e2NodeRelations))
		for _, e2nodeRelation := range e2NodeRelations {
			relations = append(relations, e2nodeRelation)
		}
		relation := relations[rand.Intn(len(relations))]

		// Increment the mastership term and assign the selected master
		mastership.Term++
		mastership.NodeId = string(relation.ID)
		err = e2NodeEntity.SetAspect(mastership)
		if err != nil {
			log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", e2Node.GetID(), err)
			return err
		}

		// Update the E2 node entity
		err = r.rnib.Update(ctx, e2NodeEntity)
		if err != nil {
			if !errors.IsNotFound(err) {
				log.Warnf("Updating MastershipState for E2Node '%s' failed: %v", e2Node.GetID(), err)
				return err
			}
			return nil
		}
		return nil
	}
	return nil
}
