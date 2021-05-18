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
)

func CreateRicServiceQueryE2apPdu() (*e2appdudescriptions.E2ApPdu, error) {

	ranFunctionsAccepted := e2appducontents.RicserviceQueryIes_RicserviceQueryIes9{
		Id:          int32(v1beta2.ProtocolIeIDRanfunctionsAccepted),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfii := &e2appducontents.RanfunctionIdItemIes{
		RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RanfunctionIdItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 123,
				},
				RanFunctionRevision: &e2apies.RanfunctionRevision{
					Value: 1,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, rfii)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceQuery: &e2appdudescriptions.RicServiceQuery{
						InitiatingMessage: &e2appducontents.RicserviceQuery{
							ProtocolIes: &e2appducontents.RicserviceQueryIes{
								RicserviceQueryIes9: &ranFunctionsAccepted, //RAN functions Accepted List
							},
						},
						ProcedureCode: &e2ap_constants.IdRicserviceQuery{
							Value: int32(v1beta2.ProcedureCodeIDRICserviceQuery),
						},
						Criticality: &e2ap_commondatatypes.CriticalityIgnore{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
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
