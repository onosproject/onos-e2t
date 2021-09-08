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

func createE2nodeComponentConfigUpdateEnbMsg() (*e2ap_ies.E2NodeComponentConfigUpdateEnb, error) {

	e2nodeComponentConfigUpdateEnb := e2ap_ies.E2NodeComponentConfigUpdateEnb{
		//S1ApconfigUpdate: "s1_AP",
		//X2ApconfigUpdate: "x2_AP",
	}

	if err := e2nodeComponentConfigUpdateEnb.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateEnb %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateEnb, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateEnb(t *testing.T) {

	e2nodeComponentConfigUpdateEnb, err := createE2nodeComponentConfigUpdateEnbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateEnb PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb)
	assert.NilError(t, err)
	//assert.Equal(t, 179, len(xer))
	t.Logf("E2nodeComponentConfigUpdateEnb XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateEnb(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateEnb XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateEnb.GetS1ApconfigUpdate(), result.GetS1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateEnb.GetX2ApconfigUpdate(), result.GetX2ApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdateEnb(t *testing.T) {

	e2nodeComponentConfigUpdateEnb, err := createE2nodeComponentConfigUpdateEnbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateEnb PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb)
	assert.NilError(t, err)
	//assert.Equal(t, 13, len(per))
	t.Logf("E2nodeComponentConfigUpdateEnb PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateEnb(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateEnb PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateEnb.GetS1ApconfigUpdate(), result.GetS1ApconfigUpdate())
	assert.Equal(t, e2nodeComponentConfigUpdateEnb.GetX2ApconfigUpdate(), result.GetX2ApconfigUpdate())
}
