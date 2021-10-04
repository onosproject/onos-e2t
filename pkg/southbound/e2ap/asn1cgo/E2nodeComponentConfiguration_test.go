// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

import (
	"encoding/hex"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentConfigurationMsg() (*e2ap_ies.E2NodeComponentConfiguration, error) {

	e2ncc := e2ap_ies.E2NodeComponentConfiguration{
		E2NodeComponentRequestPart:  []byte{0x00, 0x01},
		E2NodeComponentResponsePart: []byte{0x02, 0x03},
	}

	//if err := e2nodeComponentConfigUpdateGnb.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdateGnb %s", err.Error())
	//}
	return &e2ncc, nil
}

func Test_xerEncodingE2nodeComponentConfiguration(t *testing.T) {

	e2ncc, err := createE2nodeComponentConfigurationMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfiguration PDU")

	xer, err := xerEncodeE2nodeComponentConfiguration(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfiguration XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfiguration(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfiguration XER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetE2NodeComponentResponsePart(), result.GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2ncc.GetE2NodeComponentRequestPart(), result.GetE2NodeComponentRequestPart())
}

func Test_perEncodingE2nodeComponentConfiguration(t *testing.T) {

	e2ncc, err := createE2nodeComponentConfigurationMsg()
	assert.NilError(t, err, "Error creating E2nodeComponentConfiguration PDU")

	per, err := perEncodeE2nodeComponentConfiguration(e2ncc)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfiguration PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfiguration(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfiguration PER - decoded\n%v", result)
	assert.DeepEqual(t, e2ncc.GetE2NodeComponentResponsePart(), result.GetE2NodeComponentResponsePart())
	assert.DeepEqual(t, e2ncc.GetE2NodeComponentRequestPart(), result.GetE2NodeComponentRequestPart())
}
