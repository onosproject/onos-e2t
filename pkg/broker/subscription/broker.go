// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	subctrl "github.com/onosproject/onos-e2t/pkg/controller/subscription"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("broker", "subscription")

// NewBroker creates a new subscription broker
func NewBroker(requests *subctrl.RequestJournal, streams *stream.Manager, channels *channel.Manager) *Broker {
	return &Broker{
		requests: requests,
		streams:  streams,
		channels: channels,
	}
}

// Broker is a subscription broker
type Broker struct {
	requests *subctrl.RequestJournal
	streams  *stream.Manager
	channels *channel.Manager
	cancel   context.CancelFunc
}

// Start starts the broker
func (b *Broker) Start() error {
	log.Infof("Starting Subscription Broker")
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
	log.Infof("Received Channel %s", channel.ID())
	dispatcher, err := newDispatcher(b.requests, channel, b.streams)
	if err != nil {
		log.Errorf("Failed to create dispatcher for Channel %s: %v", channel.ID(), err)
	} else {
		go func() {
			<-channel.Context().Done()
			err := dispatcher.Close()
			if err != nil {
				log.Errorf("Failed to close dispatcher for Channel %s: %v", channel.ID(), err)
			}
		}()
	}
}

// Stop stops the broker
func (b *Broker) Stop() error {
	b.cancel()
	return nil
}
