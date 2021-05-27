// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeRicServiceQueryPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (types.RanFunctionRevisions, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	rsq := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceQuery()
	if rsq == nil {
		return nil, fmt.Errorf("error E2APpdu does not have RICserviceQuery")
	}

	rfAccepted := make(types.RanFunctionRevisions)
	ranFunctionsAcceptedIE := rsq.GetInitiatingMessage().GetProtocolIes().GetRicserviceQueryIes9()
	if ranFunctionsAcceptedIE != nil {
		// It's not mandatory
		for _, ranFunctionIDItemIe := range ranFunctionsAcceptedIE.GetValue().GetValue() {
			ranFunctionIDItem := ranFunctionIDItemIe.GetRanFunctionIdItemIes6().GetValue()
			id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
			val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
			rfAccepted[id] = val
		}
	}

	return rfAccepted, nil
}
