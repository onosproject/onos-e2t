// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-SingleContainer.h"
//#include "RICaction-ToBeSetup-Item.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func newRanFunctionItemIesSingleContainer(rfItemIes *e2appducontents.RanfunctionItemIes) (*C.ProtocolIE_SingleContainer_1547P3_t, error) {
	pIeSC1547P3, err := newRANfunctionItemIEs(rfItemIes)

	return (*C.ProtocolIE_SingleContainer_1547P3_t)(pIeSC1547P3), err
}

func decodeRanFunctionItemIesSingleContainer(rfiIeScC *C.ProtocolIE_SingleContainer_1547P3_t) (*e2appducontents.RanfunctionItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfiIeScC, rfiIeScC)
	switch id := rfiIeScC.id; id {
	case C.long(v1beta1.ProtocolIeIDRanfunctionItem):
		return decodeRANfunctionItemIes(&rfiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RanFunctionItem %v", C.long(id))
	}

}
