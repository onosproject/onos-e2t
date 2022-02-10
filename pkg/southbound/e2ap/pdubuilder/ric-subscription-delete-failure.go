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
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func CreateRicSubscriptionDeleteFailureE2apPdu(ricReq *types.RicRequest, ranFuncID types.RanFunctionID,
	cause *e2apies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDRICsubscriptionDelete),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures{
					UoValues: &e2appdudescriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{
						RicSubscriptionDelete: &e2appducontents.RicsubscriptionDeleteFailure{
							ProtocolIes: make([]*e2appducontents.RicsubscriptionDeleteFailureIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscriptionDelete().
		SetRicRequestID(ricReq).SetRanFunctionID(ranFuncID).SetCause(cause)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
