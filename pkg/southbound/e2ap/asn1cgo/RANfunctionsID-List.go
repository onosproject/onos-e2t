// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Field.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func newRanFunctionsIDList(rfIDl *e2appducontents.RanfunctionsIdList) (*C.RANfunctionsID_List_t, error) {
	rfIDlC := C.RANfunctionsID_List_t{
		list: C.struct___49{
			size:  C.int(0),
			count: C.int(0),
		},
	}
	// TODO: Implement list

	return &rfIDlC, nil
}
