// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeE2nodeConfigurationUpdateAcknowledgePdu(t *testing.T) {
	grnID, err := pdubuilder.CreateGlobalNgRanNodeIDGnb([]byte{0x01, 0x02, 0x03}, &asn1.BitString{
		Value: []byte{0xAB, 0xCD, 0xEF, 0xFF},
		Len:   32,
	})
	assert.NilError(t, err)

	e2ncID1 := pdubuilder.CreateE2NodeComponentIDF1(21)
	e2ncID2 := pdubuilder.CreateE2NodeComponentIDXn(grnID)

	e2apPdu, err := pdubuilder.CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)
	e2apPdu.GetSuccessfulOutcome().GetValue().GetE2NodeConfigurationUpdate().
		SetE2nodeComponentConfigUpdateAck([]*types.E2NodeComponentConfigUpdateAckItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1,
				E2NodeComponentID: e2ncID1,
				E2NodeComponentConfigurationAck: types.E2NodeComponentConfigurationAck{
					UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_FAILURE,
					FailureCause: &e2ap_ies.Cause{
						Cause: &e2ap_ies.Cause_Protocol{
							Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
						},
					},
				}},
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN,
				E2NodeComponentID: e2ncID2,
				E2NodeComponentConfigurationAck: types.E2NodeComponentConfigurationAck{
					UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
					FailureCause: &e2ap_ies.Cause{
						Cause: &e2ap_ies.Cause_Protocol{
							Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
						},
					},
				}}})

	transactionID, additionAckList, err := DecodeE2nodeConfigurationUpdateAcknowledgePdu(e2apPdu)
	assert.NilError(t, err)

	assert.Equal(t, additionAckList[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1.Number())
	assert.Equal(t, additionAckList[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeF1().GetGNbDuId().GetValue(), int64(21))
	assert.Equal(t, additionAckList[0].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_FAILURE.Number())

	assert.Equal(t, additionAckList[1].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN.Number())
	assert.DeepEqual(t, additionAckList[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetPlmnId().GetValue(), []byte{0x01, 0x02, 0x03})
	assert.DeepEqual(t, additionAckList[1].E2NodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetGnbId().GetGnbId().GetValue(), []byte{0xAB, 0xCD, 0xEF, 0xFF})
	assert.Equal(t, additionAckList[1].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS.Number())
	assert.Equal(t, additionAckList[1].E2NodeComponentConfigurationAck.FailureCause.GetProtocol().Number(), e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE.Number())

	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
}
