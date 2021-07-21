// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"testing"
)

func TestGetChannelID(t *testing.T) {
	channelID, err := getChannelID("0")
	t.Log(channelID, err)
}
