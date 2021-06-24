// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

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
			if object != nil && object.GetEntity().GetKindID() == topoapi.E2CELL {
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
			KindFilter: &topoapi.Filter{
				Filter: &topoapi.Filter_Equal_{
					Equal_: &topoapi.EqualFilter{
						Value: topoapi.CONTAINS,
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
			KindFilter: &topoapi.Filter{
				Filter: &topoapi.Filter_Equal_{
					Equal_: &topoapi.EqualFilter{
						Value: topoapi.CONTROLS,
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
	const maxAttempts = 30
	var err error
	var connectedNodes []topoapi.ID
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		objects, err := GetControlRelationObjects()
		if err != nil || len(objects) == 0 {
			fmt.Fprintf(os.Stderr, "Attempt %d got an error, sleeping\n", attempt)
			time.Sleep(2 * time.Second)
			continue
		} else {
			for _, obj := range objects {
				relation := obj.Obj.(*topoapi.Object_Relation)
				connectedNodes = append(connectedNodes, relation.Relation.TgtEntityID)
			}
			break
		}
	}
	return connectedNodes, err
}


func GetFirstNodeID(t *testing.T) topoapi.ID {
	const maxAttempts = 15
	var nodeIDs []topoapi.ID
	var err error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		nodeIDs, err = GetNodeIDs(t)
		if err != nil || len(nodeIDs) == 0 {
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	assert.GreaterOrEqual(t, len(nodeIDs), 1, "No nodes found")
	return nodeIDs[0]
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
