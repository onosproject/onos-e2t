// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	"testing"
	"time"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"
)

const (
	TopoServiceHost = "onos-topo"
	TopoServicePort = 5150
)

func GetCellIDsPerNode(nodeID topoapi.ID) ([]*topoapi.E2Cell, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := getTopoClient()
	if err != nil {
		return nil, err
	}
	objects, err := GetContainRelationObjects()
	if err != nil {
		return nil, err
	}
	var cells []*topoapi.E2Cell

	for _, obj := range objects {
		relation := obj.Obj.(*topoapi.Object_Relation)
		if relation.Relation.SrcEntityID == nodeID {
			targetEntity := relation.Relation.TgtEntityID

			object, err := client.Get(ctx, targetEntity)
			if err != nil {
				return nil, err
			}
			if object != nil && object.GetEntity().GetKindID() == topoapi.ID(topoapi.RANEntityKinds_E2CELL.String()) {
				cellObject := &topoapi.E2Cell{}
				object.GetAspect(cellObject)
				cells = append(cells, cellObject)
			}
		}
	}

	return cells, nil
}

func getTopoClient() (toposdk.Client, error) {
	client, err := toposdk.NewClient(toposdk.Config{
		TopoService: toposdk.ServiceConfig{
			Host: TopoServiceHost,
			Port: TopoServicePort,
		},
	})

	return client, err
}

func GetContainRelationObjects() ([]topoapi.Object, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := getTopoClient()
	if err != nil {
		return nil, err
	}

	filters := &topoapi.Filters{
		KindFilters: []*topoapi.Filter{
			{
				Filter: &topoapi.Filter_Equal_{
					Equal_: &topoapi.EqualFilter{
						Value: topoapi.RANRelationKinds_CONTAINS.String(),
					},
				},
			},
		},
	}

	response, err := client.List(ctx, filters)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func GetControlRelationObjects() ([]topoapi.Object, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := getTopoClient()
	if err != nil {
		return nil, err
	}
	filters := &topoapi.Filters{
		KindFilters: []*topoapi.Filter{
			{
				Filter: &topoapi.Filter_Equal_{
					Equal_: &topoapi.EqualFilter{
						Value: topoapi.RANRelationKinds_CONTROLS.String(),
					},
				},
			},
		},
	}

	response, err := client.List(ctx, filters)
	if err != nil {
		return nil, err
	}

	return response, nil

}

func GetNodeIDs(t *testing.T) ([]topoapi.ID, error) {
	objects, err := GetControlRelationObjects()
	if err != nil {
		return nil, err
	}
	var connectedNodes []topoapi.ID
	for _, obj := range objects {
		relation := obj.Obj.(*topoapi.Object_Relation)
		connectedNodes = append(connectedNodes, relation.Relation.TgtEntityID)

	}
	return connectedNodes, nil
}

func GetAllE2Connections(t *testing.T) ([]topoapi.ID, error) {
	objects, err := GetControlRelationObjects()
	if err != nil {
		return nil, err
	}
	var connectionIDs []topoapi.ID
	for _, obj := range objects {
		connectionIDs = append(connectionIDs, obj.ID)

	}
	return connectionIDs, nil
}
