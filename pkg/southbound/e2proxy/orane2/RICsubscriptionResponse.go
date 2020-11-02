// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICsubscriptionResponse.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

// Deprecated: do not use
func decodeRicSubscriptionResponseOld(rsrC *C.RICsubscriptionResponse_t) (*e2ctypes.RICsubscriptionResponseT, error) {
	pIEs, err := decodeProtocolIeContainer1544P1(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2ctypes.RICsubscriptionResponseT{
		ProtocolIEs: pIEs,
	}

	return &rsr, nil
}

func decodeRicSubscriptionResponse(rsrC *C.RICsubscriptionResponse_t) (*e2appducontents.RicsubscriptionResponse, error) {
	pIEs, err := decodeRicSubscriptionResponseIes(&rsrC.protocolIEs)
	if err != nil {
		return nil, err
	}

	rsr := e2appducontents.RicsubscriptionResponse{
		ProtocolIes: pIEs,
	}

	return &rsr, nil
}
