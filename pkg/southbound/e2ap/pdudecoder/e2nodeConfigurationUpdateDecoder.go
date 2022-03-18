// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"

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
	var globalE2NodeID *e2apies.GlobalE2NodeId
	e2nccual := make([]*types.E2NodeComponentConfigUpdateItem, 0)
	for _, v := range e2ncu.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTransactionId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDGlobalE2nodeID) {
			globalE2NodeID = v.GetValue().GetGlobalE2NodeId()
		}
		if v.Id == int32(v2.ProtocolIeIDE2nodeComponentConfigUpdate) {
			list := v.GetValue().GetE2NodeComponentConfigUpdate().GetValue()
			for _, ie := range list {
				e2nccuai := types.E2NodeComponentConfigUpdateItem{}
				e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentConfigUpdateItem().GetE2NodeComponentInterfaceType()
				e2nccuai.E2NodeComponentID = ie.GetValue().GetE2NodeComponentConfigUpdateItem().GetE2NodeComponentId()
				e2nccuai.E2NodeComponentConfiguration = *ie.GetValue().GetE2NodeComponentConfigUpdateItem().GetE2NodeComponentConfiguration()

				e2nccual = append(e2nccual, &e2nccuai)
			}
		}
	}

	// Extract node ID
	nodeIdentity, err := ExtractE2NodeIdentity(globalE2NodeID, e2ncu)
	if err != nil {
		return nil, nil, nil, err
	}

	return &transactionID, nodeIdentity, e2nccual, nil
}
