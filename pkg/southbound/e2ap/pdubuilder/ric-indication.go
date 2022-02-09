// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func RicIndicationE2apPdu(ricReqID types.RicRequest, ranFuncID types.RanFunctionID,
	ricAction int32, ricIndicationType e2apies.RicindicationType,
	ricIndHd types.RicIndicationHeader, ricIndMsg types.RicIndicationMessage) (
	*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDRICindication),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicIndication{
						RicIndication: &e2appducontents.Ricindication{
							ProtocolIes: make([]*e2appducontents.RicindicationIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetInitiatingMessage().GetValue().GetRicIndication().SetRicRequestID(ricReqID).SetRanFunctionID(ranFuncID).
		SetRicActionID(ricAction).SetRicIndicationType(ricIndicationType).SetRicIndicationHeader(ricIndHd).SetRicIndicationMessage(ricIndMsg)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}
