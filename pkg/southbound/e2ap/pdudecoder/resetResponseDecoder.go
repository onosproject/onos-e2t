// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeResetResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*int32, *v2.ProcedureCodeT,
	*e2ap_commondatatypes.Criticality, *e2ap_commondatatypes.TriggeringMessage,
	*types.RicRequest, []*types.CritDiag, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rr := e2apPdu.GetSuccessfulOutcome().GetValue().GetReset_()
	if rr == nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have ResetResponse")
	}

	var transactionID int32
	var pc v2.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID types.RicRequest
	var diags []*types.CritDiag
	for _, v := range rr.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDCriticalityDiagnostics) {
			critDiagnostics := v.GetValue().GetCd()
			if critDiagnostics != nil { //It's optional
				pc = v2.ProcedureCodeT(critDiagnostics.GetProcedureCode().GetValue())
				crit = critDiagnostics.GetProcedureCriticality()
				tm = critDiagnostics.GetTriggeringMessage()
				critDiagRequestID = types.RicRequest{
					RequestorID: types.RicRequestorID(critDiagnostics.GetRicRequestorId().GetRicRequestorId()),
					InstanceID:  types.RicInstanceID(critDiagnostics.GetRicRequestorId().GetRicInstanceId()),
				}
				for _, ie := range critDiagnostics.GetIEsCriticalityDiagnostics().GetValue() {
					diag := types.CritDiag{
						IECriticality: ie.IEcriticality,
						IEId:          v2.ProtocolIeID(ie.GetIEId().GetValue()),
						TypeOfError:   ie.TypeOfError,
					}
					diags = append(diags, &diag)
				}
			}
		}
	}

	return &transactionID, &pc, &crit, &tm, &critDiagRequestID, diags, nil
}
