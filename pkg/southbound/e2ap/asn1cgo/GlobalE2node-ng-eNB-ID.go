// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-ng-eNB-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbID *e2ap_ies.GlobalE2NodeNgEnbId) ([]byte, error) {
	globalE2nodeNgEnbIDCP, err := newGlobalE2nodeNgEnbID(globalE2nodeNgEnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodeNgEnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalE2node_ng_eNB_ID, unsafe.Pointer(globalE2nodeNgEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodeNgEnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbID *e2ap_ies.GlobalE2NodeNgEnbId) ([]byte, error) {
	globalE2nodeNgEnbIDCP, err := newGlobalE2nodeNgEnbID(globalE2nodeNgEnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodeNgEnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalE2node_ng_eNB_ID, unsafe.Pointer(globalE2nodeNgEnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodeNgEnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalE2nodeNgEnbID(bytes []byte) (*e2ap_ies.GlobalE2NodeNgEnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalE2node_ng_eNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalE2nodeNgEnbID((*C.GlobalE2node_ng_eNB_ID_t)(unsafePtr))
}

func perDecodeGlobalE2nodeNgEnbID(bytes []byte) (*e2ap_ies.GlobalE2NodeNgEnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalE2node_ng_eNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalE2nodeNgEnbID((*C.GlobalE2node_ng_eNB_ID_t)(unsafePtr))
}

func newGlobalE2nodeNgEnbID(globalE2nodeNgEnbID *e2ap_ies.GlobalE2NodeNgEnbId) (*C.GlobalE2node_ng_eNB_ID_t, error) {

	var err error
	globalE2nodeNgEnbIDC := C.GlobalE2node_ng_eNB_ID_t{}

	globalNgENbIDC, err := newGlobalngeNbID(globalE2nodeNgEnbID.GlobalNgENbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalngeNbID() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	globalE2nodeNgEnbIDC.global_ng_eNB_ID = *globalNgENbIDC

	return &globalE2nodeNgEnbIDC, nil
}

func decodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbIDC *C.GlobalE2node_ng_eNB_ID_t) (*e2ap_ies.GlobalE2NodeNgEnbId, error) {

	var err error
	globalE2nodeNgEnbID := e2ap_ies.GlobalE2NodeNgEnbId{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//GlobalNgENbId: globalNgENbId,
	}

	globalE2nodeNgEnbID.GlobalNgENbId, err = decodeGlobalngeNbID(&globalE2nodeNgEnbIDC.global_ng_eNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalngeNbID() %s", err.Error())
	}

	return &globalE2nodeNgEnbID, nil
}

func decodeGlobalE2nodeNgEnbIDBytes(array [8]byte) (*e2ap_ies.GlobalE2NodeNgEnbId, error) {
	globalE2nodeNgEnbIDC := (*C.GlobalE2node_ng_eNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalE2nodeNgEnbID(globalE2nodeNgEnbIDC)
}
