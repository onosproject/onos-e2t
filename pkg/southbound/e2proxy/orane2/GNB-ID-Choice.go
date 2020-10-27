// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GNB-ID-Choice.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeGnbID - used only in tests
// Deprecated: Do not use.
func XerEncodeGnbID(gnbID *e2ctypes.GNB_ID_ChoiceT) ([]byte, error) {
	gnbIDC, err := newGnbIDChoice(gnbID)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_GNB_ID_Choice, unsafe.Pointer(gnbIDC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeGnbID - used only in tests
// Deprecated: Do not use.
func PerEncodeGnbID(gnbID *e2ctypes.GNB_ID_ChoiceT) ([]byte, error) {
	gnbIDC, err := newGnbIDChoice(gnbID)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_GNB_ID_Choice, unsafe.Pointer(gnbIDC))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// Deprecated: Do not use.
func newGnbIDChoice(gnbIDCh *e2ctypes.GNB_ID_ChoiceT) (*C.GNB_ID_Choice_t, error) {
	var pr C.GNB_ID_Choice_PR

	choiceC := [48]byte{}

	switch choice := gnbIDCh.GetChoice().(type) {
	case *e2ctypes.GNB_ID_ChoiceT_Gnb_ID:
		pr = C.GNB_ID_Choice_PR_gnb_ID
		bsC := newBitString(choice.Gnb_ID)
		fmt.Printf("gNB ID %v %v %v %v\n", bsC, unsafe.Sizeof(bsC.size), unsafe.Sizeof(bsC.bits_unused), *bsC.buf)

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(bsC.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(bsC.size))
		binary.LittleEndian.PutUint32(choiceC[16:], uint32(bsC.bits_unused))
	default:
		return nil, fmt.Errorf("undhanled type %v", choice)
	}

	gnbChC := C.GNB_ID_Choice_t{
		present: pr,
		choice:  choiceC,
	}
	return &gnbChC, nil
}

// Deprecated: Do not use.
func decodeGnbIDChoice(gnbIDC *C.GNB_ID_Choice_t) (*e2ctypes.GNB_ID_ChoiceT, error) {
	result := new(e2ctypes.GNB_ID_ChoiceT)

	switch gnbIDC.present {
	case C.GNB_ID_Choice_PR_gnb_ID:
		//fmt.Printf("GNB_ID_Choice_t %+v\n", gnbIDC.choice)
		result.Choice = &e2ctypes.GNB_ID_ChoiceT_Gnb_ID{
			Gnb_ID: decodeBitString(gnbIDC.choice),
		}
	default:
		return nil, fmt.Errorf("decodeGnbIDChoice() %v not yet implemented", gnbIDC.present)
	}

	return result, nil
}
