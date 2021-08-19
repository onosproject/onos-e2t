// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
)

func CreateRicServiceUpdateE2apPdu(trID int32) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						InitiatingMessage: &e2appducontents.RicserviceUpdate{
							ProtocolIes: &e2appducontents.RicserviceUpdateIes{
								//E2ApProtocolIes10: &ranFunctionsAddedList,    //RAN functions (added) List
								//E2ApProtocolIes11: &ranFunctionsDeletedList,  //RAN functions ID (deleted) List
								//E2ApProtocolIes12: &ranFunctionsModifiedList, //RAN functions (modified) List
								E2ApProtocolIes49: &e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2apies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
							Value: int32(v2beta1.ProcedureCodeIDRICserviceUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
