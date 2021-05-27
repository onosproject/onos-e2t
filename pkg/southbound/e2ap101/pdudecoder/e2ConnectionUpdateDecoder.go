// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeE2connectionUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) ([]*types.E2ConnectionUpdateItem,
	[]*types.E2ConnectionUpdateItem, []*types.TnlInformation, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2cu := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate()
	if e2cu == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2connectionUpdate")
	}

	connAdd := make([]*types.E2ConnectionUpdateItem, 0)
	cal := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()
	for _, ie := range cal {
		item := types.E2ConnectionUpdateItem{}
		item.TnlInformation.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlInformation.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		item.TnlUsage = ie.GetValue().GetTnlUsage()
		connAdd = append(connAdd, &item)
	}

	connModify := make([]*types.E2ConnectionUpdateItem, 0)
	cml := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()
	for _, ie := range cml {
		item := types.E2ConnectionUpdateItem{}
		item.TnlInformation.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlInformation.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		item.TnlUsage = ie.GetValue().GetTnlUsage()
		connModify = append(connModify, &item)
	}

	connRemove := make([]*types.TnlInformation, 0)
	crl := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()
	for _, ie := range crl {
		item := types.TnlInformation{}
		item.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		connRemove = append(connRemove, &item)
	}

	return connAdd, connModify, connRemove, nil
}
