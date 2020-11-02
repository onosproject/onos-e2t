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
//#include "SuccessfulOutcome.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-constants"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// Deprecated: Do not use.
func newSuccessfulOutcomeOld(so *e2ctypes.SuccessfulOutcomeT) (*C.SuccessfulOutcome_t, error) {
	pcC, err := procedureCodeToC(so.GetProcedureCode())
	if err != nil {
		return nil, err
	}

	critC, err := criticalityToCOld(so.GetCriticality())
	if err != nil {
		return nil, err
	}

	soC, err := newSuccessfulOutcomeValueOld(so)
	if err != nil {
		return nil, err
	}

	imC := C.SuccessfulOutcome_t{
		procedureCode: pcC,
		criticality:   critC,
		value:         *soC,
	}

	return &imC, nil
}

// Deprecated: Do not use.
func newSuccessfulOutcomeValueOld(so *e2ctypes.SuccessfulOutcomeT) (*C.struct_SuccessfulOutcome__value, error) {
	var presentC C.SuccessfulOutcome__value_PR
	choiceC := [72]byte{} // The size of the SuccessfulOutcome__value_u union
	switch choice := so.GetChoice().(type) {
	case *e2ctypes.SuccessfulOutcomeT_E2SetupResponse:
		presentC = C.SuccessfulOutcome__value_PR_E2setupResponse

		e2srC, err := newE2setupResponseOld(choice.E2SetupResponse)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Protocol IEs %v %v %v\n", e2srC.protocolIEs.list.array, e2srC.protocolIEs.list.count, e2srC.protocolIEs.list.size)
		// Now copy the e2srC over in to the choice byte by byte - the union is [72]byte
		// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))
	default:
		return nil, fmt.Errorf("newSuccessfulOutcomeValueOld() %T not yet implemented", choice)
	}

	soC := C.struct_SuccessfulOutcome__value{
		present: presentC,
		choice:  choiceC,
	}

	return &soC, nil
}

// Deprecated: Do not use.
func decodeSuccessfulOutcomeOld(successC *C.SuccessfulOutcome_t) (*e2ctypes.SuccessfulOutcomeT, error) {
	successfulOutcome := e2ctypes.SuccessfulOutcomeT{
		ProcedureCode: e2ctypes.ProcedureCodeT(successC.procedureCode),
		Criticality:   e2ctypes.CriticalityT(successC.criticality),
	}
	listArrayAddr := successC.value.choice[0:8]

	switch successC.value.present {
	case C.SuccessfulOutcome__value_PR_RICsubscriptionResponse:
		rsrespC := C.RICsubscriptionResponse_t{
			protocolIEs: C.ProtocolIE_Container_1544P1_t{
				list: C.struct___48{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionResponse_IEs_t)(unsafe.Pointer(&listArrayAddr[0])),
					count: C.int(binary.LittleEndian.Uint32(successC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(successC.value.choice[12:16])),
				},
			},
		}
		//fmt.Printf("RICsubscriptionResponse_t %+v\n %+v\n", successC, rsrespC)
		rsresp, err := decodeRicSubscriptionResponseOld(&rsrespC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.Choice = &e2ctypes.SuccessfulOutcomeT_RICsubscriptionResponse{
			RICsubscriptionResponse: rsresp,
		}
	default:
		return nil, fmt.Errorf("decodeInitiatingMessageOld() %v not yet implemented", successC.value.present)
	}

	return &successfulOutcome, nil
}

func newSuccessfulOutcome(so *e2appdudescriptions.SuccessfulOutcome) (*C.SuccessfulOutcome_t, error) {
	var presentC C.SuccessfulOutcome__value_PR
	var pcC C.ProcedureCode_t
	var critC C.Criticality_t
	choiceC := [72]byte{} // The size of the SuccessfulOutcome__value_u union
	if pc := so.GetProcedureCode().GetE2Setup(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_E2setupResponse
		pcC = C.ProcedureCode_id_E2setup
		critC = C.long(C.Criticality_reject)
		e2srC, err := newE2setupResponse(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))
	} else {
		return nil, fmt.Errorf("newSuccessfulOutcomeValue type not yet implemented")
	}

	sovC := C.struct_SuccessfulOutcome__value{
		present: presentC,
		choice:  choiceC,
	}

	soC := C.SuccessfulOutcome_t{
		procedureCode: pcC,
		criticality:   critC,
		value:         sovC,
	}

	return &soC, nil
}

func decodeSuccessfulOutcome(successC *C.SuccessfulOutcome_t) (*e2appdudescriptions.SuccessfulOutcome, error) {
	successfulOutcome := new(e2appdudescriptions.SuccessfulOutcome)

	listArrayAddr := successC.value.choice[0:8]

	switch successC.value.present {
	case C.SuccessfulOutcome__value_PR_RICsubscriptionResponse:
		rsrespC := C.RICsubscriptionResponse_t{
			protocolIEs: C.ProtocolIE_Container_1544P1_t{
				list: C.struct___48{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionResponse_IEs_t)(unsafe.Pointer(&listArrayAddr[0])),
					count: C.int(binary.LittleEndian.Uint32(successC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(successC.value.choice[12:16])),
				},
			},
		}
		//fmt.Printf("RICsubscriptionResponse_t %+v\n %+v\n", successC, rsrespC)
		rsresp, err := decodeRicSubscriptionResponse(&rsrespC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicSubscription: &e2appdudescriptions.RicSubscription{
				SuccessfulOutcome: rsresp,
				ProcedureCode: &e2ap_constants.IdRicsubscription{
					Value: int32(v1beta1.ProcedureCodeIDRICsubscription),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}
	default:
		return nil, fmt.Errorf("decodeInitiatingMessageOld() %v not yet implemented", successC.value.present)
	}

	return successfulOutcome, nil
}
