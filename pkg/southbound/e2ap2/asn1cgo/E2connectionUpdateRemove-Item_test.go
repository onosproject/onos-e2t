// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func createE2connectionUpdateRemoveItemMsg() (*e2ap_pdu_contents.E2ConnectionUpdateRemoveItem, error) {

	e2connectionUpdateRemoveItem := e2ap_pdu_contents.E2ConnectionUpdateRemoveItem{
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
	}

	if err := e2connectionUpdateRemoveItem.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionUpdateRemoveItem %s", err.Error())
	}
	return &e2connectionUpdateRemoveItem, nil
}

func Test_xerEncodingE2connectionUpdateRemoveItem(t *testing.T) {

	e2connectionUpdateRemoveItem, err := createE2connectionUpdateRemoveItemMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateRemoveItem PDU")

	xer, err := xerEncodeE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem)
	assert.NilError(t, err)
	assert.Equal(t, 294, len(xer))
	t.Logf("E2connectionUpdateRemoveItem XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateRemoveItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateRemoveItem XER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlPort().GetValue(), result.GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlAddress().GetValue(), result.GetTnlInformation().GetTnlAddress().GetValue())
}

func Test_perEncodingE2connectionUpdateRemoveItem(t *testing.T) {

	e2connectionUpdateRemoveItem, err := createE2connectionUpdateRemoveItemMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateRemoveItem PDU")

	per, err := perEncodeE2connectionUpdateRemoveItem(e2connectionUpdateRemoveItem)
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("E2connectionUpdateRemoveItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateRemoveItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateRemoveItem PER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlPort().GetValue(), result.GetTnlInformation().GetTnlPort().GetValue())
	assert.Equal(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.DeepEqual(t, e2connectionUpdateRemoveItem.GetTnlInformation().GetTnlAddress().GetValue(), result.GetTnlInformation().GetTnlAddress().GetValue())
}
