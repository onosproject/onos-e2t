// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/device"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

type topoManager struct {
	deviceStore device.Store
}

// CreateOrUpdateE2Cells creates or update E2 cells entities and relations
func (d *topoManager) CreateOrUpdateE2Cells(deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error {
	currentDeviceObject, err := d.deviceStore.Get(deviceID)
	if currentDeviceObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
		return err
	}
	for _, e2Cell := range e2Cells {
		cellID := topoapi.ID(e2Cell.CID)
		currentCellObject, err := d.deviceStore.Get(cellID)
		if currentCellObject == nil && errors.IsNotFound(errors.FromGRPC(err)) {
			cellObject := &topoapi.Object{
				ID:   cellID,
				Type: topoapi.Object_ENTITY,
				Obj: &topoapi.Object_Entity{
					Entity: &topoapi.Entity{
						KindID: topoapi.ID(topoapi.RANEntityKinds_E2CELL),
					},
				},
				Aspects: make(map[string]*gogotypes.Any),
				Labels:  []string{topoapi.RANEntityKinds_E2CELL.String()},
			}
			e2CellEntity, err := gogotypes.MarshalAny(e2Cell)
			if err != nil {
				return err
			}
			err = cellObject.SetAspect(e2CellEntity)
			if err != nil {
				return err
			}
			err = d.deviceStore.Create(cellObject)
			if err != nil {
				return err
			}

			cellRelationID := deviceID + ":" + cellID
			cellRelation := &topoapi.Object{
				ID:   cellRelationID,
				Type: topoapi.Object_RELATION,
				Obj: &topoapi.Object_Relation{
					Relation: &topoapi.Relation{
						KindID:      topoapi.ID(topoapi.RANRelationKinds_CONTROLS.String()),
						SrcEntityID: deviceID,
						TgtEntityID: cellID,
					},
				},
			}
			err = d.deviceStore.Create(cellRelation)
			if err != nil {
				return err
			}

		} else if currentCellObject != nil && err == nil {
			e2CellEntity, err := gogotypes.MarshalAny(e2Cell)
			if err != nil {
				return err
			}
			currentCellObject.Aspects[topoapi.RANEntityKinds_E2CELL.String()] = e2CellEntity
			_, err = d.deviceStore.Update(currentCellObject)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateOrUpdateE2Device creates or updates E2 entities
func (d *topoManager) CreateOrUpdateE2Device(deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error {
	currentDeviceObject, err := d.deviceStore.Get(deviceID)
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
			Labels:  []string{topoapi.RANEntityKinds_E2NODE.String()},
		}
		e2NodeEntity, err := gogotypes.MarshalAny(&topoapi.E2Node{
			ServiceModels: serviceModels,
		})
		if err != nil {
			return err
		}
		err = deviceObject.SetAspect(e2NodeEntity)
		if err != nil {
			return err
		}
		err = d.deviceStore.Create(deviceObject)
		if err != nil {
			return err
		}
	} else if currentDeviceObject != nil && err == nil {
		e2NodeEntity, err := gogotypes.MarshalAny(&topoapi.E2Node{
			ServiceModels: serviceModels,
		})
		if err != nil {
			return err
		}
		currentDeviceObject.Aspects[topoapi.RANEntityKinds_E2NODE.String()] = e2NodeEntity
		_, err = d.deviceStore.Update(currentDeviceObject)
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
	CreateOrUpdateE2Cells(deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error
	CreateOrUpdateE2Device(deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error
}

// NewTopoManager creates topology manager
func NewTopoManager(deviceStore device.Store) *topoManager {
	return &topoManager{
		deviceStore: deviceStore,
	}
}

var _ Manager = &topoManager{}
