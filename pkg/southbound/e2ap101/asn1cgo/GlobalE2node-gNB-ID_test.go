// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createGlobalE2nodeGnbID() *e2apies.GlobalE2NodeGnbId {

	return &e2apies.GlobalE2NodeGnbId{
		GlobalGNbId: &e2apies.GlobalgNbId{
			PlmnId: &e2ap_commondatatypes.PlmnIdentity{
				Value: []byte{0x01, 0x02, 0x03},
			},
			GnbId: &e2apies.GnbIdChoice{
				GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
					GnbId: &e2ap_commondatatypes.BitString{
						Value: 0x9ABCD4,
						Len:   22,
					},
				},
			},
		},
		//GNbCuUpId: &e2apies.GnbCuUpId{
		//	Value: 21,
		//},
		//GNbDuId:   &e2apies.GnbDuId{
		//	Value: 13,
		//},
	}
}

func Test_GlobalE2nodeGnbID(t *testing.T) {

	ge2n := createGlobalE2nodeGnbID()

	xer, err := xerEncodeGlobalE2nodegNBID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodegNBID XER\n%s", xer)

	per, err := perEncodeGlobalE2nodegNBID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalE2nodegNBID PER\n%s", hex.Dump(per))

	// Now reverse the XER
	ge2nReversed, err := xerDecodeGlobalE2nodegNBID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalE2nodegNBID decoded from XER is \n%v", ge2nReversed)
	//assert.Equal(t, 2, len(rflReversed.GetValue()))

	// Now reverse the PER
	ge2nReversedFromPer, err := perDecodeGlobalE2nodegNBID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversedFromPer != nil)
	t.Logf("GlobalE2nodegNBID decoded from PER is \n%v", ge2nReversedFromPer)
	//assert.Equal(t, 2, len(rflReversedFromPer.GetValue()))

}
