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

// AddSubscription add a subscription
func (s Server) AddSubscription(ctx context.Context, req *subapi.SubscribeRequest) (*subapi.SubscribeResponse, error) {
	// TODO implement AddSubscription
	log.Info("Adding subscription")
	return &subapi.SubscribeResponse{}, nil
}

// DeleteSubscription delete a subscription
func (s Server) DeleteSubscription(ctx context.Context, req *subapi.SubscribeDeleteRequest) (*subapi.SubscribeDeleteResponse, error) {
	// TODO implement DeleteSubscription
	log.Info("Deleting subscription")
	return &subapi.SubscribeDeleteResponse{}, nil
}
