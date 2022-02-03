// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package sctp

import (
	"net"

	"github.com/onosproject/onos-lib-go/pkg/sctp/addressing"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/sctp/listener"
	"github.com/onosproject/onos-lib-go/pkg/sctp/types"
)

var log = logging.GetLogger("southbound", "sctp")

const (
	defaultSCTPPort = 36421
)

// ServerOptions is SCTP options
type ServerOptions struct {
	Port            int
	WriteBufferSize int
	ReadBufferSize  int
}

// ServerOption is an SCTP option function
type ServerOption func(*ServerOptions)

// WriteBuffer sets the write buffer size
func WriteBuffer(size int) ServerOption {
	return func(options *ServerOptions) {
		options.WriteBufferSize = size
	}
}

// ReadBuffer sets the read buffer size
func ReadBuffer(size int) ServerOption {
	return func(options *ServerOptions) {
		options.ReadBufferSize = size
	}
}

// Handler is a handler for SCTP connections
type Handler func(conn net.Conn)

// NewServer creates a new southbound server
func NewServer(opts ...ServerOption) *Server {
	return &Server{
		options: applyServerOptions(opts...),
	}
}

// Server is a southbound server
type Server struct {
	options ServerOptions
	lis     *listener.Listener
}

// Serve starts the server
func (s *Server) Serve(handler Handler) error {
	addr := &addressing.Address{
		IPAddrs: []net.IPAddr{},
		Port:    s.options.Port,
	}

	ln, err := listener.NewListener(addr,
		listener.WithMode(types.OneToOne),
		listener.WithNonBlocking(false))
	if err != nil {
		return err
	}
	s.lis = ln

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Errorf("Failed to accept connection: %v", err)
				continue
			}

			log.Infof("Accepted connection from %s", conn.RemoteAddr())
			go handler(conn)
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	return s.lis.Close()
}

func applyServerOptions(opts ...ServerOption) ServerOptions {
	options := ServerOptions{
		Port: defaultSCTPPort,
	}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}
