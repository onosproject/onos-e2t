// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"context"
	"time"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	epapi "github.com/onosproject/onos-api/go/onos/e2sub/endpoint"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	subbroker "github.com/onosproject/onos-e2t/pkg/broker/subscription"
	subctrl "github.com/onosproject/onos-e2t/pkg/controller/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
	"github.com/onosproject/onos-e2t/pkg/northbound/ricapie2"
	"github.com/onosproject/onos-e2t/pkg/northbound/stream"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("manager")

// Config is a manager configuration
type Config struct {
	CAPath              string
	KeyPath             string
	CertPath            string
	GRPCPort            int
	E2Port              int
	E2SubAddress        string
	ServiceModelPlugins []string
}

// NewManager creates a new manager
func NewManager(config Config) *Manager {
	log.Info("Creating Manager")
	opts, err := certs.HandleCertPaths(config.CAPath, config.KeyPath, config.CertPath, true)
	if err != nil {
		log.Fatal(err)
	}

	modelRegistry := modelregistry.NewModelRegistry()
	for _, smp := range config.ServiceModelPlugins {
		if _, _, err := modelRegistry.RegisterModelPlugin(smp); err != nil {
			log.Fatal(err)
		}
	}

	opts = append(opts, grpc.WithUnaryInterceptor(southbound.RetryingUnaryClientInterceptor()))
	opts = append(opts, grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(time.Second)))
	conn, err := grpc.Dial(config.E2SubAddress, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		Config:        config,
		ModelRegistry: modelRegistry,
		conn:          conn,
	}
}

// Manager is a manager for the E2T service
type Manager struct {
	Config        Config
	ModelRegistry modelregistry.ModelRegistry
	conn          *grpc.ClientConn
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
	requests := subctrl.NewRequestJournal()
	streams := stream.NewManager()
	channels := e2server.NewChannelManager()

	err := m.startSubscriptionBroker(requests, streams, channels)
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
func (m *Manager) startSubscriptionBroker(catalog *subctrl.RequestJournal, streams *stream.Manager, channels e2server.ChannelManager) error {
	controller := subctrl.NewController(catalog, subapi.NewE2SubscriptionServiceClient(m.conn), subtaskapi.NewE2SubscriptionTaskServiceClient(m.conn), channels, m.ModelRegistry)
	if err := controller.Start(); err != nil {
		return err
	}

	broker := subbroker.NewBroker(catalog, streams, channels)
	if err := broker.Start(); err != nil {
		return err
	}
	return nil
}

// startSouthboundServer starts the southbound server
func (m *Manager) startSouthboundServer(channels e2server.ChannelManager) error {
	server := e2server.NewE2Server(channels, m.ModelRegistry)
	return server.Serve()
}

// startSouthboundServer starts the northbound gRPC server
func (m *Manager) startNorthboundServer(streams *stream.Manager, channels e2server.ChannelManager) error {
	s := northbound.NewServer(northbound.NewServerCfg(
		m.Config.CAPath,
		m.Config.KeyPath,
		m.Config.CertPath,
		int16(m.Config.GRPCPort),
		true,
		northbound.SecurityConfig{}))
	s.AddService(admin.NewService(channels))
	s.AddService(logging.Service{})
	s.AddService(ricapie2.NewService(subapi.NewE2SubscriptionServiceClient(m.conn), streams, m.ModelRegistry, channels))

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
	client := epapi.NewE2RegistryServiceClient(m.conn)
	request := &epapi.AddTerminationRequest{
		Endpoint: &epapi.TerminationEndpoint{
			ID:   epapi.ID(env.GetPodID()),
			IP:   epapi.IP(env.GetPodIP()),
			Port: epapi.Port(5150),
		},
	}
	_, err := client.AddTermination(context.Background(), request)
	return err
}

// leaveSubscriptionManager removes the termination point from the subscription manager
func (m *Manager) leaveSubscriptionManager() error {
	client := epapi.NewE2RegistryServiceClient(m.conn)
	request := &epapi.RemoveTerminationRequest{
		ID: epapi.ID(env.GetPodID()),
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
