// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package xapp

import (
	xappv1 "github.com/onosproject/onos-e2t/api/xapp/v1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "xapp")

// Service is a Service implementation for E2T service.
type Service struct {
	northbound.Service
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := Server{}
	xappv1.RegisterE2TServiceServer(r, server)
}

// Server implements the gRPC service for E2AP related functions.
type Server struct {
}

func (s Server) Subscribe(req *xappv1.SubscribeRequest, stream xappv1.E2TService_SubscribeServer) error {
	panic("implement me")
}

// RegisterApp ...
func (s Server) RegisterApp(xappv1.E2TService_RegisterAppServer) error {
	log.Error("Implement me")
	return nil
}
