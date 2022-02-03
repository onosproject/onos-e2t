// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigUpdateAckItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {

	e2nodeComponentConfigUpdateAckItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem{
		E2NodeComponentInterfaceType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{
				E2NodeComponentInterfaceTypeW1: &e2ap_ies.E2NodeComponentInterfaceW1{
					NgENbDuId: &e2ap_ies.NgenbDuId{
						Value: 2,
					},
				},
			},
		},
		E2NodeComponentConfigurationAck: &e2ap_ies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
			FailureCause: &e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_E2Node{
					E2Node: e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN,
				},
			},
		},
	}

	//if err := e2nodeComponentConfigUpdateAckItem.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateAckItem %s", err.Error())
	//}
	return &e2nodeComponentConfigUpdateAckItem, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateAckItem(t *testing.T) {

	e2nodeComponentConfigUpdateAckItem, err := createE2nodeComponentConfigUpdateAckItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdateAckItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateAckItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckItem XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
}

func Test_perEncodingE2nodeComponentConfigUpdateAckItem(t *testing.T) {

	e2nodeComponentConfigUpdateAckItem, err := createE2nodeComponentConfigUpdateAckItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckItem PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdateAckItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateAckItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckItem PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckItem.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
}
