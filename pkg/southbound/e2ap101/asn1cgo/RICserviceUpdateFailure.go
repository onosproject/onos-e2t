// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceUpdateFailure.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicserviceUpdateFailure(ricserviceUpdateFailure *e2ap_pdu_contents.RicserviceUpdateFailure) ([]byte, error) {
	ricserviceUpdateFailureCP, err := newRicserviceUpdateFailure(ricserviceUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceUpdateFailure() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceUpdateFailure, unsafe.Pointer(ricserviceUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicserviceUpdateFailure(ricserviceUpdateFailure *e2ap_pdu_contents.RicserviceUpdateFailure) ([]byte, error) {
	ricserviceUpdateFailureCP, err := newRicserviceUpdateFailure(ricserviceUpdateFailure)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceUpdateFailure() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceUpdateFailure, unsafe.Pointer(ricserviceUpdateFailureCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicserviceUpdateFailure(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicserviceUpdateFailure((*C.RICserviceUpdateFailure_t)(unsafePtr))
}

func perDecodeRicserviceUpdateFailure(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicserviceUpdateFailure((*C.RICserviceUpdateFailure_t)(unsafePtr))
}

func newRicserviceUpdateFailure(ricserviceUpdateFailure *e2ap_pdu_contents.RicserviceUpdateFailure) (*C.RICserviceUpdateFailure_t, error) {

	//var err error
	ricserviceUpdateFailureC := C.RICserviceUpdateFailure_t{}

	//protocolIesC, err := newRicserviceUpdateFailureIes(ricserviceUpdateFailure.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newRicserviceUpdateFailureIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//ricserviceUpdateFailureC.protocolIEs = protocolIesC

	return &ricserviceUpdateFailureC, nil
}

func decodeRicserviceUpdateFailure(ricserviceUpdateFailureC *C.RICserviceUpdateFailure_t) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {

	//var err error
	ricserviceUpdateFailure := e2ap_pdu_contents.RicserviceUpdateFailure{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//ricserviceUpdateFailure.ProtocolIes, err = decodeRicserviceUpdateFailureIes(ricserviceUpdateFailureC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeRicserviceUpdateFailureIes() %s", err.Error())
	//}

	return &ricserviceUpdateFailure, nil
}

func decodeRicserviceUpdateFailureBytes(array [8]byte) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
	ricserviceUpdateFailureC := (*C.RICserviceUpdateFailure_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicserviceUpdateFailure(ricserviceUpdateFailureC)
}
