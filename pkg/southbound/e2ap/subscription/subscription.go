// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package subscription

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

type StreamID int

type Subscription interface {
	ID() e2api.SubscriptionID
	StreamID() StreamID
	Subscription() *e2api.Subscription
	In() chan<- *e2appducontents.Ricindication
	Out() <-chan *e2appducontents.Ricindication
	Close()
}

func newSubscriptionStream(id StreamID, sub *e2api.Subscription, manager *subscriptionManager) Subscription {
	return &subscriptionStream{
		manager: manager,
		id:      id,
		sub:     sub,
		ch:      make(chan *e2appducontents.Ricindication),
	}
}

type subscriptionStream struct {
	manager *subscriptionManager
	id      StreamID
	sub     *e2api.Subscription
	ch      chan *e2appducontents.Ricindication
}

func (s *subscriptionStream) ID() e2api.SubscriptionID {
	return s.sub.ID
}

func (s *subscriptionStream) Channel() *e2api.Subscription {
	return s.sub
}

func (s *subscriptionStream) StreamID() StreamID {
	return s.id
}

func (s *subscriptionStream) Subscription() *e2api.Subscription {
	return s.sub
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
