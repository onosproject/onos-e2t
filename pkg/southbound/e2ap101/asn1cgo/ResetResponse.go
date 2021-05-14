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
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeResetResponse(rr *e2ap_pdu_contents.ResetResponse) ([]byte, error) {
	rrCP, err := newResetResponse(rr)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetResponse() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ResetResponse, unsafe.Pointer(rrCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetResponse() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeResetResponse(rr *e2ap_pdu_contents.ResetResponse) ([]byte, error) {
	rrCP, err := newResetResponse(rr)
	if err != nil {
		return nil, fmt.Errorf("perEncodeResetResponse() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ResetResponse, unsafe.Pointer(rrCP))
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

func newResetResponse(rr *e2ap_pdu_contents.ResetResponse) (*C.ResetResponse_t, error) {

	pIeC1710P21, err := newResetResponseIe(rr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rrC := C.ResetResponse_t{
		protocolIEs: *pIeC1710P21,
	}

	return &rrC, nil
}

func decodeResetResponse(rrC *C.ResetResponse_t) (*e2ap_pdu_contents.ResetResponse, error) {

	pIEs, err := decodeResetResponseIes(&rrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rr := e2ap_pdu_contents.ResetResponse{
		ProtocolIes: pIEs,
	}

	return &rr, nil
}
