// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICserviceUpdateAcknowledge.h"
import "C"

import (
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

func newRicServiceUpdateAcknowledge(rsua *e2ap_pdu_contents.RicserviceUpdateAcknowledge) (*C.RICserviceUpdateAcknowledge_t, error) {

	pIeC1710P23, err := newRicServiceUpdateAcknowledgeIe(rsua.ProtocolIes)
	if err != nil {
		return nil, err
	}
	rsuaC := C.RICserviceUpdateAcknowledge_t{
		protocolIEs: *pIeC1710P23,
	}

	return &rsuaC, nil
}

func decodeRicServiceUpdateAcknowledge(rsuaC *C.RICserviceUpdateAcknowledge_t) (*e2ap_pdu_contents.RicserviceUpdateAcknowledge, error) {

	pIEs, err := decodeRicServiceUpdateAcknowledgeIes(&rsuaC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsua := e2ap_pdu_contents.RicserviceUpdateAcknowledge{
		ProtocolIes: pIEs,
	}

	return &rsua, nil
}
