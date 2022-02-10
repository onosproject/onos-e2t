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

func DecodeRicServiceUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (*int32, types.RanFunctions, types.RanFunctionRevisions,
	types.RanFunctions, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rsu := e2apPdu.GetInitiatingMessage().GetValue().GetRicServiceUpdate()
	if rsu == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have RICserviceUpdate")
	}

	var transactionID int32
	ranFunctionsAddedList := make(types.RanFunctions)
	ranFunctionsDeletedList := make(types.RanFunctionRevisions)
	ranFunctionsModifiedList := make(types.RanFunctions)
	for _, v := range rsu.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsAdded) {
			rfal := v.GetValue().GetRfl().GetValue()
			for _, ie := range rfal.GetValue() {
				val := ie.GetValue().GetRfi()
				ranFunctionsAddedList[types.RanFunctionID(val.GetRanFunctionId().GetValue())] = types.RanFunctionItem{
					Description: val.GetRanFunctionDefinition().GetValue(),
					Revision:    types.RanFunctionRevision(val.GetRanFunctionRevision().GetValue()),
					OID:         types.RanFunctionOID(val.GetRanFunctionOid().GetValue()),
				}
			}
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsModified) {
			rfml := v.GetValue().GetRfl().GetValue()
			for _, ie := range rfml.GetValue() {
				val := ie.GetValue().GetRfi()
				ranFunctionsModifiedList[types.RanFunctionID(val.GetRanFunctionId().GetValue())] = types.RanFunctionItem{
					Description: val.GetRanFunctionDefinition().GetValue(),
					Revision:    types.RanFunctionRevision(val.GetRanFunctionRevision().GetValue()),
					OID:         types.RanFunctionOID(val.GetRanFunctionOid().GetValue()),
				}
			}
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsDeleted) {
			rfdl := v.GetValue().GetRfidl().GetValue()
			for _, ranFunctionIDItemIe := range rfdl.GetValue() {
				ranFunctionIDItem := ranFunctionIDItemIe.GetValue().GetRfId()
				id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
				val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
				ranFunctionsDeletedList[id] = val
			}
		}
	}

	return &transactionID, ranFunctionsAddedList, ranFunctionsDeletedList, ranFunctionsModifiedList, nil
}
