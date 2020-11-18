// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	headerapi "github.com/onosproject/onos-e2t/api/ricapi/e2/headers/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/northbound/codec"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"io"

	ricapi "github.com/onosproject/onos-e2t/api/ricapi/e2/v1beta1"
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
	ricapi.RegisterE2TServiceServer(r, server)

}

// Server implements the gRPC service for E2 ricapi related functions.
type Server struct {
	streams *stream.Manager
}

func (s *Server) Stream(server ricapi.E2TService_StreamServer) error {
	// Get the application name
	request, err := server.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	log.Infof("Received StreamRequest %+v", request)
	streamCh := make(chan stream.Message)
	streamMeta := stream.Metadata{
		AppID:          request.AppID,
		InstanceID:     request.InstanceID,
		SubscriptionID: request.SubscriptionID,
	}
	stream, err := s.streams.Open(server.Context(), streamMeta, streamCh)
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

		bytes, err := codec.Encode(message.Payload.(*e2appdudescriptions.E2ApPdu), request.Encoding)
		if err != nil {
			return errors.NewInvalid(err.Error())
		}

		response := &ricapi.StreamResponse{
			Header: &headerapi.ResponseHeader{
				EncodingType: request.Encoding,
			},
			Payload: bytes,
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
