// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentInterfaceType.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentInterfaceType(E2nodeComponentInterfaceType *e2ap_ies.E2NodeComponentInterfaceType) ([]byte, error) {
	E2nodeComponentInterfaceTypeCP, err := newE2nodeComponentInterfaceType(E2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceType() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentInterfaceType, unsafe.Pointer(E2nodeComponentInterfaceTypeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentInterfaceType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentInterfaceType(E2nodeComponentInterfaceType *e2ap_ies.E2NodeComponentInterfaceType) ([]byte, error) {
	E2nodeComponentInterfaceTypeCP, err := newE2nodeComponentInterfaceType(E2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceType() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentInterfaceType, unsafe.Pointer(E2nodeComponentInterfaceTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentInterfaceType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentInterfaceType(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentInterfaceType)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentInterfaceType((*C.E2nodeComponentInterfaceType_t)(unsafePtr))
}

func perDecodeE2nodeComponentInterfaceType(bytes []byte) (*e2ap_ies.E2NodeComponentInterfaceType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentInterfaceType)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentInterfaceType((*C.E2nodeComponentInterfaceType_t)(unsafePtr))
}

func newE2nodeComponentInterfaceType(e2nodeComponentInterfaceType *e2ap_ies.E2NodeComponentInterfaceType) (*C.E2nodeComponentInterfaceType_t, error) {

	var itC C.E2nodeComponentInterfaceType_t
	switch *e2nodeComponentInterfaceType {
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_NG:
		itC = C.E2nodeComponentInterfaceType_ng
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_XN:
		itC = C.E2nodeComponentInterfaceType_xn
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1:
		itC = C.E2nodeComponentInterfaceType_e1
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_F1:
		itC = C.E2nodeComponentInterfaceType_f1
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_W1:
		itC = C.E2nodeComponentInterfaceType_w1
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_S1:
		itC = C.E2nodeComponentInterfaceType_s1
	case e2ap_ies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_X2:
		itC = C.E2nodeComponentInterfaceType_x2
	default:
		return nil, fmt.Errorf("unexpected E2nodeComponentInterfaceType %v", *e2nodeComponentInterfaceType)
	}

	return &itC, nil
}

func decodeE2nodeComponentInterfaceType(e2nodeComponentInterfaceTypeC *C.E2nodeComponentInterfaceType_t) (*e2ap_ies.E2NodeComponentInterfaceType, error) {

	e2ncit := e2ap_ies.E2NodeComponentInterfaceType(int32(*e2nodeComponentInterfaceTypeC))

	return &e2ncit, nil
}
