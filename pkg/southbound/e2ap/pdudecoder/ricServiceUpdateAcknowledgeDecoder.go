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

func DecodeRicServiceUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, types.RanFunctionRevisions,
	types.RanFunctionCauses, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rsua := e2apPdu.GetSuccessfulOutcome().GetValue().GetRicServiceUpdate()
	if rsua == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have RICserviceUpdateAcknowledge")
	}

	var transactionID int32
	ranFunctionsAcceptedList := make(types.RanFunctionRevisions)
	causes := make(types.RanFunctionCauses)
	for _, v := range rsua.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsAccepted) {
			rfal := v.GetValue().GetRfIdl().GetValue().GetValue()
			for _, ranFunctionIDItemIe := range rfal {
				ranFunctionIDItem := ranFunctionIDItemIe.GetValue().GetRfId()
				id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
				val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
				ranFunctionsAcceptedList[id] = val
			}
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsRejected) {
			rfrl := v.GetValue().GetRfIdcl().GetValue().GetValue()
			for _, rfri := range rfrl {
				causes[types.RanFunctionID(rfri.GetValue().GetRfIdci().GetRanFunctionId().GetValue())] = rfri.GetValue().GetRfIdci().GetCause()
			}
		}
	}

	return &transactionID, ranFunctionsAcceptedList, causes, nil
}
