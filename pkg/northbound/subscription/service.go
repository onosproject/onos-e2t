// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	"io"

	subapi "github.com/onosproject/onos-e2t/api/ricapi/e2/subscription/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "subscription")

// NewService creates a new subscription service
func NewService(subscriptions substore.Store) northbound.Service {
	return &Service{subs: subscriptions}
}

// Service is a Service implementation for subscription service.
type Service struct {
	northbound.Service
	subs substore.Store
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{subs: s.subs}
	subapi.RegisterSubscriptionServiceServer(r, server)

}

// Server implements the gRPC service for managing of subscriptions
type Server struct {
	subs substore.Store
}

// AddSubscription adds a subscription
func (s *Server) AddSubscription(ctx context.Context, req *subapi.AddSubscriptionRequest) (*subapi.AddSubscriptionResponse, error) {
	log.Debugf("Received AddSubscriptionRequest %+v", req)
	sub := req.Subscription
	err := s.subs.Add(ctx, sub)
	if err != nil {
		log.Warnf("AddSubscriptionRequest %+v failed: %v", req, err)
		return nil, err
	}
	res := &subapi.AddSubscriptionResponse{
		Subscription: sub,
	}
	log.Debugf("Sending AddSubscriptionResponse %+v", res)
	return res, nil
}

// UpdateSubscription updates a subscription
func (s *Server) UpdateSubscription(ctx context.Context, req *subapi.UpdateSubscriptionRequest) (*subapi.UpdateSubscriptionResponse, error) {
	log.Debugf("Received UpdateSubscriptionRequest %+v", req)
	sub := req.Subscription
	err := s.subs.Update(ctx, sub)
	if err != nil {
		log.Warnf("UpdateSubscriptionRequest %+v failed: %v", req, err)
		return nil, err
	}
	res := &subapi.UpdateSubscriptionResponse{
		Subscription: sub,
	}
	log.Debugf("Sending UpdateSubscriptionResponse %+v", res)
	return res, nil
}

// RemoveSubscription removes a subscription
func (s *Server) RemoveSubscription(ctx context.Context, req *subapi.RemoveSubscriptionRequest) (*subapi.RemoveSubscriptionResponse, error) {
	log.Debugf("Received RemoveSubscriptionRequest %+v", req)
	sub := req.Subscription
	err := s.subs.Remove(ctx, sub)
	if err != nil {
		log.Warnf("RemoveSubscriptionRequest %+v failed: %v", req, err)
		return nil, err
	}
	res := &subapi.RemoveSubscriptionResponse{
		Subscription: sub,
	}
	log.Debugf("Sending RemoveSubscriptionResponse %+v", res)
	return res, nil
}

// GetSubscription retrieves information about a specific subscription in the list of existing subscriptions
func (s *Server) GetSubscription(ctx context.Context, req *subapi.GetSubscriptionRequest) (*subapi.GetSubscriptionResponse, error) {
	log.Debugf("Received GetSubscriptionRequest %+v", req)
	sub, err := s.subs.Get(ctx, req.ID)
	if err != nil {
		log.Warnf("GetSubscriptionRequest %+v failed: %v", req, err)
		return nil, err
	}
	res := &subapi.GetSubscriptionResponse{
		Subscription: sub,
	}
	log.Debugf("Sending GetSubscriptionResponse %+v", res)
	return res, nil
}

// ListSubscriptions returns the list of current existing subscriptions
func (s *Server) ListSubscriptions(ctx context.Context, req *subapi.ListSubscriptionsRequest) (*subapi.ListSubscriptionsResponse, error) {
	log.Debugf("Received ListSubscriptionsRequest %+v", req)
	subs, err := s.subs.List(ctx)
	if err != nil {
		log.Warnf("ListSubscriptionsRequest %+v failed: %v", req, err)
		return nil, err
	}
	res := &subapi.ListSubscriptionsResponse{
		Subscriptions: subs,
	}
	log.Debugf("Sending ListSubscriptionsResponse %+v", res)
	return res, nil
}

// WatchSubscriptions streams subscription changes
func (s *Server) WatchSubscriptions(req *subapi.WatchSubscriptionsRequest, stream subapi.SubscriptionService_WatchSubscriptionsServer) error {
	log.Debugf("Received WatchSubscriptionsRequest %+v", req)
	ch := make(chan substore.Event)
	err := s.subs.Watch(stream.Context(), ch)
	if err != nil {
		log.Warnf("WatchSubscriptionsRequest %+v failed: %v", req, err)
		return err
	}

	for event := range ch {
		res := &subapi.WatchSubscriptionsResponse{
			Type:         event.Type,
			Subscription: event.Subscription,
		}
		log.Debugf("Sending WatchSubscriptionsResponse %+v", res)
		err := stream.Send(res)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Warnf("WatchSubscriptionsResponse %+v failed: %v", req, err)
			return err
		}
	}
	return nil
}
