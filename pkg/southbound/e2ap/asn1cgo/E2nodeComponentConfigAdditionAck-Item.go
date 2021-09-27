// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigAdditionAck-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func xerEncodeE2nodeComponentConfigAdditionAckItem(e2nodeComponentConfigAdditionAckItem *e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem) ([]byte, error) {
	e2nodeComponentConfigAdditionAckItemCP, err := newE2nodeComponentConfigAdditionAckItem(e2nodeComponentConfigAdditionAckItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigAdditionAckItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigAdditionAck_Item, unsafe.Pointer(e2nodeComponentConfigAdditionAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigAdditionAckItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigAdditionAckItem(e2nodeComponentConfigAdditionAckItem *e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem) ([]byte, error) {
	e2nodeComponentConfigAdditionAckItemCP, err := newE2nodeComponentConfigAdditionAckItem(e2nodeComponentConfigAdditionAckItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigAdditionAckItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigAdditionAck_Item, unsafe.Pointer(e2nodeComponentConfigAdditionAckItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigAdditionAckItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigAdditionAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigAdditionAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigAdditionAckItem((*C.E2nodeComponentConfigAdditionAck_Item_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigAdditionAckItem(bytes []byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigAdditionAck_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigAdditionAckItem((*C.E2nodeComponentConfigAdditionAck_Item_t)(unsafePtr))
}

func newE2nodeComponentConfigAdditionAckItem(e2nodeComponentConfigAdditionAckItem *e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem) (*C.E2nodeComponentConfigAdditionAck_Item_t, error) {

	e2nodeComponentInterfaceTypeC, err := newE2nodeComponentInterfaceType(&e2nodeComponentConfigAdditionAckItem.E2NodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentInterfaceType() %s", err.Error())
	}

	e2nodeComponentIDC, err := newE2nodeComponentID(e2nodeComponentConfigAdditionAckItem.E2NodeComponentId)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigurationC, err := newE2nodeComponentConfigurationAck(e2nodeComponentConfigAdditionAckItem.E2NodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfiguration() %s", err.Error())
	}

	e2nodeComponentConfigAdditionAckItemC := C.E2nodeComponentConfigAdditionAck_Item_t{
		e2nodeComponentInterfaceType:    *e2nodeComponentInterfaceTypeC,
		e2nodeComponentID:               *e2nodeComponentIDC,
		e2nodeComponentConfigurationAck: *e2nodeComponentConfigurationC,
	}

	return &e2nodeComponentConfigAdditionAckItemC, nil
}

func decodeE2nodeComponentConfigAdditionAckItem(e2nodeComponentConfigAdditionAckItemC *C.E2nodeComponentConfigAdditionAck_Item_t) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem, error) {

	var err error
	e2nodeComponentConfigAdditionAckItem := e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem{}

	componentType, err := decodeE2nodeComponentInterfaceType(&e2nodeComponentConfigAdditionAckItemC.e2nodeComponentInterfaceType)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentType() %s", err.Error())
	}
	e2nodeComponentConfigAdditionAckItem.E2NodeComponentInterfaceType = *componentType

	e2nodeComponentConfigAdditionAckItem.E2NodeComponentId, err = decodeE2nodeComponentID(&e2nodeComponentConfigAdditionAckItemC.e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentID() %s", err.Error())
	}

	e2nodeComponentConfigAdditionAckItem.E2NodeComponentConfigurationAck, err = decodeE2nodeComponentConfigurationAck(&e2nodeComponentConfigAdditionAckItemC.e2nodeComponentConfigurationAck)
	if err != nil {
		return nil, fmt.Errorf("decodeE2nodeComponentConfigAddition() %s", err.Error())
	}

	return &e2nodeComponentConfigAdditionAckItem, nil
}

func decodeE2nodeComponentConfigAdditionAckItemBytes(bytes [112]byte) (*e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItem, error) {

	e2nodeComponentConfigAdditionAckItemC := C.E2nodeComponentConfigAdditionAck_Item_t{
		e2nodeComponentInterfaceType: C.long(binary.LittleEndian.Uint64(bytes[0:8])),
		e2nodeComponentID: C.E2nodeComponentID_t{
			present: C.E2nodeComponentID_PR(binary.LittleEndian.Uint64(bytes[8:16])),
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentID --> 48
		e2nodeComponentConfigurationAck: C.E2nodeComponentConfigurationAck_t{
			updateOutcome: C.long(binary.LittleEndian.Uint64(bytes[48:56])),
			failureCause:  (*C.struct_Cause)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(bytes[56:64])))), // OPTIONAL, so it's a pinter
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfigurationAck --> 88
	}
	copy(e2nodeComponentConfigAdditionAckItemC.e2nodeComponentID.choice[:], bytes[16:24])
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfigAdditionAck_Item --> 112

	return decodeE2nodeComponentConfigAdditionAckItem(&e2nodeComponentConfigAdditionAckItemC)
}
