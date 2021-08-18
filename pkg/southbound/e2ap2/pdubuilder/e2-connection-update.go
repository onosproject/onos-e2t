// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
)

func CreateE2connectionUpdateE2apPdu(trID int32, addItems []*types.E2ConnectionUpdateItem, modifyItems []*types.E2ConnectionUpdateItem,
	removeItems []*types.TnlInformation) (*e2appdudescriptions.E2ApPdu, error) {

	if addItems == nil && modifyItems == nil && removeItems == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
						InitiatingMessage: &e2appducontents.E2ConnectionUpdate{
							ProtocolIes: &e2appducontents.E2ConnectionUpdateIes{
								//E2ApProtocolIes44: &connectionAddList,    //E2 Connection Add List
								//E2ApProtocolIes45: &connectionModifyList, //E2 Connection Modify List
								//E2ApProtocolIes46: &connectionRemoveList, //E2 Connection Remove List
								E2ApProtocolIes49: &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2ap_ies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdE2ConnectionUpdate{
							Value: int32(v2beta1.ProcedureCodeIDE2connectionUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}
	if addItems != nil {
		connectionAddList := e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes44{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateAdd),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2ConnectionUpdateList{
				Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for _, addItem := range addItems {
			cai := &e2appducontents.E2ConnectionUpdateItemIes{
				Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &e2appducontents.E2ConnectionUpdateItem{
					TnlInformation: &e2ap_ies.Tnlinformation{
						TnlPort:    &addItem.TnlInformation.TnlPort,
						TnlAddress: &addItem.TnlInformation.TnlAddress,
					},
					TnlUsage: addItem.TnlUsage,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			}
			connectionAddList.Value.Value = append(connectionAddList.Value.Value, cai)
		}
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes44 = &connectionAddList
	}

	if modifyItems != nil {
		connectionModifyList := e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes45{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateModify),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2ConnectionUpdateList{
				Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for _, modifyItem := range modifyItems {
			cmi := &e2appducontents.E2ConnectionUpdateItemIes{
				Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &e2appducontents.E2ConnectionUpdateItem{
					TnlInformation: &e2ap_ies.Tnlinformation{
						TnlPort:    &modifyItem.TnlInformation.TnlPort,
						TnlAddress: &modifyItem.TnlInformation.TnlAddress,
					},
					TnlUsage: modifyItem.TnlUsage,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			}
			connectionModifyList.Value.Value = append(connectionModifyList.Value.Value, cmi)
		}
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes45 = &connectionModifyList
	}

	if removeItems != nil {
		connectionRemoveList := e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes46{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateRemove),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2ConnectionUpdateRemoveList{
				Value: make([]*e2appducontents.E2ConnectionUpdateRemoveItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for _, removeItem := range removeItems {
			cri := &e2appducontents.E2ConnectionUpdateRemoveItemIes{
				Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateRemoveItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &e2appducontents.E2ConnectionUpdateRemoveItem{
					TnlInformation: &e2ap_ies.Tnlinformation{
						TnlPort:    &removeItem.TnlPort,
						TnlAddress: &removeItem.TnlAddress,
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			}
			connectionRemoveList.Value.Value = append(connectionRemoveList.Value.Value, cri)
		}
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes46 = &connectionRemoveList
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
