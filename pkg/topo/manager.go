// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/object"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"golang.org/x/net/context"
)

var log = logging.GetLogger("device", "manager")

type DeviceManager struct {
	deviceStore object.Store
}

func (d *DeviceManager) DeleteE2Relation(ctx context.Context, relationID topoapi.ID) error {
	return d.deviceStore.Delete(ctx, relationID)
}

func (d *DeviceManager) GetE2Relation(ctx context.Context, deviceID topoapi.ID) (topoapi.ID, error) {
	objects, err := d.deviceStore.List(ctx)
	if err != nil {
		return "", err
	}
	podID := getPodID()
	for _, object := range objects {
		if object.Type == topoapi.Object_RELATION {
			switch val := object.Obj.(type) {
			case *topoapi.Object_Relation:
				srcEntity := val.Relation.GetSrcEntityID()
				dstEntity := val.Relation.GetTgtEntityID()
				if srcEntity == topoapi.ID(podID) && dstEntity == deviceID {
					return object.ID, nil
				}

			}

		}
	}
	return "", errors.New(errors.NotFound, "E2 relation ID is not found")
}

func (d *DeviceManager) CreateOrUpdateE2Relation(ctx context.Context, deviceID topoapi.ID, relationID topoapi.ID) error {
	podID := getPodID()
	currentDeviceObject, err := d.deviceStore.Get(ctx, deviceID)
	if err != nil {
		return err
	}

	currentRelationObject, err := d.deviceStore.Get(ctx, relationID)
	if currentDeviceObject != nil && currentRelationObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		e2Relation := &topoapi.Object{
			ID:   relationID,
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.ID(topoapi.RANRelationKinds_CONTROLS.String()),
					SrcEntityID: topoapi.ID(podID),
					TgtEntityID: deviceID,
				},
			},
		}
		err = d.deviceStore.Create(ctx, e2Relation)
		if err != nil {
			return err
		}
	} else if currentDeviceObject != nil && currentRelationObject != nil && err == nil {
		err = d.deviceStore.Update(ctx, currentRelationObject)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func (d *DeviceManager) createOrUpdateE2CellRelation(ctx context.Context, deviceID topoapi.ID, cellID topoapi.ID) error {
	cellRelationID, err := getE2CellRelationID(deviceID, cellID)
	if err != nil {
		return err
	}
	currentCellRelation, err := d.deviceStore.Get(ctx, cellRelationID)
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
		err = d.deviceStore.Create(ctx, cellRelation)
		if err != nil {
			return err
		}
	} else if currentCellRelation != nil && err == nil {
		err := d.deviceStore.Update(ctx, currentCellRelation)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

// CreateOrUpdateE2Cells creates or update E2 cells entities and relations
func (d *DeviceManager) CreateOrUpdateE2Cells(ctx context.Context, deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error {
	currentDeviceObject, err := d.deviceStore.Get(ctx, deviceID)
	if currentDeviceObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		return err
	}
	for _, e2Cell := range e2Cells {
		cellID := topoapi.ID(e2Cell.CID)
		currentCellObject, err := d.deviceStore.Get(ctx, cellID)
		if currentCellObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
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
			err = d.deviceStore.Create(ctx, cellObject)
			if err != nil {
				return err
			}

			err = d.createOrUpdateE2CellRelation(ctx, deviceID, cellID)
			if err != nil {
				return err
			}

		} else if currentCellObject != nil && err == nil {

			err := currentCellObject.SetAspect(e2Cell)
			if err != nil {
				return err
			}

			err = d.deviceStore.Update(ctx, currentCellObject)
			if err != nil {
				return err
			}

			err = d.createOrUpdateE2CellRelation(ctx, deviceID, cellID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateOrUpdateE2Device creates or updates E2 entities
func (d *DeviceManager) CreateOrUpdateE2Device(ctx context.Context, deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error {
	currentDeviceObject, err := d.deviceStore.Get(ctx, deviceID)
	if currentDeviceObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		deviceObject := &topoapi.Object{
			ID:   deviceID,
			Type: topoapi.Object_ENTITY,
			Obj: &topoapi.Object_Entity{
				Entity: &topoapi.Entity{
					KindID: topoapi.ID(topoapi.RANEntityKinds_E2NODE.String()),
					Protocols: []*topoapi.ProtocolState{
						{
							Protocol: topoapi.Protocol_E2AP,
						},
					},
				},
			},
			Aspects: make(map[string]*gogotypes.Any),
			Labels:  map[string]string{},
		}
		e2Node := &topoapi.E2Node{
			ServiceModels: serviceModels,
		}

		err = deviceObject.SetAspect(e2Node)
		if err != nil {
			return err
		}
		err = d.deviceStore.Create(ctx, deviceObject)
		if err != nil {
			return err
		}
	} else if currentDeviceObject != nil && err == nil {
		e2Node := &topoapi.E2Node{
			ServiceModels: serviceModels,
		}

		err := currentDeviceObject.SetAspect(e2Node)
		if err != nil {
			return err
		}

		err = d.deviceStore.Update(ctx, currentDeviceObject)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

// Manager topology manager
type Manager interface {
	CreateOrUpdateE2Cells(ctx context.Context, deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error
	CreateOrUpdateE2Device(ctx context.Context, deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error
	CreateOrUpdateE2Relation(ctx context.Context, deviceID topoapi.ID, relationID topoapi.ID) error
	DeleteE2Relation(ctx context.Context, relationID topoapi.ID) error
	GetE2Relation(ctx context.Context, deviceID topoapi.ID) (topoapi.ID, error)
}

// NewManager creates topology manager
func NewManager(deviceStore object.Store) *DeviceManager {
	return &DeviceManager{
		deviceStore: deviceStore,
	}
}

var _ Manager = &DeviceManager{}
