// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentConfigUpdateNgeNbMsg() (*e2ap_ies.E2NodeComponentConfigUpdateNgeNb, error) {

	e2nodeComponentConfigUpdateNgeNb := e2ap_ies.E2NodeComponentConfigUpdateNgeNb{
		NgApconfigUpdate: "ng_AP",
		XnApconfigUpdate: "xn_AP",
	}

	if err := e2nodeComponentConfigUpdateNgeNb.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateNgeNb %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateNgeNb, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateNgeNb(t *testing.T) {

	e2nodeComponentConfigUpdateNgeNb, err := createE2nodeComponentConfigUpdateNgeNbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateNgeNb PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb)
	assert.NilError(t, err)
	//assert.Equal(t, 183, len(xer))
	t.Logf("E2nodeComponentConfigUpdateNgeNb XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateNgeNb(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateNgeNb XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateNgeNb.GetNgApconfigUpdate(), result.GetNgApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateNgeNb.GetXnApconfigUpdate(), result.GetXnApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdateNgeNb(t *testing.T) {

	e2nodeComponentConfigUpdateNgeNb, err := createE2nodeComponentConfigUpdateNgeNbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateNgeNb PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb)
	assert.NilError(t, err)
	//assert.Equal(t, 13, len(per))
	t.Logf("E2nodeComponentConfigUpdateNgeNb PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateNgeNb(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateNgeNb PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateNgeNb.GetNgApconfigUpdate(), result.GetNgApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateNgeNb.GetXnApconfigUpdate(), result.GetXnApconfigUpdate())
}
