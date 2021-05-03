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

func createE2nodeComponentConfigUpdateGnbMsg() (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {

	e2nodeComponentConfigUpdateGnb := e2ap_ies.E2NodeComponentConfigUpdateGnb{
		NgApconfigUpdate: "ng_AP",
		XnApconfigUpdate: "xn_AP",
		E1ApconfigUpdate: "e1_AP",
		F1ApconfigUpdate: "f1_AP",
	}

	if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateGnb %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateGnb, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateGnb(t *testing.T) {

	e2nodeComponentConfigUpdateGnb, err := createE2nodeComponentConfigUpdateGnbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateGnb PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb)
	assert.NilError(t, err)
	assert.Equal(t, 291, len(xer))
	t.Logf("E2nodeComponentConfigUpdateGnb XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateGnb(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateGnb XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetNgApconfigUpdate(), result.GetNgApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetXnApconfigUpdate(), result.GetXnApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetE1ApconfigUpdate(), result.GetE1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetF1ApconfigUpdate(), result.GetF1ApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdateGnb(t *testing.T) {

	e2nodeComponentConfigUpdateGnb, err := createE2nodeComponentConfigUpdateGnbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateGnb PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb)
	assert.NilError(t, err)
	assert.Equal(t, 25, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2nodeComponentConfigUpdateGnb PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateGnb(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateGnb PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetNgApconfigUpdate(), result.GetNgApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetXnApconfigUpdate(), result.GetXnApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetE1ApconfigUpdate(), result.GetE1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateGnb.GetF1ApconfigUpdate(), result.GetF1ApconfigUpdate())
}
