// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"

	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2ap_ies/pdubuilder"
	"testing"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"gotest.tools/assert"
)

func createE2nodeComponentIDNg() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDNg("NG-interface")

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func createE2nodeComponentIDXn() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDXn(&e2ap_ies.GlobalNgRannodeId{
		GlobalNgRannodeId: &e2ap_ies.GlobalNgRannodeId_GNb{
			GNb: &e2ap_ies.GlobalgNbId{
				PlmnId: &e2ap_commondatatypes.PlmnIdentity{
					Value: []byte{0x00, 0x00, 0x0F},
				},
				GnbId: &e2ap_ies.GnbIdChoice{
					GnbIdChoice: &e2ap_ies.GnbIdChoice_GnbId{
						GnbId: &asn1.BitString{
							Value: []byte{0x00, 0x00, 0x00, 0x01},
							Len:   32,
						},
					},
				},
			},
		},
	})

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func createE2nodeComponentIDE1() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDE1(2)

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func createE2nodeComponentIDF1() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDF1(2)

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func createE2nodeComponentIDW1() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDW1(2)

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func createE2nodeComponentIDS1() (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDS1("ONF")

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func createE2nodeComponentIDX2() (*e2ap_ies.E2NodeComponentId, error) {

	enbID, err := pdubuilder.CreateEnbIDHome(&asn1.BitString{
		Value: []byte{0x00, 0xA7, 0xDD, 0xF0},
		Len:   28,
	})
	if err != nil {
		return nil, err
	}

	gEnbID, err := pdubuilder.CreateGlobalEnbID([]byte{0xAA, 0xBB, 0xCC}, enbID)
	if err != nil {
		return nil, err
	}

	gEnGnbID, err := pdubuilder.CreateGlobalEnGnbID([]byte{0xFF, 0xCD, 0xBF}, &asn1.BitString{
		Value: []byte{0xFA, 0x2C, 0xD4, 0xF8},
		Len:   29,
	})
	if err != nil {
		return nil, err
	}

	e2nodeComponentID := pdubuilder.CreateE2NodeComponentIDX2(gEnbID, gEnGnbID)

	//if err := e2nodeComponentID.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2nodeComponentID %s", err.Error())
	//}
	return e2nodeComponentID, nil
}

func Test_xerEncodingE2nodeComponentID(t *testing.T) {

	e2nodeComponentID, err := createE2nodeComponentIDNg()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err := xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (Ng) XER\n%s", string(xer))

	result, err := xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (Ng) XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDXn()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (Xn) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (Xn) XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetPlmnId().GetValue(), result.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetPlmnId().GetValue())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetGnbId().GetGnbId().GetValue(), result.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetGnbId().GetGnbId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDE1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (E1) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (E1) XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeE1().GetGNbCuCpId().GetValue(), result.GetE2NodeComponentInterfaceTypeE1().GetGNbCuCpId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDF1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (F1) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (F1) XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeF1().GetGNbDuId().GetValue(), result.GetE2NodeComponentInterfaceTypeF1().GetGNbDuId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDW1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (W1) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (W1) XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDS1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (S1) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (S1) XER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue(), result.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDX2()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	xer, err = xerEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (X2) XER\n%s", string(xer))

	result, err = xerDecodeE2nodeComponentID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (X2) XER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetPLmnIdentity().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue())
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetLen(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetLen())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
}

func Test_perEncodingE2nodeComponentID(t *testing.T) {

	e2nodeComponentID, err := createE2nodeComponentIDNg()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err := perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (Ng) PER\n%v", hex.Dump(per))

	result, err := perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (Ng) PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue(), result.GetE2NodeComponentInterfaceTypeNg().GetAmfName().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDXn()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (Xn) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (Xn) PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetPlmnId().GetValue(), result.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetPlmnId().GetValue())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetGnbId().GetGnbId().GetValue(), result.GetE2NodeComponentInterfaceTypeXn().GetGlobalNgRanNodeId().GetGNb().GetGnbId().GetGnbId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDE1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (E1) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (E1) PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeE1().GetGNbCuCpId().GetValue(), result.GetE2NodeComponentInterfaceTypeE1().GetGNbCuCpId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDF1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (F1) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (F1) PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeF1().GetGNbDuId().GetValue(), result.GetE2NodeComponentInterfaceTypeF1().GetGNbDuId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDW1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (W1) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (W1) PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue(), result.GetE2NodeComponentInterfaceTypeW1().GetNgENbDuId().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDS1()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (S1) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (S1) PER - decoded\n%v", result)
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue(), result.GetE2NodeComponentInterfaceTypeS1().GetMmeName().GetValue())

	e2nodeComponentID, err = createE2nodeComponentIDX2()
	assert.NilError(t, err, "Error creating E2nodeComponentID PDU")

	per, err = perEncodeE2nodeComponentID(e2nodeComponentID)
	assert.NilError(t, err)
	t.Logf("E2nodeComponentID (X2) PER\n%v", hex.Dump(per))

	result, err = perDecodeE2nodeComponentID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2nodeComponentID (X2) PER - decoded\n%v", result)
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetPLmnIdentity().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetValue())
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetLen(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalENbId().GetENbId().GetHomeENbId().GetLen())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetPLmnIdentity().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
	assert.DeepEqual(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue())
	assert.Equal(t, e2nodeComponentID.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen(), result.GetE2NodeComponentInterfaceTypeX2().GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
}
