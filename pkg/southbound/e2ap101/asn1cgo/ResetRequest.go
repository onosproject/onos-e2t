// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ResetRequest.h"
import "C"

import (
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newResetRequest(rr *e2ap_pdu_contents.ResetRequest) (*C.ResetRequest_t, error) {

	pIeC1710P20, err := newResetRequestIe(rr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rrC := C.ResetRequest_t{
		protocolIEs: *pIeC1710P20,
	}

	return &rrC, nil
}

func decodeResetRequest(rrC *C.ResetRequest_t) (*e2ap_pdu_contents.ResetRequest, error) {

	pIEs, err := decodeResetRequestIes(&rrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rr := e2ap_pdu_contents.ResetRequest{
		ProtocolIes: pIEs,
	}

	return &rr, nil
}
