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

func TestNewGlobaleNBID(t *testing.T) {

	g := e2apies.GlobalEnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: []byte("ONF"),
		},
		ENbId: &e2apies.EnbId{
			EnbId: &e2apies.EnbId_HomeENbId{
				HomeENbId: &e2ap_commondatatypes.BitString{
					Value: []byte{0xf0, 0xde, 0xcb, 0xb0},
					Len:   28,
				},
			},
		},
	}

	cobject, err := newGlobaleNBID(&g)
	assert.NilError(t, err, "error converting to c struct")
	//assert.Assert(t, cobject != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 33 & 34
	assert.Equal(t, int(cobject.pLMN_Identity.size), 3, "expected plmn id to be 3 bytes")
	assert.Equal(t, int(cobject.eNB_ID.present), 2, "expected choice to be 1 (home_eNB_ID)")

	// Now do the reverse - C object back to struct
	g1, err := decodeGlobalEnbID(cobject)
	assert.NilError(t, err, "error converting back from c struct")
	//assert.Assert(t, g1 != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 40
	assert.Equal(t, string(g1.PLmnIdentity.Value), "ONF", "unexpected value for Plmn ID")
	switch choice := g1.ENbId.EnbId.(type) {
	case *e2apies.EnbId_HomeENbId:
		assert.Equal(t, int(choice.HomeENbId.Len), 28)
		assert.DeepEqual(t, choice.HomeENbId.Value, []byte{0xf0, 0xde, 0xcb, 0xb0})
	default:
		t.Fatalf("unexpected choice in EnbID %v", choice)
	}

	xer, err := xerEncodeeNBID(&g)
	assert.NilError(t, err)
	t.Logf("XER GlobalEnbId: \n%s", string(xer))

	per, err := perEncodeeNBID(&g)
	assert.NilError(t, err)
	t.Logf("PER GlobalEnbId: \n%v", hex.Dump(per))
}
