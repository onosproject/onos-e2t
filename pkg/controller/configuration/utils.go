// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package configuration

import (
	"encoding/binary"
	"net"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2apcommondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func createConnectionUpdateReq(ip string) *e2appducontents.E2ConnectionUpdate {
	connectionAddList := &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes44{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateAdd),
		Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.E2ConnectionUpdateList{
			Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2apcommondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	testIP := net.ParseIP(ip)

	portBytes := make([]byte, 2)
	port := uint16(36421)
	binary.BigEndian.PutUint16(portBytes, port)
	cai := &e2appducontents.E2ConnectionUpdateItemIes{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
		Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2ConnectionUpdateItem{
			TnlInformation: &e2apies.Tnlinformation{
				TnlPort: &asn1.BitString{
					Value: portBytes,
					Len:   16,
				},
				TnlAddress: &asn1.BitString{
					Value: testIP.To4(),
					Len:   32,
				},
			},
			TnlUsage: e2apies.Tnlusage_TNLUSAGE_BOTH,
		},
		Presence: int32(e2apcommondatatypes.Presence_PRESENCE_MANDATORY),
	}

	transactionID := &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes49{
		Id:          int32(v2beta1.ProtocolIeIDTransactionID),
		Criticality: int32(e2apcommondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.TransactionId{
			Value: 3,
		},
		Presence: int32(e2apcommondatatypes.Presence_PRESENCE_MANDATORY),
	}
	connectionAddList.Value.Value = append(connectionAddList.Value.Value, cai)
	connectionUpdateRequest := &e2appducontents.E2ConnectionUpdate{
		ProtocolIes: &e2appducontents.E2ConnectionUpdateIes{
			E2ApProtocolIes44: connectionAddList,
			E2ApProtocolIes49: transactionID,
		},
	}

	return connectionUpdateRequest
}
