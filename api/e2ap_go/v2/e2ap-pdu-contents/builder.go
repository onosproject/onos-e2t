// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2ap_pdu_contents

import (
	v2 "github.com/onosproject/onos-e2t/api/e2ap_go/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap_go/types"
)

func (m *E2SetupRequest) SetTransactionID(trID int32) *E2SetupRequest {

	ie := &E2SetupRequestIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupRequestIe{
			E2SetupRequestIe: &E2SetupRequestIe_TrId{
				TrId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupRequest) SetGlobalE2nodeID(e2NodeID *e2ap_ies.GlobalE2NodeId) *E2SetupRequest {

	ie := &E2SetupRequestIes{
		Id:          int32(v2.ProtocolIeIDGlobalE2nodeID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupRequestIe{
			E2SetupRequestIe: &E2SetupRequestIe_GE2NId{
				GE2NId: e2NodeID,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupRequest) SetRanFunctionsAdded(rfal types.RanFunctions) *E2SetupRequest {

	rfl := &RanfunctionsList{
		Value: make([]*RanfunctionItemIes, 0),
	}

	for id, ranFunctionID := range rfal {
		ranFunction := RanfunctionItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionItemIe{
				RanfunctionItemIe: &RanfunctionItemIe_Rfi{
					Rfi: &RanfunctionItem{
						RanFunctionId: &e2ap_ies.RanfunctionId{
							Value: int32(id),
						},
						RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
							Value: []byte(ranFunctionID.Description),
						},
						RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
							Value: int32(ranFunctionID.Revision),
						},
						RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
							Value: string(ranFunctionID.OID),
						},
					},
				},
			},
		}
		rfl.Value = append(rfl.Value, &ranFunction)
	}

	ie := &E2SetupRequestIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAdded),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupRequestIe{
			E2SetupRequestIe: &E2SetupRequestIe_Rfl{
				Rfl: rfl,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupRequest) SetE2nodeComponentConfigAddition(e2nccul []*types.E2NodeComponentConfigAdditionItem) *E2SetupRequest {
	configAdditionList := E2NodeComponentConfigAdditionList{
		Value: make([]*E2NodeComponentConfigAdditionItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &E2NodeComponentConfigAdditionItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigAdditionItemIe{
				E2NodeComponentConfigAdditionItemIe: &E2NodeComponentConfigAdditionItemIe_E2Nccui{
					E2Nccui: &E2NodeComponentConfigAdditionItem{
						E2NodeComponentInterfaceType: e2nccui.E2NodeComponentType,
						E2NodeComponentId:            e2nccui.E2NodeComponentID,
						E2NodeComponentConfiguration: &e2nccui.E2NodeComponentConfiguration,
					},
				},
			},
		}
		configAdditionList.Value = append(configAdditionList.Value, cui)
	}

	ie := &E2SetupRequestIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAddition),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupRequestIe{
			E2SetupRequestIe: &E2SetupRequestIe_E2Nccal{
				E2Nccal: &configAdditionList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

//func (m *E2SetupResponse) SetRanFunctionRejected(rfRejected types.RanFunctionCauses) *E2SetupResponse {
//	ranFunctionsRejected := E2SetupResponseIes_E2SetupResponseIes13{
//		Id:          int32(v2.ProtocolIeIDRanfunctionsRejected),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RanfunctionsIdcauseList{
//			Value: make([]*RanfunctionIdcauseItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for id, cause := range rfRejected {
//		rfIDcIIe := RanfunctionIdcauseItemIes{
//			RanFunctionIdcauseItemIes7: &RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
//				Id:          int32(v2.ProtocolIeIDRanfunctionIeCauseItem),
//				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//				Value: &RanfunctionIdcauseItem{
//					RanFunctionId: &e2ap_ies.RanfunctionId{
//						Value: int32(id),
//					},
//					Cause: &e2ap_ies.Cause{},
//				},
//				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//			},
//		}
//
//		switch cause.GetCause().(type) {
//		case *e2ap_ies.Cause_Misc:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Misc{
//				Misc: cause.GetMisc(),
//			}
//		case *e2ap_ies.Cause_Protocol:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Protocol{
//				Protocol: cause.GetProtocol(),
//			}
//		case *e2ap_ies.Cause_RicService:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicService{
//				RicService: cause.GetRicService(),
//			}
//		case *e2ap_ies.Cause_RicRequest:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicRequest{
//				RicRequest: cause.GetRicRequest(),
//			}
//		case *e2ap_ies.Cause_Transport:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Transport{
//				Transport: cause.GetTransport(),
//			}
//
//		default:
//			return m
//		}
//		ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, &rfIDcIIe)
//	}
//	m.GetProtocolIes().E2ApProtocolIes13 = &ranFunctionsRejected
//	return m
//}
//
//func (m *E2SetupResponse) SetRanFunctionAccepted(rfAccepted types.RanFunctionRevisions) *E2SetupResponse {
//	ranFunctionsAccepted := E2SetupResponseIes_E2SetupResponseIes9{
//		Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RanfunctionsIdList{
//			Value: make([]*RanfunctionIdItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for rfID, rfRevision := range rfAccepted {
//		rfIDiIe := RanfunctionIdItemIes{
//			RanFunctionIdItemIes6: &RanfunctionIdItemIes_RanfunctionIdItemIes6{
//				Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
//				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//				Value: &RanfunctionIdItem{
//					RanFunctionId: &e2ap_ies.RanfunctionId{
//						Value: int32(rfID),
//					},
//					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
//						Value: int32(rfRevision),
//					},
//				},
//				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//			},
//		}
//		ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, &rfIDiIe)
//	}
//	m.GetProtocolIes().E2ApProtocolIes9 = &ranFunctionsAccepted
//	return m
//}
//
//func (m *E2SetupFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2SetupFailure {
//	criticalityDiagnostics := E2SetupFailureIes_E2SetupFailureIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = &criticalityDiagnostics
//	return m
//}
//
//func (m *E2SetupFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2SetupFailure {
//	m.GetProtocolIes().E2ApProtocolIes31 = &E2SetupFailureIes_E2SetupFailureIes31{
//		Id:          int32(v2.ProtocolIeIDTimeToWait),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value:       ttw, // Could be any other value
//		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *E2SetupFailure) SetTnlInformation(tnl *e2ap_ies.Tnlinformation) *E2SetupFailure {
//	m.GetProtocolIes().E2ApProtocolIes48 = &E2SetupFailureIes_E2SetupFailureIes48{
//		Id:          int32(v2.ProtocolIeIDTNLinformation),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value:       tnl,
//		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}

func (m *E2NodeConfigurationUpdate) SetTransactionID(trID int32) *E2NodeConfigurationUpdate {

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_TrId{
				TrId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdate) SetGlobalE2nodeID(e2NodeID *e2ap_ies.GlobalE2NodeId) *E2NodeConfigurationUpdate {

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDGlobalE2nodeID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_Ge2NId{
				Ge2NId: e2NodeID,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdate) SetE2nodeComponentConfigAddition(e2nccul []*types.E2NodeComponentConfigAdditionItem) *E2NodeConfigurationUpdate {
	configAdditionList := E2NodeComponentConfigAdditionList{
		Value: make([]*E2NodeComponentConfigAdditionItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &E2NodeComponentConfigAdditionItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigAdditionItemIe{
				E2NodeComponentConfigAdditionItemIe: &E2NodeComponentConfigAdditionItemIe_E2Nccui{
					E2Nccui: &E2NodeComponentConfigAdditionItem{
						E2NodeComponentInterfaceType: e2nccui.E2NodeComponentType,
						E2NodeComponentId:            e2nccui.E2NodeComponentID,
						E2NodeComponentConfiguration: &e2nccui.E2NodeComponentConfiguration,
					},
				},
			},
		}
		configAdditionList.Value = append(configAdditionList.Value, cui)
	}

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAddition),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2Nccal{
				E2Nccal: &configAdditionList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdate) SetE2nodeComponentConfigUpdate(e2nccul []*types.E2NodeComponentConfigUpdateItem) *E2NodeConfigurationUpdate {
	configUpdateList := E2NodeComponentConfigUpdateList{
		Value: make([]*E2NodeComponentConfigUpdateItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &E2NodeComponentConfigUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigUpdateItemIe{
				E2NodeComponentConfigUpdateItemIe: &E2NodeComponentConfigUpdateItemIe_E2Nccui{
					E2Nccui: &E2NodeComponentConfigUpdateItem{
						E2NodeComponentInterfaceType: e2nccui.E2NodeComponentType,
						E2NodeComponentId:            e2nccui.E2NodeComponentID,
						E2NodeComponentConfiguration: &e2nccui.E2NodeComponentConfiguration,
					},
				},
			},
		}
		configUpdateList.Value = append(configUpdateList.Value, cui)
	}

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdate),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2Nccul{
				E2Nccul: &configUpdateList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdate) SetE2nodeComponentConfigRemoval(e2nccul []*types.E2NodeComponentConfigRemovalItem) *E2NodeConfigurationUpdate {
	configRemovalList := E2NodeComponentConfigRemovalList{
		Value: make([]*E2NodeComponentConfigRemovalItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &E2NodeComponentConfigRemovalItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigRemovalItemIe{
				E2NodeComponentConfigRemovalItemIe: &E2NodeComponentConfigRemovalItemIe_E2Nccri{
					E2Nccri: &E2NodeComponentConfigRemovalItem{
						E2NodeComponentInterfaceType: e2nccui.E2NodeComponentType,
						E2NodeComponentId:            e2nccui.E2NodeComponentID,
					},
				},
			},
		}
		configRemovalList.Value = append(configRemovalList.Value, cui)
	}

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemoval),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2Nccrl{
				E2Nccrl: &configRemovalList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdate) SetE2nodeTnlAssociationRemoval(e2ntnlar []*types.TnlAssociationRemovalItem) *E2NodeConfigurationUpdate {
	tnlar := E2NodeTnlassociationRemovalList{
		Value: make([]*E2NodeTnlassociationRemovalItemIes, 0),
	}

	for _, e2ntnlari := range e2ntnlar {
		cui := &E2NodeTnlassociationRemovalItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeTNLassociationRemovalItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeTnlassociationRemovalItemIe{
				E2NodeTnlassociationRemovalItemIe: &E2NodeTnlassociationRemovalItemIe_E2Ntnlai{
					E2Ntnlai: &E2NodeTnlassociationRemovalItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlAddress: &e2ntnlari.TnlInformation.TnlAddress,
							//TnlPort:    e2ntnlari.TnlInformation.TnlPort,
						},
						TnlInformationRic: &e2ap_ies.Tnlinformation{
							TnlAddress: &e2ntnlari.TnlInformationRic.TnlAddress,
							TnlPort:    e2ntnlari.TnlInformationRic.TnlPort,
						},
					},
				},
			},
		}
		if e2ntnlari.TnlInformation.TnlPort != nil {
			cui.GetValue().GetE2Ntnlai().GetTnlInformation().TnlPort = e2ntnlari.TnlInformation.TnlPort
		}
		if e2ntnlari.TnlInformationRic.TnlPort != nil {
			cui.GetValue().GetE2Ntnlai().GetTnlInformationRic().TnlPort = e2ntnlari.TnlInformationRic.TnlPort
		}
		tnlar.Value = append(tnlar.Value, cui)
	}

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2nodeTNLassociationRemoval),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2Ntnlarl{
				E2Ntnlarl: &tnlar,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateAcknowledge) SetTransactionID(trID int32) *E2NodeConfigurationUpdateAcknowledge {

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_TrId{
				TrId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateAcknowledge) SetE2nodeComponentConfigAdditionAck(e2nccual []*types.E2NodeComponentConfigAdditionAckItem) *E2NodeConfigurationUpdateAcknowledge {
	configAdditionAckList := E2NodeComponentConfigAdditionAckList{
		Value: make([]*E2NodeComponentConfigAdditionAckItemIes, 0),
	}

	for _, e2nccuai := range e2nccual {
		cuai := &E2NodeComponentConfigAdditionAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigAdditionAckItemIe{
				E2NodeComponentConfigAdditionAckItemIe: &E2NodeComponentConfigAdditionAckItemIe_E2Nccaai{
					E2Nccaai: &E2NodeComponentConfigAdditionAckItem{
						E2NodeComponentInterfaceType: e2nccuai.E2NodeComponentType,
						E2NodeComponentId:            e2nccuai.E2NodeComponentID,
						E2NodeComponentConfigurationAck: &e2ap_ies.E2NodeComponentConfigurationAck{
							UpdateOutcome: e2nccuai.E2NodeComponentConfigurationAck.UpdateOutcome,
							//FailureCause:  e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause,
						},
					},
				},
			},
		}
		if e2nccuai.E2NodeComponentID != nil {
			cuai.GetValue().GetE2Nccaai().E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigurationAck.FailureCause != nil {
			cuai.GetValue().GetE2Nccaai().E2NodeComponentConfigurationAck.FailureCause = e2nccuai.E2NodeComponentConfigurationAck.FailureCause
		}

		configAdditionAckList.Value = append(configAdditionAckList.Value, cuai)
	}

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_E2Nccaal{
				E2Nccaal: &configAdditionAckList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateAcknowledge) SetE2nodeComponentConfigUpdateAck(e2nccual []*types.E2NodeComponentConfigUpdateAckItem) *E2NodeConfigurationUpdateAcknowledge {
	configUpdateAckList := E2NodeComponentConfigUpdateAckList{
		Value: make([]*E2NodeComponentConfigUpdateAckItemIes, 0),
	}

	for _, e2nccuai := range e2nccual {
		cuai := &E2NodeComponentConfigUpdateAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigUpdateAckItemIe{
				E2NodeComponentConfigUpdateAckItemIe: &E2NodeComponentConfigUpdateAckItemIe_E2Nccuai{
					E2Nccuai: &E2NodeComponentConfigUpdateAckItem{
						E2NodeComponentInterfaceType: e2nccuai.E2NodeComponentType,
						E2NodeComponentId:            e2nccuai.E2NodeComponentID,
						E2NodeComponentConfigurationAck: &e2ap_ies.E2NodeComponentConfigurationAck{
							UpdateOutcome: e2nccuai.E2NodeComponentConfigurationAck.UpdateOutcome,
							//FailureCause:  e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause,
						},
					},
				},
			},
		}
		if e2nccuai.E2NodeComponentID != nil {
			cuai.GetValue().GetE2Nccuai().E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigurationAck.FailureCause != nil {
			cuai.GetValue().GetE2Nccuai().E2NodeComponentConfigurationAck.FailureCause = e2nccuai.E2NodeComponentConfigurationAck.FailureCause
		}

		configUpdateAckList.Value = append(configUpdateAckList.Value, cuai)
	}

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_E2Nccual{
				E2Nccual: &configUpdateAckList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateAcknowledge) SetE2nodeComponentConfigRemovalAck(e2nccual []*types.E2NodeComponentConfigRemovalAckItem) *E2NodeConfigurationUpdateAcknowledge {
	configRemovalAckList := E2NodeComponentConfigRemovalAckList{
		Value: make([]*E2NodeComponentConfigRemovalAckItemIes, 0),
	}

	for _, e2nccuai := range e2nccual {
		cuai := &E2NodeComponentConfigRemovalAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigRemovalAckItemIe{
				E2NodeComponentConfigRemovalAckItemIe: &E2NodeComponentConfigRemovalAckItemIe_E2Nccrai{
					E2Nccrai: &E2NodeComponentConfigRemovalAckItem{
						E2NodeComponentInterfaceType: e2nccuai.E2NodeComponentType,
						E2NodeComponentId:            e2nccuai.E2NodeComponentID,
						E2NodeComponentConfigurationAck: &e2ap_ies.E2NodeComponentConfigurationAck{
							UpdateOutcome: e2nccuai.E2NodeComponentConfigurationAck.UpdateOutcome,
							//FailureCause:  e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause,
						},
					},
				},
			},
		}
		if e2nccuai.E2NodeComponentID != nil {
			cuai.GetValue().GetE2Nccrai().E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigurationAck.FailureCause != nil {
			cuai.GetValue().GetE2Nccrai().E2NodeComponentConfigurationAck.FailureCause = e2nccuai.E2NodeComponentConfigurationAck.FailureCause
		}

		configRemovalAckList.Value = append(configRemovalAckList.Value, cuai)
	}

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_E2Nccral{
				E2Nccral: &configRemovalAckList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

//func (m *E2NodeConfigurationUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2NodeConfigurationUpdateFailure {
//	criticalityDiagnostics := E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = &criticalityDiagnostics
//	return m
//}
//
//func (m *E2NodeConfigurationUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2NodeConfigurationUpdateFailure {
//	m.GetProtocolIes().E2ApProtocolIes31 = &E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes31{
//		Id:          int32(v2.ProtocolIeIDTimeToWait),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value:       ttw, // Could be any other value
//		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}

func (m *E2ConnectionUpdateFailure) SetTransactionID(trID int32) *E2ConnectionUpdateFailure {

	ie := &E2ConnectionUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateFailureIe{
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_TrId{
				TrId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2ConnectionUpdateFailure {

	criticalityDiagnostics := &e2ap_ies.CriticalityDiagnostics{
		ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
			Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
		},
		TriggeringMessage:    failureTrigMsg,
		ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
		RicRequestorId: &e2ap_ies.RicrequestId{
			RicRequestorId: int32(reqID.RequestorID),
			RicInstanceId:  int32(reqID.InstanceID),
		},
	}

	if critDiags != nil {
		criticalityDiagnostics.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
		}

		for _, critDiag := range critDiags {
			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
				IEcriticality: critDiag.IECriticality,
				IEId: &e2ap_commondatatypes.ProtocolIeId{
					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
				},
				TypeOfError: critDiag.TypeOfError,
			}
			criticalityDiagnostics.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	ie := &E2ConnectionUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2ConnectionUpdateFailureIe{
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_Cd{
				Cd: criticalityDiagnostics,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2ConnectionUpdateFailure {

	ie := &E2ConnectionUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2ConnectionUpdateFailureIe{
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_Ttw{
				Ttw: ttw,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateFailure) SetCause(c *e2ap_ies.Cause) *E2ConnectionUpdateFailure {

	ie := &E2ConnectionUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateFailureIe{
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_C{
				C: c,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateAcknowledge) SetTransactionID(trID int32) *E2ConnectionUpdateAcknowledge {

	ie := &E2ConnectionUpdateAckIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateAckIe{
			E2ConnectionUpdateAckIe: &E2ConnectionUpdateAckIe_TrId{
				TrId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *E2ConnectionUpdateAcknowledge) SetE2ConnectionSetup(connSetup []*types.E2ConnectionUpdateItem) *E2ConnectionUpdateAcknowledge {

	ie := &E2ConnectionUpdateAckIes{
		Id:          int32(v2.ProtocolIeIDE2connectionSetup),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateAckIe{
			E2ConnectionUpdateAckIe: &E2ConnectionUpdateAckIe_E2Cul{
				E2Cul: &E2ConnectionUpdateList{
					Value: make([]*E2ConnectionUpdateItemIes, 0),
				},
			},
		},
	}

	for _, setupItem := range connSetup {
		si := &E2ConnectionUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateItemIe{
				E2ConnectionUpdateItemIe: &E2ConnectionUpdateItemIe_E2Curi{
					E2Curi: &E2ConnectionUpdateItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    setupItem.TnlInformation.TnlPort,
							TnlAddress: &setupItem.TnlInformation.TnlAddress,
						},
						TnlUsage: setupItem.TnlUsage,
					},
				},
			},
		}
		ie.GetValue().GetE2Cul().Value = append(ie.GetValue().GetE2Cul().Value, si)
	}
	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateAcknowledge) SetE2ConnectionSetupFailed(connSetFail []*types.E2ConnectionSetupFailedItem) *E2ConnectionUpdateAcknowledge {

	ie := &E2ConnectionUpdateAckIes{
		Id:          int32(v2.ProtocolIeIDE2connectionSetupFailed),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateAckIe{
			E2ConnectionUpdateAckIe: &E2ConnectionUpdateAckIe_E2Csfl{
				E2Csfl: &E2ConnectionSetupFailedList{
					Value: make([]*E2ConnectionSetupFailedItemIes, 0),
				},
			},
		},
	}

	for _, sfItem := range connSetFail {
		sfi := &E2ConnectionSetupFailedItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionSetupFailedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionSetupFailedItemIe{
				E2ConnectionSetupFailedItemIe: &E2ConnectionSetupFailedItemIe_E2Csfi{
					E2Csfi: &E2ConnectionSetupFailedItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    sfItem.TnlInformation.TnlPort,
							TnlAddress: &sfItem.TnlInformation.TnlAddress,
						},
						Cause: &sfItem.Cause,
					},
				},
			},
		}
		ie.GetValue().GetE2Csfl().Value = append(ie.GetValue().GetE2Csfl().Value, sfi)
	}
	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdate) SetTransactionID(trID int32) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_TrId{
				TrId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateAdd(addItems []*types.E2ConnectionUpdateItem) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateAdd),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_E2Cul{
				E2Cul: &E2ConnectionUpdateList{
					Value: make([]*E2ConnectionUpdateItemIes, 0),
				},
			},
		},
	}

	for _, addItem := range addItems {
		cai := &E2ConnectionUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateItemIe{
				E2ConnectionUpdateItemIe: &E2ConnectionUpdateItemIe_E2Curi{
					E2Curi: &E2ConnectionUpdateItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    addItem.TnlInformation.TnlPort,
							TnlAddress: &addItem.TnlInformation.TnlAddress,
						},
						TnlUsage: addItem.TnlUsage,
					},
				},
			},
		}
		ie.GetValue().GetE2Cul().Value = append(ie.GetValue().GetE2Cul().Value, cai)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateModify(modifyItems []*types.E2ConnectionUpdateItem) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateModify),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_E2Cul{
				E2Cul: &E2ConnectionUpdateList{
					Value: make([]*E2ConnectionUpdateItemIes, 0),
				},
			},
		},
	}

	for _, modifyItem := range modifyItems {
		cai := &E2ConnectionUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateItemIe{
				E2ConnectionUpdateItemIe: &E2ConnectionUpdateItemIe_E2Curi{
					E2Curi: &E2ConnectionUpdateItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    modifyItem.TnlInformation.TnlPort,
							TnlAddress: &modifyItem.TnlInformation.TnlAddress,
						},
						TnlUsage: modifyItem.TnlUsage,
					},
				},
			},
		}
		ie.GetValue().GetE2Cul().Value = append(ie.GetValue().GetE2Cul().Value, cai)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateRemove(removeItems []*types.TnlInformation) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemove),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_E2Curl{
				E2Curl: &E2ConnectionUpdateRemoveList{
					Value: make([]*E2ConnectionUpdateRemoveItemIes, 0),
				},
			},
		},
	}

	for _, removeItem := range removeItems {
		cai := &E2ConnectionUpdateRemoveItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemoveItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateRemoveItemIe{
				E2ConnectionUpdateRemoveItemIe: &E2ConnectionUpdateRemoveItemIe_E2Curi{
					E2Curi: &E2ConnectionUpdateRemoveItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    removeItem.TnlPort,
							TnlAddress: &removeItem.TnlAddress,
						},
					},
				},
			},
		}
		ie.GetValue().GetE2Curl().Value = append(ie.GetValue().GetE2Curl().Value, cai)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

//func (m *ErrorIndication) SetTransactionID(trID int32) *ErrorIndication {
//	m.GetProtocolIes().E2ApProtocolIes49 = &ErrorIndicationIes_ErrorIndicationIes49{
//		Id:          int32(v2.ProtocolIeIDTransactionID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_ies.TransactionId{
//			Value: trID,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *ErrorIndication) SetRicRequestID(ricReqID *types.RicRequest) *ErrorIndication {
//	m.GetProtocolIes().E2ApProtocolIes29 = &ErrorIndicationIes_ErrorIndicationIes29{
//		Id:          int32(v2.ProtocolIeIDRicrequestID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_ies.RicrequestId{
//			RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
//			RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *ErrorIndication) SetRanFunctionID(ranFuncID *types.RanFunctionID) *ErrorIndication {
//	m.GetProtocolIes().E2ApProtocolIes5 = &ErrorIndicationIes_ErrorIndicationIes5{
//		Id:          int32(v2.ProtocolIeIDRanfunctionID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_ies.RanfunctionId{
//			Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *ErrorIndication) SetCause(cause *e2ap_ies.Cause) *ErrorIndication {
//	m.GetProtocolIes().E2ApProtocolIes1 = &ErrorIndicationIes_ErrorIndicationIes1{
//		Id:          int32(v2.ProtocolIeIDCause),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value:       cause,
//		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *ErrorIndication) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *ErrorIndication {
//	criticalityDiagnostics := &ErrorIndicationIes_ErrorIndicationIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
//	return m
//}
//
//func (m *ResetResponse) SetCriticalityDiagnostics(failureProcCode v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *ResetResponse {
//	criticalityDiagnostics := &ResetResponseIes_ResetResponseIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
//	return m
//}
//
//func (m *RiccontrolAcknowledge) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolAcknowledge {
//	m.GetProtocolIes().E2ApProtocolIes20 = &RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes20{
//		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_commondatatypes.RiccallProcessId{
//			Value: ricCallPrID,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *RiccontrolAcknowledge) SetRicControlOutcome(ricCtrlOut types.RicControlOutcome) *RiccontrolAcknowledge {
//	m.GetProtocolIes().E2ApProtocolIes32 = &RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes32{
//		Id:          int32(v2.ProtocolIeIDRiccontrolOutcome),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_commondatatypes.RiccontrolOutcome{
//			Value: ricCtrlOut,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *RiccontrolFailure) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolFailure {
//	m.GetProtocolIes().E2ApProtocolIes20 = &RiccontrolFailureIes_RiccontrolFailureIes20{
//		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_commondatatypes.RiccallProcessId{
//			Value: ricCallPrID,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *RiccontrolFailure) SetRicControlOutcome(ricCtrlOut types.RicControlOutcome) *RiccontrolFailure {
//	m.GetProtocolIes().E2ApProtocolIes32 = &RiccontrolFailureIes_RiccontrolFailureIes32{
//		Id:          int32(v2.ProtocolIeIDRiccontrolOutcome),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_commondatatypes.RiccontrolOutcome{
//			Value: ricCtrlOut,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *RiccontrolRequest) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolRequest {
//	m.GetProtocolIes().E2ApProtocolIes20 = &RiccontrolRequestIes_RiccontrolRequestIes20{
//		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_commondatatypes.RiccallProcessId{
//			Value: ricCallPrID,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *RiccontrolRequest) SetRicControlAckRequest(ricCtrlAckRequest e2ap_ies.RiccontrolAckRequest) *RiccontrolRequest {
//	m.GetProtocolIes().E2ApProtocolIes21 = &RiccontrolRequestIes_RiccontrolRequestIes21{
//		Id:          int32(v2.ProtocolIeIDRiccontrolAckRequest),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value:       ricCtrlAckRequest,
//		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *Ricindication) SetRicIndicationSN(ricSn types.RicIndicationSn) *Ricindication {
//	m.GetProtocolIes().E2ApProtocolIes27 = &RicindicationIes_RicindicationIes27{
//		Id:          int32(v2.ProtocolIeIDRicindicationSn),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_ies.RicindicationSn{
//			Value: int32(ricSn),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *Ricindication) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *Ricindication {
//	m.GetProtocolIes().E2ApProtocolIes20 = &RicindicationIes_RicindicationIes20{
//		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &e2ap_commondatatypes.RiccallProcessId{
//			Value: ricCallPrID,
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}

func (m *RicserviceQuery) SetRanFunctionsAccepted(rfAccepted types.RanFunctionRevisions) *RicserviceQuery {

	ie := &RicserviceQueryIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceQueryIe{
			RicServiceQueryIe: &RicServiceQueryIe_RfidList{
				RfidList: &RanfunctionsIdListSingleContainer{
					Value: &RanfunctionsIdList{
						Value: make([]*RanfunctionIdItemIes, 0),
					},
				},
			},
		},
	}

	for rfID, rfRevision := range rfAccepted {
		rfIDiIe := &RanfunctionIdItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionIdItemIe{
				RanfunctionIdItemIe: &RanfunctionIdItemIe_RfId{
					RfId: &RanfunctionIdItem{
						RanFunctionId: &e2ap_ies.RanfunctionId{
							Value: int32(rfID),
						},
						RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
							Value: int32(rfRevision),
						},
					},
				},
			},
		}
		ie.GetValue().GetRfidList().GetValue().Value = append(ie.GetValue().GetRfidList().GetValue().Value, rfIDiIe)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceQuery) SetTransactionID(trID int32) *RicserviceQuery {

	ie := &RicserviceQueryIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceQueryIe{
			RicServiceQueryIe: &RicServiceQueryIe_Id{
				Id: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

//func (m *RicserviceUpdate) SetRanFunctionsAdded(rfal types.RanFunctions) *RicserviceUpdate {
//	ranFunctionsAddedList := RicserviceUpdateIes_RicserviceUpdateIes10{
//		Id:          int32(v2.ProtocolIeIDRanfunctionsAdded),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RanfunctionsList{
//			Value: make([]*RanfunctionItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for id, ranFunctionID := range rfal {
//		ranFunction := RanfunctionItemIes{
//			E2ApProtocolIes8: &RanfunctionItemIes_RanfunctionItemIes8{
//				Id:          int32(v2.ProtocolIeIDRanfunctionItem),
//				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//				Value: &RanfunctionItem{
//					RanFunctionId: &e2ap_ies.RanfunctionId{
//						Value: int32(id),
//					},
//					RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
//						Value: []byte(ranFunctionID.Description),
//					},
//					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
//						Value: int32(ranFunctionID.Revision),
//					},
//					RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
//						Value: string(ranFunctionID.OID),
//					},
//				},
//			},
//		}
//		ranFunctionsAddedList.Value.Value = append(ranFunctionsAddedList.Value.Value, &ranFunction)
//	}
//	m.GetProtocolIes().E2ApProtocolIes10 = &ranFunctionsAddedList
//	return m
//}
//
//func (m *RicserviceUpdate) SetRanFunctionsModified(rfml types.RanFunctions) *RicserviceUpdate {
//	ranFunctionsModifiedList := RicserviceUpdateIes_RicserviceUpdateIes12{
//		Id:          int32(v2.ProtocolIeIDRanfunctionsModified),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RanfunctionsList{
//			Value: make([]*RanfunctionItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for id, ranFunctionID := range rfml {
//		ranFunction := RanfunctionItemIes{
//			E2ApProtocolIes8: &RanfunctionItemIes_RanfunctionItemIes8{
//				Id:          int32(v2.ProtocolIeIDRanfunctionItem),
//				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//				Value: &RanfunctionItem{
//					RanFunctionId: &e2ap_ies.RanfunctionId{
//						Value: int32(id),
//					},
//					RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
//						Value: []byte(ranFunctionID.Description),
//					},
//					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
//						Value: int32(ranFunctionID.Revision),
//					},
//					RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
//						Value: string(ranFunctionID.OID),
//					},
//				},
//			},
//		}
//		ranFunctionsModifiedList.Value.Value = append(ranFunctionsModifiedList.Value.Value, &ranFunction)
//	}
//	m.GetProtocolIes().E2ApProtocolIes12 = &ranFunctionsModifiedList
//	return m
//}
//
//func (m *RicserviceUpdate) SetRanFunctionsDeleted(rfDeleted types.RanFunctionRevisions) *RicserviceUpdate {
//	ranFunctionsDeletedList := RicserviceUpdateIes_RicserviceUpdateIes11{
//		Id:          int32(v2.ProtocolIeIDRanfunctionsDeleted),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RanfunctionsIdList{
//			Value: make([]*RanfunctionIdItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for rfID, rfRevision := range rfDeleted {
//		rfIDiIe := RanfunctionIdItemIes{
//			RanFunctionIdItemIes6: &RanfunctionIdItemIes_RanfunctionIdItemIes6{
//				Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
//				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//				Value: &RanfunctionIdItem{
//					RanFunctionId: &e2ap_ies.RanfunctionId{
//						Value: int32(rfID),
//					},
//					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
//						Value: int32(rfRevision),
//					},
//				},
//				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//			},
//		}
//		ranFunctionsDeletedList.Value.Value = append(ranFunctionsDeletedList.Value.Value, &rfIDiIe)
//	}
//	m.GetProtocolIes().E2ApProtocolIes11 = &ranFunctionsDeletedList
//	return m
//}
//
//func (m *RicserviceUpdateAcknowledge) SetRanFunctionsRejected(rfRejected types.RanFunctionCauses) *RicserviceUpdateAcknowledge {
//	ranFunctionsRejected := RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes13{
//		Id:          int32(v2.ProtocolIeIDRanfunctionsRejected),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RanfunctionsIdcauseList{
//			Value: make([]*RanfunctionIdcauseItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for id, cause := range rfRejected {
//		rfIDcIIe := RanfunctionIdcauseItemIes{
//			RanFunctionIdcauseItemIes7: &RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
//				Id:          int32(v2.ProtocolIeIDRanfunctionIeCauseItem),
//				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//				Value: &RanfunctionIdcauseItem{
//					RanFunctionId: &e2ap_ies.RanfunctionId{
//						Value: int32(id),
//					},
//					Cause: &e2ap_ies.Cause{},
//				},
//				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//			},
//		}
//
//		switch cause.GetCause().(type) {
//		case *e2ap_ies.Cause_Misc:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Misc{
//				Misc: cause.GetMisc(),
//			}
//		case *e2ap_ies.Cause_Protocol:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Protocol{
//				Protocol: cause.GetProtocol(),
//			}
//		case *e2ap_ies.Cause_RicService:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicService{
//				RicService: cause.GetRicService(),
//			}
//		case *e2ap_ies.Cause_RicRequest:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicRequest{
//				RicRequest: cause.GetRicRequest(),
//			}
//		case *e2ap_ies.Cause_Transport:
//			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Transport{
//				Transport: cause.GetTransport(),
//			}
//
//		default:
//			return m
//		}
//		ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, &rfIDcIIe)
//	}
//	m.GetProtocolIes().E2ApProtocolIes13 = &ranFunctionsRejected
//	return m
//}
//
//func (m *RicserviceUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *RicserviceUpdateFailure {
//	m.GetProtocolIes().E2ApProtocolIes31 = &RicserviceUpdateFailureIes_RicserviceUpdateFailureIes31{
//		Id:          int32(v2.ProtocolIeIDTimeToWait),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value:       ttw, // Could be any other value
//		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//	return m
//}
//
//func (m *RicserviceUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicserviceUpdateFailure {
//	criticalityDiagnostics := &RicserviceUpdateFailureIes_RicserviceUpdateFailureIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
//	return m
//}
//
//func (m *RicsubscriptionDeleteFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicsubscriptionDeleteFailure {
//	criticalityDiagnostics := &RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
//	return m
//}
//
//func (m *RicsubscriptionResponse) SetRicActionNotAdmitted(ricActionsNotAccepted map[types.RicActionID]*e2ap_ies.Cause) *RicsubscriptionResponse {
//	ricActionNotAdmit := &RicsubscriptionResponseIes_RicsubscriptionResponseIes18{
//		Id:          int32(v2.ProtocolIeIDRicactionsNotAdmitted),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
//		Value: &RicactionNotAdmittedList{
//			Value: make([]*RicactionNotAdmittedItemIes, 0),
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	for ricActionID, cause := range ricActionsNotAccepted {
//		ranaIe := &RicactionNotAdmittedItemIes{
//			Id:          int32(v2.ProtocolIeIDRicactionNotAdmittedItem),
//			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//			Value: &RicactionNotAdmittedItem{
//				RicActionId: &e2ap_ies.RicactionId{
//					Value: int32(ricActionID),
//				},
//				Cause: cause,
//			},
//			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
//		}
//		ricActionNotAdmit.GetValue().Value = append(ricActionNotAdmit.GetValue().Value, ranaIe)
//	}
//	m.GetProtocolIes().E2ApProtocolIes18 = ricActionNotAdmit
//	return m
//}
//
//func (m *RicsubscriptionFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
//	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
//	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicsubscriptionFailure {
//	criticalityDiagnostics := &RicsubscriptionFailureIes_RicsubscriptionFailureIes2{
//		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
//		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
//		Value: &e2ap_ies.CriticalityDiagnostics{
//			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
//				Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
//			},
//			TriggeringMessage:    failureTrigMsg,
//			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
//			RicRequestorId: &e2ap_ies.RicrequestId{
//				RicRequestorId: int32(reqID.RequestorID),
//				RicInstanceId:  int32(reqID.InstanceID),
//			},
//		},
//		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
//	}
//
//	if critDiags != nil {
//		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
//			Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
//		}
//
//		for _, critDiag := range critDiags {
//			criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
//				IEcriticality: critDiag.IECriticality,
//				IEId: &e2ap_commondatatypes.ProtocolIeId{
//					Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
//				},
//				TypeOfError: critDiag.TypeOfError,
//			}
//			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
//		}
//	}
//
//	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
//	return m
//}
