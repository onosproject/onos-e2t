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

func newRanFunctionIDItemIesSingleContainer(rfIDItemIes *e2appducontents.RanfunctionIdItemIes) (*C.ProtocolIE_SingleContainer_1547P4_t, error) {
	pIeSC1547P4, err := newRANfunctionIDItemIEs(rfIDItemIes)

	return (*C.ProtocolIE_SingleContainer_1547P4_t)(pIeSC1547P4), err
}

func newRanFunctionIDcauseItemIesSingleContainer(rfIDcauseItemIes *e2appducontents.RanfunctionIdcauseItemIes) (*C.ProtocolIE_SingleContainer_1547P5_t, error) {
	pIeSC1547P5, err := newRANfunctionIDCauseItemIEs(rfIDcauseItemIes)

	return (*C.ProtocolIE_SingleContainer_1547P5_t)(pIeSC1547P5), err
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

func decodeRanFunctionIDItemIesSingleContainer(rfIDiIeScC *C.ProtocolIE_SingleContainer_1547P4_t) (*e2appducontents.RanfunctionIdItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIeScC, rfIDiIeScC)
	switch id := rfIDiIeScC.id; id {
	case C.long(v1beta1.ProtocolIeIDRanfunctionIDItem):
		return decodeRANfunctionIDItemIes(&rfIDiIeScC.value)
	default:
		return nil, fmt.Errorf("unexpected id for RanFunctionItem %v", C.long(id))
	}

}
