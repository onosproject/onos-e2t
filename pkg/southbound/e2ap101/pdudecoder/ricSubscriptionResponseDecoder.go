// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicSubscriptionResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RanFunctionID, *types.RicRequest, []types.RicActionID, *string, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscription := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicSubscription()
	if ricSubscription == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscription")
	}

	ranFunctionIDIe := ricSubscription.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes5()
	if ranFunctionIDIe == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
	}
	ranFunctionID := types.RanFunctionID(ranFunctionIDIe.GetValue().GetValue())

	ricRequestIDIe := ricSubscription.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes29()
	if ricRequestIDIe == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
	}
	ricRequestID := &types.RicRequest{
		RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
		InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
	}

	ricActionAdmittedIe := ricSubscription.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes17()
	if ricActionAdmittedIe == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICactions-Admitted (mandatory)")
	}
	ricActionsAdmitted := make([]types.RicActionID, 0)
	for _, actionAdmitted := range ricActionAdmittedIe.GetValue().GetValue() {
		ricActionsAdmitted = append(ricActionsAdmitted, types.RicActionID(actionAdmitted.GetValue().GetRicActionId().GetValue()))
	}

	return &ranFunctionID, ricRequestID, ricActionsAdmitted, nil, nil
}
