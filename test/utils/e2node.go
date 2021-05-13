// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"context"
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
func GetTopoConn(topoEndpoint string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(topoEndpoint, opts...)
}

func GetCellIDsPerNode(nodeID topoapi.ID) ([]*topoapi.E2Cell, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100*time.Millisecond)))

	conn, _ := GetTopoConn(OnosTopoAddress, opts...)
	client := topoapi.CreateTopoClient(conn)
	listResponse, err := client.List(ctx, &topoapi.ListRequest{})
	if err != nil {
		return nil, err
	}
	var cells []*topoapi.E2Cell

	for _, obj := range listResponse.Objects {
		if obj.Type == topoapi.Object_RELATION {
			switch relation := obj.Obj.(type) {
			case *topoapi.Object_Relation:
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

		}
	}

	return cells, nil
}

// GetNodeIDs get list of E2 node IDs using topology subsystem
func GetNodeIDs() ([]topoapi.ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100*time.Millisecond)))

	conn, _ := GetTopoConn(OnosTopoAddress, opts...)
	client := topoapi.CreateTopoClient(conn)
	listResponse, err := client.List(ctx, &topoapi.ListRequest{})
	if err != nil {
		return nil, err
	}

	var nodeIDs []topoapi.ID
	for _, obj := range listResponse.Objects {
		if obj.Type == topoapi.Object_ENTITY {
			switch entity := obj.Obj.(type) {
			case *topoapi.Object_Entity:
				if entity.Entity.GetKindID() == topoapi.ID(topoapi.RANEntityKinds_E2NODE.String()) {
					nodeIDs = append(nodeIDs, obj.ID)
				}

			}

		}
	}
	return nodeIDs, nil

}
