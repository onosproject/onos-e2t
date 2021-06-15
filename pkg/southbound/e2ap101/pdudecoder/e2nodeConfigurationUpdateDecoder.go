// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeE2nodeConfigurationUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) ([]*types.E2NodeComponentConfigUpdateItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2ncu := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2NodeConfigurationUpdate()
	if e2ncu == nil {
		return nil, fmt.Errorf("error E2APpdu does not have E2nodeConfigurationUpdate")
	}

	e2nccual := make([]*types.E2NodeComponentConfigUpdateItem, 0)
	list := e2ncu.GetInitiatingMessage().GetProtocolIes().GetValue().GetValue()
	for _, ie := range list {
		e2nccuai := types.E2NodeComponentConfigUpdateItem{}
		e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentType()
		e2nccuai.E2NodeComponentID = ie.GetValue().GetE2NodeComponentId()
		e2nccuai.E2NodeComponentConfigUpdate = *ie.GetValue().GetE2NodeComponentConfigUpdate()

		e2nccual = append(e2nccual, &e2nccuai)
	}

	return e2nccual, nil
}
