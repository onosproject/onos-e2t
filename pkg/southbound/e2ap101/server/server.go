// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"encoding/hex"
	"fmt"
	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	"time"

	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"

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
	modelRegistry modelregistry.ModelRegistry) *E2Server {
	return &E2Server{
		server:         e2.NewServer(),
		channels:       channels,
		subs:           streams,
		streamsv1beta1: streamsv1beta1,
		modelRegistry:  modelRegistry,
	}
}

type E2Server struct {
	server         *e2.Server
	channels       ChannelManager
	subs           subscription.Broker
	streamsv1beta1 subscriptionv1beta1.Broker
	modelRegistry  modelregistry.ModelRegistry
}

func (s *E2Server) Serve() error {
	return s.server.Serve(func(channel e2.ServerChannel) e2.ServerInterface {
		return &E2ChannelServer{
			serverChannel:  channel,
			manager:        s.channels,
			streams:        s.subs,
			streamsv1beta1: s.streamsv1beta1,
			modelRegistry:  s.modelRegistry,
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

	var e2Cells []*topoapi.E2Cell
	serviceModels := make(map[string]*topoapi.ServiceModelInfo)
	ranFunctions := make(map[e2smtypes.OID]RANFunction)
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
						E2Cells:                &e2Cells,
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

				ranFunction := RANFunction{
					OID:         oid,
					ID:          ranFunctionID,
					Description: ranFunctionDescriptionProto,
				}

				// TODO channel ID should be changed to e2node ID after admin API is removed
				ranFunctions[oid] = ranFunction
				if err != nil {
					log.Warn(err)
				} else {
					rfAccepted[ranFunctionID] = ranFunc.Revision
				}
			}
		}
	}

	e.e2Channel = NewE2Channel(createE2NodeURI(nodeIdentity), plmnID, nodeIdentity, e.serverChannel, e.streams, e.streamsv1beta1, serviceModels, ranFunctions, e2Cells, time.Now())
	defer e.manager.open(e.e2Channel)

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
