// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package admin

import (
	"context"
	"errors"
	"time"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2conn"

	adminapi "github.com/onosproject/onos-api/go/onos/e2t/admin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "admin")

// NewService creates a new admin service
func NewService(connections e2conn.ConnManager) northbound.Service {
	return &Service{
		connections: connections,
	}
}

// Service is a Service implementation for administration.
type Service struct {
	connections e2conn.ConnManager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{
		connections: s.connections,
	}
	adminapi.RegisterE2TAdminServiceServer(r, server)
}

var _ northbound.Service = &Service{}

// Server implements the gRPC service for administrative facilities.
type Server struct {
	connections e2conn.ConnManager
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
	connections, err := s.connections.List(stream.Context())
	if err != nil {
		return err
	}

	for _, conn := range connections {
		sctpAddr := conn.RemoteAddr()
		if sctpAddr == nil {
			log.Errorf("Found non-SCTP connection in CreateConnection: %v", conn)
			return errors.New("found non-SCTP connection")
		}
		remoteAddrs := conn.RemoteAddr().(*addressing.Address).IPAddrs
		remotePort := uint32(conn.RemoteAddr().(*addressing.Address).Port)
		var remoteAddrsStrings []string
		for _, remoteAddr := range remoteAddrs {
			remoteAddrsStrings = append(remoteAddrsStrings, remoteAddr.String())
		}

		msg := &adminapi.ListE2NodeConnectionsResponse{
			Id:         string(conn.GetID()),
			RemoteIp:   remoteAddrsStrings,
			RemotePort: remotePort,
			PlmnId:     conn.GetPlmnID(),
			NodeId:     string(conn.GetE2NodeID()),
			//ConnectionType: connectionType(conn.NodeType),
			AgeMs: int32(time.Since(conn.GetTimeAlive()).Milliseconds()),
		}

		err = stream.Send(msg)
		if err != nil {
			return err
		}
	}
	return err
}

/*func connectionType(nodeType types.E2NodeType) adminapi.E2NodeConnectionType {
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
}*/

// DropE2NodeConnections drops the specified E2 node SCTP connections
func (s *Server) DropE2NodeConnections(ctx context.Context, req *adminapi.DropE2NodeConnectionsRequest) (*adminapi.DropE2NodeConnectionsResponse, error) {
	log.Error("implement me")
	return &adminapi.DropE2NodeConnectionsResponse{}, nil
}
