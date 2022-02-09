// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2SetupResponsePdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*int32, *types.RicIdentity, types.RanFunctionRevisions,
	types.RanFunctionCauses, []*types.E2NodeComponentConfigAdditionAckItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2setup := e2apPdu.GetSuccessfulOutcome().GetValue().GetE2Setup()
	if e2setup == nil {
		return nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}

	var transactionID int32
	var ricIdentity types.RicIdentity
	rfAccepted := make(types.RanFunctionRevisions)
	rfRejected := make(types.RanFunctionCauses)
	e2nccual := make([]*types.E2NodeComponentConfigAdditionAckItem, 0)
	for _, v := range e2setup.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDGlobalRicID) {
			identifierIe := v.GetValue().GetGRicId()
			if identifierIe == nil {
				return nil, nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-GlobalE2node-ID")
			}
			ricIdentifier := types.RicIdentifier{
				RicIdentifierValue: types.RicIdentifierBits(identifierIe.GetRicId().GetValue()),
				RicIdentifierLen:   types.RicIdentifierLen(identifierIe.GetRicId().GetLen()),
			}
			plmnIDSlice := identifierIe.GetPLmnIdentity().GetValue()
			plmnID := types.PlmnID{plmnIDSlice[0], plmnIDSlice[1], plmnIDSlice[2]}

			ricIdentity.RicIdentifier = ricIdentifier
			ricIdentity.PlmnID = plmnID
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsAccepted) {
			if v.GetValue().GetRfIdl().GetValue() != nil {
				// It's not mandatory
				for _, ranFunctionIDItemIe := range v.GetValue().GetRfIdl().GetValue() {
					ranFunctionIDItem := ranFunctionIDItemIe.GetValue().GetRfId()
					id := types.RanFunctionID(ranFunctionIDItem.GetRanFunctionId().GetValue())
					val := types.RanFunctionRevision(ranFunctionIDItem.GetRanFunctionRevision().GetValue())
					rfAccepted[id] = val
				}
			}
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsRejected) {
			if v.GetValue().GetRfIdcl() != nil {
				// It's not mandatory
				for _, ranFunctionIDRejectedItemIe := range v.GetValue().GetRfIdcl().GetValue() {
					ranFunctionIDcauseItem := ranFunctionIDRejectedItemIe.GetValue().GetRfIdci()
					id := types.RanFunctionID(ranFunctionIDcauseItem.GetRanFunctionId().GetValue())
					rfRejected[id] = &e2apies.Cause{
						Cause: &e2apies.Cause_Misc{Misc: e2apies.CauseMisc_CAUSE_MISC_OM_INTERVENTION},
					}
				}
			}
		}
		if v.Id == int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck) {
			list := v.GetValue().GetE2Nccaal().GetValue()
			for _, ie := range list {
				e2nccuai := types.E2NodeComponentConfigAdditionAckItem{}
				e2nccuai.E2NodeComponentType = ie.GetValue().GetE2Nccaai().GetE2NodeComponentInterfaceType()
				e2nccuai.E2NodeComponentID = ie.GetValue().GetE2Nccaai().GetE2NodeComponentId()
				e2nccuai.E2NodeComponentConfigurationAck = e2apies.E2NodeComponentConfigurationAck{
					UpdateOutcome: ie.GetValue().GetE2Nccaai().GetE2NodeComponentConfigurationAck().GetUpdateOutcome(),
					FailureCause:  ie.GetValue().GetE2Nccaai().GetE2NodeComponentConfigurationAck().GetFailureCause(),
				}

				e2nccual = append(e2nccual, &e2nccuai)
			}
		}
	}

	return &transactionID, &ricIdentity, rfAccepted, rfRejected, e2nccual, nil
}
