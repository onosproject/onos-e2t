// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2ap_ies/pdubuilder"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2nodeComponentIDGNbCUUP() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbCuUp{
			E2NodeComponentTypeGnbCuUp: &e2ap_ies.E2NodeComponentGnbCuUpId{
				GNbCuUpId: &e2ap_ies.GnbCuUpId{
					Value: 21,
				},
			},
		},
	}

	if err := e2nodeComponentID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	}
	return &e2nodeComponentID, nil
}

func createE2nodeComponentIDGNbDU() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentTypeGnbDu{
			E2NodeComponentTypeGnbDu: &e2ap_ies.E2NodeComponentGnbDuId{
				GNbDuId: &e2ap_ies.GnbDuId{
					Value: 1234,
				},
			},
		},
	}

	if err := e2nodeComponentID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	}
	return &e2nodeComponentID, nil
}

func Test_xerEncodingE2nodeComponentID(t *testing.T) {

	e2nodeComponentID, err := createE2nodeComponentIDGNbCUUP()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err := xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	assert.Equal(t, 152, len(xer))
	t.Logf("E2nodeComponentID (GNb-CU-UP) XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (GNb-CU-UP) XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDGNbDU()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	assert.Equal(t, 142, len(xer))
	t.Logf("E2nodeComponentID (GNb-DU) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (GNb-DU) XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue(), result.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue())
}

func Test_perEncodingE2nodeComponentID(t *testing.T) {

	e2nodeComponentID, err := createE2nodeComponentIDGNbCUUP()
	assert.NilError(t, err, "Error creating E2nodeComponentId PDU")

	per, err := perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2nodeComponentID (GNb-CU-UP) PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (GNb-CU-UP) PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue(), result.GetE2NodeComponentTypeGnbCuUp().GetGNbCuUpId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDGNbDU()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("E2nodeComponentID (GNb-DU) PER\n%s", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (GNb-DU) PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue(), result.GetE2NodeComponentTypeGnbDu().GetGNbDuId().GetValue())
}
