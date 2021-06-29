// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	"github.com/onosproject/onos-api/go/onos/topo"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"time"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
)

func NewE2Channel(id ChannelID, plmnID string, nodeID *types.E2NodeIdentity, channel e2.ServerChannel,
	streams subscription.Broker, streamsv1beta1 subscriptionv1beta1.Broker, now time.Time) *E2Channel {
	nid, err := GetNodeID(nodeID.NodeIdentifier)
	if err != nil {
		log.Warn("Unable to parse node ID: %v due to %v", nodeID, err)
		nid = topo.ID(id)
	}
	return &E2Channel{
		ServerChannel:  channel,
		ID:             id,
		PlmnID:         plmnID,
		NodeID:         string(nid),
		NodeType:       nodeID.NodeType,
		TimeAlive:      now,
		streams:        streams,
		streamsv1beta1: streamsv1beta1,
	}
}

type E2Channel struct {
	e2.ServerChannel
	ID             ChannelID
	NodeID         string
	NodeType       types.E2NodeType
	PlmnID         string
	TimeAlive      time.Time
	streams        subscription.Broker
	streamsv1beta1 subscriptionv1beta1.Broker
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
