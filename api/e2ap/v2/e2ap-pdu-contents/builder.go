// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap_pdu_contents

import (
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func (m *E2SetupRequest) SetTransactionID(trID int32) *E2SetupRequest {

	ie := &E2SetupRequestIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupRequestIe{
			E2SetupRequestIe: &E2SetupRequestIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
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
			E2SetupRequestIe: &E2SetupRequestIe_GlobalE2NodeId{
				GlobalE2NodeId: e2NodeID,
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
				RanfunctionItemIe: &RanfunctionItemIe_RanfunctionItem{
					RanfunctionItem: &RanfunctionItem{
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
			E2SetupRequestIe: &E2SetupRequestIe_RanfunctionsAdded{
				RanfunctionsAdded: rfl,
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
				E2NodeComponentConfigAdditionItemIe: &E2NodeComponentConfigAdditionItemIe_E2NodeComponentConfigAdditionItem{
					E2NodeComponentConfigAdditionItem: &E2NodeComponentConfigAdditionItem{
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
			E2SetupRequestIe: &E2SetupRequestIe_E2NodeComponentConfigAddition{
				E2NodeComponentConfigAddition: &configAdditionList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupResponse) SetTransactionID(trID int32) *E2SetupResponse {

	ie := &E2SetupResponseIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupResponseIe{
			E2SetupResponseIe: &E2SetupResponseIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupResponse) SetGlobalRicID(plmnID types.PlmnID, ricID types.RicIdentifier) *E2SetupResponse {

	ie := &E2SetupResponseIes{
		Id:          int32(v2.ProtocolIeIDGlobalRicID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupResponseIe{
			E2SetupResponseIe: &E2SetupResponseIe_GlobalRicId{
				GlobalRicId: &e2ap_ies.GlobalRicId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
					},
					RicId: &asn1.BitString{
						Value: ricID.RicIdentifierValue,
						Len:   uint32(ricID.RicIdentifierLen),
					},
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupResponse) SetE2nodeComponentConfigAdditionAck(e2nccul []*types.E2NodeComponentConfigAdditionAckItem) *E2SetupResponse {
	configAdditionAckList := E2NodeComponentConfigAdditionAckList{
		Value: make([]*E2NodeComponentConfigAdditionAckItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &E2NodeComponentConfigAdditionAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigAdditionAckItemIe{
				E2NodeComponentConfigAdditionAckItemIe: &E2NodeComponentConfigAdditionAckItemIe_E2NodeComponentConfigAdditionAckItem{
					E2NodeComponentConfigAdditionAckItem: &E2NodeComponentConfigAdditionAckItem{
						E2NodeComponentInterfaceType:    e2nccui.E2NodeComponentType,
						E2NodeComponentId:               e2nccui.E2NodeComponentID,
						E2NodeComponentConfigurationAck: &e2nccui.E2NodeComponentConfigurationAck,
					},
				},
			},
		}
		configAdditionAckList.Value = append(configAdditionAckList.Value, cui)
	}

	ie := &E2SetupResponseIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupResponseIe{
			E2SetupResponseIe: &E2SetupResponseIe_E2NodeComponentConfigAdditionAck{
				E2NodeComponentConfigAdditionAck: &configAdditionAckList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupResponse) SetRanFunctionRejected(rfRejected types.RanFunctionCauses) *E2SetupResponse {

	rfrl := RanfunctionsIdcauseList{
		Value: make([]*RanfunctionIdcauseItemIes, 0),
	}

	for id, cause := range rfRejected {
		rfIDcIIe := RanfunctionIdcauseItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionIeCauseItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionIdcauseItemIe{
				RanfunctionIdcauseItemIe: &RanfunctionIdcauseItemIe_RanfunctionIecauseItem{
					RanfunctionIecauseItem: &RanfunctionIdcauseItem{
						RanFunctionId: &e2ap_ies.RanfunctionId{
							Value: int32(id),
						},
						Cause: &e2ap_ies.Cause{},
					},
				},
			},
		}

		switch cause.GetCause().(type) {
		case *e2ap_ies.Cause_Misc:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_Misc{
				Misc: cause.GetMisc(),
			}
		case *e2ap_ies.Cause_Protocol:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_Protocol{
				Protocol: cause.GetProtocol(),
			}
		case *e2ap_ies.Cause_RicService:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_RicService{
				RicService: cause.GetRicService(),
			}
		case *e2ap_ies.Cause_RicRequest:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_RicRequest{
				RicRequest: cause.GetRicRequest(),
			}
		case *e2ap_ies.Cause_Transport:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_Transport{
				Transport: cause.GetTransport(),
			}

		default:
			return m
		}
		rfrl.Value = append(rfrl.Value, &rfIDcIIe)
	}

	ie := &E2SetupResponseIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupResponseIe{
			E2SetupResponseIe: &E2SetupResponseIe_RanfunctionsRejected{
				RanfunctionsRejected: &rfrl,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupResponse) SetRanFunctionAccepted(rfAccepted types.RanFunctionRevisions) *E2SetupResponse {

	rfl := RanfunctionsIdList{
		Value: make([]*RanfunctionIdItemIes, 0),
	}

	for rfID, rfRevision := range rfAccepted {
		rfIDiIe := RanfunctionIdItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionIdItemIe{
				RanfunctionIdItemIe: &RanfunctionIdItemIe_RanfunctionIdItem{
					RanfunctionIdItem: &RanfunctionIdItem{
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
		rfl.Value = append(rfl.Value, &rfIDiIe)
	}

	ie := &E2SetupResponseIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupResponseIe{
			E2SetupResponseIe: &E2SetupResponseIe_RanfunctionsAccepted{
				RanfunctionsAccepted: &rfl,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupFailure) SetTransactionID(trID int32) *E2SetupFailure {

	ie := &E2SetupFailureIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2SetupFailureIe{
			E2SetupFailureIe: &E2SetupFailureIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupFailure) SetErrorCause(cause *e2ap_ies.Cause) *E2SetupFailure {

	ie := &E2SetupFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2SetupFailureIe{
			E2SetupFailureIe: &E2SetupFailureIe_Cause{
				Cause: cause,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2SetupFailure {

	ie := &E2SetupFailureIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2SetupFailureIe{
			E2SetupFailureIe: &E2SetupFailureIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2SetupFailure {

	ie := &E2SetupFailureIes{
		Id:          int32(v2.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2SetupFailureIe{
			E2SetupFailureIe: &E2SetupFailureIe_TimeToWait{
				TimeToWait: ttw,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2SetupFailure) SetTnlInformation(tnl *e2ap_ies.Tnlinformation) *E2SetupFailure {

	ie := &E2SetupFailureIes{
		Id:          int32(v2.ProtocolIeIDTNLinformation),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2SetupFailureIe{
			E2SetupFailureIe: &E2SetupFailureIe_TnlInformation{
				TnlInformation: tnl,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdate) SetTransactionID(trID int32) *E2NodeConfigurationUpdate {

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
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
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_GlobalE2NodeId{
				GlobalE2NodeId: e2NodeID,
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
				E2NodeComponentConfigAdditionItemIe: &E2NodeComponentConfigAdditionItemIe_E2NodeComponentConfigAdditionItem{
					E2NodeComponentConfigAdditionItem: &E2NodeComponentConfigAdditionItem{
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
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2NodeComponentConfigAddition{
				E2NodeComponentConfigAddition: &configAdditionList,
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
				E2NodeComponentConfigUpdateItemIe: &E2NodeComponentConfigUpdateItemIe_E2NodeComponentConfigUpdateItem{
					E2NodeComponentConfigUpdateItem: &E2NodeComponentConfigUpdateItem{
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
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2NodeComponentConfigUpdate{
				E2NodeComponentConfigUpdate: &configUpdateList,
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
				E2NodeComponentConfigRemovalItemIe: &E2NodeComponentConfigRemovalItemIe_E2NodeComponentConfigRemovalItem{
					E2NodeComponentConfigRemovalItem: &E2NodeComponentConfigRemovalItem{
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
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2NodeComponentConfigRemoval{
				E2NodeComponentConfigRemoval: &configRemovalList,
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
				E2NodeTnlassociationRemovalItemIe: &E2NodeTnlassociationRemovalItemIe_E2NodeTnlassociationRemovalItem{
					E2NodeTnlassociationRemovalItem: &E2NodeTnlassociationRemovalItem{
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
			cui.GetValue().GetE2NodeTnlassociationRemovalItem().GetTnlInformation().TnlPort = e2ntnlari.TnlInformation.TnlPort
		}
		if e2ntnlari.TnlInformationRic.TnlPort != nil {
			cui.GetValue().GetE2NodeTnlassociationRemovalItem().GetTnlInformationRic().TnlPort = e2ntnlari.TnlInformationRic.TnlPort
		}
		tnlar.Value = append(tnlar.Value, cui)
	}

	ie := &E2NodeConfigurationUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2nodeTNLassociationRemoval),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateIe{
			E2NodeConfigurationUpdateIe: &E2NodeConfigurationUpdateIe_E2NodeTnlassociationRemoval{
				E2NodeTnlassociationRemoval: &tnlar,
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
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
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
				E2NodeComponentConfigAdditionAckItemIe: &E2NodeComponentConfigAdditionAckItemIe_E2NodeComponentConfigAdditionAckItem{
					E2NodeComponentConfigAdditionAckItem: &E2NodeComponentConfigAdditionAckItem{
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
			cuai.GetValue().GetE2NodeComponentConfigAdditionAckItem().E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigurationAck.FailureCause != nil {
			cuai.GetValue().GetE2NodeComponentConfigAdditionAckItem().E2NodeComponentConfigurationAck.FailureCause = e2nccuai.E2NodeComponentConfigurationAck.FailureCause
		}

		configAdditionAckList.Value = append(configAdditionAckList.Value, cuai)
	}

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_E2NodeComponentConfigAdditionAck{
				E2NodeComponentConfigAdditionAck: &configAdditionAckList,
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
				E2NodeComponentConfigUpdateAckItemIe: &E2NodeComponentConfigUpdateAckItemIe_E2NodeComponentConfigUpdateAckItem{
					E2NodeComponentConfigUpdateAckItem: &E2NodeComponentConfigUpdateAckItem{
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
			cuai.GetValue().GetE2NodeComponentConfigUpdateAckItem().E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigurationAck.FailureCause != nil {
			cuai.GetValue().GetE2NodeComponentConfigUpdateAckItem().E2NodeComponentConfigurationAck.FailureCause = e2nccuai.E2NodeComponentConfigurationAck.FailureCause
		}

		configUpdateAckList.Value = append(configUpdateAckList.Value, cuai)
	}

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_E2NodeComponentConfigUpdateAck{
				E2NodeComponentConfigUpdateAck: &configUpdateAckList,
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
				E2NodeComponentConfigRemovalAckItemIe: &E2NodeComponentConfigRemovalAckItemIe_E2NodeComponentConfigRemovalAckItem{
					E2NodeComponentConfigRemovalAckItem: &E2NodeComponentConfigRemovalAckItem{
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
			cuai.GetValue().GetE2NodeComponentConfigRemovalAckItem().E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigurationAck.FailureCause != nil {
			cuai.GetValue().GetE2NodeComponentConfigRemovalAckItem().E2NodeComponentConfigurationAck.FailureCause = e2nccuai.E2NodeComponentConfigurationAck.FailureCause
		}

		configRemovalAckList.Value = append(configRemovalAckList.Value, cuai)
	}

	ie := &E2NodeConfigurationUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateAcknowledgeIe{
			E2NodeConfigurationUpdateAcknowledgeIe: &E2NodeConfigurationUpdateAcknowledgeIe_E2NodeComponentConfigRemovalAck{
				E2NodeComponentConfigRemovalAck: &configRemovalAckList,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2NodeConfigurationUpdateFailure {

	ie := &E2NodeConfigurationUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2NodeConfigurationUpdateFailureIe{
			E2NodeConfigurationUpdateFailureIe: &E2NodeConfigurationUpdateFailureIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateFailure) SetCause(c *e2ap_ies.Cause) *E2NodeConfigurationUpdateFailure {

	ie := &E2NodeConfigurationUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2NodeConfigurationUpdateFailureIe{
			E2NodeConfigurationUpdateFailureIe: &E2NodeConfigurationUpdateFailureIe_Cause{
				Cause: c,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2NodeConfigurationUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2NodeConfigurationUpdateFailure {

	ie := &E2NodeConfigurationUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &E2NodeConfigurationUpdateFailureIe{
			E2NodeConfigurationUpdateFailureIe: &E2NodeConfigurationUpdateFailureIe_TimeToWait{
				TimeToWait: ttw,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *E2NodeConfigurationUpdateFailure) SetTransactionID(trID int32) *E2NodeConfigurationUpdateFailure {

	ie := &E2NodeConfigurationUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2NodeConfigurationUpdateFailureIe{
			E2NodeConfigurationUpdateFailureIe: &E2NodeConfigurationUpdateFailureIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateFailure) SetTransactionID(trID int32) *E2ConnectionUpdateFailure {

	ie := &E2ConnectionUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateFailureIe{
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
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
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_CriticalityDiagnostics{
				CriticalityDiagnostics: criticalityDiagnostics,
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
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_TimeToWait{
				TimeToWait: ttw,
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
			E2ConnectionUpdateFailureIe: &E2ConnectionUpdateFailureIe_Cause{
				Cause: c,
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
			E2ConnectionUpdateAckIe: &E2ConnectionUpdateAckIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
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
			E2ConnectionUpdateAckIe: &E2ConnectionUpdateAckIe_E2ConnectionSetup{
				E2ConnectionSetup: &E2ConnectionUpdateList{
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
				E2ConnectionUpdateItemIe: &E2ConnectionUpdateItemIe_E2ConnectionUpdateItem{
					E2ConnectionUpdateItem: &E2ConnectionUpdateItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    setupItem.TnlInformation.TnlPort,
							TnlAddress: &setupItem.TnlInformation.TnlAddress,
						},
						TnlUsage: setupItem.TnlUsage,
					},
				},
			},
		}
		ie.GetValue().GetE2ConnectionSetup().Value = append(ie.GetValue().GetE2ConnectionSetup().Value, si)
	}
	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdateAcknowledge) SetE2ConnectionSetupFailed(connSetFail []*types.E2ConnectionSetupFailedItem) *E2ConnectionUpdateAcknowledge {

	ie := &E2ConnectionUpdateAckIes{
		Id:          int32(v2.ProtocolIeIDE2connectionSetupFailed),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateAckIe{
			E2ConnectionUpdateAckIe: &E2ConnectionUpdateAckIe_E2ConnectionSetupFailed{
				E2ConnectionSetupFailed: &E2ConnectionSetupFailedList{
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
				E2ConnectionSetupFailedItemIe: &E2ConnectionSetupFailedItemIe_E2ConnectionSetupFailedItem{
					E2ConnectionSetupFailedItem: &E2ConnectionSetupFailedItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    sfItem.TnlInformation.TnlPort,
							TnlAddress: &sfItem.TnlInformation.TnlAddress,
						},
						Cause: &sfItem.Cause,
					},
				},
			},
		}
		ie.GetValue().GetE2ConnectionSetupFailed().Value = append(ie.GetValue().GetE2ConnectionSetupFailed().Value, sfi)
	}
	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdate) SetTransactionID(trID int32) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
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
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_E2ConnectionUpdateAdd{
				E2ConnectionUpdateAdd: &E2ConnectionUpdateList{
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
				E2ConnectionUpdateItemIe: &E2ConnectionUpdateItemIe_E2ConnectionUpdateItem{
					E2ConnectionUpdateItem: &E2ConnectionUpdateItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    addItem.TnlInformation.TnlPort,
							TnlAddress: &addItem.TnlInformation.TnlAddress,
						},
						TnlUsage: addItem.TnlUsage,
					},
				},
			},
		}
		ie.GetValue().GetE2ConnectionUpdateAdd().Value = append(ie.GetValue().GetE2ConnectionUpdateAdd().Value, cai)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateModify(modifyItems []*types.E2ConnectionUpdateItem) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateModify),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_E2ConnectionUpdateModify{
				E2ConnectionUpdateModify: &E2ConnectionUpdateList{
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
				E2ConnectionUpdateItemIe: &E2ConnectionUpdateItemIe_E2ConnectionUpdateItem{
					E2ConnectionUpdateItem: &E2ConnectionUpdateItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    modifyItem.TnlInformation.TnlPort,
							TnlAddress: &modifyItem.TnlInformation.TnlAddress,
						},
						TnlUsage: modifyItem.TnlUsage,
					},
				},
			},
		}
		ie.GetValue().GetE2ConnectionUpdateModify().Value = append(ie.GetValue().GetE2ConnectionUpdateModify().Value, cai)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateRemove(removeItems []*types.TnlInformation) *E2ConnectionUpdate {

	ie := &E2ConnectionUpdateIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemove),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateIe{
			E2ConnectionUpdateIe: &E2ConnectionUpdateIe_E2ConnectionUpdateRemove{
				E2ConnectionUpdateRemove: &E2ConnectionUpdateRemoveList{
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
				E2ConnectionUpdateRemoveItemIe: &E2ConnectionUpdateRemoveItemIe_E2ConnectionUpdateRemoveItem{
					E2ConnectionUpdateRemoveItem: &E2ConnectionUpdateRemoveItem{
						TnlInformation: &e2ap_ies.Tnlinformation{
							TnlPort:    removeItem.TnlPort,
							TnlAddress: &removeItem.TnlAddress,
						},
					},
				},
			},
		}
		ie.GetValue().GetE2ConnectionUpdateRemove().Value = append(ie.GetValue().GetE2ConnectionUpdateRemove().Value, cai)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ResetRequest) SetTransactionID(trID int32) *ResetRequest {

	ie := &ResetRequestIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &ResetRequestIe{
			ResetRequestIe: &ResetRequestIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ResetRequest) SetCause(cause *e2ap_ies.Cause) *ResetRequest {

	ie := &ResetRequestIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &ResetRequestIe{
			ResetRequestIe: &ResetRequestIe_Cause{
				Cause: cause,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ErrorIndication) SetTransactionID(trID int32) *ErrorIndication {

	ie := &ErrorIndicationIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &ErrorIndicationIe{
			ErrorIndicationIe: &ErrorIndicationIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ErrorIndication) SetRicRequestID(ricReqID *types.RicRequest) *ErrorIndication {

	ie := &ErrorIndicationIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &ErrorIndicationIe{
			ErrorIndicationIe: &ErrorIndicationIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID),
					RicInstanceId:  int32(ricReqID.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ErrorIndication) SetRanFunctionID(ranFuncID *types.RanFunctionID) *ErrorIndication {

	ie := &ErrorIndicationIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &ErrorIndicationIe{
			ErrorIndicationIe: &ErrorIndicationIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ErrorIndication) SetCause(cause *e2ap_ies.Cause) *ErrorIndication {

	ie := &ErrorIndicationIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &ErrorIndicationIe{
			ErrorIndicationIe: &ErrorIndicationIe_Cause{
				Cause: cause,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ErrorIndication) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *ErrorIndication {

	ie := &ErrorIndicationIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &ErrorIndicationIe{
			ErrorIndicationIe: &ErrorIndicationIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ResetResponse) SetTransactionID(trID int32) *ResetResponse {

	ie := &ResetResponseIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &ResetResponseIe{
			ResetResponseIe: &ResetResponseIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *ResetResponse) SetCriticalityDiagnostics(failureProcCode v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *ResetResponse {

	ie := &ResetResponseIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &ResetResponseIe{
			ResetResponseIe: &ResetResponseIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolAcknowledge) SetRicRequestID(ricReqID types.RicRequest) *RiccontrolAcknowledge {

	ie := &RiccontrolAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolAcknowledgeIe{
			RiccontrolAcknowledgeIe: &RiccontrolAcknowledgeIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
					RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolAcknowledge) SetRanFunctionID(ranFuncID *types.RanFunctionID) *RiccontrolAcknowledge {

	ie := &RiccontrolAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolAcknowledgeIe{
			RiccontrolAcknowledgeIe: &RiccontrolAcknowledgeIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolAcknowledge) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolAcknowledge {

	ie := &RiccontrolAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolAcknowledgeIe{
			RiccontrolAcknowledgeIe: &RiccontrolAcknowledgeIe_RiccallProcessId{
				RiccallProcessId: &e2ap_commondatatypes.RiccallProcessId{
					Value: ricCallPrID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolAcknowledge) SetRicControlOutcome(ricCtrlOut types.RicControlOutcome) *RiccontrolAcknowledge {

	ie := &RiccontrolAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDRiccontrolOutcome),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolAcknowledgeIe{
			RiccontrolAcknowledgeIe: &RiccontrolAcknowledgeIe_RiccontrolOutcome{
				RiccontrolOutcome: &e2ap_commondatatypes.RiccontrolOutcome{
					Value: ricCtrlOut,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolFailure) SetRicRequestID(ricReqID types.RicRequest) *RiccontrolFailure {

	ie := &RiccontrolFailureIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolFailureIe{
			RiccontrolFailureIe: &RiccontrolFailureIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
					RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolFailure) SetRanFunctionID(ranFuncID *types.RanFunctionID) *RiccontrolFailure {

	ie := &RiccontrolFailureIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolFailureIe{
			RiccontrolFailureIe: &RiccontrolFailureIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolFailure) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolFailure {

	ie := &RiccontrolFailureIes{
		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolFailureIe{
			RiccontrolFailureIe: &RiccontrolFailureIe_RiccallProcessId{
				RiccallProcessId: &e2ap_commondatatypes.RiccallProcessId{
					Value: ricCallPrID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolFailure) SetCause(c *e2ap_ies.Cause) *RiccontrolFailure {

	ie := &RiccontrolFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RiccontrolFailureIe{
			RiccontrolFailureIe: &RiccontrolFailureIe_Cause{
				Cause: c,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolFailure) SetRicControlOutcome(ricCtrlOut types.RicControlOutcome) *RiccontrolFailure {

	ie := &RiccontrolFailureIes{
		Id:          int32(v2.ProtocolIeIDRiccontrolOutcome),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolFailureIe{
			RiccontrolFailureIe: &RiccontrolFailureIe_RiccontrolOutcome{
				RiccontrolOutcome: &e2ap_commondatatypes.RiccontrolOutcome{
					Value: ricCtrlOut,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolRequest) SetRicRequestID(ricReqID types.RicRequest) *RiccontrolRequest {

	ie := &RiccontrolRequestIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolRequestIe{
			RiccontrolRequestIe: &RiccontrolRequestIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
					RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolRequest) SetRanFunctionID(ranFuncID *types.RanFunctionID) *RiccontrolRequest {

	ie := &RiccontrolRequestIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolRequestIe{
			RiccontrolRequestIe: &RiccontrolRequestIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolRequest) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolRequest {

	ie := &RiccontrolRequestIes{
		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolRequestIe{
			RiccontrolRequestIe: &RiccontrolRequestIe_RiccallProcessId{
				RiccallProcessId: &e2ap_commondatatypes.RiccallProcessId{
					Value: ricCallPrID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolRequest) SetRicControlHeader(ricCtrlHdr types.RicControlHeader) *RiccontrolRequest {

	ie := &RiccontrolRequestIes{
		Id:          int32(v2.ProtocolIeIDRiccontrolHeader),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolRequestIe{
			RiccontrolRequestIe: &RiccontrolRequestIe_RiccontrolHeader{
				RiccontrolHeader: &e2ap_commondatatypes.RiccontrolHeader{
					Value: ricCtrlHdr,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolRequest) SetRicControlMessage(ricCtrlMsg types.RicControlMessage) *RiccontrolRequest {

	ie := &RiccontrolRequestIes{
		Id:          int32(v2.ProtocolIeIDRiccontrolMessage),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolRequestIe{
			RiccontrolRequestIe: &RiccontrolRequestIe_RiccontrolMessage{
				RiccontrolMessage: &e2ap_commondatatypes.RiccontrolMessage{
					Value: ricCtrlMsg,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RiccontrolRequest) SetRicControlAckRequest(ricCtrlAckRequest e2ap_ies.RiccontrolAckRequest) *RiccontrolRequest {

	ie := &RiccontrolRequestIes{
		Id:          int32(v2.ProtocolIeIDRiccontrolAckRequest),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RiccontrolRequestIe{
			RiccontrolRequestIe: &RiccontrolRequestIe_RiccontrolAckRequest{
				RiccontrolAckRequest: ricCtrlAckRequest,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicRequestID(ricReqID types.RicRequest) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
					RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRanFunctionID(ranFuncID types.RanFunctionID) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicActionID(ricAction int32) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRicactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RicactionId{
				RicactionId: &e2ap_ies.RicactionId{
					Value: ricAction,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicIndicationSN(ricSn types.RicIndicationSn) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRicindicationSn),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RicindicationSn{
				RicindicationSn: &e2ap_ies.RicindicationSn{
					Value: int32(ricSn),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicIndicationType(ricIndicationType e2ap_ies.RicindicationType) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRicindicationType),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RicindicationType{
				RicindicationType: ricIndicationType,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicIndicationHeader(ricIndHd types.RicIndicationHeader) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRicindicationHeader),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RicindicationHeader{
				RicindicationHeader: &e2ap_commondatatypes.RicindicationHeader{
					Value: []byte(ricIndHd),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicIndicationMessage(ricIndMsg types.RicIndicationMessage) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRicindicationMessage),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RicindicationMessage{
				RicindicationMessage: &e2ap_commondatatypes.RicindicationMessage{
					Value: []byte(ricIndMsg),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *Ricindication) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *Ricindication {

	ie := &RicindicationIes{
		Id:          int32(v2.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicindicationIe{
			RicindicationIe: &RicindicationIe_RiccallProcessId{
				RiccallProcessId: &e2ap_commondatatypes.RiccallProcessId{
					Value: ricCallPrID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceQuery) SetRanFunctionsAccepted(rfAccepted types.RanFunctionRevisions) *RicserviceQuery {

	ie := &RicserviceQueryIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceQueryIe{
			RicServiceQueryIe: &RicServiceQueryIe_RanfunctionsAccepted{
				RanfunctionsAccepted: &RanfunctionsIdListSingleContainer{
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
				RanfunctionIdItemIe: &RanfunctionIdItemIe_RanfunctionIdItem{
					RanfunctionIdItem: &RanfunctionIdItem{
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
		ie.GetValue().GetRanfunctionsAccepted().GetValue().Value = append(ie.GetValue().GetRanfunctionsAccepted().GetValue().Value, rfIDiIe)
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceQuery) SetTransactionID(trID int32) *RicserviceQuery {

	ie := &RicserviceQueryIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceQueryIe{
			RicServiceQueryIe: &RicServiceQueryIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *RicserviceUpdate) SetTransactionID(trID int32) *RicserviceUpdate {

	ie := &RicserviceUpdateIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicserviceUpdateIe{
			RicServiceUpdateIe: &RicserviceUpdateIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *RicserviceUpdate) SetRanFunctionsAdded(rfal types.RanFunctions) *RicserviceUpdate {

	rfa := &RanfunctionsList{
		Value: make([]*RanfunctionItemIes, 0),
	}

	for id, ranFunctionID := range rfal {
		ranFunction := RanfunctionItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionItemIe{
				RanfunctionItemIe: &RanfunctionItemIe_RanfunctionItem{
					RanfunctionItem: &RanfunctionItem{
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
		rfa.Value = append(rfa.Value, &ranFunction)
	}

	ie := &RicserviceUpdateIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAdded),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicserviceUpdateIe{
			RicServiceUpdateIe: &RicserviceUpdateIe_RanfunctionsAdded{
				RanfunctionsAdded: &RanfunctionsListSingleContainer{
					Value: rfa,
				},
			},
		},
	}
	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdate) SetRanFunctionsModified(rfml types.RanFunctions) *RicserviceUpdate {

	rfm := &RanfunctionsList{
		Value: make([]*RanfunctionItemIes, 0),
	}

	for id, ranFunctionID := range rfml {
		ranFunction := RanfunctionItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionItemIe{
				RanfunctionItemIe: &RanfunctionItemIe_RanfunctionItem{
					RanfunctionItem: &RanfunctionItem{
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
		rfm.Value = append(rfm.Value, &ranFunction)
	}

	ie := &RicserviceUpdateIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsModified),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicserviceUpdateIe{
			RicServiceUpdateIe: &RicserviceUpdateIe_RanfunctionsModified{
				RanfunctionsModified: &RanfunctionsListSingleContainer{
					Value: rfm,
				},
			},
		},
	}
	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdate) SetRanFunctionsDeleted(rfDeleted types.RanFunctionRevisions) *RicserviceUpdate {

	rfd := &RanfunctionsIdList{
		Value: make([]*RanfunctionIdItemIes, 0),
	}

	for rfID, rfRevision := range rfDeleted {
		rfIDiIe := RanfunctionIdItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionIdItemIe{
				RanfunctionIdItemIe: &RanfunctionIdItemIe_RanfunctionIdItem{
					RanfunctionIdItem: &RanfunctionIdItem{
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
		rfd.Value = append(rfd.Value, &rfIDiIe)
	}

	ie := &RicserviceUpdateIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsDeleted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicserviceUpdateIe{
			RicServiceUpdateIe: &RicserviceUpdateIe_RanfunctionsDeleted{
				RanfunctionsDeleted: &RanfunctionsIdListSingleContainer{
					Value: rfd,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdateAcknowledge) SetTransactionID(trID int32) *RicserviceUpdateAcknowledge {

	ie := &RicserviceUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceUpdateAcknowledgeIe{
			RicServiceUpdateAcknowledgeIe: &RicServiceUpdateAcknowledgeIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *RicserviceUpdateAcknowledge) SetRanFunctionsAccepted(rfAccepted types.RanFunctionRevisions) *RicserviceUpdateAcknowledge {

	rfa := &RanfunctionsIdList{
		Value: make([]*RanfunctionIdItemIes, 0),
	}

	for rfID, rfRevision := range rfAccepted {
		ranFunction := RanfunctionIdItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionIdItemIe{
				RanfunctionIdItemIe: &RanfunctionIdItemIe_RanfunctionIdItem{
					RanfunctionIdItem: &RanfunctionIdItem{
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
		rfa.Value = append(rfa.Value, &ranFunction)
	}

	ie := &RicserviceUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceUpdateAcknowledgeIe{
			RicServiceUpdateAcknowledgeIe: &RicServiceUpdateAcknowledgeIe_RanfunctionsAccepted{
				RanfunctionsAccepted: &RanfunctionsIdListSingleContainer{
					Value: rfa,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdateAcknowledge) SetRanFunctionsRejected(rfRejected types.RanFunctionCauses) *RicserviceUpdateAcknowledge {
	rfr := &RanfunctionsIdcauseList{
		Value: make([]*RanfunctionIdcauseItemIes, 0),
	}

	for id, cause := range rfRejected {
		rfIDcIIe := RanfunctionIdcauseItemIes{
			Id:          int32(v2.ProtocolIeIDRanfunctionIeCauseItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RanfunctionIdcauseItemIe{
				RanfunctionIdcauseItemIe: &RanfunctionIdcauseItemIe_RanfunctionIecauseItem{
					RanfunctionIecauseItem: &RanfunctionIdcauseItem{
						RanFunctionId: &e2ap_ies.RanfunctionId{
							Value: int32(id),
						},
						Cause: &e2ap_ies.Cause{},
					},
				},
			},
		}

		switch cause.GetCause().(type) {
		case *e2ap_ies.Cause_Misc:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_Misc{
				Misc: cause.GetMisc(),
			}
		case *e2ap_ies.Cause_Protocol:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_Protocol{
				Protocol: cause.GetProtocol(),
			}
		case *e2ap_ies.Cause_RicService:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_RicService{
				RicService: cause.GetRicService(),
			}
		case *e2ap_ies.Cause_RicRequest:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_RicRequest{
				RicRequest: cause.GetRicRequest(),
			}
		case *e2ap_ies.Cause_Transport:
			rfIDcIIe.GetValue().GetRanfunctionIecauseItem().GetCause().Cause = &e2ap_ies.Cause_Transport{
				Transport: cause.GetTransport(),
			}

		default:
			return m
		}

		rfr.Value = append(rfr.Value, &rfIDcIIe)
	}

	ie := &RicserviceUpdateAcknowledgeIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceUpdateAcknowledgeIe{
			RicServiceUpdateAcknowledgeIe: &RicServiceUpdateAcknowledgeIe_RanfunctionsRejected{
				RanfunctionsRejected: &RanfunctionsIdcauseListSingleContainer{
					Value: rfr,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdateFailure) SetTransactionID(trID int32) *RicserviceUpdateFailure {

	ie := &RicserviceUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceUpdateFailureIe{
			RicServiceUpdateFailureIe: &RicServiceUpdateFailureIe_TransactionId{
				TransactionId: &e2ap_ies.TransactionId{
					Value: trID,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)

	return m
}

func (m *RicserviceUpdateFailure) SetCause(c *e2ap_ies.Cause) *RicserviceUpdateFailure {

	ie := &RicserviceUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicServiceUpdateFailureIe{
			RicServiceUpdateFailureIe: &RicServiceUpdateFailureIe_Cause{
				Cause: c,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *RicserviceUpdateFailure {

	ie := &RicserviceUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicServiceUpdateFailureIe{
			RicServiceUpdateFailureIe: &RicServiceUpdateFailureIe_TimeToWait{
				TimeToWait: ttw,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicserviceUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicserviceUpdateFailure {

	ie := &RicserviceUpdateFailureIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicServiceUpdateFailureIe{
			RicServiceUpdateFailureIe: &RicServiceUpdateFailureIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteFailure) SetRicRequestID(ricReq *types.RicRequest) *RicsubscriptionDeleteFailure {

	ie := &RicsubscriptionDeleteFailureIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionDeleteFailureIe{
			RicsubscriptionDeleteFailureIe: &RicsubscriptionDeleteFailureIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReq.RequestorID),
					RicInstanceId:  int32(ricReq.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteFailure) SetRanFunctionID(ranFuncID types.RanFunctionID) *RicsubscriptionDeleteFailure {

	ie := &RicsubscriptionDeleteFailureIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionDeleteFailureIe{
			RicsubscriptionDeleteFailureIe: &RicsubscriptionDeleteFailureIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(ranFuncID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteFailure) SetCause(cause *e2ap_ies.Cause) *RicsubscriptionDeleteFailure {

	ie := &RicsubscriptionDeleteFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicsubscriptionDeleteFailureIe{
			RicsubscriptionDeleteFailureIe: &RicsubscriptionDeleteFailureIe_Cause{
				Cause: cause,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteFailure) SetCriticalityDiagnostics(failureProcCode *v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicsubscriptionDeleteFailure {

	ie := &RicsubscriptionDeleteFailureIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicsubscriptionDeleteFailureIe{
			RicsubscriptionDeleteFailureIe: &RicsubscriptionDeleteFailureIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteRequest) SetRicRequestID(ricReq *types.RicRequest) *RicsubscriptionDeleteRequest {

	ie := &RicsubscriptionDeleteRequestIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionDeleteRequestIe{
			RicsubscriptionDeleteRequestIe: &RicsubscriptionDeleteRequestIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReq.RequestorID),
					RicInstanceId:  int32(ricReq.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteRequest) SetRanFunctionID(ranFuncID types.RanFunctionID) *RicsubscriptionDeleteRequest {

	ie := &RicsubscriptionDeleteRequestIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionDeleteRequestIe{
			RicsubscriptionDeleteRequestIe: &RicsubscriptionDeleteRequestIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(ranFuncID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteResponse) SetRicRequestID(ricReq *types.RicRequest) *RicsubscriptionDeleteResponse {

	ie := &RicsubscriptionDeleteResponseIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionDeleteResponseIe{
			RicsubscriptionDeleteResponseIe: &RicsubscriptionDeleteResponseIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReq.RequestorID),
					RicInstanceId:  int32(ricReq.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteResponse) SetRanFunctionID(ranFuncID types.RanFunctionID) *RicsubscriptionDeleteResponse {

	ie := &RicsubscriptionDeleteResponseIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionDeleteResponseIe{
			RicsubscriptionDeleteResponseIe: &RicsubscriptionDeleteResponseIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(ranFuncID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionDeleteRequired) SetRicSubscriptionToBeRemoved(rswcl types.RicSubscriptionWithCauseList) *RicsubscriptionDeleteRequired {

	rstbrl := &RicsubscriptionListWithCause{
		Value: make([]*RicsubscriptionWithCauseItemIes, 0),
	}

	for rfid, item := range rswcl {
		rswci := &RicsubscriptionWithCauseItemIes{
			Id:          int32(v2.ProtocolIeIDRICsubscriptionWithCauseItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RicsubscriptionWithCauseItemIe{
				RicsubscriptionWithCauseItemIe: &RicsubscriptionWithCauseItemIe_RicsubscriptionWithCauseItem{
					RicsubscriptionWithCauseItem: &RicsubscriptionWithCauseItem{
						RicRequestId: &e2ap_ies.RicrequestId{
							RicRequestorId: int32(item.RicRequestID.RequestorID),
							RicInstanceId:  int32(item.RicRequestID.InstanceID),
						},
						RanFunctionId: &e2ap_ies.RanfunctionId{
							Value: int32(rfid),
						},
						Cause: item.Cause,
					},
				},
			},
		}
		rstbrl.Value = append(rstbrl.Value, rswci)
	}

	ie := &RicsubscriptionDeleteRequiredIes{
		Id:          int32(v2.ProtocolIeIDRICsubscriptionToBeRemoved),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicsubscriptionDeleteRequiredIe{
			RicsubscriptionDeleteRequiredIe: &RicsubscriptionDeleteRequiredIe_RicsubscriptionToBeRemoved{
				RicsubscriptionToBeRemoved: rstbrl,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionRequest) SetRicRequestID(ricReqID *types.RicRequest) *RicsubscriptionRequest {

	ie := &RicsubscriptionRequestIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionRequestIe{
			RicsubscriptionRequestIe: &RicsubscriptionRequestIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID),
					RicInstanceId:  int32(ricReqID.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionRequest) SetRanFunctionID(ranFuncID *types.RanFunctionID) *RicsubscriptionRequest {

	ie := &RicsubscriptionRequestIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionRequestIe{
			RicsubscriptionRequestIe: &RicsubscriptionRequestIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionRequest) SetRicSubscriptionDetails(ricEventDef types.RicEventDefintion,
	ricActionsToBeSetup map[types.RicActionID]types.RicActionDef) *RicsubscriptionRequest {

	ratbsl := &RicactionsToBeSetupList{
		Value: make([]*RicactionToBeSetupItemIes, 0),
	}

	for ricActionID, ricAction := range ricActionsToBeSetup {
		ricActionToSetup := RicactionToBeSetupItemIes{
			Id:          int32(v2.ProtocolIeIDRicactionToBeSetupItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RicactionToBeSetupItemIe{
				RicactionToBeSetupItemIe: &RicactionToBeSetupItemIe_RicactionToBeSetupItem{
					RicactionToBeSetupItem: &RicactionToBeSetupItem{
						RicActionId: &e2ap_ies.RicactionId{
							Value: int32(ricActionID), // range of Integer from e2ap-v01.00.asn1:1059, value from line 1283
						},
						RicActionType: ricAction.RicActionType,
						RicActionDefinition: &e2ap_commondatatypes.RicactionDefinition{
							Value: ricAction.RicActionDefinition,
						},
						RicSubsequentAction: &e2ap_ies.RicsubsequentAction{
							RicSubsequentActionType: ricAction.RicSubsequentAction,
							RicTimeToWait:           ricAction.Ricttw,
						},
					},
				},
			},
		}
		ratbsl.Value = append(ratbsl.Value, &ricActionToSetup)
	}

	ie := &RicsubscriptionRequestIes{
		Id:          int32(v2.ProtocolIeIDRicsubscriptionDetails),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionRequestIe{
			RicsubscriptionRequestIe: &RicsubscriptionRequestIe_RicsubscriptionDetails{
				RicsubscriptionDetails: &RicsubscriptionDetails{
					RicEventTriggerDefinition: &e2ap_commondatatypes.RiceventTriggerDefinition{
						Value: ricEventDef,
					},
					RicActionToBeSetupList: ratbsl,
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionResponse) SetRicRequestID(ricReqID *types.RicRequest) *RicsubscriptionResponse {

	ie := &RicsubscriptionResponseIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionResponseIe{
			RicsubscriptionResponseIe: &RicsubscriptionResponseIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID),
					RicInstanceId:  int32(ricReqID.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionResponse) SetRanFunctionID(ranFuncID *types.RanFunctionID) *RicsubscriptionResponse {

	ie := &RicsubscriptionResponseIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionResponseIe{
			RicsubscriptionResponseIe: &RicsubscriptionResponseIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionResponse) SetRicActionAdmitted(ricActionsAdmitted []*types.RicActionID) *RicsubscriptionResponse {
	raal := &RicactionAdmittedList{
		Value: make([]*RicactionAdmittedItemIes, 0),
	}

	for _, ricActionID := range ricActionsAdmitted {
		ranaIe := &RicactionAdmittedItemIes{
			Id:          int32(v2.ProtocolIeIDRicactionAdmittedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RicactionAdmittedItemIe{
				RicactionAdmittedItemIe: &RicactionAdmittedItemIe_RicactionAdmittedItem{
					RicactionAdmittedItem: &RicactionAdmittedItem{
						RicActionId: &e2ap_ies.RicactionId{
							Value: int32(*ricActionID),
						},
					},
				},
			},
		}
		raal.Value = append(raal.Value, ranaIe)
	}

	ie := &RicsubscriptionResponseIes{
		Id:          int32(v2.ProtocolIeIDRicactionsAdmitted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionResponseIe{
			RicsubscriptionResponseIe: &RicsubscriptionResponseIe_RicactionsAdmitted{
				RicactionsAdmitted: raal,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionResponse) SetRicActionNotAdmitted(ricActionsNotAdmitted map[types.RicActionID]*e2ap_ies.Cause) *RicsubscriptionResponse {
	ranal := &RicactionNotAdmittedList{
		Value: make([]*RicactionNotAdmittedItemIes, 0),
	}

	for ricActionID, cause := range ricActionsNotAdmitted {
		ranaIe := &RicactionNotAdmittedItemIes{
			Id:          int32(v2.ProtocolIeIDRicactionNotAdmittedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RicactionNotAdmittedItemIe{
				RicactionNotAdmittedItemIe: &RicactionNotAdmittedItemIe_RicactionNotAdmittedItem{
					RicactionNotAdmittedItem: &RicactionNotAdmittedItem{
						RicActionId: &e2ap_ies.RicactionId{
							Value: int32(ricActionID),
						},
						Cause: cause,
					},
				},
			},
		}
		ranal.Value = append(ranal.Value, ranaIe)
	}

	ie := &RicsubscriptionResponseIes{
		Id:          int32(v2.ProtocolIeIDRicactionsNotAdmitted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionResponseIe{
			RicsubscriptionResponseIe: &RicsubscriptionResponseIe_RicactionsNotAdmitted{
				RicactionsNotAdmitted: ranal,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionFailure) SetRicRequestID(ricReqID *types.RicRequest) *RicsubscriptionFailure {

	ie := &RicsubscriptionFailureIes{
		Id:          int32(v2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionFailureIe{
			RicsubscriptionFailureIe: &RicsubscriptionFailureIe_RicrequestId{
				RicrequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(ricReqID.RequestorID),
					RicInstanceId:  int32(ricReqID.InstanceID),
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionFailure) SetRanFunctionID(ranFuncID *types.RanFunctionID) *RicsubscriptionFailure {

	ie := &RicsubscriptionFailureIes{
		Id:          int32(v2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicsubscriptionFailureIe{
			RicsubscriptionFailureIe: &RicsubscriptionFailureIe_RanfunctionId{
				RanfunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
				},
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionFailure) SetCause(cause *e2ap_ies.Cause) *RicsubscriptionFailure {

	ie := &RicsubscriptionFailureIes{
		Id:          int32(v2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicsubscriptionFailureIe{
			RicsubscriptionFailureIe: &RicsubscriptionFailureIe_Cause{
				Cause: cause,
			},
		},
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}

func (m *RicsubscriptionFailure) SetCriticalityDiagnostics(failureProcCode v2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicsubscriptionFailure {

	ie := &RicsubscriptionFailureIes{
		Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &RicsubscriptionFailureIe{
			RicsubscriptionFailureIe: &RicsubscriptionFailureIe_CriticalityDiagnostics{
				CriticalityDiagnostics: &e2ap_ies.CriticalityDiagnostics{
					ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
						Value: int32(failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
					},
					TriggeringMessage:    failureTrigMsg,
					ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
					RicRequestorId: &e2ap_ies.RicrequestId{
						RicRequestorId: int32(reqID.RequestorID),
						RicInstanceId:  int32(reqID.InstanceID),
					},
				},
			},
		},
	}

	if critDiags != nil {
		cdl := &e2ap_ies.CriticalityDiagnosticsIeList{
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
			cdl.Value = append(cdl.Value, &criticDiagnostics)
		}
		ie.GetValue().GetCriticalityDiagnostics().IEsCriticalityDiagnostics = cdl
	}

	m.ProtocolIes = append(m.ProtocolIes, ie)
	return m
}
