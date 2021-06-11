// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicServiceUpdatePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (types.RanFunctions, types.RanFunctionRevisions,
	types.RanFunctions, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rsu := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate()
	if rsu == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have RICserviceUpdate")
	}

	ranFunctionsAddedList := make(types.RanFunctions)
	rfal := rsu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes10().GetRanFunctionsAddedList().GetValue()
	for _, ie := range rfal {
		val := ie.GetE2ApProtocolIes10().GetValue()
		ranFunctionsAddedList[types.RanFunctionID(val.GetRanFunctionId().GetValue())] = types.RanFunctionItem{
			Description: val.GetRanFunctionDefinition().GetValue(),
			Revision:    types.RanFunctionRevision(val.GetRanFunctionRevision().GetValue()),
			OID:         val.GetRanFunctionOid().GetValue(),
		}
	}

	ranFunctionsDeletedList := make(types.RanFunctionRevisions)
	rfdl := rsu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes11().GetRanFunctionsDeletedList().GetValue()
	for _, ranFunctionIDItemIe := range rfdl {
		ranFunctionIDItem := ranFunctionIDItemIe.GetRanFunctionIdItemIes6().GetValue()
		id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
		val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
		ranFunctionsDeletedList[id] = val
	}

	ranFunctionsModifiedList := make(types.RanFunctions)
	rfml := rsu.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes12().GetRanFunctionsModifiedList().GetValue()
	for _, ie := range rfml {
		val := ie.GetE2ApProtocolIes10().GetValue()
		ranFunctionsModifiedList[types.RanFunctionID(val.GetRanFunctionId().GetValue())] = types.RanFunctionItem{
			Description: val.GetRanFunctionDefinition().GetValue(),
			Revision:    types.RanFunctionRevision(val.GetRanFunctionRevision().GetValue()),
			OID:         val.GetRanFunctionOid().GetValue(),
		}
	}

	return ranFunctionsAddedList, ranFunctionsDeletedList, ranFunctionsModifiedList, nil
}
