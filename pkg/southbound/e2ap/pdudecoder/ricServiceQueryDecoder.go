// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeRicServiceQueryPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*int32, types.RanFunctionRevisions, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rsq := e2apPdu.GetInitiatingMessage().GetValue().GetRicServiceQuery()
	if rsq == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have RICserviceQuery")
	}

	var transactionID int32
	rfAccepted := make(types.RanFunctionRevisions)
	for _, v := range rsq.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsAccepted) {
			ranFunctionsAcceptedIE := v.GetValue().GetRfIdl().GetValue()
			if ranFunctionsAcceptedIE != nil {
				// It's not mandatory
				for _, ranFunctionIDItemIe := range ranFunctionsAcceptedIE.GetValue() {
					ranFunctionIDItem := ranFunctionIDItemIe.GetValue().GetRfId()
					id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
					val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
					rfAccepted[id] = val
				}
			}
		}
	}

	return &transactionID, rfAccepted, nil
}
