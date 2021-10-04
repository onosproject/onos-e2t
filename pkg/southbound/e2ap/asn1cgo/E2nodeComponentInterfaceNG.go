// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceNG.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceNG(e2nodeComponentInterfaceNG *e2ap_ies.E2NodeComponentInterfaceNg) ([]byte, error) {
	e2nodeComponentInterfaceNGCP, err := newE2nodeComponentInterfaceNG(e2nodeComponentInterfaceNG)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceNG() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceNG, unsafe.Pointer(e2nodeComponentInterfaceNGCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceNG() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceNG(e2nodeComponentInterfaceNG *e2ap_ies.E2NodeComponentInterfaceNg) ([]byte, error) {
	e2nodeComponentInterfaceNGCP, err := newE2nodeComponentInterfaceNG(e2nodeComponentInterfaceNG)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceNG() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceNG, unsafe.Pointer(e2nodeComponentInterfaceNGCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceNG() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceNG(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceNg, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceNG)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceNG((*C.E2nodeComponentInterfaceNG_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceNG(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceNg, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceNG)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceNG((*C.E2nodeComponentInterfaceNG_t)(unsafePtr))
}

func newE2nodeComponentInterfaceNG(e2nodeComponentInterfaceNG *e2ap_ies.E2NodeComponentInterfaceNg) (*C.E2nodeComponentInterfaceNG_t, error) {

	e2nodeComponentInterfaceNGC := C.E2nodeComponentInterfaceNG_t{
		amf_name: *newAmfName(e2nodeComponentInterfaceNG.GetAmfName()),
	}

	return &e2nodeComponentInterfaceNGC, nil
}

func decodeE2nodeComponentInterfaceNG(e2nodeComponentInterfaceNGC *C.E2nodeComponentInterfaceNG_t) (*e2ap_ies.E2NodeComponentInterfaceNg, error) {

	e2nodeComponentInterfaceNG := e2ap_ies.E2NodeComponentInterfaceNg{
		AmfName:  decodeAmfName(&e2nodeComponentInterfaceNGC.amf_name),
	}

	return &e2nodeComponentInterfaceNG, nil
}

func decodeE2nodeComponentInterfaceNGBytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceNg, error) {
	e2ncc := (*C.E2nodeComponentInterfaceNG_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceNG(e2ncc)
}
