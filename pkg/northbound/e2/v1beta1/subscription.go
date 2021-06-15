// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gogo/protobuf/proto"

	substoreapi "github.com/onosproject/onos-e2t/api/onos/e2t/store/subscription"
	subbroker "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"

	"github.com/onosproject/onos-e2t/pkg/oid"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

// NewSubscriptionService creates a new E2T subscription service
func NewSubscriptionService(subs substore.Store, streams subbroker.Broker, modelRegistry modelregistry.ModelRegistry, oidRegistry oid.Registry) northbound.Service {
	return &SubscriptionService{
		subs:          subs,
		streams:       streams,
		modelRegistry: modelRegistry,
		oidRegistry:   oidRegistry,
	}
}

// SubscriptionService is a Service implementation for E2 Subscription service.
type SubscriptionService struct {
	northbound.Service
	subs          substore.Store
	streams       subbroker.Broker
	modelRegistry modelregistry.ModelRegistry
	oidRegistry   oid.Registry
}

// Register registers the SubscriptionService with the gRPC server.
func (s SubscriptionService) Register(r *grpc.Server) {
	server := &SubscriptionServer{
		subs:          s.subs,
		streams:       s.streams,
		modelRegistry: s.modelRegistry,
		oidRegistry:   s.oidRegistry,
	}
	e2api.RegisterSubscriptionServiceServer(r, server)
}

// SubscriptionServer implements the gRPC service for E2 Subscription related functions.
type SubscriptionServer struct {
	subs          substore.Store
	streams       subbroker.Broker
	modelRegistry modelregistry.ModelRegistry
	oidRegistry   oid.Registry
}

func (s *SubscriptionServer) Subscribe(request *e2api.SubscribeRequest, server e2api.SubscriptionService_SubscribeServer) error {
	encoding := request.Headers.Encoding

	log.Infof("Received SubscribeRequest %+v", request)

	serviceModelOID, err := oid.ModelIDToOid(s.oidRegistry,
		string(request.Headers.ServiceModel.Name),
		string(request.Headers.ServiceModel.Version))
	if err != nil {
		log.Warn(err)
		return err
	}

	serviceModelPlugin, err := s.modelRegistry.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warn(err)
		return errors.Status(err).Err()
	}
	smData := serviceModelPlugin.ServiceModelData()
	log.Infof("Service model found %s %s %s", smData.Name, smData.Version, smData.OID)

	eventTriggerBytes := request.Subscription.EventTrigger.Payload
	if encoding == e2api.Encoding_PROTO {
		eventTriggerBytes, err = serviceModelPlugin.EventTriggerDefinitionProtoToASN1(eventTriggerBytes)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	eventTrigger := &substoreapi.SubscriptionEventTrigger{
		Payload: eventTriggerBytes,
	}

	actions := make([]substoreapi.SubscriptionAction, len(request.Subscription.Actions))
	for i, action := range request.Subscription.Actions {
		actionBytes := action.Payload
		if encoding == e2api.Encoding_PROTO && action.Payload != nil {
			actionBytes, err = serviceModelPlugin.ActionDefinitionProtoToASN1(actionBytes)
			if err != nil {
				log.Error(err)
				return err
			}
		}
		subAction := substoreapi.SubscriptionAction{
			ID:      action.ID,
			Type:    substoreapi.ActionType(action.Type),
			Payload: actionBytes,
		}
		if action.SubsequentAction != nil {
			subAction.SubsequentAction = &substoreapi.SubsequentAction{
				Type:       substoreapi.SubsequentActionType(action.SubsequentAction.Type),
				TimeToWait: substoreapi.TimeToWait(action.SubsequentAction.TimeToWait),
			}
		}
		actions[i] = subAction
	}

	spec := substoreapi.SubscriptionSpec{
		EventTrigger: eventTrigger,
		Actions:      actions,
	}

	subBytes, err := proto.Marshal(&spec)
	if err != nil {
		log.Error(err)
		return err
	}
	subHash := fmt.Sprintf("%x", md5.Sum(subBytes))

	subID := substoreapi.SubscriptionID{
		NodeID:     substoreapi.NodeID(request.Headers.NodeID),
		AppID:      substoreapi.AppID(request.Headers.AppID),
		InstanceID: substoreapi.InstanceID(request.Headers.InstanceID),
		RequestID:  substoreapi.RequestID(request.Subscription.ID),
		Hash:       subHash,
	}

	_, err = s.subs.Get(server.Context(), subID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Error(err)
			return errors.Status(err).Err()
		}

		sub := &substoreapi.Subscription{
			SubscriptionMeta: substoreapi.SubscriptionMeta{
				ID: subID,
				ServiceModel: substoreapi.ServiceModel{
					Name:    substoreapi.ServiceModelName(request.Headers.ServiceModel.Name),
					Version: substoreapi.ServiceModelVersion(request.Headers.ServiceModel.Version),
				},
			},
			Spec: spec,
		}
		err = s.subs.Create(server.Context(), sub)
		if err != nil {
			return errors.Status(err).Err()
		}
	}

	reader := s.streams.OpenReader(subID)
	response := &e2api.SubscribeResponse{
		Headers: e2api.ResponseHeaders{
			Encoding: encoding,
		},
		Message: &e2api.SubscribeResponse_Ack{
			Ack: &e2api.Acknowledgement{},
		},
	}

	err = server.Send(response)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Warnf("SubscribeResponse %+v failed: %v", response, err)
		return err
	}

	for {
		indication, err := reader.Recv(server.Context())
		if err != nil {
			log.Warnf("SubscribeRequest %+v failed: %v", request, err)
			return err
		}

		ranFuncID := indication.ProtocolIes.E2ApProtocolIes5.Value.Value
		ricActionID := indication.ProtocolIes.E2ApProtocolIes15.Value.Value
		indHeaderAsn1 := indication.ProtocolIes.E2ApProtocolIes25.Value.Value
		indMessageAsn1 := indication.ProtocolIes.E2ApProtocolIes26.Value.Value
		log.Infof("Ric Indication. Ran FundID: %d, Ric Action ID: %d", ranFuncID, ricActionID)

		response := &e2api.SubscribeResponse{
			Headers: e2api.ResponseHeaders{
				Encoding: encoding,
			},
		}

		switch encoding {
		case e2api.Encoding_PROTO:
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
			response.Message = &e2api.SubscribeResponse_Indication{
				Indication: &e2api.Indication{
					Header:  indHeaderProto,
					Payload: indMessageProto,
				},
			}
			log.Infof("RICIndication successfully decoded from ASN1 to Proto #Bytes - Header: %d, Message: %d",
				len(indHeaderProto), len(indMessageProto))
		case e2api.Encoding_ASN1_PER:
			response.Message = &e2api.SubscribeResponse_Indication{
				Indication: &e2api.Indication{
					Header:  indHeaderAsn1,
					Payload: indMessageAsn1,
				},
			}
		default:
			log.Errorf("encoding type %v not supported", encoding)
			return errors.Status(errors.NewInvalid("encoding type %v not supported", encoding)).Err()
		}

		err = server.Send(response)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Warnf("SubscribeResponse %+v failed: %v", response, err)
			return err
		}
	}
}
