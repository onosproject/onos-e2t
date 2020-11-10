// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Field.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"unsafe"
)

func newRanFunctionsIDcauseList(rfIDcl *e2appducontents.RanfunctionsIdcauseList) (*C.RANfunctionsIDcause_List_t, error) {
	rfIDclC := new(C.RANfunctionsIDcause_List_t)
	for _, rfIDCause := range rfIDcl.GetValue() {
		rfIDcauseC, err := newRanFunctionIDcauseItemIesSingleContainer(rfIDCause)
		if err != nil {
			return nil, fmt.Errorf("error on newRanFunctionIDcauseItemIesSingleContainer() %s", err.Error())
		}
		if _, err = C.asn_sequence_add(unsafe.Pointer(rfIDclC), unsafe.Pointer(rfIDcauseC)); err != nil {
			return nil, err
		}
	}

	return rfIDclC, nil
}
