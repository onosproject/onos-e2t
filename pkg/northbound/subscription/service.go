// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"

	subapi "github.com/onosproject/onos-e2t/api/ricapi/e2/subscription/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("northbound", "ricapi", "subscription")

// Service is a Service implementation for subscription service.
type Service struct {
	northbound.Service
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := Server{}
	subapi.RegisterSubscriptionServiceServer(r, server)

}

// Server implements the gRPC service for managing of subscriptions
type Server struct {
}

// AddSubscription adds a subscription
func (s Server) AddSubscription(ctx context.Context, req *subapi.AddSubscriptionRequest) (*subapi.AddSubscriptionResponse, error) {
	log.Info("Adding subscription")
	return &subapi.AddSubscriptionResponse{}, nil
}

// RemoveSubscription removes a subscription
func (s Server) RemoveSubscription(ctx context.Context, req *subapi.RemoveSubscriptionRequest) (*subapi.RemoveSubscriptionResponse, error) {
	log.Info("Removing subscription")
	return &subapi.RemoveSubscriptionResponse{}, nil
}

// GetSubscription  retrieves information about a specific subscription in the list of existing subscriptions
func (s Server) GetSubscription(ctx context.Context, req *subapi.GetSubscriptionRequest) (*subapi.GetSubscriptionResponse, error) {
	return &subapi.GetSubscriptionResponse{}, nil
}

// ListSubscriptions returns the list of current existing subscriptions
func (s Server) ListSubscriptions(ctx context.Context, req *subapi.ListSubscriptionsRequest) (*subapi.ListSubscriptionsResponse, error) {
	return &subapi.ListSubscriptionsResponse{}, nil
}
