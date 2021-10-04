// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "AMFName.h"
import "C"
import (
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
)

func newAmfName(amf *e2ap_commondatatypes.Amfname) *C.AMFName_t {

	return newPrintableString(amf.Value)
}

func decodeAmfName(amfC *C.AMFName_t) *e2ap_commondatatypes.Amfname {
	amf := decodePrintableString(amfC)
	result := e2ap_commondatatypes.Amfname{
		Value: amf,
	}

	return &result
}
