// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscriptionRequest.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// PerEncodeRICsubscriptionRequestOld - used only for testing
// Deprecated: Do not use.
func PerEncodeRICsubscriptionRequestOld(rsrIE *e2ctypes.RICsubscriptionRequest_IEsT) ([]byte, error) {
	rsrIEC, err := newRICsubscriptionRequestIE(rsrIE)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICsubscriptionRequest_IEs, unsafe.Pointer(rsrIEC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Deprecated: Do not use.
func newRICsubscriptionRequestOld(rsr *e2ctypes.RICsubscriptionRequestT) (*C.RICsubscriptionRequest_t, error) {
	pIeC1544P0, err := newProtocolIeContainer1544P0(rsr.GetProtocolIEs())
	if err != nil {
		return nil, err
	}
	rsrC := C.RICsubscriptionRequest_t{
		protocolIEs: *pIeC1544P0,
	}

	return &rsrC, nil
}

// Deprecated: Do not use.
func decodeRicSubscriptionRequestOld(rsrC *C.RICsubscriptionRequest_t) (*e2ctypes.RICsubscriptionRequestT, error) {
	pIEs, err := decodeProtocolIeContainer1544P0(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2ctypes.RICsubscriptionRequestT{
		ProtocolIEs: pIEs,
	}

	return &rsr, nil
}

func xerEncodeRICsubscriptionRequest(rsr *e2appducontents.RicsubscriptionRequest) ([]byte, error) {
	rsrC, err := newRICsubscriptionRequest(rsr)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICsubscriptionRequest, unsafe.Pointer(rsrC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func perEncodeRICsubscriptionRequest(rsr *e2appducontents.RicsubscriptionRequest) ([]byte, error) {
	rsrC, err := newRICsubscriptionRequest(rsr)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_RICsubscriptionRequest, unsafe.Pointer(rsrC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newRICsubscriptionRequest(rsr *e2appducontents.RicsubscriptionRequest) (*C.RICsubscriptionRequest_t, error) {
	pIeC1544P0, err := newRicSubscriptionRequestIes(rsr.GetProtocolIes())
	if err != nil {
		return nil, err
	}
	rsrC := C.RICsubscriptionRequest_t{
		protocolIEs: *pIeC1544P0,
	}

	return &rsrC, nil
}

func decodeRicSubscriptionRequest(rsrC *C.RICsubscriptionRequest_t) (*e2appducontents.RicsubscriptionRequest, error) {
	pIEs, err := decodeRicSubscriptionRequestIes(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2appducontents.RicsubscriptionRequest{
		ProtocolIes: pIEs,
	}

	return &rsr, nil
}

