// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicIndicationPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RanFunctionID, types.RicActionID, *types.RicCallProcessID,
	*types.RicIndicationHeader, *types.RicIndicationMessage,
	types.RicIndicationSn, e2apies.RicindicationType, *types.RicRequest, error) {

	if err := e2apPdu.Validate(); err != nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricIndication := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicIndication()
	if ricIndication == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have RicIndication")
	}

	ranFunctionIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes5()
	if ranFunctionIe == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RanfunctionID")
	}
	ranFunctionID := types.RanFunctionID(ranFunctionIe.GetValue().GetValue())

	ricActionIDIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes15()
	if ricActionIDIe == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICactionID")
	}
	ricActionID := types.RicActionID(ricActionIDIe.GetValue().GetValue())

	var ricCallProcessID *types.RicCallProcessID
	ricCallProcessIDIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes20()
	if ricCallProcessIDIe != nil { // Is optional
		ricCallProcessIDBytes := types.RicCallProcessID(ricCallProcessIDIe.GetValue().GetValue())
		ricCallProcessID = &ricCallProcessIDBytes
	}

	ricIndicationHeaderIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes25()
	if ricIndicationHeaderIe == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICindicationHeader")
	}
	ricIndicationHeader := types.RicIndicationHeader(ricIndicationHeaderIe.GetValue().GetValue())

	ricIndicationMessageIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes26()
	if ricIndicationMessageIe == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICindicationMessage")
	}
	ricIndicationMessage := types.RicIndicationMessage(ricIndicationMessageIe.GetValue().GetValue())

	ricIndicationSnIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes27()
	// ricIndicationSn is optional
	var ricIndicationSn types.RicIndicationSn
	if ricIndicationSnIe != nil {
		ricIndicationSn = types.RicIndicationSn(ricIndicationSnIe.GetValue().GetValue())
	}

	ricIndicationTypeIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes28()
	if ricIndicationTypeIe == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICindicationType")
	}
	ricIndicationType := ricIndicationTypeIe.GetValue()

	// ricRequestID is optional
	var ricRequestID types.RicRequest
	ricRequestIDIe := ricIndication.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes29()
	if ricRequestIDIe != nil {
		ricRequestID = types.RicRequest{
			RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
			InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
		}
	}

	return ranFunctionID, ricActionID, ricCallProcessID, &ricIndicationHeader,
		&ricIndicationMessage, ricIndicationSn, ricIndicationType, &ricRequestID, nil
}
