// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "ErrorIndication.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

// XerDecodeErrorIndication - just used for test only
func XerDecodeErrorIndication(bytes []byte) (*e2ctypes.ErrorIndicationT, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ErrorIndication)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	errorIndicationC := (*C.ErrorIndication_t)(unsafePtr)
	errorIndication, err := decodeErrorIndication(errorIndicationC)
	if err != nil {
		return nil, err
	}

	return errorIndication, nil
}

func decodeErrorIndication(errIndChoice *C.ErrorIndication_t) (*e2ctypes.ErrorIndicationT, error) {
	pIEs, err := decodeProtocolIeContainer1544P10(&errIndChoice.protocolIEs)
	if err != nil {
		return nil, err
	}

	errorIndication := &e2ctypes.ErrorIndicationT{
		ProtocolIEs: pIEs,
	}
	return errorIndication, nil
}
