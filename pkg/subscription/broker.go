// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

// NewBroker creates a new subscription broker
func NewBroker(catalog *Catalog, streams *stream.Manager, channels *channel.Manager) *Broker {
	return &Broker{
		catalog:  catalog,
		streams:  streams,
		channels: channels,
		log:      logging.GetLogger("subscription", "broker"),
	}
}

// Broker is a subscription broker
type Broker struct {
	catalog  *Catalog
	streams  *stream.Manager
	channels *channel.Manager
	cancel   context.CancelFunc
	log      logging.Logger
}

// Start starts the broker
func (b *Broker) Start() error {
	b.log.Infof("Starting Subscription Broker")
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
	b.log.Infof("Received Channel %s", channel.ID())
	dispatcher, err := newDispatcher(b.catalog, channel, b.streams, b.log)
	if err != nil {
		b.log.Errorf("Failed to create dispatcher for Channel %s: %v", channel.ID(), err)
	} else {
		go func() {
			<-channel.Context().Done()
			err := dispatcher.Close()
			if err != nil {
				b.log.Errorf("Failed to close dispatcher for Channel %s: %v", channel.ID(), err)
			}
		}()
	}
}

// Stop stops the broker
func (b *Broker) Stop() error {
	b.cancel()
	return nil
}
