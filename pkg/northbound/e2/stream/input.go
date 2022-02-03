// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package stream

type Input interface {
	Open()
	Close(err error)
}

func newInput(c *channel) *channelInput {
	return &channelInput{
		channel: c,
	}
}

type channelInput struct {
	*channel
}

func (c *channelInput) Open() {
	c.channel.open()
}

func (c *channelInput) Close(err error) {
	c.channel.close(err)
}
