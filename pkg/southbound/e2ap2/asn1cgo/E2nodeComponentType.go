// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentType.h"
import "C"
import (
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentType(e2nodeComponentType *e2ap_ies.E2NodeComponentType) ([]byte, error) {
	e2nodeComponentTypeCP, err := newE2nodeComponentType(e2nodeComponentType)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentType, unsafe.Pointer(e2nodeComponentTypeCP)) //ToDo - change name of C-encoder tag
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentType() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentType(e2nodeComponentType *e2ap_ies.E2NodeComponentType) ([]byte, error) {
	e2nodeComponentTypeCP, err := newE2nodeComponentType(e2nodeComponentType)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentType, unsafe.Pointer(e2nodeComponentTypeCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentType() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentType(bytes []byte) (*e2ap_ies.E2NodeComponentType, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentType)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentType((*C.E2nodeComponentType_t)(unsafePtr)) //ToDo - change name of C-struct
}

func perDecodeE2nodeComponentType(bytes []byte) (*e2ap_ies.E2NodeComponentType, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentType)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentType((*C.E2nodeComponentType_t)(unsafePtr))
}

func newE2nodeComponentType(e2nodeComponentType *e2ap_ies.E2NodeComponentType) (*C.E2nodeComponentType_t, error) {
	var ret C.E2nodeComponentType_t
	switch *e2nodeComponentType {
	case e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB:
		ret = C.E2nodeComponentType_gNB
	case e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB_CU_UP:
		ret = C.E2nodeComponentType_gNB_CU_UP
	case e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_G_NB_DU:
		ret = C.E2nodeComponentType_gNB_DU
	case e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_EN_G_NB:
		ret = C.E2nodeComponentType_en_gNB
	case e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_E_NB:
		ret = C.E2nodeComponentType_eNB
	case e2ap_ies.E2NodeComponentType_E2NODE_COMPONENT_TYPE_NG_E_NB:
		ret = C.E2nodeComponentType_ng_eNB
	default:
		return nil, fmt.Errorf("unexpected E2nodeComponentType %v", e2nodeComponentType)
	}

	return &ret, nil
}

func decodeE2nodeComponentType(e2nodeComponentTypeC *C.E2nodeComponentType_t) (*e2ap_ies.E2NodeComponentType, error) {

	//ToDo: int32 shouldn't be valid all the time -- investigate in data type conversion (casting) more
	e2nodeComponentType := e2ap_ies.E2NodeComponentType(int32(*e2nodeComponentTypeC))

	return &e2nodeComponentType, nil
}
