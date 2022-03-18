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

func DecodeE2connectionUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, []*types.E2ConnectionUpdateItem,
	[]*types.E2ConnectionSetupFailedItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2cua := e2apPdu.GetSuccessfulOutcome().GetValue().GetE2ConnectionUpdate()
	if e2cua == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2connectionUpdateAcknowledge")
	}

	var transactionID int32
	connSetup := make([]*types.E2ConnectionUpdateItem, 0)
	connSetupFail := make([]*types.E2ConnectionSetupFailedItem, 0)
	for _, v := range e2cua.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTransactionId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDE2connectionSetup) {
			list := v.GetValue().GetE2ConnectionSetup().GetValue()
			for _, ie := range list {
				item := types.E2ConnectionUpdateItem{}
				item.TnlInformation.TnlAddress = *ie.GetValue().GetE2ConnectionUpdateItem().GetTnlInformation().GetTnlAddress()
				item.TnlInformation.TnlPort = ie.GetValue().GetE2ConnectionUpdateItem().GetTnlInformation().GetTnlPort()
				item.TnlUsage = ie.GetValue().GetE2ConnectionUpdateItem().GetTnlUsage()
				connSetup = append(connSetup, &item)
			}
		}
		if v.Id == int32(v2.ProtocolIeIDE2connectionSetupFailed) {
			failedList := v.GetValue().GetE2ConnectionSetupFailed().GetValue()
			for _, ie := range failedList {
				item := types.E2ConnectionSetupFailedItem{}
				item.TnlInformation.TnlAddress = *ie.GetValue().GetE2ConnectionSetupFailedItem().GetTnlInformation().GetTnlAddress()
				item.TnlInformation.TnlPort = ie.GetValue().GetE2ConnectionSetupFailedItem().GetTnlInformation().GetTnlPort()
				item.Cause = *ie.GetValue().GetE2ConnectionSetupFailedItem().GetCause()
				connSetupFail = append(connSetupFail, &item)
			}
		}
	}

	return &transactionID, connSetup, connSetupFail, nil
}
