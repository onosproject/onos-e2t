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
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeResetRequest(rr *e2ap_pdu_contents.ResetRequest) ([]byte, error) {
	rrCP, err := newResetRequest(rr)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetRequest() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ResetRequest, unsafe.Pointer(rrCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeResetRequest() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeResetRequest(rr *e2ap_pdu_contents.ResetRequest) ([]byte, error) {
	rrCP, err := newResetRequest(rr)
	if err != nil {
		return nil, fmt.Errorf("perEncodeResetRequest() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ResetRequest, unsafe.Pointer(rrCP))
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

func newResetRequest(rr *e2ap_pdu_contents.ResetRequest) (*C.ResetRequest_t, error) {

	pIeC1710P20, err := newResetRequestIe(rr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rrC := C.ResetRequest_t{
		protocolIEs: *pIeC1710P20,
	}

	return &rrC, nil
}

func decodeResetRequest(rrC *C.ResetRequest_t) (*e2ap_pdu_contents.ResetRequest, error) {

	pIEs, err := decodeResetRequestIes(&rrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rr := e2ap_pdu_contents.ResetRequest{
		ProtocolIes: pIEs,
	}

	return &rr, nil
}
