// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateGNB.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
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

	ngApconfigUpdateC := newPrintableString(e2nodeComponentConfigUpdateGnb.NgApconfigUpdate)
	xnApconfigUpdateC := newPrintableString(e2nodeComponentConfigUpdateGnb.XnApconfigUpdate)
	e1ApconfigUpdateC := newPrintableString(e2nodeComponentConfigUpdateGnb.E1ApconfigUpdate)
	f1ApconfigUpdateC := newPrintableString(e2nodeComponentConfigUpdateGnb.F1ApconfigUpdate)

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2nodeComponentConfigUpdateGnbC.ngAPconfigUpdate = ngApconfigUpdateC
	e2nodeComponentConfigUpdateGnbC.xnAPconfigUpdate = xnApconfigUpdateC
	e2nodeComponentConfigUpdateGnbC.e1APconfigUpdate = e1ApconfigUpdateC
	e2nodeComponentConfigUpdateGnbC.f1APconfigUpdate = f1ApconfigUpdateC

	return &e2nodeComponentConfigUpdateGnbC, nil
}

func decodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnbC *C.E2nodeComponentConfigUpdateGNB_t) (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {

	e2nodeComponentConfigUpdateGnb := e2ap_ies.E2NodeComponentConfigUpdateGnb{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//NgApconfigUpdate: ngApconfigUpdate,
		//XnApconfigUpdate: xnApconfigUpdate,
		//E1ApconfigUpdate: e1ApconfigUpdate,
		//F1ApconfigUpdate: f1ApconfigUpdate,
	}

	e2nodeComponentConfigUpdateGnb.NgApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.ngAPconfigUpdate)
	e2nodeComponentConfigUpdateGnb.XnApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.xnAPconfigUpdate)
	e2nodeComponentConfigUpdateGnb.E1ApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.e1APconfigUpdate)
	e2nodeComponentConfigUpdateGnb.F1ApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateGnbC.f1APconfigUpdate)

	return &e2nodeComponentConfigUpdateGnb, nil
}

func decodeE2nodeComponentConfigUpdateGnbBytes(array [8]byte) (*e2ap_ies.E2NodeComponentConfigUpdateGnb, error) {
	e2nodeComponentConfigUpdateGnbC := (*C.E2nodeComponentConfigUpdateGNB_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateGnb(e2nodeComponentConfigUpdateGnbC)
}
