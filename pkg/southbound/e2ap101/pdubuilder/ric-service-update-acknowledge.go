// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
)

func CreateRicServiceUpdateAcknowledgeE2apPdu() (*e2appdudescriptions.E2ApPdu, error) {

	ranFunctionsAccepted := e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes9{
		Id:          int32(v1beta2.ProcedureCodeIDRICsubscriptionDelete),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfai := &e2appducontents.RanfunctionIdItemIes{
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
	ranFunctionsAccepted.Value.Value = append(ranFunctionsAccepted.Value.Value, rfai)

	ranFunctionsRejected := e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes13{
		Id:          int32(v1beta2.ProtocolIeIDRanfunctionsRejected),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsIdcauseList{
			Value: make([]*e2appducontents.RanfunctionIdcauseItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfri := &e2appducontents.RanfunctionIdcauseItemIes{
		RanFunctionIdcauseItemIes7: &e2appducontents.RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
			Id:          int32(v1beta2.ProcedureCodeIDRICserviceUpdate),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RanfunctionIdcauseItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 123,
				},
				//ToDo - Pass cause as a parameter
				Cause: &e2apies.Cause{
					Cause: &e2apies.Cause_RicService{
						RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
					},
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	ranFunctionsRejected.Value.Value = append(ranFunctionsRejected.Value.Value, rfri)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						SuccessfulOutcome: &e2appducontents.RicserviceUpdateAcknowledge{
							ProtocolIes: &e2appducontents.RicserviceUpdateAcknowledgeIes{
								E2ApProtocolIes9:  &ranFunctionsAccepted, //RAN functions Accepted
								E2ApProtocolIes13: &ranFunctionsRejected, //RAN functions Rejected
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
