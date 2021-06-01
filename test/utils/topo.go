// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/options"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"
)

func GetCellIDsPerNode(nodeID topoapi.ID) ([]*topoapi.E2Cell, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := GetTopoClient()
	if err != nil {
		return nil, err
	}
	objects, err := GetTopoObjects(GetContainsRelationFilter())
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

func GetTopoClient() (toposdk.Client, error) {
	client, err := toposdk.NewClient()
	return client, err
}

func GetTopoObjects(filters *topoapi.Filters) ([]topoapi.Object, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := GetTopoClient()
	if err != nil {
		return nil, err
	}

	response, err := client.List(ctx, options.WithListFilters(filters))
	if err != nil {
		return nil, err
	}
	return response, nil

}

func GetTestNodeID() (topoapi.ID, error) {
	topoClient, err := GetTopoClient()
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	topoCh := make(chan topoapi.Event)

	err = topoClient.Watch(ctx, topoCh, options.WithWatchFilters(GetControlRelationFilter()))
	if err != nil {
		return "", err
	}

	topoEvent := <-topoCh
	relation := topoEvent.Object.Obj.(*topoapi.Object_Relation)
	testNodeID := relation.Relation.TgtEntityID
	return testNodeID, nil
}

func GetNodeIDs(t *testing.T) ([]topoapi.ID, error) {
	objects, err := GetTopoObjects(GetControlRelationFilter())
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

func GetE2NodeEntityFilter() *topoapi.Filters {
	e2nodeEntityFilter := &topoapi.Filters{
		KindFilters: []*topoapi.Filter{
			{
				Filter: &topoapi.Filter_Equal_{
					Equal_: &topoapi.EqualFilter{
						Value: topoapi.RANEntityKinds_E2NODE.String(),
					},
				},
			},
		},
	}
	return e2nodeEntityFilter
}

func GetControlRelationFilter() *topoapi.Filters {
	controlRelationFilter := &topoapi.Filters{
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
	return controlRelationFilter
}

func GetContainsRelationFilter() *topoapi.Filters {
	containsRelationFilter := &topoapi.Filters{
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

	return containsRelationFilter

}
