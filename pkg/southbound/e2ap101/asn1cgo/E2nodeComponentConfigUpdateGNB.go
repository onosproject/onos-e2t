// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateGNB.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb *e2ap_ies.E2NodeComponentConfigUpdateGnb) ([]byte, error) {
	e2nodeComponentConfigUpdateGnbCP, err := newE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateGnb() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateGNB, unsafe.Pointer(e2nodeComponentConfigUpdateGnbCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateGnb() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb *e2ap_ies.E2NodeComponentConfigUpdateGnb) ([]byte, error) {
	e2nodeComponentConfigUpdateGnbCP, err := newE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateGnb() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateGNB, unsafe.Pointer(e2nodeComponentConfigUpdateGnbCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateGnb() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateGnb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateGNB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateGnb((*C.E2nodeComponentConfigUpdateGNB_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateGnb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateGNB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateGnb((*C.E2nodeComponentConfigUpdateGNB_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnb *e2ap_ies.E2NodeComponentConfigUpdateGnb) (*C.E2nodeComponentConfigUpdateGNB_t, error) {

	e2nodeComponentConfigUpdateGnbC := C.E2nodeComponentConfigUpdateGNB_t{}

	if e2nodeComponentConfigUpdateGnb.NgApconfigUpdate != "" {
		e2nodeComponentConfigUpdateGnbC.ngAPconfigUpdate = newPrintableString(e2nodeComponentConfigUpdateGnb.NgApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateGnb.XnApconfigUpdate != "" {
		e2nodeComponentConfigUpdateGnbC.xnAPconfigUpdate = newPrintableString(e2nodeComponentConfigUpdateGnb.XnApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateGnb.E1ApconfigUpdate != "" {
		e2nodeComponentConfigUpdateGnbC.e1APconfigUpdate = newPrintableString(e2nodeComponentConfigUpdateGnb.E1ApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateGnb.F1ApconfigUpdate != "" {
		e2nodeComponentConfigUpdateGnbC.f1APconfigUpdate = newPrintableString(e2nodeComponentConfigUpdateGnb.F1ApconfigUpdate)
	}

	return &e2nodeComponentConfigUpdateGnbC, nil
}

func decodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnbC *C.E2nodeComponentConfigUpdateGNB_t) (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {

	e2nodeComponentConfigUpdateGnb := e2ap_ies.E2NodeComponentConfigUpdateGnb{}

	if e2nodeComponentConfigUpdateGnbC.ngAPconfigUpdate != nil {
		e2nodeComponentConfigUpdateGnb.NgApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.ngAPconfigUpdate)
	}
	if e2nodeComponentConfigUpdateGnbC.xnAPconfigUpdate != nil {
		e2nodeComponentConfigUpdateGnb.XnApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.xnAPconfigUpdate)
	}
	if e2nodeComponentConfigUpdateGnbC.e1APconfigUpdate != nil {
		e2nodeComponentConfigUpdateGnb.E1ApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.e1APconfigUpdate)
	}
	if e2nodeComponentConfigUpdateGnbC.f1APconfigUpdate != nil {
		e2nodeComponentConfigUpdateGnb.F1ApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.f1APconfigUpdate)
	}

	return &e2nodeComponentConfigUpdateGnb, nil
}

func decodeE2nodeComponentConfigUpdateGnbBytes(array [8]byte) (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {
	e2nodeComponentConfigUpdateGnbC := (*C.E2nodeComponentConfigUpdateGNB_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnbC)
}
