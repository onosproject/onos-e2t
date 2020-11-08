// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	subapi "github.com/onosproject/onos-e2t/api/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/config"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
	"sync"
)

// newDispatcher creates a new subscription dispatcher
func newDispatcher(channel channel.Channel, subs subscription.Store, streams *stream.Manager) (*Dispatcher, error) {
	dispatcher := &Dispatcher{
		channel:   channel,
		streams:   streams,
		listeners: make(map[ListenerID]*Listener),
	}
	if err := dispatcher.open(subs); err != nil {
		return nil, err
	}
	return dispatcher, nil
}

// Dispatcher is a subscription dispatcher
type Dispatcher struct {
	channel   channel.Channel
	streams   *stream.Manager
	listeners map[ListenerID]*Listener
	mu        sync.RWMutex
}

// open opens the dispatcher
func (d *Dispatcher) open(subs subscription.Store) error {
	subCh := make(chan subscription.Event)
	if err := subs.Watch(d.channel.Context(), subCh); err != nil {
		return err
	}
	go d.processSubscriptions(subCh)

	ricRequestID := &e2apies.RicrequestId{
		RicInstanceId: config.InstanceID,
	}
	indCh := d.channel.Recv(filter.RicIndication(ricRequestID), codec.XER)
	go d.processIndications(indCh)
	return nil
}

func (d *Dispatcher) processSubscriptions(ch <-chan subscription.Event) {
	for event := range ch {
		d.processSubscription(event)
	}
}

func (d *Dispatcher) processSubscription(event subscription.Event) {
	switch event.Type {
	case subapi.EventType_UPDATED:
		if err := d.processSubscriptionUpdated(event.Subscription); err != nil {
			log.Errorf("Failed to process subscription event %v: %v", event, err)
		}
	case subapi.EventType_REMOVED:
		if err := d.processSubscriptionRemoved(event.Subscription); err != nil {
			log.Errorf("Failed to process subscription event %v: %v", event, err)
		}
	}
}

func (d *Dispatcher) processSubscriptionUpdated(sub subapi.Subscription) error {
	if channel.ID(sub.Status.E2ConnID) != d.channel.ID() || sub.Status.E2RequestID == 0 {
		return nil
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	listenerID := ListenerID(sub.Status.E2RequestID)
	_, ok := d.listeners[listenerID]
	if !ok {
		l, err := newListener(listenerID, sub, d.streams)
		if err != nil {
			return err
		}
		d.listeners[listenerID] = l
	}
	return nil
}

func (d *Dispatcher) processSubscriptionRemoved(sub subapi.Subscription) error {
	if channel.ID(sub.Status.E2ConnID) != d.channel.ID() || sub.Status.E2RequestID == 0 {
		return nil
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	listenerID := ListenerID(sub.Status.E2RequestID)
	listener, ok := d.listeners[listenerID]
	if ok {
		delete(d.listeners, listenerID)
		return listener.Close()
	}
	return nil
}

// readIndications reads indications from the connection
func (d *Dispatcher) processIndications(ch <-chan *e2appdudescriptions.E2ApPdu) {
	for indication := range ch {
		requestID := indication.GetInitiatingMessage().ProcedureCode.RicIndication.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value
		listenerID := ListenerID(requestID.RicRequestorId)
		d.mu.RLock()
		listener, ok := d.listeners[listenerID]
		d.mu.RUnlock()
		if ok {
			err := listener.Notify(indication)
			if err != nil {
				log.Errorf("Failed to process indication %+v : %v", indication, err)
			}
		}
	}
}

// Close closes the dispatcher
func (d *Dispatcher) Close() error {
	d.mu.Lock()
	var err error
	for _, listener := range d.listeners {
		if e := listener.Close(); e != nil {
			err = e
		}
	}
	d.mu.Unlock()
	return err
}
