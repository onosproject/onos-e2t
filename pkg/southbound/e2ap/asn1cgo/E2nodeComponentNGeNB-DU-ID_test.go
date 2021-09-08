// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentNgEnbDuIDMsg() (*e2ap_ies.E2NodeComponentNgeNbDuId, error) {

	e2nodeComponentNgEnbDuID := e2ap_ies.E2NodeComponentNgeNbDuId{
		NgEnbDuId: &e2ap_ies.NgenbDuId{
			Value: 11111,
		},
	}

	//if err := e2nodeComponentGnbDuID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentGnbDuId %s", err.Error())
	//}
	return &e2nodeComponentNgEnbDuID, nil
}

func Test_xerEncodingE2nodeComponentNgEnbDuID(t *testing.T) {

	e2nodeComponentNgEnbDuID, err := createE2nodeComponentNgEnbDuIDMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentNgEnbDuId PDU")

	xer, err := xerEncodeE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID)
	assert.NilError(t, err)
	//assert.Equal(t, 85, len(xer))
	t.Logf("E2nodeComponentNgEnbDuID XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentNgEnbDuID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentNgEnbDuID XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentNgEnbDuID.GetNgEnbDuId().GetValue(), result.GetNgEnbDuId().GetValue())
}

func Test_perEncodingE2nodeComponentNgEnbDuID(t *testing.T) {

	e2nodeComponentNgEnbDuID, err := createE2nodeComponentNgEnbDuIDMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentNgEnbDuId PDU")

	per, err := perEncodeE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID)
	assert.NilError(t, err)
	//assert.Equal(t, 2, len(per))
	t.Logf("E2nodeComponentNgEnbDuID PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentNgEnbDuID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentNgEnbDuID PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentNgEnbDuID.GetNgEnbDuId().GetValue(), result.GetNgEnbDuId().GetValue())
}
