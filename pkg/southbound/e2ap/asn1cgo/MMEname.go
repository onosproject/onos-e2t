// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MMEname.h"
import "C"
import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newMmeName(mme *e2ap_commondatatypes.Mmename) *C.MMEname_t {

	return newPrintableString(mme.Value)
}

func decodeMmeName(mmeC *C.MMEname_t) *e2ap_commondatatypes.Mmename {
	mme := decodePrintableString(mmeC)
	result := e2ap_commondatatypes.Mmename{
		Value: mme,
	}

	return &result
}
