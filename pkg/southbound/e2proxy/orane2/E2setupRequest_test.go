// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
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
									Value: 0x9bcd4,
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
	assert.Assert(t, e2srFedback != nil)
	ge2nId := e2srFedback.ProtocolIes.E2ApProtocolIes3.Value.GlobalE2NodeId.(*e2apies.GlobalE2NodeId_GNb)
	assert.Equal(t, "ONF", string(ge2nId.GNb.GlobalGNbId.PlmnId.Value))
	gnbId := ge2nId.GNb.GlobalGNbId.GnbId.GnbIdChoice.(*e2apies.GnbIdChoice_GnbId)
	assert.Equal(t, uint64(0x9bcd4), gnbId.GnbId.Value)
	assert.Equal(t, uint32(22), gnbId.GnbId.Len)

	xer, err := xerEncodeE2SetupRequest(e2srFedback)
	assert.NilError(t, err)
	t.Logf("XER E2SetupRequest: \n%s", string(xer))

	per, err := perEncodeE2SetupRequest(e2srFedback)
	assert.NilError(t, err)
	t.Logf("PER E2SetupRequest: \n%s", string(per))
}
