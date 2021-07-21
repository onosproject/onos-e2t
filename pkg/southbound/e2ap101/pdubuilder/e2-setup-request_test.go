// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

func TestE2SetupRequest(t *testing.T) {
	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         []byte("oid1"),
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         []byte("oid2"),
	}

	gnbID, err := CreateGnbIDchoice([]byte{0x00, 0x00, 0x04}, 22)
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu([3]byte{0x4F, 0x4E, 0x46}, gnbID, ranFunctionList)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)
}

func TestE2SetupRequestExcludeOptionalIE(t *testing.T) {
	gnbID, err := CreateGnbIDchoice([]byte{0x00, 0x00, 0x04}, 22)
	assert.NilError(t, err)

	newE2apPdu, err := CreateE2SetupRequestPdu([3]byte{0x4F, 0x4E, 0x46}, gnbID, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2apPdu != nil)

	xer, err := asn1cgo.XerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest XER\n%s", string(xer))

	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(xer)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)

	per, err := asn1cgo.PerEncodeE2apPdu(newE2apPdu)
	assert.NilError(t, err)
	t.Logf("E2SetupRequest E2AP PDU PER\n%v", hex.Dump(per))

	e2apPdu, err = asn1cgo.PerDecodeE2apPdu(per)
	assert.NilError(t, err)
	assert.DeepEqual(t, newE2apPdu, e2apPdu)
}

var adib1 = "0001008236000002000300090003f5115000000001000a00822100020008408180800001815b70184f52414e2d4532534d2d4b504d000018312e332e362e312e342e312e35333134382e312e322e322e3205004b504d206d6f6e69746f720000400003f51150000000010000000001300003f511000004001000010700506572696f646963207265706f72740001000309004532204e6f6465204d6561737572656d656e740000000742605252432e436f6e6e45737461624174742e73756d00000042805252432e436f6e6e4573746162537563632e73756d00000141a05252432e5265457374616241747400000244805252432e526545737461624174742e7265636f6e66696775726174696f6e4661696c75726500000343a05252432e526545737461624174742e68616e646f7665724661696c75726500000443405252432e526545737461624174742e6f746865724661696c75726500000541605252432e436f6e6e4d65616e00000641405252432e436f6e6e4d6178000007000100010000011c000018312e332e362e312e342e312e35333134382e312e322e322e3200084080958000026f00f04f52414e2d4532534d2d52432d50524500001a312e332e362e312e342e312e35333134382e312e322e322e313030028052432d505245600001068052432d5052452d74726967676572000100010c8050434920616e64204e52542075706461746520666f7220674e42000100010000011e00001a312e332e362e312e342e312e35333134382e312e322e322e313030"

func TestE2SetupRequestAdib1(t *testing.T) {

	perBytes, err := hexlib.Asn1BytesToByte(adib1)
	assert.NilError(t, err)
	t.Logf("These are the bytes obtained from PCAP\n%v", hex.Dump(perBytes))

	t.Logf("Attempting to decode E2SetupRequest message")
	e2sr, err := asn1cgo.PerDecodeE2apPdu(perBytes)
	assert.NilError(t, err)
	t.Logf("Here is a decoded message in protobuf\n%v", e2sr)

	xer, err := asn1cgo.XerEncodeE2apPdu(e2sr)
	assert.NilError(t, err)
	t.Logf("Here is a decoded message in XER\n%s", xer)
}

var adib2 = "0001008236000002000300090003f5115000000002000a00822100020008408180800001815b70184f52414e2d4532534d2d4b504d000018312e332e362e312e342e312e35333134382e312e322e322e3205004b504d206d6f6e69746f720000400003f51150000000020000000001300003f511000008002000010700506572696f646963207265706f72740001000309004532204e6f6465204d6561737572656d656e740000000742605252432e436f6e6e45737461624174742e73756d00000042805252432e436f6e6e4573746162537563632e73756d00000141a05252432e5265457374616241747400000244805252432e526545737461624174742e7265636f6e66696775726174696f6e4661696c75726500000343a05252432e526545737461624174742e68616e646f7665724661696c75726500000443405252432e526545737461624174742e6f746865724661696c75726500000541605252432e436f6e6e4d65616e00000641405252432e436f6e6e4d6178000007000100010000011c000018312e332e362e312e342e312e35333134382e312e322e322e3200084080958000026f00f04f52414e2d4532534d2d52432d50524500001a312e332e362e312e342e312e35333134382e312e322e322e313030028052432d505245600001068052432d5052452d74726967676572000100010c8050434920616e64204e52542075706461746520666f7220674e42000100010000011e00001a312e332e362e312e342e312e35333134382e312e322e322e313030"

func TestE2SetupRequestAdib2(t *testing.T) {

	perBytes, err := hexlib.Asn1BytesToByte(adib2)
	assert.NilError(t, err)
	t.Logf("These are the bytes obtained from PCAP\n%v", hex.Dump(perBytes))

	t.Logf("Attempting to decode E2SetupRequest message")
	e2sr, err := asn1cgo.PerDecodeE2apPdu(perBytes)
	assert.NilError(t, err)
	t.Logf("Here is a decoded message in protobuf\n%v", e2sr)

	xer, err := asn1cgo.XerEncodeE2apPdu(e2sr)
	assert.NilError(t, err)
	t.Logf("Here is a decoded message in XER\n%s", xer)
}
