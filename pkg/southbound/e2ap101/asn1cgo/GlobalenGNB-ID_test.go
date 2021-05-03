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

func createGlobalenGnbIDMsg() (*e2ap_ies.GlobalenGnbId, error) {

	globalenGnbID := e2ap_ies.GlobalenGnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: []byte{0x01, 0x02, 0x03},
		},
		GNbId: &e2ap_ies.EngnbId{
			EngnbId: &e2ap_ies.EngnbId_GNbId{
				GNbId: &e2ap_commondatatypes.BitString{
					Value: 0x98bcd,
					Len:   32, //Should be of length 22 to 32
				},
			},
		},
	}

	if err := globalenGnbID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating GlobalenGnbId %s", err.Error())
	}
	return &globalenGnbID, nil
}

func Test_xerEncodingGlobalenGnbID(t *testing.T) {

	globalenGnbID, err := createGlobalenGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalenGnbId PDU")

	xer, err := xerEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	assert.Equal(t, 186, len(xer))
	t.Logf("GlobalenGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalenGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalenGnbID XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.DeepEqual(t, globalenGnbID.GetPLmnIdentity().GetValue(), result.GetPLmnIdentity().GetValue())
	assert.Equal(t, globalenGnbID.GetGNbId().GetGNbId().GetLen(), result.GetGNbId().GetGNbId().GetLen())

}

func Test_perEncodingGlobalenGnbID(t *testing.T) {

	globalenGnbID, err := createGlobalenGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalenGnbId PDU")

	per, err := perEncodeGlobalenGnbID(globalenGnbID)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("GlobalenGnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalenGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalenGnbID PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	assert.DeepEqual(t, globalenGnbID.GetPLmnIdentity().GetValue(), result.GetPLmnIdentity().GetValue())
	assert.Equal(t, globalenGnbID.GetGNbId().GetGNbId().GetLen(), result.GetGNbId().GetGNbId().GetLen())

}
