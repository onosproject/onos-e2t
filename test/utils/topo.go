// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	"testing"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/certs"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/southbound"

	"google.golang.org/grpc"
)

const (
	OnosTopoAddress = "onos-topo:5150"
)

// GetTopoConn gets a gRPC connection to the topology service
func GetTopoConn(topoEndpoint string) (*grpc.ClientConn, error) {
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100*time.Millisecond)))
	return grpc.Dial(topoEndpoint, opts...)
}

func GetCellIDsPerNode(nodeID topoapi.ID) ([]*topoapi.E2Cell, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := GetTopoConn(OnosTopoAddress)
	if err != nil {
		return nil, err
	}
	client := topoapi.CreateTopoClient(conn)
	objects, err := GetContainRelationObjects()
	if err != nil {
		return nil, err
	}
	var cells []*topoapi.E2Cell

	for _, obj := range objects {
		relation := obj.Obj.(*topoapi.Object_Relation)
		if relation.Relation.SrcEntityID == nodeID {
			targetEntity := relation.Relation.TgtEntityID
			getRequest := &topoapi.GetRequest{
				ID: targetEntity,
			}
			response, err := client.Get(ctx, getRequest)
			if err != nil {
				return nil, err
			}
			object := response.Object
			if object != nil && object.GetEntity().GetKindID() == topoapi.ID(topoapi.RANEntityKinds_E2CELL.String()) {
				cellObject := &topoapi.E2Cell{}
				object.GetAspect(cellObject)
				cells = append(cells, cellObject)
			}

		}
	}

	return cells, nil
}

func GetContainRelationObjects() ([]topoapi.Object, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := GetTopoConn(OnosTopoAddress)
	if err != nil {
		return nil, err
	}
	client := topoapi.CreateTopoClient(conn)
	listResponse, err := client.List(ctx, &topoapi.ListRequest{
		Filters: &topoapi.Filters{
			KindFilters: []*topoapi.Filter{
				{
					Filter: &topoapi.Filter_Equal_{
						Equal_: &topoapi.EqualFilter{
							Value: topoapi.RANRelationKinds_CONTAINS.String(),
						},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return listResponse.Objects, nil

}

func GetControlRelationObjects() ([]topoapi.Object, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn, err := GetTopoConn(OnosTopoAddress)
	if err != nil {
		return nil, err
	}
	client := topoapi.CreateTopoClient(conn)
	listResponse, err := client.List(ctx, &topoapi.ListRequest{
		Filters: &topoapi.Filters{
			KindFilters: []*topoapi.Filter{
				{
					Filter: &topoapi.Filter_Equal_{
						Equal_: &topoapi.EqualFilter{
							Value: topoapi.RANRelationKinds_CONTROLS.String(),
						},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return listResponse.Objects, nil

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
