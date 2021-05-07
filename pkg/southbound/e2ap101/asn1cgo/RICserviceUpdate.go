// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceUpdate.h"
import "C"

import (
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newRicServiceUpdate(rsu *e2ap_pdu_contents.RicserviceUpdate) (*C.RICserviceUpdate_t, error) {

	pIeC1710P22, err := newRicServiceUpdateIe(rsu.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsuC := C.RICserviceUpdate_t{
		protocolIEs: *pIeC1710P22,
	}

	return &rsuC, nil
}

func decodeRicServiceUpdate(rsuC *C.RICserviceUpdate_t) (*e2ap_pdu_contents.RicserviceUpdate, error) {

	pIEs, err := decodeRicServiceUpdateIes(&rsuC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsu := e2ap_pdu_contents.RicserviceUpdate{
		ProtocolIes: pIEs,
	}

	return &rsu, nil
}
