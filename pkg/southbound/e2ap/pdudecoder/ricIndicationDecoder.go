// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicIndicationPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RanFunctionID, types.RicActionID, *types.RicCallProcessID,
	*types.RicIndicationHeader, *types.RicIndicationMessage,
	types.RicIndicationSn, e2apies.RicindicationType, *types.RicRequest, error) {

	if err := e2apPdu.Validate(); err != nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricIndication := e2apPdu.GetInitiatingMessage().GetValue().GetRicIndication()
	if ricIndication == nil {
		return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have RicIndication")
	}

	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID
	var ricActionID types.RicActionID
	var ricIndicationSn types.RicIndicationSn
	var ricIndicationType e2apies.RicindicationType
	var ricIndicationHeader types.RicIndicationHeader
	var ricIndicationMessage types.RicIndicationMessage
	var ricCallProcessID *types.RicCallProcessID

	for _, v := range ricIndication.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRicrequestId() == nil {
				return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRicrequestId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRicrequestId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRanfunctionId() == nil {
				return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID")
			}
			ranFunctionID = types.RanFunctionID(v.GetValue().GetRanfunctionId().GetValue())
		}
		if v.Id == int32(v2.ProtocolIeIDRicactionID) {
			ricActionIDIe := v.GetValue().GetRicactionId()
			if ricActionIDIe == nil {
				return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICactionID")
			}
			ricActionID = types.RicActionID(v.GetValue().GetRicactionId().GetValue())
		}
		if v.Id == int32(v2.ProtocolIeIDRicindicationSn) {
			// ricIndicationSn is optional
			if v.GetValue().GetRicindicationSn() != nil {
				ricIndicationSn = types.RicIndicationSn(v.GetValue().GetRicindicationSn().GetValue())
			}
		}
		if v.Id == int32(v2.ProtocolIeIDRicindicationType) {
			if v.GetValue().GetRicindicationType() < 0 || v.GetValue().GetRicindicationType() > 1 {
				return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICindicationType")
			}
			ricIndicationType = v.GetValue().GetRicindicationType()
		}
		if v.Id == int32(v2.ProtocolIeIDRicindicationHeader) {
			if v.GetValue().GetRicindicationHeader() == nil {
				return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICindicationHeader")
			}
			ricIndicationHeader = v.GetValue().GetRicindicationHeader().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDRicindicationMessage) {
			if v.GetValue().GetRicindicationMessage() == nil {
				return 0, 0, nil, nil, nil, 0, 0, nil, fmt.Errorf("error E2APpdu does not have id-RICindicationMessage")
			}
			ricIndicationMessage = v.GetValue().GetRicindicationMessage().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDRiccallProcessID) {
			if v.GetValue().GetRiccallProcessId() != nil { // Is optional
				ricCallProcessIDBytes := types.RicCallProcessID(v.GetValue().GetRiccallProcessId().GetValue())
				ricCallProcessID = &ricCallProcessIDBytes
			}
		}
	}

	return ranFunctionID, ricActionID, ricCallProcessID, &ricIndicationHeader,
		&ricIndicationMessage, ricIndicationSn, ricIndicationType, &ricRequestID, nil
}
