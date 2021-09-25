// SPDX-FileCopyrightText: 2020-present Open NetworkiW1 Foundation <info@opennetworkiW1.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceW1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceW1(e2nodeComponentInterfaceW1 *e2ap_ies.E2NodeComponentInterfaceW1) ([]byte, error) {
	e2nodeComponentInterfaceW1CP, err := newE2nodeComponentInterfaceW1(e2nodeComponentInterfaceW1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceW1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceW1, unsafe.Pointer(e2nodeComponentInterfaceW1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceW1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceW1(e2nodeComponentInterfaceW1 *e2ap_ies.E2NodeComponentInterfaceW1) ([]byte, error) {
	e2nodeComponentInterfaceW1CP, err := newE2nodeComponentInterfaceW1(e2nodeComponentInterfaceW1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceW1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceW1, unsafe.Pointer(e2nodeComponentInterfaceW1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceW1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceW1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceW1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceW1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceW1((*C.E2nodeComponentInterfaceW1_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceW1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceW1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceW1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceW1((*C.E2nodeComponentInterfaceW1_t)(unsafePtr))
}

func newE2nodeComponentInterfaceW1(e2nodeComponentInterfaceW1 *e2ap_ies.E2NodeComponentInterfaceW1) (*C.E2nodeComponentInterfaceW1_t, error) {

	ngEnbDuIDC, err := newNgEnbDuID(e2nodeComponentInterfaceW1.GetNgENbDuId())
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceW1C := C.E2nodeComponentInterfaceW1_t{
		ng_eNB_DU_ID: ngEnbDuIDC,
	}

	return &e2nodeComponentInterfaceW1C, nil
}

func decodeE2nodeComponentInterfaceW1(e2nodeComponentInterfaceW1C *C.E2nodeComponentInterfaceW1_t) (*e2ap_ies.E2NodeComponentInterfaceW1, error) {

	ngEnbDuID, err := decodeNgEnbDuID(e2nodeComponentInterfaceW1C.ng_eNB_DU_ID)
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceW1 := e2ap_ies.E2NodeComponentInterfaceW1{
		NgENbDuId:  ngEnbDuID,
	}

	return &e2nodeComponentInterfaceW1, nil
}

func decodeE2nodeComponentInterfaceW1Bytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceW1, error) {
	e2ncc := (*C.E2nodeComponentInterfaceW1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceW1(e2ncc)
}
