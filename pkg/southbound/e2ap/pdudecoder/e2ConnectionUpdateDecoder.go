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

func DecodeE2connectionUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, []*types.E2ConnectionUpdateItem,
	[]*types.E2ConnectionUpdateItem, []*types.TnlInformation, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2cu := e2apPdu.GetInitiatingMessage().GetValue().GetE2ConnectionUpdate()
	if e2cu == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2connectionUpdate")
	}

	var transactionID int32
	connAdd := make([]*types.E2ConnectionUpdateItem, 0)
	connModify := make([]*types.E2ConnectionUpdateItem, 0)
	connRemove := make([]*types.TnlInformation, 0)
	for _, v := range e2cu.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDE2connectionUpdateAdd) {
			cal := v.GetValue().GetE2Cul().GetValue()
			for _, ie := range cal {
				item := types.E2ConnectionUpdateItem{}
				item.TnlInformation.TnlAddress = *ie.GetValue().GetE2Curi().GetTnlInformation().GetTnlAddress()
				item.TnlInformation.TnlPort = ie.GetValue().GetE2Curi().GetTnlInformation().GetTnlPort()
				item.TnlUsage = ie.GetValue().GetE2Curi().GetTnlUsage()
				connAdd = append(connAdd, &item)
			}
		}
		if v.Id == int32(v2.ProtocolIeIDE2connectionUpdateModify) {
			cml := v.GetValue().GetE2Cul().GetValue()
			for _, ie := range cml {
				item := types.E2ConnectionUpdateItem{}
				item.TnlInformation.TnlAddress = *ie.GetValue().GetE2Curi().GetTnlInformation().GetTnlAddress()
				item.TnlInformation.TnlPort = ie.GetValue().GetE2Curi().GetTnlInformation().GetTnlPort()
				item.TnlUsage = ie.GetValue().GetE2Curi().GetTnlUsage()
				connModify = append(connModify, &item)
			}
		}
		if v.Id == int32(v2.ProtocolIeIDE2connectionUpdateRemove) {
			crl := v.GetValue().GetE2Curl().GetValue()
			for _, ie := range crl {
				item := types.TnlInformation{}
				item.TnlAddress = *ie.GetValue().GetE2Curi().GetTnlInformation().GetTnlAddress()
				item.TnlPort = ie.GetValue().GetE2Curi().GetTnlInformation().GetTnlPort()
				connRemove = append(connRemove, &item)
			}
		}
	}

	return &transactionID, connAdd, connModify, connRemove, nil
}
