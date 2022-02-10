// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/cenkalti/backoff"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	"github.com/onosproject/onos-e2t/pkg/northbound/e2/stream"
	"io"
	"time"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/gogo/protobuf/proto"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"

	channelstore "github.com/onosproject/onos-e2t/pkg/store/channel"

	"github.com/onosproject/onos-e2t/pkg/oid"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

// NewSubscriptionService creates a new E2T subscription service
func NewSubscriptionService(chans channelstore.Store, subs substore.Store, streams stream.Manager, modelRegistry modelregistry.ModelRegistry, oidRegistry oid.Registry, rnib rnib.Store) northbound.Service {
	return &SubscriptionService{
		chans:         chans,
		subs:          subs,
		streams:       streams,
		modelRegistry: modelRegistry,
		oidRegistry:   oidRegistry,
		rnib:          rnib,
	}
}

// SubscriptionService is a Service implementation for E2 Subscription service.
type SubscriptionService struct {
	northbound.Service
	chans         channelstore.Store
	subs          substore.Store
	streams       stream.Manager
	modelRegistry modelregistry.ModelRegistry
	oidRegistry   oid.Registry
	rnib          rnib.Store
}

// Register registers the SubscriptionService with the gRPC server.
func (s SubscriptionService) Register(r *grpc.Server) {
	server := &SubscriptionServer{
		chans:         s.chans,
		subs:          s.subs,
		streams:       s.streams,
		modelRegistry: s.modelRegistry,
		oidRegistry:   s.oidRegistry,
		rnib:          s.rnib,
	}
	e2api.RegisterSubscriptionServiceServer(r, server)
	e2api.RegisterSubscriptionAdminServiceServer(r, server)
}

// SubscriptionServer implements the gRPC service for E2 Subscription related functions.
type SubscriptionServer struct {
	chans         channelstore.Store
	subs          substore.Store
	streams       stream.Manager
	modelRegistry modelregistry.ModelRegistry
	oidRegistry   oid.Registry
	rnib          rnib.Store
}

func (s *SubscriptionServer) GetChannel(ctx context.Context, request *e2api.GetChannelRequest) (*e2api.GetChannelResponse, error) {
	log.Debugf("Received GetChannelRequest %+v", request)
	channel, err := s.chans.Get(ctx, request.ChannelID)
	if err != nil {
		log.Warnf("GetChannelRequest %+v failed: %s", request, err)
		return nil, errors.Status(err).Err()
	}
	response := &e2api.GetChannelResponse{
		Channel: *channel,
	}
	log.Debugf("Sending GetChannelResponse %+v", response)
	return response, nil
}

func (s *SubscriptionServer) ListChannels(ctx context.Context, request *e2api.ListChannelsRequest) (*e2api.ListChannelsResponse, error) {
	log.Debugf("Received ListChannelsRequest %+v", request)
	channels, err := s.chans.List(ctx)
	if err != nil {
		log.Warnf("ListChannelsRequest %+v failed: %s", request, err)
		return nil, errors.Status(err).Err()
	}
	response := &e2api.ListChannelsResponse{
		Channels: channels,
	}
	log.Debugf("Sending ListChannelsResponse %+v", response)
	return response, nil
}

func (s *SubscriptionServer) WatchChannels(request *e2api.WatchChannelsRequest, server e2api.SubscriptionAdminService_WatchChannelsServer) error {
	log.Debugf("Received WatchChannelsRequest %+v", request)
	eventCh := make(chan e2api.ChannelEvent)
	var opts []channelstore.WatchOption
	if !request.NoReplay {
		opts = append(opts, channelstore.WithReplay())
	}
	if err := s.chans.Watch(server.Context(), eventCh, opts...); err != nil {
		log.Warnf("WatchChannelsRequest %+v failed: %s", request, err)
		return errors.Status(err).Err()
	}

	for event := range eventCh {
		response := &e2api.WatchChannelsResponse{
			Event: event,
		}
		log.Debugf("Sending WatchChannelsResponse %+v", response)
		err := server.Send(response)
		if err != nil {
			log.Warnf("Sending WatchChannelsResponse %+v failed: %s", response, err)
			return err
		}
	}
	return nil
}

func (s *SubscriptionServer) GetSubscription(ctx context.Context, request *e2api.GetSubscriptionRequest) (*e2api.GetSubscriptionResponse, error) {
	log.Debugf("Received GetSubscriptionRequest %+v", request)
	sub, err := s.subs.Get(ctx, request.SubscriptionID)
	if err != nil {
		log.Warnf("GetSubscriptionRequest %+v failed: %s", request, err)
		return nil, errors.Status(err).Err()
	}
	response := &e2api.GetSubscriptionResponse{
		Subscription: *sub,
	}
	log.Debugf("Sending GetSubscriptionResponse %+v", response)
	return response, nil
}

func (s *SubscriptionServer) ListSubscriptions(ctx context.Context, request *e2api.ListSubscriptionsRequest) (*e2api.ListSubscriptionsResponse, error) {
	log.Debugf("Received ListSubscriptionsRequest %+v", request)
	subs, err := s.subs.List(ctx)
	if err != nil {
		log.Warnf("ListSubscriptionsRequest %+v failed: %s", request, err)
		return nil, errors.Status(err).Err()
	}
	response := &e2api.ListSubscriptionsResponse{
		Subscriptions: subs,
	}
	log.Debugf("Sending ListSubscriptionsResponse %+v", response)
	return response, nil
}

func (s *SubscriptionServer) WatchSubscriptions(request *e2api.WatchSubscriptionsRequest, server e2api.SubscriptionAdminService_WatchSubscriptionsServer) error {
	log.Debugf("Received WatchSubscriptionsRequest %+v", request)
	eventCh := make(chan e2api.SubscriptionEvent)
	var opts []substore.WatchOption
	if !request.NoReplay {
		opts = append(opts, substore.WithReplay())
	}
	if err := s.subs.Watch(server.Context(), eventCh, opts...); err != nil {
		log.Warnf("WatchSubscriptionsRequest %+v failed: %s", request, err)
		return errors.Status(err).Err()
	}

	for event := range eventCh {
		response := &e2api.WatchSubscriptionsResponse{
			Event: event,
		}
		log.Debugf("Sending WatchSubscriptionsResponse %+v", response)
		err := server.Send(response)
		if err != nil {
			log.Warnf("Sending WatchSubscriptionResponse %+v failed: %s", response, err)
			return err
		}
	}
	return nil
}

func (s *SubscriptionServer) Subscribe(request *e2api.SubscribeRequest, server e2api.SubscriptionService_SubscribeServer) error {
	log.Debugf("Received SubscribeRequest %+v", request)
	encoding := request.Headers.Encoding

	serviceModelOID, err := oid.ModelIDToOid(s.oidRegistry,
		string(request.Headers.ServiceModel.Name),
		string(request.Headers.ServiceModel.Version))
	if err != nil {
		log.Warnf("SubscribeRequest %+v failed", request, err)
		return err
	}

	serviceModelPlugin, err := s.modelRegistry.GetPlugin(serviceModelOID)
	if err != nil {
		log.Warnf("SubscribeRequest %+v failed", request, err)
		return errors.Status(errors.NewNotFound(err.Error())).Err()
	}

	smData := serviceModelPlugin.ServiceModelData()
	log.Debugf("Service model found %s %s %s", smData.Name, smData.Version, smData.OID)

	if encoding != e2api.Encoding_PROTO && encoding != e2api.Encoding_ASN1_PER {
		err = errors.NewNotSupported("encoding type %s is not supported", encoding)
		log.Warnf("SubscribeRequest %+v failed: %s", request, err)
		return errors.Status(errors.NewInvalid(err.Error())).Err()
	}

	subSpec := request.Subscription
	if encoding == e2api.Encoding_PROTO {
		eventTriggerBytes, err := serviceModelPlugin.EventTriggerDefinitionProtoToASN1(subSpec.EventTrigger.Payload)
		if err != nil {
			log.Warnf("SubscribeRequest %+v failed", request, err)
			return errors.Status(errors.NewInvalid(err.Error())).Err()
		}
		subSpec.EventTrigger.Payload = eventTriggerBytes
	}

	for i, action := range subSpec.Actions {
		if encoding == e2api.Encoding_PROTO && action.Payload != nil {
			actionBytes, err := serviceModelPlugin.ActionDefinitionProtoToASN1(action.Payload)
			if err != nil {
				log.Warnf("SubscribeRequest %+v failed", request, err)
				return errors.Status(errors.NewInvalid(err.Error())).Err()
			}
			action.Payload = actionBytes
			subSpec.Actions[i] = action
		}
	}

	subBytes, err := proto.Marshal(&subSpec)
	if err != nil {
		log.Warnf("SubscribeRequest %+v failed", request, err)
		return errors.Status(errors.NewInvalid(err.Error())).Err()
	}

	subID := e2api.SubscriptionID(fmt.Sprintf("%x:%s", md5.Sum(subBytes), request.Headers.E2NodeID))

	channelID := e2api.ChannelID(fmt.Sprintf("%s:%s:%s:%s",
		request.Headers.AppID,
		request.Headers.AppInstanceID,
		request.Headers.E2NodeID,
		request.TransactionID))
	channelMeta := e2api.ChannelMeta{
		AppID:          request.Headers.AppID,
		AppInstanceID:  request.Headers.AppInstanceID,
		E2NodeID:       request.Headers.E2NodeID,
		TransactionID:  request.TransactionID,
		SubscriptionID: subID,
		ServiceModel:   request.Headers.ServiceModel,
		Encoding:       e2api.Encoding_ASN1_PER,
	}

	err = backoff.Retry(func() error {
		if channel, err := s.chans.Get(server.Context(), channelID); err == nil {
			// Change the channel phase to OPEN if necessary
			if channel.Status.Phase != e2api.ChannelPhase_CHANNEL_OPEN {
				channel.Status.Phase = e2api.ChannelPhase_CHANNEL_OPEN
				channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
				channel.Status.Error = nil
				now := time.Now()
				channel.Status.Timestamp = &now
				if err := s.chans.Update(server.Context(), channel); err != nil {
					if !errors.IsNotFound(err) && !errors.IsConflict(err) {
						return backoff.Permanent(err)
					}
					return err
				}
			}
		} else if errors.IsNotFound(err) {
			// Create the channel if necessary
			now := time.Now()
			channel := &e2api.Channel{
				ID:          channelID,
				ChannelMeta: channelMeta,
				Spec: e2api.ChannelSpec{
					SubscriptionSpec:   subSpec,
					TransactionTimeout: request.TransactionTimeout,
				},
				Status: e2api.ChannelStatus{
					Phase:     e2api.ChannelPhase_CHANNEL_OPEN,
					State:     e2api.ChannelState_CHANNEL_PENDING,
					Timestamp: &now,
				},
			}

			if err := s.chans.Create(server.Context(), channel); err != nil {
				if !errors.IsAlreadyExists(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		} else {
			return backoff.Permanent(err)
		}
		return nil
	}, backoff.WithContext(backoff.NewExponentialBackOff(), server.Context()))
	if err != nil {
		log.Warnf("SubscribeRequest %+v failed", request, err)
		return errors.Status(err).Err()
	}

	stream := s.streams.Open(channelID, channelMeta).Output().Open(server.Context())

	select {
	case <-stream.Ready():
		err := server.Send(&e2api.SubscribeResponse{
			Message: &e2api.SubscribeResponse_Ack{
				Ack: &e2api.Acknowledgement{
					ChannelID: channelID,
				},
			},
		})
		if err != nil {
			log.Warnf("SubscribeRequest %+v failed", request, err)
			return err
		}
	case <-stream.Done():
		err := stream.Err()
		if err != nil {
			log.Warnf("SubscribeRequest %+v failed", request, err)
			if _, ok := err.(*errors.TypedError); ok {
				return errors.Status(err).Err()
			}
			return err
		}
		log.Debugf("SubscribeRequest %+v complete", request)
		return nil
	case <-server.Context().Done():
		log.Debugf("SubscribeRequest %+v closed", request)
		return server.Context().Err()
	}

	for {
		select {
		case ind, ok := <-stream.Indications():
			if !ok {
				return errors.Status(errors.NewUnavailable("stream closed")).Err()
			}

			var ranFuncID int32
			var ricActionID int32
			var indHeaderAsn1 []byte
			var indMessageAsn1 []byte
			for _, v := range ind.GetProtocolIes() {
				if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
					ranFuncID = v.GetValue().GetRfId().GetValue()
				}
				if v.Id == int32(v2.ProtocolIeIDRicactionID) {
					ricActionID = v.GetValue().GetRaId().GetValue()
				}
				if v.Id == int32(v2.ProtocolIeIDRicindicationHeader) {
					indHeaderAsn1 = v.GetValue().GetRih().GetValue()
				}
				if v.Id == int32(v2.ProtocolIeIDRicindicationMessage) {
					indMessageAsn1 = v.GetValue().GetRim().GetValue()
				}
			}
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
				log.Errorf("Encoding type %v not supported", encoding)
				return errors.Status(errors.NewInvalid("encoding type %v not supported", encoding)).Err()
			}

			log.Debugf("Sending SubscribeResponse %+v", response)
			err = server.Send(response)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				log.Warnf("Sending SubscribeResponse %+v failed", response, err)
				return err
			}
		case <-stream.Done():
			err := stream.Err()
			if err != nil {
				log.Warnf("SubscribeRequest %+v failed", request, err)
				if _, ok := err.(*errors.TypedError); ok {
					return errors.Status(err).Err()
				}
				return err
			}
			log.Debugf("SubscribeRequest %+v complete", request)
			return nil
		case <-server.Context().Done():
			log.Debugf("SubscribeRequest %+v closed", request)
			return server.Context().Err()
		}
	}
}

func (s *SubscriptionServer) Unsubscribe(ctx context.Context, request *e2api.UnsubscribeRequest) (*e2api.UnsubscribeResponse, error) {
	log.Debugf("Received UnsubscribeRequest %+v", request)
	channelID := e2api.ChannelID(fmt.Sprintf("%s:%s:%s:%s",
		request.Headers.AppID,
		request.Headers.AppInstanceID,
		request.Headers.E2NodeID,
		request.TransactionID))

	err := backoff.Retry(func() error {
		channel, err := s.chans.Get(ctx, channelID)
		if err != nil {
			if errors.IsNotFound(err) {
				return nil
			}
			return err
		}
		if channel.Status.Phase != e2api.ChannelPhase_CHANNEL_CLOSED {
			channel.Status.Phase = e2api.ChannelPhase_CHANNEL_CLOSED
			channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
			channel.Status.Error = nil
			err := s.chans.Update(ctx, channel)
			if err != nil {
				if errors.IsNotFound(err) {
					return nil
				}
				if errors.IsConflict(err) {
					return err
				}
				return backoff.Permanent(err)
			}
		}
		return nil
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		log.Warnf("UnsubscribeRequest %+v failed: %s", request, err)
		return nil, errors.Status(err).Err()
	}

	response := &e2api.UnsubscribeResponse{}
	log.Debugf("Sending UnsubscribeResponse %+v", response)
	return response, nil
}
