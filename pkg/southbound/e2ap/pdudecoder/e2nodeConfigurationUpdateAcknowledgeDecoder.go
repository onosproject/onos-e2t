// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2nodeConfigurationUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32,
	[]*types.E2NodeComponentConfigUpdateAckItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2ncua := e2apPdu.GetSuccessfulOutcome().GetValue().GetE2NodeConfigurationUpdate()
	if e2ncua == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdateAcknowledge")
	}

	var transactionID int32
	e2nccual := make([]*types.E2NodeComponentConfigUpdateAckItem, 0)
	for _, v := range e2ncua.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck) {
			list := v.GetValue().GetE2Nccual().GetValue()
			for _, ie := range list {
				e2nccuai := types.E2NodeComponentConfigUpdateAckItem{}
				e2nccuai.E2NodeComponentType = ie.GetValue().GetE2Nccuai().GetE2NodeComponentInterfaceType()
				e2nccuai.E2NodeComponentID = ie.GetValue().GetE2Nccuai().GetE2NodeComponentId()
				e2nccuai.E2NodeComponentConfigurationAck = types.E2NodeComponentConfigurationAck{
					UpdateOutcome: ie.GetValue().GetE2Nccuai().GetE2NodeComponentConfigurationAck().GetUpdateOutcome(),
					FailureCause:  ie.GetValue().GetE2Nccuai().GetE2NodeComponentConfigurationAck().GetFailureCause(),
				}

				e2nccual = append(e2nccual, &e2nccuai)
			}
		}
	}

	return &transactionID, e2nccual, nil
}
