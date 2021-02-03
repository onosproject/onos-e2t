// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"io"
	"sync"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
)

// newListener creates a new subscription listener
func newListener(id ListenerID, sub subapi.Subscription, streams *stream.Manager) (*Listener, error) {
	listener := &Listener{
		ID:  id,
		sub: sub,
	}
	if err := listener.open(streams); err != nil {
		return nil, err
	}
	return listener, nil
}

// ListenerID is a subscription listener identifier
type ListenerID int32

// Listener is a subscription listener
type Listener struct {
	ID      ListenerID
	sub     subapi.Subscription
	streams []stream.Stream
	mu      sync.RWMutex
	cancel  context.CancelFunc
}

// open opens the listener
func (l *Listener) open(streams *stream.Manager) error {
	ctx, cancel := context.WithCancel(context.Background())
	l.cancel = cancel

	streamCh := make(chan stream.Stream)
	if err := streams.Watch(ctx, streamCh); err != nil {
		return err
	}
	go l.processStreams(streamCh)
	return nil
}

// processStreams processes stream events
func (l *Listener) processStreams(ch <-chan stream.Stream) {
	for stream := range ch {
		l.processStream(stream)
	}
}

// processStream processes a stream event
func (l *Listener) processStream(s stream.Stream) {
	if subapi.ID(s.Metadata().SubscriptionID) == l.sub.ID {
		log.Infof("Opened new stream %d for listener %d", s.ID(), l.ID)
		l.mu.Lock()
		l.streams = append(l.streams, s)
		l.mu.Unlock()
		go func() {
			<-s.Context().Done()
			log.Infof("Closed stream %d for listener %d", s.ID(), l.ID)
			l.mu.Lock()
			if len(l.streams) > 0 {
				streams := make([]stream.Stream, 0, len(l.streams)-1)
				for _, s2 := range l.streams {
					if s2.ID() != s.ID() {
						streams = append(streams, s2)
					}
				}
				l.streams = streams
			}
			l.mu.Unlock()
		}()
	}
}

// Notify notifies the listener of the given indication
func (l *Listener) Notify(indication e2appducontents.Ricindication) error {
	l.mu.RLock()
	streams := l.streams
	l.mu.RUnlock()

	id := stream.MessageID(indication.ProtocolIes.E2ApProtocolIes27.Value.Value)
	log.Infof("Notifying indication %d for listener %d", id, l.ID)
	for _, s := range streams {
		err := s.Send(stream.Value(id, indication))
		if err != nil {
			log.Errorf("Failed to indicate %d for listener %d: %s", id, l.ID, err)
			return err
		}
	}
	return nil
}

// Close closes the listener
func (l *Listener) Close() error {
	l.cancel()
	return nil
}

var _ io.Closer = &Listener{}
