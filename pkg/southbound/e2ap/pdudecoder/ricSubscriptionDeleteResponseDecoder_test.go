// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"gotest.tools/assert"
	"testing"
)

func Test_DecodeRicSubscriptionDeleteResponsePdu(t *testing.T) {
	e2apPdu, err := pdubuilder.CreateRicSubscriptionDeleteResponseE2apPdu(&types.RicRequest{
		RequestorID: 22,
		InstanceID:  6,
	}, 9)
	assert.NilError(t, err)
	assert.Assert(t, e2apPdu != nil)

	rfID, rrID, err := DecodeRicSubscriptionDeleteResponsePdu(e2apPdu)
	assert.NilError(t, err)
	//assert.Assert(t, rfID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on line 23
	assert.Equal(t, 9, int(*rfID))

	//assert.Assert(t, rrID != nil) //Commented due to the Linters (v1.34.1) error - possible nil pointer dereference (https://staticcheck.io/docs/checks#SA5011) on lines 26 & 27
	assert.Equal(t, 22, int(rrID.RequestorID))
	assert.Equal(t, 6, int(rrID.InstanceID))
}
