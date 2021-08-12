// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2connectionUpdatePdu(t *testing.T) {
	e2ncuXer, err := ioutil.ReadFile("../test/E2connectionUpdate.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2ncuXer)
	assert.NilError(t, err)

	connSetup, connModify, connRemove, err := DecodeE2connectionUpdatePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25

	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlAddress.GetLen()), int32(64))
	assert.Equal(t, int32(connSetup[0].TnlInformation.TnlPort.GetLen()), int32(16))
	assert.Equal(t, int32(connSetup[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_BOTH))
	assert.Equal(t, int32(connModify[0].TnlInformation.TnlAddress.GetLen()), int32(64))
	assert.Equal(t, int32(connModify[0].TnlInformation.TnlPort.GetLen()), int32(16))
	assert.Equal(t, int32(connModify[0].TnlUsage), int32(e2ap_ies.Tnlusage_TNLUSAGE_RIC_SERVICE))
	assert.Equal(t, int32(connRemove[0].TnlAddress.GetLen()), int32(64))
	assert.Equal(t, int32(connRemove[0].TnlPort.GetLen()), int32(16))
}
