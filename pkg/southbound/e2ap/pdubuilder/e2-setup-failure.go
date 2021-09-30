// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SetupFailurePdu(trID int32, cause *e2apies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	errorCause := e2appducontents.E2SetupFailureIes_E2SetupFailureIes1{
		Id:          int32(v2beta1.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       cause, // Probably, could be any other reason
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						UnsuccessfulOutcome: &e2appducontents.E2SetupFailure{
							ProtocolIes: &e2appducontents.E2SetupFailureIes{
								E2ApProtocolIes1: &errorCause, // RIC Requestor & RIC Instance ID
								//E2ApProtocolIes31: &timeToWait,             // RAN function ID
								//E2ApProtocolIes2:  &criticalityDiagnostics, // CriticalityDiagnostics
								E2ApProtocolIes49: &e2appducontents.E2SetupFailureIes_E2SetupFailureIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2apies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdE2Setup{
							Value: int32(v2beta1.ProcedureCodeIDE2setup),
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

func CreateTnlInformation(tnlAddress *asn1.BitString) (*e2apies.Tnlinformation, error) {

	return &e2apies.Tnlinformation{
		TnlAddress: tnlAddress,
	}, nil
}
