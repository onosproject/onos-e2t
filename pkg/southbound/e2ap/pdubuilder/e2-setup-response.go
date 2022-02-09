// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func NewE2SetupResponse(trID int32, plmnID types.PlmnID, ricID types.RicIdentifier,
	e2nccaal []*types.E2NodeComponentConfigAdditionAckItem) (*e2appducontents.E2SetupResponse, error) {

	// Expecting 20 bits for ric ID
	if len(ricID.RicIdentifierValue) != 3 {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricID)
	}

	res := &e2appducontents.E2SetupResponse{
		ProtocolIes: make([]*e2appducontents.E2SetupResponseIes, 0),
	}

	res.SetTransactionID(trID).SetGlobalRicID(plmnID, ricID).SetE2nodeComponentConfigAdditionAck(e2nccaal)

	return res, nil
}

func CreateResponseE2apPdu(response *e2appducontents.E2SetupResponse) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: int32(v2.ProcedureCodeIDE2setup),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures{
					SoValues: &e2appdudescriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup{
						E2Setup: response,
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
