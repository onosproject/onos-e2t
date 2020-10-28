// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"encoding/binary"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
)

const mask20bit = 0xFFFFF

func CreateResponseE2apPdu(plmnID string, ricID uint32) (*e2appdudescriptions.E2ApPdu, error) {
	if len(plmnID) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	// Expecting 20 bits for ric ID
	if ricID|mask20bit > mask20bit {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricID)
	}

	gricIDIe := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte(plmnID),
			},
			RicId: make([]byte, 4),
		},
	}
	binary.LittleEndian.PutUint32(gricIDIe.Value.RicId, ricID)

	ranFunctions := e2appducontents.E2SetupResponseIes_E2SetupResponseIes9{
		Value: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
	}

	rfIDiIe100 := e2appducontents.RanfunctionIdItemIes{
		RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
			Value: &e2appducontents.RanfunctionIdItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 100,
				},
				RanFunctionRevision: &e2apies.RanfunctionRevision{
					Value: 1,
				},
			},
		},
	}
	ranFunctions.Value.Value = append(ranFunctions.Value.Value, &rfIDiIe100)

	ranfunctionsIdcauseList := e2appducontents.E2SetupResponseIes_E2SetupResponseIes13{
		Value: &e2appducontents.RanfunctionsIdcauseList{
			Value: make([]*e2appducontents.RanfunctionIdcauseItemIes, 0),
		},
	}

	rfIDcLi100 := e2appducontents.RanfunctionIdcauseItemIes{
		RanFunctionIdcauseItemIes7: &e2appducontents.RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
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
		},
	}
	ranfunctionsIdcauseList.Value.Value = append(ranfunctionsIdcauseList.Value.Value, &rfIDcLi100)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						SuccessfulOutcome: &e2appducontents.E2SetupResponse{
							ProtocolIes: &e2appducontents.E2SetupResponseIes{
								E2ApProtocolIes4:  &gricIDIe,                //global RIC ID
								E2ApProtocolIes9:  &ranFunctions,            //RanFunctionIdList
								E2ApProtocolIes13: &ranfunctionsIdcauseList, //RanFunctionIdCauseList
							},
						},
					},
				},
			},
		},
	}
	return &e2apPdu, nil
}
