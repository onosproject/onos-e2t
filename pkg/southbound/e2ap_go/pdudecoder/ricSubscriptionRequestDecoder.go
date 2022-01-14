// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap_go/v2"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
)

func DecodeRicSubscriptionRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RicRequest,
	types.RanFunctionID, types.RicEventDefintion,
	map[types.RicActionID]types.RicActionDef, error) {

	if err := e2apPdu.Validate(); err != nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscription := e2apPdu.GetInitiatingMessage().GetValue().GetRicSubscription()
	if ricSubscription == nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscription")
	}

	var ricRequestID types.RicRequest
	var ranFunctionID types.RanFunctionID

	var ricEventTriggerDef []byte
	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)

	for _, v := range ricSubscription.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			if v.GetValue().GetRrId() == nil {
				return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}
			ricRequestID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			ricRequestID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicInstanceId())
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			if v.GetValue().GetRfId() == nil {
				return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
			}
			ranFunctionID = types.RanFunctionID(v.GetValue().GetRfId().GetValue())
		}
		if v.Id == int32(v2.ProtocolIeIDRicsubscriptionDetails) {
			ricSubscriptionDetailsIe := v.GetValue().GetRsd()
			if ricSubscriptionDetailsIe == nil {
				return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
			}

			ricEventTriggerDef = ricSubscriptionDetailsIe.GetRicEventTriggerDefinition().GetValue()

			for raID, actionToBeSetup := range ricSubscriptionDetailsIe.GetRicActionToBeSetupList().GetValue() {
				ricActionsToBeSetup[types.RicActionID(raID)] = types.RicActionDef{
					RicActionID:         types.RicActionID(actionToBeSetup.GetValue().GetRatbsi().GetRicActionId().GetValue()),
					RicActionType:       actionToBeSetup.GetValue().GetRatbsi().GetRicActionType(),
					RicSubsequentAction: actionToBeSetup.GetValue().GetRatbsi().GetRicSubsequentAction().GetRicSubsequentActionType(),
					Ricttw:              actionToBeSetup.GetValue().GetRatbsi().GetRicSubsequentAction().GetRicTimeToWait(),
					RicActionDefinition: actionToBeSetup.GetValue().GetRatbsi().GetRicActionDefinition().GetValue(),
				}
			}
		}
	}

	return ricRequestID, ranFunctionID, ricEventTriggerDef, ricActionsToBeSetup, nil
}
