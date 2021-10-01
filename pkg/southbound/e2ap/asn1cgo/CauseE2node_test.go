// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createCauseE2nodeComponentUnknown() e2ap_ies.CauseE2Node {
	return e2ap_ies.CauseE2Node_CAUSE_E2NODE_E2NODE_COMPONENT_UNKNOWN
}

func Test_xerEncodingCauseE2node(t *testing.T) {

	causeE2node := createCauseE2nodeComponentUnknown()

	xer, err := xerEncodeCauseE2node(&causeE2node)
	assert.NilError(t, err)
	t.Logf("CauseE2node (Component unknown) XER\n%s", string(xer))

	result, err := xerDecodeCauseE2node(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseE2node (Component unknown) XER - decoded\n%v", result)
	assert.Equal(t, causeE2node.Number(), result.Number())
}

func Test_perEncodingCauseE2node(t *testing.T) {

	causeE2node := createCauseE2nodeComponentUnknown()

	per, err := perEncodeCauseE2node(&causeE2node)
	assert.NilError(t, err)
	t.Logf("CauseE2node (Component unknown) PER\n%v", hex.Dump(per))

	result, err := perDecodeCauseE2node(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("CauseE2node (Component unknown) PER - decoded\n%v", result)
	assert.Equal(t, causeE2node.Number(), result.Number())
}
