// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeConfigurationUpdate.h"
import "C"

import (
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate *e2ap_pdu_contents.E2NodeConfigurationUpdate) ([]byte, error) {
	e2nodeConfigurationUpdateCP, err := newE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdate() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeConfigurationUpdate, unsafe.Pointer(e2nodeConfigurationUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdate() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeConfigurationUpdate(e2nodeConfigurationUpdate *e2ap_pdu_contents.E2NodeConfigurationUpdate) ([]byte, error) {
	e2nodeConfigurationUpdateCP, err := newE2nodeConfigurationUpdate(e2nodeConfigurationUpdate)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdate() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeConfigurationUpdate, unsafe.Pointer(e2nodeConfigurationUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdate() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeConfigurationUpdate(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeConfigurationUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeConfigurationUpdate((*C.E2nodeConfigurationUpdate_t)(unsafePtr))
}

func perDecodeE2nodeConfigurationUpdate(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeConfigurationUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeConfigurationUpdate((*C.E2nodeConfigurationUpdate_t)(unsafePtr))
}

func newE2nodeConfigurationUpdate(e2cu *e2ap_pdu_contents.E2NodeConfigurationUpdate) (*C.E2nodeConfigurationUpdate_t, error) {

	pIeC1710P17, err := newE2nodeConfigurationUpdateIe(e2cu.ProtocolIes)
	if err != nil {
		return nil, err
	}
	e2cuC := C.E2nodeConfigurationUpdate_t{
		protocolIEs: *pIeC1710P17,
	}

	return &e2cuC, nil
}

func decodeE2nodeConfigurationUpdate(e2cuC *C.E2nodeConfigurationUpdate_t) (*e2ap_pdu_contents.E2NodeConfigurationUpdate, error) {

	pIEs, err := decodeE2nodeConfigurationUpdateIes(&e2cuC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2cu := e2ap_pdu_contents.E2NodeConfigurationUpdate{
		ProtocolIes: pIEs,
	}

	return &e2cu, nil
}
