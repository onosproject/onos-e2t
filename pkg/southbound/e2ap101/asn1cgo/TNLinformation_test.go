// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createTnlinformationMsg() (*e2ap_ies.Tnlinformation, error) {

	bs1 := &e2ap_commondatatypes.BitString{
		Value: 0x89abcdef01234567,
		Len:   64,
	}

	bs2 := &e2ap_commondatatypes.BitString{
		Value: 0x89bc,
		Len:   16,
	}

	tnlinformation := e2ap_ies.Tnlinformation{
		TnlAddress: bs1,
		TnlPort:    bs2,
	}

	if err := tnlinformation.Validate(); err != nil {
		return nil, fmt.Errorf("error validating Tnlinformation %s", err.Error())
	}
	return &tnlinformation, nil
}

func Test_xerEncodingTnlinformation(t *testing.T) {

	tnlinformation, err := createTnlinformationMsg()
	assert.NilError(t, err, "Error creating TNLinformation PDU")
	t.Logf("TNLinformation (message)\n%v", tnlinformation)

	xer, err := xerEncodeTnlinformation(tnlinformation)
	assert.NilError(t, err)
	assert.Equal(t, 197, len(xer))
	t.Logf("TNLinformation XER\n%s", string(xer))

	result, err := xerDecodeTnlinformation(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLinformation XER - decoded\n%v", result)

	assert.Equal(t, tnlinformation.GetTnlAddress().GetValue(), result.GetTnlAddress().GetValue())
	assert.Equal(t, tnlinformation.GetTnlAddress().GetLen(), result.GetTnlAddress().GetLen())
	assert.Equal(t, tnlinformation.GetTnlPort().GetValue(), result.GetTnlPort().GetValue())
	assert.Equal(t, tnlinformation.GetTnlPort().GetLen(), result.GetTnlPort().GetLen())
}

func Test_perEncodingTnlinformation(t *testing.T) {

	tnlinformation, err := createTnlinformationMsg()
	assert.NilError(t, err, "Error creating TNLinformation PDU")
	t.Logf("TNLinformation (message)\n%v", tnlinformation)

	per, err := perEncodeTnlinformation(tnlinformation)
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("TNLinformation PER\n%v", hex.Dump(per))

	result, err := perDecodeTnlinformation(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("TNLinformation PER - decoded\n%v", result)

	assert.Equal(t, tnlinformation.GetTnlAddress().GetValue(), result.GetTnlAddress().GetValue())
	assert.Equal(t, tnlinformation.GetTnlAddress().GetLen(), result.GetTnlAddress().GetLen())
	assert.Equal(t, tnlinformation.GetTnlPort().GetValue(), result.GetTnlPort().GetValue())
	assert.Equal(t, tnlinformation.GetTnlPort().GetLen(), result.GetTnlPort().GetLen())
}
