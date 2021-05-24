// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package admin

import (
	"context"

	"github.com/onosproject/onos-e2t/pkg/ranfunctions"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	adminapi "github.com/onosproject/onos-api/go/onos/e2t/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "admin")

// NewService creates a new admin service
func NewService(channels e2server.ChannelManager, ranFunctionRegistry ranfunctions.Registry) northbound.Service {
	return &Service{
		channels:            channels,
		ranFunctionRegistry: ranFunctionRegistry}
}

// Service is a Service implementation for administration.
type Service struct {
	channels            e2server.ChannelManager
	ranFunctionRegistry ranfunctions.Registry
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{
		channels:            s.channels,
		ranFunctionRegistry: s.ranFunctionRegistry}
	adminapi.RegisterE2TAdminServiceServer(r, server)
}

var _ northbound.Service = &Service{}

// Server implements the gRPC service for administrative facilities.
type Server struct {
	channels            e2server.ChannelManager
	ranFunctionRegistry ranfunctions.Registry
}

// UploadRegisterServiceModel uploads and adds the model plugin to the list of supported models
func (s *Server) UploadRegisterServiceModel(stream adminapi.E2TAdminService_UploadRegisterServiceModelServer) error {
	log.Error("implement me")

	return nil
}

// ListRegisteredServiceModels returns a stream of registered service models.
func (s *Server) ListRegisteredServiceModels(req *adminapi.ListRegisteredServiceModelsRequest, stream adminapi.E2TAdminService_ListRegisteredServiceModelsServer) error {
	log.Error("implement me")
	return nil
}

// ListE2NodeConnections returns a stream of existing SCTP connections.
// Deprecated
func (s *Server) ListE2NodeConnections(req *adminapi.ListE2NodeConnectionsRequest, stream adminapi.E2TAdminService_ListE2NodeConnectionsServer) error {
	return nil
}

// DropE2NodeConnections drops the specified E2 node SCTP connections
func (s *Server) DropE2NodeConnections(ctx context.Context, req *adminapi.DropE2NodeConnectionsRequest) (*adminapi.DropE2NodeConnectionsResponse, error) {
	log.Error("implement me")
	return &adminapi.DropE2NodeConnectionsResponse{}, nil
}
