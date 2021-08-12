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

func createE2nodeComponentGnbDuIDMsg() (*e2ap_ies.E2NodeComponentGnbDuId, error) {

	// e2nodeComponentGnbDuID := pdubuilder.CreateE2nodeComponentGnbDuID() //ToDo - fill in arguments here(if this function exists

	e2nodeComponentGnbDuID := e2ap_ies.E2NodeComponentGnbDuId{
		GNbDuId: &e2ap_ies.GnbDuId{
			Value: 11,
		},
	}

	if err := e2nodeComponentGnbDuID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentGnbDuId %s", err.Error())
	}
	return &e2nodeComponentGnbDuID, nil
}

func Test_xerEncodingE2nodeComponentGnbDuID(t *testing.T) {

	e2nodeComponentGnbDuID, err := createE2nodeComponentGnbDuIDMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentGnbDuId PDU")

	xer, err := xerEncodeE2nodeComponentGnbDuID(e2nodeComponentGnbDuID)
	assert.NilError(t, err)
	assert.Equal(t, 85, len(xer))
	t.Logf("E2nodeComponentGnbDuID XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentGnbDuID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentGnbDuID XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentGnbDuID.GetGNbDuId().GetValue(), result.GetGNbDuId().GetValue())
}

func Test_perEncodingE2nodeComponentGnbDuID(t *testing.T) {

	e2nodeComponentGnbDuID, err := createE2nodeComponentGnbDuIDMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentGnbDuId PDU")

	per, err := perEncodeE2nodeComponentGnbDuID(e2nodeComponentGnbDuID)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2nodeComponentGnbDuID PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentGnbDuID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentGnbDuID PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentGnbDuID.GetGNbDuId().GetValue(), result.GetGNbDuId().GetValue())
}
