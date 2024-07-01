// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"context"
	"io"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var log = logging.GetLogger()

const e2NodeIDHeader = "e2-node-id"

// NewProxyService creates a new E2T control and subscription proxy service
func NewProxyService(clientConn *grpc.ClientConn) northbound.Service {
	return &SubscriptionService{
		conn: clientConn,
	}
}

// SubscriptionService is a Service implementation for E2 Subscription service.
type SubscriptionService struct {
	northbound.Service
	conn *grpc.ClientConn
}

// Register registers the SubscriptionService with the gRPC server.
func (s SubscriptionService) Register(r *grpc.Server) {
	server := &ProxyServer{
		conn: s.conn,
	}
	e2api.RegisterSubscriptionServiceServer(r, server)
	e2api.RegisterControlServiceServer(r, server)
}

// ProxyServer implements the gRPC service for E2 Subscription related functions.
type ProxyServer struct {
	conn *grpc.ClientConn
}

func (s *ProxyServer) Control(ctx context.Context, request *e2api.ControlRequest) (*e2api.ControlResponse, error) {
	log.Debugf("ControlRequest %+v", request)
	client := e2api.NewControlServiceClient(s.conn)
	ctx = metadata.AppendToOutgoingContext(ctx, e2NodeIDHeader, string(request.Headers.E2NodeID))
	response, err := client.Control(ctx, request)
	if err != nil {
		log.Warnf("ControlRequest %+v error: %s", request, err)
		return nil, err
	}
	log.Debugf("ControlResponse %+v", response)
	return response, nil
}

func (s *ProxyServer) Subscribe(request *e2api.SubscribeRequest, server e2api.SubscriptionService_SubscribeServer) error {
	log.Debugf("SubscribeRequest %+v", request)
	client := e2api.NewSubscriptionServiceClient(s.conn)
	ctx := metadata.AppendToOutgoingContext(server.Context(), e2NodeIDHeader, string(request.Headers.E2NodeID))
	clientStream, err := client.Subscribe(ctx, request)
	if err != nil {
		log.Warnf("SubscribeRequest %+v error: %s", request, err)
		return err
	}

	for {
		response, err := clientStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Warnf("SubscribeRequest %+v error: %s", request, err)
			return err
		}
		log.Debugf("SubscribeResponse %+v", response)
		err = server.Send(response)
		if err != nil {
			log.Warnf("SubscribeResponse %+v error: %s", response, err)
			return err
		}
	}
}

func (s *ProxyServer) Unsubscribe(ctx context.Context, request *e2api.UnsubscribeRequest) (*e2api.UnsubscribeResponse, error) {
	log.Debugf("UnsubscribeRequest %+v", request)
	client := e2api.NewSubscriptionServiceClient(s.conn)
	ctx = metadata.AppendToOutgoingContext(ctx, e2NodeIDHeader, string(request.Headers.E2NodeID))
	response, err := client.Unsubscribe(ctx, request)
	if err != nil {
		log.Warnf("UnsubscribeRequest %+v error: %s", request, err)
		return nil, err
	}
	log.Debugf("UnsubscribeResponse %+v", response)
	return response, nil
}
