// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"github.com/atomix/atomix-go-client/pkg/atomix"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	e2v1beta1service "github.com/onosproject/onos-e2t/pkg/northbound/e2/v1beta1"
	chanstore "github.com/onosproject/onos-e2t/pkg/store/channel"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/oid"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	topoctrl "github.com/onosproject/onos-e2t/pkg/controller/topo"
	subctrlv1beta1 "github.com/onosproject/onos-e2t/pkg/controller/v1beta1/channel"
	taskctrlv1beta1 "github.com/onosproject/onos-e2t/pkg/controller/v1beta1/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
)

var log = logging.GetLogger("manager")

// Config is a manager configuration
type Config struct {
	CAPath              string
	KeyPath             string
	CertPath            string
	GRPCPort            int
	E2Port              int
	TopoAddress         string
	ServiceModelPlugins []string
}

// NewManager creates a new manager
func NewManager(config Config) *Manager {
	log.Info("Creating Manager")
	modelRegistry := modelregistry.NewModelRegistry()
	for _, smp := range config.ServiceModelPlugins {
		if _, _, err := modelRegistry.RegisterModelPlugin(smp); err != nil {
			log.Fatal(err)
		}
	}

	oidRegistry := oid.NewOidRegistry()

	return &Manager{
		Config:        config,
		ModelRegistry: modelRegistry,
		OidRegistry:   oidRegistry,
	}
}

// Manager is a manager for the E2T service
type Manager struct {
	Config        Config
	ModelRegistry modelregistry.ModelRegistry
	OidRegistry   oid.Registry
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
	opts, err := certs.HandleCertPaths(m.Config.CAPath, m.Config.KeyPath, m.Config.CertPath, true)
	if err != nil {
		return err
	}
	rnibStore, err := rnib.NewStore(m.Config.TopoAddress, opts...)
	if err != nil {
		return err
	}

	atomixClient := atomix.NewClient(atomix.WithClientID(env.GetPodName()))

	chanStore, err := chanstore.NewAtomixStore(atomixClient)
	if err != nil {
		return err
	}
	subStore, err := substore.NewAtomixStore(atomixClient)
	if err != nil {
		return err
	}

	streams := subscription.NewBroker()
	streamsv1beta1 := subscriptionv1beta1.NewBroker()
	channels := e2server.NewChannelManager()

	err = m.startTopov1alpha1Controller(rnibStore, channels)
	if err != nil {
		return err
	}
	err = m.startChannelv1beta1Controller(chanStore, subStore, streamsv1beta1)
	if err != nil {
		return err
	}
	err = m.startSubscriptionv1beta1Controller(subStore, streamsv1beta1, rnibStore, channels)
	if err != nil {
		return err
	}

	err = m.startSouthboundServer(channels, streams, streamsv1beta1)
	if err != nil {
		return err
	}

	err = m.startNorthboundServer(chanStore, subStore, streamsv1beta1, rnibStore, channels)
	if err != nil {
		return err
	}
	return nil
}

// startTopov1alpha1Controller starts the topo controller
func (m *Manager) startTopov1alpha1Controller(topo rnib.Store, channels e2server.ChannelManager) error {
	subsv1beta1 := topoctrl.NewController(topo, channels)
	return subsv1beta1.Start()
}

// startChannelv1beta1Controller starts the subscription controllers
func (m *Manager) startChannelv1beta1Controller(chans chanstore.Store, subs substore.Store, streams subscriptionv1beta1.Broker) error {
	subsv1beta1 := subctrlv1beta1.NewController(chans, subs, streams)
	return subsv1beta1.Start()
}

// startSubscriptionv1beta1Controller starts the subscription controllers
func (m *Manager) startSubscriptionv1beta1Controller(subs substore.Store, streams subscriptionv1beta1.Broker, topo rnib.Store, channels e2server.ChannelManager) error {
	tasksv1beta1 := taskctrlv1beta1.NewController(streams, subs, topo, channels, m.ModelRegistry, m.OidRegistry)
	return tasksv1beta1.Start()
}

// startSouthboundServer starts the southbound server
func (m *Manager) startSouthboundServer(channels e2server.ChannelManager, streams subscription.Broker,
	streamsv1beta1 subscriptionv1beta1.Broker) error {
	server := e2server.NewE2Server(channels, streams, streamsv1beta1, m.ModelRegistry)
	return server.Serve()
}

// startSouthboundServer starts the northbound gRPC server
func (m *Manager) startNorthboundServer(chans chanstore.Store, subs substore.Store, streamsv1beta1 subscriptionv1beta1.Broker,
	topo rnib.Store, channels e2server.ChannelManager) error {
	s := northbound.NewServer(northbound.NewServerCfg(
		m.Config.CAPath,
		m.Config.KeyPath,
		m.Config.CertPath,
		int16(m.Config.GRPCPort),
		true,
		northbound.SecurityConfig{}))
	s.AddService(admin.NewService(channels))
	s.AddService(logging.Service{})
	s.AddService(e2v1beta1service.NewControlService(m.ModelRegistry, channels, m.OidRegistry, topo))
	s.AddService(e2v1beta1service.NewSubscriptionService(chans, subs, streamsv1beta1, m.ModelRegistry, m.OidRegistry))

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
	if err := m.Stop(); err != nil {
		log.Fatal("Unable to Close Manager", err)
	}
}

// Stop stops the manager
func (m *Manager) Stop() error {
	return nil
}
