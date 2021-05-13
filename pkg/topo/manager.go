// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"os"

	"github.com/google/uuid"

	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/device"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("topo", "manager")

type DeviceManager struct {
	deviceStore device.Store
}

func (d *DeviceManager) DeleteE2Relation(relationID topoapi.ID) error {
	return d.deviceStore.Delete(relationID)
}

func (d *DeviceManager) CreateOrUpdateE2Relation(deviceID topoapi.ID) error {
	podID := os.Getenv("POD_ID")
	currentDeviceObject, err := d.deviceStore.Get(deviceID)
	if err != nil {
		return err
	}

	if currentDeviceObject != nil {
		e2Relation := &topoapi.Object{
			ID:   topoapi.ID(uuid.New().String()),
			Type: topoapi.Object_RELATION,
			Obj: &topoapi.Object_Relation{
				Relation: &topoapi.Relation{
					KindID:      topoapi.ID(topoapi.RANRelationKinds_CONTROLS.String()),
					SrcEntityID: topoapi.ID(podID),
					TgtEntityID: deviceID,
				},
			},
		}
		err = d.deviceStore.Create(e2Relation)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateOrUpdateE2Cells creates or update E2 cells entities and relations
func (d *DeviceManager) CreateOrUpdateE2Cells(deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error {
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
						KindID: topoapi.ID(topoapi.RANEntityKinds_E2CELL.String()),
					},
				},
				Aspects: make(map[string]*gogotypes.Any),
				Labels:  []string{topoapi.RANEntityKinds_E2CELL.String()},
			}

			err := cellObject.SetAspect(e2Cell)
			if err != nil {
				log.Warn(err)
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

			err := currentCellObject.SetAspect(e2Cell)
			if err != nil {
				return err
			}

			err = d.deviceStore.Update(currentCellObject)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateOrUpdateE2Device creates or updates E2 entities
func (d *DeviceManager) CreateOrUpdateE2Device(deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error {
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
		e2Node := &topoapi.E2Node{
			ServiceModels: serviceModels,
		}

		err = deviceObject.SetAspect(e2Node)
		if err != nil {
			return err
		}
		err = d.deviceStore.Create(deviceObject)
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

		err = d.deviceStore.Update(currentDeviceObject)
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
	CreateOrUpdateE2Relation(deviceID topoapi.ID) error
	DeleteE2Relation(relationID topoapi.ID) error
}

// NewDeviceManager creates topology manager
func NewDeviceManager(deviceStore device.Store) *DeviceManager {
	return &DeviceManager{
		deviceStore: deviceStore,
	}
}

var _ Manager = &DeviceManager{}
