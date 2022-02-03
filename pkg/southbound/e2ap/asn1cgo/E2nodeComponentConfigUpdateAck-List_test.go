// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigUpdateAckListMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList, error) {

	e2nodeComponentConfigUpdateAckItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem{
		E2NodeComponentInterfaceType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{
				E2NodeComponentInterfaceTypeW1: &e2ap_ies.E2NodeComponentInterfaceW1{
					NgENbDuId: &e2ap_ies.NgenbDuId{
						Value: 21,
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

	item := &e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &e2nodeComponentConfigUpdateAckItem,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2nodeComponentConfigUpdateAckList := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIes, 0),
	}
	e2nodeComponentConfigUpdateAckList.Value = append(e2nodeComponentConfigUpdateAckList.Value, item)

	//if err := e2nodeComponentConfigUpdateAckList.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateAckList %s", err.Error())
	//}
	return &e2nodeComponentConfigUpdateAckList, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateAckList(t *testing.T) {

	e2nodeComponentConfigUpdateAckList, err := createE2nodeComponentConfigUpdateAckListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckList PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdateAckList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateAckList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckList XER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())

}

func Test_perEncodingE2nodeComponentConfigUpdateAckList(t *testing.T) {

	e2nodeComponentConfigUpdateAckList, err := createE2nodeComponentConfigUpdateAckListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateAckList PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateAckList(e2nodeComponentConfigUpdateAckList)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdateAckList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateAckList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateAckList PER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetUpdateOutcome().Number())
	assert.Equal(t, e2nodeComponentConfigUpdateAckList.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentConfigurationAck().GetFailureCause().GetE2Node().Number())
}
