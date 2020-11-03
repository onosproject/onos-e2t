// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
