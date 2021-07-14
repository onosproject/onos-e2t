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

func NewE2SetupResponse(plmnID types.PlmnID, ricID types.RicIdentifier, rfAccepted types.RanFunctionRevisions, rfRejected types.RanFunctionCauses) (*e2appducontents.E2SetupResponse, error) {
	if len(plmnID) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	// Expecting 20 bits for ric ID
	if len(ricID.RicIdentifierValue) != 3 {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricID)
	}

	globalRicID := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Id:          int32(v1beta2.ProtocolIeIDGlobalRicID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
			},
			RicId: &e2ap_commondatatypes.BitString{
				Value: ricID.RicIdentifierValue,
				Len:   uint32(ricID.RicIdentifierLen),
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	res := &e2appducontents.E2SetupResponse{
		ProtocolIes: &e2appducontents.E2SetupResponseIes{
			E2ApProtocolIes4: &globalRicID, //global RIC ID
			//E2ApProtocolIes9:  &ranFunctionsAccepted, //RanFunctionIdList
			//E2ApProtocolIes13: &ranFunctionsRejected, //RanFunctionIdCauseList
		},
	}

	if rfAccepted != nil {
		ranFunctionsAccepted := e2appducontents.E2SetupResponseIes_E2SetupResponseIes9{
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
		res.GetProtocolIes().E2ApProtocolIes9 = &ranFunctionsAccepted
	}

	if rfRejected != nil {
		ranFunctionsRejected := e2appducontents.E2SetupResponseIes_E2SetupResponseIes13{
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
		res.GetProtocolIes().E2ApProtocolIes13 = &ranFunctionsRejected
	}

	return res, nil
}

func CreateResponseE2apPdu(plmnID types.PlmnID, ricID types.RicIdentifier,
	rfAccepted types.RanFunctionRevisions, rfRejected types.RanFunctionCauses) (*e2appdudescriptions.E2ApPdu, error) {
	response, err := NewE2SetupResponse(plmnID, ricID, rfAccepted, rfRejected)
	if err != nil {
		return nil, err
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						SuccessfulOutcome: response,
						ProcedureCode: &e2ap_constants.IdE2Setup{
							Value: int32(v1beta2.ProcedureCodeIDE2setup),
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
