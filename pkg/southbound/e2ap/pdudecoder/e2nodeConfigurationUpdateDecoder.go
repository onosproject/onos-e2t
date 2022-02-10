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

func DecodeE2nodeConfigurationUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, *types.E2NodeIdentity, []*types.E2NodeComponentConfigUpdateItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2ncu := e2apPdu.GetInitiatingMessage().GetValue().GetE2NodeConfigurationUpdate()
	if e2ncu == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdate")
	}

	var err error
	var transactionID int32
	var nodeIdentity *types.E2NodeIdentity
	e2nccual := make([]*types.E2NodeComponentConfigUpdateItem, 0)
	for _, v := range e2ncu.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDGlobalE2nodeID) {
			globalE2NodeID := v.GetValue().GetGe2NId()
			nodeIdentity, err = ExtractE2NodeIdentity(globalE2NodeID)
			if err != nil {
				return nil, nil, nil, err
			}
		}
		if v.Id == int32(v2.ProtocolIeIDE2nodeComponentConfigUpdate) {
			list := v.GetValue().GetE2Nccul().GetValue()
			for _, ie := range list {
				e2nccuai := types.E2NodeComponentConfigUpdateItem{}
				e2nccuai.E2NodeComponentType = ie.GetValue().GetE2Nccui().GetE2NodeComponentInterfaceType()
				e2nccuai.E2NodeComponentID = ie.GetValue().GetE2Nccui().GetE2NodeComponentId()
				e2nccuai.E2NodeComponentConfiguration = *ie.GetValue().GetE2Nccui().GetE2NodeComponentConfiguration()

				e2nccual = append(e2nccual, &e2nccuai)
			}
		}
	}

	return &transactionID, nodeIdentity, e2nccual, nil
}
