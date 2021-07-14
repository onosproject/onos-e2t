// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ENB-ID-Choice.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeEnbIDChoice(enbIDChoice *e2ap_ies.EnbIdChoice) ([]byte, error) {
	enbIDChoiceCP, err := newEnbIDChoice(enbIDChoice)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbIDChoice() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_ENB_ID_Choice, unsafe.Pointer(enbIDChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEnbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEnbIDChoice(enbIDChoice *e2ap_ies.EnbIdChoice) ([]byte, error) {
	enbIDChoiceCP, err := newEnbIDChoice(enbIDChoice)
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbIDChoice() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_ENB_ID_Choice, unsafe.Pointer(enbIDChoiceCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEnbIDChoice() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeEnbIDChoice(bytes []byte) (*e2ap_ies.EnbIdChoice, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_ENB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeEnbIDChoice((*C.ENB_ID_Choice_t)(unsafePtr))
}

func perDecodeEnbIDChoice(bytes []byte) (*e2ap_ies.EnbIdChoice, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_ENB_ID_Choice)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeEnbIDChoice((*C.ENB_ID_Choice_t)(unsafePtr))
}

func newEnbIDChoice(enbIDChoice *e2ap_ies.EnbIdChoice) (*C.ENB_ID_Choice_t, error) {

	var pr C.ENB_ID_Choice_PR
	choiceC := [48]byte{}

	switch choice := enbIDChoice.EnbIdChoice.(type) {
	case *e2ap_ies.EnbIdChoice_EnbIdMacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_macro

		bsC, err := newBitString(choice.EnbIdMacro)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	case *e2ap_ies.EnbIdChoice_EnbIdShortmacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_shortmacro

		bsC, err := newBitString(choice.EnbIdShortmacro)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	case *e2ap_ies.EnbIdChoice_EnbIdLongmacro:
		pr = C.ENB_ID_Choice_PR_enb_ID_longmacro

		bsC, err := newBitString(choice.EnbIdLongmacro)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("newEnbIDChoice() %T not yet implemented", choice)
	}

	enbIDChoiceC := C.ENB_ID_Choice_t{
		present: pr,
		choice:  choiceC,
	}

	return &enbIDChoiceC, nil
}

func decodeEnbIDChoice(enbIDChoiceC *C.ENB_ID_Choice_t) (*e2ap_ies.EnbIdChoice, error) {

	enbIDChoice := new(e2ap_ies.EnbIdChoice)

	switch enbIDChoiceC.present {
	case C.ENB_ID_Choice_PR_enb_ID_macro:
		enbIDChoicestructC := newBitStringFromArray(enbIDChoiceC.choice)
		enbID, err := decodeBitString(enbIDChoicestructC)
		if err != nil{
			return nil, err
		}
		enbIDChoice.EnbIdChoice = &e2ap_ies.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: enbID,
		}
	case C.ENB_ID_Choice_PR_enb_ID_shortmacro:
		enbIDChoicestructC := newBitStringFromArray(enbIDChoiceC.choice)
		enbID, err := decodeBitString(enbIDChoicestructC)
		if err != nil{
			return nil, err
		}
		enbIDChoice.EnbIdChoice = &e2ap_ies.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: enbID,
		}
	case C.ENB_ID_Choice_PR_enb_ID_longmacro:
		enbIDChoicestructC := newBitStringFromArray(enbIDChoiceC.choice)
		enbID, err := decodeBitString(enbIDChoicestructC)
		if err != nil{
			return nil, err
		}
		enbIDChoice.EnbIdChoice = &e2ap_ies.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: enbID,
		}
	default:
		return nil, fmt.Errorf("decodeEnbIDChoice() %v not yet implemented", enbIDChoiceC.present)
	}

	return enbIDChoice, nil
}
