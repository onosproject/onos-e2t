// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

import (
	"encoding/binary"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

var connectionsByID = make(map[uint32]SctpConnection)

// CreateConnection logs creation of a connection by a setup request
func CreateConnection(setupRequest *e2ctypes.E2SetupRequestT) SctpConnection {
	id := binary.BigEndian.Uint32(setupRequest.ProtocolIEs.List[0].GetGlobalE2Node_ID().GetGNB().GetGlobalGNB_ID().GnbId.GetGnb_ID().BitString)
	plmnID := setupRequest.ProtocolIEs.List[0].GetGlobalE2Node_ID().GetGNB().GetGlobalGNB_ID().PlmnId

	// TODO - these values need to be extracted from the inbound connection
	ipAddress := "127.0.0.1"
	port := uint32(1234)

	connection := SctpConnection{
		ID:              id,
		PlmnID:          plmnID,
		RemoteIPAddress: ipAddress,
		RemotePort:      port,
	}
	connectionsByID[id] = connection
	return connection
}

// FindConnection looks up and returns a connection given its ID
func FindConnection(id uint32) SctpConnection {
	return connectionsByID[id]
}

// DeleteConnection removes a connection
func DeleteConnection(id uint32) {
	delete(connectionsByID, id)
}

// ListConnection returns an array of all of the active connections
func ListConnections() []SctpConnection {
	var connections []SctpConnection

	for _, value := range connectionsByID {
		connections = append(connections, value)
	}
	return connections
}
