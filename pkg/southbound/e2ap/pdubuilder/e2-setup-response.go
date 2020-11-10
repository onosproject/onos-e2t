// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

const mask20bit = 0xFFFFF

func CreateResponseE2apPdu(plmnID string, ricID uint32,
	rfAccepted types.RanFunctionIDs, rfRejected map[int32]e2apies.Cause) (*e2appdudescriptions.E2ApPdu, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	// Expecting 20 bits for ric ID
	if ricID|mask20bit > mask20bit {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricID)
	}

	globalRicID := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Id:          int32(v1beta1.ProtocolIeIDGlobalRicID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte(plmnID),
			},
			RicId: &e2ap_commondatatypes.BitString{
				Value: uint64(ricID),
				Len:   20,
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionsAccepted := e2appducontents.E2SetupResponseIes_E2SetupResponseIes9{
		Id:          int32(v1beta1.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for rfID, rfRevision := range rfAccepted {
		rfIDiIe := e2appducontents.RanfunctionIdItemIes{
			RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
				Id:          int32(v1beta1.ProtocolIeIDRanfunctionIDItem),
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

	ranFunctionsRejected := e2appducontents.E2SetupResponseIes_E2SetupResponseIes13{
		Id:          int32(v1beta1.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsIdcauseList{
			Value: make([]*e2appducontents.RanfunctionIdcauseItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfIDcLi100 := e2appducontents.RanfunctionIdcauseItemIes{
		RanFunctionIdcauseItemIes7: &e2appducontents.RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionIeCauseItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RanfunctionIdcauseItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 100,
				},
				Cause: &e2apies.Cause{
					Cause: &e2apies.Cause_RicService{
						RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
					},
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, &rfIDcLi100)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						SuccessfulOutcome: &e2appducontents.E2SetupResponse{
							ProtocolIes: &e2appducontents.E2SetupResponseIes{
								E2ApProtocolIes4:  &globalRicID,          //global RIC ID
								E2ApProtocolIes9:  &ranFunctionsAccepted, //RanFunctionIdList
								E2ApProtocolIes13: &ranFunctionsRejected, //RanFunctionIdCauseList
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
