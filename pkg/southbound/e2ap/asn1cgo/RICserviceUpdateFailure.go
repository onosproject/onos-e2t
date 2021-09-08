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
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeRicServiceUpdateFailure(rsuf *e2ap_pdu_contents.RicserviceUpdateFailure) ([]byte, error) {
	rsufCP, err := newRicServiceUpdateFailure(rsuf)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicServiceUpdateFailure() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_RICserviceUpdateFailure, unsafe.Pointer(rsufCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeRicServiceUpdateFailure() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeRicServiceUpdateFailure(rsuf *e2ap_pdu_contents.RicserviceUpdateFailure) ([]byte, error) {
	rsufCP, err := newRicServiceUpdateFailure(rsuf)
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicServiceFailure() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICserviceUpdateFailure, unsafe.Pointer(rsufCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeRicServiceFailure() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeRicServiceUpdateFailure(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_RICserviceUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeRicServiceUpdateFailure((*C.RICserviceUpdateFailure_t)(unsafePtr))
}

func perDecodeRicServiceUpdateFailure(bytes []byte) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_RICserviceUpdateFailure)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeRicServiceUpdateFailure((*C.RICserviceUpdateFailure_t)(unsafePtr))
}

func newRicServiceUpdateFailure(rsuf *e2ap_pdu_contents.RicserviceUpdateFailure) (*C.RICserviceUpdateFailure_t, error) {

	pIeC1710P24, err := newRicServiceUpdateFailureIe(rsuf.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsufC := C.RICserviceUpdateFailure_t{
		protocolIEs: *pIeC1710P24,
	}

	return &rsufC, nil
}

func decodeRicServiceUpdateFailure(rsufC *C.RICserviceUpdateFailure_t) (*e2ap_pdu_contents.RicserviceUpdateFailure, error) {

	pIEs, err := decodeRicServiceUpdateFailureIes(&rsufC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsuf := e2ap_pdu_contents.RicserviceUpdateFailure{
		ProtocolIes: pIEs,
	}

	return &rsuf, nil
}
