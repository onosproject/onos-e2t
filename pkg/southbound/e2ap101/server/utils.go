// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package server

import (
	"github.com/google/uuid"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdudecoder"
)

func GetNodeID(nodeID []byte, length int) (topoapi.ID, error) {

	e2NodeID := pdudecoder.GetE2NodeID(nodeID, length)

	e2NodeTopoID := topoapi.ID(e2NodeID)
	return e2NodeTopoID, nil
}

func getChannelID(deviceID topoapi.ID) (ChannelID, error) {
	bs := make([]byte, 16)
	copy(bs, deviceID)
	id, err := uuid.FromBytes(bs)
	if err != nil {
		return "", err
	}

	return ChannelID(id.String()), nil
}
