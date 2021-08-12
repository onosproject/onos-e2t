// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateNGeNB.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb *e2ap_ies.E2NodeComponentConfigUpdateNgeNb) ([]byte, error) {
	e2nodeComponentConfigUpdateNgeNbCP, err := newE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateNgeNb() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateNGeNB, unsafe.Pointer(e2nodeComponentConfigUpdateNgeNbCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateNgeNb() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb *e2ap_ies.E2NodeComponentConfigUpdateNgeNb) ([]byte, error) {
	e2nodeComponentConfigUpdateNgeNbCP, err := newE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateNgeNb() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateNGeNB, unsafe.Pointer(e2nodeComponentConfigUpdateNgeNbCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateNgeNb() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateNgeNb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateNgeNb, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateNGeNB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateNgeNb((*C.E2nodeComponentConfigUpdateNGeNB_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateNgeNb(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateNgeNb, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateNGeNB)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateNgeNb((*C.E2nodeComponentConfigUpdateNGeNB_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNb *e2ap_ies.E2NodeComponentConfigUpdateNgeNb) (*C.E2nodeComponentConfigUpdateNGeNB_t, error) {

	e2nodeComponentConfigUpdateNgeNbC := C.E2nodeComponentConfigUpdateNGeNB_t{}

	if e2nodeComponentConfigUpdateNgeNb.NgApconfigUpdate != "" {
		e2nodeComponentConfigUpdateNgeNbC.ngAPconfigUpdate = newPrintableString(e2nodeComponentConfigUpdateNgeNb.NgApconfigUpdate)
	}
	if e2nodeComponentConfigUpdateNgeNb.XnApconfigUpdate != "" {
		e2nodeComponentConfigUpdateNgeNbC.xnAPconfigUpdate = newPrintableString(e2nodeComponentConfigUpdateNgeNb.XnApconfigUpdate)
	}

	return &e2nodeComponentConfigUpdateNgeNbC, nil
}

func decodeE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNbC *C.E2nodeComponentConfigUpdateNGeNB_t) (*e2ap_ies.E2NodeComponentConfigUpdateNgeNb, error) {

	e2nodeComponentConfigUpdateNgeNb := e2ap_ies.E2NodeComponentConfigUpdateNgeNb{}

	if e2nodeComponentConfigUpdateNgeNbC.ngAPconfigUpdate != nil {
		e2nodeComponentConfigUpdateNgeNb.NgApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateNgeNbC.ngAPconfigUpdate)
	}
	if e2nodeComponentConfigUpdateNgeNbC.xnAPconfigUpdate != nil {
		e2nodeComponentConfigUpdateNgeNb.XnApconfigUpdate = decodePrintableString(e2nodeComponentConfigUpdateNgeNbC.xnAPconfigUpdate)
	}

	return &e2nodeComponentConfigUpdateNgeNb, nil
}

func decodeE2nodeComponentConfigUpdateNgeNbBytes(array [8]byte) (*e2ap_ies.E2NodeComponentConfigUpdateNgeNb, error) {
	e2nodeComponentConfigUpdateNgeNbC := (*C.E2nodeComponentConfigUpdateNGeNB_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateNgeNb(e2nodeComponentConfigUpdateNgeNbC)
}
