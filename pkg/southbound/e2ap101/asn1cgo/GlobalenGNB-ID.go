// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalenGNB-ID.h"
import "C"

import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeGlobalenGnbID(globalenGnbID *e2ap_ies.GlobalenGnbId) ([]byte, error) {
	globalenGnbIDCP, err := newGlobalenGnbID(globalenGnbID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalenGnbID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_GlobalenGNB_ID, unsafe.Pointer(globalenGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeGlobalenGnbID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeGlobalenGnbID(globalenGnbID *e2ap_ies.GlobalenGnbId) ([]byte, error) {
	globalenGnbIDCP, err := newGlobalenGnbID(globalenGnbID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalenGnbID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GlobalenGNB_ID, unsafe.Pointer(globalenGnbIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeGlobalenGnbID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeGlobalenGnbID(bytes []byte) (*e2ap_ies.GlobalenGnbId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_GlobalenGNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeGlobalenGnbID((*C.GlobalenGNB_ID_t)(unsafePtr))
}

func perDecodeGlobalenGnbID(bytes []byte) (*e2ap_ies.GlobalenGnbId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_GlobalenGNB_ID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeGlobalenGnbID((*C.GlobalenGNB_ID_t)(unsafePtr))
}

func newGlobalenGnbID(globalenGnbID *e2ap_ies.GlobalenGnbId) (*C.GlobalenGNB_ID_t, error) {

	var err error
	globalenGnbIDC := C.GlobalenGNB_ID_t{}

	pLmnIdentityC, err := newPlmnIdentity(globalenGnbID.PLmnIdentity)
	if err != nil {
		return nil, fmt.Errorf("newPlmnIdentity() %s", err.Error())
	}

	gNbIDC, err := newEngnbID(globalenGnbID.GNbId)
	if err != nil {
		return nil, fmt.Errorf("newEngnbID() %s", err.Error())
	}

	globalenGnbIDC.pLMN_Identity = *pLmnIdentityC
	globalenGnbIDC.gNB_ID = *gNbIDC

	return &globalenGnbIDC, nil
}

func decodeGlobalenGnbID(globalenGnbIDC *C.GlobalenGNB_ID_t) (*e2ap_ies.GlobalenGnbId, error) {

	var err error
	globalenGnbID := e2ap_ies.GlobalenGnbId{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		//PLmnIdentity: pLmnIdentity,
		//GNbId: gNbId,
	}

	globalenGnbID.PLmnIdentity, err = decodePlmnIdentity(&globalenGnbIDC.pLMN_Identity)
	if err != nil {
		return nil, fmt.Errorf("decodePlmnIdentity() %s", err.Error())
	}

	globalenGnbID.GNbId, err = decodeEngnbID(&globalenGnbIDC.gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("decodeEngnbID() %s", err.Error())
	}

	return &globalenGnbID, nil
}
