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
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeRICsubscriptionRequest - used only for testing
func XerEncodeRICsubscriptionRequest(rsrIE *e2ctypes.RICsubscriptionRequest_IEsT) ([]byte, error) {
	rsrIEC, err := newRICsubscriptionRequestIE(rsrIE)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_RICsubscriptionRequest_IEs, unsafe.Pointer(rsrIEC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeRICsubscriptionRequest - used only for testing
func PerEncodeRICsubscriptionRequest(rsrIE *e2ctypes.RICsubscriptionRequest_IEsT) ([]byte, error) {
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

func newRICsubscriptionRequest(rsr *e2ctypes.RICsubscriptionRequestT) (*C.RICsubscriptionRequest_t, error) {
	pIeC1544P0, err := newProtocolIeContainer1544P0(rsr.GetProtocolIEs())
	if err != nil {
		return nil, err
	}
	rsrC := C.RICsubscriptionRequest_t{
		protocolIEs: *pIeC1544P0,
	}

	return &rsrC, nil
}

func decodeRicSubscriptionRequest(rsrC *C.RICsubscriptionRequest_t) (*e2ctypes.RICsubscriptionRequestT, error) {
	pIEs, err := decodeProtocolIeContainer1544P0(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2ctypes.RICsubscriptionRequestT{
		ProtocolIEs: pIEs,
	}

	return &rsr, nil
}
