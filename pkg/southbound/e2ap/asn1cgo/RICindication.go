// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

import "C"
import "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICindication.h"

func newRicIndication(ri *e2appducontents.Ricindication) (*C.RICindication_t, error) {
	pIeC1544P6, err := newRicIndicationIEs(ri.ProtocolIes)
	if err != nil {
		return nil, err
	}
	riC := C.RICindication_t{
		protocolIEs: *pIeC1544P6,
	}

	return &riC, nil
}
