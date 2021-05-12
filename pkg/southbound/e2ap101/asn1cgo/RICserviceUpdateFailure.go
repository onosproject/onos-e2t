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
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

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
