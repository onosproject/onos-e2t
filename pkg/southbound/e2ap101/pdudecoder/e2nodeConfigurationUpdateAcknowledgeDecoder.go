// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeE2nodeConfigurationUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) ([]*types.E2NodeComponentConfigUpdateAckItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2ncua := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate()
	if e2ncua == nil {
		return nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdateAcknowledge")
	}

	e2nccual := make([]*types.E2NodeComponentConfigUpdateAckItem, 0)
	list := e2ncua.GetSuccessfulOutcome().GetProtocolIes().GetValue().GetValue()
	for _, ie := range list {
		e2nccuai := types.E2NodeComponentConfigUpdateAckItem{}
		e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentType()
		e2nccuai.E2NodeComponentID = *ie.GetValue().GetE2NodeComponentId()
		e2nccuai.E2NodeComponentConfigUpdateAck = types.E2NodeComponentConfigUpdateAck{
			UpdateOutcome: ie.GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(),
			FailureCause:  *ie.GetValue().GetE2NodeComponentConfigUpdateAck().GetFailureCause(),
		}

		e2nccual = append(e2nccual, &e2nccuai)
	}

	return e2nccual, nil
}
