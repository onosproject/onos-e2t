// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/cenkalti/backoff/v4"

	"github.com/onosproject/onos-e2t/pkg/topo"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
	"github.com/onosproject/onos-e2t/pkg/ranfunctions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdudecoder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

// TODO: Change the RIC ID to something appropriate
var ricID = types.RicIdentifier{
	RicIdentifierValue: 0xABCDE,
	RicIdentifierLen:   20,
}

func NewE2Server(channels ChannelManager,
	subs subscription.Broker,
	modelRegistry modelregistry.ModelRegistry,
	ranFunctionRegistry ranfunctions.Registry,
	topoManager topo.Manager) *E2Server {
	return &E2Server{
		server:              e2.NewServer(),
		channels:            channels,
		subs:                subs,
		modelRegistry:       modelRegistry,
		ranFunctionRegistry: ranFunctionRegistry,
		topoManager:         topoManager,
	}
}

type E2Server struct {
	server              *e2.Server
	channels            ChannelManager
	subs                subscription.Broker
	modelRegistry       modelregistry.ModelRegistry
	ranFunctionRegistry ranfunctions.Registry
	topoManager         topo.Manager
}

func (s *E2Server) Serve() error {
	return s.server.Serve(func(channel e2.ServerChannel) e2.ServerInterface {
		return &E2ChannelServer{
			serverChannel:       channel,
			manager:             s.channels,
			subs:                s.subs,
			modelRegistry:       s.modelRegistry,
			ranFunctionRegistry: s.ranFunctionRegistry,
			topoManager:         s.topoManager,
		}
	})
}

func (s *E2Server) Stop() error {
	return s.server.Stop()
}

type E2ChannelServer struct {
	manager             ChannelManager
	subs                subscription.Broker
	serverChannel       e2.ServerChannel
	e2Channel           *E2Channel
	modelRegistry       modelregistry.ModelRegistry
	ranFunctionRegistry ranfunctions.Registry
	topoManager         topo.Manager
}

func (e *E2ChannelServer) processRANFunctions(ranFuncs *types.RanFunctions,
	e2NodeID topoapi.ID,
	serviceModels map[string]*topoapi.ServiceModelInfo,
	e2cells *[]*topoapi.E2Cell) (types.RanFunctionRevisions, types.RanFunctionCauses, error) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfRejected := make(types.RanFunctionCauses)
	plugins := e.modelRegistry.GetPlugins()
	for ranFunctionID, ranFunc := range *ranFuncs {
		for smOid, sm := range plugins {
			oid := e2smtypes.OID(ranFunc.OID)
			if smOid == oid {
				serviceModels[string(oid)] = &topoapi.ServiceModelInfo{
					OID: string(oid),
				}

				if setup, ok := sm.(modelregistry.E2Setup); ok {
					onSetupRequest := &e2smtypes.OnSetupRequest{
						ServiceModels:          serviceModels,
						E2Cells:                e2cells,
						RANFunctionDescription: ranFunc.Description,
					}
					err := setup.OnSetup(onSetupRequest)
					if err != nil {
						log.Warn(err)
						log.Debugf("Length of RAN function Description Bytes is: %d", len(onSetupRequest.RANFunctionDescription))
						log.Debugf("RAN Function Description Bytes in hex format: %v", hex.Dump(onSetupRequest.RANFunctionDescription))
					}
				}

				ranFunction := ranfunctions.RANFunction{
					OID: oid,
					ID:  ranFunctionID,
				}

				id := ranfunctions.NewID(oid, string(e2NodeID))
				err := e.ranFunctionRegistry.Add(id, ranFunction)
				if err != nil {
					log.Warn(err)
				} else {
					rfAccepted[ranFunctionID] = ranFunc.Revision
				}
			}
		}
	}
	return rfAccepted, rfRejected, nil

}

func (e *E2ChannelServer) updateRNIB(ctx context.Context, e2NodeID topoapi.ID,
	serviceModels map[string]*topoapi.ServiceModelInfo, e2Cells []*topoapi.E2Cell, relationID topoapi.ID) error {

	notifier := func(err error, t time.Duration) {
		log.Infof("Updating R-NIB is failed: %v", err)
	}
	// create or update E2 node entities
	err := backoff.RetryNotify(func() error {
		err := e.topoManager.CreateOrUpdateE2Node(ctx, e2NodeID, serviceModels)
		return err

	}, newExpBackoff(), notifier)

	// Add E2 cells if there are any associated cells with an E2 node
	if len(e2Cells) != 0 {
		err = backoff.RetryNotify(func() error {
			err := e.topoManager.CreateOrUpdateE2Cells(ctx, e2NodeID, e2Cells)
			return err
		}, newExpBackoff(), notifier)

	}

	// create E2T to E2 node relation
	err = backoff.RetryNotify(func() error {
		err = e.topoManager.CreateOrUpdateE2Relation(ctx, e2NodeID, relationID)
		return err
	}, newExpBackoff(), notifier)

	return err
}

func (e *E2ChannelServer) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (*e2appducontents.E2SetupResponse, *e2appducontents.E2SetupFailure, error) {
	nodeID, ranFuncs, err := pdudecoder.DecodeE2SetupRequest(request)
	if err != nil {
		return nil, nil, err
	}

	e2NodeID, err := GetNodeID(nodeID.NodeIdentifier)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	serviceModels := make(map[string]*topoapi.ServiceModelInfo)
	var e2Cells []*topoapi.E2Cell
	rfAccepted, rfRejected, err := e.processRANFunctions(ranFuncs, e2NodeID, serviceModels, &e2Cells)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	channelID, err := getChannelID(e2NodeID)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	e.e2Channel = NewE2Channel(channelID, e.serverChannel, e.subs)
	e.manager.Open(channelID, e.e2Channel)

	err = e.updateRNIB(ctx, e2NodeID, serviceModels, e2Cells, topoapi.ID(channelID))
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	// Create an E2 setup response
	response, err := pdubuilder.NewE2SetupResponse(nodeID.Plmn, ricID, rfAccepted, rfRejected)
	if err != nil {
		return nil, nil, err
	}
	return response, nil, nil
}

func (e *E2ChannelServer) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	return e.e2Channel.ricIndication(ctx, request)
}
