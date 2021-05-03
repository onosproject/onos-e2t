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

func CreateE2NodeConfigurationUpdateE2apPdu() (*e2appdudescriptions.E2ApPdu, error) {

	configUpdateList := e2appducontents.E2NodeComponentConfigUpdateList{
		Value: make([]*e2appducontents.E2NodeComponentConfigUpdateItemIes, 0),
	}

	cui := &e2appducontents.E2NodeComponentConfigUpdateItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2NodeComponentConfigUpdateItem{
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
			E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate{ //ToDo - pass as a parameter
				E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate{
					GNbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateGnb{
						NgApconfigUpdate: "ng_AP",
						XnApconfigUpdate: "xn_AP",
						E1ApconfigUpdate: "e1_AP",
						F1ApconfigUpdate: "f1_AP",
					},
				},
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}
	configUpdateList.Value = append(configUpdateList.Value, cui)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
						InitiatingMessage: &e2appducontents.E2NodeConfigurationUpdate{
							ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateIes{
								Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdate),
								Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
								Value:       &configUpdateList,
								Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
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
