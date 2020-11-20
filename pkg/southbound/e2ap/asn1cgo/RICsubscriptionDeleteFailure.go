// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscriptionDeleteFailure.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func newRicSubscriptionDeleteFailure(rsr *e2appducontents.RicsubscriptionDeleteFailure) (*C.RICsubscriptionDeleteFailure_t, error) {
	pIeC1544P5, err := newRicSubscriptionDeleteFailureIe(rsr.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsrC := C.RICsubscriptionDeleteFailure_t{
		protocolIEs: *pIeC1544P5,
	}

	return &rsrC, nil
}

func decodeRicSubscriptionDeleteFailure(rsrC *C.RICsubscriptionDeleteFailure_t) (*e2appducontents.RicsubscriptionDeleteFailure, error) {
	pIEs, err := decodeRicSubscriptionDeleteFailureIes(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2appducontents.RicsubscriptionDeleteFailure{
		ProtocolIes: pIEs,
	}

	return &rsr, nil
}
