// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"
	subscriptionv1beta1 "github.com/onosproject/onos-e2t/pkg/broker/subscription/v1beta1"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
)

func NewE2Channel(id ChannelID, plmnID string, channel e2.ServerChannel, streams subscription.Broker, streamsv1beta1 subscriptionv1beta1.Broker) *E2Channel {
	return &E2Channel{
		ServerChannel:  channel,
		ID:             id,
		PlmnID:         plmnID,
		streams:        streams,
		streamsv1beta1: streamsv1beta1,
	}
}

type E2Channel struct {
	e2.ServerChannel
	ID             ChannelID
	PlmnID         string
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
