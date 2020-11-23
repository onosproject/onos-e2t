// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connection

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"net"
)

const defaultRecvBufSize = 1024 * 4

var log = logging.GetLogger("protocols", "e2")

// Connection is an E2 connection
type Connection interface {
	io.Closer

	// Send sends a message on the connection
	Send(msg *e2appdudescriptions.E2ApPdu) error
	// Recv receives a message on the connection
	Recv() (*e2appdudescriptions.E2ApPdu, error)
}

// Options is E2 connection options
type Options struct {
	SendBufferSize  int
	RecvBufferSize  int
	WriteBufferSize int
	ReadBufferSize  int
}

// Option is an E2 connection option
type Option func(*Options)

// WithSendBuffer sets the connection send buffer size
func WithSendBuffer(size int) Option {
	return func(options *Options) {
		options.SendBufferSize = size
	}
}

// WithRecvBuffer sets the connection receive buffer size
func WithRecvBuffer(size int) Option {
	return func(options *Options) {
		options.RecvBufferSize = size
	}
}

// New creates a new connection for the given net.Conn
func New(conn net.Conn, opts ...Option) Connection {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}
	return &netConnection{
		conn:    conn,
		options: options,
	}
}

// netConnection is a net.Conn based Connection implementation
type netConnection struct {
	conn    net.Conn
	options Options
}

func (c *netConnection) Send(msg *e2appdudescriptions.E2ApPdu) error {
	bytes, err := asn1cgo.PerEncodeE2apPdu(msg)
	if err != nil {
		return err
	}
	return c.write(bytes)
}

// write writes the given bytes to the connection
func (c *netConnection) write(payload []byte) error {
	_, err := c.conn.Write(payload)
	if err == io.EOF {
		c.Close()
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *netConnection) Recv() (*e2appdudescriptions.E2ApPdu, error) {
	bytes, err := c.read()
	if err != nil {
		return nil, err
	}
	return asn1cgo.PerDecodeE2apPdu(bytes)
}

// read reads bytes from the connection
func (c *netConnection) read() ([]byte, error) {
	var buf []byte
	if c.options.RecvBufferSize != 0 {
		buf = make([]byte, c.options.RecvBufferSize)
	} else {
		buf = make([]byte, defaultRecvBufSize)
	}
	n, err := c.conn.Read(buf)
	if err == io.EOF {
		c.Close()
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (c *netConnection) Close() error {
	return c.conn.Close()
}
