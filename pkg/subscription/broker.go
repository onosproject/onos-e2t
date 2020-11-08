// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/store/subscription"
)

// NewBroker creates a new subscription broker
func NewBroker(subs subscription.Store, streams *stream.Manager, channels *channel.Manager) *Broker {
	return &Broker{
		subs:     subs,
		streams:  streams,
		channels: channels,
	}
}

// Broker is a subscription broker
type Broker struct {
	subs     subscription.Store
	streams  *stream.Manager
	channels *channel.Manager
	cancel   context.CancelFunc
}

// Start starts the broker
func (b *Broker) Start() error {
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

// Stop stops the broker
func (b *Broker) Stop() error {
	b.cancel()
	return nil
}
