// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
)

const mask20bitric = 0xFFFFF

func CreateRicSubscriptionRequestE2apPdu(ricReqID int32, ricInstanceID int32, ranFuncID int32, ricActionID int32, ricEventDef byte, ricActionDef byte) (*e2appdudescriptions.E2ApPdu, error) {

	if ricReqID|mask20bitric > mask20bitric {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricReqID)
	}
	if ricInstanceID|mask20bitric > mask20bitric {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricInstanceID)
	}

	ricRequestId := e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
		Value: &e2apies.RicrequestId{
			RicRequestorId: ricReqID,      // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  ricInstanceID, // sequence from e2ap-v01.00.asn1:1127
		},
	}

	ranFunctionId := e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes5{
		Value: &e2apies.RanfunctionId{
			Value: ranFuncID, // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
	}

	ricSubscriptionDetails := e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes30{
		Value: &e2appducontents.RicsubscriptionDetails{
			RicEventTriggerDefinition: &e2ap_commondatatypes.RiceventTriggerDefinition{
				Value: make([]byte, 3), // Octet String definition from e2ap-v01.00.asn1:337
			},
			RicActionToBeSetupList: &e2appducontents.RicactionsToBeSetupList{
				Value: make([]*e2appducontents.RicactionToBeSetupItemIes, 0),
			},
		},
	}
	ricSubscriptionDetails.Value.RicEventTriggerDefinition.Value = append(ricSubscriptionDetails.Value.RicEventTriggerDefinition.Value, ricEventDef)
	// ricEventDef value taken from e2ap-v01.00.asn1:1297

	ricActionListToSetup := e2appducontents.RicactionToBeSetupItemIes{
		Value: &e2appducontents.RicactionToBeSetupItem{
			RicActionId: &e2apies.RicactionId{
				Value: ricActionID, // range of Integer from e2ap-v01.00.asn1:1059, value from line 1283
			},
			RicActionType: e2apies.RicactionType_RICACTION_TYPE_POLICY,
			RicActionDefinition: &e2ap_commondatatypes.RicactionDefinition{
				Value: make([]byte, 3), // Octet String definition from e2ap-v01.00.asn1:1057
			},
			RicSubsequentAction: &e2apies.RicsubsequentAction{
				RicSubsequentActionType: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
				RicTimeToWait:           e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO,
			},
		},
	}
	ricActionListToSetup.Value.RicActionDefinition.Value = append(ricActionListToSetup.Value.RicActionDefinition.Value, ricActionDef)
	// ricEventDef value taken from e2ap-v01.00.asn1:1285
	ricSubscriptionDetails.Value.RicActionToBeSetupList.Value = append(ricSubscriptionDetails.Value.RicActionToBeSetupList.Value, &ricActionListToSetup)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicSubscription: &e2appdudescriptions.RicSubscription{
						InitiatingMessage: &e2appducontents.RicsubscriptionRequest{
							ProtocolIes: &e2appducontents.RicsubscriptionRequestIes{
								E2ApProtocolIes29: &ricRequestId,           //RIC request ID
								E2ApProtocolIes5:  &ranFunctionId,          //RAN function ID
								E2ApProtocolIes30: &ricSubscriptionDetails, // RIC subscription details
							},
						},
					},
				},
			},
		},
	}
	return &e2apPdu, nil
}
