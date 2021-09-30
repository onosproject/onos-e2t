// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
)

func CreateE2connectionUpdateAcknowledgeE2apPdu(trID int32) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
						SuccessfulOutcome: &e2appducontents.E2ConnectionUpdateAcknowledge{
							ProtocolIes: &e2appducontents.E2ConnectionUpdateAckIes{
								//E2ApProtocolIes39: &connectionSetup,       //E2 Connection Setup List
								//E2ApProtocolIes40: &connectionSetupFailed, //E2 Connection Setup Failed List
								E2ApProtocolIes49: &e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2ap_ies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdE2ConnectionUpdate{
							Value: int32(v2beta1.ProcedureCodeIDE2connectionUpdate),
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
