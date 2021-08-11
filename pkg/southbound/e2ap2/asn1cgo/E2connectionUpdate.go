// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdate.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionUpdate(e2connectionUpdate *e2ap_pdu_contents.E2ConnectionUpdate) ([]byte, error) {
	e2connectionUpdateCP, err := newE2connectionUpdate(e2connectionUpdate)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdate() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionUpdate, unsafe.Pointer(e2connectionUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdate() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionUpdate(e2connectionUpdate *e2ap_pdu_contents.E2ConnectionUpdate) ([]byte, error) {
	e2connectionUpdateCP, err := newE2connectionUpdate(e2connectionUpdate)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdate() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionUpdate, unsafe.Pointer(e2connectionUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdate() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionUpdate(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdate, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionUpdate((*C.E2connectionUpdate_t)(unsafePtr))
}

func perDecodeE2connectionUpdate(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdate, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionUpdate((*C.E2connectionUpdate_t)(unsafePtr))
}

func newE2connectionUpdate(e2cu *e2ap_pdu_contents.E2ConnectionUpdate) (*C.E2connectionUpdate_t, error) {

	pIeC1710P14, err := newE2connectionUpdateIe(e2cu.ProtocolIes)
	if err != nil {
		return nil, err
	}
	e2cuC := C.E2connectionUpdate_t{
		protocolIEs: *pIeC1710P14,
	}

	return &e2cuC, nil
}

func decodeE2connectionUpdate(e2cuC *C.E2connectionUpdate_t) (*e2ap_pdu_contents.E2ConnectionUpdate, error) {

	pIEs, err := decodeE2connectionUpdateIes(&e2cuC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2cu := e2ap_pdu_contents.E2ConnectionUpdate{
		ProtocolIes: pIEs,
	}

	return &e2cu, nil
}
