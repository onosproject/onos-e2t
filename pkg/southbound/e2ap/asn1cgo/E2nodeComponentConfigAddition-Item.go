// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigAddition-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigAdditionItem(e2nodeComponentConfigAdditionItem *e2ap_pdu_contents.E2NodeComponentConfigAdditionItem) ([]byte, error) {
	e2nodeComponentConfigAdditionItemCP, err := newE2nodeComponentConfigAdditionItem(e2nodeComponentConfigAdditionItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigAdditionItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigAddition_Item, unsafe.Pointer(e2nodeComponentConfigAdditionItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigAdditionItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigAdditionItem(e2nodeComponentConfigAdditionItem *e2ap_pdu_contents.E2NodeComponentConfigAdditionItem) ([]byte, error) {
	e2nodeComponentConfigAdditionItemCP, err := newE2nodeComponentConfigAdditionItem(e2nodeComponentConfigAdditionItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigAdditionItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigAddition_Item, unsafe.Pointer(e2nodeComponentConfigAdditionItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigAdditionItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigAdditionItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigAddition_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigAdditionItem((*C.E2nodeComponentConfigAddition_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigAdditionItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigAddition_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigAdditionItem((*C.E2nodeComponentConfigAddition_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigAdditionItem(e2nodeComponentConfigAdditionItem *e2ap_pdu_contents.E2NodeComponentConfigAdditionItem) (*C.E2nodeComponentConfigAddition_Item_t, error) {

	e2nodeComponentInterfaceTypeC, err := newE2nodeComponentInterfaceType(&e2nodeComponentConfigAdditionItem.E2NodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentInterfaceType() %s", err.Error())
	}

	e2nodeComponentIDC, err := newE2nodeComponentID(e2nodeComponentConfigAdditionItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigurationC, err := newE2nodeComponentConfiguration(e2nodeComponentConfigAdditionItem.E2NodeComponentConfiguration)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfiguration() %s", err.Error())
	}

	e2nodeComponentConfigAdditionItemC := C.E2nodeComponentConfigAddition_Item_t{
		e2nodeComponentInterfaceType: *e2nodeComponentInterfaceTypeC,
		e2nodeComponentID:            *e2nodeComponentIDC,
		e2nodeComponentConfiguration: *e2nodeComponentConfigurationC,
	}

	return &e2nodeComponentConfigAdditionItemC, nil
}

func decodeE2nodeComponentConfigAdditionItem(e2nodeComponentConfigAdditionItemC *C.E2nodeComponentConfigAddition_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionItem, error) {

	var err error
	e2nodeComponentConfigAdditionItem := e2ap_pdu_contents.E2NodeComponentConfigAdditionItem{}

	componentType, err := decodeE2nodeComponentInterfaceType(&e2nodeComponentConfigAdditionItemC.e2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigAdditionItem.E2NodeComponentInterfaceType = *componentType

	e2nodeComponentConfigAdditionItem.E2NodeComponentId, err = decodeE2nodeComponentID(&e2nodeComponentConfigAdditionItemC.e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigAdditionItem.E2NodeComponentConfiguration, err = decodeE2nodeComponentConfiguration(&e2nodeComponentConfigAdditionItemC.e2nodeComponentConfiguration)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigAddition() %s", err.Error())
	}

	return &e2nodeComponentConfigAdditionItem, nil
}

func decodeE2nodeComponentConfigAdditionItemBytes(bytes [176]byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionItem, error) {

	requestSize := binary.LittleEndian.Uint64(bytes[56:64])
	requestGobytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[48:56]))), C.int(requestSize))
	responseSize := binary.LittleEndian.Uint64(bytes[96:104])
	responseGobytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[88:96]))), C.int(requestSize))

	e2nodeComponentConfigAdditionItemC := C.E2nodeComponentConfigAddition_Item_t{
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
	copy(e2nodeComponentConfigAdditionItemC.e2nodeComponentID.choice[:], bytes[16:24])
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfigAddition_Item --> 176

	return decodeE2nodeComponentConfigAdditionItem(&e2nodeComponentConfigAdditionItemC)
}
