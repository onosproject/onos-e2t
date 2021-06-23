// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"golang.org/x/net/context"
)

var log = logging.GetLogger("topo", "manager")

type Rnib struct {
	store rnib.Store
}

func (r *Rnib) DeleteE2Relation(ctx context.Context, relationID topoapi.ID) error {
	return r.store.Delete(ctx, relationID)
}

func (r *Rnib) GetE2Relation(ctx context.Context, deviceID topoapi.ID) (topoapi.ID, error) {
	objects, err := r.store.List(ctx, &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.RANRelationKinds_CONTROLS.String(),
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	podID := getPodID()
	for _, object := range objects {
		val := object.Obj.(*topoapi.Object_Relation)
		srcEntity := val.Relation.GetSrcEntityID()
		dstEntity := val.Relation.GetTgtEntityID()
		if srcEntity == topoapi.ID(podID) && dstEntity == deviceID {
			return object.ID, nil
		}
	}

	return "", errors.New(errors.NotFound, "E2 relation ID is not found")
}

func (r *Rnib) CreateOrUpdateE2Relation(ctx context.Context, deviceID topoapi.ID, relationID topoapi.ID) error {
	currentE2NodeObject, err := r.store.Get(ctx, deviceID)
	if err != nil {
		return err
	}

	currentRelationObject, err := r.store.Get(ctx, relationID)
	if currentE2NodeObject != nil && currentRelationObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		e2Relation := &topoapi.Object{
			ID:   relationID,
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.ID(topoapi.RANRelationKinds_CONTROLS.String()),
					SrcEntityID: topoapi.ID(getPodID()),
					TgtEntityID: deviceID,
				},
			},
		}
		err = r.store.Create(ctx, e2Relation)
		if err != nil {
			return err
		}
	} else if err == nil {
		currentRelationObject.Obj.(*topoapi.Object_Relation).Relation.SrcEntityID = topoapi.ID(getPodID())
		currentRelationObject.Obj.(*topoapi.Object_Relation).Relation.TgtEntityID = deviceID
		currentRelationObject.Obj.(*topoapi.Object_Relation).Relation.KindID = topoapi.ID(topoapi.RANRelationKinds_CONTROLS.String())
		err = r.store.Update(ctx, currentRelationObject)
		if err != nil {
			return err
		}
	}

	return err
}

func (r *Rnib) CreateOrUpdateE2CellRelation(ctx context.Context, deviceID topoapi.ID, cellID topoapi.ID) error {
	cellRelationID, err := getE2CellRelationID(deviceID, cellID)
	if err != nil {
		return err
	}
	currentCellRelation, err := r.store.Get(ctx, cellRelationID)
	if currentCellRelation == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		cellRelation := &topoapi.Object{
			ID:   cellRelationID,
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.ID(topoapi.RANRelationKinds_CONTAINS.String()),
					SrcEntityID: deviceID,
					TgtEntityID: cellID,
				},
			},
		}
		err = r.store.Create(ctx, cellRelation)
		if err != nil {
			return err
		}
	} else if err == nil {
		err := r.store.Update(ctx, currentCellRelation)
		if err != nil {
			return err
		}
	}
	return err
}

// CreateOrUpdateE2Cells creates or update E2 cells entities and relations
func (r *Rnib) CreateOrUpdateE2Cells(ctx context.Context, deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error {
	currentE2NodeObject, err := r.store.Get(ctx, deviceID)
	if currentE2NodeObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		return err
	}
	for _, e2Cell := range e2Cells {
		cellID := topoapi.ID(e2Cell.CellGlobalID.Value)
		currentCellObject, err := r.store.Get(ctx, cellID)
		if errors.IsNotFound(errors.FromGRPC(err)) {
			cellObject := &topoapi.Object{
				ID:   cellID,
				Type: topoapi.Object_ENTITY,
				Obj: &topoapi.Object_Entity{
					Entity: &topoapi.Entity{
						KindID: topoapi.ID(topoapi.RANEntityKinds_E2CELL.String()),
					},
				},
				Aspects: make(map[string]*gogotypes.Any),
				Labels:  map[string]string{},
			}

			err := cellObject.SetAspect(e2Cell)
			if err != nil {
				log.Warn(err)
				return err
			}
			err = r.store.Create(ctx, cellObject)
			if err != nil {
				return err
			}

			err = r.CreateOrUpdateE2CellRelation(ctx, deviceID, cellID)
			if err != nil {
				return err
			}

		} else if err == nil {
			err := currentCellObject.SetAspect(e2Cell)
			if err != nil {
				return err
			}

			err = r.store.Update(ctx, currentCellObject)
			if err != nil {
				return err
			}

			err = r.CreateOrUpdateE2CellRelation(ctx, deviceID, cellID)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// CreateOrUpdateE2Node creates or updates E2 entities
func (r *Rnib) CreateOrUpdateE2Node(ctx context.Context, deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error {
	e2NodeAspects := &topoapi.E2Node{
		ServiceModels: serviceModels,
	}
	currentE2NodeObject, err := r.store.Get(ctx, deviceID)
	if errors.IsNotFound(errors.FromGRPC(err)) {
		e2NodeObject := &topoapi.Object{
			ID:   deviceID,
			Type: topoapi.Object_ENTITY,
			Obj: &topoapi.Object_Entity{
				Entity: &topoapi.Entity{
					KindID: topoapi.ID(topoapi.RANEntityKinds_E2NODE.String()),
				},
			},
			Aspects: make(map[string]*gogotypes.Any),
			Labels:  map[string]string{},
		}

		err = e2NodeObject.SetAspect(e2NodeAspects)
		if err != nil {
			return err
		}
		err = r.store.Create(ctx, e2NodeObject)
		return err
	} else if err == nil {
		err := currentE2NodeObject.SetAspect(e2NodeAspects)
		if err != nil {
			return err
		}

		err = r.store.Update(ctx, currentE2NodeObject)
		if err != nil {
			return err
		}
	}

	return err
}

// Manager topology manager
type Manager interface {
	CreateOrUpdateE2Cells(ctx context.Context, deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error
	CreateOrUpdateE2CellRelation(ctx context.Context, deviceID topoapi.ID, cellID topoapi.ID) error
	CreateOrUpdateE2Node(ctx context.Context, deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error
	CreateOrUpdateE2Relation(ctx context.Context, deviceID topoapi.ID, relationID topoapi.ID) error
	DeleteE2Relation(ctx context.Context, relationID topoapi.ID) error
	GetE2Relation(ctx context.Context, deviceID topoapi.ID) (topoapi.ID, error)
}

// NewManager creates topology manager
func NewManager(store rnib.Store) *Rnib {
	return &Rnib{
		store: store,
	}
}

var _ Manager = &Rnib{}
