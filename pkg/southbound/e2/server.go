// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	"github.com/onosproject/onos-e2t/pkg/southbound/sctp"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("southbound", "e2")

const (
	defaultE2Port = 36421
)

// Config is the server configuration
type Config struct {
	Port int
}

// GetPort returns the server port
func (c Config) GetPort() int {
	if c.Port == 0 {
		return defaultE2Port
	}
	return c.Port
}

// NewServer creates a new E2 server
func NewServer(config Config, channels *channel.Manager) *Server {
	s := sctp.NewServer(newHandler(channels), sctp.Config{
		Port: config.GetPort(),
	})
	return &Server{
		Server: s,
	}
}

// Server is a southbound server
type Server struct {
	*sctp.Server
}
