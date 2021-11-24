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

func CreateRicServiceQueryE2apPdu(trID int32) (*e2appdudescriptions.E2ApPdu, error) {

	pIes := &e2appducontents.RicserviceQueryIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RicServiceQueryIe{
			RicServiceQueryIe: &e2appducontents.RicServiceQueryIe_Id{
				Id: &e2apies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	//pIes2 := &e2appducontents.RicserviceQueryIes{
	//	Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
	//	Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
	//	Value: &e2appducontents.RicServiceQueryIeValue{
	//		RicServiceQueryIeValue: &e2appducontents.RicServiceQueryIeValue_RfidList{
	//			RfidList: &e2appducontents.RanfunctionsIdListSingleContainer{
	//				//Value: nil,
	//			},
	//		},
	//	},
	//}

	//ie9 := &e2appducontents.RicserviceQueryIes_RicserviceQueryIes9{ //RAN functions Accepted List
	//	Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
	//	Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
	//	Value:       nil,
	//	//&e2appducontents.RanfunctionsIdList{
	//	//Value: nil, //make([]*e2appducontents.RanfunctionIdItemIes, 0),
	//	//},
	//	Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	//}

	//ie49 := &e2appducontents.RicserviceQueryIes_RicserviceQueryIes49{
	//	Id:          int32(v2.ProtocolIeIDTransactionID),
	//	Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
	//	Value: &e2apies.TransactionId{
	//		Value: trID,
	//	},
	//	Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	//}

	//pIes.E2ApProtocolIes9 = append(pIes.E2ApProtocolIes9, ie9)
	//pIes.E2ApProtocolIes49 = append(pIes.E2ApProtocolIes49, ie49)

	//rfIDie := &e2appducontents.RanfunctionIdItemIes{
	//	RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
	//		Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
	//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
	//		Value: &e2appducontents.RanfunctionIdItem{
	//			RanFunctionId: &e2apies.RanfunctionId{
	//				Value: 0,
	//			},
	//			RanFunctionRevision: &e2apies.RanfunctionRevision{
	//				Value: 0,
	//			},
	//		},
	//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	//	},
	//}
	//pIes.GetE2ApProtocolIes9().GetValue().Value = append(pIes.GetE2ApProtocolIes9().GetValue().Value, rfIDie)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicServiceQuery{
						RicServiceQuery: &e2appducontents.RicserviceQuery{
							ProtocolIes: make([]*e2appducontents.RicserviceQueryIes, 0),
						},
					},
				},
				ProcedureCode: int32(v2.ProcedureCodeIDRICserviceQuery),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
			},
		},
	}

	e2apPdu.GetInitiatingMessage().GetValue().GetRicServiceQuery().ProtocolIes = append(e2apPdu.GetInitiatingMessage().GetValue().GetRicServiceQuery().ProtocolIes, pIes)
	//e2apPdu.GetInitiatingMessage().GetValue().ProtocolIes = append(e2apPdu.GetInitiatingMessage().GetValue().ProtocolIes, pIes2)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
