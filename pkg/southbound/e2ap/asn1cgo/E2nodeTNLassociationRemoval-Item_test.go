// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeTnlAssociationRemovalItemMsg() (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {

	e2ncc := e2ap_pdu_contents.E2NodeTnlassociationRemovalItem{
		TnlInformation: &e2ap_ies.Tnlinformation{
			TnlAddress: &asn1.BitString{
				Value: []byte{0xF0, 0xAB, 0x34, 0x9F},
				Len:   32,
			},
			TnlPort: &asn1.BitString{
				Value: []byte{0x00, 0x02},
				Len:   16,
			},
		},
		TnlInformationRic: &e2ap_ies.Tnlinformation{
			TnlAddress: &asn1.BitString{
				Value: []byte{0x00, 0x00, 0x00, 0x10},
				Len:   28,
			},
			TnlPort: &asn1.BitString{
				Value: []byte{0x00, 0x01},
				Len:   16,
			},
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validatiXn E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodiXnE2nodeTnlAssociationRemovalItem(t *testing.T) {

	e2ncc, err := createE2nodeTnlAssociationRemovalItemMsg()
	assert.NilError(t, err, "Error creatiXn E2nodeTnlAssociationRemovalItem PDU")

	xer, err := xerEncodeE2nodeTNLassociationRemovalItem(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeTnlAssociationRemovalItem XER\n%s", string(xer))

	result, err := xerDecodeE2nodeTNLassociationRemovalItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeTnlAssociationRemovalItem XER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetTnlInformation().GetTnlAddress().GetValue(), result.GetTnlInformation().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ncc.GetTnlInformation().GetTnlPort().GetValue(), result.GetTnlInformation().GetTnlPort().GetValue())
	assert.DeepEqual(t, e2ncc.GetTnlInformationRic().GetTnlAddress().GetValue(), result.GetTnlInformationRic().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ncc.GetTnlInformationRic().GetTnlPort().GetValue(), result.GetTnlInformationRic().GetTnlPort().GetValue())
}

func Test_perEncodiXnE2nodeTnlAssociationRemovalItem(t *testing.T) {

	e2ncc, err := createE2nodeTnlAssociationRemovalItemMsg()
	assert.NilError(t, err, "Error creatiXn E2nodeTnlAssociationRemovalItem PDU")

	per, err := perEncodeE2nodeTNLassociationRemovalItem(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeTnlAssociationRemovalItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeTNLassociationRemovalItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeTnlAssociationRemovalItem PER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetTnlInformation().GetTnlAddress().GetValue(), result.GetTnlInformation().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ncc.GetTnlInformation().GetTnlPort().GetValue(), result.GetTnlInformation().GetTnlPort().GetValue())
	assert.DeepEqual(t, e2ncc.GetTnlInformationRic().GetTnlAddress().GetValue(), result.GetTnlInformationRic().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ncc.GetTnlInformationRic().GetTnlPort().GetValue(), result.GetTnlInformationRic().GetTnlPort().GetValue())
}
