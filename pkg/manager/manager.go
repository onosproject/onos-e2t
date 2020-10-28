// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
	"github.com/onosproject/onos-e2t/pkg/northbound/ricapie2"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
)

type Manager struct {
	caPath   string
	keyPath  string
	certPath string
	// many things to be added
}

var mgr Manager
var log = logging.GetLogger("manager")

func NewManager(options ...func(*Manager)) (*Manager, error) {
	log.Info("Creating Manager")
	for _, option := range options {
		option(&mgr)
	}

	return &mgr, nil
}

// Run starts the manager and the associated services
func (m *Manager) Run() error {
	log.Info("Starting Manager")
	err := m.startGRPCServer()
	if err != nil {
		return err
	}

	return nil
}

// Close kills the channels and manager related objects
func (m *Manager) Close() {
	log.Info("Closing Manager")
}

// GetManager returns the initialized and running instance of manager.
// Should be called only after NewManager and Run are done.
func GetManager() *Manager {
	return &mgr
}

// Creates gRPC server and registers various services; then serves.
func (m *Manager) startGRPCServer() error {
	s := northbound.NewServer(northbound.NewServerCfg(m.caPath, m.keyPath, m.certPath, 5150, true, northbound.SecurityConfig{}))
	s.AddService(admin.Service{})
	s.AddService(logging.Service{})
	s.AddService(ricapie2.Service{})

	return s.Serve(func(started string) {
		log.Info("Started NBI on ", started)
	})
}

// WithCAPath overrides the default path for the certificate authority file
func WithCAPath(caPath string) func(*Manager) {
	return func(mgr *Manager) {
		mgr.caPath = caPath
	}
}

// WithKeyPath overrides the default path for the key file
func WithKeyPath(keyPath string) func(*Manager) {
	return func(mgr *Manager) {
		mgr.keyPath = keyPath
	}
}

// WithCertPath overrides the default path for the certificate file
func WithCertPath(certPath string) func(*Manager) {
	return func(mgr *Manager) {
		mgr.certPath = certPath
	}
}
