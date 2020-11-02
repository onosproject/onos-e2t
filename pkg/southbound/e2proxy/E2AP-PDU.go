// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2proxy

import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

// NewRICsubscriptionRequest - create sub request
// Deprecated - use new proto structures from API instead
func NewRICsubscriptionRequest(requestorID int, ricInstanceID int, ranFunctionID int, eventTrigger []byte) *e2ctypes.E2AP_PDUT {
	ricSubscriptionRequest := e2ctypes.RICsubscriptionRequestT{
		ProtocolIEs: &e2ctypes.ProtocolIE_Container_1544P0T{
			List: make([]*e2ctypes.RICsubscriptionRequest_IEsT, 0),
		},
	}

	rsrIe1 := e2ctypes.RICsubscriptionRequest_IEsT{
		Id:          e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICrequestID,
		Criticality: e2ctypes.CriticalityT_Criticality_reject,
		Choice: &e2ctypes.RICsubscriptionRequest_IEsT_RICrequestID{
			RICrequestID: &e2ctypes.RICrequestIDT{
				RicRequestorID: int64(requestorID),
				RicInstanceID:  int64(ricInstanceID),
			},
		},
	}
	ricSubscriptionRequest.ProtocolIEs.List = append(ricSubscriptionRequest.ProtocolIEs.List, &rsrIe1)

	rsrIe2 := e2ctypes.RICsubscriptionRequest_IEsT{
		Id:          e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionID,
		Criticality: e2ctypes.CriticalityT_Criticality_reject,
		Choice: &e2ctypes.RICsubscriptionRequest_IEsT_RANfunctionID{
			RANfunctionID: int64(ranFunctionID),
		},
	}
	ricSubscriptionRequest.ProtocolIEs.List = append(ricSubscriptionRequest.ProtocolIEs.List, &rsrIe2)

	tbs1 := e2ctypes.RICaction_ToBeSetup_ItemIEsT{
		Id:          e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICaction_ToBeSetup_Item,
		Criticality: e2ctypes.CriticalityT_Criticality_reject,
		Choice: &e2ctypes.RICaction_ToBeSetup_ItemIEsT_RICaction_ToBeSetup_Item{
			RICaction_ToBeSetup_Item: &e2ctypes.RICaction_ToBeSetup_ItemT{
				RicActionID:         5,
				RicActionType:       e2ctypes.RICactionTypeT_RICactionType_report,
				RicActionDefinition: string([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}),
				RicSubsequentAction: &e2ctypes.RICsubsequentActionT{
					RicSubsequentActionType: e2ctypes.RICsubsequentActionTypeT_RICsubsequentActionType_continue,
					RicTimeToWait:           e2ctypes.RICtimeToWaitT_RICtimeToWait_w500ms,
				},
			},
		},
	}
	rsd1 := e2ctypes.RICsubscriptionDetailsT{
		RicAction_ToBeSetup_List: &e2ctypes.RICactions_ToBeSetup_ListT{
			List: []*e2ctypes.RICaction_ToBeSetup_ItemIEsT{&tbs1},
		},
		RicEventTriggerDefinition: eventTrigger,
	}

	rsrIe3 := e2ctypes.RICsubscriptionRequest_IEsT{
		Id:          e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICsubscriptionDetails,
		Criticality: e2ctypes.CriticalityT_Criticality_reject,
		Choice: &e2ctypes.RICsubscriptionRequest_IEsT_RICsubscriptionDetails{
			RICsubscriptionDetails: &rsd1,
		},
	}
	ricSubscriptionRequest.ProtocolIEs.List = append(ricSubscriptionRequest.ProtocolIEs.List, &rsrIe3)

	e2apPdu := e2ctypes.E2AP_PDUT{
		Choice: &e2ctypes.E2AP_PDUT_InitiatingMessage{
			InitiatingMessage: &e2ctypes.InitiatingMessageT{
				ProcedureCode: e2ctypes.ProcedureCodeT_ProcedureCode_id_RICsubscription,
				Criticality:   e2ctypes.CriticalityT_Criticality_reject,
				Choice: &e2ctypes.InitiatingMessageT_RICsubscriptionRequest{
					RICsubscriptionRequest: &ricSubscriptionRequest,
				},
			},
		},
	}

	return &e2apPdu
}

// GetE2apPduType - get the type of a E2AP_PDU
// Deprecated: Do not use.
func GetE2apPduType(e2apPdu *e2ctypes.E2AP_PDUT) (e2ctypes.ProcedureCodeT, error) {
	switch choice := e2apPdu.GetChoice().(type) {
	case *e2ctypes.E2AP_PDUT_InitiatingMessage:
		return choice.InitiatingMessage.GetProcedureCode(), nil
	case *e2ctypes.E2AP_PDUT_SuccessfulOutcome:
		return choice.SuccessfulOutcome.GetProcedureCode(), nil
	case *e2ctypes.E2AP_PDUT_UnsuccessfulOutcome:
		return choice.UnsuccessfulOutcome.GetProcedureCode(), nil
	default:
		return e2ctypes.ProcedureCodeT_ProcedureCode_id_dummy, fmt.Errorf("GetE2apPduType() unexpected type %T", choice)
	}
}
