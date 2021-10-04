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

func createE2nodeComponentInterfaceE1Msg() (*e2ap_ies.E2NodeComponentInterfaceE1, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceE1{
		GNbCuCpId: &e2ap_ies.GnbCuUpId{
			Value: 96,
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingE2nodeComponentInterfaceE1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceE1Msg()
	assert.NilError(t, err, "Error creating E2nodeComponentInterfaceE1 PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceE1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceE1 XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceE1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceE1 XER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetGNbCuCpId().GetValue(), result.GetGNbCuCpId().GetValue())
}

func Test_perEncodingE2nodeComponentInterfaceE1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceE1Msg()
	assert.NilError(t, err, "Error creating E2nodeComponentInterfaceE1 PDU")

	per, err := perEncodeE2nodeComponentInterfaceE1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceE1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceE1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceE1 PER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetGNbCuCpId().GetValue(), result.GetGNbCuCpId().GetValue())
}
