// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeConfigurationUpdateAcknowledge.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge *e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge) ([]byte, error) {
	e2nodeConfigurationUpdateAcknowledgeCP, err := newE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeConfigurationUpdateAcknowledge, unsafe.Pointer(e2nodeConfigurationUpdateAcknowledgeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge *e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge) ([]byte, error) {
	e2nodeConfigurationUpdateAcknowledgeCP, err := newE2nodeConfigurationUpdateAcknowledge(e2nodeConfigurationUpdateAcknowledge)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeConfigurationUpdateAcknowledge, unsafe.Pointer(e2nodeConfigurationUpdateAcknowledgeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeConfigurationUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeConfigurationUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeConfigurationUpdateAcknowledge((*C.E2nodeConfigurationUpdateAcknowledge_t)(unsafePtr))
}

func perDecodeE2nodeConfigurationUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeConfigurationUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeConfigurationUpdateAcknowledge((*C.E2nodeConfigurationUpdateAcknowledge_t)(unsafePtr))
}

func newE2nodeConfigurationUpdateAcknowledge(e2cua *e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge) (*C.E2nodeConfigurationUpdateAcknowledge_t, error) {

	pIeC1710P18, err := newE2nodeConfigurationUpdateAcknowledgeIe(e2cua.ProtocolIes)
	if err != nil {
		return nil, err
	}
	e2cuaC := C.E2nodeConfigurationUpdateAcknowledge_t{
		protocolIEs: *pIeC1710P18,
	}

	return &e2cuaC, nil
}

func decodeE2nodeConfigurationUpdateAcknowledge(e2cuaC *C.E2nodeConfigurationUpdateAcknowledge_t) (*e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge, error) {

	pIEs, err := decodeE2nodeConfigurationUpdateAcknowledgeIes(&e2cuaC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2cua := e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledge{
		ProtocolIes: pIEs,
	}

	return &e2cua, nil
}
