// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICactionTypeVone.h"
import "C"
import (
	"fmt"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
)

func newRicActionType(rat e2apies.RicactionType) (*C.RICactionTypeVone_t, error) {
	var ret C.RICactionTypeVone_t
	switch rat {
	case e2apies.RicactionType_RICACTION_TYPE_REPORT:
		ret = C.RICactionTypeVone_report
	case e2apies.RicactionType_RICACTION_TYPE_INSERT:
		ret = C.RICactionTypeVone_insert
	case e2apies.RicactionType_RICACTION_TYPE_POLICY:
		ret = C.RICactionTypeVone_policy
	default:
		return nil, fmt.Errorf("unexpected RicActionType %v", rat)
	}
	return &ret, nil
}

//func decodeRicActionTypeBytes(bytes []byte) e2apies.RicactionType {
//	raIDC := C.long(binary.LittleEndian.Uint64(bytes[:8]))
//	return decodeRicActionType(&raIDC)
//}

func decodeRicActionType(ratC *C.RICactionTypeVone_t) e2apies.RicactionType {
	return e2apies.RicactionType(*ratC)
}
