// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicSubscriptionFailurePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RicRequest, *types.RanFunctionID, v1beta2.ProcedureCodeT,
	e2ap_commondatatypes.Criticality, e2ap_commondatatypes.TriggeringMessage,
	*types.RicRequest, map[types.RicActionID]*e2apies.Cause, []*types.CritDiag, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscription := e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicSubscription()
	if ricSubscription == nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscription")
	}

	ranFunctionIDIe := ricSubscription.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes5()
	if ranFunctionIDIe == nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
	}
	ranFunctionID := types.RanFunctionID(ranFunctionIDIe.GetValue().GetValue())

	ricRequestIDIe := ricSubscription.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes29()
	if ricRequestIDIe == nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
	}
	ricRequestID := types.RicRequest{
		RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
		InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
	}

	ricActionsNotAdmittedList := ricSubscription.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes18()
	if ricActionsNotAdmittedList == nil {
		return nil, nil, 0, 0, 0, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICactions-NotAdmitted (mandatory)")
	}
	causes := make(map[types.RicActionID]*e2apies.Cause)
	for _, ranai := range ricActionsNotAdmittedList.GetValue().GetValue() {
		causes[types.RicActionID(ranai.GetValue().GetRicActionId().GetValue())] = ranai.GetValue().GetCause()
	}

	critDiagnostics := ricSubscription.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2()
	var pc v1beta2.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID *types.RicRequest
	var diags []*types.CritDiag
	if critDiagnostics != nil { //It's optional
		pc = v1beta2.ProcedureCodeT(critDiagnostics.GetValue().GetProcedureCode().GetValue())
		crit = critDiagnostics.GetValue().GetProcedureCriticality()
		tm = critDiagnostics.GetValue().GetTriggeringMessage()
		critDiagRequestID = &types.RicRequest{
			RequestorID: types.RicRequestorID(critDiagnostics.GetValue().GetRicRequestorId().GetRicRequestorId()),
			InstanceID:  types.RicInstanceID(critDiagnostics.GetValue().GetRicRequestorId().GetRicInstanceId()),
		}
		// TODO: handle Diags
	}

	return &ricRequestID, &ranFunctionID, pc, crit, tm, critDiagRequestID, causes, diags, nil
}
