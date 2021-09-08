// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package admin

import (
	"context"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
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
	log.Error("implement me")

	return nil
}

// ListRegisteredServiceModels returns a stream of registered service models.
func (s *Server) ListRegisteredServiceModels(req *adminapi.ListRegisteredServiceModelsRequest, stream adminapi.E2TAdminService_ListRegisteredServiceModelsServer) error {
	log.Error("implement me")
	return nil
}

// ListE2NodeConnections returns a stream of existing SCTP connections.
func (s *Server) ListE2NodeConnections(req *adminapi.ListE2NodeConnectionsRequest, stream adminapi.E2TAdminService_ListE2NodeConnectionsServer) error {
	return errors.Status(errors.NewNotSupported("ListE2NodeConnections not supported: use onos-topo instead")).Err()
}

func connectionType(nodeType types.E2NodeType) adminapi.E2NodeConnectionType {
	switch nodeType {
	case types.E2NodeTypeGNB:
		return adminapi.E2NodeConnectionType_G_NB
	case types.E2NodeTypeENB:
		return adminapi.E2NodeConnectionType_E_NB
	case types.E2NodeTypeEnGNB:
		return adminapi.E2NodeConnectionType_ENG_MB
	case types.E2NodeTypeNgENB:
		return adminapi.E2NodeConnectionType_NGE_NB
	default:
		return adminapi.E2NodeConnectionType_G_NB
	}
}

// DropE2NodeConnections drops the specified E2 node SCTP connections
func (s *Server) DropE2NodeConnections(ctx context.Context, req *adminapi.DropE2NodeConnectionsRequest) (*adminapi.DropE2NodeConnectionsResponse, error) {
	log.Error("implement me")
	return &adminapi.DropE2NodeConnectionsResponse{}, nil
}
