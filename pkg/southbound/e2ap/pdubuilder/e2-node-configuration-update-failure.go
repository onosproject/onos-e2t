// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
)

func CreateE2NodeConfigurationUpdateFailureE2apPdu(trID int32, c *e2ap_ies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDE2nodeConfigurationUpdate),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
					UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{
						E2NodeConfigurationUpdate: &e2appducontents.E2NodeConfigurationUpdateFailure{
							ProtocolIes: make([]*e2appducontents.E2NodeConfigurationUpdateFailureIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetE2NodeConfigurationUpdate().SetTransactionID(trID).SetCause(c)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
