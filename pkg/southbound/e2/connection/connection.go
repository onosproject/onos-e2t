// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connection

import (
	"context"
	"errors"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/orane2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io"
	"net"
)

var log = logging.GetLogger("southbound", "e2", "connection")

const readBufSize = 4096

// TODO: Change the RIC ID to something appropriate
const ricID = 0x01 // µONOS RIC is #1!!!

// ID is a connection identifier
type ID uint64

// PlmnID is a connection identifier
type PlmnID string

// NewConnection creates a new connection
func NewConnection(conn net.Conn) (*Connection, error) {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Connection{
		conn:   conn,
		ctx:    ctx,
		cancel: cancel,
	}
	if err := c.setup(); err != nil {
		return nil, err
	}
	return c, nil
}

// Connection is an E2 connection
type Connection struct {
	ID     ID
	PlmnID PlmnID
	conn   net.Conn
	ctx    context.Context
	cancel context.CancelFunc
}

// LocalAddr returns the local connection address
func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote connection address
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// setup sets up the connection
func (c *Connection) setup() error {
	buf := make([]byte, readBufSize)
	n, err := c.conn.Read(buf)
	if err != nil {
		defer c.cancel()
		return err
	}

	// Decode the E2 request in XER encoding
	// TODO: E2AP messages are supposed to use PER encoding
	e2PDUReqBytes := buf[:n]
	e2PDUReq, err := orane2.XerDecodeE2apPdu(e2PDUReqBytes)
	if err != nil {
		defer c.cancel()
		return err
	}

	// Verify this is a setup request
	e2SetupReq := e2PDUReq.GetInitiatingMessage().GetProcedureCode().GetE2Setup()
	if e2SetupReq == nil {
		defer c.cancel()
		return errors.New("unexpected message type")
	}

	// Verify an E2 node ID is provided
	e2NodeID := e2SetupReq.InitiatingMessage.ProtocolIes.E2ApProtocolIes3.Value.GlobalE2NodeId
	globalE2NodeID, ok := e2NodeID.(*e2apies.GlobalE2NodeId_GNb)
	if !ok {
		defer c.cancel()
		return errors.New("unexpected message format")
	}

	// Verify a gNB ID is provided
	gnbID, ok := globalE2NodeID.GNb.GlobalGNbId.GnbId.GnbIdChoice.(*e2apies.GnbIdChoice_GnbId)
	if !ok {
		defer c.cancel()
		return errors.New("unexpected message format")
	}

	// Create a connection ID from the gNB ID and plmn ID
	connID := gnbID.GnbId.Value
	plmnID := globalE2NodeID.GNb.GlobalGNbId.PlmnId.Value
	for i := range plmnID {
		b := plmnID[(len(plmnID)-i)-1]
		connID = connID & (uint64(b) << (gnbID.GnbId.Len + uint32(i*8)))
	}

	// Set the connection ID
	c.ID = ID(connID)
	c.PlmnID = PlmnID(plmnID)

	// Create an E2 setup response
	e2SetupResp := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte(plmnID),
			},
			RicId: &e2ap_commondatatypes.BitString{
				Value: ricID,
				Len:   20,
			},
		},
	}

	// Create an E2 response
	e2PDUResp := &e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						SuccessfulOutcome: &e2appducontents.E2SetupResponse{
							ProtocolIes: &e2appducontents.E2SetupResponseIes{
								E2ApProtocolIes4: &e2SetupResp,
							},
						},
					},
				},
			},
		},
	}

	// Encode the setup response in XER
	// TODO: E2AP messages are supposed to use PER encoding
	e2PDURespBytes, err := orane2.XerEncodeE2apPdu(e2PDUResp)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(e2PDURespBytes)
	if err != nil {
		defer c.cancel()
		return err
	}
	return nil
}

// Send sends a message on the connection
func (c *Connection) Send(msg *e2appdudescriptions.E2ApPdu) error {
	// TODO: This encodes all messages in PER encoding
	bytes, err := orane2.XerEncodeE2apPdu(msg)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(bytes)
	if err == io.EOF {
		c.cancel()
	}
	if err != nil {
		return err
	}
	return nil
}

// Recv receives a message on the connection
func (c *Connection) Recv() (*e2appdudescriptions.E2ApPdu, error) {
	buf := make([]byte, readBufSize)
	n, err := c.conn.Read(buf)
	if err == io.EOF {
		c.cancel()
	}
	if err != nil {
		return nil, err
	}

	bytes := buf[:n]

	// TODO: This decodes all messages in PER encoding
	msg, err := orane2.PerDecodeE2apPdu(bytes)
	if err == io.EOF {
		c.cancel()
	}
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// Context returns the connection context
func (c *Connection) Context() context.Context {
	return c.ctx
}

// Close closes the connection
func (c *Connection) Close() error {
	err := c.Close()
	c.cancel()
	return err
}

var _ io.Closer = &Connection{}
