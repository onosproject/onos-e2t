// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func Test_newE2setupResponseE2APpdu(t *testing.T) {
	rfAccepted := make(types.RanFunctionRevisions)
	rfAccepted[100] = 2
	rfAccepted[200] = 2

	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0xDE, 0xBC, 0xA0},
		RicIdentifierLen:   20,
	}

	e2SetupResponseE2APpdu, err := pdubuilder.CreateResponseE2apPdu(plmnID, ricID, rfAccepted, nil)
	assert.NilError(t, err)

	e2SetupResponseE2APpduC, err := newE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	assert.Assert(t, e2SetupResponseE2APpduC != nil)

	xer, err := XerEncodeE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	t.Logf("XER of E2AP: %s\n", string(xer))

	per, err := PerEncodeE2apPdu(e2SetupResponseE2APpdu)
	assert.NilError(t, err)
	t.Logf("PER of E2AP: %v\n", hex.Dump(per))

}
