// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap101/channels"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap101/procedures"

	sctp "github.com/onosproject/onos-lib-go/pkg/sctp"
)

// ClientHandler is a client handler function
type ClientHandler func(channel ClientChannel) ClientInterface

// ClientInterface is an interface for E2 client procedures
type ClientInterface procedures.E2NodeProcedures

// ClientChannel is an interface for initiating client procedures
type ClientChannel channels.E2NodeChannel

// Connect connects to the given address
func Connect(ctx context.Context, address string, handler ClientHandler) (ClientChannel, error) {
	addr, err := addressing.ResolveAddress(types.Sctp4, address)
	if err != nil {
		return nil, err
	}
	conn, err := sctp.DialSCTP(addr,
		sctp.WithAddressFamily(addr.AddressFamily),
		sctp.WithNonBlocking(false),
		sctp.WithMode(types.OneToOne),
		sctp.WithInitMsg(types.InitMsg{}))
	if err != nil {
		return nil, err
	}
	channel := channels.NewE2NodeChannel(conn, func(channel channels.E2NodeChannel) procedures.E2NodeProcedures {
		return handler(channel)
	})
	return channel, nil
}
