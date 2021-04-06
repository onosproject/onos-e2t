// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICcontrolFailure.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"fmt"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRICcontrolFailure(rcf *e2appducontents.RiccontrolFailure) ([]byte, error) {
	rcfC, err := newRicControlFailure(rcf)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICcontrolFailure, unsafe.Pointer(rcfC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func xerDecodeRICcontrolFailure(bytes []byte) (*e2appducontents.RiccontrolFailure, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICcontrolFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicControlFailure((*C.RICcontrolFailure_t)(unsafePtr))
}

func perEncodeRICcontrolFailure(rcf *e2appducontents.RiccontrolFailure) ([]byte, error) {
	rcfC, err := newRicControlFailure(rcf)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICcontrolFailure, unsafe.Pointer(rcfC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perDecodeRICcontrolFailure(bytes []byte) (*e2appducontents.RiccontrolFailure, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICcontrolFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicControlFailure((*C.RICcontrolFailure_t)(unsafePtr))
}

func newRicControlFailure(rcf *e2appducontents.RiccontrolFailure) (*C.RICcontrolFailure_t, error) {
	pIeC1710P9, err := newRicControlFailureIEs(rcf.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rcfC := C.RICcontrolFailure_t{
		protocolIEs: *pIeC1710P9,
	}

	return &rcfC, nil
}

func decodeRicControlFailure(rcfC *C.RICcontrolFailure_t) (*e2appducontents.RiccontrolFailure, error) {
	pIEs, err := decodeRicControlFailureIes(&rcfC.protocolIEs)
	if err != nil {
		return nil, err
	}

	ricControlFailure := e2appducontents.RiccontrolFailure{
		ProtocolIes: pIEs,
	}
	return &ricControlFailure, nil
}
