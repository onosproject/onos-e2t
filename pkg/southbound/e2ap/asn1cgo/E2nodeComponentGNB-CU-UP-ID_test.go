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

func createE2nodeComponentGnbCuUpIDMsg() (*e2ap_ies.E2NodeComponentGnbCuUpId, error) {

	// e2nodeComponentGnbCuUpID := pdubuilder.CreateE2nodeComponentGnbCuUpID() //ToDo - fill in arguments here(if this function exists

	e2nodeComponentGnbCuUpID := e2ap_ies.E2NodeComponentGnbCuUpId{
		GNbCuUpId: &e2ap_ies.GnbCuUpId{
			Value: 2,
		},
	}

	if err := e2nodeComponentGnbCuUpID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentGnbCuUpId %s", err.Error())
	}
	return &e2nodeComponentGnbCuUpID, nil
}

func Test_xerEncodingE2nodeComponentGnbCuUpID(t *testing.T) {

	e2nodeComponentGnbCuUpID, err := createE2nodeComponentGnbCuUpIDMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentGnbCuUpId PDU")

	xer, err := xerEncodeE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID)
	assert.NilError(t, err)
	assert.Equal(t, 96, len(xer))
	t.Logf("E2nodeComponentGnbCuUpID XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentGnbCuUpID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentGnbCuUpID XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentGnbCuUpID.GetGNbCuUpId().GetValue(), result.GetGNbCuUpId().GetValue())
}

func Test_perEncodingE2nodeComponentGnbCuUpID(t *testing.T) {

	e2nodeComponentGnbCuUpID, err := createE2nodeComponentGnbCuUpIDMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentGnbCuUpId PDU")

	per, err := perEncodeE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("E2nodeComponentGnbCuUpID PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentGnbCuUpID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentGnbCuUpID PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentGnbCuUpID.GetGNbCuUpId().GetValue(), result.GetGNbCuUpId().GetValue())
}
