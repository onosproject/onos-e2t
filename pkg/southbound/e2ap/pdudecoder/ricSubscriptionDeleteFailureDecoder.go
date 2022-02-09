// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicSubscriptionDeleteFailurePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RanFunctionID, *types.RicRequest, *e2ap_ies.Cause,
	*v2.ProcedureCodeT, *e2ap_commondatatypes.Criticality, *e2ap_commondatatypes.TriggeringMessage, *types.RicRequest,
	[]*types.CritDiag, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscriptionDelete := e2apPdu.GetUnsuccessfulOutcome().GetValue().GetRicSubscriptionDelete()
	if ricSubscriptionDelete == nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscriptionDelete")
	}

	var cause *e2ap_ies.Cause
	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID

	var pc v2.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID types.RicRequest
	var diags []*types.CritDiag

	for _, v := range ricSubscriptionDelete.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRrId() == nil {
				return nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRfId() == nil {
				return nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
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

	return &ranFunctionID, &ricRequestID, cause, &pc, &crit, &tm, &critDiagRequestID, diags, nil
}
