// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"encoding/hex"
	"sync"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/uri"

	"github.com/google/uuid"
	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/modelregistry"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
)

func NewE2Channel(nodeID topoapi.ID, plmnID string, nodeIdentity *types.E2NodeIdentity, channel e2.ServerChannel,
	streams subscription.Broker, streamsv1beta1 subscriptionv1beta1.Broker,
	modelRegistry modelregistry.ModelRegistry, now time.Time) *E2Channel {

	channelID := ChannelID(uri.NewURI(
		uri.WithScheme("uuid"),
		uri.WithOpaque(uuid.New().String())).String())

	return &E2Channel{
		ServerChannel:  channel,
		ID:             channelID,
		E2NodeID:       nodeID,
		PlmnID:         plmnID,
		NodeID:         string(nodeID),
		NodeType:       nodeIdentity.NodeType,
		TimeAlive:      now,
		streams:        streams,
		streamsv1beta1: streamsv1beta1,
		modelRegistry:  modelRegistry,
		ranFunctions:   make(map[e2smtypes.OID]RANFunction),
	}
}

// RANFunction RAN function information
type RANFunction struct {
	OID e2smtypes.OID
	ID  types.RanFunctionID
	// protobuf encoded description
	Description []byte
}

type E2Channel struct {
	e2.ServerChannel
	ID             ChannelID
	E2NodeID       topoapi.ID
	NodeID         string
	NodeType       types.E2NodeType
	PlmnID         string
	TimeAlive      time.Time
	streams        subscription.Broker
	streamsv1beta1 subscriptionv1beta1.Broker
	modelRegistry  modelregistry.ModelRegistry
	ranFunctions   map[e2smtypes.OID]RANFunction
	mu             sync.RWMutex
}

func (c *E2Channel) GetRANFunctions() []RANFunction {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ranFunctions := make([]RANFunction, 0, len(c.ranFunctions))
	for _, ranFunction := range c.ranFunctions {
		ranFunctions = append(ranFunctions, ranFunction)
	}
	return ranFunctions
}

func (c *E2Channel) GetRANFunction(oid e2smtypes.OID) (RANFunction, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ranFunction, ok := c.ranFunctions[oid]
	return ranFunction, ok
}

func (c *E2Channel) processRANFunctions(
	ranFuncs *types.RanFunctions, serviceModels map[string]*topoapi.ServiceModelInfo,
	e2cells *[]*topoapi.E2Cell) (types.RanFunctionRevisions, types.RanFunctionCauses, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	rfAccepted := make(types.RanFunctionRevisions)
	rfRejected := make(types.RanFunctionCauses)
	plugins := c.modelRegistry.GetPlugins()
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

				ranFunction := RANFunction{
					OID:         oid,
					ID:          ranFunctionID,
					Description: ranFunctionDescriptionProto,
				}

				// TODO channel ID should be changed to e2node ID after admin API is removed
				c.ranFunctions[oid] = ranFunction
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

func (c *E2Channel) ricIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	streamID := subscriptionv1beta1.StreamID(request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	stream, ok := c.streamsv1beta1.GetWriter(streamID)
	if !ok {
		deprecatedStreamID := subscription.StreamID(request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
		deprecatedStream, err := c.streams.GetStream(deprecatedStreamID)
		if err != nil {
			return err
		}
		return deprecatedStream.Send(request)
	}
	return stream.Send(request)
}
