// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"testing"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigRemovalAckListMsg() (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList, error) {

	e2nccal := e2ap_pdu_contents.E2NodeComponentConfigRemovalAckList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItemIes, 0),
	}

	e2nccai := &e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItemIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalAckItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem{
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
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}
	e2nccal.Value = append(e2nccal.Value, e2nccai)

	//if err := e2nccai.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2NodeComponentConfigAdditionItem %s", err.Error())
	//}
	return &e2nccal, nil
}

func Test_xerEncodingE2nodeComponentConfigRemovalAckList(t *testing.T) {

	e2nccal, err := createE2nodeComponentConfigRemovalAckListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigRemovalAckList PDU")

	xer, err := xerEncodeE2nodeComponentConfigRemovalAckList(e2nccal)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigRemovalAckList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigRemovalAckList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigRemovalAckList XER - decoded\n%v", result)
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.Value[0].Value.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.Value[0].Value.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.Value[0].Value.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentInterfaceType().Number(), result.Value[0].Value.GetE2NodeComponentInterfaceType().Number())
}

func Test_perEncodingE2nodeComponentConfigRemovalAckList(t *testing.T) {

	e2nccal, err := createE2nodeComponentConfigRemovalAckListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigRemovalAckList PDU")

	per, err := perEncodeE2nodeComponentConfigRemovalAckList(e2nccal)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigRemovalItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigRemovalAckList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigRemovalItem PER - decoded\n%v", result)
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.Value[0].Value.GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.Value[0].Value.GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.Value[0].Value.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccal.Value[0].Value.GetE2NodeComponentInterfaceType().Number(), result.Value[0].Value.GetE2NodeComponentInterfaceType().Number())
}
