// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RANfunctionIDcause-Item.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func newRanFunctionIDCauseItem(rfIDi *e2appducontents.RanfunctionIdcauseItem) (*C.RANfunctionIDcause_Item_t, error) {
	cause, err := newCause(rfIDi.GetCause())
	if err != nil {
		return nil, fmt.Errorf("newCause() error %s", err.Error())
	}

	rfIDiC := C.RANfunctionIDcause_Item_t{
		ranFunctionID: newRanFunctionID(rfIDi.GetRanFunctionId()),
		cause:         *cause,
	}

	return &rfIDiC, nil
}
