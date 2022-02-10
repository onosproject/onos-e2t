// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func NewRicSubscriptionDeleteRequest(ricReq types.RicRequest,
	ranFuncID types.RanFunctionID) (
	*e2appducontents.RicsubscriptionDeleteRequest, error) {

	rsdr := &e2appducontents.RicsubscriptionDeleteRequest{
		ProtocolIes: make([]*e2appducontents.RicsubscriptionDeleteRequestIes, 0),
	}

	rsdr.SetRicRequestID(&ricReq).SetRanFunctionID(ranFuncID)

	return rsdr, nil
}

func CreateRicSubscriptionDeleteRequestE2apPdu(ricReq types.RicRequest,
	ranFuncID types.RanFunctionID) (
	*e2appdudescriptions.E2ApPdu, error) {
	request, err := NewRicSubscriptionDeleteRequest(ricReq, ranFuncID)
	if err != nil {
		return nil, err
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDRICsubscriptionDelete),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDelete{
						RicSubscriptionDelete: request,
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
