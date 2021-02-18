// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func TestNewGlobalgNBID(t *testing.T) {

	g := e2apies.GlobalgNbId{
		PlmnId: &e2ap_commondatatypes.PlmnIdentity{
			Value: []byte("ONF"),
		},
		GnbId: &e2apies.GnbIdChoice{
			GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
				GnbId: &e2ap_commondatatypes.BitString{
					Value: 0x9ABCDEF0,
					Len:   28,
				},
			},
		},
	}

	cobject, err := newGlobalgNBID(&g)
	assert.NilError(t, err, "error converting to c struct")
	//assert.Assert(t, cobject != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 33 & 34
	assert.Equal(t, int(cobject.plmn_id.size), 3, "expected plmn id to be 3 bytes")
	assert.Equal(t, int(cobject.gnb_id.present), 1, "expected choice to be 1 (gnb_id)")

	// Now do the reverse - C object back to struct
	g1, err := decodeGlobalGnbID(cobject)
	assert.NilError(t, err, "error converting back from c struct")
	//assert.Assert(t, g1 != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 40
	assert.Equal(t, string(g1.PlmnId.Value), "ONF", "unexpected value for Plmn ID")
	switch choice := g1.GnbId.GnbIdChoice.(type) {
	case *e2apies.GnbIdChoice_GnbId:
		assert.Equal(t, int(choice.GnbId.Len), 28)
		assert.Equal(t, choice.GnbId.Value, uint64(0x9ABCDEF0))
	default:
		t.Fatalf("unexpected choice in GnbIdChoice %v", choice)
	}

	xer, err := xerEncodegNBID(&g)
	assert.NilError(t, err)
	t.Logf("XER GlobalgNbId: \n%s", string(xer))

	per, err := perEncodegNBID(&g)
	assert.NilError(t, err)
	t.Logf("PER GlobalgNbId: \n%s", string(per))
}
