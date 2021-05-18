// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "E2nodeComponentConfigUpdate.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	"unsafe"
)

func xerEncodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate *e2ap_ies.E2NodeComponentConfigUpdate) ([]byte, error) {
	e2nodeComponentConfigUpdateCP, err := newE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdate() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_E2nodeComponentConfigUpdate, unsafe.Pointer(e2nodeComponentConfigUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeE2nodeComponentConfigUpdate() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate *e2ap_ies.E2NodeComponentConfigUpdate) ([]byte, error) {
	e2nodeComponentConfigUpdateCP, err := newE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate)
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdate() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2nodeComponentConfigUpdate, unsafe.Pointer(e2nodeComponentConfigUpdateCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeE2nodeComponentConfigUpdate() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeE2nodeComponentConfigUpdate(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdate, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2nodeComponentConfigUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2nodeComponentConfigUpdate((*C.E2nodeComponentConfigUpdate_t)(unsafePtr))
}

func perDecodeE2nodeComponentConfigUpdate(bytes []byte) (*e2ap_ies.E2NodeComponentConfigUpdate, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2nodeComponentConfigUpdate)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeE2nodeComponentConfigUpdate((*C.E2nodeComponentConfigUpdate_t)(unsafePtr))
}

func newE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdate *e2ap_ies.E2NodeComponentConfigUpdate) (*C.E2nodeComponentConfigUpdate_t, error) {

	var pr C.E2nodeComponentConfigUpdate_PR //ToDo - verify correctness of the name
	choiceC := [8]byte{}                    //ToDo - Check if number of bytes is sufficient

	switch choice := e2nodeComponentConfigUpdate.E2NodeComponentConfigUpdate.(type) {
	case *e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate:
		pr = C.E2nodeComponentConfigUpdate_PR_gNBconfigUpdate //ToDo - Check if it's correct PR's name

		im, err := newE2nodeComponentConfigUpdateGnb(choice.GNbconfigUpdate)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentConfigUpdateGnb() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentConfigUpdate_EnGNbconfigUpdate:
		pr = C.E2nodeComponentConfigUpdate_PR_en_gNBconfigUpdate //ToDo - Check if it's correct PR's name

		im, err := newE2nodeComponentConfigUpdateEngNb(choice.EnGNbconfigUpdate)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentConfigUpdateEngNb() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentConfigUpdate_NgENbconfigUpdate:
		pr = C.E2nodeComponentConfigUpdate_PR_ng_eNBconfigUpdate //ToDo - Check if it's correct PR's name

		im, err := newE2nodeComponentConfigUpdateNgeNb(choice.NgENbconfigUpdate)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentConfigUpdateNgeNb() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ap_ies.E2NodeComponentConfigUpdate_ENbconfigUpdate:
		pr = C.E2nodeComponentConfigUpdate_PR_eNBconfigUpdate //ToDo - Check if it's correct PR's name

		im, err := newE2nodeComponentConfigUpdateEnb(choice.ENbconfigUpdate)
		if err != nil {
			return nil, fmt.Errorf("newE2nodeComponentConfigUpdateEnb() %s", err.Error())
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	default:
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdate() %T not yet implemented", choice)
	}

	e2nodeComponentConfigUpdateC := C.E2nodeComponentConfigUpdate_t{
		present: pr,
		choice:  choiceC,
	}

	return &e2nodeComponentConfigUpdateC, nil
}

func decodeE2nodeComponentConfigUpdate(e2nodeComponentConfigUpdateC *C.E2nodeComponentConfigUpdate_t) (*e2ap_ies.E2NodeComponentConfigUpdate, error) {

	e2nodeComponentConfigUpdate := new(e2ap_ies.E2NodeComponentConfigUpdate)

	switch e2nodeComponentConfigUpdateC.present {
	case C.E2nodeComponentConfigUpdate_PR_gNBconfigUpdate:
		e2nodeComponentConfigUpdatestructC, err := decodeE2nodeComponentConfigUpdateGnbBytes(e2nodeComponentConfigUpdateC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateGnbBytes() %s", err.Error())
		}
		e2nodeComponentConfigUpdate.E2NodeComponentConfigUpdate = &e2ap_ies.E2NodeComponentConfigUpdate_GNbconfigUpdate{
			GNbconfigUpdate: e2nodeComponentConfigUpdatestructC,
		}
	case C.E2nodeComponentConfigUpdate_PR_en_gNBconfigUpdate:
		e2nodeComponentConfigUpdatestructC, err := decodeE2nodeComponentConfigUpdateEngNbBytes(e2nodeComponentConfigUpdateC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateEngNbBytes() %s", err.Error())
		}
		e2nodeComponentConfigUpdate.E2NodeComponentConfigUpdate = &e2ap_ies.E2NodeComponentConfigUpdate_EnGNbconfigUpdate{
			EnGNbconfigUpdate: e2nodeComponentConfigUpdatestructC,
		}
	case C.E2nodeComponentConfigUpdate_PR_ng_eNBconfigUpdate:
		e2nodeComponentConfigUpdatestructC, err := decodeE2nodeComponentConfigUpdateNgeNbBytes(e2nodeComponentConfigUpdateC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateNgeNbBytes() %s", err.Error())
		}
		e2nodeComponentConfigUpdate.E2NodeComponentConfigUpdate = &e2ap_ies.E2NodeComponentConfigUpdate_NgENbconfigUpdate{
			NgENbconfigUpdate: e2nodeComponentConfigUpdatestructC,
		}
	case C.E2nodeComponentConfigUpdate_PR_eNBconfigUpdate:
		e2nodeComponentConfigUpdatestructC, err := decodeE2nodeComponentConfigUpdateEnbBytes(e2nodeComponentConfigUpdateC.choice) //ToDo - Verify if decodeSmthBytes function exists
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateEnbBytes() %s", err.Error())
		}
		e2nodeComponentConfigUpdate.E2NodeComponentConfigUpdate = &e2ap_ies.E2NodeComponentConfigUpdate_ENbconfigUpdate{
			ENbconfigUpdate: e2nodeComponentConfigUpdatestructC,
		}
	default:
		return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdate() %v not yet implemented", e2nodeComponentConfigUpdateC.present)
	}

	return e2nodeComponentConfigUpdate, nil
}
