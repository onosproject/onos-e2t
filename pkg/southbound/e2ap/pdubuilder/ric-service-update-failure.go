// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
)

func CreateRicServiceUpdateFailureE2apPdu(trID int32, cause *e2apies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	if cause == nil {
		return nil, fmt.Errorf("Cause was not passed - it is mandatory parameter to set")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDRICserviceUpdate),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
					UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{
						RicServiceUpdate: &e2appducontents.RicserviceUpdateFailure{
							ProtocolIes: make([]*e2appducontents.RicserviceUpdateFailureIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicServiceUpdate().SetTransactionID(trID).SetCause(cause)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
