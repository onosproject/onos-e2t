// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createCauseTransportUnspecified() e2ap_ies.CauseTransport {
	return e2ap_ies.CauseTransport_CAUSE_TRANSPORT_UNSPECIFIED
}

func createCauseTransportResourceUnavailable() e2ap_ies.CauseTransport {
	return e2ap_ies.CauseTransport_CAUSE_TRANSPORT_TRANSPORT_RESOURCE_UNAVAILABLE
}

func Test_xerEncodingCauseTransport(t *testing.T) {

	causeTransport := createCauseTransportUnspecified()

	xer, err := xerEncodeCauseTransport(&causeTransport)
	assert.NilError(t, err)
	assert.Equal(t, 48, len(xer))
	t.Logf("CauseTransport (Unspecified) XER\n%s", string(xer))

	result, err := xerDecodeCauseTransport(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseTransport (Unspecified) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeTransport, *result)
	}

	causeTransport = createCauseTransportResourceUnavailable()

	xer, err = xerEncodeCauseTransport(&causeTransport)
	assert.NilError(t, err)
	assert.Equal(t, 67, len(xer))
	t.Logf("CauseTransport (Resource Unavailable) XER\n%s", string(xer))

	result, err = xerDecodeCauseTransport(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseTransport (Resource Unavailable) XER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeTransport, *result)
	}
}

func Test_perEncodingCauseTransport(t *testing.T) {

	causeTransport := createCauseTransportUnspecified()

	per, err := perEncodeCauseTransport(&causeTransport)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseTransport (Unspecified) PER\n%v", hex.Dump(per))

	result, err := perDecodeCauseTransport(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseTransport (Unspecified) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeTransport, *result)
	}

	causeTransport = createCauseTransportResourceUnavailable()

	per, err = perEncodeCauseTransport(&causeTransport)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("CauseTransport (Resource Unavailable) PER\n%v", hex.Dump(per))

	result, err = perDecodeCauseTransport(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseTransport (Resource Unavailable) PER - decoded\n%v", result)
	if result != nil {
		assert.Equal(t, causeTransport, *result)
	}
}
