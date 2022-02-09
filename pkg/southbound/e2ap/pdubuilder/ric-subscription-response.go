// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

const mask20bitricResponse = 0xFFFFF

func CreateRicSubscriptionResponseE2apPdu(ricReq *types.RicRequest, ranFuncID types.RanFunctionID,
	ricActionsAdmitted []*types.RicActionID) (*e2appdudescriptions.E2ApPdu, error) {

	if ricReq.RequestorID|mask20bitricResponse > mask20bitricResponse {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricReq.RequestorID)
	}
	if ricReq.InstanceID|mask20bitricResponse > mask20bitricResponse {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricReq.InstanceID)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDRICsubscription),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
					SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{
						RicSubscription: &e2appducontents.RicsubscriptionResponse{
							ProtocolIes: make([]*e2appducontents.RicsubscriptionResponseIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetSuccessfulOutcome().GetValue().GetRicSubscription().
		SetRicRequestID(ricReq).SetRanFunctionID(&ranFuncID).SetRicActionAdmitted(ricActionsAdmitted)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
