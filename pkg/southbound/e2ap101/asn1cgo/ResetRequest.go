// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ResetRequest.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeResetRequest(resetRequest *e2ap_pdu_contents.ResetRequest) ([]byte, error) {
	resetRequestCP, err := newResetRequest(resetRequest)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetRequest() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ResetRequest, unsafe.Pointer(resetRequestCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetRequest() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeResetRequest(resetRequest *e2ap_pdu_contents.ResetRequest) ([]byte, error) {
	resetRequestCP, err := newResetRequest(resetRequest)
	if err != nil {
		return nil, fmt.Errorf("perEncodeResetRequest() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ResetRequest, unsafe.Pointer(resetRequestCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeResetRequest() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeResetRequest(bytes []byte) (*e2ap_pdu_contents.ResetRequest, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ResetRequest)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeResetRequest((*C.ResetRequest_t)(unsafePtr))
}

func perDecodeResetRequest(bytes []byte) (*e2ap_pdu_contents.ResetRequest, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ResetRequest)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeResetRequest((*C.ResetRequest_t)(unsafePtr))
}

func newResetRequest(resetRequest *e2ap_pdu_contents.ResetRequest) (*C.ResetRequest_t, error) {

	//var err error
	resetRequestC := C.ResetRequest_t{}

	//protocolIesC, err := newResetRequestIes(resetRequest.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newResetRequestIes() %s", err.Error())
	//}

	//resetRequestC.protocolIEs = protocolIesC

	return &resetRequestC, nil
}

func decodeResetRequest(resetRequestC *C.ResetRequest_t) (*e2ap_pdu_contents.ResetRequest, error) {

	//var err error
	resetRequest := e2ap_pdu_contents.ResetRequest{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//resetRequest.ProtocolIes, err = decodeResetRequestIes(resetRequestC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeResetRequestIes() %s", err.Error())
	//}

	return &resetRequest, nil
}

func decodeResetRequestBytes(array [8]byte) (*e2ap_pdu_contents.ResetRequest, error) {
	resetRequestC := (*C.ResetRequest_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeResetRequest(resetRequestC)
}
