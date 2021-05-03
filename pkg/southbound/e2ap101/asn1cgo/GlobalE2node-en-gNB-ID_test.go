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

func createGlobalE2nodeEnGnbIDMsg() (*e2ap_ies.GlobalE2NodeEnGnbId, error) {

	globalE2nodeEnGnbID := e2ap_ies.GlobalE2NodeEnGnbId{
		GlobalGNbId: &e2ap_ies.GlobalenGnbId{
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
		},
	}

	if err := globalE2nodeEnGnbID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating GlobalE2nodeEnGnbId %s", err.Error())
	}
	return &globalE2nodeEnGnbID, nil
}

func Test_xerEncodingGlobalE2nodeEnGnbID(t *testing.T) {

	globalE2nodeEnGnbID, err := createGlobalE2nodeEnGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeEnGnbId PDU")

	xer, err := xerEncodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbID)
	assert.NilError(t, err)
	assert.Equal(t, 267, len(xer))
	t.Logf("GlobalE2nodeEnGnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalE2nodeEnGnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeEnGnbID XER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeEnGnbID.GetGlobalGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalGNbId().GetPLmnIdentity().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetGlobalGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetGlobalGNbId().GetGNbId().GetGNbId().GetLen(), result.GetGlobalGNbId().GetGNbId().GetGNbId().GetLen())
}

func Test_perEncodingGlobalE2nodeEnGnbID(t *testing.T) {

	globalE2nodeEnGnbID, err := createGlobalE2nodeEnGnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeEnGnbId PDU")

	per, err := perEncodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbID)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per))
	t.Logf("GlobalE2nodeEnGnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalE2nodeEnGnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeEnGnbID PER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeEnGnbID.GetGlobalGNbId().GetPLmnIdentity().GetValue(), result.GetGlobalGNbId().GetPLmnIdentity().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetGlobalGNbId().GetGNbId().GetGNbId().GetValue(), result.GetGlobalGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, globalE2nodeEnGnbID.GetGlobalGNbId().GetGNbId().GetGNbId().GetLen(), result.GetGlobalGNbId().GetGNbId().GetGNbId().GetLen())
}
