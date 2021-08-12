// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2connectionUpdate-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2connectionUpdateItem(e2connectionUpdateItem *e2ap_pdu_contents.E2ConnectionUpdateItem) ([]byte, error) {
	e2connectionUpdateItemCP, err := newE2connectionUpdateItem(e2connectionUpdateItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2connectionUpdate_Item, unsafe.Pointer(e2connectionUpdateItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2connectionUpdateItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2connectionUpdateItem(e2connectionUpdateItem *e2ap_pdu_contents.E2ConnectionUpdateItem) ([]byte, error) {
	e2connectionUpdateItemCP, err := newE2connectionUpdateItem(e2connectionUpdateItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2connectionUpdate_Item, unsafe.Pointer(e2connectionUpdateItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2connectionUpdateItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2connectionUpdateItem(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2connectionUpdate_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2connectionUpdateItem((*C.E2connectionUpdate_Item_t)(unsafePtr))
}

func perDecodeE2connectionUpdateItem(bytes []byte) (*e2ap_pdu_contents.E2ConnectionUpdateItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2connectionUpdate_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2connectionUpdateItem((*C.E2connectionUpdate_Item_t)(unsafePtr))
}

func newE2connectionUpdateItem(e2connectionUpdateItem *e2ap_pdu_contents.E2ConnectionUpdateItem) (*C.E2connectionUpdate_Item_t, error) {

	var err error
	e2connectionUpdateItemC := C.E2connectionUpdate_Item_t{}

	tnlInformationC, err := newTnlinformation(e2connectionUpdateItem.TnlInformation)
	if err != nil {
		return nil, fmt.Errorf("newTnlinformation() %s", err.Error())
	}

	tnlUsageC, err := newTnlusage(&e2connectionUpdateItem.TnlUsage)
	if err != nil {
		return nil, fmt.Errorf("newTnlusage() %s", err.Error())
	}

	//ToDo - check whether pointers passed correctly with regard to C-struct's definition .h file
	e2connectionUpdateItemC.tnlInformation = *tnlInformationC
	e2connectionUpdateItemC.tnlUsage = *tnlUsageC

	return &e2connectionUpdateItemC, nil
}

func decodeE2connectionUpdateItem(e2connectionUpdateItemC *C.E2connectionUpdate_Item_t) (*e2ap_pdu_contents.E2ConnectionUpdateItem, error) {

	var err error
	e2connectionUpdateItem := e2ap_pdu_contents.E2ConnectionUpdateItem{}

	e2connectionUpdateItem.TnlInformation, err = decodeTnlinformation(&e2connectionUpdateItemC.tnlInformation)
	if err != nil {
		return nil, fmt.Errorf("decodeTnlinformation() %s", err.Error())
	}

	tnlUsage, err := decodeTnlusage(&e2connectionUpdateItemC.tnlUsage)
	if err != nil {
		return nil, fmt.Errorf("decodeTnlusage() %s", err.Error())
	}
	e2connectionUpdateItem.TnlUsage = *tnlUsage

	return &e2connectionUpdateItem, nil
}

func decodeE2connectionUpdateItemBytes(array [112]byte) (*e2ap_pdu_contents.E2ConnectionUpdateItem, error) {

	tnlAddrsize := binary.LittleEndian.Uint64(array[8:16])
	tnlAddrbitsUnused := int(binary.LittleEndian.Uint32(array[16:20]))
	tnlAddrbytes := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[:8]))), C.int(tnlAddrsize))
	tnlPortPtrC := (*C.BIT_STRING_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[48:56]))))

	e2cuItemC := C.E2connectionUpdate_Item_t{
		tnlInformation: C.TNLinformation_t{
			tnlAddress: C.BIT_STRING_t{
				buf:         (*C.uchar)(C.CBytes(tnlAddrbytes)),
				size:        C.ulong(tnlAddrsize),
				bits_unused: C.int(tnlAddrbitsUnused),
			},
			tnlPort: tnlPortPtrC,
		},
		tnlUsage: C.long(binary.LittleEndian.Uint64(array[80:])),
	}

	return decodeE2connectionUpdateItem(&e2cuItemC)
}
