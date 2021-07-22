// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package ricapie2

import (
	"context"
	"io"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-e2t/pkg/topo"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/oid"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"

	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"

	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "e2")

// NewService creates a new E2T service
func NewService(subs subapi.E2SubscriptionServiceClient, streams subscription.Broker, modelRegistry modelregistry.ModelRegistry,
	channels e2server.ChannelManager, oidRegistry oid.Registry, topoManager topo.Manager) northbound.Service {
	return &Service{
		subs:          subs,
		streams:       streams,
		modelRegistry: modelRegistry,
		channels:      channels,
		oidRegistry:   oidRegistry,
		topoManager:   topoManager,
	}
}

// Service is a Service implementation for E2T service.
type Service struct {
	northbound.Service
	subs          subapi.E2SubscriptionServiceClient
	streams       subscription.Broker
	modelRegistry modelregistry.ModelRegistry
	channels      e2server.ChannelManager
	oidRegistry   oid.Registry
	topoManager   topo.Manager
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{subs: s.subs,
		streams:       s.streams,
		modelRegistry: s.modelRegistry,
		channels:      s.channels,
		oidRegistry:   s.oidRegistry,
		topoManager:   s.topoManager}
	e2api.RegisterE2TServiceServer(r, server)
}

// Server implements the gRPC service for E2 ricapi related functions.
type Server struct {
	subs             subapi.E2SubscriptionServiceClient
	streams          subscription.Broker
	modelRegistry    modelregistry.ModelRegistry
	channels         e2server.ChannelManager
	controlRequestID RequestID
	oidRegistry      oid.Registry
	topoManager      topo.Manager
}

func getControlAckRequest(request *e2api.ControlRequest) e2apies.RiccontrolAckRequest {
	var controlAckRequest e2apies.RiccontrolAckRequest
	switch request.ControlAckRequest {
	case e2api.ControlAckRequest_ACK:
		controlAckRequest = e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_ACK
	case e2api.ControlAckRequest_NACK:
		controlAckRequest = e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_N_ACK
	case e2api.ControlAckRequest_NO_ACK:
		controlAckRequest = e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_NO_ACK
	default:
		controlAckRequest = e2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_ACK
	}
	return controlAckRequest
}

func (s *Server) Control(ctx context.Context, request *e2api.ControlRequest) (*e2api.ControlResponse, error) {
	log.Infof("Received E2 Control Request %v", request)

	channel, err := s.channels.Get(ctx, topoapi.ID(request.GetE2NodeID()))
	if err != nil {
		return nil, errors.Status(err).Err()
	}
	response := &e2api.ControlResponse{}
	serviceModelOID, err := oid.ModelIDToOid(s.oidRegistry,
		string(request.Header.ServiceModel.Name),
		string(request.Header.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}
	serviceModelPlugin, err := s.modelRegistry.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}
	s.controlRequestID++
	requestID := s.controlRequestID

	ricRequest := types.RicRequest{
		RequestorID: types.RicRequestorID(requestID),
		InstanceID:  config.InstanceID,
	}

	var controlHeaderBytes []byte
	var controlMessageBytes []byte

	if request.Header.EncodingType == e2api.EncodingType_ASN1_PER ||
		request.Header.EncodingType == e2api.EncodingType_ASN1_XER {
		controlHeaderBytes = request.ControlHeader
		controlMessageBytes = request.ControlMessage
	} else if request.Header.EncodingType == e2api.EncodingType_PROTO {
		controlHeaderBytes = request.ControlHeader
		controlMessageBytes = request.ControlMessage
		controlHeader, err := serviceModelPlugin.ControlHeaderProtoToASN1(controlHeaderBytes)
		if err != nil {
			log.Warnf("Error transforming Control Header Proto bytes to ASN: %s", err.Error())
			return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
		}
		controlHeaderBytes = controlHeader
		controlMessage, err := serviceModelPlugin.ControlMessageProtoToASN1(controlMessageBytes)
		if err != nil {
			log.Warnf("Error transforming Control Message Proto bytes to ASN: %s", err.Error())
			return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
		}
		controlMessageBytes = controlMessage
	} else {
		err = errors.New(errors.Invalid, "invalid encoding type")
		log.Warn(err)
		return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
	}

	// TODO to keep admin API the channel ID is used for ran function registry mapping but
	//  should be changed to e2nodeID later one
	ranFuncID, ok := channel.GetRANFunction(serviceModelOID)
	if !ok {
		log.Warn("RAN function not found for SM %s", serviceModelOID)
	}

	controlAckRequest := getControlAckRequest(request)
	controlRequest, err := pdubuilder.NewControlRequest(ricRequest, ranFuncID.ID, nil, controlHeaderBytes, controlMessageBytes, &controlAckRequest)

	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}

	ack, failure, err := channel.RICControl(ctx, controlRequest)
	if err != nil {
		log.Warn(err)
		return nil, errors.Status(err).Err()
	}

	if ack != nil {
		if request.Header.EncodingType == e2api.EncodingType_PROTO {
			outcomeProtoBytes, err := serviceModelPlugin.ControlOutcomeASN1toProto(ack.ProtocolIes.E2ApProtocolIes32.Value.Value)
			if err != nil {
				log.Warnf("Error transforming Control Outcome ASN1 to Proto bytes: %s", err.Error())
				return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
			}
			response = &e2api.ControlResponse{
				Response: &e2api.ControlResponse_ControlAcknowledge{
					ControlAcknowledge: &e2api.ControlAcknowledge{
						ControlOutcome: outcomeProtoBytes,
					},
				},
			}
		} else {
			response = &e2api.ControlResponse{
				Response: &e2api.ControlResponse_ControlAcknowledge{
					ControlAcknowledge: &e2api.ControlAcknowledge{
						ControlOutcome: ack.ProtocolIes.E2ApProtocolIes32.Value.Value,
					},
				},
			}
		}
	}

	if failure != nil {
		if request.Header.EncodingType == e2api.EncodingType_PROTO {
			outcomeProtoBytes, err := serviceModelPlugin.ControlOutcomeASN1toProto(failure.ProtocolIes.E2ApProtocolIes32.Value.Value)
			if err != nil {
				log.Warnf("Error transforming Control Outcome ASN1 to Proto bytes: %s", err.Error())
				return nil, errors.Status(errors.NewInvalid(err.Error())).Err()
			}
			response = &e2api.ControlResponse{
				Response: &e2api.ControlResponse_ControlFailure{
					ControlFailure: &e2api.ControlFailure{
						ControlOutcome: outcomeProtoBytes,
					},
				},
			}
		} else {
			response = &e2api.ControlResponse{
				Response: &e2api.ControlResponse_ControlAcknowledge{
					ControlAcknowledge: &e2api.ControlAcknowledge{
						ControlOutcome: failure.ProtocolIes.E2ApProtocolIes32.Value.Value,
					},
				},
			}
		}
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
		log.Warnf("StreamRequest %+v failed: %v", request, err)
		return err
	}
	encodingType := request.GetHeader().GetEncodingType()

	log.Infof("Received StreamRequest %+v", request)
	sub, err := s.subs.GetSubscription(server.Context(), &subapi.GetSubscriptionRequest{ID: subapi.ID(request.SubscriptionID)})
	if err != nil {
		log.Warnf("StreamRequest %+v failed: %v", request, err)
		return err
	}

	reader, err := s.streams.OpenStream(subapi.ID(request.SubscriptionID))
	if err != nil {
		log.Warnf("StreamRequest %+v failed: %v", request, err)
		return err
	}

	for {
		indication, err := reader.Recv(server.Context())
		if err != nil {
			log.Warnf("StreamRequest %+v failed: %v", request, err)
			return err
		}

		ranFuncID := indication.ProtocolIes.E2ApProtocolIes5.Value.Value
		ricActionID := indication.ProtocolIes.E2ApProtocolIes15.Value.Value
		indHeaderAsn1 := indication.ProtocolIes.E2ApProtocolIes25.Value.Value
		indMessageAsn1 := indication.ProtocolIes.E2ApProtocolIes26.Value.Value
		log.Infof("Ric Indication. Ran FundID: %d, Ric Action ID: %d", ranFuncID, ricActionID)

		response := &e2api.StreamResponse{
			Header: &e2api.ResponseHeader{
				EncodingType: encodingType,
			},
		}

		serviceModelOID, err := oid.ModelIDToOid(s.oidRegistry,
			string(sub.Subscription.Details.ServiceModel.Name),
			string(sub.Subscription.Details.ServiceModel.Version))
		if err != nil {
			log.Warn(err)
			return err
		}
		switch encodingType {
		case e2api.EncodingType_PROTO:
			serviceModelPlugin, err := s.modelRegistry.GetPlugin(serviceModelOID)
			if err != nil {
				log.Warn(err)
				return errors.Status(err).Err()
			}
			smData := serviceModelPlugin.ServiceModelData()
			log.Infof("Service model found %s %s %s", smData.Name, smData.Version, smData.OID)

			indHeaderProto, err := serviceModelPlugin.IndicationHeaderASN1toProto(indHeaderAsn1)
			if err != nil {
				log.Errorf("Error transforming Header ASN Bytes to Proto %s", err.Error())
				return errors.Status(errors.NewInvalid(err.Error())).Err()
			}
			log.Infof("Indication Header %d bytes", len(indHeaderProto))

			indMessageProto, err := serviceModelPlugin.IndicationMessageASN1toProto(indMessageAsn1)
			if err != nil {
				log.Errorf("Error transforming Message ASN Bytes to Proto %s", err.Error())
				return errors.Status(errors.NewInvalid(err.Error())).Err()
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
			err = errors.NewInvalid("encoding type %v not supported", request.Header.EncodingType)
			return errors.Status(err).Err()
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
