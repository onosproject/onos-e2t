// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn1/v1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn1/v1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn1/v1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap_v01_00_00_asn1/v1/e2appdudescriptions"
)

func CreateResponseE2apPdu(plmnID string) (*e2appdudescriptions.E2ApPdu, error) {
	if len(plmnID) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}

	gricIDIe := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte(plmnID),
			},
		},
	}

	ranFunctions := e2appducontents.E2SetupResponseIes_E2SetupResponseIes9{
		Value: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
	}

	ranfunctionsIdcauseList := e2appducontents.E2SetupResponseIes_E2SetupResponseIes13{
		Value: &e2appducontents.RanfunctionsIdcauseList{
			Value: make([]*e2appducontents.RanfunctionIdcauseItemIes, 0),
		},
	}

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
