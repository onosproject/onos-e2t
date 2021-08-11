// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicSubscriptionDeleteResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RanFunctionID, *types.RicRequest, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscriptionDelete := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicSubscriptionDelete()
	if ricSubscriptionDelete == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscriptionDelete")
	}

	ranFunctionIDIe := ricSubscriptionDelete.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes5()
	if ranFunctionIDIe == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
	}
	ranFunctionID := types.RanFunctionID(ranFunctionIDIe.GetValue().GetValue())

	ricRequestIDIe := ricSubscriptionDelete.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes29()
	if ricRequestIDIe == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
	}
	ricRequestID := &types.RicRequest{
		RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
		InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
	}

	return &ranFunctionID, ricRequestID, nil
}
