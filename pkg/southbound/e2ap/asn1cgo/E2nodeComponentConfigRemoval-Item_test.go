// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"testing"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigRemovalItemMsg() (*e2ap_pdu_contents.E2NodeComponentConfigRemovalItem, error) {

	e2nccri := e2ap_pdu_contents.E2NodeComponentConfigRemovalItem{
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
	return &e2nccri, nil
}

func Test_xerEncodingE2nodeComponentRemovalItem(t *testing.T) {

	e2nccri, err := createE2nodeComponentConfigRemovalItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigRemovalItem PDU")

	xer, err := xerEncodeE2nodeComponentConfigRemovalItem(e2nccri)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigRemovalItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigRemovalItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigRemovalItem XER - decoded\n%v", result)
	assert.Equal(t, e2nccri.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccri.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
}

func Test_perEncodingE2nodeComponentConfigRemovalItem(t *testing.T) {

	e2nccri, err := createE2nodeComponentConfigRemovalItemMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigRemovalItem PDU")

	per, err := perEncodeE2nodeComponentConfigRemovalItem(e2nccri)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigRemovalItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigRemovalItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigRemovalItem PER - decoded\n%v", result)
	assert.Equal(t, e2nccri.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentId().GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())
	assert.Equal(t, e2nccri.GetE2NodeComponentInterfaceType().Number(), result.GetE2NodeComponentInterfaceType().Number())
}
