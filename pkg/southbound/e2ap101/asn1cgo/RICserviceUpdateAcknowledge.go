// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceUpdateAcknowledge.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge *e2ap_pdu_contents.RicserviceUpdateAcknowledge) ([]byte, error) {
	ricserviceUpdateAcknowledgeCP, err := newRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceUpdateAcknowledge, unsafe.Pointer(ricserviceUpdateAcknowledgeCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge *e2ap_pdu_contents.RicserviceUpdateAcknowledge) ([]byte, error) {
	ricserviceUpdateAcknowledgeCP, err := newRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceUpdateAcknowledge() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceUpdateAcknowledge, unsafe.Pointer(ricserviceUpdateAcknowledgeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceUpdateAcknowledge() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicserviceUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicserviceUpdateAcknowledge((*C.RICserviceUpdateAcknowledge_t)(unsafePtr))
}

func perDecodeRicserviceUpdateAcknowledge(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceUpdateAcknowledge)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicserviceUpdateAcknowledge((*C.RICserviceUpdateAcknowledge_t)(unsafePtr))
}

func newRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledge *e2ap_pdu_contents.RicserviceUpdateAcknowledge) (*C.RICserviceUpdateAcknowledge_t, error) {

	//var err error
	ricserviceUpdateAcknowledgeC := C.RICserviceUpdateAcknowledge_t{}

	//protocolIesC, err := newRicserviceUpdateAcknowledgeIes(ricserviceUpdateAcknowledge.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newRicserviceUpdateAcknowledgeIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//ricserviceUpdateAcknowledgeC.protocolIEs = protocolIesC

	return &ricserviceUpdateAcknowledgeC, nil
}

func decodeRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledgeC *C.RICserviceUpdateAcknowledge_t) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {

	//var err error
	ricserviceUpdateAcknowledge := e2ap_pdu_contents.RicserviceUpdateAcknowledge{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//ricserviceUpdateAcknowledge.ProtocolIes, err = decodeRicserviceUpdateAcknowledgeIes(ricserviceUpdateAcknowledgeC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeRicserviceUpdateAcknowledgeIes() %s", err.Error())
	//}

	return &ricserviceUpdateAcknowledge, nil
}

func decodeRicserviceUpdateAcknowledgeBytes(array [8]byte) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {
	ricserviceUpdateAcknowledgeC := (*C.RICserviceUpdateAcknowledge_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicserviceUpdateAcknowledge(ricserviceUpdateAcknowledgeC)
}
