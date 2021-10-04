// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdate-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateItem) ([]byte, error) {
	e2nodeComponentConfigUpdateItemCP, err := newE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdate_Item, unsafe.Pointer(e2nodeComponentConfigUpdateItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateItem) ([]byte, error) {
	e2nodeComponentConfigUpdateItemCP, err := newE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdate_Item, unsafe.Pointer(e2nodeComponentConfigUpdateItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdate_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateItem((*C.E2nodeComponentConfigUpdate_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdate_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateItem((*C.E2nodeComponentConfigUpdate_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateItem) (*C.E2nodeComponentConfigUpdate_Item_t, error) {

	e2nodeComponentInterfaceTypeC, err := newE2nodeComponentInterfaceType(&e2nodeComponentConfigUpdateItem.E2NodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentInterfaceType() %s", err.Error())
	}

	e2nodeComponentIDC, err := newE2nodeComponentID(e2nodeComponentConfigUpdateItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigurationC, err := newE2nodeComponentConfiguration(e2nodeComponentConfigUpdateItem.E2NodeComponentConfiguration)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfiguration() %s", err.Error())
	}

	e2nodeComponentConfigUpdateItemC := C.E2nodeComponentConfigUpdate_Item_t{
		e2nodeComponentInterfaceType: *e2nodeComponentInterfaceTypeC,
		e2nodeComponentID:            *e2nodeComponentIDC,
		e2nodeComponentConfiguration: *e2nodeComponentConfigurationC,
	}

	return &e2nodeComponentConfigUpdateItemC, nil
}

func decodeE2nodeComponentConfigUpdateItem(e2nodeComponentConfigUpdateItemC *C.E2nodeComponentConfigUpdate_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {

	var err error
	e2nodeComponentConfigUpdateItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateItem{}

	componentType, err := decodeE2nodeComponentInterfaceType(&e2nodeComponentConfigUpdateItemC.e2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigUpdateItem.E2NodeComponentInterfaceType = *componentType

	e2nodeComponentConfigUpdateItem.E2NodeComponentId, err = decodeE2nodeComponentID(&e2nodeComponentConfigUpdateItemC.e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigUpdateItem.E2NodeComponentConfiguration, err = decodeE2nodeComponentConfiguration(&e2nodeComponentConfigUpdateItemC.e2nodeComponentConfiguration)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigAddition() %s", err.Error())
	}

	return &e2nodeComponentConfigUpdateItem, nil
}

func decodeE2nodeComponentConfigUpdateItemBytes(bytes [176]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateItem, error) {

	requestSize := binary.LittleEndian.Uint64(bytes[56:64])
	requestGobytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[48:56]))), C.int(requestSize))
	responseSize := binary.LittleEndian.Uint64(bytes[96:104])
	responseGobytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[88:96]))), C.int(requestSize))

	e2nodeComponentConfigUpdateItemC := C.E2nodeComponentConfigUpdate_Item_t{
		e2nodeComponentInterfaceType: C.long(binary.LittleEndian.Uint64(bytes[0:8])),
		e2nodeComponentID: C.E2nodeComponentID_t{
			present: C.E2nodeComponentID_PR(binary.LittleEndian.Uint64(bytes[8:16])),
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentID --> 48
		e2nodeComponentConfiguration: C.E2nodeComponentConfiguration_t{
			e2nodeComponentRequestPart: C.OCTET_STRING_t{
				buf:  (*C.uchar)(C.CBytes(requestGobytes)),
				size: C.ulong(requestSize),
			},
			// Gap of 24 for the asn_struct_ctx_t belonging to OCTET_STRING --> 88
			e2nodeComponentResponsePart: C.OCTET_STRING_t{
				buf:  (*C.uchar)(C.CBytes(responseGobytes)),
				size: C.ulong(responseSize),
			},
			// Gap of 24 for the asn_struct_ctx_t belonging to OCTET_STRING --> 128
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration --> 152
	}
	copy(e2nodeComponentConfigUpdateItemC.e2nodeComponentID.choice[:], bytes[16:24])
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfigUpdate_Item --> 176

	return decodeE2nodeComponentConfigUpdateItem(&e2nodeComponentConfigUpdateItemC)
}
