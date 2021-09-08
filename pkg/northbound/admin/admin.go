// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package admin

import (
	"context"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	adminapi "github.com/onosproject/onos-api/go/onos/e2t/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "admin")

// NewService creates a new admin service
func NewService(channels e2server.ConnManager) northbound.Service {
	return &Service{
		channels: channels,
	}
}

// Service is a Service implementation for administration.
type Service struct {
	channels e2server.ConnManager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{
		channels: s.channels,
	}
	adminapi.RegisterE2TAdminServiceServer(r, server)
}

var _ northbound.Service = &Service{}

// Server implements the gRPC service for administrative facilities.
type Server struct {
	channels e2server.ConnManager
}

// UploadRegisterServiceModel uploads and adds the model plugin to the list of supported models
func (s *Server) UploadRegisterServiceModel(stream adminapi.E2TAdminService_UploadRegisterServiceModelServer) error {
	return errors.Status(errors.NewNotSupported("UploadRegisterServiceModel not implemented")).Err()
}

// ListRegisteredServiceModels returns a stream of registered service models.
func (s *Server) ListRegisteredServiceModels(req *adminapi.ListRegisteredServiceModelsRequest, stream adminapi.E2TAdminService_ListRegisteredServiceModelsServer) error {
	return errors.Status(errors.NewNotSupported("ListRegisteredServiceModels not implemented")).Err()
}

// ListE2NodeConnections returns a stream of existing SCTP connections.
func (s *Server) ListE2NodeConnections(req *adminapi.ListE2NodeConnectionsRequest, stream adminapi.E2TAdminService_ListE2NodeConnectionsServer) error {
	return errors.Status(errors.NewNotSupported("ListE2NodeConnections not supported: use onos-topo instead")).Err()
}

// DropE2NodeConnections drops the specified E2 node SCTP connections
func (s *Server) DropE2NodeConnections(ctx context.Context, req *adminapi.DropE2NodeConnectionsRequest) (*adminapi.DropE2NodeConnectionsResponse, error) {
	return nil, errors.Status(errors.NewNotSupported("DropE2NodeConnections not supported: use onos-topo instead")).Err()
}
