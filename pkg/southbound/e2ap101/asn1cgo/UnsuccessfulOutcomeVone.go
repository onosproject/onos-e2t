// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "UnsuccessfulOutcomeVone.h"
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

func newUnsuccessfulOutcome(uso *e2appdudescriptions.UnsuccessfulOutcome) (*C.UnsuccessfulOutcomeVone_t, error) {
	var presentC C.UnsuccessfulOutcomeVone__value_PR
	var pcC C.ProcedureCodeVone_t
	var critC C.CriticalityVone_t
	choiceC := [72]byte{} // The size of the UnsuccessfulOutcome__value_u union
	if pc := uso.GetProcedureCode().GetRicSubscriptionDelete(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_RICsubscriptionDeleteFailureVone
		pcC = C.ProcedureCodeVone_id_RICsubscriptionDeleteVone
		critC = C.long(C.CriticalityVone_reject)
		rsdfC, err := newRicSubscriptionDeleteFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsdfC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsdfC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsdfC.protocolIEs.list.size))

	} else if pc := uso.GetProcedureCode().GetRicControl(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_RICcontrolFailureVone
		pcC = C.ProcedureCodeVone_id_RICcontrolVone
		critC = C.long(C.CriticalityVone_reject)
		rsfC, err := newRicControlFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsfC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsfC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsfC.protocolIEs.list.size))

	} else if pc := uso.GetProcedureCode().GetRicSubscription(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_RICsubscriptionFailureVone
		pcC = C.ProcedureCodeVone_id_RICsubscriptionVone
		critC = C.long(C.CriticalityVone_reject)
		rsfC, err := newRicSubscriptionFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsfC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsfC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsfC.protocolIEs.list.size))

	} else if pc := uso.GetProcedureCode().GetE2Setup(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_E2setupFailureVone
		pcC = C.ProcedureCodeVone_id_E2setupVone
		critC = C.long(C.CriticalityVone_reject)
		e2sfC, err := newE2setupFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2sfC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2sfC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2sfC.protocolIEs.list.size))

	} else if pc := uso.GetProcedureCode().GetRicServiceUpdate(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_RICserviceUpdateFailureVone
		pcC = C.ProcedureCodeVone_id_RICserviceUpdateVone
		critC = C.long(C.CriticalityVone_reject)
		rsufC, err := newRicServiceUpdateFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsufC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsufC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsufC.protocolIEs.list.size))

	} else if pc := uso.GetProcedureCode().GetE2NodeConfigurationUpdate(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_E2nodeConfigurationUpdateFailureVone
		pcC = C.ProcedureCodeVone_id_E2nodeConfigurationUpdateVone
		critC = C.long(C.CriticalityVone_reject)
		rsufC, err := newE2nodeConfigurationUpdateFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsufC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsufC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsufC.protocolIEs.list.size))

	} else if pc := uso.GetProcedureCode().GetE2ConnectionUpdate(); pc != nil &&
		pc.GetUnsuccessfulOutcome() != nil {

		presentC = C.UnsuccessfulOutcomeVone__value_PR_E2connectionUpdateFailureVone
		pcC = C.ProcedureCodeVone_id_E2connectionUpdateVone
		critC = C.long(C.CriticalityVone_reject)
		rsufC, err := newE2connectionUpdateFailure(pc.GetUnsuccessfulOutcome())
		if err != nil {
			return nil, err
		}

		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsufC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsufC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsufC.protocolIEs.list.size))

	} else {
		return nil, fmt.Errorf("newUnsuccessfulOutcomeValue type not yet implemented")
	}

	soC := C.UnsuccessfulOutcomeVone_t{
		procedureCode: pcC,
		criticality:   critC,
		value: C.struct_UnsuccessfulOutcomeVone__value{
			present: presentC,
			choice:  choiceC,
		},
	}

	return &soC, nil
}

func decodeUnsuccessfulOutcome(failureC *C.UnsuccessfulOutcomeVone_t) (*e2appdudescriptions.UnsuccessfulOutcome, error) {
	uso := new(e2appdudescriptions.UnsuccessfulOutcome)

	listArrayAddr := unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(failureC.value.choice[0:8])))
	count := C.int(binary.LittleEndian.Uint32(failureC.value.choice[8:12]))
	size := C.int(binary.LittleEndian.Uint32(failureC.value.choice[12:16]))

	switch failureC.value.present {
	case C.UnsuccessfulOutcomeVone__value_PR_RICsubscriptionFailureVone:
		rsfC := C.RICsubscriptionFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P2_t{
				list: C.struct___146{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionFailureVone_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionDeleteFailure %+v\n %+v\n", failureC, rsfC)
		rsf, err := decodeRicSubscriptionFailure(&rsfC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicSubscription: &e2appdudescriptions.RicSubscription{
				UnsuccessfulOutcome: rsf,
				ProcedureCode: &e2ap_constants.IdRicsubscription{
					Value: int32(v1beta2.ProcedureCodeIDRICsubscription),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.UnsuccessfulOutcomeVone__value_PR_RICcontrolFailureVone:
		rcfC := C.RICcontrolFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P9_t{
				list: C.struct___148{ // TODO: tie this down with a predictable name
					array: (**C.RICcontrolFailureVone_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICcontrolFailure %+v\n %+v\n", failureC, rsfC)
		rcf, err := decodeRicControlFailure(&rcfC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicControl: &e2appdudescriptions.RicControl{
				UnsuccessfulOutcome: rcf,
				ProcedureCode: &e2ap_constants.IdRiccontrol{
					Value: int32(v1beta2.ProcedureCodeIDRICcontrol),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.UnsuccessfulOutcomeVone__value_PR_RICsubscriptionDeleteFailureVone:
		rsdfC := C.RICsubscriptionDeleteFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P5_t{
				list: C.struct___147{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionDeleteFailureVone_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionDeleteFailure %+v\n %+v\n", failureC, rsfC)
		rsdf, err := decodeRicSubscriptionDeleteFailure(&rsdfC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicSubscriptionDelete: &e2appdudescriptions.RicSubscriptionDelete{
				UnsuccessfulOutcome: rsdf,
				ProcedureCode: &e2ap_constants.IdRicsubscriptionDelete{
					Value: int32(v1beta2.ProcedureCodeIDRICsubscriptionDelete),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.UnsuccessfulOutcomeVone__value_PR_E2setupFailureVone:
		e2sfC := C.E2setupFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P13_t{
				list: C.struct___144{ // TODO: tie this down with a predictable name
					array: (**C.E2setupFailureIEsVone_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionDeleteFailure %+v\n %+v\n", failureC, rsfC)
		e2sf, err := decodeE2setupFailure(&e2sfC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2Setup: &e2appdudescriptions.E2Setup{
				UnsuccessfulOutcome: e2sf,
				ProcedureCode: &e2ap_constants.IdE2Setup{
					Value: int32(v1beta2.ProcedureCodeIDE2setup),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.UnsuccessfulOutcomeVone__value_PR_RICserviceUpdateFailureVone:
		rsufC := C.RICserviceUpdateFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P24_t{
				list: C.struct___145{ // TODO: tie this down with a predictable name
					array: (**C.RICserviceUpdateFailureVone_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionDeleteFailure %+v\n %+v\n", failureC, rsfC)
		rsuf, err := decodeRicServiceUpdateFailure(&rsufC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
				UnsuccessfulOutcome: rsuf,
				ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
					Value: int32(v1beta2.ProcedureCodeIDRICserviceUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.UnsuccessfulOutcomeVone__value_PR_E2nodeConfigurationUpdateFailureVone:
		e2ncufC := C.E2nodeConfigurationUpdateFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P19_t{
				list: C.struct___143{ // TODO: tie this down with a predictable name
					array: (**C.E2nodeConfigurationUpdateFailureVone_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionDeleteFailure %+v\n %+v\n", failureC, rsfC)
		e2ncuf, err := decodeE2nodeConfigurationUpdateFailure(&e2ncufC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
				UnsuccessfulOutcome: e2ncuf,
				ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
					Value: int32(v1beta2.ProcedureCodeIDE2nodeConfigurationUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.UnsuccessfulOutcomeVone__value_PR_E2connectionUpdateFailureVone:
		e2cufC := C.E2connectionUpdateFailureVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P16_t{
				list: C.struct___142{ // TODO: tie this down with a predictable name
					array: (**C.E2connectionUpdateFailureVone_IEs_t)(listArrayAddr),
					count: count,
					size:  size,
				},
			},
		}
		//fmt.Printf("RICsubscriptionDeleteFailure %+v\n %+v\n", failureC, rsfC)
		e2cuf, err := decodeE2connectionUpdateFailure(&e2cufC)
		if err != nil {
			return nil, err
		}
		uso.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
				UnsuccessfulOutcome: e2cuf,
				ProcedureCode: &e2ap_constants.IdE2ConnectionUpdate{
					Value: int32(v1beta2.ProcedureCodeIDE2connectionUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	default:
		return nil, fmt.Errorf("decodeSuccessfulOutcome() %v not yet implemented", failureC.value.present)
	}

	return uso, nil
}
