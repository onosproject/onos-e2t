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
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func newE2SetupRequest(esr *e2ctypes.E2SetupRequestT) (*C.E2setupRequest_t, error) {
	pIeC1544P11, err := newProtocolIeContainer1544P11(esr.GetProtocolIEs())
	if err != nil {
		return nil, err
	}
	esC := C.E2setupRequest_t{
		protocolIEs: *pIeC1544P11,
	}

	return &esC, nil
}

func decodeE2setupRequest(e2setupRequestC *C.E2setupRequest_t) (*e2ctypes.E2SetupRequestT, error) {
	pIEs, err := decodeProtocolIeContainer1544P11(&e2setupRequestC.protocolIEs)
	if err != nil {
		return nil, err
	}

	e2setupRequest := e2ctypes.E2SetupRequestT{
		ProtocolIEs: pIEs,
	}
	return &e2setupRequest, nil
}
