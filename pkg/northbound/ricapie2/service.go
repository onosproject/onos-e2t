// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/ricapi/e2/headers/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"io"

	ricapie2v1beta1 "github.com/onosproject/onos-e2t/api/ricapi/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "e2")

// NewService creates a new E2T service
func NewService(streams *stream.Manager) northbound.Service {
	return &Service{streams: streams}
}

// Service is a Service implementation for E2T service.
type Service struct {
	northbound.Service
	streams *stream.Manager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{streams: s.streams}
	ricapie2v1beta1.RegisterE2TServiceServer(r, server)

}

// Server implements the gRPC service for E2 ricapi related functions.
type Server struct {
	streams *stream.Manager
}

func (s *Server) Stream(server ricapie2v1beta1.E2TService_StreamServer) error {
	// Get the application name
	request, err := server.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	log.Debugf("Received StreamRequest %+v", request)
	streamID := stream.ID(fmt.Sprintf("%s:%s", request.AppID, request.InstanceID))
	streamCh := make(chan stream.Message)
	stream, err := s.streams.Open(server.Context(), streamID, streamCh)
	if err != nil {
		log.Warnf("StreamRequest %+v failed: %v", request, err)
		return err
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Warnf("StreamRequest %+v failed: %v", request, err)
			return err
		}

		response := &ricapie2v1beta1.StreamResponse{
			Header: &v1beta1.ResponseHeader{
				EncodingType: v1beta1.EncodingType_ENCODING_TYPE_PROTO,
			},
			Payload: message.Payload,
		}

		err = server.Send(response)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Warnf("StreamResponse %+v failed: %v", response, err)
			return err
		}
	}
}
