// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func createE2connectionUpdateRemoveListMsg() (*e2ap_pdu_contents.E2ConnectionUpdateRemoveList, error) {

	e2connectionUpdateRemoveList := e2ap_pdu_contents.E2ConnectionUpdateRemoveList{
		Value: make([]*e2ap_pdu_contents.E2ConnectionUpdateRemoveItemIes, 0),
	}

	e2connectionUpdateRemoveItem := &e2ap_pdu_contents.E2ConnectionUpdateRemoveItemIes{
		Id:          int32(v1beta2.ProtocolIeIDE2connectionUpdateRemoveItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_pdu_contents.E2ConnectionUpdateRemoveItem{
			TnlInformation: &e2ap_ies.Tnlinformation{
				TnlAddress: &e2ap_commondatatypes.BitString{
					Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
					Len:   64,
				},
				TnlPort: &e2ap_commondatatypes.BitString{
					Value: []byte{0xae, 0x89},
					Len:   16,
				},
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	e2connectionUpdateRemoveList.Value = append(e2connectionUpdateRemoveList.Value, e2connectionUpdateRemoveItem)

	if err := e2connectionUpdateRemoveList.Validate(); err != nil {
		return nil, fmt.Errorf("error validating e2connectionUpdateRemoveList %s", err.Error())
	}
	return &e2connectionUpdateRemoveList, nil
}

func Test_xerEncodingE2connectionUpdateRemoveList(t *testing.T) {

	e2connectionUpdateRemoveList, err := createE2connectionUpdateRemoveListMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateRemoveItem PDU")

	xer, err := xerEncodeE2connectionUpdateRemoveList(e2connectionUpdateRemoveList)
	assert.NilError(t, err)
	assert.Equal(t, 644, len(xer))
	t.Logf("E2connectionUpdateRemoveList XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateRemoveList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateRemoveList XER - decoded\n%v", result)
	assert.Equal(t, len(e2connectionUpdateRemoveList.GetValue()), 1)
	assert.Equal(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
}

func Test_perEncodingE2connectionUpdateRemoveList(t *testing.T) {

	e2connectionUpdateRemoveList, err := createE2connectionUpdateRemoveListMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateRemoveList PDU")

	per, err := perEncodeE2connectionUpdateRemoveList(e2connectionUpdateRemoveList)
	assert.NilError(t, err)
	assert.Equal(t, 17, len(per))
	t.Logf("E2connectionUpdateRemoveList PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateRemoveList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateRemoveList PER - decoded\n%v", result)
	assert.Equal(t, len(e2connectionUpdateRemoveList.GetValue()), 1)
	assert.Equal(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveList.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue(), result.GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetValue())
}
