// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"google.golang.org/grpc/codes"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	modelapi "github.com/onosproject/onos-api/go/onos/ransim/model"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ConnectRansimServiceHost connects to ransim service
func (s *TestSuite) ConnectRansimServiceHost() (*grpc.ClientConn, error) {
	services, err := s.CoreV1().Services(s.Namespace()).List(s.Context(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	tlsConfig, err := creds.GetClientCredentials()
	if err != nil {
		return nil, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable, codes.Unknown))),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable, codes.Unknown))),
	}

	ransimServiceAddress := utils.GetRansimServiceAddress(services.Items[0].Name)
	return grpc.DialContext(s.Context(), ransimServiceAddress, opts...)
}

func (s *TestSuite) GetNodes(nodeClient modelapi.NodeModelClient) []*ransimtypes.Node {
	ctx, cancel := context.WithCancel(s.Context())
	defer cancel()
	stream, err := nodeClient.ListNodes(ctx, &modelapi.ListNodesRequest{})
	s.NoError(err)
	var nodes []*ransimtypes.Node
	for {
		e2node, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return []*ransimtypes.Node{}
		}
		nodes = append(nodes, e2node.Node)
	}
	return nodes
}

func (s *TestSuite) GetCells(cellClient modelapi.CellModelClient) []*ransimtypes.Cell {
	ctx, cancel := context.WithCancel(s.Context())
	defer cancel()
	stream, err := cellClient.ListCells(ctx, &modelapi.ListCellsRequest{})
	s.NoError(err)
	var cellsList []*ransimtypes.Cell
	for {
		cell, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return []*ransimtypes.Cell{}
		}

		cellsList = append(cellsList, cell.Cell)
	}
	return cellsList
}

func (s *TestSuite) GetNumCells(cellClient modelapi.CellModelClient) int {
	stream, err := cellClient.ListCells(s.Context(), &modelapi.ListCellsRequest{})
	s.NoError(err)
	numCells := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0
		}
		numCells++
	}
	return numCells
}

func (s *TestSuite) GetNumNodes(nodeClient modelapi.NodeModelClient) int {
	stream, err := nodeClient.ListNodes(s.Context(), &modelapi.ListNodesRequest{})
	s.NoError(err)
	numNodes := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0
		}
		numNodes++
	}
	return numNodes
}

func (s *TestSuite) GetRansimCellClient() modelapi.CellModelClient {
	conn, err := s.ConnectRansimServiceHost()
	s.NoError(err)
	s.NotNil(conn)
	return modelapi.NewCellModelClient(conn)
}

func (s *TestSuite) GetRansimNodeClient() modelapi.NodeModelClient {
	conn, err := s.ConnectRansimServiceHost()
	s.NoError(err)
	s.NotNil(conn)
	return modelapi.NewNodeModelClient(conn)
}
