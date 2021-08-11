// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"fmt"
	"time"

	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"

	"github.com/onosproject/onos-e2t/pkg/topo"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdudecoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

// TODO: Change the RIC ID to something appropriate
var ricID = types.RicIdentifier{
	RicIdentifierValue: []byte{0xDE, 0xBC, 0xA0},
	RicIdentifierLen:   20,
}

func NewE2Server(channels ChannelManager,
	streams subscription.Broker,
	streamsv1beta1 subscriptionv1beta1.Broker,
	modelRegistry modelregistry.ModelRegistry,
	topoManager topo.Manager) *E2Server {
	return &E2Server{
		server:         e2.NewServer(),
		channels:       channels,
		subs:           streams,
		streamsv1beta1: streamsv1beta1,
		modelRegistry:  modelRegistry,
		topoManager:    topoManager,
	}
}

type E2Server struct {
	server         *e2.Server
	channels       ChannelManager
	subs           subscription.Broker
	streamsv1beta1 subscriptionv1beta1.Broker
	modelRegistry  modelregistry.ModelRegistry
	topoManager    topo.Manager
}

func (s *E2Server) Serve() error {
	return s.server.Serve(func(channel e2.ServerChannel) e2.ServerInterface {
		return &E2ChannelServer{
			serverChannel:  channel,
			manager:        s.channels,
			streams:        s.subs,
			streamsv1beta1: s.streamsv1beta1,
			modelRegistry:  s.modelRegistry,
			topoManager:    s.topoManager,
		}
	})
}

func (s *E2Server) Stop() error {
	return s.server.Stop()
}

type E2ChannelServer struct {
	manager        ChannelManager
	streams        subscription.Broker
	streamsv1beta1 subscriptionv1beta1.Broker
	serverChannel  e2.ServerChannel
	e2Channel      *E2Channel
	modelRegistry  modelregistry.ModelRegistry
	topoManager    topo.Manager
}

func (e *E2ChannelServer) updateRNIB(ctx context.Context, e2NodeID topoapi.ID,
	serviceModels map[string]*topoapi.ServiceModelInfo, e2Cells []*topoapi.E2Cell, relationID topoapi.ID) error {
	log.Infof("Adding channel '%s' relation to R-NIB", relationID)
	err := e.topoManager.CreateOrUpdateE2T(ctx)
	if err != nil {
		log.Warnf("Updating R-NIB is failed: %v", err)
		return err
	}

	// create or update E2 node entities
	err = e.topoManager.CreateOrUpdateE2Node(ctx, e2NodeID, serviceModels)
	if err != nil {
		log.Warnf("Updating R-NIB is failed: %v", err)
		return err
	}

	// Add E2 cells if there are any associated cells with an E2 node
	if len(e2Cells) != 0 {
		err := e.topoManager.CreateOrUpdateE2Cells(ctx, e2NodeID, e2Cells)
		if err != nil {
			log.Warnf("Updating R-NIB is failed: %v", err)
			return err
		}
	}

	// create E2T to E2 node relation
	err = e.topoManager.CreateOrUpdateE2Relation(ctx, e2NodeID, relationID)
	if err != nil {
		log.Warnf("Updating R-NIB is failed: %v", err)
		return err
	}

	go func() {
		<-e.e2Channel.Context().Done()
		log.Infof("Removing channel '%s' relation from R-NIB", relationID)
		err := e.topoManager.DeleteE2Relation(context.Background(), relationID)
		if err != nil {
			log.Warnf("Updating R-NIB is failed: %v", err)
		}
	}()
	return nil
}

// uint24ToUint32 converts uint24 uint32
func uint24ToUint32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 3; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}

func (e *E2ChannelServer) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (*e2appducontents.E2SetupResponse, *e2appducontents.E2SetupFailure, error) {
	nodeIdentity, ranFuncs, err := pdudecoder.DecodeE2SetupRequest(request)
	if err != nil {
		return nil, nil, err
	}

	rawPlmnid := []byte{nodeIdentity.Plmn[0], nodeIdentity.Plmn[1], nodeIdentity.Plmn[2]}
	plmnID := fmt.Sprintf("%x", uint24ToUint32(rawPlmnid))

	e2NodeID := createE2NodeURI(nodeIdentity)

	serviceModels := make(map[string]*topoapi.ServiceModelInfo)
	var e2Cells []*topoapi.E2Cell
	e.e2Channel = NewE2Channel(e2NodeID, plmnID, nodeIdentity, e.serverChannel, e.streams, e.streamsv1beta1, e.modelRegistry, time.Now())
	rfAccepted, rfRejected, err := e.e2Channel.processRANFunctions(ranFuncs, serviceModels, &e2Cells)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	e.manager.Open(e2NodeID, e.e2Channel)

	err = e.updateRNIB(ctx, e2NodeID, serviceModels, e2Cells, topoapi.ID(e.e2Channel.ID))
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	// Create an E2 setup response
	response, err := pdubuilder.NewE2SetupResponse(nodeIdentity.Plmn, ricID, rfAccepted, rfRejected)
	if err != nil {
		return nil, nil, err
	}
	return response, nil, nil
}

func (e *E2ChannelServer) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	return e.e2Channel.ricIndication(ctx, request)
}
