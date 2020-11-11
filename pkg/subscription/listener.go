// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"context"
	"github.com/gogo/protobuf/proto"
	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"io"
	"sync"
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
	ID     ListenerID
	sub    subapi.Subscription
	stream stream.Stream
	mu     sync.RWMutex
	cancel context.CancelFunc
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
func (l *Listener) processStream(stream stream.Stream) {
	l.mu.Lock()
	defer l.mu.Unlock()
	streamAppID := subapi.AppID(stream.ID())
	if l.sub.AppID == streamAppID {
		l.stream = stream
	}
}

// Notify notifies the listener of the given indication
func (l *Listener) Notify(indication *e2appdudescriptions.E2ApPdu) error {
	l.mu.RLock()
	s := l.stream
	l.mu.RUnlock()

	if s == nil {
		return nil
	}

	bytes, err := proto.Marshal(indication)
	if err != nil {
		return err
	}
	id := stream.MessageID(indication.GetInitiatingMessage().ProcedureCode.RicIndication.InitiatingMessage.ProtocolIes.E2ApProtocolIes27.Value.Value)
	return s.Send(stream.Value(id, bytes))
}

// Close closes the listener
func (l *Listener) Close() error {
	l.cancel()
	return nil
}

var _ io.Closer = &Listener{}
