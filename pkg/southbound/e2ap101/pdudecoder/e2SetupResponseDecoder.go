// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func DecodeE2SetupResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*types.RicIdentity, types.RanFunctionRevisions, types.RanFunctionCauses, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2setup := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2Setup()
	if e2setup == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}

	identifierIe := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes4()
	if identifierIe == nil {
		return nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-GlobalE2node-ID")
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

	rfAccepted := make(types.RanFunctionRevisions)
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

	rfRejected := make(types.RanFunctionCauses)
	ranFunctionsRejectedIE := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes13()
	if ranFunctionsRejectedIE != nil {
		// It's not mandatory
		for _, ranFunctionIDRejectedItemIe := range ranFunctionsRejectedIE.GetValue().GetValue() {
			ranFunctionIDcauseItem := ranFunctionIDRejectedItemIe.GetRanFunctionIdcauseItemIes7().GetValue()
			id := types.RanFunctionID(ranFunctionIDcauseItem.GetRanFunctionId().GetValue())
			rfRejected[id] = &e2apies.Cause{
				Cause: &e2apies.Cause_Misc{Misc: e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION},
			}
		}
	}

	return &ricIdentity, rfAccepted, rfRejected, nil
}
