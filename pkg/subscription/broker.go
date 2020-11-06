// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
	"io"
)

// NewBroker creates a new subscription broker
func NewBroker(channel *channel.Manager, subs subscription.Store, streams *stream.Manager) (*Broker, error) {
	broker := &Broker{
		channels: channel,
		subs:     subs,
		streams:  streams,
	}
	if err := broker.open(); err != nil {
		return nil, err
	}
	return broker, nil
}

// Broker is a subscription broker
type Broker struct {
	channels *channel.Manager
	subs     subscription.Store
	streams  *stream.Manager
	cancel   context.CancelFunc
}

// open opens the broker
func (b *Broker) open() error {
	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel

	channelCh := make(chan channel.Channel)
	if err := b.channels.Watch(ctx, channelCh); err != nil {
		cancel()
		return err
	}
	go b.processChannels(channelCh)
	return nil
}

// processChannels processes channel events
func (b *Broker) processChannels(ch <-chan channel.Channel) {
	for conn := range ch {
		b.processChannel(conn)
	}
}

// processChannel processes a channel event
func (b *Broker) processChannel(channel channel.Channel) {
	dispatcher, err := newDispatcher(channel, b.subs, b.streams)
	if err != nil {
		log.Errorf("Failed to create dispatcher: %v", err)
	} else {
		go func() {
			<-channel.Context().Done()
			err := dispatcher.Close()
			if err != nil {
				log.Errorf("Failed to close dispatcher: %v", err)
			}
		}()
	}
}

// Close closes the broker
func (b *Broker) Close() error {
	b.cancel()
	return nil
}

var _ io.Closer = &Broker{}
