// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
)

func CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu() (*e2appdudescriptions.E2ApPdu, error) {

	configUpdateAckList := e2appducontents.E2NodeComponentConfigUpdateAckList{
		Value: make([]*e2appducontents.E2NodeComponentConfigUpdateAckItemIes, 0),
	}

	cui := &e2appducontents.E2NodeComponentConfigUpdateAckItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2NodeComponentConfigUpdateAckItem{
			E2NodeComponentType: e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB, //ToDo - pass as a parameter
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId{ //ToDo - pass as a parameter
				E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp{
					E2NodeComponentTypeGnbCuUp: &e2ap_ies.E2NodeComponentGnbCuUpId{
						GNbCuUpId: &e2ap_ies.GnbCuUpId{
							Value: 21,
						},
					},
				},
			},
			E2NodeComponentConfigUpdateAck: &e2ap_ies.E2NodeComponentConfigUpdateAck{ //ToDo - pass as a parameter
				UpdateOutcome: 1,
				FailureCause: &e2ap_ies.Cause{ //ToDo - pass as a parameter
					Cause: &e2ap_ies.Cause_Protocol{
						Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
					},
				},
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}
	configUpdateAckList.Value = append(configUpdateAckList.Value, cui)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
						SuccessfulOutcome: &e2appducontents.E2NodeConfigurationUpdateAcknowledge{
							ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes{
								Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAck),
								Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
								Value:       &configUpdateAckList,
								Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
							},
						},
					},
				},
			},
		},
	}
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
