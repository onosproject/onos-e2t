// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap_go/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-descriptions"
)

func CreateRicServiceQueryE2apPdu(trID int32) (*e2appdudescriptions.E2ApPduRicServiceQuery, error) {

	pIes := &e2appducontents.RicserviceQueryIes{
		//E2ApProtocolIes9: &ranFunctionsAccepted, //RAN functions Accepted List
		E2ApProtocolIes49: &e2appducontents.RicserviceQueryIes_RicserviceQueryIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2apies.TransactionId{
				Value: trID,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}

	e2apPdu := e2appdudescriptions.E2ApPduRicServiceQuery{
		E2ApPdu: &e2appdudescriptions.E2ApPduRicServiceQuery_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessageRicServiceQuery{
				Value: &e2appducontents.RicserviceQuery{
					ProtocolIes: make([]*e2appducontents.RicserviceQueryIes, 0),
				},
				ProcedureCode: int32(v2.ProcedureCodeIDRICserviceQuery),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
			},
		},
	}

	e2apPdu.GetInitiatingMessage().GetValue().ProtocolIes = append(e2apPdu.GetInitiatingMessage().GetValue().ProtocolIes, pIes)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
