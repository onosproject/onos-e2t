// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdateAck-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem) ([]byte, error) {
	e2nodeComponentConfigUpdateAckItemCP, err := newE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdateAck_Item, unsafe.Pointer(e2nodeComponentConfigUpdateAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem) ([]byte, error) {
	e2nodeComponentConfigUpdateAckItemCP, err := newE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdateAck_Item, unsafe.Pointer(e2nodeComponentConfigUpdateAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdateAckItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdateAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdateAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAckItem((*C.E2nodeComponentConfigUpdateAck_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdateAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdateAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdateAckItem((*C.E2nodeComponentConfigUpdateAck_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItem *e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem) (*C.E2nodeComponentConfigUpdateAck_Item_t, error) {

	e2nodeComponentInterfaceTypeC, err := newE2nodeComponentInterfaceType(&e2nodeComponentConfigUpdateAckItem.E2NodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentInterfaceType() %s", err.Error())
	}

	e2nodeComponentIDC, err := newE2nodeComponentID(e2nodeComponentConfigUpdateAckItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigurationAckC, err := newE2nodeComponentConfigurationAck(e2nodeComponentConfigUpdateAckItem.E2NodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigurationAck() %s", err.Error())
	}

	e2nodeComponentConfigUpdateAckItemC := C.E2nodeComponentConfigUpdateAck_Item_t{
		e2nodeComponentInterfaceType:    *e2nodeComponentInterfaceTypeC,
		e2nodeComponentID:               *e2nodeComponentIDC,
		e2nodeComponentConfigurationAck: *e2nodeComponentConfigurationAckC,
	}

	return &e2nodeComponentConfigUpdateAckItemC, nil
}

func decodeE2nodeComponentConfigUpdateAckItem(e2nodeComponentConfigUpdateAckItemC *C.E2nodeComponentConfigUpdateAck_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {

	var err error
	e2nodeComponentConfigUpdateAckItem := e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem{}

	componentType, err := decodeE2nodeComponentInterfaceType(&e2nodeComponentConfigUpdateAckItemC.e2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigUpdateAckItem.E2NodeComponentInterfaceType = *componentType

	e2nodeComponentConfigUpdateAckItem.E2NodeComponentId, err = decodeE2nodeComponentID(&e2nodeComponentConfigUpdateAckItemC.e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigUpdateAckItem.E2NodeComponentConfigurationAck, err = decodeE2nodeComponentConfigurationAck(&e2nodeComponentConfigUpdateAckItemC.e2nodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigAddition() %s", err.Error())
	}

	return &e2nodeComponentConfigUpdateAckItem, nil
}

func decodeE2nodeComponentConfigUpdateAckItemBytes(bytes [112]byte) (*e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItem, error) {

	e2nodeComponentConfigUpdateItemC := C.E2nodeComponentConfigUpdateAck_Item_t{
		e2nodeComponentInterfaceType: C.long(binary.LittleEndian.Uint64(bytes[0:8])),
		e2nodeComponentID: C.E2nodeComponentID_t{
			present: C.E2nodeComponentID_PR(binary.LittleEndian.Uint64(bytes[8:16])),
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentID --> 48
		e2nodeComponentConfigurationAck: C.E2nodeComponentConfigurationAck_t{
			updateOutcome: C.long(binary.LittleEndian.Uint64(bytes[48:56])),
			failureCause:  (*C.struct_Cause)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[56:64])))), // OPTIONAL, so it's a pinter
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration --> 88
	}
	copy(e2nodeComponentConfigUpdateItemC.e2nodeComponentID.choice[:], bytes[16:24])
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfigAdditionAck_Item --> 112

	return decodeE2nodeComponentConfigUpdateAckItem(&e2nodeComponentConfigUpdateItemC)
}
