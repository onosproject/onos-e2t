// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package configuration

import (
	"encoding/binary"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"net"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	e2server "github.com/onosproject/onos-e2t/pkg/southbound/e2ap/server"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2apcommondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func getConnToRemoveList(mgmtConn *e2server.ManagementConn, e2tInterfaces []*topoapi.Interface) []topoapi.Interface {
	var connToRemoveList []topoapi.Interface

	// TODO use E2 node config aspect
	/*for _, conn := range mgmtConn.E2NodeConfig.Connections {
		exist := false
		for _, e2tIface := range e2tInterfaces {
			if e2tIface.Type == topoapi.Interface_INTERFACE_E2AP200 {
				if conn.IP == e2tIface.IP && conn.Port == e2tIface.Port && e2tIface.Type == conn.Type {
					exist = true
				}
			}

		}
		if !exist {
			connToRemoveList = append(connToRemoveList, conn)
		}
	}*/
	return connToRemoveList
}

// getConnToAddList finds list of all missing connections that should be added
func getConnToAddList(mgmtConn *e2server.ManagementConn, e2tInterfaces []*topoapi.Interface) []topoapi.Interface {
	var connToAddList []topoapi.Interface
	if len(mgmtConn.E2NodeConfig.Connections) == 0 {
		log.Debugf("No configured connection is available, adding connections for all of the E2T instances")
		for _, e2tIface := range e2tInterfaces {
			if e2tIface.Type == topoapi.Interface_INTERFACE_E2AP200 {
				conn := topoapi.Interface{
					IP:   e2tIface.IP,
					Port: e2tIface.Port,
					Type: topoapi.Interface_INTERFACE_E2AP200,
				}
				connToAddList = append(connToAddList, conn)
			}
		}
		return connToAddList
	}

	for _, e2tIface := range e2tInterfaces {
		exist := false
		if e2tIface.Type == topoapi.Interface_INTERFACE_E2AP200 {
			for _, e2NodeConn := range mgmtConn.E2NodeConfig.Connections {
				if e2NodeConn.IP == e2tIface.IP && e2NodeConn.Port == e2tIface.Port && e2NodeConn.Type == e2tIface.Type {
					log.Debugf("Connection %+v already exists for e2node: %s", e2NodeConn, mgmtConn.E2NodeID)
					exist = true
				}
			}
			if !exist {
				conn := topoapi.Interface{
					IP:   e2tIface.IP,
					Port: e2tIface.Port,
					Type: topoapi.Interface_INTERFACE_E2AP200,
				}
				connToAddList = append(connToAddList, conn)
			}
		}
	}
	return connToAddList
}

func createConnectionAddListIE(connToAddList []topoapi.Interface) *e2appducontents.E2ConnectionUpdateIes {
	connectionAddList := &e2appducontents.E2ConnectionUpdateList{
		Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
	}

	for _, connToAdd := range connToAddList {
		parsedIP := net.ParseIP(connToAdd.IP)
		portBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(portBytes, uint16(connToAdd.Port))
		cai := &e2appducontents.E2ConnectionUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.E2ConnectionUpdateItemIe{
				E2ConnectionUpdateItemIe: &e2appducontents.E2ConnectionUpdateItemIe_E2Curi{
					E2Curi: &e2appducontents.E2ConnectionUpdateItem{
						TnlInformation: &e2apies.Tnlinformation{
							TnlPort: &asn1.BitString{
								Value: portBytes,
								Len:   16,
							},
							TnlAddress: &asn1.BitString{
								Value: parsedIP.To4(),
								Len:   32,
							},
						},
						TnlUsage: e2apies.Tnlusage_TNLUSAGE_BOTH,
					},
				},
			},
		}
		connectionAddList.Value = append(connectionAddList.Value, cai)
	}

	ie := &e2appducontents.E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateAdd),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &e2appducontents.E2ConnectionUpdateIe_E2Cul{
				E2Cul: connectionAddList,
			},
		},
	}

	return ie
}

func createConnectionRemoveList(connToRemoveList []topoapi.Interface) *e2appducontents.E2ConnectionUpdateIes {
	connectionRemoveList := &e2appducontents.E2ConnectionUpdateRemoveList{
		Value: make([]*e2appducontents.E2ConnectionUpdateRemoveItemIes, 0),
	}

	for _, connToRemove := range connToRemoveList {
		parsedIP := net.ParseIP(connToRemove.IP)
		portBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(portBytes, uint16(connToRemove.Port))
		cri := &e2appducontents.E2ConnectionUpdateRemoveItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemoveItem),
			Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.E2ConnectionUpdateRemoveItemIe{
				E2ConnectionUpdateRemoveItemIe: &e2appducontents.E2ConnectionUpdateRemoveItemIe_E2Curi{
					E2Curi: &e2appducontents.E2ConnectionUpdateRemoveItem{
						TnlInformation: &e2apies.Tnlinformation{
							TnlPort: &asn1.BitString{
								Value: portBytes,
								Len:   16,
							},
							TnlAddress: &asn1.BitString{
								Value: parsedIP.To4(),
								Len:   32,
							},
						},
					},
				},
			},
		}
		connectionRemoveList.Value = append(connectionRemoveList.Value, cri)
	}

	ie := &e2appducontents.E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemove),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &e2appducontents.E2ConnectionUpdateIe_E2Curl{
				E2Curl: connectionRemoveList,
			},
		},
	}

	return ie
}

// ConnectionUpdate connection update request
type ConnectionUpdate struct {
	transactionID        int32
	connectionAddList    *e2appducontents.E2ConnectionUpdateIes
	connectionRemoveList *e2appducontents.E2ConnectionUpdateIes
}

// NewConnectionUpdate creates a new instance of connection update request
func NewConnectionUpdate(options ...func(update *ConnectionUpdate)) *ConnectionUpdate {
	connUpdate := &ConnectionUpdate{}

	for _, option := range options {
		option(connUpdate)
	}
	return connUpdate
}

// WithTransactionID sets transaction ID
func WithTransactionID(transID int32) func(update *ConnectionUpdate) {
	return func(connUpdate *ConnectionUpdate) {
		connUpdate.transactionID = transID
	}
}

// WithConnectionAddList sets connection add list IE
func WithConnectionAddList(connectionAddList *e2appducontents.E2ConnectionUpdateIes) func(update *ConnectionUpdate) {
	return func(connUpdate *ConnectionUpdate) {
		connUpdate.connectionAddList = connectionAddList
	}
}

// WithConnectionRemoveList sets connection remove list IE
func WithConnectionRemoveList(connectionRemoveList *e2appducontents.E2ConnectionUpdateIes) func(update *ConnectionUpdate) {
	return func(connUpdate *ConnectionUpdate) {
		connUpdate.connectionRemoveList = connectionRemoveList
	}
}

// Build creates a E2 node connection update request
func (c *ConnectionUpdate) Build() *e2appducontents.E2ConnectionUpdate {

	connectionUpdateRequest := &e2appducontents.E2ConnectionUpdate{
		ProtocolIes: make([]*e2appducontents.E2ConnectionUpdateIes, 0),
	}
	connectionUpdateRequest.SetTransactionID(c.transactionID)

	if c.connectionAddList != nil && len(c.connectionAddList.GetValue().GetE2Cul().GetValue()) != 0 {
		connectionUpdateRequest.ProtocolIes = append(connectionUpdateRequest.ProtocolIes, c.connectionAddList)

	}
	if c.connectionRemoveList != nil && len(c.connectionRemoveList.GetValue().GetE2Curl().GetValue()) != 0 {
		connectionUpdateRequest.ProtocolIes = append(connectionUpdateRequest.ProtocolIes, c.connectionRemoveList)
	}

	return connectionUpdateRequest
}
