// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SetupResponsePdu(t *testing.T) {
	e2setupResponseXer, err := ioutil.ReadFile("../test/E2setupResponse.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupResponseXer)
	assert.NilError(t, err)

	identifier, _, err := DecodeE2SetupResponsePdu(e2apPdu)
	assert.NilError(t, err)
	assert.Assert(t, identifier != nil)
	assert.Equal(t, "ONF", string([]byte{identifier.PlmnID[0], identifier.PlmnID[1], identifier.PlmnID[2]}))
	assert.Equal(t, 20, int(identifier.RicIdentifier.RicIdentifierLen))
	assert.Equal(t, 0xBCDE, int(identifier.RicIdentifier.RicIdentifierValue))

}
