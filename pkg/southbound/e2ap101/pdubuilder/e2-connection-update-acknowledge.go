// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func CreateE2connectionUpdateAcknowledgeE2apPdu(connSetup []*types.E2ConnectionUpdateItem,
	connSetFail []*types.E2ConnectionSetupFailedItem) (*e2appdudescriptions.E2ApPdu, error) {

	connectionSetup := e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes39{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionSetup),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		ConnectionSetup: &e2appducontents.E2ConnectionUpdateList{
			Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, setupItem := range connSetup {
		si := &e2appducontents.E2ConnectionUpdateItemIes{
			Id:          int32(v1beta2.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.E2ConnectionUpdateItem{
				TnlInformation: &e2ap_ies.Tnlinformation{
					TnlPort:    &setupItem.TnlInformation.TnlPort,
					TnlAddress: &setupItem.TnlInformation.TnlAddress,
				},
				TnlUsage: setupItem.TnlUsage,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		connectionSetup.ConnectionSetup.Value = append(connectionSetup.ConnectionSetup.Value, si)
	}

	connectionSetupFailed := e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes40{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionSetupFailed),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		ConnectionSetupFailed: &e2appducontents.E2ConnectionSetupFailedList{
			Value: make([]*e2appducontents.E2ConnectionSetupFailedItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, sfItem := range connSetFail {
		sfi := &e2appducontents.E2ConnectionSetupFailedItemIes{
			Id:          int32(v1beta2.ProtocolIeIDE2connectionSetupFailedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.E2ConnectionSetupFailedItem{
				TnlInformation: &e2ap_ies.Tnlinformation{
					TnlPort:    &sfItem.TnlInformation.TnlPort,
					TnlAddress: &sfItem.TnlInformation.TnlAddress,
				},
				Cause: &sfItem.Cause,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		connectionSetupFailed.ConnectionSetupFailed.Value = append(connectionSetupFailed.ConnectionSetupFailed.Value, sfi)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
						SuccessfulOutcome: &e2appducontents.E2ConnectionUpdateAcknowledge{
							ProtocolIes: &e2appducontents.E2ConnectionUpdateAckIes{
								E2ApProtocolIes39: &connectionSetup,       //E2 Connection Setup List
								E2ApProtocolIes40: &connectionSetupFailed, //E2 Connection Setup Failed List
							},
						},
						ProcedureCode: &e2ap_constants.IdE2ConnectionUpdate{
							Value: int32(v1beta2.ProcedureCodeIDE2connectionUpdate),
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
