// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeConfigurationUpdateFailure.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure *e2ap_pdu_contents.E2NodeConfigurationUpdateFailure) ([]byte, error) {
	e2nodeConfigurationUpdateFailureCP, err := newE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeConfigurationUpdateFailure, unsafe.Pointer(e2nodeConfigurationUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure *e2ap_pdu_contents.E2NodeConfigurationUpdateFailure) ([]byte, error) {
	e2nodeConfigurationUpdateFailureCP, err := newE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeConfigurationUpdateFailure, unsafe.Pointer(e2nodeConfigurationUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeConfigurationUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeConfigurationUpdateFailure(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeConfigurationUpdateFailure((*C.E2nodeConfigurationUpdateFailure_t)(unsafePtr))
}

func perDecodeE2nodeConfigurationUpdateFailure(bytes []byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeConfigurationUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeConfigurationUpdateFailure((*C.E2nodeConfigurationUpdateFailure_t)(unsafePtr))
}

func newE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailure *e2ap_pdu_contents.E2NodeConfigurationUpdateFailure) (*C.E2nodeConfigurationUpdateFailure_t, error) {

	//var err error
	e2nodeConfigurationUpdateFailureC := C.E2nodeConfigurationUpdateFailure_t{}

	//protocolIesC, err := newE2nodeConfigurationUpdateFailureIes(e2nodeConfigurationUpdateFailure.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newE2nodeConfigurationUpdateFailureIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//e2nodeConfigurationUpdateFailureC.protocolIEs = protocolIesC

	return &e2nodeConfigurationUpdateFailureC, nil
}

func decodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailureC *C.E2nodeConfigurationUpdateFailure_t) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {

	//var err error
	e2nodeConfigurationUpdateFailure := e2ap_pdu_contents.E2NodeConfigurationUpdateFailure{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//e2nodeConfigurationUpdateFailure.ProtocolIes, err = decodeE2nodeConfigurationUpdateFailureIes(e2nodeConfigurationUpdateFailureC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateFailureIes() %s", err.Error())
	//}

	return &e2nodeConfigurationUpdateFailure, nil
}

func decodeE2nodeConfigurationUpdateFailureBytes(array [8]byte) (*e2ap_pdu_contents.E2NodeConfigurationUpdateFailure, error) {
	e2nodeConfigurationUpdateFailureC := (*C.E2nodeConfigurationUpdateFailure_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeConfigurationUpdateFailure(e2nodeConfigurationUpdateFailureC)
}
