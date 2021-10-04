// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func createGlobalNgRanNodeIDGNb() *e2apies.GlobalNgRannodeId {

	return &e2apies.GlobalNgRannodeId{
		GlobalNgRannodeId: &e2apies.GlobalNgRannodeId_GNb{
			GNb: &e2apies.GlobalgNbId{
				PlmnId: &e2ap_commondatatypes.PlmnIdentity{
					Value: []byte{0x01, 0x02, 0x03},
				},
				GnbId: &e2apies.GnbIdChoice{
					GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
						GnbId: &asn1.BitString{
							Value: []byte{0xd4, 0xbc, 0x9c},
							Len:   22,
						},
					},
				},
			},
		},
	}
}

func createGlobalNgRanNodeIDNgENb() *e2apies.GlobalNgRannodeId {

	return &e2apies.GlobalNgRannodeId{
		GlobalNgRannodeId: &e2apies.GlobalNgRannodeId_NgENb{
			NgENb: &e2apies.GlobalngeNbId{
				PlmnId: &e2ap_commondatatypes.PlmnIdentity{
					Value: []byte{0x01, 0x02, 0x03},
				},
				EnbId: &e2apies.EnbIdChoice{
					EnbIdChoice: &e2apies.EnbIdChoice_EnbIdLongmacro{
						EnbIdLongmacro: &asn1.BitString{
							Value: []byte{0xd4, 0xbc, 0x98},
							Len:   21,
						},
					},
				},
			},
		},
	}
}

func Test_xerDecodeGlobalNgRanNodeID(t *testing.T) {

	ge2n := createGlobalNgRanNodeIDGNb()

	xer, err := xerEncodeGlobalNgRanNodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalNgRanNodeID (GNb) XER\n%s", xer)

	// Now reverse the XER
	ge2nReversed, err := xerDecodeGlobalNgRanNodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalNgRanNodeID (GNb) decoded from XER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetGNb().GetGnbId().GetGnbId().GetLen(), ge2nReversed.GetGNb().GetGnbId().GetGnbId().GetLen())
	assert.DeepEqual(t, ge2n.GetGNb().GetGnbId().GetGnbId().GetValue(), ge2nReversed.GetGNb().GetGnbId().GetGnbId().GetValue())
	assert.DeepEqual(t, ge2n.GetGNb().GetPlmnId().GetValue(), ge2nReversed.GetGNb().GetPlmnId().GetValue())

	ge2n = createGlobalNgRanNodeIDNgENb()

	xer, err = xerEncodeGlobalNgRanNodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalNgRanNodeID (ng-ENb) XER\n%s", xer)

	// Now reverse the XER
	ge2nReversed, err = xerDecodeGlobalNgRanNodeID(xer)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalNgRanNodeID (ng-ENb) decoded from XER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetLen(), ge2nReversed.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetLen())
	assert.DeepEqual(t, ge2n.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetValue(), ge2nReversed.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetValue())
	assert.DeepEqual(t, ge2n.GetNgENb().GetPlmnId().GetValue(), ge2nReversed.GetNgENb().GetPlmnId().GetValue())
}

func Test_perDecodeGlobalNgRanNodeID(t *testing.T) {

	ge2n := createGlobalNgRanNodeIDGNb()

	per, err := perEncodeGlobalNgRanNodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalNgRanNodeID (GNb) PER\n%v", hex.Dump(per))

	ge2nReversed, err := perDecodeGlobalNgRanNodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalNgRanNodeID (GNb) decoded from PER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetGNb().GetGnbId().GetGnbId().GetLen(), ge2nReversed.GetGNb().GetGnbId().GetGnbId().GetLen())
	assert.DeepEqual(t, ge2n.GetGNb().GetGnbId().GetGnbId().GetValue(), ge2nReversed.GetGNb().GetGnbId().GetGnbId().GetValue())
	assert.DeepEqual(t, ge2n.GetGNb().GetPlmnId().GetValue(), ge2nReversed.GetGNb().GetPlmnId().GetValue())

	ge2n = createGlobalNgRanNodeIDNgENb()

	per, err = perEncodeGlobalNgRanNodeID(ge2n)
	assert.NilError(t, err)
	t.Logf("GlobalNgRanNodeID (ng-ENb) XER\n%v", hex.Dump(per))

	ge2nReversed, err = perDecodeGlobalNgRanNodeID(per)
	assert.NilError(t, err)
	assert.Assert(t, ge2nReversed != nil)
	t.Logf("GlobalNgRanNodeID (ng-ENb) decoded from PER is \n%v", ge2nReversed)
	assert.Equal(t, ge2n.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetLen(), ge2nReversed.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetLen())
	assert.DeepEqual(t, ge2n.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetValue(), ge2nReversed.GetNgENb().GetEnbId().GetEnbIdLongmacro().GetValue())
	assert.DeepEqual(t, ge2n.GetNgENb().GetPlmnId().GetValue(), ge2nReversed.GetNgENb().GetPlmnId().GetValue())
}
