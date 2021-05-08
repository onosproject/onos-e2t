// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "SuccessfulOutcome.h"
//#include "ProtocolIE-Field.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"unsafe"
)

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

	} else if pc := so.GetProcedureCode().GetRicSubscription(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_RICsubscriptionResponse
		pcC = C.ProcedureCode_id_RICsubscription
		critC = C.long(C.Criticality_reject)
		e2srC, err := newRicSubscriptionResponse(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))

	} else if pc := so.GetProcedureCode().GetRicSubscriptionDelete(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_RICsubscriptionDeleteResponse
		pcC = C.ProcedureCode_id_RICsubscriptionDelete
		critC = C.long(C.Criticality_reject)
		e2srC, err := newRicSubscriptionDeleteResponse(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))

	} else if pc := so.GetProcedureCode().GetRicControl(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_RICcontrolAcknowledge
		pcC = C.ProcedureCode_id_RICcontrol
		critC = C.long(C.Criticality_reject)
		e2srC, err := newRicControlAcknowledge(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))

	} else if pc := so.GetProcedureCode().GetReset_(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_ResetResponse
		pcC = C.ProcedureCode_id_Reset
		critC = C.long(C.Criticality_reject)
		e2srC, err := newResetResponse(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))

	} else if pc := so.GetProcedureCode().GetRicServiceUpdate(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_RICserviceUpdateAcknowledge
		pcC = C.ProcedureCode_id_RICserviceUpdate
		critC = C.long(C.Criticality_reject)
		e2srC, err := newRicServiceUpdateAcknowledge(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))

	} else if pc := so.GetProcedureCode().GetE2NodeConfigurationUpdate(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_E2nodeConfigurationUpdateAcknowledge
		pcC = C.ProcedureCode_id_E2nodeConfigurationUpdate
		critC = C.long(C.Criticality_reject)
		e2srC, err := newE2nodeConfigurationUpdateAcknowledge(pc.GetSuccessfulOutcome())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2srC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2srC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2srC.protocolIEs.list.size))

	} else if pc := so.GetProcedureCode().GetE2ConnectionUpdate(); pc != nil &&
		pc.GetSuccessfulOutcome() != nil {

		presentC = C.SuccessfulOutcome__value_PR_E2connectionUpdateAcknowledge
		pcC = C.ProcedureCode_id_E2connectionUpdate
		critC = C.long(C.Criticality_reject)
		e2srC, err := newE2connectionUpdateAcknowledge(pc.GetSuccessfulOutcome())
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

	listArrayAddr := unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(successC.value.choice[0:8])))
	count := C.int(binary.LittleEndian.Uint32(successC.value.choice[8:12]))
	size := C.int(binary.LittleEndian.Uint32(successC.value.choice[12:16]))

	switch successC.value.present {
	case C.SuccessfulOutcome__value_PR_RICsubscriptionResponse:
		rsrespC := C.RICsubscriptionResponse_t{
			protocolIEs: C.ProtocolIE_Container_1710P1_t{
				list: C.struct___135{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionResponse_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
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
					Value: int32(v1beta2.ProcedureCodeIDRICsubscription),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}
	case C.SuccessfulOutcome__value_PR_E2setupResponse:
		e2SrC := C.E2setupResponse_t{
			protocolIEs: C.ProtocolIE_Container_1710P12_t{
				list: C.struct___134{ // TODO: tie this down with a predictable name
					array: (**C.E2setupResponseIEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		e2Sr, err := decodeE2setupResponse(&e2SrC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2Setup: &e2appdudescriptions.E2Setup{
				SuccessfulOutcome: e2Sr,
				ProcedureCode: &e2ap_constants.IdE2Setup{
					Value: int32(v1beta2.ProcedureCodeIDE2setup),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}
	case C.SuccessfulOutcome__value_PR_RICsubscriptionDeleteResponse:
		rsdrC := C.RICsubscriptionDeleteResponse_t{
			protocolIEs: C.ProtocolIE_Container_1710P4_t{
				list: C.struct___138{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionDeleteResponse_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionResponse_t %+v\n %+v\n", successC, rsrespC)
		rsdr, err := decodeRicSubscriptionDeleteResponse(&rsdrC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicSubscriptionDelete: &e2appdudescriptions.RicSubscriptionDelete{
				SuccessfulOutcome: rsdr,
				ProcedureCode: &e2ap_constants.IdRicsubscriptionDelete{
					Value: int32(v1beta2.ProcedureCodeIDRICsubscriptionDelete),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}
	case C.SuccessfulOutcome__value_PR_RICcontrolAcknowledge:
		rcaC := C.RICcontrolAcknowledge_t{
			protocolIEs: C.ProtocolIE_Container_1710P8_t{
				list: C.struct___139{ // TODO: tie this down with a predictable name
					array: (**C.RICcontrolAcknowledge_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionResponse_t %+v\n %+v\n", successC, rsrespC)
		rca, err := decodeRicControlAcknowledge(&rcaC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicControl: &e2appdudescriptions.RicControl{
				SuccessfulOutcome: rca,
				ProcedureCode: &e2ap_constants.IdRiccontrol{
					Value: int32(v1beta2.ProcedureCodeIDRICcontrol),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}
	case C.SuccessfulOutcome__value_PR_ResetResponse:
		rrC := C.ResetResponse_t{
			protocolIEs: C.ProtocolIE_Container_1710P21_t{
				list: C.struct___136{ // TODO: tie this down with a predictable name
					array: (**C.ResetResponseIEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("ResetResponse_t %+v\n %+v\n", successC, rsrespC)
		rr, err := decodeResetResponse(&rrC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			Reset_: &e2appdudescriptions.Reset{
				SuccessfulOutcome: rr,
				ProcedureCode: &e2ap_constants.IdReset{
					Value: int32(v1beta2.ProcedureCodeIDReset),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.SuccessfulOutcome__value_PR_RICserviceUpdateAcknowledge:
		rsuaC := C.RICserviceUpdateAcknowledge_t{
			protocolIEs: C.ProtocolIE_Container_1710P23_t{
				list: C.struct___137{ // TODO: tie this down with a predictable name
					array: (**C.RICserviceUpdateAcknowledge_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("ResetResponse_t %+v\n %+v\n", successC, rsrespC)
		rsua, err := decodeRicServiceUpdateAcknowledge(&rsuaC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
				SuccessfulOutcome: rsua,
				ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
					Value: int32(v1beta2.ProcedureCodeIDRICserviceUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.SuccessfulOutcome__value_PR_E2nodeConfigurationUpdateAcknowledge:
		e2ncuaC := C.E2nodeConfigurationUpdateAcknowledge_t{
			protocolIEs: C.ProtocolIE_Container_1710P18_t{
				list: C.struct___133{ // TODO: tie this down with a predictable name
					array: (**C.E2nodeConfigurationUpdateAcknowledge_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("ResetResponse_t %+v\n %+v\n", successC, rsrespC)
		e2ncua, err := decodeE2nodeConfigurationUpdateAcknowledge(&e2ncuaC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
				SuccessfulOutcome: e2ncua,
				ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
					Value: int32(v1beta2.ProcedureCodeIDE2nodeConfigurationUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.SuccessfulOutcome__value_PR_E2connectionUpdateAcknowledge:
		e2cuaC := C.E2connectionUpdateAcknowledge_t{
			protocolIEs: C.ProtocolIE_Container_1710P15_t{
				list: C.struct___132{ // TODO: tie this down with a predictable name
					array: (**C.E2connectionUpdateAck_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("ResetResponse_t %+v\n %+v\n", successC, rsrespC)
		e2cua, err := decodeE2connectionUpdateAcknowledge(&e2cuaC)
		if err != nil {
			return nil, err
		}
		successfulOutcome.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
				SuccessfulOutcome: e2cua,
				ProcedureCode: &e2ap_constants.IdE2ConnectionUpdate{
					Value: int32(v1beta2.ProcedureCodeIDE2connectionUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	default:
		return nil, fmt.Errorf("decodeSuccessfulOutcome() %v not yet implemented", successC.value.present)
	}

	return successfulOutcome, nil
}
