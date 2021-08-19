// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
)

func CreateE2NodeConfigurationUpdateE2apPdu(trID int32) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
						InitiatingMessage: &e2appducontents.E2NodeConfigurationUpdate{
							ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateIes{
								E2ApProtocolIes49: &e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2ap_ies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
							Value: int32(v2beta1.ProcedureCodeIDE2nodeConfigurationUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
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

func CreateE2NodeComponentIDNgEnbDu(value int64) e2ap_ies.E2NodeComponentId {
	return e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeNgeNbDu{
			E2NodeComponentTypeNgeNbDu: &e2ap_ies.E2NodeComponentNgeNbDuId{
				NgEnbDuId: &e2ap_ies.NgenbDuId{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentConfigUpdateGnb(ng []byte, xn []byte, e1 []byte, f1 []byte, x2 []byte) e2ap_ies.E2NodeComponentConfigUpdate {
	return e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate{
			GNbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateGnb{
				NgApconfigUpdate: ng,
				XnApconfigUpdate: xn,
				E1ApconfigUpdate: e1,
				F1ApconfigUpdate: f1,
				X2ApconfigUpdate: x2,
			},
		},
	}
}

func CreateE2NodeComponentConfigUpdateEnb(ng []byte, xn []byte, w1 []byte, s1 []byte, x2 []byte) e2ap_ies.E2NodeComponentConfigUpdate {
	return e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_ENbconfigUpdate{
			ENbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateEnb{
				NgApconfigUpdate: ng,
				XnApconfigUpdate: xn,
				W1ApconfigUpdate: w1,
				S1ApconfigUpdate: s1,
				X2ApconfigUpdate: x2,
			},
		},
	}
}
