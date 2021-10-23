// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package channel

import (
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

type Channel interface {
	ID() e2api.ChannelID
	Channel() *e2api.Channel
	Indications() <-chan *e2appducontents.Ricindication
	Done() <-chan error
	Close(err error)
}

func newChannelStream(channel *e2api.Channel, buffer Buffer, manager *channelManager) Channel {
	return &channelStream{
		manager: manager,
		channel: channel,
		buffer:  buffer,
		errs:    make(chan error, 1),
	}
}

type channelStream struct {
	manager *channelManager
	channel *e2api.Channel
	buffer  Buffer
	errs    chan error
}

func (s *channelStream) ID() e2api.ChannelID {
	return s.channel.ID
}

func (s *channelStream) Channel() *e2api.Channel {
	return s.channel
}

func (s *channelStream) Indications() <-chan *e2appducontents.Ricindication {
	return s.buffer.Out()
}

func (s *channelStream) Done() <-chan error {
	return s.errs
}

func (s *channelStream) Close(err error) {
	if err != nil {
		s.errs <- err
	}
	s.manager.close(s)
	close(s.errs)
}
