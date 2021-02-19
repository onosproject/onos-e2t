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

func CreateRicControlRequestE2apPdu(ricReqID types.RicRequest, ranFuncID types.RanFunctionID,
	ricCallPrID types.RicCallProcessID, ricCtrlHdr types.RicControlHeader, ricCtrlMsg types.RicControlMessage,
	ricCtrlAckRequest e2apies.RiccontrolAckRequest) (*e2appdudescriptions.E2ApPdu, error) {

	ricRequestID := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes29{
		Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RicrequestId{
			RicRequestorId: int32(ricReqID.RequestorID), // sequence from e2ap-v01.00.asn1:1126
			RicInstanceId:  int32(ricReqID.InstanceID),  // sequence from e2ap-v01.00.asn1:1127
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ranFunctionID := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes5{
		Id:          int32(v1beta2.ProtocolIeIDRanfunctionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.RanfunctionId{
			Value: int32(ranFuncID), // range of Integer from e2ap-v01.00.asn1:1050, value from line 1277
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricCallProcessID := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes20{
		Id:          int32(v1beta2.ProtocolIeIDRiccallProcessID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccallProcessId{
			Value: []byte(ricCallPrID),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	ricControlHeader := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes22{
		Id:          int32(v1beta2.ProtocolIeIDRiccontrolHeader),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccontrolHeader{
			Value: []byte(ricCtrlHdr),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricControlMessage := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes23{
		Id:          int32(v1beta2.ProtocolIeIDRiccontrolMessage),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_commondatatypes.RiccontrolMessage{
			Value: []byte(ricCtrlMsg),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	ricControlAckRequest := e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes21{
		Id:          int32(v1beta2.ProtocolIeIDRiccontrolAckRequest),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ricCtrlAckRequest, //2apies.RiccontrolAckRequest_RICCONTROL_ACK_REQUEST_ACK,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicControl: &e2appdudescriptions.RicControl{
						InitiatingMessage: &e2appducontents.RiccontrolRequest{
							ProtocolIes: &e2appducontents.RiccontrolRequestIes{
								E2ApProtocolIes29: &ricRequestID,         // RIC Requestor & RIC Instance ID
								E2ApProtocolIes5:  &ranFunctionID,        // RAN function ID
								E2ApProtocolIes20: &ricCallProcessID,     // RIC Call Process ID
								E2ApProtocolIes22: &ricControlHeader,     // RIC Control Header
								E2ApProtocolIes23: &ricControlMessage,    // RIC Control Message
								E2ApProtocolIes21: &ricControlAckRequest, // RIC Control Ack Request
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
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
