// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeE2nodeConfigurationUpdateFailurePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*e2ap_ies.Cause, *e2ap_ies.TimeToWait,
	*v1beta2.ProcedureCodeT, *e2ap_commondatatypes.Criticality, *e2ap_commondatatypes.TriggeringMessage, *types.RicRequest,
	[]*types.CritDiag, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2ncuf := e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate()
	if e2ncuf == nil {
		return nil, nil, nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdateFailure")
	}

	cause := e2ncuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes1().GetValue()
	ttw := e2ncuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes31().GetValue()

	critDiagnostics := e2ncuf.GetUnsuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes2()
	var pc v1beta2.ProcedureCodeT
	var crit e2ap_commondatatypes.Criticality
	var tm e2ap_commondatatypes.TriggeringMessage
	var critDiagRequestID types.RicRequest
	var diags []*types.CritDiag
	if critDiagnostics != nil { //It's optional
		pc = v1beta2.ProcedureCodeT(critDiagnostics.GetValue().GetProcedureCode().GetValue())
		crit = critDiagnostics.GetValue().GetProcedureCriticality()
		tm = critDiagnostics.GetValue().GetTriggeringMessage()
		critDiagRequestID = types.RicRequest{
			RequestorID: types.RicRequestorID(critDiagnostics.GetValue().GetRicRequestorId().GetRicRequestorId()),
			InstanceID:  types.RicInstanceID(critDiagnostics.GetValue().GetRicRequestorId().GetRicInstanceId()),
		}
		for _, ie := range critDiagnostics.GetValue().GetIEsCriticalityDiagnostics().GetValue() {
			diag := types.CritDiag{
				IECriticality: ie.IEcriticality,
				IEId:          v1beta2.ProtocolIeID(ie.GetIEId().GetValue()),
				TypeOfError:   ie.TypeOfError,
			}
			diags = append(diags, &diag)
		}
	}

	return cause, &ttw, &pc, &crit, &tm, &critDiagRequestID, diags, nil
}
