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
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func CreateE2SetupFailurePdu(cause e2apies.Cause, ttw *e2apies.TimeToWait, failureProcCode *v1beta2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) (*e2appdudescriptions.E2ApPdu, error) {

	errorCause := e2appducontents.E2SetupFailureIes_E2SetupFailureIes1{
		Id:          int32(v1beta2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       &cause, // Probably, could be any other reason
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						UnsuccessfulOutcome: &e2appducontents.E2SetupFailure{
							ProtocolIes: &e2appducontents.E2SetupFailureIes{
								E2ApProtocolIes1: &errorCause, // RIC Requestor & RIC Instance ID
								//E2ApProtocolIes31: &timeToWait,             // RAN function ID
								//E2ApProtocolIes2:  &criticalityDiagnostics, // CriticalityDiagnostics
							},
						},
						ProcedureCode: &e2ap_constants.IdE2Setup{
							Value: int32(v1beta2.ProcedureCodeIDE2setup),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	if failureProcCode != nil && failureTrigMsg != nil && failureCrit != nil && reqID != nil {
		criticalityDiagnostics := e2appducontents.E2SetupFailureIes_E2SetupFailureIes2{
			Id:          int32(v1beta2.ProtocolIeIDCriticalityDiagnostics),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2apies.CriticalityDiagnostics{
				ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
					Value: int32(*failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
				},
				TriggeringMessage:    *failureTrigMsg,
				ProcedureCriticality: *failureCrit, // from e2ap-v01.00.asn1:153
				RicRequestorId: &e2apies.RicrequestId{
					RicRequestorId: int32(reqID.RequestorID),
					RicInstanceId:  int32(reqID.InstanceID),
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		if critDiags != nil {
			criticalityDiagnostics.Value.IEsCriticalityDiagnostics = &e2apies.CriticalityDiagnosticsIeList{
				Value: make([]*e2apies.CriticalityDiagnosticsIeItem, 0),
			}

			for _, critDiag := range critDiags {
				criticDiagnostics := e2apies.CriticalityDiagnosticsIeItem{
					IEcriticality: critDiag.IECriticality,
					IEId: &e2ap_commondatatypes.ProtocolIeId{
						Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
					},
					TypeOfError: critDiag.TypeOfError,
				}
				criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
			}
		}

		e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().GetProtocolIes().E2ApProtocolIes2 = &criticalityDiagnostics
	}

	if ttw != nil {
		e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetE2Setup().GetUnsuccessfulOutcome().GetProtocolIes().E2ApProtocolIes31 = &e2appducontents.E2SetupFailureIes_E2SetupFailureIes31{
			Id:          int32(v1beta2.ProtocolIeIDTimeToWait),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       *ttw, // Could be any other value
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	}

	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
