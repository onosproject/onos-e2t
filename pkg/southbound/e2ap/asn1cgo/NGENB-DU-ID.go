// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "NGENB-DU-ID.h"
import "C"

import (
	"fmt"
	"unsafe"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
)

func xerEncodeNgEnbDuID(ngEnbDuID *e2apies.NgenbDuId) ([]byte, error) {
	ngEnbDuIDCP, err := newNgEnbDuID(ngEnbDuID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNgEnbDuID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_NGENB_DU_ID, unsafe.Pointer(ngEnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeNgEnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeNgEnbDuID(ngEnbDuID *e2apies.NgenbDuId) ([]byte, error) {
	ngEnbDuIDCP, err := newNgEnbDuID(ngEnbDuID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeNgEnbDuID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_NGENB_DU_ID, unsafe.Pointer(ngEnbDuIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeNgEnbDuID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeNgEnbDuID(bytes []byte) (*e2apies.NgenbDuId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_NGENB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeNgEnbDuID((*C.NGENB_DU_ID_t)(unsafePtr))
}

func perDecodeNgEnbDuID(bytes []byte) (*e2apies.NgenbDuId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_NGENB_DU_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeNgEnbDuID((*C.NGENB_DU_ID_t)(unsafePtr))
}

func newNgEnbDuID(ngEnbDuID *e2apies.NgenbDuId) (*C.NGENB_DU_ID_t, error) {

	ngEnbDuIDC, err := newInteger(ngEnbDuID.Value)
	if err != nil {
		return nil, fmt.Errorf("newInteger() %s", err.Error())
	}

	return ngEnbDuIDC, nil
}

func decodeNgEnbDuID(ngEnbDuIDC *C.NGENB_DU_ID_t) (*e2apies.NgenbDuId, error) {

	ngEnbDuID := new(e2apies.NgenbDuId)
	out, err := decodeInteger(ngEnbDuIDC)
	if err != nil {
		return nil, fmt.Errorf("decodeInteger() %s", err.Error())
	}
	ngEnbDuID.Value = out

	return ngEnbDuID, nil
}
