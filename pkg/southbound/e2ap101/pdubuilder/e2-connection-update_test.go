// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"gotest.tools/assert"
	"testing"
)

func TestE2connectionUpdate(t *testing.T) {
	newE2apPdu, err := CreateE2connectionUpdateE2apPdu([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
		TnlPort: e2ap_commondatatypes.BitString{
			Value: 0x89ae,
			Len:   16,
		},
		TnlAddress: e2ap_commondatatypes.BitString{
			Value: 0x89abdcdf01234567,
			Len:   64,
		}},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}},
		[]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: e2ap_commondatatypes.BitString{
				Value: 0x91ab,
				Len:   16,
			},
			TnlAddress: e2ap_commondatatypes.BitString{
				Value: 0x65abcdef01234567,
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE}},
		[]*types.TnlInformation{
			{TnlPort: e2ap_commondatatypes.BitString{
				Value: 0x89ab,
				Len:   16,
			},
				TnlAddress: e2ap_commondatatypes.BitString{
					Value: 0x89abcdef01234567,
					Len:   64,
				}},
			{TnlPort: e2ap_commondatatypes.BitString{
				Value: 0x89cd,
				Len:   16,
			},
				TnlAddress: e2ap_commondatatypes.BitString{
					Value: 0x89abcdef12345678,
					Len:   64,
				}},
		})
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdate E2AP PDU XER\n%s", string(xer))

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdate E2AP PDU PER\n%v", hex.Dump(per))
}
