// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeTnlAssociationRemovalListMsg() (*e2ap_pdu_contents.E2NodeTnlassociationRemovalList, error) {

	e2ntarl := e2ap_pdu_contents.E2NodeTnlassociationRemovalList{
		Value: make([]*e2ap_pdu_contents.E2NodeTnlassociationRemovalItemIes, 0),
	}

	e2ntari := &e2ap_pdu_contents.E2NodeTnlassociationRemovalItemIes{
		Id:          int32(v2.ProtocolIeIDE2nodeTNLassociationRemovalItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2ap_pdu_contents.E2NodeTnlassociationRemovalItem{
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
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2ntarl.Value = append(e2ntarl.Value, e2ntari)

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validatiXn E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ntarl, nil
}

func Test_xerEncodiXnE2nodeTnlAssociationRemovalList(t *testing.T) {

	e2ntarl, err := createE2nodeTnlAssociationRemovalListMsg()
	assert.NilError(t, err, "Error creatiXn E2nodeTnlAssociationRemovalList PDU")

	xer, err := xerEncodeE2nodeTNLassociationRemovalList(e2ntarl)
	assert.NilError(t, err)
	t.Logf("E2nodeTnlAssociationRemovalList XER\n%s", string(xer))

	result, err := xerDecodeE2nodeTNLassociationRemovalList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeTnlAssociationRemovalList XER - decoded\n%v", result)
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlPort().GetValue())
}

func Test_perEncodiXnE2nodeTnlAssociationRemovalList(t *testing.T) {

	e2ntarl, err := createE2nodeTnlAssociationRemovalListMsg()
	assert.NilError(t, err, "Error creatiXn E2nodeTnlAssociationRemovalList PDU")

	per, err := perEncodeE2nodeTNLassociationRemovalList(e2ntarl)
	assert.NilError(t, err)
	t.Logf("E2nodeTnlAssociationRemovalList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeTNLassociationRemovalList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeTnlAssociationRemovalList PER - decoded\n%v", result)
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlAddress().GetValue())
	assert.DeepEqual(t, e2ntarl.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformationRic().GetTnlPort().GetValue())
}
