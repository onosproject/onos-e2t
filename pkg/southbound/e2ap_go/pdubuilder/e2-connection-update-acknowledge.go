// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap_go/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-commondatatypes"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-descriptions"
)

func CreateE2connectionUpdateAcknowledgeE2apPdu(trID int32) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDE2connectionUpdate),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
					SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{
						E2ConnectionUpdate: &e2appducontents.E2ConnectionUpdateAcknowledge{
							ProtocolIes: make([]*e2appducontents.E2ConnectionUpdateAckIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetSuccessfulOutcome().GetValue().GetE2ConnectionUpdate().SetTransactionID(trID)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}