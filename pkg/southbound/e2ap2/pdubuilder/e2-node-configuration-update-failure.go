// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
)

func CreateE2NodeConfigurationUpdateFailureE2apPdu(trID int32, c *e2ap_ies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	cause := e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes1{
		Id:          int32(v2beta1.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       c,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
						UnsuccessfulOutcome: &e2appducontents.E2NodeConfigurationUpdateFailure{
							ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateFailureIes{
								E2ApProtocolIes1: &cause, //Cause
								//E2ApProtocolIes31: &timeToWait,             //E2 Connection Setup Failed List
								//E2ApProtocolIes2:  &criticalityDiagnostics, //E2 Connection Setup Failed List
								E2ApProtocolIes49: &e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2ap_ies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
							Value: int32(v2beta1.ProcedureCodeIDE2nodeConfigurationUpdate),
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
