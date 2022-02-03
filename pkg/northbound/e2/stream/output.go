// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import (
	"context"
	"github.com/google/uuid"
	"sync"
	"sync/atomic"
)

type Output interface {
	Open(ctx context.Context) Stream
	Streams() []Stream
}

func newOutput(c *channel) *channelOutput {
	return &channelOutput{
		channel: c,
		streams: make(map[uuid.UUID]Stream),
		readyCh: make(chan struct{}),
		doneCh:  make(chan struct{}),
	}
}

type channelOutput struct {
	*channel
	streams map[uuid.UUID]Stream
	readyCh chan struct{}
	ready   bool
	doneCh  chan struct{}
	done    bool
	err     atomic.Value
	mu      sync.RWMutex
}

func (c *channelOutput) open() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ready {
		return
	}
	defer close(c.readyCh)
	c.ready = true
	go c.notify()
}

func (c *channelOutput) close(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.done {
		return
	}
	if err != nil {
		c.err.Store(err)
	}
	defer close(c.doneCh)
	defer c.manager.close(c)
	c.done = true
	go c.notify()
}

func (c *channelOutput) Ready() <-chan struct{} {
	return c.readyCh
}

func (c *channelOutput) Done() <-chan struct{} {
	return c.doneCh
}

func (c *channelOutput) Err() error {
	err := c.err.Load()
	if err != nil {
		return err.(error)
	}
	return nil
}

func (c *channelOutput) Open(ctx context.Context) Stream {
	stream := newStream(c)
	c.mu.Lock()
	id := uuid.New()
	c.streams[id] = stream
	c.mu.Unlock()
	go c.notify()

	go func() {
		<-ctx.Done()
		c.mu.Lock()
		delete(c.streams, id)
		c.mu.Unlock()
		c.notify()
	}()
	return stream
}

func (c *channelOutput) Streams() []Stream {
	c.mu.RLock()
	defer c.mu.RUnlock()
	streams := make([]Stream, 0, len(c.streams))
	for _, stream := range c.streams {
		streams = append(streams, stream)
	}
	return streams
}
