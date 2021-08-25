// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentConfigUpdateGnb() (*e2ap_ies.E2NodeComponentConfigUpdate, error) {

	e2nodeComponentConfigUpdate := e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate{
			GNbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateGnb{
				NgApconfigUpdate: []byte("ng_AP"),
				XnApconfigUpdate: []byte("xn_AP"),
				E1ApconfigUpdate: []byte("e1_AP"),
				F1ApconfigUpdate: []byte("f1_AP"),
				X2ApconfigUpdate: []byte("x2_AP"),
			},
		},
	}

	//if err := e2nodeComponentConfigUpdate.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdate %s", err.Error())
	//}
	return &e2nodeComponentConfigUpdate, nil
}

func createE2nodeComponentConfigUpdateEnb() (*e2ap_ies.E2NodeComponentConfigUpdate, error) {

	e2nodeComponentConfigUpdate := e2ap_ies.E2NodeComponentConfigUpdate{
		E2NodeComponentConfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdate_ENbconfigUpdate{
			ENbconfigUpdate: &e2ap_ies.E2NodeComponentConfigUpdateEnb{
				NgApconfigUpdate: []byte("ng_AP"),
				XnApconfigUpdate: []byte("xn_AP"),
				W1ApconfigUpdate: []byte("w1_AP"),
				S1ApconfigUpdate: []byte("s1_AP"),
				X2ApconfigUpdate: []byte("x2_AP"),
			},
		},
	}

	//if err := e2nodeComponentConfigUpdate.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentConfigUpdate %s", err.Error())
	//}
	return &e2nodeComponentConfigUpdate, nil
}

func Test_xerEncodingE2nodeComponentConfigUpdate(t *testing.T) {

	e2nodeComponentConfigUpdate, err := createE2nodeComponentConfigUpdateGnb()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdate (GNb) PDU")

	xer, err := xerEncodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdate (GNb) XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentConfigUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdate (GNb) XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetE1ApconfigUpdate(), result.GetGNbconfigUpdate().GetE1ApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetF1ApconfigUpdate(), result.GetGNbconfigUpdate().GetF1ApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetNgApconfigUpdate(), result.GetGNbconfigUpdate().GetNgApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetGNbconfigUpdate().GetXnApconfigUpdate())

	e2nodeComponentConfigUpdate, err = createE2nodeComponentConfigUpdateEnb()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdate (ENb) PDU")

	xer, err = xerEncodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdate (ENb) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentConfigUpdate(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdate (ENb) XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetENbconfigUpdate().GetX2ApconfigUpdate(), result.GetENbconfigUpdate().GetX2ApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetENbconfigUpdate().GetS1ApconfigUpdate(), result.GetENbconfigUpdate().GetS1ApconfigUpdate())
}

func Test_perEncodingE2nodeComponentConfigUpdate(t *testing.T) {

	e2nodeComponentConfigUpdate, err := createE2nodeComponentConfigUpdateGnb()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdate (GNb) PDU")

	per, err := perEncodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdate (GNb) PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentConfigUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdate (GNb) PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetE1ApconfigUpdate(), result.GetGNbconfigUpdate().GetE1ApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetF1ApconfigUpdate(), result.GetGNbconfigUpdate().GetF1ApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetNgApconfigUpdate(), result.GetGNbconfigUpdate().GetNgApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetGNbconfigUpdate().GetXnApconfigUpdate(), result.GetGNbconfigUpdate().GetXnApconfigUpdate())

	e2nodeComponentConfigUpdate, err = createE2nodeComponentConfigUpdateEnb()
	assert.NilError(t, err, "Error creating E2nodeComponentConfigUpdate (ENb) PDU")

	per, err = perEncodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentConfigUpdate (ENb) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentConfigUpdate(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentConfigUpdate (ENb) PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetENbconfigUpdate().GetX2ApconfigUpdate(), result.GetENbconfigUpdate().GetX2ApconfigUpdate())
	assert.DeepEqual(t, e2nodeComponentConfigUpdate.GetENbconfigUpdate().GetS1ApconfigUpdate(), result.GetENbconfigUpdate().GetS1ApconfigUpdate())
}
