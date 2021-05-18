// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdateFailure.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionUpdateFailure(e2connectionUpdateFailure *e2ap_pdu_contents.E2ConnectionUpdateFailure) ([]byte, error) {
	e2connectionUpdateFailureCP, err := newE2connectionUpdateFailure(e2connectionUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateFailure() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionUpdateFailure, unsafe.Pointer(e2connectionUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionUpdateFailure(e2connectionUpdateFailure *e2ap_pdu_contents.E2ConnectionUpdateFailure) ([]byte, error) {
	e2connectionUpdateFailureCP, err := newE2connectionUpdateFailure(e2connectionUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateFailure() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionUpdateFailure, unsafe.Pointer(e2connectionUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionUpdateFailure(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateFailure, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionUpdateFailure((*C.E2connectionUpdateFailure_t)(unsafePtr))
}

func perDecodeE2connectionUpdateFailure(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateFailure, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionUpdateFailure((*C.E2connectionUpdateFailure_t)(unsafePtr))
}

func newE2connectionUpdateFailure(e2cuf *e2ap_pdu_contents.E2ConnectionUpdateFailure) (*C.E2connectionUpdateFailure_t, error) {

	pIeC1710P16, err := newE2connectionUpdateFailureIe(e2cuf.ProtocolIes)
	if err != nil {
		return nil, err
	}
	e2cufC := C.E2connectionUpdateFailure_t{
		protocolIEs: *pIeC1710P16,
	}

	return &e2cufC, nil
}

func decodeE2connectionUpdateFailure(e2cufC *C.E2connectionUpdateFailure_t) (*e2ap_pdu_contents.E2ConnectionUpdateFailure, error) {

	pIEs, err := decodeE2connectionUpdateFailureIes(&e2cufC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2cuf := e2ap_pdu_contents.E2ConnectionUpdateFailure{
		ProtocolIes: pIEs,
	}

	return &e2cuf, nil
}
