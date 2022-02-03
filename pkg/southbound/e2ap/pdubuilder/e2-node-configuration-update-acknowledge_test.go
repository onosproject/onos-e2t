// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
)

func TestE2NodeConfigurationUpdateAck(t *testing.T) {

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
	newE2apPdu.GetSuccessfulOutcome().GetProcedureCode().GetE2NodeConfigurationUpdate().GetSuccessfulOutcome().
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

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU XER\n%s", string(xer))

	result1, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU XER - decoded\n%v", result1)
	assert.DeepEqual(t, newE2apPdu.String(), result1.String())

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU PER\n%v", hex.Dump(per))

	resultPer, err := asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	t.Logf("E2NodeConfigurationUpdateAck E2AP PDU PER - decoded\n%v", resultPer)
	assert.DeepEqual(t, newE2apPdu.String(), resultPer.String())
}
