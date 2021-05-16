// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"

	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2ap_pdu_contents/pdubuilder"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func createE2connectionUpdateItemMsg() (*e2ap_pdu_contents.E2ConnectionUpdateItem, error) {

	e2connectionUpdateItem := e2ap_pdu_contents.E2ConnectionUpdateItem{
		TnlInformation: &e2ap_ies.Tnlinformation{
			TnlPort: &e2ap_commondatatypes.BitString{
				Value: 0x89bcd,
				Len:   16,
			},
			TnlAddress: &e2ap_commondatatypes.BitString{
				Value: 0x89abcdef01234567,
				Len:   64,
			},
		},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH,
	}

	if err := e2connectionUpdateItem.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionUpdateItem %s", err.Error())
	}
	return &e2connectionUpdateItem, nil
}

func Test_xerEncodingE2connectionUpdateItem(t *testing.T) {

	e2connectionUpdateItem, err := createE2connectionUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateItem PDU")

	xer, err := xerEncodeE2connectionUpdateItem(e2connectionUpdateItem)
	assert.NilError(t, err)
	assert.Equal(t, 315, len(xer))
	t.Logf("E2connectionUpdateItem XER\n%s", string(xer))

	result, err := xerDecodeE2connectionUpdateItem(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateItem XER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionUpdateItem.GetTnlUsage(), result.GetTnlUsage())
}

func Test_perEncodingE2connectionUpdateItem(t *testing.T) {

	e2connectionUpdateItem, err := createE2connectionUpdateItemMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdateItem PDU")

	per, err := perEncodeE2connectionUpdateItem(e2connectionUpdateItem)
	assert.NilError(t, err)
	assert.Equal(t, 13, len(per))
	t.Logf("E2connectionUpdateItem PER\n%v", hex.Dump(per))

	result, err := perDecodeE2connectionUpdateItem(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2connectionUpdateItem PER - decoded\n%v", result)
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlPort().GetLen(), result.GetTnlInformation().GetTnlPort().GetLen())
	assert.Equal(t, e2connectionUpdateItem.GetTnlInformation().GetTnlAddress().GetLen(), result.GetTnlInformation().GetTnlAddress().GetLen())
	assert.Equal(t, e2connectionUpdateItem.GetTnlUsage(), result.GetTnlUsage())
}
