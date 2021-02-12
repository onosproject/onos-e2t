// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"context"
	"io"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"

	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "e2")

// NewService creates a new E2T service
func NewService(subs subapi.E2SubscriptionServiceClient, streams *stream.Manager, modelRegistry *modelregistry.ModelRegistry, channels *e2server.ChannelManager) northbound.Service {
	return &Service{
		subs:          subs,
		streams:       streams,
		modelRegistry: modelRegistry,
		channels:      channels,
	}
}

// Service is a Service implementation for E2T service.
type Service struct {
	northbound.Service
	subs          subapi.E2SubscriptionServiceClient
	streams       *stream.Manager
	modelRegistry *modelregistry.ModelRegistry
	channels      *e2server.ChannelManager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{subs: s.subs, streams: s.streams, modelRegistry: s.modelRegistry,
		channels: s.channels, controlRequestID: RequestID(0)}
	e2api.RegisterE2TServiceServer(r, server)
}

// Server implements the gRPC service for E2 ricapi related functions.
type Server struct {
	subs             subapi.E2SubscriptionServiceClient
	streams          *stream.Manager
	modelRegistry    *modelregistry.ModelRegistry
	channels         *e2server.ChannelManager
	controlRequestID RequestID
}

func (s *Server) Control(ctx context.Context, request *e2api.ControlRequest) (*e2api.ControlResponse, error) {
	log.Info("Received Control Request %v", request)
	channel, err := s.channels.Get(ctx, e2server.ChannelID(request.E2NodeID))
	response := &e2api.ControlResponse{}
	if err != nil {
		return response, err
	}
	serviceModelID := modelregistry.ModelFullName(request.Header.ServiceModel.ID)
	_, ok := s.modelRegistry.ModelPlugins[serviceModelID]
	if !ok {
		response := &e2api.ControlResponse{}
		return response, err
	}
	s.controlRequestID++
	requestID := s.controlRequestID

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(requestID),
		InstanceID:  config.InstanceID,
	}

	ranFuncID := channel.GetRANFunctionID(serviceModelID)

	ricRequestID := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes29{
		Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(ricRequest.RequestorID), // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  int32(ricRequest.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionID := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes5{
		Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: int32(ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	controlRequest := e2appducontents.RiccontrolRequest{
		ProtocolIes: &e2appducontents.RiccontrolRequestIes{
			E2ApProtocolIes29: &ricRequestID,
			E2ApProtocolIes5:  &ranFunctionID,
		},
	}

	log.Info("Call Ric Channel")
	_, _, err = channel.RICControl(ctx, &controlRequest)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
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
	sub, err := s.subs.GetSubscription(server.Context(), &subapi.GetSubscriptionRequest{ID: subapi.ID(request.SubscriptionID)})
	if err != nil {
		return err
	}

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

		serviceModelID := modelregistry.ModelFullName(sub.Subscription.Details.ServiceModel.ID)
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
			response.IndicationHeader = indHeaderProto
			response.IndicationMessage = indMessageProto
			log.Infof("RICIndication successfully decoded from ASN1 to Proto #Bytes - Header: %d, Message: %d",
				len(indHeaderProto), len(indMessageProto))
		case e2api.EncodingType_ASN1_PER:
			response.IndicationHeader = indHeaderAsn1
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
