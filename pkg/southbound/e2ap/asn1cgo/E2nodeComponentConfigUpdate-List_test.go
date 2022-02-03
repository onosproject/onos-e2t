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

func createE2nodeComponentConfigUpdateListMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateList, error) {

	e2nodeComponentConfigUpdateItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateItem{
		E2NodeComponentInterfaceType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{
				E2NodeComponentInterfaceTypeW1: &e2ap_ies.E2NodeComponentInterfaceW1{
					NgENbDuId: &e2ap_ies.NgenbDuId{
						Value: 153,
					},
				},
			},
		},
		E2NodeComponentConfiguration: &e2ap_ies.E2NodeComponentConfiguration{
			E2NodeComponentRequestPart:  []byte{0x00, 0x01},
			E2NodeComponentResponsePart: []byte{0x02, 0x03},
		},
	}

	item := e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       &e2nodeComponentConfigUpdateItem,
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2nodeComponentConfigUpdateList := e2ap_pdu_contents.E2NodeComponentConfigUpdateList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIes, 0),
	}
	e2nodeComponentConfigUpdateList.Value = append(e2nodeComponentConfigUpdateList.Value, &item)

	//if err := e2nodeComponentConfigUpdateList.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateList %s", err.Error())
	//}
	return &e2nodeComponentConfigUpdateList, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateList(t *testing.T) {

	e2nodeComponentConfigUpdateList, err := createE2nodeComponentConfigUpdateListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateList PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdateList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateList XER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
}

func Test_perEncodingE2nodeComponentConfigUpdateList(t *testing.T) {

	e2nodeComponentConfigUpdateList, err := createE2nodeComponentConfigUpdateListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateList PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateList(e2nodeComponentConfigUpdateList)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdateList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateList PER - decoded\n%v", result)
	assert.Equal(t, 1, len(result.GetValue()))
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetValue()[0].GetValue().GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.Equal(t, e2nodeComponentConfigUpdateList.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
}
