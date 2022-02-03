// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

import e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"

type Channel interface {
	ID() e2api.ChannelID
	Meta() e2api.ChannelMeta
	Input() Input
	Output() Output
}

func newChannel(id e2api.ChannelID, meta e2api.ChannelMeta, buffer *channelBuffer, manager *channelManager) Channel {
	c := &channel{
		manager: manager,
		buffer:  buffer,
		id:      id,
		meta:    meta,
	}
	c.input = newInput(c)
	c.output = newOutput(c)
	return c
}

type channel struct {
	manager *channelManager
	buffer  *channelBuffer
	id      e2api.ChannelID
	meta    e2api.ChannelMeta
	input   *channelInput
	output  *channelOutput
}

func (c *channel) ID() e2api.ChannelID {
	return c.id
}

func (c *channel) Meta() e2api.ChannelMeta {
	return c.meta
}

func (c *channel) Input() Input {
	return c.input
}

func (c *channel) Output() Output {
	return c.output
}

func (c *channel) open() {
	c.output.open()
}

func (c *channel) close(err error) {
	c.output.close(err)
}

func (c *channel) notify() {
	c.manager.notify(c)
}
