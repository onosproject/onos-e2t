// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ResetResponse.h"
import "C"

import (
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newResetResponse(rr *e2ap_pdu_contents.ResetResponse) (*C.ResetResponse_t, error) {

	pIeC1710P21, err := newResetResponseIe(rr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rrC := C.ResetResponse_t{
		protocolIEs: *pIeC1710P21,
	}

	return &rrC, nil
}

func decodeResetResponse(rrC *C.ResetResponse_t) (*e2ap_pdu_contents.ResetResponse, error) {

	pIEs, err := decodeResetResponseIes(&rrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rr := e2ap_pdu_contents.ResetResponse{
		ProtocolIes: pIEs,
	}

	return &rr, nil
}
