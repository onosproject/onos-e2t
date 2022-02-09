// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func CreateRicControlAcknowledgeE2apPdu(ricReqID types.RicRequest, ranFuncID types.RanFunctionID) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDRICcontrol),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
					SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicControl{
						RicControl: &e2appducontents.RiccontrolAcknowledge{
							ProtocolIes: make([]*e2appducontents.RiccontrolAcknowledgeIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetSuccessfulOutcome().GetValue().GetRicControl().SetRicRequestID(ricReqID).SetRanFunctionID(&ranFuncID)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
