// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2appducontents

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap201/types"
)

func (m *E2SetupResponse) SetE2nodeComponentConfigUpdateAck(e2nccual []*types.E2NodeComponentConfigUpdateAckItem) *E2SetupResponse {
	configUpdateAckList := E2NodeComponentConfigUpdateAckList{
		Value: make([]*E2NodeComponentConfigUpdateAckItemIes, 0),
	}

	for _, e2nccuai := range e2nccual {
		cuai := &E2NodeComponentConfigUpdateAckItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigUpdateAckItem{
				E2NodeComponentType: e2nccuai.E2NodeComponentType,
				//E2NodeComponentId:   e2nccuai.E2NodeComponentID,
				E2NodeComponentConfigUpdateAck: &e2ap_ies.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: e2nccuai.E2NodeComponentConfigUpdateAck.UpdateOutcome,
					//FailureCause:  e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		if e2nccuai.E2NodeComponentID != nil {
			cuai.Value.E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause != nil {
			cuai.Value.E2NodeComponentConfigUpdateAck.FailureCause = e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause
		}

		configUpdateAckList.Value = append(configUpdateAckList.Value, cuai)
	}
	m.GetProtocolIes().E2ApProtocolIes35 = &E2SetupResponseIes_E2SetupResponseIes35{
		Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &configUpdateAckList,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2SetupResponse) SetRanFunctionRejected(rfRejected types.RanFunctionCauses) *E2SetupResponse {
	ranFunctionsRejected := E2SetupResponseIes_E2SetupResponseIes13{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsIdcauseList{
			Value: make([]*RanfunctionIdcauseItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for id, cause := range rfRejected {
		rfIDcIIe := RanfunctionIdcauseItemIes{
			RanFunctionIdcauseItemIes7: &RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionIeCauseItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionIdcauseItem{
					RanFunctionId: &e2ap_ies.RanfunctionId{
						Value: int32(id),
					},
					Cause: &e2ap_ies.Cause{},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}

		switch cause.GetCause().(type) {
		case *e2ap_ies.Cause_Misc:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Misc{
				Misc: cause.GetMisc(),
			}
		case *e2ap_ies.Cause_Protocol:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Protocol{
				Protocol: cause.GetProtocol(),
			}
		case *e2ap_ies.Cause_RicService:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicService{
				RicService: cause.GetRicService(),
			}
		case *e2ap_ies.Cause_RicRequest:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicRequest{
				RicRequest: cause.GetRicRequest(),
			}
		case *e2ap_ies.Cause_Transport:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Transport{
				Transport: cause.GetTransport(),
			}

		default:
			return m
		}
		ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, &rfIDcIIe)
	}
	m.GetProtocolIes().E2ApProtocolIes13 = &ranFunctionsRejected
	return m
}

func (m *E2SetupResponse) SetRanFunctionAccepted(rfAccepted types.RanFunctionRevisions) *E2SetupResponse {
	ranFunctionsAccepted := E2SetupResponseIes_E2SetupResponseIes9{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsIdList{
			Value: make([]*RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for rfID, rfRevision := range rfAccepted {
		rfIDiIe := RanfunctionIdItemIes{
			RanFunctionIdItemIes6: &RanfunctionIdItemIes_RanfunctionIdItemIes6{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionIDItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionIdItem{
					RanFunctionId: &e2ap_ies.RanfunctionId{
						Value: int32(rfID),
					},
					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
						Value: int32(rfRevision),
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, &rfIDiIe)
	}
	m.GetProtocolIes().E2ApProtocolIes9 = &ranFunctionsAccepted
	return m
}

func (m *E2SetupFailure) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2SetupFailure {
	criticalityDiagnostics := E2SetupFailureIes_E2SetupFailureIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = &criticalityDiagnostics
	return m
}

func (m *E2SetupFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2SetupFailure {
	m.GetProtocolIes().E2ApProtocolIes31 = &E2SetupFailureIes_E2SetupFailureIes31{
		Id:          int32(v2beta1.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       ttw, // Could be any other value
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2SetupFailure) SetTnlInformation(tnl *e2ap_ies.Tnlinformation) *E2SetupFailure {
	m.GetProtocolIes().E2ApProtocolIes48 = &E2SetupFailureIes_E2SetupFailureIes48{
		Id:          int32(v2beta1.ProtocolIeIDTNLinformation),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       tnl,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2NodeConfigurationUpdate) SetGlobalE2nodeID(e2NodeID *e2ap_ies.GlobalE2NodeId) *E2NodeConfigurationUpdate {
	m.GetProtocolIes().E2ApProtocolIes3 = &E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes3{
		Id:          int32(v2beta1.ProtocolIeIDGlobalE2nodeID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       e2NodeID,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2NodeConfigurationUpdate) SetE2nodeComponentConfigUpdate(e2nccul []*types.E2NodeComponentConfigUpdateItem) *E2NodeConfigurationUpdate {
	configUpdateList := E2NodeComponentConfigUpdateList{
		Value: make([]*E2NodeComponentConfigUpdateItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &E2NodeComponentConfigUpdateItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigUpdateItem{
				E2NodeComponentType:         e2nccui.E2NodeComponentType,
				E2NodeComponentId:           e2nccui.E2NodeComponentID,
				E2NodeComponentConfigUpdate: &e2nccui.E2NodeComponentConfigUpdate,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		configUpdateList.Value = append(configUpdateList.Value, cui)
	}
	m.GetProtocolIes().E2ApProtocolIes33 = &E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes33{
		Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdate),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &configUpdateList,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2NodeConfigurationUpdateAcknowledge) SetE2nodeComponentConfigUpdate(e2nccual []*types.E2NodeComponentConfigUpdateAckItem) *E2NodeConfigurationUpdateAcknowledge {
	configUpdateAckList := E2NodeComponentConfigUpdateAckList{
		Value: make([]*E2NodeComponentConfigUpdateAckItemIes, 0),
	}

	for _, e2nccuai := range e2nccual {
		cuai := &E2NodeComponentConfigUpdateAckItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &E2NodeComponentConfigUpdateAckItem{
				E2NodeComponentType: e2nccuai.E2NodeComponentType,
				//E2NodeComponentId:   e2nccuai.E2NodeComponentID,
				E2NodeComponentConfigUpdateAck: &e2ap_ies.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: e2nccuai.E2NodeComponentConfigUpdateAck.UpdateOutcome,
					//FailureCause:  e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		if e2nccuai.E2NodeComponentID != nil {
			cuai.Value.E2NodeComponentId = e2nccuai.E2NodeComponentID
		}
		if e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause != nil {
			cuai.Value.E2NodeComponentConfigUpdateAck.FailureCause = e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause
		}

		configUpdateAckList.Value = append(configUpdateAckList.Value, cuai)
	}
	m.GetProtocolIes().E2ApProtocolIes35 = &E2NodeConfigurationUpdateAcknowledgeIes_E2NodeConfigurationUpdateAcknowledgeIes35{
		Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateAck),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &configUpdateAckList,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2NodeConfigurationUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2NodeConfigurationUpdateFailure {
	criticalityDiagnostics := E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = &criticalityDiagnostics
	return m
}

func (m *E2NodeConfigurationUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2NodeConfigurationUpdateFailure {
	m.GetProtocolIes().E2ApProtocolIes31 = &E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes31{
		Id:          int32(v2beta1.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       ttw, // Could be any other value
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2ConnectionUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *E2ConnectionUpdateFailure {
	criticalityDiagnostics := E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = &criticalityDiagnostics
	return m
}

func (m *E2ConnectionUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *E2ConnectionUpdateFailure {
	m.GetProtocolIes().E2ApProtocolIes31 = &E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes31{
		Id:          int32(v2beta1.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       ttw, // Could be any other value
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *E2ConnectionUpdateFailure) SetCause(c *e2ap_ies.Cause) *E2ConnectionUpdateFailure {
	cause := E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes1{
		Id:          int32(v2beta1.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       c,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	m.GetProtocolIes().E2ApProtocolIes1 = &cause
	return m
}

func (m *E2ConnectionUpdateAcknowledge) SetE2ConnectionSetup(connSetup []*types.E2ConnectionUpdateItem) *E2ConnectionUpdateAcknowledge {
	connectionSetup := E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes39{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionSetup),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateList{
			Value: make([]*E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, setupItem := range connSetup {
		si := &E2ConnectionUpdateItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateItem{
				TnlInformation: &e2ap_ies.Tnlinformation{
					TnlPort:    &setupItem.TnlInformation.TnlPort,
					TnlAddress: &setupItem.TnlInformation.TnlAddress,
				},
				TnlUsage: setupItem.TnlUsage,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		connectionSetup.Value.Value = append(connectionSetup.Value.Value, si)
	}
	m.GetProtocolIes().E2ApProtocolIes39 = &connectionSetup
	return m
}

func (m *E2ConnectionUpdateAcknowledge) SetE2ConnectionSetupFailed(connSetFail []*types.E2ConnectionSetupFailedItem) *E2ConnectionUpdateAcknowledge {
	connectionSetupFailed := E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes40{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionSetupFailed),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionSetupFailedList{
			Value: make([]*E2ConnectionSetupFailedItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, sfItem := range connSetFail {
		sfi := &E2ConnectionSetupFailedItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionSetupFailedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionSetupFailedItem{
				TnlInformation: &e2ap_ies.Tnlinformation{
					TnlPort:    &sfItem.TnlInformation.TnlPort,
					TnlAddress: &sfItem.TnlInformation.TnlAddress,
				},
				Cause: &sfItem.Cause,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		connectionSetupFailed.Value.Value = append(connectionSetupFailed.Value.Value, sfi)
	}
	m.GetProtocolIes().E2ApProtocolIes40 = &connectionSetupFailed
	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateAdd(addItems []*types.E2ConnectionUpdateItem) *E2ConnectionUpdate {
	connectionAddList := E2ConnectionUpdateIes_E2ConnectionUpdateIes44{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateAdd),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateList{
			Value: make([]*E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, addItem := range addItems {
		cai := &E2ConnectionUpdateItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateItem{
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
	m.GetProtocolIes().E2ApProtocolIes44 = &connectionAddList
	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateModify(modifyItems []*types.E2ConnectionUpdateItem) *E2ConnectionUpdate {
	connectionModifyList := E2ConnectionUpdateIes_E2ConnectionUpdateIes45{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateModify),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateList{
			Value: make([]*E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, modifyItem := range modifyItems {
		cmi := &E2ConnectionUpdateItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateItem{
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
	m.GetProtocolIes().E2ApProtocolIes45 = &connectionModifyList
	return m
}

func (m *E2ConnectionUpdate) SetE2ConnectionUpdateRemove(removeItems []*types.TnlInformation) *E2ConnectionUpdate {
	connectionRemoveList := E2ConnectionUpdateIes_E2ConnectionUpdateIes46{
		Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateRemove),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &E2ConnectionUpdateRemoveList{
			Value: make([]*E2ConnectionUpdateRemoveItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, removeItem := range removeItems {
		cri := &E2ConnectionUpdateRemoveItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2connectionUpdateRemoveItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &E2ConnectionUpdateRemoveItem{
				TnlInformation: &e2ap_ies.Tnlinformation{
					TnlPort:    &removeItem.TnlPort,
					TnlAddress: &removeItem.TnlAddress,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		connectionRemoveList.Value.Value = append(connectionRemoveList.Value.Value, cri)
	}
	m.GetProtocolIes().E2ApProtocolIes46 = &connectionRemoveList
	return m
}

func (m *ErrorIndication) SetTransactionID(trID int32) *ErrorIndication {
	m.GetProtocolIes().E2ApProtocolIes49 = &ErrorIndicationIes_ErrorIndicationIes49{
		Id:          int32(v2beta1.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_ies.TransactionId{
			Value: trID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *ErrorIndication) SetRicRequestID(ricReqID *types.RicRequest) *ErrorIndication {
	m.GetProtocolIes().E2ApProtocolIes29 = &ErrorIndicationIes_ErrorIndicationIes29{
		Id:          int32(v2beta1.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_ies.RicrequestId{
			RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *ErrorIndication) SetRanFunctionID(ranFuncID *types.RanFunctionID) *ErrorIndication {
	m.GetProtocolIes().E2ApProtocolIes5 = &ErrorIndicationIes_ErrorIndicationIes5{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_ies.RanfunctionId{
			Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *ErrorIndication) SetCause(cause *e2ap_ies.Cause) *ErrorIndication {
	m.GetProtocolIes().E2ApProtocolIes1 = &ErrorIndicationIes_ErrorIndicationIes1{
		Id:          int32(v2beta1.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       cause,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *ErrorIndication) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *ErrorIndication {
	criticalityDiagnostics := &ErrorIndicationIes_ErrorIndicationIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
	return m
}

func (m *ResetResponse) SetCriticalityDiagnostics(failureProcCode v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *ResetResponse {
	criticalityDiagnostics := &ResetResponseIes_ResetResponseIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
	return m
}

func (m *RiccontrolAcknowledge) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolAcknowledge {
	m.GetProtocolIes().E2ApProtocolIes20 = &RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes20{
		Id:          int32(v2beta1.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccallProcessId{
			Value: ricCallPrID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RiccontrolAcknowledge) SetRicControlOutcome(ricCtrlOut types.RicControlOutcome) *RiccontrolAcknowledge {
	m.GetProtocolIes().E2ApProtocolIes32 = &RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes32{
		Id:          int32(v2beta1.ProtocolIeIDRiccontrolOutcome),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccontrolOutcome{
			Value: ricCtrlOut,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RiccontrolFailure) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolFailure {
	m.GetProtocolIes().E2ApProtocolIes20 = &RiccontrolFailureIes_RiccontrolFailureIes20{
		Id:          int32(v2beta1.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccallProcessId{
			Value: ricCallPrID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RiccontrolFailure) SetRicControlOutcome(ricCtrlOut types.RicControlOutcome) *RiccontrolFailure {
	m.GetProtocolIes().E2ApProtocolIes32 = &RiccontrolFailureIes_RiccontrolFailureIes32{
		Id:          int32(v2beta1.ProtocolIeIDRiccontrolOutcome),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccontrolOutcome{
			Value: ricCtrlOut,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RiccontrolRequest) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *RiccontrolRequest {
	m.GetProtocolIes().E2ApProtocolIes20 = &RiccontrolRequestIes_RiccontrolRequestIes20{
		Id:          int32(v2beta1.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccallProcessId{
			Value: ricCallPrID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RiccontrolRequest) SetRicControlAckRequest(ricCtrlAckRequest e2ap_ies.RiccontrolAckRequest) *RiccontrolRequest {
	m.GetProtocolIes().E2ApProtocolIes21 = &RiccontrolRequestIes_RiccontrolRequestIes21{
		Id:          int32(v2beta1.ProtocolIeIDRiccontrolAckRequest),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ricCtrlAckRequest,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *Ricindication) SetRicIndicationSN(ricSn types.RicIndicationSn) *Ricindication {
	m.GetProtocolIes().E2ApProtocolIes27 = &RicindicationIes_RicindicationIes27{
		Id:          int32(v2beta1.ProtocolIeIDRicindicationSn),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_ies.RicindicationSn{
			Value: int32(ricSn),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *Ricindication) SetRicCallProcessID(ricCallPrID types.RicCallProcessID) *Ricindication {
	m.GetProtocolIes().E2ApProtocolIes20 = &RicindicationIes_RicindicationIes20{
		Id:          int32(v2beta1.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccallProcessId{
			Value: ricCallPrID,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RicserviceQuery) SetRanFunctionsAccepted(rfAccepted types.RanFunctionRevisions) *RicserviceQuery {
	ranFunctionsAccepted := RicserviceQueryIes_RicserviceQueryIes9{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsIdList{
			Value: make([]*RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for rfID, rfRevision := range rfAccepted {
		rfIDiIe := RanfunctionIdItemIes{
			RanFunctionIdItemIes6: &RanfunctionIdItemIes_RanfunctionIdItemIes6{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionIDItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionIdItem{
					RanFunctionId: &e2ap_ies.RanfunctionId{
						Value: int32(rfID),
					},
					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
						Value: int32(rfRevision),
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, &rfIDiIe)
	}
	m.GetProtocolIes().E2ApProtocolIes9 = &ranFunctionsAccepted
	return m
}

func (m *RicserviceUpdate) SetRanFunctionsAdded(rfal types.RanFunctions) *RicserviceUpdate {
	ranFunctionsAddedList := RicserviceUpdateIes_RicserviceUpdateIes10{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsAdded),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsList{
			Value: make([]*RanfunctionItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for id, ranFunctionID := range rfal {
		ranFunction := RanfunctionItemIes{
			E2ApProtocolIes10: &RanfunctionItemIes_RanfunctionItemIes8{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionItem),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionItem{
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
		}
		ranFunctionsAddedList.Value.Value = append(ranFunctionsAddedList.Value.Value, &ranFunction)
	}
	m.GetProtocolIes().E2ApProtocolIes10 = &ranFunctionsAddedList
	return m
}

func (m *RicserviceUpdate) SetRanFunctionsModified(rfml types.RanFunctions) *RicserviceUpdate {
	ranFunctionsModifiedList := RicserviceUpdateIes_RicserviceUpdateIes12{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsModified),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsList{
			Value: make([]*RanfunctionItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for id, ranFunctionID := range rfml {
		ranFunction := RanfunctionItemIes{
			E2ApProtocolIes10: &RanfunctionItemIes_RanfunctionItemIes8{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionItem),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionItem{
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
		}
		ranFunctionsModifiedList.Value.Value = append(ranFunctionsModifiedList.Value.Value, &ranFunction)
	}
	m.GetProtocolIes().E2ApProtocolIes12 = &ranFunctionsModifiedList
	return m
}

func (m *RicserviceUpdate) SetRanFunctionsDeleted(rfDeleted types.RanFunctionRevisions) *RicserviceUpdate {
	ranFunctionsDeletedList := RicserviceUpdateIes_RicserviceUpdateIes11{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsDeleted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsIdList{
			Value: make([]*RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for rfID, rfRevision := range rfDeleted {
		rfIDiIe := RanfunctionIdItemIes{
			RanFunctionIdItemIes6: &RanfunctionIdItemIes_RanfunctionIdItemIes6{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionIDItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionIdItem{
					RanFunctionId: &e2ap_ies.RanfunctionId{
						Value: int32(rfID),
					},
					RanFunctionRevision: &e2ap_ies.RanfunctionRevision{
						Value: int32(rfRevision),
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		ranFunctionsDeletedList.Value.Value = append(ranFunctionsDeletedList.Value.Value, &rfIDiIe)
	}
	m.GetProtocolIes().E2ApProtocolIes11 = &ranFunctionsDeletedList
	return m
}

func (m *RicserviceUpdateAcknowledge) SetRanFunctionsRejected(rfRejected types.RanFunctionCauses) *RicserviceUpdateAcknowledge {
	ranFunctionsRejected := RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes13{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RanfunctionsIdcauseList{
			Value: make([]*RanfunctionIdcauseItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for id, cause := range rfRejected {
		rfIDcIIe := RanfunctionIdcauseItemIes{
			RanFunctionIdcauseItemIes7: &RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionIeCauseItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &RanfunctionIdcauseItem{
					RanFunctionId: &e2ap_ies.RanfunctionId{
						Value: int32(id),
					},
					Cause: &e2ap_ies.Cause{},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}

		switch cause.GetCause().(type) {
		case *e2ap_ies.Cause_Misc:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Misc{
				Misc: cause.GetMisc(),
			}
		case *e2ap_ies.Cause_Protocol:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Protocol{
				Protocol: cause.GetProtocol(),
			}
		case *e2ap_ies.Cause_RicService:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicService{
				RicService: cause.GetRicService(),
			}
		case *e2ap_ies.Cause_RicRequest:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_RicRequest{
				RicRequest: cause.GetRicRequest(),
			}
		case *e2ap_ies.Cause_Transport:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2ap_ies.Cause_Transport{
				Transport: cause.GetTransport(),
			}

		default:
			return m
		}
		ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, &rfIDcIIe)
	}
	m.GetProtocolIes().E2ApProtocolIes13 = &ranFunctionsRejected
	return m
}

func (m *RicserviceUpdateFailure) SetTimeToWait(ttw e2ap_ies.TimeToWait) *RicserviceUpdateFailure {
	m.GetProtocolIes().E2ApProtocolIes31 = &RicserviceUpdateFailureIes_RicserviceUpdateFailureIes31{
		Id:          int32(v2beta1.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       ttw, // Could be any other value
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}
	return m
}

func (m *RicserviceUpdateFailure) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicserviceUpdateFailure {
	criticalityDiagnostics := &RicserviceUpdateFailureIes_RicserviceUpdateFailureIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
	return m
}

func (m *RicsubscriptionDeleteFailure) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicsubscriptionDeleteFailure {
	criticalityDiagnostics := &RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
	return m
}

func (m *RicsubscriptionResponse) SetRicActionNotAdmitted(ricActionsNotAccepted map[types.RicActionID]*e2ap_ies.Cause) *RicsubscriptionResponse {
	ricActionNotAdmit := &RicsubscriptionResponseIes_RicsubscriptionResponseIes18{
		Id:          int32(v2beta1.ProtocolIeIDRicactionsNotAdmitted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &RicactionNotAdmittedList{
			Value: make([]*RicactionNotAdmittedItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for ricActionID, cause := range ricActionsNotAccepted {
		ranaIe := &RicactionNotAdmittedItemIes{
			Id:          int32(v2beta1.ProtocolIeIDRicactionNotAdmittedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &RicactionNotAdmittedItem{
				RicActionId: &e2ap_ies.RicactionId{
					Value: int32(ricActionID),
				},
				Cause: cause,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		ricActionNotAdmit.GetValue().Value = append(ricActionNotAdmit.GetValue().Value, ranaIe)
	}
	m.GetProtocolIes().E2ApProtocolIes18 = ricActionNotAdmit
	return m
}

func (m *RicsubscriptionFailure) SetCriticalityDiagnostics(failureProcCode *v2beta1.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) *RicsubscriptionFailure {
	criticalityDiagnostics := &RicsubscriptionFailureIes_RicsubscriptionFailureIes2{
		Id:          int32(v2beta1.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
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
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	if critDiags != nil {
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2ap_ies.CriticalityDiagnosticsIeList{
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
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
		}
	}

	m.GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
	return m
}
