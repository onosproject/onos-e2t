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

func DecodeRicSubscriptionResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	*types.RanFunctionID, *types.RicRequest, []types.RicActionID, map[types.RicActionID]*e2apies.Cause, error) {

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscription := e2apPdu.GetSuccessfulOutcome().GetValue().GetRicSubscription()
	if ricSubscription == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscription")
	}

	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID

	ricActionsAdmitted := make([]types.RicActionID, 0)
	causes := make(map[types.RicActionID]*e2apies.Cause)

	for _, v := range ricSubscription.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRicrequestId() == nil {
				return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRicrequestId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRicrequestId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRanfunctionId() == nil {
				return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
			}
			ranFunctionID = types.RanFunctionID(v.GetValue().GetRanfunctionId().GetValue())
		}
		if v.Id == int32(v2.ProtocolIeIDRicactionsAdmitted) {
			if v.GetValue().GetRicactionsAdmitted() == nil {
				return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICactions-Admitted (mandatory)")
			}
			for _, actionAdmitted := range v.GetValue().GetRicactionsAdmitted().GetValue() {
				ricActionsAdmitted = append(ricActionsAdmitted, types.RicActionID(actionAdmitted.GetValue().GetRicactionAdmittedItem().GetRicActionId().GetValue()))
			}
		}
		if v.Id == int32(v2.ProtocolIeIDRicactionsNotAdmitted) {
			if v.GetValue().GetRicactionsNotAdmitted() == nil {
				return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICactions-NotAdmitted (mandatory)")
			}

			for _, ranai := range v.GetValue().GetRicactionsNotAdmitted().GetValue() {
				causes[types.RicActionID(ranai.GetValue().GetRicactionNotAdmittedItem().GetRicActionId().GetValue())] = ranai.GetValue().GetRicactionNotAdmittedItem().GetCause()
			}
		}
	}

	return &ranFunctionID, &ricRequestID, ricActionsAdmitted, causes, nil
}
