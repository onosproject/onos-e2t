// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentGNB-CU-UP-ID.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID *e2ap_ies.E2NodeComponentGnbCuUpId) ([]byte, error) {
	e2nodeComponentGnbCuUpIDCP, err := newE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentGnbCuUpID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentGNB_CU_UP_ID, unsafe.Pointer(e2nodeComponentGnbCuUpIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentGnbCuUpID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID *e2ap_ies.E2NodeComponentGnbCuUpId) ([]byte, error) {
	e2nodeComponentGnbCuUpIDCP, err := newE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentGnbCuUpID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentGNB_CU_UP_ID, unsafe.Pointer(e2nodeComponentGnbCuUpIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentGnbCuUpID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentGnbCuUpID(bytes []byte) (*e2ap_ies.E2NodeComponentGnbCuUpId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentGNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentGnbCuUpID((*C.E2nodeComponentGNB_CU_UP_ID_t)(unsafePtr))
}

func perDecodeE2nodeComponentGnbCuUpID(bytes []byte) (*e2ap_ies.E2NodeComponentGnbCuUpId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentGNB_CU_UP_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentGnbCuUpID((*C.E2nodeComponentGNB_CU_UP_ID_t)(unsafePtr))
}

func newE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpID *e2ap_ies.E2NodeComponentGnbCuUpId) (*C.E2nodeComponentGNB_CU_UP_ID_t, error) {

	var err error
	e2nodeComponentGnbCuUpIDC := C.E2nodeComponentGNB_CU_UP_ID_t{}

	gNbCuUpIDC, err := newGnbCuUpID(e2nodeComponentGnbCuUpID.GNbCuUpId)
	if err != nil {
		return nil, fmt.Errorf("newGnbCuUpID() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2nodeComponentGnbCuUpIDC.gNB_CU_UP_ID = *gNbCuUpIDC

	return &e2nodeComponentGnbCuUpIDC, nil
}

func decodeE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpIDC *C.E2nodeComponentGNB_CU_UP_ID_t) (*e2ap_ies.E2NodeComponentGnbCuUpId, error) {

	var err error
	e2nodeComponentGnbCuUpID := e2ap_ies.E2NodeComponentGnbCuUpId{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//GNbCuUpId: gNbCuUpId,
	}

	e2nodeComponentGnbCuUpID.GNbCuUpId, err = decodeGnbCuUpID(&e2nodeComponentGnbCuUpIDC.gNB_CU_UP_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeGnbCuUpID() %s", err.Error())
	}

	return &e2nodeComponentGnbCuUpID, nil
}

func decodeE2nodeComponentGnbCuUpIDBytes(array [8]byte) (*e2ap_ies.E2NodeComponentGnbCuUpId, error) {
	e2nodeComponentGnbCuUpIDC := (*C.E2nodeComponentGNB_CU_UP_ID_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentGnbCuUpID(e2nodeComponentGnbCuUpIDC)
}
