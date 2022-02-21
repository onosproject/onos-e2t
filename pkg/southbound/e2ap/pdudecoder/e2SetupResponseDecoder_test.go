// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeE2SetupResponsePdu(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2
	rfAccepted[150] = 2

	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Misc{
			Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	rfRejected[102] = &e2ap_ies.Cause{
		Cause: &e2ap_ies.Cause_Protocol{
			Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
		},
	}

	e2ncID3 := pdubuilder.CreateE2NodeComponentIDS1("S1-component")
	e2nccaal := make([]*types.E2NodeComponentConfigAdditionAckItem, 0)
	ie1 := types.E2NodeComponentConfigAdditionAckItem{
		E2NodeComponentConfigurationAck: e2ap_ies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
		},
		E2NodeComponentID:   e2ncID3,
		E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	}
	e2nccaal = append(e2nccaal, &ie1)

	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0x4d, 0x20, 0x00},
		RicIdentifierLen:   20,
	}
	response, err := pdubuilder.NewE2SetupResponse(1, plmnID, ricID, e2nccaal)
	assert.NilError(t, err)
	assert.Assert(t, response != nil)
	response.SetRanFunctionAccepted(rfAccepted).SetRanFunctionRejected(rfRejected)

	e2apPdu, err := pdubuilder.CreateResponseE2apPdu(response)
	assert.NilError(t, err)

	transactionID, ricIdentity, ranFunctionsAccepted, ranFunctionsRejected, additionAckList, err := DecodeE2SetupResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
	assert.DeepEqual(t, []byte{0x79, 0x78, 0x70}, []byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]})
	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
	assert.DeepEqual(t, []byte{0x4d, 0x20, 0x00}, []byte(ricIdentity.RicIdentifier.RicIdentifierValue))

	assert.Equal(t, 3, len(ranFunctionsAccepted))
	rfa100, ok := ranFunctionsAccepted[100]
	assert.Assert(t, ok, "expected a key '100'")
	assert.Equal(t, 2, int(rfa100))
	rfa200, ok := ranFunctionsAccepted[200]
	assert.Assert(t, ok, "expected a key '200'")
	assert.Equal(t, 2, int(rfa200))

	assert.Equal(t, 2, len(ranFunctionsRejected))
	rfr101, ok := ranFunctionsRejected[101]
	assert.Assert(t, ok, "expected a key '101'")
	assert.Equal(t, "CAUSE_MISC_OM_INTERVENTION", rfr101.GetMisc().String())
	rfr102, ok := ranFunctionsRejected[102]
	assert.Assert(t, ok, "expected a key '102'")
	assert.Equal(t, "CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR", rfr102.GetProtocol().String())

	if transactionID != nil {
		assert.Equal(t, int32(1), *transactionID)
	}
	assert.Equal(t, additionAckList[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1.Number())
	assert.Equal(t, additionAckList[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue(), "S1-component")
	assert.Equal(t, additionAckList[0].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS.Number())
	//assert.Equal(t, int32(additionAckList[0].E2NodeComponentConfigurationAck.FailureCause.GetProtocol()), int32(e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR))
}

func Test_DecodeE2SetupResponsePduNoOptional(t *testing.T) {
	e2ncID3 := pdubuilder.CreateE2NodeComponentIDS1("S1-component")
	e2nccaal := make([]*types.E2NodeComponentConfigAdditionAckItem, 0)
	ie1 := types.E2NodeComponentConfigAdditionAckItem{
		E2NodeComponentConfigurationAck: e2ap_ies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
		},
		E2NodeComponentID:   e2ncID3,
		E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	}
	e2nccaal = append(e2nccaal, &ie1)

	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0x4d, 0x20, 0x00},
		RicIdentifierLen:   20,
	}
	response, err := pdubuilder.NewE2SetupResponse(11, plmnID, ricID, e2nccaal)
	assert.NilError(t, err)
	assert.Assert(t, response != nil)

	e2apPdu, err := pdubuilder.CreateResponseE2apPdu(response)
	assert.NilError(t, err)

	transactionID, ricIdentity, ranFunctionsAccepted, ranFunctionsRejected, additionAckList, err := DecodeE2SetupResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 54, 55 & 56
	assert.DeepEqual(t, []byte{0x79, 0x78, 0x70}, []byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]})
	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
	assert.DeepEqual(t, []byte{0x4d, 0x20, 0x00}, []byte(ricIdentity.RicIdentifier.RicIdentifierValue))

	assert.Equal(t, 0, len(ranFunctionsAccepted))
	assert.Equal(t, 0, len(ranFunctionsRejected))

	assert.Equal(t, additionAckList[0].E2NodeComponentType.Number(), e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1.Number())
	assert.Equal(t, additionAckList[0].E2NodeComponentID.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue(), "S1-component")
	assert.Equal(t, additionAckList[0].E2NodeComponentConfigurationAck.UpdateOutcome.Number(), e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS.Number())

	if transactionID != nil {
		assert.Equal(t, int32(11), *transactionID)
	}
}
