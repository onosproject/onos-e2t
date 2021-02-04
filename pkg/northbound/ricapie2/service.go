// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"io"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"

	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "e2")

// NewService creates a new E2T service
func NewService(streams *stream.Manager, modelRegistry *modelregistry.ModelRegistry) northbound.Service {
	return &Service{
		streams:       streams,
		modelRegistry: modelRegistry,
	}
}

// Service is a Service implementation for E2T service.
type Service struct {
	northbound.Service
	streams       *stream.Manager
	modelRegistry *modelregistry.ModelRegistry
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{streams: s.streams, modelRegistry: s.modelRegistry}
	e2api.RegisterE2TServiceServer(r, server)

}

// Server implements the gRPC service for E2 ricapi related functions.
type Server struct {
	streams       *stream.Manager
	modelRegistry *modelregistry.ModelRegistry
}

func (s *Server) Stream(server e2api.E2TService_StreamServer) error {
	// Get the application name
	request, err := server.Recv()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	encodingType := request.GetHeader().GetEncodingType()

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

		ricIndication, ok := message.Payload.(e2appducontents.Ricindication)
		if !ok {
			return errors.NewInvalid("payload cannot be converted to E2AP PDU", message.Payload)
		}
		ranFuncID := ricIndication.ProtocolIes.E2ApProtocolIes5.Value.Value
		ricActionID := ricIndication.ProtocolIes.E2ApProtocolIes15.Value.Value
		indHeaderAsn1 := ricIndication.ProtocolIes.E2ApProtocolIes25.Value.Value
		indMessageAsn1 := ricIndication.ProtocolIes.E2ApProtocolIes26.Value.Value
		log.Infof("Ric Indication. Ran FundID: %d, Ric Action ID: %d", ranFuncID, ricActionID)

		response := &e2api.StreamResponse{
			Header: &e2api.ResponseHeader{
				EncodingType: encodingType,
			},
		}

		const serviceModelID = "e2sm_kpm-v1beta1" // TODO: Remove hardcoded value
		switch encodingType {
		case e2api.EncodingType_PROTO:
			serviceModelPlugin, ok := s.modelRegistry.ModelPlugins[serviceModelID]
			if !ok {
				log.Errorf("Service Model Plugin cannot be loaded %s", serviceModelID)
				return errors.NewInvalid("Service Model Plugin cannot be loaded", serviceModelID)
			}
			a, b, c := serviceModelPlugin.ServiceModelData()
			log.Infof("Service model found %s %s %s", a, b, c)

			indHeaderProto, err := serviceModelPlugin.IndicationHeaderASN1toProto(indHeaderAsn1)
			if err != nil {
				log.Errorf("Error transforming Header ASN Bytes to Proto %s", err.Error())
				return errors.NewInvalid(err.Error())
			}
			log.Infof("Indication Header %d bytes", len(indHeaderProto))

			indMessageProto, err := serviceModelPlugin.IndicationMessageASN1toProto(indMessageAsn1)
			if err != nil {
				log.Errorf("Error transforming Message ASN Bytes to Proto %s", err.Error())
				return errors.NewInvalid(err.Error())
			}
			log.Infof("Indication Message %d bytes", len(indMessageProto))
			response.Header.IndicationHeader = indHeaderProto
			response.IndicationMessage = indMessageProto
			log.Infof("RICIndication successfully decoded from ASN1 to Proto #Bytes - Header: %d, Message: %d",
				len(indHeaderProto), len(indMessageProto))
		case e2api.EncodingType_ASN1_PER:
			response.Header.IndicationHeader = indHeaderAsn1
			response.IndicationMessage = indMessageAsn1
		default:
			log.Errorf("encoding type %v not supported", request.Header.EncodingType)
			return errors.NewInvalid("encoding type %v not supported", request.Header.EncodingType)
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
