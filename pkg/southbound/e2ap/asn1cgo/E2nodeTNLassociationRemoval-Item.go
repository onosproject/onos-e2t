// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeTNLassociationRemoval-Item.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"unsafe"
)

func xerEncodeE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItem *e2ap_pdu_contents.E2NodeTnlassociationRemovalItem) ([]byte, error) {
	e2nodeTNLassociationRemovalItemCP, err := newE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeTNLassociationRemovalItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeTNLassociationRemoval_Item, unsafe.Pointer(e2nodeTNLassociationRemovalItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeTNLassociationRemovalItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItem *e2ap_pdu_contents.E2NodeTnlassociationRemovalItem) ([]byte, error) {
	e2nodeTNLassociationRemovalItemCP, err := newE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeTNLassociationRemovalItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeTNLassociationRemoval_Item, unsafe.Pointer(e2nodeTNLassociationRemovalItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeTNLassociationRemovalItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeTNLassociationRemovalItem(bytes []byte) (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeTNLassociationRemoval_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeTNLassociationRemovalItem((*C.E2nodeTNLassociationRemoval_Item_t)(unsafePtr))
}

func perDecodeE2nodeTNLassociationRemovalItem(bytes []byte) (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeTNLassociationRemoval_Item)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeTNLassociationRemovalItem((*C.E2nodeTNLassociationRemoval_Item_t)(unsafePtr))
}

func newE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItem *e2ap_pdu_contents.E2NodeTnlassociationRemovalItem) (*C.E2nodeTNLassociationRemoval_Item_t, error) {

	tnlC, err := newTnlinformation(e2nodeTNLassociationRemovalItem.GetTnlInformation())
	if err != nil {
		return nil, err
	}

	tnlRicC, err := newTnlinformation(e2nodeTNLassociationRemovalItem.GetTnlInformationRic())
	if err != nil {
		return nil, err
	}

	e2nodeTNLassociationRemovalItemC := C.E2nodeTNLassociationRemoval_Item_t{
		tnlInformation:    *tnlC,
		tnlInformationRIC: *tnlRicC,
	}

	return &e2nodeTNLassociationRemovalItemC, nil
}

func decodeE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItemC *C.E2nodeTNLassociationRemoval_Item_t) (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {

	tnl, err := decodeTnlinformation(&e2nodeTNLassociationRemovalItemC.tnlInformation)
	if err != nil {
		return nil, err
	}

	tnlRic, err := decodeTnlinformation(&e2nodeTNLassociationRemovalItemC.tnlInformationRIC)
	if err != nil {
		return nil, err
	}

	e2nodeTNLassociationRemovalItem := e2ap_pdu_contents.E2NodeTnlassociationRemovalItem{
		TnlInformation: tnl,
		TnlInformationRic: tnlRic,
	}

	return &e2nodeTNLassociationRemovalItem, nil
}

func decodeE2nodeTNLassociationRemovalItemBytes(array [184]byte) (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {

	tnlAddrsize1 := binary.LittleEndian.Uint64(array[8:16])
	tnlAddrbitsUnused1 := int(binary.LittleEndian.Uint32(array[16:20]))
	tnlAddrbytes1 := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[:8]))), C.int(tnlAddrsize1))
	tnlPort1PtrC := (*C.BIT_STRING_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[48:56]))))

	tnlAddrsize2 := binary.LittleEndian.Uint64(array[88:96])
	tnlAddrbitsUnused2 := int(binary.LittleEndian.Uint32(array[96:100]))
	tnlAddrbytes2 := C.GoBytes(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[80:88]))), C.int(tnlAddrsize2))
	tnlPort2PtrC := (*C.BIT_STRING_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[128:136]))))

	e2nccC := C.E2nodeTNLassociationRemoval_Item_t{
		tnlInformation: C.TNLinformation_t{
			tnlAddress: C.BIT_STRING_t{
				buf:         (*C.uchar)(C.CBytes(tnlAddrbytes1)),
				size:        C.ulong(tnlAddrsize1),
				bits_unused: C.int(tnlAddrbitsUnused1),
			},
			// Gap of 24 for the asn_struct_ctx_t belonging to BIT_STRING --> 48
			tnlPort: tnlPort1PtrC,
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to TNLinformation --> 80
		tnlInformationRIC: C.TNLinformation_t{
			tnlAddress: C.BIT_STRING_t{
				buf:         (*C.uchar)(C.CBytes(tnlAddrbytes2)),
				size:        C.ulong(tnlAddrsize2),
				bits_unused: C.int(tnlAddrbitsUnused2),
			},
			// Gap of 24 for the asn_struct_ctx_t belonging to BIT_STRING --> 128
			tnlPort: tnlPort2PtrC,
		},
		// Gap of 24 for the asn_struct_ctx_t belonging to TNLinformation --> 160
	}
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeTNLassociationRemoval_Item --> 184


	return decodeE2nodeTNLassociationRemovalItem(&e2nccC)
}
