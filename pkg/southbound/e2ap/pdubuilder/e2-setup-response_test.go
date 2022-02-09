// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestE2SetupResponse(t *testing.T) {
	//rfAccepted1 := make(types1.RanFunctionRevisions)
	//rfAccepted1[100] = 2
	////rfAccepted1[200] = 2
	////rfAccepted1[150] = 2
	//
	//rfRejected1 := make(types1.RanFunctionCauses)
	//rfRejected1[101] = &e2ap_ies.Cause{
	//	Cause: &e2ap_ies.Cause_Misc{
	//		Misc: e2ap_ies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
	//	},
	//}
	////rfRejected1[102] = &e2ap_ies.Cause{
	////	Cause: &e2ap_ies.Cause_Protocol{
	////		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
	////	},
	////}
	//
	//e2ncID13 := pdubuilder.CreateE2NodeComponentIDS1("S1-component")
	//e2nccaal1 := make([]*types1.E2NodeComponentConfigAdditionAckItem, 0)
	//ie11 := types1.E2NodeComponentConfigAdditionAckItem{
	//	E2NodeComponentConfigurationAck: e2ap_ies.E2NodeComponentConfigurationAck{
	//		UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
	//	},
	//	E2NodeComponentID:   e2ncID13,
	//	E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	//}
	//e2nccaal1 = append(e2nccaal1, &ie11)
	//
	//plmnID1 := [3]byte{0x79, 0x78, 0x70}
	//ricID1 := types1.RicIdentifier{
	//	RicIdentifierValue: []byte{0x4d, 0x20, 0x00},
	//	RicIdentifierLen:   20,
	//}
	//response1, err := pdubuilder.NewE2SetupResponse(1, plmnID1, ricID1, e2nccaal1)
	//assert.NilError(t, err)
	//assert.Assert(t, response1 != nil)
	//response1.SetRanFunctionAccepted(rfAccepted1).SetRanFunctionRejected(rfRejected1)
	//
	//e2apPdu, err := pdubuilder.CreateResponseE2apPdu(response1)
	//assert.NilError(t, err)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2SetupResponse E2AP PDU PER\n%v", hex.Dump(per))

	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	//rfAccepted[200] = 2
	//rfAccepted[150] = 2

	rfRejected := make(types.RanFunctionCauses)
	rfRejected[101] = &e2apies.Cause{
		Cause: &e2apies.Cause_Misc{
			Misc: e2apies.CauseMisc_CAUSE_MISC_HARDWARE_FAILURE,
		},
	}
	//rfRejected[102] = &e2apies.Cause{
	//	Cause: &e2apies.Cause_Protocol{
	//		Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_SEMANTIC_ERROR,
	//	},
	//}

	e2ncID3 := CreateE2NodeComponentIDS1("S1-component")
	e2nccaal := make([]*types.E2NodeComponentConfigAdditionAckItem, 0)
	ie1 := types.E2NodeComponentConfigAdditionAckItem{
		E2NodeComponentConfigurationAck: e2apies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2apies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
		},
		E2NodeComponentID:   e2ncID3,
		E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1,
	}
	e2nccaal = append(e2nccaal, &ie1)

	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0x4d, 0x20, 0x00},
		RicIdentifierLen:   20,
	}
	response, err := NewE2SetupResponse(1, plmnID, ricID, e2nccaal)
	assert.NilError(t, err)
	assert.Assert(t, response != nil)
	response.SetRanFunctionAccepted(rfAccepted).SetRanFunctionRejected(rfRejected)

	newE2apPdu, err := CreateResponseE2apPdu(response)
	assert.NilError(t, err)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupResponse E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, result1.String(), e2apPdu.String())
}
