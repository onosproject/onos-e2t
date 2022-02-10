// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicSubscriptionFailurePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RicRequest, *types.RanFunctionID, v2.ProcedureCodeT,
	e2ap_commondatatypes.Criticality, e2ap_commondatatypes.TriggeringMessage,
	*types.RicRequest, *e2apies.Cause, []*types.CritDiag, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscription := e2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscription()
	if ricSubscription == nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscription")
	}

	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID
	var cause *e2ap_ies.Cause

	var pc v2.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID types.RicRequest
	var diags []*types.CritDiag

	for _, v := range ricSubscription.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRrId() == nil {
				return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRfId() == nil {
				return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
			}
			ranFunctionID = types.RanFunctionID(v.GetValue().GetRfId().GetValue())
		}
		if v.Id == int32(v2.ProtocolIeIDCause) {
			cause = v.GetValue().GetC()
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

	return &ricRequestID, &ranFunctionID, pc, crit, tm, &critDiagRequestID, cause, diags, nil
}
