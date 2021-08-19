// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap2/types"
)

func DecodeE2SetupResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*int32, *types.RicIdentity, types.RanFunctionRevisions,
	types.RanFunctionCauses, []*types.E2NodeComponentConfigUpdateAckItem, error) {
	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	//}

	e2setup := e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2Setup()
	if e2setup == nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}

	identifierIe := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes4()
	if identifierIe == nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-GlobalE2node-ID")
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

	e2nccual := make([]*types.E2NodeComponentConfigUpdateAckItem, 0)
	list := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes35().GetValue().GetValue()
	for _, ie := range list {
		e2nccuai := types.E2NodeComponentConfigUpdateAckItem{}
		e2nccuai.E2NodeComponentType = ie.GetValue().GetE2NodeComponentType()
		e2nccuai.E2NodeComponentID = ie.GetValue().GetE2NodeComponentId()
		e2nccuai.E2NodeComponentConfigUpdateAck = types.E2NodeComponentConfigUpdateAck{
			UpdateOutcome: ie.GetValue().GetE2NodeComponentConfigUpdateAck().GetUpdateOutcome(),
			FailureCause:  ie.GetValue().GetE2NodeComponentConfigUpdateAck().GetFailureCause(),
		}

		e2nccual = append(e2nccual, &e2nccuai)
	}


	transactionID := e2setup.GetSuccessfulOutcome().GetProtocolIes().GetE2ApProtocolIes49().GetValue().GetValue()

	return &transactionID, &ricIdentity, rfAccepted, rfRejected, e2nccual, nil
}
