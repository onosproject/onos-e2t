// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentID.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
)

func xerEncodeE2nodeComponentID(e2nodeComponentID *e2ap_ies.E2NodeComponentId) ([]byte, error) {
	e2nodeComponentIDCP, err := newE2nodeComponentID(e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentID() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentID, unsafe.Pointer(e2nodeComponentIDCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentID() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentID(e2nodeComponentID *e2ap_ies.E2NodeComponentId) ([]byte, error) {
	e2nodeComponentIDCP, err := newE2nodeComponentID(e2nodeComponentID)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentID() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentID, unsafe.Pointer(e2nodeComponentIDCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentID() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentID(bytes []byte) (*e2ap_ies.E2NodeComponentId, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentID((*C.E2nodeComponentID_t)(unsafePtr))
}

func perDecodeE2nodeComponentID(bytes []byte) (*e2ap_ies.E2NodeComponentId, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentID)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentID((*C.E2nodeComponentID_t)(unsafePtr))
}

func newE2nodeComponentID(e2nodeComponentID *e2ap_ies.E2NodeComponentId) (*C.E2nodeComponentID_t, error) {

	var pr C.E2nodeComponentID_PR
	choiceC := [8]byte{}

	switch choice := e2nodeComponentID.E2NodeComponentId.(type) {
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeNg:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeNG

		im, err := newE2nodeComponentInterfaceNG(choice.E2NodeComponentInterfaceTypeNg)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceNG() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeXn:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeXn

		im, err := newE2nodeComponentInterfaceXn(choice.E2NodeComponentInterfaceTypeXn)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceXn() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeE1

		im, err := newE2nodeComponentInterfaceE1(choice.E2NodeComponentInterfaceTypeE1)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceE1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeF1

		im, err := newE2nodeComponentInterfaceF1(choice.E2NodeComponentInterfaceTypeF1)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceF1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeW1

		im, err := newE2nodeComponentInterfaceW1(choice.E2NodeComponentInterfaceTypeW1)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceW1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeS1:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeS1

		im, err := newE2nodeComponentInterfaceS1(choice.E2NodeComponentInterfaceTypeS1)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceS1() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeX2:
		pr = C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeX2

		im, err := newE2nodeComponentInterfaceX2(choice.E2NodeComponentInterfaceTypeX2)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentInterfaceX2() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2nodeComponentID() %T not yet implemented", choice)
	}

	e2nodeComponentIDC := C.E2nodeComponentID_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2nodeComponentIDC, nil
}

func decodeE2nodeComponentID(e2nodeComponentIDC *C.E2nodeComponentID_t) (*e2ap_ies.E2NodeComponentId, error) {

	e2nodeComponentID := new(e2ap_ies.E2NodeComponentId)

	switch e2nodeComponentIDC.present {
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeNG:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceNGBytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceNGBytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeNg{
			E2NodeComponentInterfaceTypeNg: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeXn:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceXnBytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceXnBytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeXn{
			E2NodeComponentInterfaceTypeXn: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeE1:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceE1Bytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceE1Bytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1{
			E2NodeComponentInterfaceTypeE1: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeF1:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceF1Bytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceF1Bytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1{
			E2NodeComponentInterfaceTypeF1: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeW1:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceW1Bytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceW1Bytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{
			E2NodeComponentInterfaceTypeW1: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeS1:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceS1Bytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceS1Bytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeS1{
			E2NodeComponentInterfaceTypeS1: e2nodeComponentIDstructC,
		}
	case C.E2nodeComponentID_PR_e2nodeComponentInterfaceTypeX2:
		e2nodeComponentIDstructC, err := decodeE2nodeComponentInterfaceX2Bytes(e2nodeComponentIDC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentInterfaceX2Bytes() %s", err.Error())
		}
		e2nodeComponentID.E2NodeComponentId = &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeX2{
			E2NodeComponentInterfaceTypeX2: e2nodeComponentIDstructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2nodeComponentID() %v not yet implemented", e2nodeComponentIDC.present)
	}

	return e2nodeComponentID, nil
}
