// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"context"
	regapi "github.com/onosproject/onos-e2sub/api/e2/registry/v1beta1"
	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	subtaskapi "github.com/onosproject/onos-e2sub/api/e2/task/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
	"github.com/onosproject/onos-e2t/pkg/northbound/ricapie2"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel"
	sub "github.com/onosproject/onos-e2t/pkg/subscription"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("manager")

// Config is a manager configuration
type Config struct {
	CAPath       string
	KeyPath      string
	CertPath     string
	GRPCPort     int
	E2Port       int
	E2SubAddress string
}

// NewManager creates a new manager
func NewManager(config Config) *Manager {
	log.Info("Creating Manager")
	opts, err := certs.HandleCertPaths(config.CAPath, config.KeyPath, config.CertPath, true)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(config.E2SubAddress, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		Config: config,
		conn:   conn,
	}
}

// Manager is a manager for the E2T service
type Manager struct {
	Config Config
	conn   *grpc.ClientConn
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
	catalog := sub.NewCatalog()
	streams := stream.NewManager()
	channels := channel.NewManager()

	err := m.startSubscriptionBroker(catalog, streams, channels)
	if err != nil {
		return err
	}

	err = m.startSouthboundServer(channels)
	if err != nil {
		return err
	}

	err = m.startNorthboundServer(streams, channels)
	if err != nil {
		return err
	}
	return m.joinSubscriptionManager()
}

// startSubscriptionBroker starts the subscription broker
func (m *Manager) startSubscriptionBroker(catalog *sub.Catalog, streams *stream.Manager, channels *channel.Manager) error {
	controller := sub.NewController(catalog, subapi.NewE2SubscriptionServiceClient(m.conn), subtaskapi.NewE2SubscriptionTaskServiceClient(m.conn), channels)
	if err := controller.Start(); err != nil {
		return err
	}

	broker := sub.NewBroker(catalog, streams, channels)
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
func (m *Manager) startNorthboundServer(streams *stream.Manager, channels *channel.Manager) error {
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

// joinSubscriptionManager joins the termination point to the subscription manager
func (m *Manager) joinSubscriptionManager() error {
	client := regapi.NewE2RegistryServiceClient(m.conn)
	request := &regapi.AddTerminationRequest{
		EndPoint: &regapi.TerminationEndPoint{
			ID:   regapi.ID(env.GetPodID()),
			IP:   regapi.IP(env.GetPodIP()),
			Port: regapi.Port(5150),
		},
	}
	_, err := client.AddTermination(context.Background(), request)
	return err
}

// leaveSubscriptionManager removes the termination point from the subscription manager
func (m *Manager) leaveSubscriptionManager() error {
	client := regapi.NewE2RegistryServiceClient(m.conn)
	request := &regapi.RemoveTerminationRequest{
		ID: regapi.ID(env.GetPodID()),
	}
	_, err := client.RemoveTermination(context.Background(), request)
	return err
}

// Close kills the channels and manager related objects
func (m *Manager) Close() {
	log.Info("Closing Manager")
	if err := m.Stop(); err != nil {
		log.Fatal("Unable to Close Manager", err)
	}
}

// Stop stops the manager
func (m *Manager) Stop() error {
	err := m.leaveSubscriptionManager()
	_ = m.conn.Close()
	return err
}
