// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigUpdateItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {

	e2nodeComponentConfigUpdateItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateItem{
		E2NodeComponentInterfaceType: e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1,
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId{
			E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{
				E2NodeComponentInterfaceTypeW1: &e2ap_ies.E2NodeComponentInterfaceW1{
					NgENbDuId: &e2ap_ies.NgenbDuId{
						Value: 31,
					},
				},
			},
		},
		E2NodeComponentConfiguration: &e2ap_ies.E2NodeComponentConfiguration{
			E2NodeComponentRequestPart:  []byte{0x00, 0x01},
			E2NodeComponentResponsePart: []byte{0x02, 0x03},
		},
	}

	//if err := e2nodeComponentConfigUpdateItem.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateItem %s", err.Error())
	//}
	return &e2nodeComponentConfigUpdateItem, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateItem(t *testing.T) {

	e2nodeComponentConfigUpdateItem, err := createE2nodeComponentConfigUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	assert.NilError(t, err)
	//assert.Equal(t, 675, len(xer))
	t.Logf("E2nodeComponentConfigUpdateItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateItem XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
}

func Test_perEncodingE2nodeComponentConfigUpdateItem(t *testing.T) {

	e2nodeComponentConfigUpdateItem, err := createE2nodeComponentConfigUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateItem PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	assert.NilError(t, err)
	//assert.Equal(t, 28, len(per))
	t.Logf("E2nodeComponentConfigUpdateItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateItem PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.Equal(t, e2nodeComponentConfigUpdateItem.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())
}
