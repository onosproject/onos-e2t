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
func NewBroker() *Broker {
	subs := &SubscriptionManager{
		streams: &StreamManager{
			streams: make(map[StreamID]*Stream),
		},
		subs:     make(map[e2api.SubscriptionID]*SubscriptionStream),
		watchers: make(map[uuid.UUID]chan<- e2api.SubscriptionID),
	}
	return &Broker{
		streams: subs.streams,
		subs:    subs,
	}
}

// Broker is a subscription broker
type Broker struct {
	streams *StreamManager
	subs    *SubscriptionManager
}

func (s *Broker) Streams() *StreamManager {
	return s.streams
}

func (s *Broker) Subscriptions() *SubscriptionManager {
	return s.subs
}
