// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-en-gNB-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbID *e2ap_ies.GlobalE2NodeEnGnbId) ([]byte, error) {
	globalE2nodeEnGnbIDCP, err := newGlobalE2nodeEnGnbID(globalE2nodeEnGnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodeEnGnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalE2node_en_gNB_ID, unsafe.Pointer(globalE2nodeEnGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalE2nodeEnGnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbID *e2ap_ies.GlobalE2NodeEnGnbId) ([]byte, error) {
	globalE2nodeEnGnbIDCP, err := newGlobalE2nodeEnGnbID(globalE2nodeEnGnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodeEnGnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalE2node_en_gNB_ID, unsafe.Pointer(globalE2nodeEnGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalE2nodeEnGnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalE2nodeEnGnbID(bytes []byte) (*e2ap_ies.GlobalE2NodeEnGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalE2node_en_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalE2nodeEnGnbID((*C.GlobalE2node_en_gNB_ID_t)(unsafePtr))
}

func perDecodeGlobalE2nodeEnGnbID(bytes []byte) (*e2ap_ies.GlobalE2NodeEnGnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalE2node_en_gNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalE2nodeEnGnbID((*C.GlobalE2node_en_gNB_ID_t)(unsafePtr))
}

func newGlobalE2nodeEnGnbID(globalE2nodeEnGnbID *e2ap_ies.GlobalE2NodeEnGnbId) (*C.GlobalE2node_en_gNB_ID_t, error) {

	var err error
	globalE2nodeEnGnbIDC := C.GlobalE2node_en_gNB_ID_t{}

	globalGNbIDC, err := newGlobalenGnbID(globalE2nodeEnGnbID.GlobalGNbId)
	if err != nil {
		return nil, fmt.Errorf("newGlobalenGnbID() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	globalE2nodeEnGnbIDC.global_gNB_ID = *globalGNbIDC

	return &globalE2nodeEnGnbIDC, nil
}

func decodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbIDC *C.GlobalE2node_en_gNB_ID_t) (*e2ap_ies.GlobalE2NodeEnGnbId, error) {

	var err error
	globalE2nodeEnGnbID := e2ap_ies.GlobalE2NodeEnGnbId{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//GlobalGNbId: globalGNbId,
	}

	globalE2nodeEnGnbID.GlobalGNbId, err = decodeGlobalenGnbID(&globalE2nodeEnGnbIDC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGlobalenGnbID() %s", err.Error())
	}

	return &globalE2nodeEnGnbID, nil
}

func decodeGlobalE2nodeEnGnbIDBytes(array [8]byte) (*e2ap_ies.GlobalE2NodeEnGnbId, error) {
	globalE2nodeEnGnbIDC := (*C.GlobalE2node_en_gNB_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeGlobalE2nodeEnGnbID(globalE2nodeEnGnbIDC)
}
