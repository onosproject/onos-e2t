// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceE1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceE1(e2nodeComponentInterfaceE1 *e2ap_ies.E2NodeComponentInterfaceE1) ([]byte, error) {
	e2nodeComponentInterfaceE1CP, err := newE2nodeComponentInterfaceE1(e2nodeComponentInterfaceE1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceE1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceE1, unsafe.Pointer(e2nodeComponentInterfaceE1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceE1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceE1(e2nodeComponentInterfaceE1 *e2ap_ies.E2NodeComponentInterfaceE1) ([]byte, error) {
	e2nodeComponentInterfaceE1CP, err := newE2nodeComponentInterfaceE1(e2nodeComponentInterfaceE1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceE1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceE1, unsafe.Pointer(e2nodeComponentInterfaceE1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceE1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceE1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceE1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceE1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceE1((*C.E2nodeComponentInterfaceE1_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceE1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceE1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceE1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceE1((*C.E2nodeComponentInterfaceE1_t)(unsafePtr))
}

func newE2nodeComponentInterfaceE1(e2nodeComponentInterfaceE1 *e2ap_ies.E2NodeComponentInterfaceE1) (*C.E2nodeComponentInterfaceE1_t, error) {

	gnbCuUpID, err := newGnbCuUpID(e2nodeComponentInterfaceE1.GetGNbCuCpId())
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceE1C := C.E2nodeComponentInterfaceE1_t{
		gNB_CU_CP_ID: *gnbCuUpID,
	}

	return &e2nodeComponentInterfaceE1C, nil
}

func decodeE2nodeComponentInterfaceE1(e2nodeComponentInterfaceE1C *C.E2nodeComponentInterfaceE1_t) (*e2ap_ies.E2NodeComponentInterfaceE1, error) {

	gnbCuUpID, err := decodeGnbCuUpID(&e2nodeComponentInterfaceE1C.gNB_CU_CP_ID)
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceE1 := e2ap_ies.E2NodeComponentInterfaceE1{
		GNbCuCpId:  gnbCuUpID,
	}

	return &e2nodeComponentInterfaceE1, nil
}

func decodeE2nodeComponentInterfaceE1Bytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceE1, error) {
	e2ncc := (*C.E2nodeComponentInterfaceE1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceE1(e2ncc)
}
