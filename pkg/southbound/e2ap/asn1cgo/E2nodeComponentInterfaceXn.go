// SPDX-FileCopyrightText: 2020-present Open NetworkiXn Foundation <info@opennetworkiXn.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceXn.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceXn(e2nodeComponentInterfaceXn *e2ap_ies.E2NodeComponentInterfaceXn) ([]byte, error) {
	e2nodeComponentInterfaceXnCP, err := newE2nodeComponentInterfaceXn(e2nodeComponentInterfaceXn)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceXn() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceXn, unsafe.Pointer(e2nodeComponentInterfaceXnCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceXn() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceXn(e2nodeComponentInterfaceXn *e2ap_ies.E2NodeComponentInterfaceXn) ([]byte, error) {
	e2nodeComponentInterfaceXnCP, err := newE2nodeComponentInterfaceXn(e2nodeComponentInterfaceXn)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceXn() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceXn, unsafe.Pointer(e2nodeComponentInterfaceXnCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceXn() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceXn(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceXn, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceXn)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceXn((*C.E2nodeComponentInterfaceXn_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceXn(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceXn, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceXn)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceXn((*C.E2nodeComponentInterfaceXn_t)(unsafePtr))
}

func newE2nodeComponentInterfaceXn(e2nodeComponentInterfaceXn *e2ap_ies.E2NodeComponentInterfaceXn) (*C.E2nodeComponentInterfaceXn_t, error) {

	gNgRanNodeIDC, err := newGlobalNgRanNodeID(e2nodeComponentInterfaceXn.GetGlobalNgRanNodeId())
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceXnC := C.E2nodeComponentInterfaceXn_t{
		global_NG_RAN_Node_ID: *gNgRanNodeIDC,
	}

	return &e2nodeComponentInterfaceXnC, nil
}

func decodeE2nodeComponentInterfaceXn(e2nodeComponentInterfaceXnC *C.E2nodeComponentInterfaceXn_t) (*e2ap_ies.E2NodeComponentInterfaceXn, error) {

	gNgRanNodeID, err := decodeGlobalNgRanNodeID(&e2nodeComponentInterfaceXnC.global_NG_RAN_Node_ID)
	if err != nil {
		return nil, err
	}

	e2nodeComponentInterfaceXn := e2ap_ies.E2NodeComponentInterfaceXn{
		GlobalNgRanNodeId: gNgRanNodeID,
	}

	return &e2nodeComponentInterfaceXn, nil
}

func decodeE2nodeComponentInterfaceXnBytes(array [8]byte) (*e2ap_ies.E2NodeComponentInterfaceXn, error) {
	e2ncc := (*C.E2nodeComponentInterfaceXn_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentInterfaceXn(e2ncc)
}
