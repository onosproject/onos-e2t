// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalngeNB-ID.h"
import "C"

import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeGlobalngeNbID(globalngeNbID *e2ap_ies.GlobalngeNbId) ([]byte, error) {
	globalngeNbIDCP, err := newGlobalngeNbID(globalngeNbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalngeNbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalngeNB_ID, unsafe.Pointer(globalngeNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalngeNbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalngeNbID(globalngeNbID *e2ap_ies.GlobalngeNbId) ([]byte, error) {
	globalngeNbIDCP, err := newGlobalngeNbID(globalngeNbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalngeNbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalngeNB_ID, unsafe.Pointer(globalngeNbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalngeNbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalngeNbID(bytes []byte) (*e2ap_ies.GlobalngeNbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalngeNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalngeNbID((*C.GlobalngeNB_ID_t)(unsafePtr))
}

func perDecodeGlobalngeNbID(bytes []byte) (*e2ap_ies.GlobalngeNbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalngeNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalngeNbID((*C.GlobalngeNB_ID_t)(unsafePtr))
}

func newGlobalngeNbID(globalngeNbID *e2ap_ies.GlobalngeNbId) (*C.GlobalngeNB_ID_t, error) {

	var err error
	globalngeNbIDC := C.GlobalngeNB_ID_t{}

	plmnIDC, err := newPlmnIdentity(globalngeNbID.PlmnId)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	enbIDC, err := newEnbIDChoice(globalngeNbID.EnbId)
	if err != nil {
		return nil, fmt.Errorf("newEnbIDChoice() %s", err.Error())
	}

	globalngeNbIDC.plmn_id = *plmnIDC
	globalngeNbIDC.enb_id = *enbIDC

	return &globalngeNbIDC, nil
}

func decodeGlobalngeNbID(globalngeNbIDC *C.GlobalngeNB_ID_t) (*e2ap_ies.GlobalngeNbId, error) {

	var err error
	globalngeNbID := e2ap_ies.GlobalngeNbId{}

	globalngeNbID.PlmnId, err = decodePlmnIdentity(&globalngeNbIDC.plmn_id)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	globalngeNbID.EnbId, err = decodeEnbIDChoice(&globalngeNbIDC.enb_id)
	if err != nil {
		return nil, fmt.Errorf("decodeEnbIDChoice() %s", err.Error())
	}

	return &globalngeNbID, nil
}
