// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceUpdate.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicserviceUpdate(ricserviceUpdate *e2ap_pdu_contents.RicserviceUpdate) ([]byte, error) {
	ricserviceUpdateCP, err := newRicserviceUpdate(ricserviceUpdate)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceUpdate() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceUpdate, unsafe.Pointer(ricserviceUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceUpdate() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicserviceUpdate(ricserviceUpdate *e2ap_pdu_contents.RicserviceUpdate) ([]byte, error) {
	ricserviceUpdateCP, err := newRicserviceUpdate(ricserviceUpdate)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceUpdate() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceUpdate, unsafe.Pointer(ricserviceUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceUpdate() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicserviceUpdate(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdate, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicserviceUpdate((*C.RICserviceUpdate_t)(unsafePtr))
}

func perDecodeRicserviceUpdate(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdate, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicserviceUpdate((*C.RICserviceUpdate_t)(unsafePtr))
}

func newRicserviceUpdate(ricserviceUpdate *e2ap_pdu_contents.RicserviceUpdate) (*C.RICserviceUpdate_t, error) {

	//var err error
	ricserviceUpdateC := C.RICserviceUpdate_t{}

	//protocolIesC, err := newRicserviceUpdateIes(ricserviceUpdate.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newRicserviceUpdateIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//ricserviceUpdateC.protocolIEs = protocolIesC

	return &ricserviceUpdateC, nil
}

func decodeRicserviceUpdate(ricserviceUpdateC *C.RICserviceUpdate_t) (*e2ap_pdu_contents.RicserviceUpdate, error) {

	//var err error
	ricserviceUpdate := e2ap_pdu_contents.RicserviceUpdate{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//ricserviceUpdate.ProtocolIes, err = decodeRicserviceUpdateIes(ricserviceUpdateC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeRicserviceUpdateIes() %s", err.Error())
	//}

	return &ricserviceUpdate, nil
}

func decodeRicserviceUpdateBytes(array [8]byte) (*e2ap_pdu_contents.RicserviceUpdate, error) {
	ricserviceUpdateC := (*C.RICserviceUpdate_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicserviceUpdate(ricserviceUpdateC)
}
