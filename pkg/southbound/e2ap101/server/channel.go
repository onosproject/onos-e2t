// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"context"

	"github.com/onosproject/onos-e2t/pkg/broker/subscription"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2 "github.com/onosproject/onos-e2t/pkg/protocols/e2ap101"
)

func NewE2Channel(id ChannelID, plmdID string, channel e2.ServerChannel, streams subscription.Broker) *E2Channel {
	return &E2Channel{
		ServerChannel: channel,
		ID:            id,
		PlmnID:        plmdID,
		streams:       streams,
	}
}

type E2Channel struct {
	e2.ServerChannel
	ID      ChannelID
	PlmnID  string
	streams subscription.Broker
}

func (c *E2Channel) ricIndication(ctx context.Context, request *e2appducontents.Ricindication) error {
	streamID := subscription.StreamID(request.ProtocolIes.E2ApProtocolIes29.Value.RicRequestorId)
	stream, err := c.streams.GetStream(streamID)
	if err != nil {
		return err
	}
	return stream.Send(request)
}
