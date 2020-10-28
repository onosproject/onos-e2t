// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"io"

	ricapie2v1beta1 "github.com/onosproject/onos-e2t/api/ricapi/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "e2")

// Service is a Service implementation for E2T service.
type Service struct {
	northbound.Service
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := Server{}
	ricapie2v1beta1.RegisterE2TServiceServer(r, server)

}

// Server implements the gRPC service for E2 ricapi related functions.
type Server struct {
}

// RegisterApp process and handle the incoming requests from xApps
func (s Server) RegisterApp(stream ricapie2v1beta1.E2TService_RegisterAppServer) error {
	ctx := stream.Context()
	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Error("End of file error", err)
			return nil
		}
		if err != nil {
			log.Error(err)
			continue
		}

		err = s.appRequestHandler(req)
		if err != nil {
			log.Error(err)
			continue
		}

	}

}
