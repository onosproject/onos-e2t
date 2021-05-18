// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import (
	"encoding/hex"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/pdubuilder"
	"gotest.tools/assert"
	"testing"
)

func createE2connectionUpdateMsg() (*e2ap_pdu_contents.E2ConnectionUpdate, error) {

	e2connectionUpdate, err := pdubuilder.CreateE2connectionUpdateE2apPdu()
	if err != nil {
		return nil, err
	}

	if err := e2connectionUpdate.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2connectionUpdate %s", err.Error())
	}
	return e2connectionUpdate.GetInitiatingMessage().GetProcedureCode().GetE2ConnectionUpdate().GetInitiatingMessage(), nil
}

func Test_xerEncodingE2connectionUpdate(t *testing.T) {

	e2connectionUpdate, err := createE2connectionUpdateMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdate PDU")

	xer, err := xerEncodeE2connectionUpdate(e2connectionUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 3500, len(xer))
	t.Logf("E2connectionUpdate XER\n%s", string(xer))

	//result, err := xerDecodeE2connectionUpdate(xer)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("E2connectionUpdate XER - decoded\n%v", result)
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
}

func Test_perEncodingE2connectionUpdate(t *testing.T) {

	e2connectionUpdate, err := createE2connectionUpdateMsg()
	assert.NilError(t, err, "Error creating E2connectionUpdate PDU")

	per, err := perEncodeE2connectionUpdate(e2connectionUpdate)
	assert.NilError(t, err)
	assert.Equal(t, 68, len(per))
	t.Logf("E2connectionUpdate PER\n%v", hex.Dump(per))

	//result, err := perDecodeE2connectionUpdate(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("E2connectionUpdate PER - decoded\n%v", result)
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlUsage())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes44().GetConnectionAdd().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage(), result.GetProtocolIes().GetE2ApProtocolIes45().GetConnectionModify().GetValue()[0].GetValue().GetTnlUsage())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlPort().GetLen())
	//assert.Equal(t, e2connectionUpdate.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen(), result.GetProtocolIes().GetE2ApProtocolIes46().GetConnectionRemove().GetValue()[0].GetValue().GetTnlInformation().GetTnlAddress().GetLen())
}
