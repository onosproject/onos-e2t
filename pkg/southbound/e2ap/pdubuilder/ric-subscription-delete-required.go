// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func CreateRicSubscriptionDeleteRequiredE2apPdu(rswcl types.RicSubscriptionWithCauseList) (*e2appdudescriptions.E2ApPdu, error) {

	rstbrl := &e2appducontents.RicsubscriptionListWithCause{
		Value: make([]*e2appducontents.RicsubscriptionWithCauseItemIes, 0),
	}

	for rfid, item := range rswcl {
		rswci := &e2appducontents.RicsubscriptionWithCauseItemIes{
			Id:          int32(v2.ProtocolIeIDRICsubscriptionWithCauseItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RicsubscriptionWithCauseItem{
				RicRequestId: &e2ap_ies.RicrequestId{
					RicRequestorId: int32(item.RicRequestID.RequestorID),
					RicInstanceId:  int32(item.RicRequestID.InstanceID),
				},
				RanFunctionId: &e2ap_ies.RanfunctionId{
					Value: int32(rfid),
				},
				Cause: item.Cause,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

		rstbrl.Value = append(rstbrl.Value, rswci)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicSubscriptionDeleteRequired: &e2appdudescriptions.RicSubscriptionDeleteRequired{
						InitiatingMessage: &e2appducontents.RicsubscriptionDeleteRequired{
							ProtocolIes: &e2appducontents.RicsubscriptionDeleteRequiredIes{
								Id:          int32(v2.ProtocolIeIDRICsubscriptionToBeRemoved),
								Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
								Value:       rstbrl,
								Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
							},
						},
						ProcedureCode: &e2ap_constants.IdRicsubscriptionDeleteRequired{
							Value: int32(v2.ProcedureCodeIDRICsubscriptionDeleteRequired),
						},
						Criticality: &e2ap_commondatatypes.CriticalityIgnore{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
						},
					},
				},
			},
		},
	}

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
