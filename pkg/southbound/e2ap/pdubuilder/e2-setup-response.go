// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func NewE2SetupResponse(trID int32, plmnID types.PlmnID, ricID types.RicIdentifier) (*e2appducontents.E2SetupResponse, error) {
	if len(plmnID) != 3 {
		return nil, fmt.Errorf("error: Plmn ID should be 3 chars")
	}
	// Expecting 20 bits for ric ID
	if len(ricID.RicIdentifierValue) != 3 {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricID)
	}

	globalRicID := e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
		Id:          int32(v2beta1.ProtocolIeIDGlobalRicID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.GlobalRicId{
			PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
			},
			RicId: &asn1.BitString{
				Value: ricID.RicIdentifierValue,
				Len:   uint32(ricID.RicIdentifierLen),
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	res := &e2appducontents.E2SetupResponse{
		ProtocolIes: &e2appducontents.E2SetupResponseIes{
			E2ApProtocolIes4: &globalRicID, //global RIC ID
			//E2ApProtocolIes9:  &ranFunctionsAccepted, //RanFunctionIdList
			//E2ApProtocolIes13: &ranFunctionsRejected, //RanFunctionIdCauseList
			E2ApProtocolIes49: &e2appducontents.E2SetupResponseIes_E2SetupResponseIes49{
				Id:          int32(v2beta1.ProtocolIeIDTransactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.TransactionId{
					Value: trID,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}

	return res, nil
}

func CreateResponseE2apPdu(response *e2appducontents.E2SetupResponse) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						SuccessfulOutcome: response,
						ProcedureCode: &e2ap_constants.IdE2Setup{
							Value: int32(v2beta1.ProcedureCodeIDE2setup),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
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
