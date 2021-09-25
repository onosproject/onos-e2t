// SPDX-FileCopyrightText: 2020-present Open NetworkiX2 Foundation <info@opennetworkiX2.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceX2.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceX2(e2nodeComponentInterfaceX2 *e2ap_ies.E2NodeComponentInterfaceX2) ([]byte, error) {
	e2nodeComponentInterfaceX2CP, err := newE2nodeComponentInterfaceX2(e2nodeComponentInterfaceX2)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceX2() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceX2, unsafe.Pointer(e2nodeComponentInterfaceX2CP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceX2() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceX2(e2nodeComponentInterfaceX2 *e2ap_ies.E2NodeComponentInterfaceX2) ([]byte, error) {
	e2nodeComponentInterfaceX2CP, err := newE2nodeComponentInterfaceX2(e2nodeComponentInterfaceX2)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceX2() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceX2, unsafe.Pointer(e2nodeComponentInterfaceX2CP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceX2() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceX2(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceX2, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceX2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceX2((*C.E2nodeComponentInterfaceX2_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceX2(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceX2, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceX2)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceX2((*C.E2nodeComponentInterfaceX2_t)(unsafePtr))
}

func newE2nodeComponentInterfaceX2(e2nodeComponentInterfaceX2 *e2ap_ies.E2NodeComponentInterfaceX2) (*C.E2nodeComponentInterfaceX2_t, error) {

	enbIDC, err := newGlobaleNBID(e2nodeComponentInterfaceX2.GetGlobalENbId())
	if err != nil {
		return nil, err
	}
	enGnbIDC, err := newGlobalenGnbID(e2nodeComponentInterfaceX2.GetGlobalEnGNbId())
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceX2C := C.E2nodeComponentInterfaceX2_t{
		global_eNB_ID:    enbIDC,
		global_en_gNB_ID: enGnbIDC,
	}

	return &e2nodeComponentInterfaceX2C, nil
}

func decodeE2nodeComponentInterfaceX2(e2nodeComponentInterfaceX2C *C.E2nodeComponentInterfaceX2_t) (*e2ap_ies.E2NodeComponentInterfaceX2, error) {

	enbID, err := decodeGlobalEnbID(e2nodeComponentInterfaceX2C.global_eNB_ID)
	if err != nil {
		return nil, err
	}
	enGnbID, err := decodeGlobalenGnbID(e2nodeComponentInterfaceX2C.global_en_gNB_ID)

	e2nodeComponentInterfaceX2 := e2ap_ies.E2NodeComponentInterfaceX2{
		GlobalENbId:   enbID,
		GlobalEnGNbId: enGnbID,
	}

	return &e2nodeComponentInterfaceX2, nil
}

func decodeE2nodeComponentInterfaceX2Bytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceX2, error) {
	e2ncc := (*C.E2nodeComponentInterfaceX2_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceX2(e2ncc)
}
