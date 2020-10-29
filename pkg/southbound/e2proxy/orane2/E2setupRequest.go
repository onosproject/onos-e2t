// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2setupRequest.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// Deprecated: Do not use.
func newE2SetupRequestOld(esr *e2ctypes.E2SetupRequestT) (*C.E2setupRequest_t, error) {
	pIeC1544P11, err := newProtocolIeContainer1544P11(esr.GetProtocolIEs())
	if err != nil {
		return nil, err
	}
	esC := C.E2setupRequest_t{
		protocolIEs: *pIeC1544P11,
	}

	return &esC, nil
}

// Deprecated: Do not use.
func decodeE2setupRequestOld(e2setupRequestC *C.E2setupRequest_t) (*e2ctypes.E2SetupRequestT, error) {
	pIEs, err := decodeProtocolIeContainer1544P11(&e2setupRequestC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2setupRequest := e2ctypes.E2SetupRequestT{
		ProtocolIEs: pIEs,
	}
	return &e2setupRequest, nil
}

func xerEncodeE2SetupRequest(e2SetupRequest *e2appducontents.E2SetupRequest) ([]byte, error) {
	e2SetupRequestC, err := newE2SetupRequest(e2SetupRequest)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2setupRequest, unsafe.Pointer(e2SetupRequestC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perEncodeE2SetupRequest(e2SetupRequest *e2appducontents.E2SetupRequest) ([]byte, error) {
	e2SetupRequestC, err := newE2SetupRequest(e2SetupRequest)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2setupRequest, unsafe.Pointer(e2SetupRequestC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func newE2SetupRequest(esr *e2appducontents.E2SetupRequest) (*C.E2setupRequest_t, error) {
	pIeC1544P11, err := newE2SetupRequestIes(esr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	esC := C.E2setupRequest_t{
		protocolIEs: *pIeC1544P11,
	}

	return &esC, nil
}

func decodeE2setupRequest(e2setupRequestC *C.E2setupRequest_t) (*e2appducontents.E2SetupRequest, error) {
	pIEs, err := decodeE2SetupRequestIes(&e2setupRequestC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2setupRequest := e2appducontents.E2SetupRequest{
		ProtocolIes: pIEs,
	}
	return &e2setupRequest, nil
}
