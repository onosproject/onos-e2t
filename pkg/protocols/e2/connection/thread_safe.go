// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connection

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"io"
)

// NewThreadSafeConnection wraps the given connection in a thread safe implementation
func NewThreadSafeConnection(conn Connection) Connection {
	tsc := &threadSafeConnection{
		conn:   conn,
		sendCh: make(chan e2appdudescriptions.E2ApPdu),
		recvCh: make(chan e2appdudescriptions.E2ApPdu),
	}
	tsc.open()
	return conn
}

type threadSafeConnection struct {
	conn   Connection
	sendCh chan e2appdudescriptions.E2ApPdu
	recvCh chan e2appdudescriptions.E2ApPdu
}

func (c *threadSafeConnection) open() {
	go c.send()
	go c.recv()
}

func (c *threadSafeConnection) Send(msg *e2appdudescriptions.E2ApPdu) error {
	c.sendCh <- *msg
	return nil
}

func (c *threadSafeConnection) send() {
	for msg := range c.sendCh {
		err := c.Send(&msg)
		if err == io.EOF {
			c.Close()
		}
		if err != nil {
			log.Error(err)
		}
	}
}

func (c *threadSafeConnection) Recv() (*e2appdudescriptions.E2ApPdu, error) {
	msg, ok := <-c.recvCh
	if !ok {
		return nil, io.EOF
	}
	return &msg, nil
}

func (c *threadSafeConnection) recv() {
	for {
		msg, err := c.Recv()
		if err == io.EOF {
			c.Close()
			return
		}
		if err != nil {
			log.Error(err)
		} else {
			c.recvCh <- *msg
		}
	}
}

func (c *threadSafeConnection) Close() error {
	close(c.sendCh)
	close(c.recvCh)
	return c.conn.Close()
}

var _ Connection = &threadSafeConnection{}
