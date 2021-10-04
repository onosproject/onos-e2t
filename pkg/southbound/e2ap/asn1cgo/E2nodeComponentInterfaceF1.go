// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceF1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceF1(e2nodeComponentInterfaceF1 *e2ap_ies.E2NodeComponentInterfaceF1) ([]byte, error) {
	e2nodeComponentInterfaceF1CP, err := newE2nodeComponentInterfaceF1(e2nodeComponentInterfaceF1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceF1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceF1, unsafe.Pointer(e2nodeComponentInterfaceF1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceF1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceF1(e2nodeComponentInterfaceF1 *e2ap_ies.E2NodeComponentInterfaceF1) ([]byte, error) {
	e2nodeComponentInterfaceF1CP, err := newE2nodeComponentInterfaceF1(e2nodeComponentInterfaceF1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceF1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceF1, unsafe.Pointer(e2nodeComponentInterfaceF1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceF1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceF1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceF1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceF1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceF1((*C.E2nodeComponentInterfaceF1_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceF1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceF1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceF1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceF1((*C.E2nodeComponentInterfaceF1_t)(unsafePtr))
}

func newE2nodeComponentInterfaceF1(e2nodeComponentInterfaceF1 *e2ap_ies.E2NodeComponentInterfaceF1) (*C.E2nodeComponentInterfaceF1_t, error) {

	gnbDuID, err := newGnbDuID(e2nodeComponentInterfaceF1.GetGNbDuId())
	if err != nil {
		return nil, err
	}
	e2nodeComponentInterfaceF1C := C.E2nodeComponentInterfaceF1_t{
		gNB_DU_ID: *gnbDuID,
	}

	return &e2nodeComponentInterfaceF1C, nil
}

func decodeE2nodeComponentInterfaceF1(e2nodeComponentInterfaceF1C *C.E2nodeComponentInterfaceF1_t) (*e2ap_ies.E2NodeComponentInterfaceF1, error) {

	gnbDuID, err := decodeGnbDuID(&e2nodeComponentInterfaceF1C.gNB_DU_ID)
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceF1 := e2ap_ies.E2NodeComponentInterfaceF1{
		GNbDuId:  gnbDuID,
	}

	return &e2nodeComponentInterfaceF1, nil
}

func decodeE2nodeComponentInterfaceF1Bytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceF1, error) {
	e2ncc := (*C.E2nodeComponentInterfaceF1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceF1(e2ncc)
}
