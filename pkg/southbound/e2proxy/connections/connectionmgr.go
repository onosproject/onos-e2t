// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package connections

import (
	"encoding/binary"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

var connectionsById = make(map[uint32]SctpConnection)

func CreateConnection(setupRequest *e2ctypes.E2SetupRequestT) {
	id := binary.BigEndian.Uint32(setupRequest.ProtocolIEs.List[0].GetGlobalE2Node_ID().GetGNB().GetGlobalGNB_ID().GnbId.GetGnb_ID().BitString)
	plmnId := setupRequest.ProtocolIEs.List[0].GetGlobalE2Node_ID().GetGNB().GetGlobalGNB_ID().PlmnId

	connection := SctpConnection{
		Id: id,
		PlmnId: plmnId,
	}
	connectionsById[id] = connection
}

func FindConnection(id uint32) SctpConnection {
	return connectionsById[id]
}

func DeleteConnection() {

}

func ListConnections() []SctpConnection {
	var connections []SctpConnection

	for _,value := range connectionsById {
		connections = append(connections, value)
	}
	return connections
}