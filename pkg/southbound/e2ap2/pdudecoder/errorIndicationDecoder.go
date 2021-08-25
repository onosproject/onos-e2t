// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
)

func DecodeErrorIndicationPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*int32, *e2ap_ies.Cause, *types.RanFunctionID,
	*types.RicRequest, *v2beta1.ProcedureCodeT, *e2ap_commondatatypes.Criticality, *e2ap_commondatatypes.TriggeringMessage, *types.RicRequest,
	[]*types.CritDiag, error) {
	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	ei := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetErrorIndication()
	if ei == nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have ErrorIndication")
	}

	transactionID := ei.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue()
	cause := ei.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes1().GetValue()

	ricRequestIDIe := ei.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes29()
	ricRequestID := &types.RicRequest{
		RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
		InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
	}

	ranFunctionIDIe := ei.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes5()
	ranFunctionID := types.RanFunctionID(ranFunctionIDIe.GetValue().GetValue())

	critDiagnostics := ei.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes2()
	var pc v2beta1.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID types.RicRequest
	var diags []*types.CritDiag
	if critDiagnostics != nil { //It's optional
		pc = v2beta1.ProcedureCodeT(critDiagnostics.GetValue().GetProcedureCode().GetValue())
		crit = critDiagnostics.GetValue().GetProcedureCriticality()
		tm = critDiagnostics.GetValue().GetTriggeringMessage()
		critDiagRequestID = types.RicRequest{
			RequestorID: types.RicRequestorID(critDiagnostics.GetValue().GetRicRequestorId().GetRicRequestorId()),
			InstanceID:  types.RicInstanceID(critDiagnostics.GetValue().GetRicRequestorId().GetRicInstanceId()),
		}
		for _, ie := range critDiagnostics.GetValue().GetIEsCriticalityDiagnostics().GetValue() {
			diag := types.CritDiag{
				IECriticality: ie.IEcriticality,
				IEId:          v2beta1.ProtocolIeID(ie.GetIEId().GetValue()),
				TypeOfError:   ie.TypeOfError,
			}
			diags = append(diags, &diag)
		}
	}

	return &transactionID, cause, &ranFunctionID, ricRequestID, &pc, &crit, &tm, &critDiagRequestID, diags, nil
}
