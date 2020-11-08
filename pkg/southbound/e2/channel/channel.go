// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"io"
	"net"
	"sync"
)

const readBufSize = 4096
const channelBufSize = 100

// newChannel creates a new connection
func newChannel(ctx context.Context, conn net.Conn, meta Metadata) Channel {
	ctx, cancel := context.WithCancel(ctx)
	channel := &netChannel{
		meta:      meta,
		conn:      conn,
		receivers: make(map[ReceiverID]Receiver),
		ctx:       ctx,
		cancel:    cancel,
	}
	channel.open()
	return channel
}

// Channel is an E2 channel
type Channel interface {
	io.Closer

	// ID returns the channel identifier
	ID() ID
	// Metadata returns the channel metadata
	Metadata() Metadata
	// Context returns the connection context
	Context() context.Context
	// LocalAddr returns the local connection address
	LocalAddr() net.Addr
	// RemoteAddr returns the remote connection address
	RemoteAddr() net.Addr
	// Send sends a message
	Send(message *e2appdudescriptions.E2ApPdu, codec codec.Codec) error
	// SendRecv sends a request-response message
	SendRecv(message *e2appdudescriptions.E2ApPdu, filter filter.Filter, codec codec.Codec) (*e2appdudescriptions.E2ApPdu, error)
	// Recv returns the receive channel
	Recv(filter filter.Filter, codec codec.Codec) <-chan *e2appdudescriptions.E2ApPdu
}

// netChannel is an E2 channel
type netChannel struct {
	meta      Metadata
	conn      net.Conn
	receivers map[ReceiverID]Receiver
	mu        sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
}

func (c *netChannel) open() {
	go c.recvMessages()
}

func (c *netChannel) ID() ID {
	return c.meta.ID
}

func (c *netChannel) Metadata() Metadata {
	return c.meta
}

func (c *netChannel) Context() context.Context {
	return c.ctx
}

func (c *netChannel) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *netChannel) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *netChannel) recvMessages() {
	for {
		bytes, err := c.recv()
		if err != nil {
			log.Error(err)
		} else {
			c.mu.RLock()
			receivers := c.receivers
			c.mu.RUnlock()

			for id, receiver := range receivers {
				response, err := receiver.Decode(bytes)
				if err != nil {
					continue
				}
				if receiver.Match(response) {
					err := receiver.Receive(response)
					if err != nil {
						log.Error(err)
					}
					if receiver.Done() {
						c.mu.Lock()
						delete(c.receivers, id)
						c.mu.Unlock()
					}
					break
				}
			}
		}
	}
}

func (c *netChannel) Send(request *e2appdudescriptions.E2ApPdu, codec codec.Codec) error {
	bytes, err := codec.Encode(request)
	if err != nil {
		return err
	}
	return c.send(bytes)
}

func (c *netChannel) SendRecv(request *e2appdudescriptions.E2ApPdu, filter filter.Filter, codec codec.Codec) (*e2appdudescriptions.E2ApPdu, error) {
	bytes, err := codec.Encode(request)
	if err != nil {
		return nil, err
	}

	ch := make(chan *e2appdudescriptions.E2ApPdu, 1)
	receiver := newUnaryReceiver(ch, filter, codec)
	c.mu.Lock()
	c.receivers[receiver.ID()] = receiver
	c.mu.Unlock()

	defer func() {
		receiver.Close()
		c.mu.Lock()
		delete(c.receivers, receiver.ID())
		c.mu.Unlock()
	}()

	err = c.send(bytes)
	if err != nil {
		return nil, err
	}

	response := <-ch
	return response, nil
}

func (c *netChannel) Recv(filter filter.Filter, codec codec.Codec) <-chan *e2appdudescriptions.E2ApPdu {
	ch := make(chan *e2appdudescriptions.E2ApPdu, channelBufSize)
	receiver := newStreamReceiver(ch, filter, codec)
	c.mu.Lock()
	c.receivers[receiver.ID()] = receiver
	c.mu.Unlock()
	return ch
}

// send sends a message on the channel
func (c *netChannel) send(payload []byte) error {
	_, err := c.conn.Write(payload)
	if err == io.EOF {
		c.cancel()
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

// recv receives a message on the channel
func (c *netChannel) recv() ([]byte, error) {
	buf := make([]byte, readBufSize)
	n, err := c.conn.Read(buf)
	if err == io.EOF {
		c.cancel()
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

// Close closes the connection
func (c *netChannel) Close() error {
	err := c.conn.Close()
	c.cancel()
	return err
}

var _ Channel = &netChannel{}
