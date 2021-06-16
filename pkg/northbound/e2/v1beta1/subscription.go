// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import (
	"context"
	"crypto/md5"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"

	"github.com/gogo/protobuf/proto"

	channelbroker "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	channelstore "github.com/onosproject/onos-e2t/pkg/store/channel"

	"github.com/onosproject/onos-e2t/pkg/oid"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

// NewSubscriptionService creates a new E2T subscription service
func NewSubscriptionService(subs channelstore.Store, streams channelbroker.Broker, modelRegistry modelregistry.ModelRegistry, oidRegistry oid.Registry) northbound.Service {
	return &SubscriptionService{
		channels:      subs,
		streams:       streams,
		modelRegistry: modelRegistry,
		oidRegistry:   oidRegistry,
	}
}

// SubscriptionService is a Service implementation for E2 Subscription service.
type SubscriptionService struct {
	northbound.Service
	channels      channelstore.Store
	streams       channelbroker.Broker
	modelRegistry modelregistry.ModelRegistry
	oidRegistry   oid.Registry
}

// Register registers the SubscriptionService with the gRPC server.
func (s SubscriptionService) Register(r *grpc.Server) {
	server := &SubscriptionServer{
		channels:      s.channels,
		streams:       s.streams,
		modelRegistry: s.modelRegistry,
		oidRegistry:   s.oidRegistry,
	}
	e2api.RegisterSubscriptionServiceServer(r, server)
}

// SubscriptionServer implements the gRPC service for E2 Subscription related functions.
type SubscriptionServer struct {
	channels      channelstore.Store
	streams       channelbroker.Broker
	modelRegistry modelregistry.ModelRegistry
	oidRegistry   oid.Registry
}

func (s *SubscriptionServer) Unsubscribe(ctx context.Context, request *e2api.UnsubscribeRequest) (*e2api.UnsubscribeResponse, error) {
	panic("implement me")
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

	subSpec := request.Subscription
	if encoding == e2api.Encoding_PROTO {
		eventTriggerBytes, err := serviceModelPlugin.EventTriggerDefinitionProtoToASN1(subSpec.EventTrigger.Payload)
		if err != nil {
			log.Error(err)
			return err
		}
		subSpec.EventTrigger.Payload = eventTriggerBytes
	}

	for i, action := range subSpec.Actions {
		if encoding == e2api.Encoding_PROTO && action.Payload != nil {
			actionBytes, err := serviceModelPlugin.ActionDefinitionProtoToASN1(action.Payload)
			if err != nil {
				log.Error(err)
				return err
			}
			action.Payload = actionBytes
			subSpec.Actions[i] = action
		}
	}

	subBytes, err := proto.Marshal(&subSpec)
	if err != nil {
		log.Error(err)
		return err
	}
	subID := e2api.SubscriptionID(fmt.Sprintf("%x", md5.Sum(subBytes)))

	channelID := e2api.ChannelID(fmt.Sprintf("%s:%s:%s:%s",
		request.Headers.AppID,
		request.Headers.InstanceID,
		request.Headers.NodeID,
		request.TransactionID))

	_, err = s.channels.Get(server.Context(), channelID)
	if err != nil {
		if !errors.IsNotFound(err) {
			log.Error(err)
			return errors.Status(err).Err()
		}

		channel := &e2api.Channel{
			ID: channelID,
			ChannelMeta: e2api.ChannelMeta{
				AppID:          request.Headers.AppID,
				InstanceID:     request.Headers.InstanceID,
				NodeID:         request.Headers.NodeID,
				TransactionID:  request.TransactionID,
				SubscriptionID: subID,
				ServiceModel:   request.Headers.ServiceModel,
				Encoding:       e2api.Encoding_ASN1_PER,
			},
			Spec: e2api.ChannelSpec{
				SubscriptionSpec: subSpec,
			},
		}
		err = s.channels.Create(server.Context(), channel)
		if err != nil {
			return errors.Status(err).Err()
		}
	}

	// Open a stream reader for the app instance
	reader := s.streams.OpenReader(subID, request.Headers.AppID, request.Headers.InstanceID)

	// Watch the channel store for changes
	eventCh := make(chan e2api.ChannelEvent)
	ctx, cancel := context.WithCancel(server.Context())
	if err := s.channels.Watch(ctx, eventCh); err != nil {
		cancel()
		return errors.Status(err).Err()
	}

	// Wait for the channel state to indicate the subscription has been established
	for event := range eventCh {
		if event.Channel.ID == channelID && event.Channel.Status.Phase == e2api.ChannelPhase_CHANNEL_OPEN {
			switch event.Channel.Status.State {
			case e2api.ChannelState_CHANNEL_COMPLETE:
				cancel()

				// If the channel open is complete, send an ack response to the client
				response := &e2api.SubscribeResponse{
					Headers: e2api.ResponseHeaders{
						Encoding: encoding,
					},
					Message: &e2api.SubscribeResponse_Ack{
						Ack: &e2api.Acknowledgement{
							ChannelID: channelID,
						},
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
				break
			case e2api.ChannelState_CHANNEL_FAILED:
				// If the channel open failed, send the failure to the client as an error
				errStat := status.New(codes.Aborted, "an E2AP failure occurred")
				errStat, err := errStat.WithDetails(event.Channel.Status.Error)
				if err != nil {
					return err
				}
				cancel()
				return errStat.Err()
			}
		}
	}

	// Read indications from the stream and send them to the client
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

func (s *SubscriptionServer) Unsubscribe(ctx context.Context, request *e2api.UnsubscribeRequest) (*e2api.UnsubscribeResponse, error) {
	channelID := e2api.ChannelID(fmt.Sprintf("%s:%s:%s:%s",
		request.Headers.AppID,
		request.Headers.InstanceID,
		request.Headers.NodeID,
		request.TransactionID))

	// Get the channel for the subscription/app/instance
	channel, err := s.channels.Get(ctx, channelID)
	if err != nil {
		return nil, errors.Status(err).Err()
	}

	// Ensure the channel phase is CLOSED
	if channel.Status.Phase != e2api.ChannelPhase_CHANNEL_CLOSED {
		channel.Status.Phase = e2api.ChannelPhase_CHANNEL_CLOSED
		channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
		channel.Status.Error = nil
		if err := s.channels.Update(ctx, channel); err != nil {
			return nil, errors.Status(err).Err()
		}
	}

	// Watch the channel store for changes
	eventCh := make(chan e2api.ChannelEvent)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := s.channels.Watch(ctx, eventCh); err != nil {
		return nil, errors.Status(err).Err()
	}

	// Wait for the channel state to indicate the subscription has been established
	for event := range eventCh {
		if event.Channel.ID == channelID && event.Channel.Status.Phase == e2api.ChannelPhase_CHANNEL_CLOSED {
			switch event.Channel.Status.State {
			case e2api.ChannelState_CHANNEL_COMPLETE:
				s.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.InstanceID)
				return &e2api.UnsubscribeResponse{}, nil
			case e2api.ChannelState_CHANNEL_FAILED:
				s.streams.CloseReader(channel.SubscriptionID, channel.AppID, channel.InstanceID)
				errStat := status.New(codes.Aborted, "an E2AP failure occurred")
				errStat, err := errStat.WithDetails(event.Channel.Status.Error)
				if err != nil {
					return nil, err
				}
				return nil, errStat.Err()
			}
		}
	}
	return nil, ctx.Err()
}
