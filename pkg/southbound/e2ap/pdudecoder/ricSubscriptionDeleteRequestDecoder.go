// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicSubscriptionDeleteRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RicRequest, types.RanFunctionID, error) {

	if err := e2apPdu.Validate(); err != nil {
		return types.RicRequest{}, 0, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscriptionDelete := e2apPdu.GetInitiatingMessage().GetValue().GetRicSubscriptionDelete()
	if ricSubscriptionDelete == nil {
		return types.RicRequest{}, 0, fmt.Errorf("error E2APpdu does not have RicSubscriptionDelete")
	}

	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID

	for _, v := range ricSubscriptionDelete.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRicrequestId() == nil {
				return types.RicRequest{}, 0, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRicrequestId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRicrequestId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRanfunctionId() == nil {
				return types.RicRequest{}, 0, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
			}
			ranFunctionID = types.RanFunctionID(v.GetValue().GetRanfunctionId().GetValue())
		}
	}

	return ricRequestID, ranFunctionID, nil
}
