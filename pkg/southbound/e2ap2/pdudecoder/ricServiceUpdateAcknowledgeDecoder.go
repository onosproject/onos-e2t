// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicServiceUpdateAcknowledgePdu(e2apPdu *e2ap_pdu_descriptions.E2ApPdu) (types.RanFunctionRevisions,
	types.RanFunctionCauses, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rsua := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetRicServiceUpdate()
	if rsua == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have RICserviceUpdateAcknowledge")
	}

	ranFunctionsAcceptedList := make(types.RanFunctionRevisions)
	rfal := rsua.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes9().GetValue().GetValue()
	for _, ranFunctionIDItemIe := range rfal {
		ranFunctionIDItem := ranFunctionIDItemIe.GetRanFunctionIdItemIes6().GetValue()
		id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
		val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
		ranFunctionsAcceptedList[id] = val
	}

	rfrl := rsua.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes13().GetValue().GetValue()
	causes := make(types.RanFunctionCauses)
	for _, rfri := range rfrl {
		causes[types.RanFunctionID(rfri.GetRanFunctionIdcauseItemIes7().GetValue().GetRanFunctionId().GetValue())] = rfri.GetRanFunctionIdcauseItemIes7().GetValue().GetCause()
	}

	return ranFunctionsAcceptedList, causes, nil
}
