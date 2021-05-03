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
	"encoding/binary"
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

func newE2connectionUpdate(e2connectionUpdate *e2ap_pdu_contents.E2ConnectionUpdate) (*C.E2connectionUpdate_t, error) {

	//var err error
	e2connectionUpdateC := C.E2connectionUpdate_t{}

	//protocolIesC, err := newE2connectionUpdateIes(e2connectionUpdate.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newE2connectionUpdateIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//e2connectionUpdateC.protocolIEs = protocolIesC

	return &e2connectionUpdateC, nil
}

func decodeE2connectionUpdate(e2connectionUpdateC *C.E2connectionUpdate_t) (*e2ap_pdu_contents.E2ConnectionUpdate, error) {

	//var err error
	e2connectionUpdate := e2ap_pdu_contents.E2ConnectionUpdate{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//e2connectionUpdate.ProtocolIes, err = decodeE2connectionUpdateIes(e2connectionUpdateC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeE2connectionUpdateIes() %s", err.Error())
	//}

	return &e2connectionUpdate, nil
}

func decodeE2connectionUpdateBytes(array [8]byte) (*e2ap_pdu_contents.E2ConnectionUpdate, error) {
	e2connectionUpdateC := (*C.E2connectionUpdate_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2connectionUpdate(e2connectionUpdateC)
}
