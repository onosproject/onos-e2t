// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicSubscriptionDeleteRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RicRequest, types.RanFunctionID, error) {

	if err := e2apPdu.Validate(); err != nil {
		return types.RicRequest{}, 0, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscriptionDelete := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicSubscriptionDelete()
	if ricSubscriptionDelete == nil {
		return types.RicRequest{}, 0, fmt.Errorf("error E2APpdu does not have RicSubscriptionDelete")
	}

	ranFunctionIDIe := ricSubscriptionDelete.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes5()
	if ranFunctionIDIe == nil {
		return types.RicRequest{}, 0, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
	}
	ranFunctionID := types.RanFunctionID(ranFunctionIDIe.GetValue().GetValue())

	ricRequestIDIe := ricSubscriptionDelete.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes29()
	if ricRequestIDIe == nil {
		return types.RicRequest{}, 0, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
	}
	ricRequestID := types.RicRequest{
		RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
		InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
	}

	return ricRequestID, ranFunctionID, nil
}
