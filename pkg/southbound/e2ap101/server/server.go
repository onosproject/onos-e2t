// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"encoding/hex"

	"encoding/binary"
	"strconv"

	"github.com/onosproject/onos-e2t/pkg/topo"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"
	"github.com/onosproject/onos-e2t/pkg/ranfunctions"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
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

// uint24ToUint32 converts uint24 uint32
func uint24ToUint32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 3; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}

func getPlmnID(plmn types.PlmnID) string {
	rawPlmnid := []byte{plmn[0], plmn[1], plmn[2]}
	plmnID := strconv.FormatUint(uint64(uint24ToUint32(rawPlmnid)), 10)
	return plmnID
}

func (e *E2ChannelServer) processRANFunctions(ranFuncs *types.RanFunctions,
	deviceID topoapi.ID,
	serviceModels map[string]*topoapi.ServiceModelInfo,
	e2cells *[]*topoapi.E2Cell) (types.RanFunctionRevisions, types.RanFunctionCauses, error) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfRejected := make(types.RanFunctionCauses)
	plugins := e.modelRegistry.GetPlugins()
	for ranFunctionID, ranFunc := range *ranFuncs {
		for smOid, sm := range plugins {
			oid := e2smtypes.OID(ranFunc.OID)
			if smOid == oid {
				ranFunctionDescriptionProto, err := sm.RanFuncDescriptionASN1toProto(ranFunc.Description)
				if err != nil {
					log.Warn(err)
					log.Warnf("Following set of bytes of length %v were pushed to the decoder \n%v\n", len(ranFunc.Description), hex.Dump(ranFunc.Description))
					continue

				serviceModels[string(oid)] = &topoapi.ServiceModelInfo{
					OID: string(oid),
				}

				if setup, ok := sm.(modelregistry.Setup); ok {
					onSetupRequest := &e2smtypes.OnSetupRequest{
						ServiceModels:          serviceModels,
						E2Cells:                e2cells,
						RANFunctionDescription: ranFunc.Description,
					}
					err := setup.OnSetup(onSetupRequest)
					if err != nil {
						log.Warn(err)
					}
				}

				ranFunction := ranfunctions.RANFunction{
					OID: oid,
					ID:  ranFunctionID,
				}

				id := ranfunctions.NewID(oid, string(deviceID))
				err := e.ranFunctionRegistry.Add(id, ranFunction)
				if err != nil {
					log.Warn(err)
				}
				rfAccepted[ranFunctionID] = ranFunc.Revision
			}
		}
	}
	return rfAccepted, rfRejected, nil

}

func (e *E2ChannelServer) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (*e2appducontents.E2SetupResponse, *e2appducontents.E2SetupFailure, error) {
	nodeID, ranFuncs, err := pdudecoder.DecodeE2SetupRequest(request)
	if err != nil {
		return nil, nil, err
	}

	deviceID := topoapi.ID(strconv.FormatUint(binary.BigEndian.Uint64(nodeID.NodeIdentifier), 10))
	channelID := ChannelID(deviceID)
	plmnID := getPlmnID(nodeID.Plmn)

	serviceModels := make(map[string]*topoapi.ServiceModelInfo)
	var e2Cells []*topoapi.E2Cell
	rfAccepted, rfRejected, err := e.processRANFunctions(ranFuncs, deviceID, serviceModels, &e2Cells)

	err = e.topoManager.CreateOrUpdateE2Device(deviceID, serviceModels)
	if err != nil {
		return nil, nil, err
	}

	if len(e2Cells) != 0 {
		err := e.topoManager.CreateOrUpdateE2Cells(deviceID, e2Cells)
		if err != nil {
			return nil, nil, err
		}
	}

	e.e2Channel = NewE2Channel(channelID, plmnID, e.serverChannel, e.subs)
	e.manager.Open(channelID, e.e2Channel)

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
