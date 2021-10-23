// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/subscription"
	chanstore "github.com/onosproject/onos-e2t/pkg/store/channel"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"sync"
)

type Manager interface {
	Get(channelID e2api.ChannelID) (Channel, error)
	Open(ctx context.Context, channel *e2api.Channel) (Channel, error)
	Watch(ctx context.Context, ch <-chan Channel) error
}

type channelManager struct {
	chans       chanstore.Store
	subs        subscription.Manager
	chanStreams map[e2api.ChannelID]Channel
	openChans   map[e2api.SubscriptionID]chan struct{}
	buffers     map[BufferID]Buffer
	bufferChans map[BufferID]map[e2api.ChannelID]bool
	subBuffers  map[e2api.SubscriptionID]map[BufferID]bool
	subStreams  map[e2api.SubscriptionID]subscription.Stream
	streamsMu   sync.RWMutex
	watchers    map[uuid.UUID]chan<- Channel
	watchersMu  sync.RWMutex
	cancel      context.CancelFunc
}

func (m *channelManager) open() error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	subsCh := make(chan subscription.Stream)
	if err := m.subs.Watch(ctx, subsCh); err != nil {
		return err
	}
	go func() {
		for subStream := range subsCh {
			go func(subStream subscription.Stream) {
				m.propagateStream(subStream)
			}(subStream)
		}
	}()
	return nil
}

func (m *channelManager) propagateStream(subStream subscription.Stream) {
	m.streamsMu.Lock()
	if ch, ok := m.openChans[subStream.Subscription().ID]; ok {
		close(ch)
		delete(m.openChans, subStream.Subscription().ID)
	}
	m.subStreams[subStream.Subscription().ID] = subStream
	m.streamsMu.Unlock()
	for ind := range subStream.Reader().Indications() {
		m.propagateIndication(subStream, ind)
	}
	m.streamsMu.Lock()
	delete(m.subStreams, subStream.Subscription().ID)
	m.streamsMu.Unlock()
}

func (m *channelManager) propagateIndication(subStream subscription.Stream, ind *e2appducontents.Ricindication) {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()
	buffers := m.subBuffers[subStream.Subscription().ID]
	for bufferID := range buffers {
		if buffer, ok := m.buffers[bufferID]; ok {
			buffer.In() <- ind
		}
	}
}

func (m *channelManager) Get(channelID e2api.ChannelID) (Channel, error) {
	m.streamsMu.RLock()
	defer m.streamsMu.RUnlock()
	channel, ok := m.chanStreams[channelID]
	if !ok {
		return nil, errors.NewNotFound("channel %s not found", channelID)
	}
	return channel, nil
}

func (m *channelManager) Open(ctx context.Context, channel *e2api.Channel) (Channel, error) {
	channel.Status.Phase = e2api.ChannelPhase_CHANNEL_OPEN
	channel.Status.State = e2api.ChannelState_CHANNEL_PENDING
	if err := m.chans.Create(ctx, channel); err != nil && !errors.IsAlreadyExists(err) {
		return nil, err
	}

	m.streamsMu.Lock()

	stream, ok := m.chanStreams[channel.ID]
	if ok {
		stream.Close(errors.NewUnavailable("stream closed"))
	}

	bufferID := BufferID(fmt.Sprintf("%s:%s:%s", channel.E2NodeID, channel.AppID, channel.TransactionID))
	buffer, ok := m.buffers[bufferID]
	if !ok {
		buffer = newChannelBuffer(bufferID)
		m.buffers[bufferID] = buffer
	}

	stream = newChannelStream(channel, buffer, m)
	m.chanStreams[channel.ID] = stream

	bufferChans, ok := m.bufferChans[bufferID]
	if !ok {
		bufferChans = make(map[e2api.ChannelID]bool)
		m.bufferChans[bufferID] = bufferChans
	}
	bufferChans[channel.ID] = true

	subBuffers, ok := m.subBuffers[channel.SubscriptionID]
	if !ok {
		subBuffers = make(map[BufferID]bool)
		m.subBuffers[channel.SubscriptionID] = subBuffers
	}
	subBuffers[bufferID] = true

	var openCh chan struct{}
	if _, ok = m.subStreams[channel.SubscriptionID]; ok {
		openCh = make(chan struct{})
		close(openCh)
	} else {
		openCh, ok = m.openChans[channel.SubscriptionID]
		if !ok {
			openCh = make(chan struct{})
			m.openChans[channel.SubscriptionID] = openCh
		}
	}

	m.streamsMu.Unlock()

	go m.notify(stream)

	select {
	case <-openCh:
		return stream, nil
	case err := <-stream.Done():
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (m *channelManager) close(stream Channel) {
	m.streamsMu.Lock()
	defer m.streamsMu.Unlock()

	delete(m.chanStreams, stream.Channel().ID)

	bufferID := BufferID(fmt.Sprintf("%s:%s:%s", stream.Channel().E2NodeID, stream.Channel().AppID, stream.Channel().TransactionID))
	bufferChans := m.bufferChans[bufferID]
	delete(bufferChans, stream.Channel().ID)
	if len(bufferChans) == 0 {
		if buffer, ok := m.buffers[bufferID]; ok {
			buffer.Close()
			delete(m.buffers, bufferID)
		}
		delete(m.bufferChans, bufferID)
		for _, subBuffers := range m.subBuffers {
			delete(subBuffers, bufferID)
		}
	}
	go m.notify(stream)
}

func (m *channelManager) Watch(ctx context.Context, ch chan<- Channel) error {
	m.watchersMu.Lock()
	id := uuid.New()
	m.watchers[id] = ch
	m.watchersMu.Unlock()

	m.streamsMu.RLock()
	streams := make([]Channel, 0, len(m.chanStreams))
	for _, stream := range m.chanStreams {
		streams = append(streams, stream)
	}
	m.streamsMu.RUnlock()

	go func() {
		for _, stream := range streams {
			ch <- stream
		}
		<-ctx.Done()
		m.watchersMu.Lock()
		delete(m.watchers, id)
		m.watchersMu.Unlock()
	}()
	return nil
}

func (m *channelManager) notify(stream Channel) {
	m.watchersMu.RLock()
	for _, watcher := range m.watchers {
		watcher <- stream
	}
	m.watchersMu.RUnlock()
}

func (m *channelManager) Close() error {
	m.cancel()
	return nil
}
