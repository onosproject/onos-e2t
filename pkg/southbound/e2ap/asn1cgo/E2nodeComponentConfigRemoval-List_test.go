// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

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

func createE2nodeComponentConfigRemovalListMsg() (*e2ap_pdu_contents.E2NodeComponentConfigRemovalList, error) {

	e2nccrl := e2ap_pdu_contents.E2NodeComponentConfigRemovalList{
		Value: make([]*e2ap_pdu_contents.E2NodeComponentConfigRemovalItemIes, 0),
	}
	e2nccri := &e2ap_pdu_contents.E2NodeComponentConfigRemovalItemIes{
		Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_pdu_contents.E2NodeComponentConfigRemovalItem{
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

	e2nccrl.Value = append(e2nccrl.Value, e2nccri)

	//if err := e2nccai.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2NodeComponentConfigAdditionItem %s", err.Error())
	//}
	return &e2nccrl, nil
}

func Test_xerEncodingE2nodeComponentRemovalList(t *testing.T) {

	e2nccrl, err := createE2nodeComponentConfigRemovalListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigRemovalList PDU")

	xer, err := xerEncodeE2nodeComponentConfigRemovalList(e2nccrl)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigRemovalList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigRemovalList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigRemovalList XER - decoded\n%v", result)
	assert.Equal(t, e2nccrl.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccrl.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
}

func Test_perEncodingE2nodeComponentConfigRemovalList(t *testing.T) {

	e2nccrl, err := createE2nodeComponentConfigRemovalListMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigRemovalList PDU")

	per, err := perEncodeE2nodeComponentConfigRemovalList(e2nccrl)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigRemovalList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigRemovalList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigRemovalList PER - decoded\n%v", result)
	assert.Equal(t, e2nccrl.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetValue()[0].GetValue().GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccrl.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number(), result.GetValue()[0].GetValue().GetE2NodeComponentInterfaceType().Number())
}
