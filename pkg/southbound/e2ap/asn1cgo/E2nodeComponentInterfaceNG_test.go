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

func createE2nodeComponentInterfaceNGMsg() (*e2ap_ies.E2NodeComponentInterfaceNg, error) {

	e2ncc := e2ap_ies.E2NodeComponentInterfaceNg{
		AmfName: &e2ap_commondatatypes.Amfname{
			Value: "ONF",
		},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingE2nodeComponentInterfaceNG(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceNGMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentInterfaceNG PDU")

	xer, err := xerEncodeE2nodeComponentInterfaceNG(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceNG XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentInterfaceNG(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceNG XER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetAmfName().GetValue(), result.GetAmfName().GetValue())
}

func Test_perEncodingE2nodeComponentInterfaceNG(t *testing.T) {

	e2ncc, err := createE2nodeComponentInterfaceNGMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentInterfaceNG PDU")

	per, err := perEncodeE2nodeComponentInterfaceNG(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentInterfaceNG PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentInterfaceNG(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentInterfaceNG PER - decoded\n%v", result)
	assert.Equal(t, e2ncc.GetAmfName().GetValue(), result.GetAmfName().GetValue())
}
