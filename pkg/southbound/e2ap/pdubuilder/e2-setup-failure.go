// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SetupFailurePdu(trID int32, cause *e2apies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDE2setup),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
					UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup{
						E2Setup: &e2appducontents.E2SetupFailure{
							ProtocolIes: make([]*e2appducontents.E2SetupFailureIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetE2Setup().SetTransactionID(trID).SetErrorCause(cause)
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
