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

func CreateRicServiceUpdateAcknowledgeE2apPdu(rfAccepted types.RanFunctionRevisions, rfRejected types.RanFunctionCauses) (*e2appdudescriptions.E2ApPdu, error) {

	if rfAccepted == nil && rfRejected == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						SuccessfulOutcome: &e2appducontents.RicserviceUpdateAcknowledge{
							ProtocolIes: &e2appducontents.RicserviceUpdateAcknowledgeIes{
								//E2ApProtocolIes9:  &ranFunctionsAccepted, //RAN functions Accepted
								//E2ApProtocolIes13: &ranFunctionsRejected, //RAN functions Rejected
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

	if rfAccepted != nil {
		ranFunctionsAccepted := e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes9{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionsAccepted),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.RanfunctionsIdList{
				Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for rfID, rfRevision := range rfAccepted {
			rfIDiIe := e2appducontents.RanfunctionIdItemIes{
				RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
					Id:          int32(v1beta2.ProtocolIeIDRanfunctionIDItem),
					Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
					Value: &e2appducontents.RanfunctionIdItem{
						RanFunctionId: &e2apies.RanfunctionId{
							Value: int32(rfID),
						},
						RanFunctionRevision: &e2apies.RanfunctionRevision{
							Value: int32(rfRevision),
						},
					},
					Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
				},
			}
			ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, &rfIDiIe)
		}
		e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetSuccessfulOutcome().GetProtocolIes().E2ApProtocolIes9 = &ranFunctionsAccepted
	}

	if rfRejected != nil {
		ranFunctionsRejected := e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes13{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionsRejected),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
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
		e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate().GetSuccessfulOutcome().GetProtocolIes().E2ApProtocolIes13 = &ranFunctionsRejected
	}

	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
