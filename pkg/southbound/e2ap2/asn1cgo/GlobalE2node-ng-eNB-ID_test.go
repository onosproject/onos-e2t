// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"

	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2ap_ies/pdubuilder"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createGlobalE2nodeNgEnbIDMsg() (*e2ap_ies.GlobalE2NodeNgEnbId, error) {

	// globalE2nodeNgEnbID := pdubuilder.CreateGlobalE2nodeNgEnbID() //ToDo - fill in arguments here(if this function exists

	globalE2nodeNgEnbID := e2ap_ies.GlobalE2NodeNgEnbId{
		GlobalNgENbId: &e2ap_ies.GlobalngeNbId{
			PlmnId: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			EnbId: &e2ap_ies.EnbIdChoice{
				EnbIdChoice: &e2ap_ies.EnbIdChoice_EnbIdMacro{
					EnbIdMacro: &e2ap_commondatatypes.BitString{
						Value: []byte{0x4d, 0xcb, 0xb0},
						Len:   20,
					},
				},
			},
		},
	}

	if err := globalE2nodeNgEnbID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating GlobalE2nodeNgEnbId %s", err.Error())
	}
	return &globalE2nodeNgEnbID, nil
}

func Test_xerEncodingGlobalE2nodeNgEnbID(t *testing.T) {

	globalE2nodeNgEnbID, err := createGlobalE2nodeNgEnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeNgEnbId PDU")

	xer, err := xerEncodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbID)
	assert.NilError(t, err)
	assert.Equal(t, 261, len(xer))
	t.Logf("GlobalE2nodeNgEnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalE2nodeNgEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeNgEnbID XER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetPlmnId().GetValue(), result.GetGlobalNgENbId().GetPlmnId().GetValue())
	assert.Equal(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue())
}

func Test_perEncodingGlobalE2nodeNgEnbID(t *testing.T) {

	globalE2nodeNgEnbID, err := createGlobalE2nodeNgEnbIDMsg()
	assert.NilError(t, err, "Error creating GlobalE2nodeNgEnbId PDU")

	per, err := perEncodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbID)
	assert.NilError(t, err)
	assert.Equal(t, 8, len(per))
	t.Logf("GlobalE2nodeNgEnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalE2nodeNgEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalE2nodeNgEnbID PER - decoded\n%v", result)
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetPlmnId().GetValue(), result.GetGlobalNgENbId().GetPlmnId().GetValue())
	assert.Equal(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetLen())
	assert.DeepEqual(t, globalE2nodeNgEnbID.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue(), result.GetGlobalNgENbId().GetEnbId().GetEnbIdMacro().GetValue())
}
