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

func CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu(e2nccual []*types.E2NodeComponentConfigUpdateAckItem) (*e2appdudescriptions.E2ApPdu, error) {

	configUpdateAckList := e2appducontents.E2NodeComponentConfigUpdateAckList{
		Value: make([]*e2appducontents.E2NodeComponentConfigUpdateAckItemIes, 0),
	}

	for _, e2nccuai := range e2nccual {
		cuai := &e2appducontents.E2NodeComponentConfigUpdateAckItemIes{
			Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2NodeComponentConfigUpdateAckItem{
				E2NodeComponentType: e2nccuai.E2NodeComponentType,
				//E2NodeComponentId:   &e2nccuai.E2NodeComponentID,
				E2NodeComponentConfigUpdateAck: &e2ap_ies.E2NodeComponentConfigUpdateAck{
					UpdateOutcome: e2nccuai.E2NodeComponentConfigUpdateAck.UpdateOutcome,
					//FailureCause:  &e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		if &e2nccuai.E2NodeComponentID != nil {
			cuai.Value.E2NodeComponentId = &e2nccuai.E2NodeComponentID
		} else {
			cuai.Value.E2NodeComponentId = nil
		}
		if &e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause != nil {
			cuai.Value.E2NodeComponentConfigUpdateAck.FailureCause = &e2nccuai.E2NodeComponentConfigUpdateAck.FailureCause
		} else {
			cuai.Value.E2NodeComponentConfigUpdateAck.FailureCause = nil
		}

		fmt.Printf("Composed ID is \n%v of type %T\n", *cuai.Value.GetE2NodeComponentId(), cuai.Value.GetE2NodeComponentId())
		fmt.Printf("FailureCause is \n%v of type %T\n", *cuai.Value.E2NodeComponentConfigUpdateAck.GetFailureCause(), cuai.Value.E2NodeComponentConfigUpdateAck.GetFailureCause())
		fmt.Printf("E2nodeComponentConfigUpdateAckItem is \n%v\n", *cuai.GetValue())

		configUpdateAckList.Value = append(configUpdateAckList.Value, cuai)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
						SuccessfulOutcome: &e2appducontents.E2NodeConfigurationUpdateAcknowledge{
							ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes{
								Id:          int32(v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAck),
								Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
								Value:       &configUpdateAckList,
								Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
							},
						},
						ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
							Value: int32(v1beta2.ProcedureCodeIDE2nodeConfigurationUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	fmt.Printf("Composed message is \n%v", configUpdateAckList)

	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
