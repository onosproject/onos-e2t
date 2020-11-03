// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICaction-Admitted-List.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
)

func decodeRicActionAdmittedListBytes(raalBytes []byte) (*e2appducontents.RicactionAdmittedList, error) {
	raalC := new(C.RICaction_Admitted_List_t)
	// TODO: Implement the rest of it

	return decodeRicActionAdmittedList(raalC)
}

func decodeRicActionAdmittedList(raalC *C.RICaction_Admitted_List_t) (*e2appducontents.RicactionAdmittedList, error) {

	raal := e2appducontents.RicactionAdmittedList{
		Value: make([]*e2appducontents.RicactionAdmittedItemIes, 0),
	}
	return &raal, nil
}
