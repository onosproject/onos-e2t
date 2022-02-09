// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"context"
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"io"
	"net"
	"sync"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

const defaultRecvBufSize = 1024 * 4

var log = logging.GetLogger("protocols", "e2")

// Options is E2 connection options
type Options struct {
	RecvBufferSize int
}

// Option is an E2 connection option
type Option func(*Options)

// WithRecvBuffer sets the connection receive buffer size
func WithRecvBuffer(size int) Option {
	return func(options *Options) {
		options.RecvBufferSize = size
	}
}

// Conn is the base interface for E2 connections
type Conn interface {
	io.Closer
	// Context returns the connection context
	Context() context.Context
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

// newThreadSafeConn creates a new thread safe connection
func newThreadSafeConn(c net.Conn, opts ...Option) *threadSafeConn {
	options := Options{
		RecvBufferSize: defaultRecvBufSize,
	}
	for _, opt := range opts {
		opt(&options)
	}
	ctx, cancel := context.WithCancel(context.Background())
	conn := &threadSafeConn{
		conn:    c,
		sendCh:  make(chan asyncMessage, 1000),
		recvCh:  make(chan e2appdudescriptions.E2ApPdu),
		options: options,
		ctx:     ctx,
		cancel:  cancel,
	}
	conn.open()
	return conn
}

// threadSafeConn is a thread-safe Conn implementation
type threadSafeConn struct {
	conn    net.Conn
	sendCh  chan asyncMessage
	recvCh  chan e2appdudescriptions.E2ApPdu
	options Options
	ctx     context.Context
	cancel  context.CancelFunc
	closed  bool
	mu      sync.RWMutex
}

func (c *threadSafeConn) Context() context.Context {
	return c.ctx
}

func (c *threadSafeConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *threadSafeConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *threadSafeConn) open() {
	go c.processSends()
	go c.processRecvs()
}

// send sends a message on the connection
func (c *threadSafeConn) send(msg *e2appdudescriptions.E2ApPdu) error {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return c.ctx.Err()
	}

	errCh := make(chan error, 1)
	c.sendCh <- asyncMessage{
		msg:   msg,
		errCh: errCh,
	}
	c.mu.RUnlock()
	return <-errCh
}

// processSends processes the send channel
func (c *threadSafeConn) processSends() {
	for msg := range c.sendCh {
		err := c.processSend(msg.msg)
		if err == io.EOF {
			log.Warn(err)
			c.Close()
		} else if err != nil {
			msg.errCh <- err
		}
		close(msg.errCh)
	}
}

// processSend processes a send
func (c *threadSafeConn) processSend(msg *e2appdudescriptions.E2ApPdu) error {
	log.Debugf("Obtained message to encode is:\n%v", msg)
	bytes, err := encoder.PerEncodeE2ApPdu(msg)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Debugf("Encoded message is:\n%v", hex.Dump(bytes))
	_, err = c.conn.Write(bytes)
	return err
}

// recv receives a message on the connection
func (c *threadSafeConn) recv() (*e2appdudescriptions.E2ApPdu, error) {
	msg, ok := <-c.recvCh
	if !ok {
		log.Warn("no more messages to receive")
		return nil, io.EOF
	}
	return &msg, nil
}

// processRecvs processes the receive channel
func (c *threadSafeConn) processRecvs() {
	buf := make([]byte, c.options.RecvBufferSize)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			log.Warn(err)
			c.Close()
			return
		}

		err = c.processRecv(buf[:n])
		if err != nil {
			log.Error(err)
		}
	}
}

// processRecvs processes the receive channel
func (c *threadSafeConn) processRecv(bytes []byte) error {
	log.Debugf("Obtained bytes to decode are\n%v", hex.Dump(bytes))
	msg, err := encoder.PerDecodeE2ApPdu(bytes)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Debugf("Decoded message is:\n%v", msg)
	c.recvCh <- *msg
	return nil
}

func (c *threadSafeConn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.closed {
		close(c.sendCh)
		close(c.recvCh)
		c.cancel()
		c.closed = true
	}
	return c.conn.Close()
}

type asyncMessage struct {
	msg   *e2appdudescriptions.E2ApPdu
	errCh chan error
}
