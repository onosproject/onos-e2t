// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
)

func CreateRicServiceUpdateFailureE2apPdu(trID int32, cause *e2apies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	if cause == nil {
		return nil, fmt.Errorf("Cause was not passed - it is mandatory parameter to set")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						UnsuccessfulOutcome: &e2appducontents.RicserviceUpdateFailure{
							ProtocolIes: &e2appducontents.RicserviceUpdateFailureIes{
								E2ApProtocolIes1: &e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes1{
									Id:          int32(v2beta1.ProtocolIeIDCause),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value:       cause,
									Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
								//E2ApProtocolIes31: &timeToWait,             //Time to Wait
								//E2ApProtocolIes2:  &criticalityDiagnostics, //Criticality Diagnostics
								E2ApProtocolIes49: &e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2apies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
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
