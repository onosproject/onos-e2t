// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2ap_ies/pdubuilder"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentConfigUpdateEngNbMsg() (*e2ap_ies.E2NodeComponentConfigUpdateEngNb, error) {

	// e2nodeComponentConfigUpdateEngNb := pdubuilder.CreateE2nodeComponentConfigUpdateEngNb() //ToDo - fill in arguments here(if this function exists

	e2nodeComponentConfigUpdateEngNb := e2ap_ies.E2NodeComponentConfigUpdateEngNb{
		X2ApconfigUpdate: "x2_AP",
	}

	if err := e2nodeComponentConfigUpdateEngNb.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateEngNb %s", err.Error())
	}
	return &e2nodeComponentConfigUpdateEngNb, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdateEngNb(t *testing.T) {

	e2nodeComponentConfigUpdateEngNb, err := createE2nodeComponentConfigUpdateEngNbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateEngNb PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb)
	assert.NilError(t, err)
	assert.Equal(t, 127, len(xer))
	t.Logf("E2nodeComponentConfigUpdateEngNb XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdateEngNb(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateEngNb XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateEngNb.GetX2ApconfigUpdate(), result.GetX2ApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdateEngNb(t *testing.T) {

	e2nodeComponentConfigUpdateEngNb, err := createE2nodeComponentConfigUpdateEngNbMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdateEngNb PDU")

	per, err := perEncodeE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb)
	assert.NilError(t, err)
	assert.Equal(t, 7, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2nodeComponentConfigUpdateEngNb PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdateEngNb(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdateEngNb PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentConfigUpdateEngNb.GetX2ApconfigUpdate(), result.GetX2ApconfigUpdate())
}
