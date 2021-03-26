// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package broker

import (
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"sync"
)

var log = logging.GetLogger("broker", "stream")

// NewStreamBroker creates a new subscription broker
func NewStreamBroker() StreamBroker {
	return &streamBroker{
		subs:    make(map[SubscriptionID]Stream),
		streams: make(map[StreamID]Stream),
	}
}

// StreamBroker is a subscription stream broker
type StreamBroker interface {
	io.Closer

	// OpenStream opens a subscription stream
	OpenStream(id SubscriptionID) (StreamReader, error)

	// CloseStream closes a subscription stream
	CloseStream(id SubscriptionID) (StreamReader, error)

	// GetStream gets a write stream by ID
	GetStream(id StreamID) (StreamWriter, error)
}

type streamBroker struct {
	subs     map[SubscriptionID]Stream
	streams  map[StreamID]Stream
	streamID StreamID
	mu       sync.RWMutex
}

func (b *streamBroker) OpenStream(id SubscriptionID) (StreamReader, error) {
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

func (b *streamBroker) CloseStream(id SubscriptionID) (StreamReader, error) {
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

var _ StreamBroker = &streamBroker{}
