// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestE2NodeConfigurationUpdateAck(t *testing.T) {
	//grnID1, err := pdubuilder.CreateGlobalNgRanNodeIDGnb([]byte{0x01, 0x02, 0x03}, &asn1.BitString{
	//	Value: []byte{0xAB, 0xCD, 0xEF, 0xFF},
	//	Len:   32,
	//})
	//assert.NilError(t, err)
	//
	//e2ncID11 := pdubuilder.CreateE2NodeComponentIDF1(21)
	//e2ncID12 := pdubuilder.CreateE2NodeComponentIDXn(grnID1)
	//
	//e2apPdu, err := pdubuilder.CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//e2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetSuccessfulOutcome().
	//	SetE2nodeComponentConfigUpdateAck([]*types1.E2NodeComponentConfigUpdateAckItem{
	//		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1,
	//			E2NodeComponentID: e2ncID11,
	//			E2NodeComponentConfigurationAck: types1.E2NodeComponentConfigurationAck{
	//				UpdateOutcome: e2apies.UpdateOutcome_UPDATE_OUTCOME_FAILURE,
	//				//FailureCause: e2apies.Cause{
	//				//	Cause: &e2ap_ies.CauseProtocol{
	//				//		Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
	//				//	},
	//				//},
	//			}},
	//		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN,
	//			E2NodeComponentID: e2ncID12,
	//			E2NodeComponentConfigurationAck: types1.E2NodeComponentConfigurationAck{
	//				UpdateOutcome: e2apies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
	//				FailureCause: &e2apies.Cause{
	//					Cause: &e2apies.Cause_Protocol{
	//						Protocol: e2apies.CauseProtocol_CAUSE_PROTOCOL_ABSTRACT_SYNTAX_ERROR_FALSELY_CONSTRUCTED_MESSAGE,
	//					},
	//				},
	//			}}})
	//t.Logf("E2NodeConfigurationUpdateAck E2AP PDU \n%v", e2apPdu)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2NodeConfigurationUpdateAck E2AP PDU PER\n%v", hex.Dump(per))

	grnID, err := CreateGlobalNgRanNodeIDGnb([]byte{0x01, 0x02, 0x03}, &asn1.BitString{
		Value: []byte{0xAB, 0xCD, 0xEF, 0xFF},
		Len:   32,
	})
	assert.NilError(t, err)

	e2ncID1 := CreateE2NodeComponentIDF1(21)
	e2ncID2 := CreateE2NodeComponentIDXn(grnID)

	newE2apPdu, err := CreateE2NodeConfigurationUpdateAcknowledgeE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)
	newE2apPdu.GetSuccessfulOutcome().GetValue().GetE2NodeConfigurationUpdate().
		SetE2nodeComponentConfigUpdateAck([]*types.E2NodeComponentConfigUpdateAckItem{
			{E2NodeComponentType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1,
				E2NodeComponentID: e2ncID1,
				E2NodeComponentConfigurationAck: types.E2NodeComponentConfigurationAck{
					UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_FAILURE,
					//FailureCause: e2ap_ies.Cause{
					//	Cause: &e2ap_ies.Cause_Protocol{
					//		Protocol: e2ap_ies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR,
					//	},
					//},
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
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU \n%v", newE2apPdu)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

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
	//t.Logf("E2NodeConfigurationUpdateAck E2AP PDU PER - decoded\n%v", result)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}
