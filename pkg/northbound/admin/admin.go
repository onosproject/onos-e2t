// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package admin

import (
	"context"
	"errors"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	sctpnet "github.com/ishidawataru/sctp"
	adminapi "github.com/onosproject/onos-api/go/onos/e2t/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "admin")

// NewService creates a new admin service
func NewService(channels e2server.ChannelManager) northbound.Service {
	return &Service{channels}
}

// Service is a Service implementation for administration.
type Service struct {
	channels e2server.ChannelManager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{s.channels}
	adminapi.RegisterE2TAdminServiceServer(r, server)
}

var _ northbound.Service = &Service{}

// Server implements the gRPC service for administrative facilities.
type Server struct {
	channels e2server.ChannelManager
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
	channels, err := s.channels.List(stream.Context())
	if err != nil {
		return err
	}

	for _, channel := range channels {
		sctpAddr := channel.RemoteAddr().(*sctpnet.SCTPAddr)
		if sctpAddr == nil {
			log.Errorf("Found non-SCTP connection in CreateConnection: %v", channel)
			return errors.New("found non-SCTP connection")
		}
		remoteAddrs := channel.RemoteAddr().(*sctpnet.SCTPAddr).IPAddrs
		remotePort := uint32(channel.RemoteAddr().(*sctpnet.SCTPAddr).Port)
		var remoteAddrsStrings []string
		for _, remoteAddr := range remoteAddrs {
			remoteAddrsStrings = append(remoteAddrsStrings, remoteAddr.String())
		}
		msg := &adminapi.ListE2NodeConnectionsResponse{
			RemoteIp:   remoteAddrsStrings,
			RemotePort: remotePort,
			Id:         string(channel.ID),
			PlmnId:     channel.PlmnID,
			// TODO: This should come from the connection data
			ConnectionType: adminapi.E2NodeConnectionType_G_NB,
		}

		err = stream.Send(msg)
		if err != nil {
			return err
		}
	}
	return err
}

// DropE2NodeConnections drops the specified E2 node SCTP connections
func (s *Server) DropE2NodeConnections(ctx context.Context, req *adminapi.DropE2NodeConnectionsRequest) (*adminapi.DropE2NodeConnectionsResponse, error) {
	log.Error("implement me")
	return &adminapi.DropE2NodeConnectionsResponse{}, nil
}
