// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/asn1cgo"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeRicIndicationPdu(t *testing.T) {
	e2setupRequestXer, err := ioutil.ReadFile("../test/RicIndication.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2apPdu, err := asn1cgo.XerDecodeE2apPdu(e2setupRequestXer)
	assert.NilError(t, err)

	ranFunctionID, ricActionID, ricCallProcessID, ricIndicationHeader, ricIndicationMessage, ricIndicationSn,
		ricIndicationType, ricRequest, err := DecodeRicIndicationPdu(e2apPdu)
	assert.NilError(t, err)
	assert.Equal(t, 9, int(ranFunctionID), "unexpected ranFunctionID")
	assert.Equal(t, 2, int(ricActionID), "unexpected ricActionID")
	assert.DeepEqual(t, []byte{'1', '2', '3'}, []byte(*ricIndicationHeader))
	assert.DeepEqual(t, []byte{'4', '5', '6'}, []byte(*ricIndicationMessage))
	assert.DeepEqual(t, []byte{'7', '8', '9'}, []byte(*ricCallProcessID))
	assert.Equal(t, 1, int(ricIndicationSn), "unexpected ricIndicationSn")
	assert.Equal(t, e2apies.RicindicationType_RICINDICATION_TYPE_INSERT, ricIndicationType, "unexpected ricIndicationType")
	assert.Assert(t, ricRequest != nil)

}
