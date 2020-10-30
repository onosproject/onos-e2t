// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/connection"
	"github.com/onosproject/onos-e2t/pkg/southbound/sctp"
	"net"
)

// newHandler creates a new E2 SCTP handler
func newHandler(conns *connection.Manager) sctp.Handler {
	return &Handler{
		conns: conns,
	}
}

// Handler is the SCTP E2 handler
type Handler struct {
	conns *connection.Manager
}

func (h *Handler) Accept(conn net.Conn) {
	log.Infof("Initializing E2 connection from %s", conn.RemoteAddr())
	c, err := connection.NewConnection(conn)
	if err != nil {
		log.Infof("Failed to initialize E2 connection from %s: %v", conn.RemoteAddr(), err)
	} else {
		log.Infof("Established E2 connection from %s", conn.RemoteAddr())
		h.conns.Register(c)
		<-c.Context().Done()
		h.conns.Unregister(c)
	}
}

var _ sctp.Handler = &Handler{}
