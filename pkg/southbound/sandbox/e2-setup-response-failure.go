// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package sandbox

import (
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
)

const mask20bitricid = 0xFFFFF

func CreateResponseFailureE2apPdu(ricReqID int32) (*e2appdudescriptions.E2ApPdu, error) {

	if ricReqID|mask20bitricid > mask20bitricid {
		return nil, fmt.Errorf("expecting 20 bit identifier for RIC. Got %0x", ricReqID)
	}

	causeOfFailure := e2appducontents.E2SetupFailureIes_E2SetupFailureIes1{
		Value: &e2apies.Cause{
			Cause: &e2apies.Cause_RicService{
				RicService: e2apies.CauseRicservice_CAUSE_RICSERVICE_RIC_RESOURCE_LIMIT,
			},
		},
	}

	timeToWait := e2appducontents.E2SetupFailureIes_E2SetupFailureIes31{
		Value: e2apies.TimeToWait_TIME_TO_WAIT_V1S,
	}

	criticality := e2appducontents.E2SetupFailureIes_E2SetupFailureIes2{
		Value: &e2apies.CriticalityDiagnostics{
			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
				Value: 255, // range of Integer from e2ap-v01.00.asn1:1206
			},
			TriggeringMessage:    e2ap_commondatatypes.TriggeringMessage_TRIGGERING_MESSAGE_INITIATING_MESSAGE,
			ProcedureCriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
			RicRequestorId: &e2apies.RicrequestId{
				RicRequestorId: ricReqID,
			},
			IEsCriticalityDiagnostics: &e2apies.CriticalityDiagnosticsIeList{
				Value: make([]*e2apies.CriticalityDiagnosticsIeItem, 0),
			},
		},
	}
	//binary.LittleEndian.PutUint32(criticality.Value.RicRequestorId.RicRequestorId, &ricReqID)

	criticDiagnostics := e2apies.CriticalityDiagnosticsIeItem{
		IEcriticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
		IEId: &e2ap_commondatatypes.ProtocolIeId{
			Value: 65535, // range of Integer from e2ap-v01.00.asn1:1208
		},
		TypeOfError: e2apies.TypeOfError_TYPE_OF_ERROR_NOT_UNDERSTOOD,
	}
	criticality.Value.IEsCriticalityDiagnostics.Value = append(criticality.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						UnsuccessfulOutcome: &e2appducontents.E2SetupFailure{
							ProtocolIes: &e2appducontents.E2SetupFailureIes{
								E2ApProtocolIes1:  &causeOfFailure, //Cause of failure
								E2ApProtocolIes31: &timeToWait,     //Time to wait
								E2ApProtocolIes2:  &criticality,    //Criticality diagnostics details
							},
						},
					},
				},
			},
		},
	}
	return &e2apPdu, nil
}
