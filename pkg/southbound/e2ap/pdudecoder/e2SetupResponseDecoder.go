// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2SetupResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*types.RicIdentity, types.RanFunctionIDs, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2setup := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2Setup()
	if e2setup == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}

	identifierIe := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes4()
	if identifierIe == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have id-GlobalE2node-ID")
	}
	ricIdentifier := types.RicIdentifier{
		RicIdentifierValue: types.RicIdentifierBits(identifierIe.GetValue().GetRicId().GetValue()),
		RicIdentifierLen:   types.RicIdentifierLen(identifierIe.GetValue().GetRicId().GetLen()),
	}
	plmnIDSlice := identifierIe.GetValue().GetPLmnIdentity().GetValue()
	plmnID := types.PlmnID{plmnIDSlice[0], plmnIDSlice[1], plmnIDSlice[2]}

	ricIdentity := types.RicIdentity{
		RicIdentifier: ricIdentifier,
		PlmnID:        plmnID,
	}

	rfAccepted := make(types.RanFunctionIDs)

	ranFunctionsAcceptedIE := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes9()
	if ranFunctionsAcceptedIE != nil {
		// It's not mandatory
		for _, ranFunctionIDItemIe := range ranFunctionsAcceptedIE.GetValue().GetValue() {
			ranFunctionIDItem := ranFunctionIDItemIe.GetRanFunctionIdItemIes6().GetValue()
			id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
			val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
			rfAccepted[id] = val
		}
	}

	return &ricIdentity, rfAccepted, nil
}
