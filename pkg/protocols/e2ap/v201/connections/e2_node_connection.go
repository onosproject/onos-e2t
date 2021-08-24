// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

import (
	"context"
	"io"
	"net"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/connection"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/v201/procedures"
	"github.com/onosproject/onos-e2t/pkg/utils/async"
)

// E2NodeHandler is a function for wrapping an E2NodeChannel
type E2NodeHandler func(channel E2NodeConn) procedures.E2NodeProcedures

// E2NodeConn is a connection for an E2 node
type E2NodeConn interface {
	connection.Conn
	procedures.RICProcedures
}

// NewE2NodeConn creates a new E2 node connection
func NewE2NodeConn(conn net.Conn, handler E2NodeHandler, opts ...connection.Option) E2NodeConn {
	parent := newThreadSafeConn(conn, opts...)
	e2NodConn := &e2NodeConn{
		threadSafeConn: parent,
	}
	procs := handler(e2NodConn)
	e2NodConn.e2Setup = procedures.NewE2SetupInitiator(parent.send)
	e2NodConn.e2ConfigurationUpdate = procedures.NewE2ConfigurationUpdateInitiator(parent.send)
	e2NodConn.e2ConnectionUpdate = procedures.NewE2ConnectionUpdateProcedure(parent.send, procs)
	e2NodConn.ricControl = procedures.NewRICControlProcedure(parent.send, procs)
	e2NodConn.ricIndication = procedures.NewRICIndicationInitiator(parent.send)
	e2NodConn.ricSubscription = procedures.NewRICSubscriptionProcedure(parent.send, procs)
	e2NodConn.ricSubscriptionDelete = procedures.NewRICSubscriptionDeleteProcedure(parent.send, procs)
	e2NodConn.open()
	return e2NodConn
}

// e2NodeConn is an E2 node connection
type e2NodeConn struct {
	*threadSafeConn
	e2Setup               *procedures.E2SetupInitiator
	e2ConfigurationUpdate *procedures.E2ConfigurationUpdateInitiator
	e2ConnectionUpdate    *procedures.E2ConnectionUpdateProcedure
	ricControl            *procedures.RICControlProcedure
	ricIndication         *procedures.RICIndicationInitiator
	ricSubscription       *procedures.RICSubscriptionProcedure
	ricSubscriptionDelete *procedures.RICSubscriptionDeleteProcedure
}

func (c *e2NodeConn) open() {
	go c.recvPDUs()
}

func (c *e2NodeConn) recvPDUs() {
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

func (c *e2NodeConn) recvPDU(pdu *e2appdudescriptions.E2ApPdu) {
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

func (c *e2NodeConn) E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error) {
	return c.e2ConfigurationUpdate.Initiate(ctx, request)
}

func (c *e2NodeConn) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (response *e2appducontents.E2SetupResponse, failure *e2appducontents.E2SetupFailure, err error) {
	return c.e2Setup.Initiate(ctx, request)
}

func (c *e2NodeConn) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) (err error) {
	return c.ricIndication.Initiate(ctx, request)
}

func (c *e2NodeConn) Close() error {
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

var _ E2NodeConn = &e2NodeConn{}
