// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentGNB-DU-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentGnbDuID(e2nodeComponentGnbDuID *e2ap_ies.E2NodeComponentGnbDuId) ([]byte, error) {
	e2nodeComponentGnbDuIDCP, err := newE2nodeComponentGnbDuID(e2nodeComponentGnbDuID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentGnbDuID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentGNB_DU_ID, unsafe.Pointer(e2nodeComponentGnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentGnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentGnbDuID(e2nodeComponentGnbDuID *e2ap_ies.E2NodeComponentGnbDuId) ([]byte, error) {
	e2nodeComponentGnbDuIDCP, err := newE2nodeComponentGnbDuID(e2nodeComponentGnbDuID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentGnbDuID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentGNB_DU_ID, unsafe.Pointer(e2nodeComponentGnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentGnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentGnbDuID(bytes []byte) (*e2ap_ies.E2NodeComponentGnbDuId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentGNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentGnbDuID((*C.E2nodeComponentGNB_DU_ID_t)(unsafePtr))
}

func perDecodeE2nodeComponentGnbDuID(bytes []byte) (*e2ap_ies.E2NodeComponentGnbDuId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentGNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentGnbDuID((*C.E2nodeComponentGNB_DU_ID_t)(unsafePtr))
}

func newE2nodeComponentGnbDuID(e2nodeComponentGnbDuID *e2ap_ies.E2NodeComponentGnbDuId) (*C.E2nodeComponentGNB_DU_ID_t, error) {

	var err error
	e2nodeComponentGnbDuIDC := C.E2nodeComponentGNB_DU_ID_t{}

	gNbDuIDC, err := newGnbDuID(e2nodeComponentGnbDuID.GNbDuId)
	if err != nil {
		return nil, fmt.Errorf("newGnbDuID() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2nodeComponentGnbDuIDC.gNB_DU_ID = *gNbDuIDC

	return &e2nodeComponentGnbDuIDC, nil
}

func decodeE2nodeComponentGnbDuID(e2nodeComponentGnbDuIDC *C.E2nodeComponentGNB_DU_ID_t) (*e2ap_ies.E2NodeComponentGnbDuId, error) {

	var err error
	e2nodeComponentGnbDuID := e2ap_ies.E2NodeComponentGnbDuId{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//GNbDuId: gNbDuId,

	}

	e2nodeComponentGnbDuID.GNbDuId, err = decodeGnbDuID(&e2nodeComponentGnbDuIDC.gNB_DU_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGnbDuID() %s", err.Error())
	}

	return &e2nodeComponentGnbDuID, nil
}

func decodeE2nodeComponentGnbDuIDBytes(array [8]byte) (*e2ap_ies.E2NodeComponentGnbDuId, error) {
	e2nodeComponentGnbDuIDC := (*C.E2nodeComponentGNB_DU_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentGnbDuID(e2nodeComponentGnbDuIDC)
}
