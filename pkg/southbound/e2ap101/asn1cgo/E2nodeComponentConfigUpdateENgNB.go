// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateENgNB.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb *e2ap_ies.E2NodeComponentConfigUpdateEngNb) ([]byte, error) {
	e2nodeComponentConfigUpdateEngNbCP, err := newE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateEngNb() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateENgNB, unsafe.Pointer(e2nodeComponentConfigUpdateEngNbCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateEngNb() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb *e2ap_ies.E2NodeComponentConfigUpdateEngNb) ([]byte, error) {
	e2nodeComponentConfigUpdateEngNbCP, err := newE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateEngNb() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateENgNB, unsafe.Pointer(e2nodeComponentConfigUpdateEngNbCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateEngNb() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateEngNb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateEngNb, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateENgNB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateEngNb((*C.E2nodeComponentConfigUpdateENgNB_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateEngNb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateEngNb, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateENgNB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateEngNb((*C.E2nodeComponentConfigUpdateENgNB_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNb *e2ap_ies.E2NodeComponentConfigUpdateEngNb) (*C.E2nodeComponentConfigUpdateENgNB_t, error) {

	//var err error
	e2nodeComponentConfigUpdateEngNbC := C.E2nodeComponentConfigUpdateENgNB_t{}

	x2ApconfigUpdateC := newPrintableString(e2nodeComponentConfigUpdateEngNb.X2ApconfigUpdate)

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2nodeComponentConfigUpdateEngNbC.x2APconfigUpdate = x2ApconfigUpdateC

	return &e2nodeComponentConfigUpdateEngNbC, nil
}

func decodeE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNbC *C.E2nodeComponentConfigUpdateENgNB_t) (*e2ap_ies.E2NodeComponentConfigUpdateEngNb, error) {

	e2nodeComponentConfigUpdateEngNb := e2ap_ies.E2NodeComponentConfigUpdateEngNb{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//X2ApconfigUpdate: x2ApconfigUpdate,
	}

	e2nodeComponentConfigUpdateEngNb.X2ApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateEngNbC.x2APconfigUpdate)

	return &e2nodeComponentConfigUpdateEngNb, nil
}

func decodeE2nodeComponentConfigUpdateEngNbBytes(array [8]byte) (*e2ap_ies.E2NodeComponentConfigUpdateEngNb, error) {
	e2nodeComponentConfigUpdateEngNbC := (*C.E2nodeComponentConfigUpdateENgNB_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateEngNb(e2nodeComponentConfigUpdateEngNbC)
}
