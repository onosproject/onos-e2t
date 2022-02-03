// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"context"
	"io"
	"net"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/procedures"
	"github.com/onosproject/onos-e2t/pkg/utils/async"
)

// ClientConn is a connection for an E2AP client
type ClientConn interface {
	Conn
	procedures.RICProcedures
}

// NewClientConn creates a new client connection
func NewClientConn(c net.Conn, handler ClientHandler, opts ...Option) ClientConn {
	parent := newThreadSafeConn(c, opts...)
	cc := &clientConn{
		threadSafeConn: parent,
	}
	procs := handler(cc)
	cc.e2Setup = procedures.NewE2SetupInitiator(parent.send)
	cc.e2ConnectionUpdate = procedures.NewE2ConnectionUpdateProcedure(parent.send, procs)
	cc.e2ConfigurationUpdate = procedures.NewE2ConfigurationUpdateInitiator(parent.send)
	cc.ricControl = procedures.NewRICControlProcedure(parent.send, procs)
	cc.ricIndication = procedures.NewRICIndicationInitiator(parent.send)
	cc.ricSubscription = procedures.NewRICSubscriptionProcedure(parent.send, procs)
	cc.ricSubscriptionDelete = procedures.NewRICSubscriptionDeleteProcedure(parent.send, procs)
	cc.open()
	return cc
}

// clientConn is an E2 node client connection
type clientConn struct {
	*threadSafeConn
	e2Setup               *procedures.E2SetupInitiator
	e2ConfigurationUpdate *procedures.E2ConfigurationUpdateInitiator
	e2ConnectionUpdate    *procedures.E2ConnectionUpdateProcedure
	ricControl            *procedures.RICControlProcedure
	ricIndication         *procedures.RICIndicationInitiator
	ricSubscription       *procedures.RICSubscriptionProcedure
	ricSubscriptionDelete *procedures.RICSubscriptionDeleteProcedure
}

func (c *clientConn) open() {
	go c.recvPDUs()
}

func (c *clientConn) recvPDUs() {
	for {
		pdu, err := c.recv()
		if err == io.EOF {
			c.Close()
			return
		}
		if err != nil {
			log.Error(err)
		} else {
			c.recvPDU(pdu)
		}
	}
}

func (c *clientConn) recvPDU(pdu *e2appdudescriptions.E2ApPdu) {
	if c.e2Setup.Matches(pdu) {
		go c.e2Setup.Handle(pdu)
	} else if c.e2ConfigurationUpdate.Matches(pdu) {
		go c.e2ConfigurationUpdate.Handle(pdu)
	} else if c.e2ConnectionUpdate.Matches(pdu) {
		go c.e2ConnectionUpdate.Handle(pdu)
	} else if c.ricControl.Matches(pdu) {
		go c.ricControl.Handle(pdu)
	} else if c.ricIndication.Matches(pdu) {
		c.ricIndication.Handle(pdu)
	} else if c.ricSubscription.Matches(pdu) {
		go c.ricSubscription.Handle(pdu)
	} else if c.ricSubscriptionDelete.Matches(pdu) {
		go c.ricSubscriptionDelete.Handle(pdu)
	} else {
		log.Errorf("Unsupported E2AP message: %+v", pdu)
	}
}

func (c *clientConn) E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error) {
	return c.e2ConfigurationUpdate.Initiate(ctx, request)
}

func (c *clientConn) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (response *e2appducontents.E2SetupResponse, failure *e2appducontents.E2SetupFailure, err error) {
	return c.e2Setup.Initiate(ctx, request)
}

func (c *clientConn) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) (err error) {
	return c.ricIndication.Initiate(ctx, request)
}

func (c *clientConn) Close() error {
	procedures := []procedures.ElementaryProcedure{
		c.e2Setup,
		c.e2ConnectionUpdate,
		c.e2ConfigurationUpdate,
		c.ricControl,
		c.ricIndication,
		c.ricSubscription,
		c.ricSubscriptionDelete,
	}
	err := async.Apply(len(procedures), func(i int) error {
		return procedures[i].Close()
	})
	if err != nil {
		return err
	}
	return c.threadSafeConn.Close()
}

var _ ClientConn = &clientConn{}
