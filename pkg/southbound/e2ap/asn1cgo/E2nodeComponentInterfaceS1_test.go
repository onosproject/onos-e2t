// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentInterfaceS1Msg() (*e2ap_ies.E2NodeComponentInterfaceS1, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceS1{
		MmeName: &e2ap_commondatatypes.Mmename{
			Value: "ONF",
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validatiS1 E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodiS1E2nodeComponentInterfaceS1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceS1Msg()
	assert.NilError(t, err, "Error creatiS1 E2nodeComponentInterfaceS1 PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceS1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceS1 XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceS1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceS1 XER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetMmeName().GetValue(), result.GetMmeName().GetValue())
}

func Test_perEncodiS1E2nodeComponentInterfaceS1(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceS1Msg()
	assert.NilError(t, err, "Error creatiS1 E2nodeComponentInterfaceS1 PDU")

	per, err := perEncodeE2nodeComponentInterfaceS1(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceS1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceS1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceS1 PER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetMmeName().GetValue(), result.GetMmeName().GetValue())
}
