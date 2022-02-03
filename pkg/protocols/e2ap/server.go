// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap/procedures"
	"github.com/onosproject/onos-e2t/pkg/protocols/sctp"
	"net"
)

// ServerHandler is a server connection handler
type ServerHandler func(conn ServerConn) ServerInterface

// ServerInterface is an E2 server interface
type ServerInterface procedures.RICProcedures

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
	return s.server.Serve(func(c net.Conn) {
		NewServerConn(c, func(conn ServerConn) ServerInterface {
			return handler(conn)
		})
	})
}

// Stop stops the server serving
func (s *Server) Stop() error {
	return s.server.Stop()
}
