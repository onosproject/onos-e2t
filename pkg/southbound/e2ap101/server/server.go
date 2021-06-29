// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"encoding/hex"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	"strconv"
	"time"

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
	streams subscription.Broker,
	streamsv1beta1 subscriptionv1beta1.Broker,
	modelRegistry modelregistry.ModelRegistry,
	ranFunctionRegistry ranfunctions.Registry,
	topoManager topo.Manager) *E2Server {
	return &E2Server{
		server:              e2.NewServer(),
		channels:            channels,
		subs:                streams,
		streamsv1beta1:      streamsv1beta1,
		modelRegistry:       modelRegistry,
		ranFunctionRegistry: ranFunctionRegistry,
		topoManager:         topoManager,
	}
}

type E2Server struct {
	server              *e2.Server
	channels            ChannelManager
	subs                subscription.Broker
	streamsv1beta1      subscriptionv1beta1.Broker
	modelRegistry       modelregistry.ModelRegistry
	ranFunctionRegistry ranfunctions.Registry
	topoManager         topo.Manager
}

func (s *E2Server) Serve() error {
	return s.server.Serve(func(channel e2.ServerChannel) e2.ServerInterface {
		return &E2ChannelServer{
			serverChannel:       channel,
			manager:             s.channels,
			streams:             s.subs,
			streamsv1beta1:      s.streamsv1beta1,
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
	streams             subscription.Broker
	streamsv1beta1      subscriptionv1beta1.Broker
	serverChannel       e2.ServerChannel
	e2Channel           *E2Channel
	modelRegistry       modelregistry.ModelRegistry
	ranFunctionRegistry ranfunctions.Registry
	topoManager         topo.Manager
}

func (e *E2ChannelServer) processRANFunctions(ranFuncs *types.RanFunctions,
	e2NodeID topoapi.ID,
	serviceModels map[string]*topoapi.ServiceModelInfo,
	e2cells *[]*topoapi.E2Cell, channelID ChannelID) (types.RanFunctionRevisions, types.RanFunctionCauses, error) {
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

				ranFunctionDescriptionProto, err := sm.RanFuncDescriptionASN1toProto(ranFunc.Description)
				if err != nil {
					log.Warn(err)
					log.Warnf("Following set of bytes of length %v were pushed to the decoder \n%v\n", len(ranFunc.Description), hex.Dump(ranFunc.Description))
					continue
				}

				ranFunction := ranfunctions.RANFunction{
					OID:         oid,
					ID:          ranFunctionID,
					Description: ranFunctionDescriptionProto,
				}

				// TODO channel ID should be changed to e2node ID after admin API is removed
				id := ranfunctions.NewID(oid, string(channelID))
				err = e.ranFunctionRegistry.Add(id, ranFunction)
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

	// create or update E2 node entities
	err := e.topoManager.CreateOrUpdateE2Node(ctx, e2NodeID, serviceModels)
	if err != nil {
		log.Infof("Updating R-NIB is failed: %v", err)
		return err
	}

	// Add E2 cells if there are any associated cells with an E2 node
	if len(e2Cells) != 0 {
		err := e.topoManager.CreateOrUpdateE2Cells(ctx, e2NodeID, e2Cells)
		if err != nil {
			log.Infof("Updating R-NIB is failed: %v", err)
			return err
		}
	}

	// create E2T to E2 node relation
	err = e.topoManager.CreateOrUpdateE2Relation(ctx, e2NodeID, relationID)
	if err != nil {
		log.Infof("Updating R-NIB is failed: %v", err)
		return err
	}
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
	nodeID, ranFuncs, err := pdudecoder.DecodeE2SetupRequest(request)
	if err != nil {
		return nil, nil, err
	}

	e2NodeID, err := GetNodeID(nodeID.NodeIdentifier)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}
	rawPlmnid := []byte{nodeID.Plmn[0], nodeID.Plmn[1], nodeID.Plmn[2]}
	plmnID := strconv.FormatUint(uint64(uint24ToUint32(rawPlmnid)), 10)

	serviceModels := make(map[string]*topoapi.ServiceModelInfo)
	var e2Cells []*topoapi.E2Cell
	channelID, err := getChannelID(e2NodeID)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	rfAccepted, rfRejected, err := e.processRANFunctions(ranFuncs, e2NodeID, serviceModels, &e2Cells, channelID)
	if err != nil {
		log.Warn(err)
		return nil, nil, err
	}

	e.e2Channel = NewE2Channel(channelID, plmnID, nodeID, e.serverChannel, e.streams, e.streamsv1beta1, time.Now())
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
