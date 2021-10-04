// SPDX-FileCopyrightText: 2020-present Open NetworkiS1 Foundation <info@opennetworkiS1.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceS1.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceS1(e2nodeComponentInterfaceS1 *e2ap_ies.E2NodeComponentInterfaceS1) ([]byte, error) {
	e2nodeComponentInterfaceS1CP, err := newE2nodeComponentInterfaceS1(e2nodeComponentInterfaceS1)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceS1() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceS1, unsafe.Pointer(e2nodeComponentInterfaceS1CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceS1() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceS1(e2nodeComponentInterfaceS1 *e2ap_ies.E2NodeComponentInterfaceS1) ([]byte, error) {
	e2nodeComponentInterfaceS1CP, err := newE2nodeComponentInterfaceS1(e2nodeComponentInterfaceS1)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceS1() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceS1, unsafe.Pointer(e2nodeComponentInterfaceS1CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceS1() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceS1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceS1, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceS1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceS1((*C.E2nodeComponentInterfaceS1_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceS1(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceS1, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceS1)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceS1((*C.E2nodeComponentInterfaceS1_t)(unsafePtr))
}

func newE2nodeComponentInterfaceS1(e2nodeComponentInterfaceS1 *e2ap_ies.E2NodeComponentInterfaceS1) (*C.E2nodeComponentInterfaceS1_t, error) {

	e2nodeComponentInterfaceS1C := C.E2nodeComponentInterfaceS1_t{
		mme_name: *newMmeName(e2nodeComponentInterfaceS1.GetMmeName()),
	}

	return &e2nodeComponentInterfaceS1C, nil
}

func decodeE2nodeComponentInterfaceS1(e2nodeComponentInterfaceS1C *C.E2nodeComponentInterfaceS1_t) (*e2ap_ies.E2NodeComponentInterfaceS1, error) {

	e2nodeComponentInterfaceS1 := e2ap_ies.E2NodeComponentInterfaceS1{
		MmeName:  decodeMmeName(&e2nodeComponentInterfaceS1C.mme_name),
	}

	return &e2nodeComponentInterfaceS1, nil
}

func decodeE2nodeComponentInterfaceS1Bytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceS1, error) {
	e2ncc := (*C.E2nodeComponentInterfaceS1_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceS1(e2ncc)
}
