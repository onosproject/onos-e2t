// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdateAcknowledge.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge *e2ap_pdu_contents.E2ConnectionUpdateAcknowledge) ([]byte, error) {
	e2connectionUpdateAcknowledgeCP, err := newE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionUpdateAcknowledge, unsafe.Pointer(e2connectionUpdateAcknowledgeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge *e2ap_pdu_contents.E2ConnectionUpdateAcknowledge) ([]byte, error) {
	e2connectionUpdateAcknowledgeCP, err := newE2connectionUpdateAcknowledge(e2connectionUpdateAcknowledge)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionUpdateAcknowledge, unsafe.Pointer(e2connectionUpdateAcknowledgeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateAcknowledge, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionUpdateAcknowledge((*C.E2connectionUpdateAcknowledge_t)(unsafePtr))
}

func perDecodeE2connectionUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateAcknowledge, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionUpdateAcknowledge((*C.E2connectionUpdateAcknowledge_t)(unsafePtr))
}

func newE2connectionUpdateAcknowledge(e2cua *e2ap_pdu_contents.E2ConnectionUpdateAcknowledge) (*C.E2connectionUpdateAcknowledge_t, error) {

	pIeC1710P15, err := newE2connectionUpdateAcknowledgeIe(e2cua.ProtocolIes)
	if err != nil {
		return nil, err
	}
	e2cuaC := C.E2connectionUpdateAcknowledge_t{
		protocolIEs: *pIeC1710P15,
	}

	return &e2cuaC, nil
}

func decodeE2connectionUpdateAcknowledge(e2cuaC *C.E2connectionUpdateAcknowledge_t) (*e2ap_pdu_contents.E2ConnectionUpdateAcknowledge, error) {

	pIEs, err := decodeE2connectionUpdateAcknowledgeIes(&e2cuaC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2cua := e2ap_pdu_contents.E2ConnectionUpdateAcknowledge{
		ProtocolIes: pIEs,
	}

	return &e2cua, nil
}
