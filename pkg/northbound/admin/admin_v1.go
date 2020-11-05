// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package admin

import (
	"context"
	"errors"
	sctpnet "github.com/ishidawataru/sctp"
	adminv1 "github.com/onosproject/onos-e2t/api/admin/v1"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/connection"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "admin")

// NewService creates a new admin service
func NewService(conns *connection.Manager) northbound.Service {
	return &Service{conns}
}

// Service is a Service implementation for administration.
type Service struct {
	conns *connection.Manager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{s.conns}
	adminv1.RegisterE2TAdminServiceServer(r, server)
}

var _ northbound.Service = &Service{}

// Server implements the gRPC service for administrative facilities.
type Server struct {
	conns *connection.Manager
}

// UploadRegisterServiceModel uploads and adds the model plugin to the list of supported models
func (s *Server) UploadRegisterServiceModel(stream adminv1.E2TAdminService_UploadRegisterServiceModelServer) error {
	log.Error("implement me")

	return nil
}

// ListRegisteredServiceModels returns a stream of registered service models.
func (s *Server) ListRegisteredServiceModels(req *adminv1.ListRegisteredServiceModelsRequest, stream adminv1.E2TAdminService_ListRegisteredServiceModelsServer) error {
	log.Error("implement me")
	return nil
}

// ListE2NodeConnections returns a stream of existing SCTP connections.
func (s *Server) ListE2NodeConnections(req *adminv1.ListE2NodeConnectionsRequest, stream adminv1.E2TAdminService_ListE2NodeConnectionsServer) error {
	conns := s.conns.List()
	var err error
	for _, conn := range conns {
		sctpAddr := conn.RemoteAddr().(*sctpnet.SCTPAddr)
		if sctpAddr == nil {
			log.Errorf("Found non-SCTP connection in CreateConnection: %v", conn)
			return errors.New("found non-SCTP connection")
		}
		remoteAddrs := conn.RemoteAddr().(*sctpnet.SCTPAddr).IPAddrs
		remotePort := uint32(conn.RemoteAddr().(*sctpnet.SCTPAddr).Port)
		var remoteAddrsStrings []string
		for _, remoteAddr := range remoteAddrs {
			remoteAddrsStrings = append(remoteAddrsStrings, remoteAddr.String())
		}
		msg := &adminv1.ListE2NodeConnectionsResponse{
			RemoteIp:   remoteAddrsStrings,
			RemotePort: remotePort,
			Id:         uint32(conn.ID),
			PlmnId:     string(conn.PlmnID),
			// TODO: This should come from the connection data
			ConnectionType: adminv1.E2NodeConnectionType_G_NB,
		}

		err = stream.Send(msg)
		if err != nil {
			return err
		}
	}
	return err
}

// DropE2NodeConnections drops the specified E2 node SCTP connections
func (s *Server) DropE2NodeConnections(ctx context.Context, req *adminv1.DropE2NodeConnectionsRequest) (*adminv1.DropE2NodeConnectionsResponse, error) {
	log.Error("implement me")
	return &adminv1.DropE2NodeConnectionsResponse{}, nil
}
