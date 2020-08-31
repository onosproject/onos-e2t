// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "RICindication.h"
//#include "Criticality.h"
//#include "InitiatingMessage.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

func newInitiatingMessage(im *e2ctypes.InitiatingMessageT) (*C.InitiatingMessage_t, error) {

	pcC, err := procedureCodeToC(im.GetProcedureCode())
	if err != nil {
		return nil, err
	}

	critC, err := criticalityToC(im.GetCriticality())
	if err != nil {
		return nil, err
	}

	imvC, err := newInitiatingMessageValue(im)
	if err != nil {
		return nil, err
	}

	imC := C.InitiatingMessage_t{
		procedureCode: pcC,
		criticality:   critC,
		value:         *imvC,
	}

	return &imC, nil
}

func newInitiatingMessageValue(im *e2ctypes.InitiatingMessageT) (*C.struct_InitiatingMessage__value, error) {

	var presentC C.InitiatingMessage__value_PR
	choiceC := [72]byte{} // The size of the InitiatingMessage__value_u union
	switch choice := im.GetChoice().(type) {
	case *e2ctypes.InitiatingMessageT_RICsubscriptionRequest:
		presentC = C.InitiatingMessage__value_PR_RICsubscriptionRequest

		rsrC, err := newRICsubscriptionRequest(choice.RICsubscriptionRequest)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsrC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsrC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsrC.protocolIEs.list.size))
	case *e2ctypes.InitiatingMessageT_RICindication:
		presentC = C.InitiatingMessage__value_PR_RICindication
		riC, err := newRicIndication(choice.RICindication)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Protocol IEs %v %v %v\n", riC.protocolIEs.list.array, riC.protocolIEs.list.count, riC.protocolIEs.list.size)

		// Now copy the riC over in to the choice byte by byte - the union is [72]byte
		// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(riC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(riC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(riC.protocolIEs.list.size))

	case *e2ctypes.InitiatingMessageT_E2SetupRequest:
		presentC = C.InitiatingMessage__value_PR_E2setupRequest
		esC, err := newE2SetupRequest(choice.E2SetupRequest)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Protocol IEs %v %v %v %d %d\n", esC.protocolIEs.list.array,
			esC.protocolIEs.list.count, esC.protocolIEs.list.size,
			unsafe.Sizeof(esC.protocolIEs.list.count), unsafe.Sizeof(esC.protocolIEs.list.size))

		// Now copy the esC over in to the choice byte by byte - the union is [72]byte
		// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(esC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(esC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(esC.protocolIEs.list.size))
	default:
		return nil, fmt.Errorf("newInitiatingMessageValue %T not yet implemented", choice)
	}

	imvC := C.struct_InitiatingMessage__value{
		present: presentC,
		choice:  choiceC,
	}

	return &imvC, nil
}

func decodeInitiatingMessage(initMsgC *C.InitiatingMessage_t) (*e2ctypes.InitiatingMessageT, error) {

	initiatingMessage := e2ctypes.InitiatingMessageT{
		ProcedureCode: e2ctypes.ProcedureCodeT(initMsgC.procedureCode),
		Criticality:   e2ctypes.CriticalityT(initMsgC.criticality),
	}
	listArrayAddr := initMsgC.value.choice[0:8]

	switch initMsgC.value.present {
	case C.InitiatingMessage__value_PR_ErrorIndication:
		// TODO change this to be like E2SetupRequest below - at present it cannot extract array
		errIndC := *(**C.ErrorIndication_t)(unsafe.Pointer(&initMsgC.value.choice[0]))
		errInd, err := decodeErrorIndication(errIndC)
		if err != nil {
			return nil, err
		}
		initiatingMessage.Choice = &e2ctypes.InitiatingMessageT_ErrorIndication{
			ErrorIndication: errInd,
		}
	case C.InitiatingMessage__value_PR_E2setupRequest:
		e2srC := *(**C.E2setupRequestIEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		esC := C.E2setupRequest_t{
			protocolIEs: C.ProtocolIE_Container_1544P11_t{
				list: C.struct___32{ // TODO: tie this down with a predictable name
					array: (**C.E2setupRequestIEs_t)(unsafe.Pointer(e2srC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		//fmt.Printf("E2SetupRequestC %+v\n %+v\n", initMsgC, srC)
		e2sr, err := decodeE2setupRequest(&esC)
		if err != nil {
			return nil, err
		}
		initiatingMessage.Choice = &e2ctypes.InitiatingMessageT_E2SetupRequest{
			E2SetupRequest: e2sr,
		}
	case C.InitiatingMessage__value_PR_RICsubscriptionRequest:
		ricsrC := *(**C.RICsubscriptionRequest_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		srC := C.RICsubscriptionRequest_t{
			protocolIEs: C.ProtocolIE_Container_1544P0_t{
				list: C.struct___40{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionRequest_IEs_t)(unsafe.Pointer(ricsrC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("RICsubscriptionRequest_t %+v\n %+v\n", initMsgC, srC)
		sr, err := decodeRicSubscriptionRequest(&srC)
		if err != nil {
			return nil, err
		}
		initiatingMessage.Choice = &e2ctypes.InitiatingMessageT_RICsubscriptionRequest{
			RICsubscriptionRequest: sr,
		}
	default:
		return nil, fmt.Errorf("decodeInitiatingMessage() %v not yet implemented", initMsgC.value.present)
	}

	return &initiatingMessage, nil
}
