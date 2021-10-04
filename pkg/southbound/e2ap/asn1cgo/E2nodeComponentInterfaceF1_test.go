// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentInterfaceF1Msg() (*e2ap_ies.E2NodeComponentInterfaceF1, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceF1{
		GNbDuId: &e2ap_ies.GnbDuId{
			Value: 96,
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingE2nodeComponentInterfaceF1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceF1Msg()
	assert.NilError(t, err, "Error creating E2nodeComponentInterfaceF1 PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceF1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceF1 XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceF1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceF1 XER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetGNbDuId().GetValue(), result.GetGNbDuId().GetValue())
}

func Test_perEncodingE2nodeComponentInterfaceF1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceF1Msg()
	assert.NilError(t, err, "Error creating E2nodeComponentInterfaceF1 PDU")

	per, err := perEncodeE2nodeComponentInterfaceF1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceF1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceF1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceF1 PER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetGNbDuId().GetValue(), result.GetGNbDuId().GetValue())
}
