// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentInterfaceW1Msg() (*e2ap_ies.E2NodeComponentInterfaceW1, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceW1{
		NgENbDuId: &e2ap_ies.NgenbDuId{
			Value: 33,
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validatiW1 E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingW1E2nodeComponentInterfaceW1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceW1Msg()
	assert.NilError(t, err, "Error creatiW1 E2nodeComponentInterfaceW1 PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceW1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceW1 XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceW1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceW1 XER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetNgENbDuId().GetValue(), result.GetNgENbDuId().GetValue())
}

func Test_perEncodingW1E2nodeComponentInterfaceW1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceW1Msg()
	assert.NilError(t, err, "Error creatiW1 E2nodeComponentInterfaceW1 PDU")

	per, err := perEncodeE2nodeComponentInterfaceW1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceW1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceW1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceW1 PER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetNgENbDuId().GetValue(), result.GetNgENbDuId().GetValue())
}
