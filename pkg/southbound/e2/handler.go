// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/southbound/sctp"
	"net"
)

// newHandler creates a new E2 SCTP handler
func newHandler(channels *channel.Manager) sctp.Handler {
	return &Handler{
		channels: channels,
	}
}

// Handler is the SCTP E2 handler
type Handler struct {
	channels *channel.Manager
}

func (h *Handler) Accept(conn net.Conn) {
	channel, err := h.channels.Open(context.Background(), conn)
	if err != nil {
		log.Infof("Failed to initialize E2 connection from %s: %v", conn.RemoteAddr(), err)
	} else {
		log.Infof("Established E2 connection from %s", conn.RemoteAddr())
		<-channel.Context().Done()
	}
}

var _ sctp.Handler = &Handler{}
