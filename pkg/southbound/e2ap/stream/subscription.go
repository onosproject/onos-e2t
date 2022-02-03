// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

type ID int

type Subscription interface {
	ID() e2api.SubscriptionID
	StreamID() ID
	In() chan<- *e2appducontents.Ricindication
	Out() <-chan *e2appducontents.Ricindication
	Close()
}

func newSubscriptionStream(id e2api.SubscriptionID, streamID ID, manager *subscriptionManager) Subscription {
	return &subscriptionStream{
		manager:  manager,
		id:       id,
		streamID: streamID,
		ch:       make(chan *e2appducontents.Ricindication),
	}
}

type subscriptionStream struct {
	manager  *subscriptionManager
	streamID ID
	id       e2api.SubscriptionID
	ch       chan *e2appducontents.Ricindication
}

func (s *subscriptionStream) ID() e2api.SubscriptionID {
	return s.id
}

func (s *subscriptionStream) StreamID() ID {
	return s.streamID
}

func (s *subscriptionStream) In() chan<- *e2appducontents.Ricindication {
	return s.ch
}

func (s *subscriptionStream) Out() <-chan *e2appducontents.Ricindication {
	return s.ch
}

func (s *subscriptionStream) Close() {
	close(s.ch)
	s.manager.close(s)
}
