// SPDX-FileCopyrightText: 2020-present Open NetworkiXn Foundation <info@opennetworkiXn.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

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
		tnlInformation:    tnlC,
		tnlInformationRIC: tnlRicC,
	}

	return &e2nodeTNLassociationRemovalItemC, nil
}

func decodeE2nodeTNLassociationRemovalItem(e2nodeTNLassociationRemovalItemC *C.E2nodeTNLassociationRemoval_Item_t) (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {

	tnl, err := decodeTnlinformation(e2nodeTNLassociationRemovalItemC.tnlInformation)
	if err != nil {
		return nil, err
	}

	tnlRic, err := decodeTnlinformation(e2nodeTNLassociationRemovalItemC.tnlInformationRIC)
	if err != nil {
		return nil, err
	}

	e2nodeTNLassociationRemovalItem := e2ap_pdu_contents.E2NodeTnlassociationRemovalItem{
		TnlInformation: tnl,
		TnlInformationRic: tnlRic,
	}

	return &e2nodeTNLassociationRemovalItem, nil
}

func decodeE2nodeTNLassociationRemovalItemBytes(array [8]byte) (*e2ap_pdu_contents.E2NodeTnlassociationRemovalItem, error) {
	e2ncc := (*C.E2nodeTNLassociationRemoval_Item_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeE2nodeTNLassociationRemovalItem(e2ncc)
}
