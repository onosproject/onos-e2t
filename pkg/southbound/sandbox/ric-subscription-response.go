// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
)

const mask20bitric_response = 0xFFFFF

func CreateRicSubscriptionResponseE2apPdu(ricReqID int32, ricInstanceID int32, ranFuncID int32) (*e2appdudescriptions.E2ApPdu, error) {

	if ricReqID|mask20bitric_response > mask20bitric_response {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricReqID)
	}
	if ricInstanceID|mask20bitric_response > mask20bitric_response {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricInstanceID)
	}

	ricRequestId := e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes29{
		Value: &e2apies.RicrequestId{
			RicRequestorId: ricReqID,      // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  ricInstanceID, // sequence from e2ap-v01.00.asn1:1127
		},
	}

	ranFunctionId := e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes5{
		Value: &e2apies.RanfunctionId{
			Value: ranFuncID, // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
	}

	ricActionAdmit := e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes17{
		Value: &e2appducontents.RicactionAdmittedList{
			Value: make([]*e2appducontents.RicactionAdmittedItemIes, 0),
		},
	}

	ricActionNotAdmit := e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes18{
		Value: &e2appducontents.RicactionNotAdmittedList{
			Value: make([]*e2appducontents.RicactionNotAdmittedItemIes, 0),
		},
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicSubscription: &e2appdudescriptions.RicSubscription{
						SuccessfulOutcome: &e2appducontents.RicsubscriptionResponse{
							ProtocolIes: &e2appducontents.RicsubscriptionResponseIes{
								E2ApProtocolIes29: &ricRequestId,      //RIC request ID
								E2ApProtocolIes5:  &ranFunctionId,     //RAN function ID
								E2ApProtocolIes17: &ricActionAdmit,    // RIC action Admitted items ---> EMPTY !!
								E2ApProtocolIes18: &ricActionNotAdmit, // RIC action not Admitted items --> EMPTY !!
							},
						},
					},
				},
			},
		},
	}
	return &e2apPdu, nil
}
