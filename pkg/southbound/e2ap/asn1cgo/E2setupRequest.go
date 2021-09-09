// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2setupRequest.h"
import "C"
import (
	"fmt"
	"unsafe"

	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
)

func xerEncodeE2SetupRequest(e2SetupRequest *e2appducontents.E2SetupRequest) ([]byte, error) {
	e2SetupRequestC, err := newE2SetupRequest(e2SetupRequest)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2setupRequest, unsafe.Pointer(e2SetupRequestC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perEncodeE2SetupRequest(e2SetupRequest *e2appducontents.E2SetupRequest) ([]byte, error) {
	e2SetupRequestC, err := newE2SetupRequest(e2SetupRequest)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2setupRequest, unsafe.Pointer(e2SetupRequestC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func xerDecodeE2SetupRequest(bytes []byte) (*e2appducontents.E2SetupRequest, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2setupRequest)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2setupRequest((*C.E2setupRequest_t)(unsafePtr))
}

func perDecodeE2SetupRequest(bytes []byte) (*e2appducontents.E2SetupRequest, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2setupRequest)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2setupRequest((*C.E2setupRequest_t)(unsafePtr))
}

func newE2SetupRequest(esr *e2appducontents.E2SetupRequest) (*C.E2setupRequest_t, error) {
	pIeC1751P11, err := newE2SetupRequestIes(esr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	esC := C.E2setupRequest_t{
		protocolIEs: *pIeC1751P11,
	}

	return &esC, nil
}

func decodeE2setupRequest(e2setupRequestC *C.E2setupRequest_t) (*e2appducontents.E2SetupRequest, error) {
	pIEs, err := decodeE2SetupRequestIes(&e2setupRequestC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2setupRequest := e2appducontents.E2SetupRequest{
		ProtocolIes: pIEs,
	}
	return &e2setupRequest, nil
}
