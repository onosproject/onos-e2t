// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunctions-List.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"unsafe"
)

// XerEncodeGlobalE2nodeID - used only in tests
func XerEncodeanFunctionsList(ranFunctionsList *e2appducontents.RanfunctionsList) ([]byte, error) {
	ranFunctionsListC, err := newRanFunctionsList(ranFunctionsList)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RANfunctions_List, unsafe.Pointer(ranFunctionsListC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeGlobalE2nodeID - used only in tests
func PerEncodeGanFunctionsList(ranFunctionsList *e2appducontents.RanfunctionsList) ([]byte, error) {
	ranFunctionsListC, err := newRanFunctionsList(ranFunctionsList)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RANfunctions_List, unsafe.Pointer(ranFunctionsListC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func newRanFunctionsList(rfl *e2appducontents.RanfunctionsList) (*C.RANfunctions_List_t, error) {

	return nil, fmt.Errorf("newRanFunctionsList() Not yet implemented")
}

func decodeRanFunctionsList(ranFunctionsListC [48]byte) (*e2appducontents.RanfunctionsList, error) {

	return nil, fmt.Errorf("decodeRanFunctionsList() Not yet implemented")
}
