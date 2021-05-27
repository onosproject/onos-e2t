// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func CreateRicServiceUpdateFailureE2apPdu(rfRejected types.RanFunctionCauses, ttw e2apies.TimeToWait, failureProcCode v1beta2.ProcedureCodeT, failureCrit e2ap_commondatatypes.Criticality,
	failureTrigMsg e2ap_commondatatypes.TriggeringMessage, reqID *types.RicRequest,
	critDiags []*types.CritDiag) (*e2appdudescriptions.E2ApPdu, error) {

	ranFunctionsRejected := e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes13{
		Id:          int32(v1beta2.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.RanfunctionsIdcauseList{
			Value: make([]*e2appducontents.RanfunctionIdcauseItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for id, cause := range rfRejected {
		rfIDcIIe := e2appducontents.RanfunctionIdcauseItemIes{
			RanFunctionIdcauseItemIes7: &e2appducontents.RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
				Id:          int32(v1beta2.ProtocolIeIDRanfunctionIeCauseItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &e2appducontents.RanfunctionIdcauseItem{
					RanFunctionId: &e2apies.RanfunctionId{
						Value: int32(id),
					},
					Cause: &e2apies.Cause{},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}

		switch causeType := cause.GetCause().(type) {
		case *e2apies.Cause_Misc:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2apies.Cause_Misc{
				Misc: cause.GetMisc(),
			}
		case *e2apies.Cause_Protocol:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2apies.Cause_Protocol{
				Protocol: cause.GetProtocol(),
			}
		case *e2apies.Cause_RicService:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2apies.Cause_RicService{
				RicService: cause.GetRicService(),
			}
		case *e2apies.Cause_RicRequest:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2apies.Cause_RicRequest{
				RicRequest: cause.GetRicRequest(),
			}
		case *e2apies.Cause_Transport:
			rfIDcIIe.GetRanFunctionIdcauseItemIes7().GetValue().GetCause().Cause = &e2apies.Cause_Transport{
				Transport: cause.GetTransport(),
			}

		default:
			return nil, fmt.Errorf("unexpected cause type %v", causeType)
		}
		ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, &rfIDcIIe)
	}

	timeToWait := e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes31{
		Id:          int32(v1beta2.ProtocolIeIDTimeToWait),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       ttw,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	criticalityDiagnostics := e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes2{
		Id:          int32(v1beta2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2apies.CriticalityDiagnostics{
			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
				Value: int32(failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
			},
			TriggeringMessage:    failureTrigMsg,
			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
			RicRequestorId: &e2apies.RicrequestId{
				RicRequestorId: int32(reqID.RequestorID),
				RicInstanceId:  int32(reqID.InstanceID),
			},
			IEsCriticalityDiagnostics: &e2apies.CriticalityDiagnosticsIeList{
				Value: make([]*e2apies.CriticalityDiagnosticsIeItem, 0),
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, critDiag := range critDiags {
		criticDiagnostics := e2apies.CriticalityDiagnosticsIeItem{
			IEcriticality: critDiag.IECriticality,
			IEId: &e2ap_commondatatypes.ProtocolIeId{
				Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
			},
			TypeOfError: critDiag.TypeOfError,
		}
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						UnsuccessfulOutcome: &e2appducontents.RicserviceUpdateFailure{
							ProtocolIes: &e2appducontents.RicserviceUpdateFailureIes{
								E2ApProtocolIes13: &ranFunctionsRejected,   //RAN functions Rejected
								E2ApProtocolIes31: &timeToWait,             //Time to Wait
								E2ApProtocolIes2:  &criticalityDiagnostics, //Criticality Diagnostics
							},
						},
						ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
							Value: int32(v1beta2.ProcedureCodeIDRICserviceUpdate),
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
