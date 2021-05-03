// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createEngnbIDMsg() (*e2ap_ies.EngnbId, error) {

	engnbID := e2ap_ies.EngnbId{
		EngnbId: &e2ap_ies.EngnbId_GNbId{
			GNbId: &e2ap_commondatatypes.BitString{
				Value: 0x98bcd,
				Len:   22, //Should be of length 22 to 32
			},
		},
	}

	if err := engnbID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating EngnbId %s", err.Error())
	}
	return &engnbID, nil
}

func Test_xerEncodingEngnbID(t *testing.T) {

	engnbID, err := createEngnbIDMsg()
	assert.NilError(t, err, "Error creating EngnbId PDU")

	xer, err := xerEncodeEngnbID(engnbID)
	assert.NilError(t, err)
	assert.Equal(t, 81, len(xer))
	t.Logf("EngnbID XER\n%s", string(xer))

	result, err := xerDecodeEngnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EngnbID XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.Equal(t, engnbID.GetGNbId().GetLen(), result.GetGNbId().GetLen())
}

func Test_perEncodingEngnbID(t *testing.T) {

	engnbID, err := createEngnbIDMsg()
	assert.NilError(t, err, "Error creating EngnbId PDU")

	per, err := perEncodeEngnbID(engnbID)
	assert.NilError(t, err)
	assert.Equal(t, 4, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("EngnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeEngnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("EngnbID PER - decoded\n%v", result)
	assert.Equal(t, engnbID.GetGNbId().GetLen(), result.GetGNbId().GetLen())
}
