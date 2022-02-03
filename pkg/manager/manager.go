// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package manager

import (
	"github.com/atomix/atomix-go-client/pkg/atomix"
	"github.com/onosproject/onos-e2t/pkg/controller/configuration"
	"github.com/onosproject/onos-e2t/pkg/controller/controlrelation"
	"github.com/onosproject/onos-e2t/pkg/controller/e2t"
	nbstream "github.com/onosproject/onos-e2t/pkg/northbound/e2/stream"
	e2v1beta1service "github.com/onosproject/onos-e2t/pkg/northbound/e2/v1beta1"
	sbstream "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/stream"
	chanstore "github.com/onosproject/onos-e2t/pkg/store/channel"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-e2t/pkg/oid"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"

	"github.com/onosproject/onos-e2t/pkg/controller/mastership"
	subctrlv1beta1 "github.com/onosproject/onos-e2t/pkg/controller/v1beta1/channel"
	taskctrlv1beta1 "github.com/onosproject/onos-e2t/pkg/controller/v1beta1/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
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

	subscriptions, err := sbstream.NewManager()
	if err != nil {
		return err
	}
	streams, err := nbstream.NewManager(subscriptions)
	if err != nil {
		return err
	}
	e2apConns := e2server.NewE2APConnManager()
	mgmtConns := e2server.NewMgmtConnManager()

	err = m.startE2TController(rnibStore)
	if err != nil {
		return err
	}

	err = m.startMastershipController(rnibStore)
	if err != nil {
		return err
	}

	err = m.startChannelv1beta1Controller(chanStore, subStore, streams, rnibStore)
	if err != nil {
		return err
	}
	err = m.startSubscriptionv1beta1Controller(subStore, subscriptions, rnibStore, e2apConns)
	if err != nil {
		return err
	}

	err = m.startConfigurationController(rnibStore, mgmtConns, e2apConns)
	if err != nil {
		return err
	}

	err = m.startControlRelationController(rnibStore, e2apConns)
	if err != nil {
		return err
	}

	err = m.startSouthboundServer(e2apConns, mgmtConns, subscriptions, rnibStore)
	if err != nil {
		return err
	}

	err = m.startNorthboundServer(chanStore, subStore, streams, rnibStore, e2apConns)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) startControlRelationController(rnib rnib.Store, e2apConn e2server.E2APConnManager) error {
	ctrlRelationController := controlrelation.NewController(rnib, e2apConn)
	return ctrlRelationController.Start()
}

func (m *Manager) startConfigurationController(rnib rnib.Store, mgmtConns e2server.MgmtConnManager, e2apConn e2server.E2APConnManager) error {
	connController := configuration.NewController(rnib, mgmtConns, e2apConn)
	return connController.Start()
}

func (m *Manager) startE2TController(rnib rnib.Store) error {
	e2tController := e2t.NewController(rnib)
	return e2tController.Start()
}

// startTopov1alpha1Controller starts the topo controller
func (m *Manager) startMastershipController(topo rnib.Store) error {
	mastershipController := mastership.NewController(topo)
	return mastershipController.Start()
}

// startChannelv1beta1Controller starts the subscription controllers
func (m *Manager) startChannelv1beta1Controller(chans chanstore.Store, subs substore.Store, streams nbstream.Manager, topo rnib.Store) error {
	subsv1beta1 := subctrlv1beta1.NewController(chans, subs, streams, topo)
	return subsv1beta1.Start()
}

// startSubscriptionv1beta1Controller starts the subscription controllers
func (m *Manager) startSubscriptionv1beta1Controller(subs substore.Store, streams sbstream.Manager, topo rnib.Store, e2apConns e2server.E2APConnManager) error {
	tasksv1beta1 := taskctrlv1beta1.NewController(streams, subs, topo, e2apConns, m.ModelRegistry, m.OidRegistry)
	return tasksv1beta1.Start()
}

// startSouthboundServer starts the southbound server
func (m *Manager) startSouthboundServer(e2apConns e2server.E2APConnManager, mgmtConns e2server.MgmtConnManager,
	streams sbstream.Manager, rnib rnib.Store) error {
	server := e2server.NewE2Server(e2apConns, mgmtConns, streams, m.ModelRegistry, rnib)
	return server.Serve()
}

// startSouthboundServer starts the northbound gRPC server
func (m *Manager) startNorthboundServer(chans chanstore.Store, subs substore.Store, streams nbstream.Manager,
	rnib rnib.Store, e2apConns e2server.E2APConnManager) error {
	s := northbound.NewServer(northbound.NewServerCfg(
		m.Config.CAPath,
		m.Config.KeyPath,
		m.Config.CertPath,
		int16(m.Config.GRPCPort),
		true,
		northbound.SecurityConfig{}))
	s.AddService(logging.Service{})
	s.AddService(e2v1beta1service.NewControlService(m.ModelRegistry, e2apConns, m.OidRegistry, rnib))
	s.AddService(e2v1beta1service.NewSubscriptionService(chans, subs, streams, m.ModelRegistry, m.OidRegistry, rnib))

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

// Close kills the connections and manager related objects
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
