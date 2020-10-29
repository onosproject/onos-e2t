// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/sctp"
	"net"
)

// newHandler creates a new E2 SCTP handler
func newHandler() sctp.Handler {
	return &Handler{}
}

// Handler is the SCTP E2 handler
type Handler struct{}

func (h *Handler) Accept(conn net.Conn) {
	panic("implement me")
}

var _ sctp.Handler = &Handler{}
