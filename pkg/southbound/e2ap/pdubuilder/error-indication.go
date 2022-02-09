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
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func CreateErrorIndicationE2apPduEmpty() *e2appdudescriptions.E2ApPdu {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDErrorIndication),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_ErrorIndication{
						ErrorIndication: &e2appducontents.ErrorIndication{
							ProtocolIes: make([]*e2appducontents.ErrorIndicationIes, 0),
						},
					},
				},
			},
		},
	}
	return &e2apPdu
}

func CreateErrorIndicationE2apPdu(trID *int32, ricReqID *types.RicRequest, ranFuncID *types.RanFunctionID,
	cause *e2apies.Cause, failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) (*e2appdudescriptions.E2ApPdu, error) {

	if ricReqID == nil && ranFuncID == nil && cause == nil && failureProcCode == nil && failureCrit == nil && failureTrigMsg == nil &&
		reqID == nil && critDiags == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDErrorIndication),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_ErrorIndication{
						ErrorIndication: &e2appducontents.ErrorIndication{
							ProtocolIes: make([]*e2appducontents.ErrorIndicationIes, 0),
						},
					},
				},
			},
		},
	}

	if trID != nil {
		e2apPdu.GetInitiatingMessage().GetValue().GetErrorIndication().SetTransactionID(*trID)
	}

	if cause != nil {
		e2apPdu.GetInitiatingMessage().GetValue().GetErrorIndication().SetCause(cause)
	}

	if ricReqID != nil {
		e2apPdu.GetInitiatingMessage().GetValue().GetErrorIndication().SetRicRequestID(ricReqID)
	}

	if ranFuncID != nil {
		e2apPdu.GetInitiatingMessage().GetValue().GetErrorIndication().SetRanFunctionID(ranFuncID)
	}

	if failureProcCode != nil && failureTrigMsg != nil && failureCrit != nil && reqID != nil {
		e2apPdu.GetInitiatingMessage().GetValue().GetErrorIndication().
			SetCriticalityDiagnostics(failureProcCode, failureCrit, failureTrigMsg, reqID, critDiags)
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
