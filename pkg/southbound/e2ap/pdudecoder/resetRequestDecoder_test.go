// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

//import (
//	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
//	"io/ioutil"
//	"testing"
//
//	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
//	"gotest.tools/assert"
//)
//
//func Test_DecodeResetRequestPdu(t *testing.T) {
//	rrXer, err := ioutil.ReadFile("../test/ResetRequest.xml")
//	assert.NilError(t, err, "Unexpected error when loading file")
//	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(rrXer)
//	assert.NilError(t, err)
//
//	cause, transactionID, err := DecodeResetRequestPdu(e2apPdu)
//	assert.NilError(t, err)
//	//assert.Assert(t, ricIdentity != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 23, 24 & 25
//
//	assert.Equal(t, e2apies.CauseProtocol_CAUSE_PROTOCOL_TRANSFER_SYNTAX_ERROR, cause.GetProtocol())
//	if transactionID != nil {
//		assert.Equal(t, int32(1), *transactionID)
//	}
//}
