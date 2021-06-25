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

func CreateErrorIndicationE2apPdu(ricReqID *types.RicRequest, ranFuncID *types.RanFunctionID,
	cause *e2apies.Cause, failureProcCode *v1beta2.ProcedureCodeT,
	failureCrit *e2ap_commondatatypes.Criticality, failureTrigMsg *e2ap_commondatatypes.TriggeringMessage,
	reqID *types.RicRequest, critDiags []*types.CritDiag) (*e2appdudescriptions.E2ApPdu, error) {

	if ricReqID == nil && ranFuncID == nil && cause == nil && failureProcCode == nil && failureCrit == nil && failureTrigMsg == nil &&
		reqID == nil && critDiags == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					ErrorIndication: &e2appdudescriptions.ErrorIndicationEp{
						InitiatingMessage: &e2appducontents.ErrorIndication{
							ProtocolIes: &e2appducontents.ErrorIndicationIes{
								//E2ApProtocolIes29: &ricRequestID,           // RIC Requestor & RIC Instance ID
								//E2ApProtocolIes5:  &ranFunctionID,          // RAN function ID
								//E2ApProtocolIes1:  &errorCause,             // Cause
								//E2ApProtocolIes2:  &criticalityDiagnostics, // CriticalityDiagnostics
							},
						},
						ProcedureCode: &e2ap_constants.IdErrorIndication{
							Value: int32(v1beta2.ProcedureCodeIDErrorIndication),
						},
						Criticality: &e2ap_commondatatypes.CriticalityIgnore{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
						},
					},
				},
			},
		},
	}

	if ricReqID != nil {
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetErrorIndication().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes29 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes29{
			Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2apies.RicrequestId{
				RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
				RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	}

	if ranFuncID != nil {
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetErrorIndication().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes5 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes5{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2apies.RanfunctionId{
				Value: int32(*ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	}

	if cause != nil {
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetErrorIndication().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes1 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes1{
			Id:          int32(v1beta2.ProtocolIeIDCause),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cause,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	}

	if failureProcCode != nil && failureTrigMsg != nil && failureCrit != nil && reqID != nil {
		criticalityDiagnostics := &e2appducontents.ErrorIndicationIes_ErrorIndicationIes2{
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

		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetErrorIndication().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes2 = criticalityDiagnostics
	}

	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
