// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connection

import (
	"context"
	"io"
	"net"
)

const defaultRecvBufSize = 1024 * 4

// Options is connection options
type Options struct {
	RecvBufferSize int
}

// Option is an  connection option
type Option func(*Options)

// WithRecvBuffer sets the connection receive buffer size
func WithRecvBuffer(size int) Option {
	return func(options *Options) {
		options.RecvBufferSize = size
	}
}

// Conn connection interface
type Conn interface {
	io.Closer
	// Context returns the channel context
	Context() context.Context
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

// NewBaseConnection creates a base connection
func NewBaseConnection(conn net.Conn, opts ...Option) *Connection {
	ctx, cancel := context.WithCancel(context.Background())

	options := Options{
		RecvBufferSize: defaultRecvBufSize,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return &Connection{
		conn:    conn,
		options: options,
		ctx:     ctx,
		cancel:  cancel,
	}
}

// Connection base connection information
type Connection struct {
	conn    net.Conn
	options Options
	ctx     context.Context
	cancel  context.CancelFunc
}

// Conn returns connection
func (c *Connection) Conn() net.Conn {
	return c.conn
}

// Options return connection options
func (c *Connection) Options() Options {
	return c.options
}

// Context returns the connection context
func (c *Connection) Context() context.Context {
	return c.ctx
}

// LocalAddr returns connection local address
func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the connection remote address
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// Cancel cancel connection context
func (c *Connection) Cancel() {
	c.cancel()
}
