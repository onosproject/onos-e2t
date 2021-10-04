// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigRemovalAck-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigRemovalAckItem(e2nodeComponentConfigRemovalAckItem *e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem) ([]byte, error) {
	e2nodeComponentConfigRemovalAckItemCP, err := newE2nodeComponentConfigRemovalAckItem(e2nodeComponentConfigRemovalAckItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigRemovalAckItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigRemovalAck_Item, unsafe.Pointer(e2nodeComponentConfigRemovalAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigRemovalAckItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigRemovalAckItem(e2nodeComponentConfigRemovalAckItem *e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem) ([]byte, error) {
	e2nodeComponentConfigRemovalAckItemCP, err := newE2nodeComponentConfigRemovalAckItem(e2nodeComponentConfigRemovalAckItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigRemovalAckItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigRemovalAck_Item, unsafe.Pointer(e2nodeComponentConfigRemovalAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigRemovalAckItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigRemovalAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigRemovalAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigRemovalAckItem((*C.E2nodeComponentConfigRemovalAck_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigRemovalAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigRemovalAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigRemovalAckItem((*C.E2nodeComponentConfigRemovalAck_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigRemovalAckItem(e2nodeComponentConfigRemovalAckItem *e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem) (*C.E2nodeComponentConfigRemovalAck_Item_t, error) {

	e2nodeComponentInterfaceTypeC, err := newE2nodeComponentInterfaceType(&e2nodeComponentConfigRemovalAckItem.E2NodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentInterfaceType() %s", err.Error())
	}

	e2nodeComponentIDC, err := newE2nodeComponentID(e2nodeComponentConfigRemovalAckItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigurationC, err := newE2nodeComponentConfigurationAck(e2nodeComponentConfigRemovalAckItem.E2NodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfiguration() %s", err.Error())
	}

	e2nodeComponentConfigRemovalAckItemC := C.E2nodeComponentConfigRemovalAck_Item_t{
		e2nodeComponentInterfaceType:    *e2nodeComponentInterfaceTypeC,
		e2nodeComponentID:               *e2nodeComponentIDC,
		e2nodeComponentConfigurationAck: *e2nodeComponentConfigurationC,
	}
	return &e2nodeComponentConfigRemovalAckItemC, nil
}

func decodeE2nodeComponentConfigRemovalAckItem(e2nodeComponentConfigRemovalAckItemC *C.E2nodeComponentConfigRemovalAck_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem, error) {

	var err error
	e2nodeComponentConfigRemovalAckItem := e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem{}

	componentType, err := decodeE2nodeComponentInterfaceType(&e2nodeComponentConfigRemovalAckItemC.e2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigRemovalAckItem.E2NodeComponentInterfaceType = *componentType

	e2nodeComponentConfigRemovalAckItem.E2NodeComponentId, err = decodeE2nodeComponentID(&e2nodeComponentConfigRemovalAckItemC.e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigRemovalAckItem.E2NodeComponentConfigurationAck, err = decodeE2nodeComponentConfigurationAck(&e2nodeComponentConfigRemovalAckItemC.e2nodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigAddition() %s", err.Error())
	}

	return &e2nodeComponentConfigRemovalAckItem, nil
}

func decodeE2nodeComponentConfigRemovalAckItemBytes(bytes [112]byte) (*e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItem, error) {

	e2nodeComponentConfigRemovalAckItemC := C.E2nodeComponentConfigRemovalAck_Item_t{
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
	copy(e2nodeComponentConfigRemovalAckItemC.e2nodeComponentID.choice[:], bytes[16:24])
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfigAdditionAck_Item --> 112

	return decodeE2nodeComponentConfigRemovalAckItem(&e2nodeComponentConfigRemovalAckItemC)
}
