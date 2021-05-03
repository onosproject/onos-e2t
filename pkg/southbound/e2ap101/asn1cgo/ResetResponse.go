// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ResetResponse.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeResetResponse(resetResponse *e2ap_pdu_contents.ResetResponse) ([]byte, error) {
	resetResponseCP, err := newResetResponse(resetResponse)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetResponse() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ResetResponse, unsafe.Pointer(resetResponseCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetResponse() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeResetResponse(resetResponse *e2ap_pdu_contents.ResetResponse) ([]byte, error) {
	resetResponseCP, err := newResetResponse(resetResponse)
	if err != nil {
		return nil, fmt.Errorf("perEncodeResetResponse() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ResetResponse, unsafe.Pointer(resetResponseCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeResetResponse() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeResetResponse(bytes []byte) (*e2ap_pdu_contents.ResetResponse, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ResetResponse)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeResetResponse((*C.ResetResponse_t)(unsafePtr))
}

func perDecodeResetResponse(bytes []byte) (*e2ap_pdu_contents.ResetResponse, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ResetResponse)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeResetResponse((*C.ResetResponse_t)(unsafePtr))
}

func newResetResponse(resetResponse *e2ap_pdu_contents.ResetResponse) (*C.ResetResponse_t, error) {

	//var err error
	resetResponseC := C.ResetResponse_t{}

	//protocolIesC, err := newResetResponseIes(resetResponse.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newResetResponseIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//resetResponseC.protocolIEs = protocolIesC

	return &resetResponseC, nil
}

func decodeResetResponse(resetResponseC *C.ResetResponse_t) (*e2ap_pdu_contents.ResetResponse, error) {

	//var err error
	resetResponse := e2ap_pdu_contents.ResetResponse{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,

	}

	//resetResponse.ProtocolIes, err = decodeResetResponseIes(resetResponseC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeResetResponseIes() %s", err.Error())
	//}

	return &resetResponse, nil
}

func decodeResetResponseBytes(array [8]byte) (*e2ap_pdu_contents.ResetResponse, error) {
	resetResponseC := (*C.ResetResponse_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeResetResponse(resetResponseC)
}
