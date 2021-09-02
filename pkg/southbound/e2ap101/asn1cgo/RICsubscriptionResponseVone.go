// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscriptionResponseVone.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newRicSubscriptionResponse(rsr *e2appducontents.RicsubscriptionResponse) (*C.RICsubscriptionResponseVone_t, error) {
	pIeC1710P1, err := newRicSubscriptionResponseIe(rsr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsrC := C.RICsubscriptionResponseVone_t{
		protocolIEs: *pIeC1710P1,
	}

	return &rsrC, nil
}

func decodeRicSubscriptionResponse(rsrC *C.RICsubscriptionResponseVone_t) (*e2appducontents.RicsubscriptionResponse, error) {
	pIEs, err := decodeRicSubscriptionResponseIes(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2appducontents.RicsubscriptionResponse{
		ProtocolIes: pIEs,
	}

	return &rsr, nil
}
