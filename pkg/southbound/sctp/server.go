// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package sctp

import (
	"github.com/ishidawataru/sctp"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"net"
)

var log = logging.GetLogger("southbound", "sctp")

const (
	defaultSCTPIP = "0.0.0.0"
)

const (
	defaultSCTPPort = 36421
)

// Config is the server configuration
type Config struct {
	Port int
}

// GetPort returns the server port
func (c Config) GetPort() int {
	if c.Port == 0 {
		return defaultSCTPPort
	}
	return c.Port
}

// Handler is a handler for SCTP connections
type Handler interface {
	// Accept handles the given connection
	Accept(conn net.Conn)
}

// NewServer creates a new southbound server
func NewServer(handler Handler, config Config) *Server {
	return &Server{
		Config:  config,
		handler: handler,
	}
}

// Server is a southbound server
type Server struct {
	Config  Config
	handler Handler
}

// Serve starts the server
func (s *Server) Serve(servingCh chan<- error) {
	ip, err := net.ResolveIPAddr("ip", defaultSCTPIP)
	if err != nil {
		servingCh <- err
		close(servingCh)
		return
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: []net.IPAddr{*ip},
		Port:    s.Config.GetPort(),
	}

	ln, err := sctp.ListenSCTP("sctp", addr)
	if err != nil {
		servingCh <- err
		close(servingCh)
		return
	}

	close(servingCh)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Errorf("Failed to accept connection: %v", err)
		} else {
			log.Infof("Accepted connection from %s", conn.RemoteAddr())
			go s.handler.Accept(conn)
		}
	}
}
