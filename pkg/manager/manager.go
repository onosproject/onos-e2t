// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
	"github.com/onosproject/onos-e2t/pkg/northbound/ricapie2"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/northbound/subscription"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	sub "github.com/onosproject/onos-e2t/pkg/subscription"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
)

var log = logging.GetLogger("manager")

// Config is a manager configuration
type Config struct {
	CAPath   string
	KeyPath  string
	CertPath string
	GRPCPort int
	E2Port   int
}

// NewManager creates a new manager
func NewManager(config Config) *Manager {
	log.Info("Creating Manager")
	return &Manager{
		Config: config,
	}
}

// Manager is a manager for the E2T service
type Manager struct {
	Config Config
}

// Run starts the manager and the associated services
func (m *Manager) Run() {
	log.Info("Running Manager")
	if err := m.Start(); err != nil {
		log.Fatal("Unable to run Manager", err)
	}
}

// Start starts the manager
func (m *Manager) Start() error {
	subs, err := substore.NewStore()
	if err != nil {
		return err
	}

	streams := stream.NewManager()
	channels := channel.NewManager()

	err = m.startSubscriptionBroker(subs, streams, channels)
	if err != nil {
		return err
	}

	err = m.startSouthboundServer(channels)
	if err != nil {
		return err
	}

	err = m.startNorthboundServer(subs, streams, channels)
	if err != nil {
		return err
	}
	return nil
}

// startSubscriptionBroker starts the subscription broker
func (m *Manager) startSubscriptionBroker(subs substore.Store, streams *stream.Manager, channels *channel.Manager) error {
	controller := sub.NewController(subs, channels)
	if err := controller.Start(); err != nil {
		return err
	}

	broker := sub.NewBroker(subs, streams, channels)
	if err := broker.Start(); err != nil {
		return err
	}
	return nil
}

// startSouthboundServer starts the southbound server
func (m *Manager) startSouthboundServer(channels *channel.Manager) error {
	config := e2.Config{
		Port: m.Config.E2Port,
	}
	server := e2.NewServer(config, channels)
	doneCh := make(chan error)
	go server.Serve(doneCh)
	return <-doneCh
}

// startSouthboundServer starts the northbound gRPC server
func (m *Manager) startNorthboundServer(subs substore.Store, streams *stream.Manager, channels *channel.Manager) error {
	s := northbound.NewServer(northbound.NewServerCfg(
		m.Config.CAPath,
		m.Config.KeyPath,
		m.Config.CertPath,
		int16(m.Config.GRPCPort),
		true,
		northbound.SecurityConfig{}))
	s.AddService(admin.NewService(channels))
	s.AddService(logging.Service{})
	s.AddService(ricapie2.NewService(streams))
	s.AddService(subscription.NewService(subs))

	doneCh := make(chan error)
	go func() {
		err := s.Serve(func(started string) {
			log.Info("Started NBI on ", started)
			close(doneCh)
		})
		if err != nil {
			doneCh <- err
		}
	}()
	return <-doneCh
}

// Close kills the channels and manager related objects
func (m *Manager) Close() {
	log.Info("Closing Manager")
}
