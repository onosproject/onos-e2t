// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func CreateE2NodeConfigurationUpdateE2apPdu(e2nccul []*types.E2NodeComponentConfigUpdateItem) (*e2appdudescriptions.E2ApPdu, error) {

	if e2nccul == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	configUpdateList := e2appducontents.E2NodeComponentConfigUpdateList{
		Value: make([]*e2appducontents.E2NodeComponentConfigUpdateItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &e2appducontents.E2NodeComponentConfigUpdateItemIes{
			Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2NodeComponentConfigUpdateItem{
				E2NodeComponentType:         e2nccui.E2NodeComponentType,
				E2NodeComponentId:           e2nccui.E2NodeComponentID,
				E2NodeComponentConfigUpdate: &e2nccui.E2NodeComponentConfigUpdate,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		configUpdateList.Value = append(configUpdateList.Value, cui)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
						InitiatingMessage: &e2appducontents.E2NodeConfigurationUpdate{
							ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateIes{
								Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdate),
								Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
								Value:       &configUpdateList,
								Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
							},
						},
						ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
							Value: int32(v1beta2.ProcedureCodeIDE2nodeConfigurationUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
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

func CreateE2NodeComponentIDGnbCuUp(value int64) e2ap_ies.E2NodeComponentId {
	return e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp{
			E2NodeComponentTypeGnbCuUp: &e2ap_ies.E2NodeComponentGnbCuUpId{
				GNbCuUpId: &e2ap_ies.GnbCuUpId{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentIDGnbDu(value int64) e2ap_ies.E2NodeComponentId {
	return e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbDu{
			E2NodeComponentTypeGnbDu: &e2ap_ies.E2NodeComponentGnbDuId{
				GNbDuId: &e2ap_ies.GnbDuId{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentConfigUpdateGnb(ngAp string, xnAp string, e1Ap string, f1Ap string) e2ap_ies.E2NodeComponentConfigUpdate {
	return e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate{
			GNbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateGnb{
				NgApconfigUpdate: ngAp,
				XnApconfigUpdate: xnAp,
				E1ApconfigUpdate: e1Ap,
				F1ApconfigUpdate: f1Ap,
			},
		},
	}
}

func CreateE2NodeComponentConfigUpdateEnb(s1 string, x2 string) e2ap_ies.E2NodeComponentConfigUpdate {
	return e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_ENbconfigUpdate{
			ENbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateEnb{
				S1ApconfigUpdate: s1,
				X2ApconfigUpdate: x2,
			},
		},
	}
}

func CreateE2NodeComponentConfigUpdateEnGnb(x2 string) e2ap_ies.E2NodeComponentConfigUpdate {
	return e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_EnGNbconfigUpdate{
			EnGNbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateEngNb{
				X2ApconfigUpdate: x2,
			},
		},
	}
}

func CreateE2NodeComponentConfigUpdateNgEnb(ngAp string, xnAp string) e2ap_ies.E2NodeComponentConfigUpdate {
	return e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_NgENbconfigUpdate{
			NgENbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateNgeNb{
				NgApconfigUpdate: ngAp,
				XnApconfigUpdate: xnAp,
			},
		},
	}
}
