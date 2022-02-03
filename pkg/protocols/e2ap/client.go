// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"context"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/procedures"

	sctp "github.com/onosproject/onos-lib-go/pkg/sctp"
)

// ClientHandler is a client handler function
type ClientHandler func(conn ClientConn) ClientInterface

// ClientInterface is an interface for E2 client procedures
type ClientInterface procedures.E2NodeProcedures

// Connect connects to the given address
func Connect(ctx context.Context, address string, handler ClientHandler) (ClientConn, error) {
	addr, err := addressing.ResolveAddress(types.Sctp4, address)
	if err != nil {
		return nil, err
	}
	c, err := sctp.DialSCTP(addr,
		sctp.WithAddressFamily(addr.AddressFamily),
		sctp.WithNonBlocking(false),
		sctp.WithMode(types.OneToOne),
		sctp.WithInitMsg(types.InitMsg{}))
	if err != nil {
		return nil, err
	}
	conn := NewClientConn(c, func(conn ClientConn) ClientInterface {
		return handler(conn)
	})
	return conn, nil
}
