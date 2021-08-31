// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

	"github.com/onosproject/onos-e2t/pkg/modelregistry"

	"github.com/google/uuid"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/uri"

	"sync"
	"time"

	e2smtypes "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
)

func NewE2Channel(nodeID topoapi.ID, plmnID string, nodeIdentity *types.E2NodeIdentity, channel e2.ServerChannel,
	streamsv1beta1 subscriptionv1beta1.Broker,
	serviceModels map[string]*topoapi.ServiceModelInfo, ranFunctions map[e2smtypes.OID]RANFunction, e2Cells []*topoapi.E2Cell, now time.Time, modelRegistry modelregistry.ModelRegistry) *E2Channel {

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
		streamsv1beta1: streamsv1beta1,
		ServiceModels:  serviceModels,
		RANFunctions:   ranFunctions,
		E2Cells:        e2Cells,
		modelRegistry:  modelRegistry,
	}
}

// RANFunction RAN function information
type RANFunction struct {
	OID e2smtypes.OID
	ID  types.RanFunctionID
	// protobuf encoded description
	Description []byte
}

type ChannelID string

type E2Channel struct {
	e2.ServerChannel
	ID             ChannelID
	E2NodeID       topoapi.ID
	NodeID         string
	NodeType       types.E2NodeType
	PlmnID         string
	TimeAlive      time.Time
	streamsv1beta1 subscriptionv1beta1.Broker
	ServiceModels  map[string]*topoapi.ServiceModelInfo
	RANFunctions   map[e2smtypes.OID]RANFunction
	E2Cells        []*topoapi.E2Cell
	modelRegistry  modelregistry.ModelRegistry
	mu             sync.RWMutex
}

func (c *E2Channel) GetRANFunctions() []RANFunction {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ranFunctions := make([]RANFunction, 0, len(c.RANFunctions))
	for _, ranFunction := range c.RANFunctions {
		ranFunctions = append(ranFunctions, ranFunction)
	}
	return ranFunctions
}

func (c *E2Channel) GetRANFunction(oid e2smtypes.OID) (RANFunction, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ranFunction, ok := c.RANFunctions[oid]
	return ranFunction, ok
}

func (c *E2Channel) ricIndication(ctx context.Context, ricIndication *e2appducontents.Ricindication) error {
	streamID := subscriptionv1beta1.StreamID(ricIndication.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	stream, ok := c.streamsv1beta1.GetWriter(streamID)
	if !ok {
		return errors.NewNotFound("stream %s not found", streamID)
	}

	ranFuncID := ricIndication.ProtocolIes.E2ApProtocolIes5.Value.Value
	ricActionID := ricIndication.ProtocolIes.E2ApProtocolIes15.Value.Value
	indHeaderAsn1 := ricIndication.ProtocolIes.E2ApProtocolIes25.Value.Value
	indMessageAsn1 := ricIndication.ProtocolIes.E2ApProtocolIes26.Value.Value
	indication := &e2api.Indication{}
	encoding := stream.GetEncoding()
	log.Debugf("Received RIC Indication; RAN FundID: %d, RIC Action ID: %d", ranFuncID, ricActionID)
	switch encoding {
	case e2api.Encoding_PROTO:
		for oid, ranFunction := range c.RANFunctions {
			if ranFunction.ID == types.RanFunctionID(ranFuncID) {
				serviceModelPlugin, err := c.modelRegistry.GetPlugin(oid)
				if err != nil {
					return err
				}
				indHeaderProto, err := serviceModelPlugin.IndicationHeaderASN1toProto(indHeaderAsn1)
				if err != nil {
					log.Errorf("Error transforming Header ASN.1 Bytes to Proto %s", err.Error())
					return errors.NewInvalid(err.Error())
				}
				indMessageProto, err := serviceModelPlugin.IndicationMessageASN1toProto(indMessageAsn1)
				if err != nil {
					log.Errorf("Error transforming Message ASN.1 Bytes to Proto %s", err.Error())
					return errors.NewInvalid(err.Error())
				}
				indication.Header = indHeaderProto
				indication.Payload = indMessageProto
				log.Infof("RICIndication successfully decoded from ASN.1 to Proto #Bytes - Header: %d, Message: %d", len(indHeaderProto), len(indMessageProto))
				break
			}
		}
	case e2api.Encoding_ASN1_PER:
		indication.Header = indHeaderAsn1
		indication.Payload = indMessageAsn1
	default:
		log.Errorf("encoding type %v not supported", stream.GetEncoding())
		return errors.NewInvalid("encoding type %v not supported", encoding)
	}

	return stream.Send(indication)
}
