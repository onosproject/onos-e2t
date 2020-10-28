// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
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

func NewManager() *Manager {
	log.Info("Creating Manager")
	mgr = Manager{
		caPath:   "",
		keyPath:  "",
		certPath: "",
	}
	return &mgr
}

// Run starts the manager and the associated services
func (m *Manager) Run() {
	log.Info("Starting Manager")
	grpcErr := m.startGRPCServer()
	if grpcErr != nil {
		log.Fatal("Unable to start et2 northbound grpc server ", grpcErr)
	}
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

	return s.Serve(func(started string) {
		log.Info("Started NBI on ", started)
	})
}

// UseCAPath overrides the default path for the certificate authority file
func (m *Manager) UseCAPath(caPath string) {
	m.caPath = caPath
}

// UseKeyPath overrides the default path for the key file
func (m *Manager) UseKeyPath(keyPath string) {
	m.keyPath = keyPath
}

// UseCertPath overrides the default path for the certificate file
func (m *Manager) UseCertPath(certPath string) {
	m.caPath = certPath
}
