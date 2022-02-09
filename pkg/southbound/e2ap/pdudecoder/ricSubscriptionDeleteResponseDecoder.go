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

func DecodeRicSubscriptionDeleteResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RanFunctionID, *types.RicRequest, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscriptionDelete := e2apPdu.GetSuccessfulOutcome().GetValue().GetRicSubscriptionDelete()
	if ricSubscriptionDelete == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscriptionDelete")
	}

	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID

	for _, v := range ricSubscriptionDelete.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRrId() == nil {
				return nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRfId() == nil {
				return nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
			}
			ranFunctionID = types.RanFunctionID(v.GetValue().GetRfId().GetValue())
		}
	}

	return &ranFunctionID, &ricRequestID, nil
}
