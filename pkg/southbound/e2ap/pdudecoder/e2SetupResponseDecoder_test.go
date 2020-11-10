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

	ricIdentity, ranFunctionsAccepted, err := DecodeE2SetupResponsePdu(e2apPdu)
	assert.NilError(t, err)
	assert.Assert(t, ricIdentity != nil)
	assert.Equal(t, "ONF", string([]byte{ricIdentity.PlmnID[0], ricIdentity.PlmnID[1], ricIdentity.PlmnID[2]}))
	assert.Equal(t, 20, int(ricIdentity.RicIdentifier.RicIdentifierLen))
	assert.Equal(t, 0xBCDE, int(ricIdentity.RicIdentifier.RicIdentifierValue))

	assert.Equal(t, 2, len(ranFunctionsAccepted))
	rfa100, ok := ranFunctionsAccepted[100]
	assert.Assert(t, ok, "expected a key '100'")
	assert.Equal(t, 2, int(rfa100))
	rfa200, ok := ranFunctionsAccepted[200]
	assert.Assert(t, ok, "expected a key '200'")
	assert.Equal(t, 2, int(rfa200))
}
