// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscriptionDeleteRequired.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func newRicSubscriptionDeleteRequired(rsr *e2appducontents.RicsubscriptionDeleteRequired) (*C.RICsubscriptionDeleteRequired_t, error) {
	pIeC1908P6, err := newRicSubscriptionDeleteRequiredIes(rsr.GetProtocolIes())
	if err != nil {
		return nil, err
	}
	rsrC := C.RICsubscriptionDeleteRequired_t{
		protocolIEs: *pIeC1908P6,
	}

	return &rsrC, nil
}

func decodeRicSubscriptionDeleteRequired(ricSubscriptionDeleteRequiredC *C.RICsubscriptionDeleteRequired_t) (*e2appducontents.RicsubscriptionDeleteRequired, error) {
	pIEs, err := decodeRicSubscriptionDeleteRequiredIes(&ricSubscriptionDeleteRequiredC.protocolIEs)
	if err != nil {
		return nil, err
	}

	ricSubscriptionDeleteRequired := e2appducontents.RicsubscriptionDeleteRequired{
		ProtocolIes: pIEs,
	}
	return &ricSubscriptionDeleteRequired, nil
}
