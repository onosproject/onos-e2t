// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"gotest.tools/assert"
	"testing"
)

func Test_E2setupRequest(t *testing.T) {

	gnbIDIe := e2appducontents.E2SetupRequestIes_E2SetupRequestIes3{
		Value: &e2apies.GlobalE2NodeId{
			GlobalE2NodeId: &e2apies.GlobalE2NodeId_GNb{
				GNb: &e2apies.GlobalE2NodeGnbId{
					GlobalGNbId: &e2apies.GlobalgNbId{
						PlmnId: &e2ap_commondatatypes.PlmnIdentity{
							Value: []byte("ONF"),
						},
						GnbId: &e2apies.GnbIdChoice{
							GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
								GnbId: &e2ap_commondatatypes.BitString{
									Value: []byte{0xd4, 0xcb, 0x8c},
									Len:   22,
								}},
						},
					},
				},
			},
		},
	}

	e2sr := e2appducontents.E2SetupRequest{
		ProtocolIes: &e2appducontents.E2SetupRequestIes{
			E2ApProtocolIes3: &gnbIDIe,
			// TODO add in a RANfunctionList
		},
	}

	// Convert this Go struct in to a C struct
	e2srC, err := newE2SetupRequest(&e2sr)
	assert.NilError(t, err)
	assert.Assert(t, e2srC != nil)

	// Now reverse it and decode the other way round to a Go struct
	e2srFedback, err := decodeE2setupRequest(e2srC)
	assert.NilError(t, err)
	//assert.Assert(t, e2srFedback != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 54
	ge2nID := e2srFedback.ProtocolIes.E2ApProtocolIes3.Value.GlobalE2NodeId.(*e2apies.GlobalE2NodeId_GNb)
	assert.Equal(t, "ONF", string(ge2nID.GNb.GlobalGNbId.PlmnId.Value))
	gnbID := ge2nID.GNb.GlobalGNbId.GnbId.GnbIdChoice.(*e2apies.GnbIdChoice_GnbId)
	assert.DeepEqual(t, []byte{0xd4, 0xcb, 0x8c}, gnbID.GnbId.Value)
	assert.Equal(t, uint32(22), gnbID.GnbId.Len)

	xer, err := xerEncodeE2SetupRequest(e2srFedback)
	assert.NilError(t, err)
	t.Logf("XER E2SetupRequest: \n%s", string(xer))

	per, err := perEncodeE2SetupRequest(e2srFedback)
	assert.NilError(t, err)
	t.Logf("PER E2SetupRequest: \n%v", hex.Dump(per))
}
