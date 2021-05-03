// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateAck.h" //ToDo - if there is an anonymous C-struct option, it would require linking additional C-struct file definition (the one above or before)
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck *e2ap_ies.E2NodeComponentConfigUpdateAck) ([]byte, error) {
	e2nodeComponentConfigUpdateAckCP, err := newE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAck() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateAck, unsafe.Pointer(e2nodeComponentConfigUpdateAckCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAck() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck *e2ap_ies.E2NodeComponentConfigUpdateAck) ([]byte, error) {
	e2nodeComponentConfigUpdateAckCP, err := newE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAck() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateAck, unsafe.Pointer(e2nodeComponentConfigUpdateAckCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAck() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateAck(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateAck, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAck((*C.E2nodeComponentConfigUpdateAck_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateAck(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdateAck, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAck((*C.E2nodeComponentConfigUpdateAck_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAck *e2ap_ies.E2NodeComponentConfigUpdateAck) (*C.E2nodeComponentConfigUpdateAck_t, error) {

	var err error
	e2nodeComponentConfigUpdateAckC := C.E2nodeComponentConfigUpdateAck_t{}

	updateOutcomeC := C.long(e2nodeComponentConfigUpdateAck.UpdateOutcome)
	failureCauseC, err := newCause(e2nodeComponentConfigUpdateAck.FailureCause)
	if err != nil {
		return nil, fmt.Errorf("newCause() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2nodeComponentConfigUpdateAckC.updateOutcome = updateOutcomeC
	e2nodeComponentConfigUpdateAckC.failureCause = failureCauseC

	return &e2nodeComponentConfigUpdateAckC, nil
}

func decodeE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAckC *C.E2nodeComponentConfigUpdateAck_t) (*e2ap_ies.E2NodeComponentConfigUpdateAck, error) {

	var err error
	e2nodeComponentConfigUpdateAck := e2ap_ies.E2NodeComponentConfigUpdateAck{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//UpdateOutcome: updateOutcome,
		//FailureCause: failureCause,
	}

	e2nodeComponentConfigUpdateAck.UpdateOutcome = int32(e2nodeComponentConfigUpdateAckC.updateOutcome)
	e2nodeComponentConfigUpdateAck.FailureCause, err = decodeCause(e2nodeComponentConfigUpdateAckC.failureCause)
	if err != nil {
		return nil, fmt.Errorf("decodeCause() %s", err.Error())
	}

	return &e2nodeComponentConfigUpdateAck, nil
}

func decodeE2nodeComponentConfigUpdateAckBytes(array [8]byte) (*e2ap_ies.E2NodeComponentConfigUpdateAck, error) {
	e2nodeComponentConfigUpdateAckC := (*C.E2nodeComponentConfigUpdateAck_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeComponentConfigUpdateAck(e2nodeComponentConfigUpdateAckC)
}
