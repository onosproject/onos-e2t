// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("broker")

const bufferMaxSize = 10000

type StreamID int

// NewBroker creates a new subscription stream broker
func NewBroker() Broker {
	subs := &subscriptionManager{
		streams: &streamManager{
			streams: make(map[StreamID]*Stream),
		},
		subs:     make(map[e2api.SubscriptionID]*SubscriptionStream),
		watchers: make(map[uuid.UUID]chan<- e2api.SubscriptionID),
	}
	return &broker{
		streams: subs.streams,
		subs:    subs,
	}
}

// Broker is a subscription stream broker
type Broker interface {
	Streams() StreamManager
	Subscriptions() SubscriptionManager
}

// broker is a subscription broker
type broker struct {
	streams StreamManager
	subs    SubscriptionManager
}

func (s *broker) Streams() StreamManager {
	return s.streams
}

func (s *broker) Subscriptions() SubscriptionManager {
	return s.subs
}

var _ Broker = &broker{}
