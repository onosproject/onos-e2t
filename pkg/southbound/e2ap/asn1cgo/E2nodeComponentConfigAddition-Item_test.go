// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"testing"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigAdditionItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigAdditionItem, error) {

	e2nccai := e2ap_pdu_contents.E2NodeComponentConfigAdditionItem{
		E2NodeComponentConfiguration: &e2ap_ies.E2NodeComponentConfiguration{
			E2NodeComponentRequestPart:  []byte{0x01, 0x02, 0x03},
			E2NodeComponentResponsePart: []byte{0x03, 0x02, 0x01},
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

func Test_xerEncodingE2nodeComponentConfigAdditionItem(t *testing.T) {

	e2nccai, err := createE2nodeComponentConfigAdditionItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigAdditionItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigAdditionItem(e2nccai)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigAdditionItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigAdditionItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigAdditionItem XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nccai.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2nccai.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.Equal(t, e2nccai.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccai.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
}

func Test_perEncodingE2nodeComponentConfigAdditionItem(t *testing.T) {

	e2nccai, err := createE2nodeComponentConfigAdditionItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigAdditionItem PDU")

	per, err := perEncodeE2nodeComponentConfigAdditionItem(e2nccai)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigAdditionItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigAdditionItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigAdditionItem PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nccai.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2nccai.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart(), result.GetE2NodeComponentConfiguration().GetE2NodeComponentRequestPart())
	assert.Equal(t, e2nccai.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccai.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
}
