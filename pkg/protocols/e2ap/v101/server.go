// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"net"

	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/v101/connections"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/v101/procedures"
	"github.com/onosproject/onos-e2t/pkg/protocols/sctp"
)

// ServerHandler is a server channel handler
type ServerHandler func(channel ServerConn) ServerInterface

// ServerInterface is an E2 server interface
type ServerInterface procedures.RICProcedures

// ServerConn is an interface for initiating E2 server procedures
type ServerConn connections.RICConn

// NewServer creates a new E2 server
func NewServer(opts ...sctp.ServerOption) *Server {
	return &Server{
		server: sctp.NewServer(opts...),
	}
}

// Server is an E2 server
type Server struct {
	server *sctp.Server
}

// Serve starts the server
func (s *Server) Serve(handler ServerHandler) error {
	return s.server.Serve(func(conn net.Conn) {
		connections.NewRICConn(conn, func(channel connections.RICConn) procedures.RICProcedures {
			return handler(channel)
		})
	})
}

// Stop stops the server serving
func (s *Server) Stop() error {
	return s.server.Stop()
}
