// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channels

import (
	"context"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2/procedures"
	"github.com/onosproject/onos-e2t/pkg/utils/async"
	"io"
	"net"
)

// RICChannel is a channel for an E2 node
type RICChannel interface {
	Channel
	procedures.E2NodeProcedures
}

// NewRICChannel creates a new E2 node channel
func NewRICChannel(conn net.Conn, procs procedures.RICProcedures, opts ...Option) RICChannel {
	parent := newThreadSafeChannel(conn, opts...)
	channel := &ricChannel{
		threadSafeChannel:     parent,
		e2Setup:               procedures.NewE2SetupProcedure(parent.send, procs),
		ricControl:            procedures.NewRICControlInitiator(parent.send),
		ricIndication:         procedures.NewRICIndicationProcedure(parent.send, procs),
		ricSubscription:       procedures.NewRICSubscriptionInitiator(parent.send),
		ricSubscriptionDelete: procedures.NewRICSubscriptionDeleteInitiator(parent.send),
	}
	channel.open()
	return channel
}

// ricChannel is an E2 node channel
type ricChannel struct {
	*threadSafeChannel
	e2Setup               *procedures.E2SetupProcedure
	ricControl            *procedures.RICControlInitiator
	ricIndication         *procedures.RICIndicationProcedure
	ricSubscription       *procedures.RICSubscriptionInitiator
	ricSubscriptionDelete *procedures.RICSubscriptionDeleteInitiator
}

func (c *ricChannel) open() {
	go c.recvPDUs()
}

func (c *ricChannel) recvPDUs() {
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

func (c *ricChannel) recvPDU(pdu *e2appdudescriptions.E2ApPdu) {
	if c.e2Setup.Matches(pdu) {
		c.e2Setup.Handle(pdu)
	} else if c.ricControl.Matches(pdu) {
		c.ricControl.Handle(pdu)
	} else if c.ricIndication.Matches(pdu) {
		c.ricIndication.Handle(pdu)
	} else if c.ricSubscription.Matches(pdu) {
		c.ricSubscription.Handle(pdu)
	} else if c.ricSubscriptionDelete.Matches(pdu) {
		c.ricSubscriptionDelete.Handle(pdu)
	} else {
		log.Errorf("Unsupported E2AP message: %+v", pdu)
	}
}

func (c *ricChannel) RICControl(ctx context.Context, request *e2appducontents.RiccontrolRequest) (response *e2appducontents.RiccontrolAcknowledge, failure *e2appducontents.RiccontrolFailure, err error) {
	return c.ricControl.Initiate(ctx, request)
}

func (c *ricChannel) RICSubscription(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (response *e2appducontents.RicsubscriptionResponse, failure *e2appducontents.RicsubscriptionFailure, err error) {
	return c.ricSubscription.Initiate(ctx, request)
}

func (c *ricChannel) RICSubscriptionDelete(ctx context.Context, request *e2appducontents.RicsubscriptionDeleteRequest) (response *e2appducontents.RicsubscriptionDeleteResponse, failure *e2appducontents.RicsubscriptionDeleteFailure, err error) {
	return c.ricSubscriptionDelete.Initiate(ctx, request)
}

func (c *ricChannel) Close() error {
	defer c.conn.Close()
	procedures := []procedures.ElementaryProcedure{
		c.e2Setup,
		c.ricControl,
		c.ricIndication,
		c.ricSubscription,
		c.ricSubscriptionDelete,
	}
	return async.Apply(len(procedures), func(i int) error {
		return procedures[i].Close()
	})
}

var _ RICChannel = &ricChannel{}
