// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentNGeNB-CU-UP-ID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID *e2ap_ies.E2NodeComponentNgeNbDuId) ([]byte, error) {
	e2nodeComponentNgEnbDuIDCP, err := newE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentNgEnbDuID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentNGeNB_DU_ID, unsafe.Pointer(e2nodeComponentNgEnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentNgEnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID *e2ap_ies.E2NodeComponentNgeNbDuId) ([]byte, error) {
	e2nodeComponentNgEnbDuIDCP, err := newE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentNgEnbDuID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentNGeNB_DU_ID, unsafe.Pointer(e2nodeComponentNgEnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentNgEnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentNgEnbDuID(bytes []byte) (*e2ap_ies.E2NodeComponentNgeNbDuId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentNGeNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentNgEnbDuID((*C.E2nodeComponentNGeNB_DU_ID_t)(unsafePtr))
}

func perDecodeE2nodeComponentNgEnbDuID(bytes []byte) (*e2ap_ies.E2NodeComponentNgeNbDuId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentNGeNB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentNgEnbDuID((*C.E2nodeComponentNGeNB_DU_ID_t)(unsafePtr))
}

func newE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuID *e2ap_ies.E2NodeComponentNgeNbDuId) (*C.E2nodeComponentNGeNB_DU_ID_t, error) {

	var err error
	e2nodeComponentNgEnbDuIDC := C.E2nodeComponentNGeNB_DU_ID_t{}

	ngEnbDuIDC, err := newNgEnbDuID(e2nodeComponentNgEnbDuID.NgEnbDuId)
	if err != nil {
		return nil, fmt.Errorf("newNgEnbDuID() %s", err.Error())
	}

	e2nodeComponentNgEnbDuIDC.ngENB_DU_ID = *ngEnbDuIDC

	return &e2nodeComponentNgEnbDuIDC, nil
}

func decodeE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuIDC *C.E2nodeComponentNGeNB_DU_ID_t) (*e2ap_ies.E2NodeComponentNgeNbDuId, error) {

	var err error
	e2nodeComponentNgEnbDuID := e2ap_ies.E2NodeComponentNgeNbDuId{}

	e2nodeComponentNgEnbDuID.NgEnbDuId, err = decodeNgEnbDuID(&e2nodeComponentNgEnbDuIDC.ngENB_DU_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeNgEnbDuID() %s", err.Error())
	}

	return &e2nodeComponentNgEnbDuID, nil
}

func decodeE2nodeComponentNgEnbDuIDBytes(array [8]byte) (*e2ap_ies.E2NodeComponentNgeNbDuId, error) {
	e2nodeComponentNgEnbDuIDC := (*C.E2nodeComponentNGeNB_DU_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentNgEnbDuID(e2nodeComponentNgEnbDuIDC)
}
