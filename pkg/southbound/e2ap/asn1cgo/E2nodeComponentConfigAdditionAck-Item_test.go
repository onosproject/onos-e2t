// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"testing"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigAdditionAckItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem, error) {

	e2nccai := e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem{
		E2NodeComponentConfigurationAck: &e2ap_ies.E2NodeComponentConfigurationAck{
			UpdateOutcome: e2ap_ies.UpdateOutcome_UPDATE_OUTCOME_SUCCESS,
			FailureCause: &e2ap_ies.Cause{
				Cause: &e2ap_ies.Cause_E2Node{
					E2Node: e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN,
				},
			},
		},
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeNg{
				E2NodeComponentInterfaceTypeNg: &e2ap_ies.E2NodeComponentInterfaceNg{
					AmfName: &e2ap_commondatatypes.Amfname{
						Value: "NgInterface",
					},
				},
			},
		},
		E2NodeComponentInterfaceType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG,
	}

	//if err := e2nccai.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2NodeComponentConfigAdditionItem %s", err.Error())
	//}
	return &e2nccai, nil
}

func Test_xerEncodingE2nodeComponentConfigAdditionAckItem(t *testing.T) {

	e2nccaai, err := createE2nodeComponentConfigAdditionAckItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigAdditionAckItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigAdditionAckItem(e2nccaai)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigAdditionAckItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigAdditionAckItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigAdditionAckItem XER - decoded\n%v", result)
	assert.Equal(t, e2nccaai.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nccaai.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
	assert.Equal(t, e2nccaai.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccaai.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
}

func Test_perEncodingE2nodeComponentConfigAdditionAckItem(t *testing.T) {

	e2nccaai, err := createE2nodeComponentConfigAdditionAckItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigAdditionAckItem PDU")

	per, err := perEncodeE2nodeComponentConfigAdditionAckItem(e2nccaai)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigAdditionAckItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigAdditionAckItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigAdditionAckItem PER - decoded\n%v", result)
	assert.Equal(t, e2nccaai.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nccaai.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
	assert.Equal(t, e2nccaai.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccaai.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
}
