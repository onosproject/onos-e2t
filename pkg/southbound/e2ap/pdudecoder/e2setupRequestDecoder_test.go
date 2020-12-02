// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SetupRequestPdu(t *testing.T) {
	e2setupRequestXer, err := ioutil.ReadFile("../test/E2setupRequest-eNB.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
	assert.NilError(t, err)

	identifier, ranFunctions, err := DecodeE2SetupRequestPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Assert(t, identifier != nil)
	assert.DeepEqual(t, []byte{0x00, 0x02, 0x10}, []byte{identifier.Plmn[0], identifier.Plmn[1], identifier.Plmn[2]})
	assert.Equal(t, types.E2NodeTypeENB, identifier.NodeType)
	assert.DeepEqual(t, []byte{0x00, 0xE0, 0x00}, identifier.NodeIdentifier)

	assert.Assert(t, ranFunctions != nil)
	assert.Equal(t, 1, len(*ranFunctions))
}
