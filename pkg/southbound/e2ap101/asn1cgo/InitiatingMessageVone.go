// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "InitiatingMessageVone.h"
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

func newInitiatingMessage(im *e2appdudescriptions.InitiatingMessage) (*C.struct_InitiatingMessageVone, error) {

	var presentC C.InitiatingMessageVone__value_PR
	var pcC C.ProcedureCodeVone_t
	var critC C.CriticalityVone_t
	choiceC := [72]byte{} // The size of the InitiatingMessage__value_u union

	if pc := im.GetProcedureCode().GetE2Setup(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_E2setupRequestVone
		pcC = C.ProcedureCodeVone_id_E2setupVone
		critC = C.long(C.CriticalityVone_reject)
		e2sC, err := newE2SetupRequest(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2sC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2sC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2sC.protocolIEs.list.size))
	} else if pc := im.GetProcedureCode().GetRicSubscription(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_RICsubscriptionRequestVone
		pcC = C.ProcedureCodeVone_id_RICsubscriptionVone
		critC = C.long(C.CriticalityVone_reject)
		rsC, err := newRICsubscriptionRequest(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetRicSubscriptionDelete(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_RICsubscriptionDeleteRequestVone
		pcC = C.ProcedureCodeVone_id_RICsubscriptionDeleteVone
		critC = C.long(C.CriticalityVone_reject)
		rsdC, err := newRICsubscriptionDeleteRequest(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsdC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsdC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsdC.protocolIEs.list.size))
	} else if pc := im.GetProcedureCode().GetRicIndication(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_RICindicationVone
		pcC = C.ProcedureCodeVone_id_RICindicationVone
		critC = C.long(C.CriticalityVone_reject)
		riC, err := newRicIndication(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(riC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(riC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(riC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetRicControl(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_RICcontrolRequestVone
		pcC = C.ProcedureCodeVone_id_RICcontrolVone
		critC = C.long(C.CriticalityVone_reject)
		rcC, err := newRicControlRequest(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rcC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rcC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rcC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetErrorIndication(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_ErrorIndicationVone
		pcC = C.ProcedureCodeVone_id_ErrorIndicationVone
		critC = C.long(C.CriticalityVone_ignore)
		eiC, err := newErrorIndication(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(eiC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(eiC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(eiC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetRicServiceQuery(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_RICserviceQueryVone
		pcC = C.ProcedureCodeVone_id_RICserviceQueryVone
		critC = C.long(C.CriticalityVone_ignore)
		rsqC, err := newRicServiceQuery(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsqC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsqC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsqC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetReset_(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_ResetRequestVone
		pcC = C.ProcedureCodeVone_id_ResetVone
		critC = C.long(C.CriticalityVone_reject)
		rrC, err := newResetRequest(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rrC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rrC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rrC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetRicServiceUpdate(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_RICserviceUpdateVone
		pcC = C.ProcedureCodeVone_id_RICserviceUpdateVone
		critC = C.long(C.CriticalityVone_reject)
		rsuC, err := newRicServiceUpdate(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsuC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(rsuC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(rsuC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetE2NodeConfigurationUpdate(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_E2nodeConfigurationUpdateVone
		pcC = C.ProcedureCodeVone_id_E2nodeConfigurationUpdateVone
		critC = C.long(C.CriticalityVone_reject)
		e2ncuC, err := newE2nodeConfigurationUpdate(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuC.protocolIEs.list.size))

	} else if pc := im.GetProcedureCode().GetE2ConnectionUpdate(); pc != nil &&
		pc.GetInitiatingMessage() != nil {

		presentC = C.InitiatingMessageVone__value_PR_E2connectionUpdateVone
		pcC = C.ProcedureCodeVone_id_E2connectionUpdateVone
		critC = C.long(C.CriticalityVone_reject)
		e2cuC, err := newE2connectionUpdate(pc.GetInitiatingMessage())
		if err != nil {
			return nil, err
		}
		//	//fmt.Printf("Protocol IEs %v %v %v\n", rsrC.protocolIEs.list.array, rsrC.protocolIEs.list.count, rsrC.protocolIEs.list.size)
		//	// Now copy the rsrC over in to the choice byte by byte - the union is [72]byte
		//	// It's A_SET_OF, so has <address(8), count(4), size(4)>
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2cuC.protocolIEs.list.array))))
		binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2cuC.protocolIEs.list.count))
		binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2cuC.protocolIEs.list.size))

	} else {
		return nil, fmt.Errorf("newInitiatingMessageValue type not yet implemented")
	}

	imC := C.InitiatingMessageVone_t{
		procedureCode: pcC,
		criticality:   critC,
		value: C.struct_InitiatingMessageVone__value{
			present: presentC,
			choice:  choiceC,
		},
	}

	return &imC, nil
}

func decodeInitiatingMessage(initMsgC *C.InitiatingMessageVone_t) (*e2appdudescriptions.InitiatingMessage, error) {

	initiatingMessage := new(e2appdudescriptions.InitiatingMessage)

	listArrayAddr := initMsgC.value.choice[0:8]

	switch initMsgC.value.present {
	case C.InitiatingMessageVone__value_PR_E2setupRequestVone:
		e2srC := *(**C.E2setupRequestIEsVone_t)(unsafe.Pointer(&listArrayAddr[0]))
		esC := C.E2setupRequestVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P11_t{
				list: C.struct___70{ // TODO: tie this down with a predictable name
					array: (**C.E2setupRequestIEsVone_t)(unsafe.Pointer(e2srC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		//fmt.Printf("E2SetupRequestC %+v\n %+v\n", initMsgC, riC)
		e2sr, err := decodeE2setupRequest(&esC)
		if err != nil {
			return nil, err
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2Setup: &e2appdudescriptions.E2Setup{
				InitiatingMessage: e2sr,
				ProcedureCode: &e2ap_constants.IdE2Setup{
					Value: int32(v1beta2.ProcedureCodeIDE2setup),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}
	case C.InitiatingMessageVone__value_PR_RICsubscriptionRequestVone:
		ricsrC := *(**C.RICsubscriptionRequestVone_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		srC := C.RICsubscriptionRequestVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P0_t{
				list: C.struct___125{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionRequestVone_IEs_t)(unsafe.Pointer(ricsrC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		//fmt.Printf("RICsubscriptionRequest_t %+v\n %+v\n", initMsgC, sdrC)

		sr, err := decodeRicSubscriptionRequest(&srC)
		if err != nil {
			return nil, err
		}

		// TODO: Get the value
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicSubscription: &e2appdudescriptions.RicSubscription{
				InitiatingMessage: sr,
				ProcedureCode: &e2ap_constants.IdRicsubscription{
					Value: int32(v1beta2.ProcedureCodeIDRICsubscription),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.InitiatingMessageVone__value_PR_RICsubscriptionDeleteRequestVone:
		ricsdrC := *(**C.RICsubscriptionDeleteRequestVone_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		sdrC := C.RICsubscriptionDeleteRequestVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P3_t{
				list: C.struct___119{ // TODO: tie this down with a predictable name
					array: (**C.RICsubscriptionDeleteRequestVone_IEs_t)(unsafe.Pointer(ricsdrC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		//fmt.Printf("RICsubscriptionRequest_t %+v\n %+v\n", initMsgC, sdrC)

		sdr, err := decodeRicSubscriptionDeleteRequest(&sdrC)
		if err != nil {
			return nil, err
		}

		// TODO: Get the value
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicSubscriptionDelete: &e2appdudescriptions.RicSubscriptionDelete{
				InitiatingMessage: sdr,
				ProcedureCode: &e2ap_constants.IdRicsubscriptionDelete{
					Value: int32(v1beta2.ProcedureCodeIDRICsubscriptionDelete),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{},
			},
		}

	case C.InitiatingMessageVone__value_PR_RICindicationVone:
		riIesC := *(**C.RICindicationVone_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		riC := C.RICindicationVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P6_t{
				list: C.struct___107{ // TODO: tie this down with a predictable name
					array: (**C.RICindicationVone_IEs_t)(unsafe.Pointer(riIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("RICindication_t %+v\n %+v\n", initMsgC, riC)

		ri, err := decodeRicIndication(&riC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicIndication() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicIndication: &e2appdudescriptions.RicIndication{
				InitiatingMessage: ri,
				ProcedureCode: &e2ap_constants.IdRicindication{
					Value: int32(v1beta2.ProcedureCodeIDRICindication),
				},
				Criticality: &e2ap_commondatatypes.CriticalityIgnore{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_RICcontrolRequestVone:
		rcrIesC := *(**C.RICcontrolRequestVone_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		rcrC := C.RICcontrolRequestVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P7_t{
				list: C.struct___106{ // TODO: tie this down with a predictable name
					array: (**C.RICcontrolRequestVone_IEs_t)(unsafe.Pointer(rcrIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("RICcontrolRequest_t %+v\n %+v\n", initMsgC, rcrC)

		rcr, err := decodeRicControlRequest(&rcrC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicControlRequest() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicControl: &e2appdudescriptions.RicControl{
				InitiatingMessage: rcr,
				ProcedureCode: &e2ap_constants.IdRiccontrol{
					Value: int32(v1beta2.ProcedureCodeIDRICcontrol),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_ErrorIndicationVone:
		eiIesC := *(**C.ErrorIndicationVone_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		eiC := C.ErrorIndicationVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P10_t{
				list: C.struct___69{ // TODO: tie this down with a predictable name
					array: (**C.ErrorIndicationVone_IEs_t)(unsafe.Pointer(eiIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("ErrorIndication_t %+v\n %+v\n", initMsgC, eiC)

		ei, err := decodeErrorIndication(&eiC)
		if err != nil {
			return nil, fmt.Errorf("decodeErrorIndication() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			ErrorIndication: &e2appdudescriptions.ErrorIndicationEp{
				InitiatingMessage: ei,
				ProcedureCode: &e2ap_constants.IdErrorIndication{
					Value: int32(v1beta2.ProcedureCodeIDErrorIndication),
				},
				Criticality: &e2ap_commondatatypes.CriticalityIgnore{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_RICserviceQueryVone:
		rsqIesC := *(**C.RICserviceQueryVone_IEs_t)(unsafe.Pointer(&listArrayAddr[0]))
		rsqC := C.RICserviceQueryVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P25_t{
				list: C.struct___108{ // TODO: tie this down with a predictable name
					array: (**C.RICserviceQueryVone_IEs_t)(unsafe.Pointer(rsqIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("RICserviceQuery_t %+v\n %+v\n", initMsgC, rsqC)

		rsq, err := decodeRicServiceQuery(&rsqC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicServiceQuery() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicServiceQuery: &e2appdudescriptions.RicServiceQuery{
				InitiatingMessage: rsq,
				ProcedureCode: &e2ap_constants.IdRicserviceQuery{
					Value: int32(v1beta2.ProcedureCodeIDRICserviceQuery),
				},
				Criticality: &e2ap_commondatatypes.CriticalityIgnore{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_ResetRequestVone:
		rrIesC := *(**C.ResetRequestVone_t)(unsafe.Pointer(&listArrayAddr[0]))
		rrC := C.ResetRequestVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P20_t{
				list: C.struct___129{ // TODO: tie this down with a predictable name
					array: (**C.ResetRequestIEsVone_t)(unsafe.Pointer(rrIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("ResetRequest_t %+v\n %+v\n", initMsgC, rrC)

		rr, err := decodeResetRequest(&rrC)
		if err != nil {
			return nil, fmt.Errorf("decodeResetRequest() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			Reset_: &e2appdudescriptions.Reset{
				InitiatingMessage: rr,
				ProcedureCode: &e2ap_constants.IdReset{
					Value: int32(v1beta2.ProcedureCodeIDReset),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_RICserviceUpdateVone:
		riIesC := *(**C.RICserviceUpdateVone_t)(unsafe.Pointer(&listArrayAddr[0]))
		rsuC := C.RICserviceUpdateVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P22_t{
				list: C.struct___115{ // TODO: tie this down with a predictable name
					array: (**C.RICserviceUpdateVone_IEs_t)(unsafe.Pointer(riIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("RICserviceUpdate_t %+v\n %+v\n", initMsgC, rsuC)

		rsu, err := decodeRicServiceUpdate(&rsuC)
		if err != nil {
			return nil, fmt.Errorf("decodeRicServiceUpdate() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
				InitiatingMessage: rsu,
				ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
					Value: int32(v1beta2.ProcedureCodeIDRICserviceUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_E2nodeConfigurationUpdateVone:
		e2ncuIesC := *(**C.E2nodeConfigurationUpdateVone_t)(unsafe.Pointer(&listArrayAddr[0]))
		e2ncuC := C.E2nodeConfigurationUpdateVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P17_t{
				list: C.struct___76{ // TODO: tie this down with a predictable name
					array: (**C.E2nodeConfigurationUpdateVone_IEs_t)(unsafe.Pointer(e2ncuIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("E2nodeConfigurationUpdate_t %+v\n %+v\n", initMsgC, e2ncuC)

		e2ncu, err := decodeE2nodeConfigurationUpdate(&e2ncuC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeConfigurationUpdate() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2NodeConfigurationUpdate: &e2appdudescriptions.E2NodeConfigurationUpdateEp{
				InitiatingMessage: e2ncu,
				ProcedureCode: &e2ap_constants.IdE2NodeConfigurationUpdate{
					Value: int32(v1beta2.ProcedureCodeIDE2nodeConfigurationUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				},
			},
		}

	case C.InitiatingMessageVone__value_PR_E2connectionUpdateVone:
		e2cuIesC := *(**C.E2connectionUpdateVone_t)(unsafe.Pointer(&listArrayAddr[0]))
		e2cuC := C.E2connectionUpdateVone_t{
			protocolIEs: C.ProtocolIE_Container_1710P14_t{
				list: C.struct___73{ // TODO: tie this down with a predictable name
					array: (**C.E2connectionUpdateVone_IEs_t)(unsafe.Pointer(e2cuIesC)),
					count: C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[8:12])),
					size:  C.int(binary.LittleEndian.Uint32(initMsgC.value.choice[12:16])),
				},
			},
		}
		fmt.Printf("E2connectionUpdate_t %+v\n %+v\n", initMsgC, e2cuC)

		e2cu, err := decodeE2connectionUpdate(&e2cuC)
		if err != nil {
			return nil, fmt.Errorf("decodeE2connectionUpdate() %s", err.Error())
		}
		initiatingMessage.ProcedureCode = &e2appdudescriptions.E2ApElementaryProcedures{
			E2ConnectionUpdate: &e2appdudescriptions.E2ConnectionUpdateEp{
				InitiatingMessage: e2cu,
				ProcedureCode: &e2ap_constants.IdE2ConnectionUpdate{
					Value: int32(v1beta2.ProcedureCodeIDE2connectionUpdate),
				},
				Criticality: &e2ap_commondatatypes.CriticalityReject{
					Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				},
			},
		}

	default:
		return nil, fmt.Errorf("decodeInitiatingMessage() %v not yet implemented", initMsgC.value.present)
	}

	return initiatingMessage, nil
}
