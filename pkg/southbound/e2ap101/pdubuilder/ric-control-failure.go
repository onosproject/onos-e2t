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

func CreateRicControlFailureE2apPdu(ricReqID types.RicRequest, ranFuncID types.RanFunctionID,
	ricCallPrID types.RicCallProcessID, cause e2apies.Cause,
	ricCtrlOut types.RicControlOutcome) (*e2appdudescriptions.E2ApPdu, error) {

	ricRequestID := e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes29{
		Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionID := e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes5{
		Id:          int32(v1beta2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: int32(ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricCause := e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes1{
		Id:          int32(v1beta2.ProtocolIeIDCause),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value:       &cause,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_UnsuccessfulOutcome{
			UnsuccessfulOutcome: &e2appdudescriptions.UnsuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicControl: &e2appdudescriptions.RicControl{
						UnsuccessfulOutcome: &e2appducontents.RiccontrolFailure{
							ProtocolIes: &e2appducontents.RiccontrolFailureIes{
								E2ApProtocolIes29: &ricRequestID,  // RIC Requestor & RIC Instance ID
								E2ApProtocolIes5:  &ranFunctionID, // RAN function ID
								//E2ApProtocolIes20: &ricCallProcessID,  // RIC Call Process ID
								E2ApProtocolIes1: &ricCause, // Cause
								//E2ApProtocolIes32: &ricControlOutcome, // RIC Control Outcome
							},
						},
						ProcedureCode: &e2ap_constants.IdRiccontrol{
							Value: int32(v1beta2.ProcedureCodeIDRICcontrol),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	if ricCallPrID != nil {
		e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome().GetProtocolIes().E2ApProtocolIes20 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes20{
			Id:          int32(v1beta2.ProtocolIeIDRiccallProcessID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2ap_commondatatypes.RiccallProcessId{
				Value: []byte(ricCallPrID),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	}

	if ricCtrlOut != nil {
		e2apPdu.GetUnsuccessfulOutcome().GetProcedureCode().GetRicControl().GetUnsuccessfulOutcome().GetProtocolIes().E2ApProtocolIes32 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes32{
			Id:          int32(v1beta2.ProtocolIeIDRiccontrolOutcome),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2ap_commondatatypes.RiccontrolOutcome{
				Value: []byte(ricCtrlOut),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	}

	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
