// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/encoder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
)

func TestE2connectionUpdate(t *testing.T) {
	//e2apPdu, err := pdubuilder.CreateE2connectionUpdateE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage().
	//	SetE2ConnectionUpdateAdd([]*types1.E2ConnectionUpdateItem{{TnlInformation: types1.TnlInformation{
	//		TnlPort: &asn1.BitString{
	//			Value: []byte{0xae, 0x89},
	//			Len:   16,
	//		},
	//		TnlAddress: asn1.BitString{
	//			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
	//			Len:   64,
	//		}},
	//		TnlUsage: e2apies.Tnlusage_TNLUSAGE_BOTH}}).SetE2ConnectionUpdateModify([]*types1.E2ConnectionUpdateItem{{TnlInformation: types1.TnlInformation{
	//	TnlPort: &asn1.BitString{
	//		Value: []byte{0xba, 0x91},
	//		Len:   16,
	//	},
	//	TnlAddress: asn1.BitString{
	//		Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62},
	//		Len:   64,
	//	}},
	//	TnlUsage: e2apies.Tnlusage_TNLUSAGE_RIC_SERVICE}}).SetE2ConnectionUpdateRemove([]*types1.TnlInformation{
	//	{TnlPort: &asn1.BitString{
	//		Value: []byte{0xba, 0x98},
	//		Len:   16,
	//	},
	//		TnlAddress: asn1.BitString{
	//			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76},
	//			Len:   64,
	//		}},
	//	{TnlPort: &asn1.BitString{
	//		Value: []byte{0xdc, 0x98},
	//		Len:   16,
	//	},
	//		TnlAddress: asn1.BitString{
	//			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x78},
	//			Len:   64,
	//		}},
	//})
	//t.Logf("That's what we are going to encode\n%v", e2apPdu)
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2connectionUpdate E2AP PDU PER\n%v", hex.Dump(per))

	newE2apPdu, err := CreateE2connectionUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	newE2apPdu.GetInitiatingMessage().GetValue().GetE2ConnectionUpdate().
		SetE2ConnectionUpdateAdd([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xae, 0x89},
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_BOTH}}).SetE2ConnectionUpdateModify([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
		TnlPort: &asn1.BitString{
			Value: []byte{0xba, 0x91},
			Len:   16,
		},
		TnlAddress: asn1.BitString{
			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62},
			Len:   64,
		}},
		TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE}}).SetE2ConnectionUpdateRemove([]*types.TnlInformation{
		{TnlPort: &asn1.BitString{
			Value: []byte{0xba, 0x98},
			Len:   16,
		},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76},
				Len:   64,
			}},
		{TnlPort: &asn1.BitString{
			Value: []byte{0xdc, 0x98},
			Len:   16,
		},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x78},
				Len:   64,
			}},
	})
	t.Logf("That's what we are goint to encode\n%v", newE2apPdu)

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.Assert(t, result1 != nil)
	//t.Logf("E2connectionUpdate E2AP PDU PER - decoded is \n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

func TestE2connectionUpdateExcludeOptionalIEs(t *testing.T) {
	//e2apPdu, err := pdubuilder.CreateE2connectionUpdateE2apPdu(1)
	//assert.NilError(t, err)
	//assert.Assert(t, e2apPdu != nil)
	//
	//e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage().
	//	SetE2ConnectionUpdateModify([]*types1.E2ConnectionUpdateItem{{TnlInformation: types1.TnlInformation{
	//		TnlPort: &asn1.BitString{
	//			Value: []byte{0xba, 0x19},
	//			Len:   16,
	//		},
	//		TnlAddress: asn1.BitString{
	//			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
	//			Len:   64,
	//		}},
	//		TnlUsage: e2apies.Tnlusage_TNLUSAGE_RIC_SERVICE}}).SetE2ConnectionUpdateRemove([]*types1.TnlInformation{
	//	{TnlPort: &asn1.BitString{
	//		Value: []byte{0xba, 0x98},
	//		Len:   16,
	//	},
	//		TnlAddress: asn1.BitString{
	//			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76},
	//			Len:   64,
	//		}},
	//	{TnlPort: &asn1.BitString{
	//		Value: []byte{0xdc, 0x98},
	//		Len:   16,
	//	},
	//		TnlAddress: asn1.BitString{
	//			Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62},
	//			Len:   64,
	//		}},
	//})
	//
	//per, err := asn1cgo.PerEncodeE2apPdu(e2apPdu)
	//assert.NilError(t, err)
	//t.Logf("E2connectionUpdate E2AP PDU PER\n%v", hex.Dump(per))

	newE2apPdu, err := CreateE2connectionUpdateE2apPdu(1)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	newE2apPdu.GetInitiatingMessage().GetValue().GetE2ConnectionUpdate().
		SetE2ConnectionUpdateModify([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: []byte{0xba, 0x19},
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x67},
				Len:   64,
			}},
			TnlUsage: e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE}}).SetE2ConnectionUpdateRemove([]*types.TnlInformation{
		{TnlPort: &asn1.BitString{
			Value: []byte{0xba, 0x98},
			Len:   16,
		},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x76},
				Len:   64,
			}},
		{TnlPort: &asn1.BitString{
			Value: []byte{0xdc, 0x98},
			Len:   16,
		},
			TnlAddress: asn1.BitString{
				Value: []byte{0x89, 0xab, 0xdc, 0xdf, 0x01, 0x23, 0x45, 0x62},
				Len:   64,
			}},
	})

	perNew, err := encoder.PerEncodeE2ApPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2connectionUpdate E2AP PDU PER with Go APER library\n%v", hex.Dump(perNew))

	//Comparing reference PER bytes with Go APER library produced
	//assert.DeepEqual(t, per, perNew)

	result, err := encoder.PerDecodeE2ApPdu(perNew)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu.String(), result.String())

	// Decoding the message from the APER bytes produced by CGo
	//result11, err := encoder.PerDecodeE2ApPdu(per)
	//assert.NilError(t, err)
	//assert.DeepEqual(t, newE2apPdu.String(), result11.String())
	//
	//result1, err := asn1cgo.PerDecodeE2apPdu(perNew)
	//assert.NilError(t, err)
	//assert.Assert(t, result1 != nil)
	//t.Logf("E2connectionUpdate E2AP PDU PER - decoded is\n%v", result1)
	//assert.DeepEqual(t, e2apPdu.String(), result1.String())
}

var bytes = []byte{
	0x20, 0x0b, 0x00, 0x28, 0x00, 0x00, 0x02, 0x00, 0x31, 0x00, 0x02, 0x00, 0x03, 0x00, 0x27, 0x00,
	0x1b, 0x08, 0x00, 0x2b, 0x40, 0x09, 0x21, 0xf0, 0x0a, 0xf4, 0x00, 0xa1, 0x8e, 0x45, 0x40, 0x00,
	0x2b, 0x40, 0x09, 0x21, 0xf0, 0x0a, 0xf4, 0x00, 0xa1, 0x8e, 0x45, 0x40,
	//0x20, 0x0b, 0x00, 0x28, 0x00, 0x00, 0x02, 0x00, 0x31, 0x00, 0x02, 0x00, 0x03, 0x00, 0x27, 0x00,
	//0x1b, 0x08, 0x00, 0x2b, 0x40, 0x09, 0x21, 0xf0, 0x0a, 0xf4, 0x00, 0x18, 0x8e, 0x45, 0x40, 0x00,
	//0x2b, 0x40, 0x09, 0x21, 0xf0, 0x0a, 0xf4, 0x00, 0x18, 0x8e, 0x45, 0x40,
}

func Test1(t *testing.T) {
	// Decoding the message from the APER bytes produced by CGo
	go func() {
		result11, err := encoder.PerDecodeE2ApPdu(bytes)
		assert.NilError(t, err)
		t.Logf("Decoded message is\n%v", result11)
	}()
	go func() {
		result11, err := encoder.PerDecodeE2ApPdu(bytes)
		assert.NilError(t, err)
		t.Logf("Decoded message is\n%v", result11)
	}()
	go func() {
		result11, err := encoder.PerDecodeE2ApPdu(bytes)
		assert.NilError(t, err)
		t.Logf("Decoded message is\n%v", result11)
	}()
}
