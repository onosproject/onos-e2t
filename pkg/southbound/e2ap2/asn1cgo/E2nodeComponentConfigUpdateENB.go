// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateENB.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb *e2ap_ies.E2NodeComponentConfigUpdateEnb) ([]byte, error) {
	e2nodeComponentConfigUpdateEnbCP, err := newE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateEnb() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateENB, unsafe.Pointer(e2nodeComponentConfigUpdateEnbCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateEnb() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb *e2ap_ies.E2NodeComponentConfigUpdateEnb) ([]byte, error) {
	e2nodeComponentConfigUpdateEnbCP, err := newE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateEnb() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateENB, unsafe.Pointer(e2nodeComponentConfigUpdateEnbCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateEnb() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateEnb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateEnb, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateENB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateEnb((*C.E2nodeComponentConfigUpdateENB_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateEnb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateEnb, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateENB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateEnb((*C.E2nodeComponentConfigUpdateENB_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnb *e2ap_ies.E2NodeComponentConfigUpdateEnb) (*C.E2nodeComponentConfigUpdateENB_t, error) {

	e2nodeComponentConfigUpdateEnbC := C.E2nodeComponentConfigUpdateENB_t{}

	if e2nodeComponentConfigUpdateEnb.NgApconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnbC.ngAPconfigUpdate = newOctetString(e2nodeComponentConfigUpdateEnb.NgApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnb.XnApconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnbC.xnAPconfigUpdate = newOctetString(e2nodeComponentConfigUpdateEnb.XnApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnb.W1ApconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnbC.w1APconfigUpdate = newOctetString(e2nodeComponentConfigUpdateEnb.W1ApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnb.S1ApconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnbC.s1APconfigUpdate = newOctetString(e2nodeComponentConfigUpdateEnb.S1ApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnb.X2ApconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnbC.x2APconfigUpdate = newOctetString(e2nodeComponentConfigUpdateEnb.X2ApconfigUpdate)
	}

	return &e2nodeComponentConfigUpdateEnbC, nil
}

func decodeE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnbC *C.E2nodeComponentConfigUpdateENB_t) (*e2ap_ies.E2NodeComponentConfigUpdateEnb, error) {

	e2nodeComponentConfigUpdateEnb := e2ap_ies.E2NodeComponentConfigUpdateEnb{}

	if e2nodeComponentConfigUpdateEnbC.ngAPconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnb.NgApconfigUpdate = decodeOctetString(e2nodeComponentConfigUpdateEnbC.ngAPconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnbC.xnAPconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnb.XnApconfigUpdate = decodeOctetString(e2nodeComponentConfigUpdateEnbC.xnAPconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnbC.w1APconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnb.W1ApconfigUpdate = decodeOctetString(e2nodeComponentConfigUpdateEnbC.w1APconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnbC.s1APconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnb.S1ApconfigUpdate = decodeOctetString(e2nodeComponentConfigUpdateEnbC.s1APconfigUpdate)
	}
	if e2nodeComponentConfigUpdateEnbC.x2APconfigUpdate != nil {
		e2nodeComponentConfigUpdateEnb.X2ApconfigUpdate = decodeOctetString(e2nodeComponentConfigUpdateEnbC.x2APconfigUpdate)
	}

	return &e2nodeComponentConfigUpdateEnb, nil
}

func decodeE2nodeComponentConfigUpdateEnbBytes(array [8]byte) (*e2ap_ies.E2NodeComponentConfigUpdateEnb, error) {
	e2nodeComponentConfigUpdateEnbC := (*C.E2nodeComponentConfigUpdateENB_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateEnb(e2nodeComponentConfigUpdateEnbC)
}
