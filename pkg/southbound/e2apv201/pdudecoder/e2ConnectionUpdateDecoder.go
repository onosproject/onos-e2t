// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"

	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2apv201/types"
)

func DecodeE2connectionUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, []*types.E2ConnectionUpdateItem,
	[]*types.E2ConnectionUpdateItem, []*types.TnlInformation, error) {
	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	e2cu := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate()
	if e2cu == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2connectionUpdate")
	}

	connAdd := make([]*types.E2ConnectionUpdateItem, 0)
	cal := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes44().GetValue().GetValue()
	for _, ie := range cal {
		item := types.E2ConnectionUpdateItem{}
		item.TnlInformation.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlInformation.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		item.TnlUsage = ie.GetValue().GetTnlUsage()
		connAdd = append(connAdd, &item)
	}

	connModify := make([]*types.E2ConnectionUpdateItem, 0)
	cml := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes45().GetValue().GetValue()
	for _, ie := range cml {
		item := types.E2ConnectionUpdateItem{}
		item.TnlInformation.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlInformation.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		item.TnlUsage = ie.GetValue().GetTnlUsage()
		connModify = append(connModify, &item)
	}

	connRemove := make([]*types.TnlInformation, 0)
	crl := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes46().GetValue().GetValue()
	for _, ie := range crl {
		item := types.TnlInformation{}
		item.TnlAddress = *ie.GetValue().GetTnlInformation().GetTnlAddress()
		item.TnlPort = *ie.GetValue().GetTnlInformation().GetTnlPort()
		connRemove = append(connRemove, &item)
	}

	transactionID := e2cu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue()

	return &transactionID, connAdd, connModify, connRemove, nil
}
