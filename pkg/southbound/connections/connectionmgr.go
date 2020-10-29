// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

var connectionsByID = make(map[uint32]SctpConnection)

// CreateConnection logs creation of a connection by a setup request
func CreateConnection(setupRequest *e2appducontents.E2SetupRequest) SctpConnection {
	globalGNbID := setupRequest.ProtocolIes.E2ApProtocolIes3.Value.GetGlobalE2NodeId().(*e2apies.GlobalE2NodeId_GNb).GNb.GlobalGNbId
	globalID := uint32(globalGNbID.GnbId.GetGnbId().Value)
	plmnID := globalGNbID.PlmnId.String()

	// TODO - these values need to be extracted from the inbound connection
	ipAddress := "127.0.0.1"
	port := uint32(1234)

	connection := SctpConnection{
		ID:              globalID,
		PlmnID:          plmnID,
		RemoteIPAddress: ipAddress,
		RemotePort:      port,
	}
	connectionsByID[globalID] = connection
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
