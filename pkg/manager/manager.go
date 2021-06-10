// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package manager

import (
	"context"
	"github.com/atomix/atomix-go-client/pkg/atomix"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	e2v1beta1service "github.com/onosproject/onos-e2t/pkg/northbound/e2/v1beta1"
	substore "github.com/onosproject/onos-e2t/pkg/store/subscription"
	taskstore "github.com/onosproject/onos-e2t/pkg/store/task"
	"time"

	"github.com/onosproject/onos-e2t/pkg/store/rnib"

	"github.com/onosproject/onos-e2t/pkg/topo"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/ranfunctions"
	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/onosproject/onos-e2t/pkg/oid"

	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/server"

	epapi "github.com/onosproject/onos-api/go/onos/e2sub/endpoint"
	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"
	subctrl "github.com/onosproject/onos-e2t/pkg/controller/subscription"
	subctrlv1beta1 "github.com/onosproject/onos-e2t/pkg/controller/v1beta1/subscription"
	taskctrlv1beta1 "github.com/onosproject/onos-e2t/pkg/controller/v1beta1/task"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/northbound/admin"
	"github.com/onosproject/onos-e2t/pkg/northbound/ricapie2"
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
	TopoAddress         string
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

	oidRegistry := oid.NewOidRegistry()

	opts = append(opts, grpc.WithUnaryInterceptor(southbound.RetryingUnaryClientInterceptor()))
	opts = append(opts, grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(time.Second)))
	conn, err := grpc.Dial(config.E2SubAddress, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return &Manager{
		Config:        config,
		ModelRegistry: modelRegistry,
		OidRegistry:   oidRegistry,
		conn:          conn,
	}
}

// Manager is a manager for the E2T service
type Manager struct {
	Config        Config
	ModelRegistry modelregistry.ModelRegistry
	OidRegistry   oid.Registry
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
	opts, err := certs.HandleCertPaths(m.Config.CAPath, m.Config.KeyPath, m.Config.CertPath, true)
	if err != nil {
		return err
	}
	rnibStore, err := rnib.NewStore(m.Config.TopoAddress, opts...)
	if err != nil {
		return err
	}

	atomixClient := atomix.NewClient(atomix.WithClientID(env.GetPodName()))

	subStore, err := substore.NewAtomixStore(atomixClient)
	if err != nil {
		return err
	}
	taskStore, err := taskstore.NewAtomixStore(atomixClient)
	if err != nil {
		return err
	}

	topoManager := topo.NewManager(rnibStore)
	streams := subscription.NewBroker()
	streamsv1beta1 := subscriptionv1beta1.NewBroker()
	channels := e2server.NewChannelManager(topoManager)
	ranFunctionRegistry := ranfunctions.NewRegistry()

	err = m.startSubscriptionController(streams, channels, ranFunctionRegistry, topoManager)
	if err != nil {
		return err
	}
	err = m.startSubscriptionv1beta1Controller(subStore, taskStore)
	if err != nil {
		return err
	}
	err = m.startTaskv1beta1Controller(subStore, taskStore, streamsv1beta1, channels, ranFunctionRegistry, topoManager)
	if err != nil {
		return err
	}

	err = m.startSouthboundServer(channels, streams, streamsv1beta1, ranFunctionRegistry, topoManager)
	if err != nil {
		return err
	}

	err = m.startNorthboundServer(subStore, streams, streamsv1beta1, channels, ranFunctionRegistry, topoManager)
	if err != nil {
		return err
	}
	return m.joinSubscriptionManager()
}

// startSubscriptionController starts the subscription controllers
func (m *Manager) startSubscriptionController(streams subscription.Broker,
	channels e2server.ChannelManager, ranFunctionRegistry ranfunctions.Registry, deviceManager topo.Manager) error {
	controller := subctrl.NewController(streams, subapi.NewE2SubscriptionServiceClient(m.conn),
		subtaskapi.NewE2SubscriptionTaskServiceClient(m.conn),
		channels, m.ModelRegistry, m.OidRegistry, ranFunctionRegistry, deviceManager)
	if err := controller.Start(); err != nil {
		return err
	}
	return nil
}

// startSubscriptionv1beta1Controller starts the subscription controllers
func (m *Manager) startSubscriptionv1beta1Controller(subs substore.Store, tasks taskstore.Store) error {
	subsv1beta1 := subctrlv1beta1.NewController(subs, tasks)
	if err := subsv1beta1.Start(); err != nil {
		return err
	}
	return nil
}

// startTaskv1beta1Controller starts the subscription controllers
func (m *Manager) startTaskv1beta1Controller(subs substore.Store, tasks taskstore.Store, streams subscriptionv1beta1.Broker,
	channels e2server.ChannelManager, ranFunctionRegistry ranfunctions.Registry, deviceManager topo.Manager) error {
	tasksv1beta1 := taskctrlv1beta1.NewController(streams, subs, tasks,
		channels, m.ModelRegistry, m.OidRegistry, ranFunctionRegistry, deviceManager)
	if err := tasksv1beta1.Start(); err != nil {
		return err
	}
	return nil
}

// startSouthboundServer starts the southbound server
func (m *Manager) startSouthboundServer(channels e2server.ChannelManager, streams subscription.Broker,
	streamsv1beta1 subscriptionv1beta1.Broker, ranFunctionRegistry ranfunctions.Registry, topoManager topo.Manager) error {
	server := e2server.NewE2Server(channels, streams, streamsv1beta1, m.ModelRegistry, ranFunctionRegistry, topoManager)
	return server.Serve()
}

// startSouthboundServer starts the northbound gRPC server
func (m *Manager) startNorthboundServer(subs substore.Store, streams subscription.Broker, streamsv1beta1 subscriptionv1beta1.Broker,
	channels e2server.ChannelManager, ranFunctionRegistry ranfunctions.Registry, topoManager topo.Manager) error {
	s := northbound.NewServer(northbound.NewServerCfg(
		m.Config.CAPath,
		m.Config.KeyPath,
		m.Config.CertPath,
		int16(m.Config.GRPCPort),
		true,
		northbound.SecurityConfig{}))
	s.AddService(admin.NewService(channels, ranFunctionRegistry))
	s.AddService(logging.Service{})
	s.AddService(ricapie2.NewService(subapi.NewE2SubscriptionServiceClient(m.conn), streams, m.ModelRegistry,
		channels, m.OidRegistry, ranFunctionRegistry, topoManager))
	s.AddService(e2v1beta1service.NewControlService(m.ModelRegistry,
		channels, m.OidRegistry, ranFunctionRegistry, topoManager))
	s.AddService(e2v1beta1service.NewSubscriptionService(subs, streamsv1beta1, m.ModelRegistry, m.OidRegistry))

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
	if err != nil {
		err = errors.FromGRPC(err)
		if errors.IsAlreadyExists(err) {
			return nil
		}
		return err
	}
	return nil
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
