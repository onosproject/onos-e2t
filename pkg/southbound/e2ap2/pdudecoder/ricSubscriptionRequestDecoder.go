// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicSubscriptionRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (
	types.RicRequest,
	types.RanFunctionID, types.RicEventDefintion,
	map[types.RicActionID]types.RicActionDef, error) {

	if err := e2apPdu.Validate(); err != nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	ricSubscription := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicSubscription()
	if ricSubscription == nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have RicSubscription")
	}

	ranFunctionIDIe := ricSubscription.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes5()
	if ranFunctionIDIe == nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionID (mandatory)")
	}
	ranFunctionID := types.RanFunctionID(ranFunctionIDIe.GetValue().GetValue())

	ricRequestIDIe := ricSubscription.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes29()
	if ricRequestIDIe == nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
	}
	ricRequestID := types.RicRequest{
		RequestorID: types.RicRequestorID(ricRequestIDIe.GetValue().GetRicRequestorId()),
		InstanceID:  types.RicInstanceID(ricRequestIDIe.GetValue().GetRicInstanceId()),
	}

	ricSubscriptionDetailsIe := ricSubscription.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes30()
	if ricSubscriptionDetailsIe == nil {
		return types.RicRequest{}, 0, nil, nil, fmt.Errorf("error E2APpdu does not have id-RICrequestID (mandatory)")
	}

	ricEventTriggerDef := ricSubscriptionDetailsIe.GetValue().GetRicEventTriggerDefinition().GetValue()

	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	for raID, actionToBeSetup := range ricSubscriptionDetailsIe.GetValue().GetRicActionToBeSetupList().GetValue() {
		ricActionsToBeSetup[types.RicActionID(raID)] = types.RicActionDef{
			RicActionID:         types.RicActionID(actionToBeSetup.GetValue().GetRicActionId().GetValue()),
			RicActionType:       actionToBeSetup.GetValue().GetRicActionType(),
			RicSubsequentAction: actionToBeSetup.GetValue().GetRicSubsequentAction().GetRicSubsequentActionType(),
			Ricttw:              actionToBeSetup.GetValue().GetRicSubsequentAction().GetRicTimeToWait(),
			RicActionDefinition: actionToBeSetup.GetValue().GetRicActionDefinition().GetValue(),
		}
	}

	return ricRequestID, ranFunctionID, ricEventTriggerDef, ricActionsToBeSetup, nil
}
