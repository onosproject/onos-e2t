// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

import (
	"context"
	"io"
	"net"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/connection"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap201/asn1cgo"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("protocols", "e2", "v201")

// newThreadSafeConn creates a new thread safe connection
func newThreadSafeConn(conn net.Conn, opts ...connection.Option) *threadSafeConn {
	e2Connection := &threadSafeConn{
		sendCh: make(chan asyncMessage),
		recvCh: make(chan e2appdudescriptions.E2ApPdu),
	}
	baseConn := connection.NewBaseConnection(conn, opts...)
	e2Connection.baseConn = baseConn

	e2Connection.open()
	return e2Connection
}

// threadSafeCon is a thread-safe Connection implementation
type threadSafeConn struct {
	baseConn *connection.Connection
	sendCh   chan asyncMessage
	recvCh   chan e2appdudescriptions.E2ApPdu
}

func (c *threadSafeConn) Context() context.Context {
	return c.baseConn.Context()
}

func (c *threadSafeConn) LocalAddr() net.Addr {
	return c.baseConn.LocalAddr()
}

func (c *threadSafeConn) RemoteAddr() net.Addr {
	return c.baseConn.RemoteAddr()
}

func (c *threadSafeConn) open() {
	go c.processSends()
	go c.processRecvs()
}

// send sends a message on the connection
func (c *threadSafeConn) send(msg *e2appdudescriptions.E2ApPdu) error {
	errCh := make(chan error, 1)
	c.sendCh <- asyncMessage{
		msg:   msg,
		errCh: errCh,
	}
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
	bytes, err := asn1cgo.PerEncodeE2apPdu(msg)
	if err != nil {
		log.Warn(err)
		return err
	}
	_, err = c.baseConn.Conn().Write(bytes)
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
	buf := make([]byte, c.baseConn.Options().RecvBufferSize)
	for {
		n, err := c.baseConn.Conn().Read(buf)
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
	msg, err := asn1cgo.PerDecodeE2apPdu(bytes)
	if err != nil {
		log.Warn(err)
		return err
	}
	c.recvCh <- *msg
	return nil
}

func (c *threadSafeConn) Close() error {
	defer func() {
		if err := recover(); err != nil {
			log.Debug("recovering from panic:", err)
		}
	}()
	close(c.sendCh)
	close(c.recvCh)
	c.baseConn.Cancel()
	return c.baseConn.Conn().Close()
}

type asyncMessage struct {
	msg   *e2appdudescriptions.E2ApPdu
	errCh chan error
}

var _ connection.Conn = &threadSafeConn{}
