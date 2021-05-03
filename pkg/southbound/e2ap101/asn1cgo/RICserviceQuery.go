// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceQuery.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicserviceQuery(ricserviceQuery *e2ap_pdu_contents.RicserviceQuery) ([]byte, error) {
	ricserviceQueryCP, err := newRicserviceQuery(ricserviceQuery)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceQuery() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceQuery, unsafe.Pointer(ricserviceQueryCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicserviceQuery() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicserviceQuery(ricserviceQuery *e2ap_pdu_contents.RicserviceQuery) ([]byte, error) {
	ricserviceQueryCP, err := newRicserviceQuery(ricserviceQuery)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceQuery() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceQuery, unsafe.Pointer(ricserviceQueryCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicserviceQuery() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicserviceQuery(bytes []byte) (*e2ap_pdu_contents.RicserviceQuery, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceQuery)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicserviceQuery((*C.RICserviceQuery_t)(unsafePtr))
}

func perDecodeRicserviceQuery(bytes []byte) (*e2ap_pdu_contents.RicserviceQuery, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceQuery)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicserviceQuery((*C.RICserviceQuery_t)(unsafePtr))
}

func newRicserviceQuery(ricserviceQuery *e2ap_pdu_contents.RicserviceQuery) (*C.RICserviceQuery_t, error) {

	//var err error
	ricserviceQueryC := C.RICserviceQuery_t{}

	//protocolIesC, err := newRicserviceQueryIes(ricserviceQuery.ProtocolIes)
	//if err != nil {
	//	return nil, fmt.Errorf("newRicserviceQueryIes() %s", err.Error())
	//}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	//ricserviceQueryC.protocolIEs = protocolIesC

	return &ricserviceQueryC, nil
}

func decodeRicserviceQuery(ricserviceQueryC *C.RICserviceQuery_t) (*e2ap_pdu_contents.RicserviceQuery, error) {

	//var err error
	ricserviceQuery := e2ap_pdu_contents.RicserviceQuery{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//ProtocolIes: protocolIes,
	}

	//ricserviceQuery.ProtocolIes, err = decodeRicserviceQueryIes(ricserviceQueryC.protocolIEs)
	//if err != nil {
	//	return nil, fmt.Errorf("decodeRicserviceQueryIes() %s", err.Error())
	//}

	return &ricserviceQuery, nil
}

func decodeRicserviceQueryBytes(array [8]byte) (*e2ap_pdu_contents.RicserviceQuery, error) {
	ricserviceQueryC := (*C.RICserviceQuery_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeRicserviceQuery(ricserviceQueryC)
}
