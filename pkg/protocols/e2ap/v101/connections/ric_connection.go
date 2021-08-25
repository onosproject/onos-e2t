// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

import (
	"context"
	"io"
	"net"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/connection"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/v101/procedures"
	"github.com/onosproject/onos-e2t/pkg/utils/async"
)

// RICHandler is a function for wrapping an RICConn
type RICHandler func(channel RICConn) procedures.RICProcedures

// RICConn is a connection for an E2 node
type RICConn interface {
	connection.Conn
	procedures.E2NodeProcedures
	BaseConn() *connection.Connection
}

// NewRICConn creates a new E2 node channel
func NewRICConn(conn net.Conn, handler RICHandler, opts ...connection.Option) RICConn {
	parent := newThreadSafeConn(conn, opts...)
	ricConn := &ricConn{
		threadSafeConn: parent,
	}
	procs := handler(ricConn)
	ricConn.e2Setup = procedures.NewE2SetupProcedure(parent.send, procs)
	ricConn.e2ConnectionUpdate = procedures.NewE2ConnectionUpdateInitiator(parent.send)
	ricConn.ricControl = procedures.NewRICControlInitiator(parent.send)
	ricConn.ricIndication = procedures.NewRICIndicationProcedure(parent.send, procs)
	ricConn.ricSubscription = procedures.NewRICSubscriptionInitiator(parent.send)
	ricConn.ricSubscriptionDelete = procedures.NewRICSubscriptionDeleteInitiator(parent.send)
	ricConn.open()
	return ricConn
}

// ricConn is an E2 node connection
type ricConn struct {
	*threadSafeConn
	e2Setup               *procedures.E2SetupProcedure
	e2ConnectionUpdate    *procedures.E2ConnectionUpdateInitiator
	ricControl            *procedures.RICControlInitiator
	ricIndication         *procedures.RICIndicationProcedure
	ricSubscription       *procedures.RICSubscriptionInitiator
	ricSubscriptionDelete *procedures.RICSubscriptionDeleteInitiator
	ricIndicationCh       chan e2appdudescriptions.E2ApPdu
}

func (c *ricConn) BaseConn() *connection.Connection {
	return c.baseConn
}

func (c *ricConn) open() {
	c.ricIndicationCh = make(chan e2appdudescriptions.E2ApPdu)
	go c.recvPDUs()
	go c.recvIndications()
}

func (c *ricConn) recvPDUs() {
	for {
		pdu, err := c.recv()
		if err == io.EOF {
			log.Warn(err)
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

func (c *ricConn) recvPDU(pdu *e2appdudescriptions.E2ApPdu) {
	if c.e2Setup.Matches(pdu) {
		go c.e2Setup.Handle(pdu)
	} else if c.e2ConnectionUpdate.Matches(pdu) {
		go c.e2ConnectionUpdate.Handle(pdu)
	} else if c.ricControl.Matches(pdu) {
		go c.ricControl.Handle(pdu)
	} else if c.ricIndication.Matches(pdu) {
		c.ricIndicationCh <- *pdu
	} else if c.ricSubscription.Matches(pdu) {
		go c.ricSubscription.Handle(pdu)
	} else if c.ricSubscriptionDelete.Matches(pdu) {
		go c.ricSubscriptionDelete.Handle(pdu)
	} else {
		log.Errorf("Unsupported E2AP message: %+v", pdu)
	}
}

func (c *ricConn) recvIndications() {
	for pdu := range c.ricIndicationCh {
		c.recvIndication(pdu)
	}
}

func (c *ricConn) recvIndication(pdu e2appdudescriptions.E2ApPdu) {
	c.ricIndication.Handle(&pdu)
}

func (c *ricConn) E2ConnectionUpdate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (response *e2appducontents.E2ConnectionUpdateAcknowledge, failure *e2appducontents.E2ConnectionUpdateFailure, err error) {
	return c.e2ConnectionUpdate.Initiate(ctx, request)
}

func (c *ricConn) RICControl(ctx context.Context, request *e2appducontents.RiccontrolRequest) (response *e2appducontents.RiccontrolAcknowledge, failure *e2appducontents.RiccontrolFailure, err error) {
	return c.ricControl.Initiate(ctx, request)
}

func (c *ricConn) RICSubscription(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (response *e2appducontents.RicsubscriptionResponse, failure *e2appducontents.RicsubscriptionFailure, err error) {
	return c.ricSubscription.Initiate(ctx, request)
}

func (c *ricConn) RICSubscriptionDelete(ctx context.Context, request *e2appducontents.RicsubscriptionDeleteRequest) (response *e2appducontents.RicsubscriptionDeleteResponse, failure *e2appducontents.RicsubscriptionDeleteFailure, err error) {
	return c.ricSubscriptionDelete.Initiate(ctx, request)
}

func (c *ricConn) Close() error {
	procedures := []procedures.ElementaryProcedure{
		c.e2Setup,
		c.e2ConnectionUpdate,
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
	return nil
}

var _ RICConn = &ricConn{}
