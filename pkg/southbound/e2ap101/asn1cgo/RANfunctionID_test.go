// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_RanFunctionID(t *testing.T) {

	rfID := &e2ap_ies.RanfunctionId{
		Value: 1,
	}

	xer, err := xerEncodeRanFunctionID(rfID)
	assert.NilError(t, err)
	t.Logf("RanFunctionID XER\n%s", xer)

	per, err := perEncodeRanFunctionID(rfID)
	assert.NilError(t, err)
	t.Logf("RanFunctionID PER\n%v", hex.Dump(per))

	// Now reverse the XER
	rfIDReversed, err := xerDecodeRanFunctionID(xer)
	assert.NilError(t, err)
	assert.Assert(t, rfIDReversed != nil)
	t.Logf("RanFunctionID decoded from XER is \n%v", rfIDReversed)
	//assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	rfIDReversedFromPer, err := perDecodeRanFunctionID(per)
	assert.NilError(t, err)
	assert.Assert(t, rfIDReversedFromPer != nil)
	t.Logf("RanFunctionID decoded from PER is \n%v", rfIDReversedFromPer)
	//assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
