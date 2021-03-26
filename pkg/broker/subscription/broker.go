// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"sync"
)

var log = logging.GetLogger("broker", "subscription")

// NewBroker creates a new subscription stream broker
func NewBroker() Broker {
	return &streamBroker{
		subs:    make(map[subscription.ID]Stream),
		streams: make(map[StreamID]Stream),
	}
}

// Broker is a subscription stream broker
// The Broker is responsible for managing Streams for propagating indications from the southbound API
// to the northbound API.
type Broker interface {
	io.Closer

	// OpenStream opens a subscription Stream
	// If a stream already exists for the subscription, the existing stream will be returned.
	// If no stream exists, a new stream will be allocated with a unique StreamID.
	OpenStream(id subscription.ID) (StreamReader, error)

	// CloseStream closes a subscription Stream
	// The associated Stream will be closed gracefully: the reader will continue receiving pending indications
	// until the buffer is empty.
	CloseStream(id subscription.ID) (StreamReader, error)

	// GetStream gets a write stream by its StreamID
	// If no Stream exists for the given StreamID, a NotFound error will be returned.
	GetStream(id StreamID) (StreamWriter, error)
}

type streamBroker struct {
	subs     map[subscription.ID]Stream
	streams  map[StreamID]Stream
	streamID StreamID
	mu       sync.RWMutex
}

func (b *streamBroker) OpenStream(id subscription.ID) (StreamReader, error) {
	b.mu.RLock()
	stream, ok := b.subs[id]
	b.mu.RUnlock()
	if ok {
		return stream, nil
	}

	b.mu.Lock()
	defer b.mu.Unlock()
	stream, ok = b.subs[id]
	if ok {
		return stream, nil
	}

	b.streamID++
	streamID := b.streamID
	stream = newBufferedStream(id, streamID)
	b.subs[id] = stream
	b.streams[streamID] = stream
	log.Infof("Opened new stream %d for subscription '%s'", streamID, id)
	return stream, nil
}

func (b *streamBroker) CloseStream(id subscription.ID) (StreamReader, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	stream, ok := b.subs[id]
	if !ok {
		return nil, errors.NewNotFound("subscription '%s' not found", id)
	}
	delete(b.subs, stream.SubscriptionID())
	delete(b.streams, stream.StreamID())
	log.Infof("Closed stream %d for subscription '%s'", stream.StreamID(), id)
	return stream, stream.Close()
}

func (b *streamBroker) GetStream(id StreamID) (StreamWriter, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	stream, ok := b.streams[id]
	if !ok {
		return nil, errors.NewNotFound("stream %d not found", id)
	}
	return stream, nil
}

func (b *streamBroker) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	var err error
	for _, stream := range b.streams {
		if e := stream.Close(); e != nil {
			err = e
		}
	}
	return err
}

var _ Broker = &streamBroker{}
