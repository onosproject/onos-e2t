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

func NewControlRequest(ricReqID types.RicRequest, ranFuncID types.RanFunctionID,
	ricCtrlHdr types.RicControlHeader, ricCtrlMsg types.RicControlMessage) (*e2appducontents.RiccontrolRequest, error) {

	controlRequest := &e2appducontents.RiccontrolRequest{
		ProtocolIes: make([]*e2appducontents.RiccontrolRequestIes, 0),
	}

	controlRequest.SetRicRequestID(ricReqID).SetRanFunctionID(&ranFuncID).SetRicControlHeader(ricCtrlHdr).SetRicControlMessage(ricCtrlMsg)

	return controlRequest, nil

}

func CreateRicControlRequestE2apPdu(request *e2appducontents.RiccontrolRequest) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDRICcontrol),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_RicControl{
						RicControl: request,
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
