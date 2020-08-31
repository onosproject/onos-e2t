// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Field.h"
//#include "RICaction-ToBeSetup-Item.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerEncodeE2setRequest - used only for testing
func XerEncodeE2setRequest(e2SetupReqIe *e2ctypes.E2SetupRequestIEsT) ([]byte, error) {
	e2SetupRequestIeC, err := newE2setupRequestIe(e2SetupReqIe)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2setupRequestIEs, unsafe.Pointer(e2SetupRequestIeC))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func newRicIndicationIe(riIe *e2ctypes.RICindication_IEsT) (*C.RICindication_IEs_t, error) {

	critC, err := criticalityToC(riIe.GetCriticality())
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(riIe.GetId())
	if err != nil {
		return nil, err
	}
	var vpr C.RICindication_IEs__value_PR

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u
	switch choice := riIe.GetChoice().(type) {
	case *e2ctypes.RICindication_IEsT_RICrequestID:
		vpr = C.RICindication_IEs__value_PR_RICrequestID
		reqID := C.RICrequestID_t{
			ricRequestorID: C.long(choice.RICrequestID.RicRequestorID),
			ricInstanceID:  C.long(choice.RICrequestID.RicInstanceID)}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(reqID.ricRequestorID))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(reqID.ricInstanceID))

	case *e2ctypes.RICindication_IEsT_RANfunctionID:
		vpr = C.RICindication_IEs__value_PR_RANfunctionID
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(choice.RANfunctionID))

	case *e2ctypes.RICindication_IEsT_RICactionID:
		vpr = C.RICindication_IEs__value_PR_RICactionID
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(choice.RICactionID))

	case *e2ctypes.RICindication_IEsT_RICindicationSN:
		vpr = C.RICindication_IEs__value_PR_RICindicationSN
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(choice.RICindicationSN))

	case *e2ctypes.RICindication_IEsT_RICindicationType:
		vpr = C.RICindication_IEs__value_PR_RICindicationType
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(choice.RICindicationType))

	case *e2ctypes.RICindication_IEsT_RICindicationHeader:
		vpr = C.RICindication_IEs__value_PR_RICindicationHeader
		// Encoded as OCTET_STRING
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&choice.RICindicationHeader))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(len(choice.RICindicationHeader)))

	case *e2ctypes.RICindication_IEsT_RICindicationMessage:
		vpr = C.RICindication_IEs__value_PR_RICindicationMessage
		// Encoded as OCTET_STRING
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(&choice.RICindicationMessage))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(len(choice.RICindicationMessage)))

	default:
		return nil, fmt.Errorf("not yet implemented IE type %v", choice)
	}

	ricIndIeC := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: vpr,
			choice:  choiceC,
		},
	}

	return &ricIndIeC, nil
}

func newRICsubscriptionRequestIE(rsrIe *e2ctypes.RICsubscriptionRequest_IEsT) (*C.RICsubscriptionRequest_IEs_t, error) {
	critC, err := criticalityToC(rsrIe.GetCriticality())
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(rsrIe.GetId())
	if err != nil {
		return nil, err
	}
	var vpr C.RICsubscriptionRequest_IEs__value_PR

	choiceC := [112]byte{} // The size of the RICsubscriptionRequest_IEs__value_u
	switch choice := rsrIe.GetChoice().(type) {
	case *e2ctypes.RICsubscriptionRequest_IEsT_RICrequestID:
		vpr = C.RICsubscriptionRequest_IEs__value_PR_RICrequestID
		reqID := C.RICrequestID_t{
			ricRequestorID: C.long(choice.RICrequestID.RicRequestorID),
			ricInstanceID:  C.long(choice.RICrequestID.RicInstanceID)}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(reqID.ricRequestorID))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(reqID.ricInstanceID))

	case *e2ctypes.RICsubscriptionRequest_IEsT_RANfunctionID:
		vpr = C.RICsubscriptionRequest_IEs__value_PR_RANfunctionID
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(choice.RANfunctionID))

	case *e2ctypes.RICsubscriptionRequest_IEsT_RICsubscriptionDetails:
		vpr = C.RICsubscriptionRequest_IEs__value_PR_RICsubscriptionDetails
		ricIeC, err := newRicSubscriptionDetails(choice.RICsubscriptionDetails)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("encode RICsubscriptionDetails [112]byte %v %v %v %v %v %v %v %v %v\n",
		//	ricIeC,
		//	ricIeC.ricAction_ToBeSetup_List.list.array,
		//	ricIeC.ricAction_ToBeSetup_List.list.count,
		//	ricIeC.ricAction_ToBeSetup_List.list.size,
		//	unsafe.Sizeof(*ricIeC),
		//	unsafe.Sizeof(ricIeC.ricEventTriggerDefinition.size),
		//	unsafe.Sizeof(ricIeC.ricAction_ToBeSetup_List.list.count),
		//	unsafe.Sizeof(ricIeC.ricAction_ToBeSetup_List.list.size),
		//	unsafe.Sizeof(ricIeC._asn_ctx))
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricIeC.ricEventTriggerDefinition.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricIeC.ricEventTriggerDefinition.size))
		binary.LittleEndian.PutUint64(choiceC[40:], uint64(uintptr(unsafe.Pointer(ricIeC.ricAction_ToBeSetup_List.list.array))))
		binary.LittleEndian.PutUint32(choiceC[48:], uint32(ricIeC.ricAction_ToBeSetup_List.list.count))
		binary.LittleEndian.PutUint32(choiceC[52:], uint32(ricIeC.ricAction_ToBeSetup_List.list.size))

	default:
		return nil, fmt.Errorf("not yet implemented IE type %v", choice)
	}

	ieC := C.RICsubscriptionRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionRequest_IEs__value{
			present: vpr,
			choice:  choiceC,
		},
	}

	return &ieC, nil
}

func newE2setupRequestIe(esIe *e2ctypes.E2SetupRequestIEsT) (*C.E2setupRequestIEs_t, error) {

	critC, err := criticalityToC(esIe.GetCriticality())
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(esIe.GetId())
	if err != nil {
		return nil, err
	}
	var vpr C.E2setupRequestIEs__value_PR

	choiceC := [48]byte{} // The size of the E2setupRequestIEs__value_u
	switch choice := esIe.Choice.(type) {
	case *e2ctypes.E2SetupRequestIEsT_GlobalE2Node_ID:
		vpr = C.E2setupRequestIEs__value_PR_GlobalE2node_ID
		globalNodeIDC, err := newGlobalE2nodeID(choice.GlobalE2Node_ID)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Assigning to choice of E2setupRequestIE %v %v %v %v %v\n",
		//	globalNodeIDC, globalNodeIDC.present, &globalNodeIDC.choice,
		//	unsafe.Sizeof(globalNodeIDC.present), unsafe.Sizeof(globalNodeIDC.choice))
		binary.LittleEndian.PutUint32(choiceC[0:], uint32(globalNodeIDC.present))
		for i := 0; i < 8; i++ {
			choiceC[i+8] = globalNodeIDC.choice[i]
		}
	default:
		return nil, fmt.Errorf("newE2setupRequestIe() %T not yet implemented", choice)
	}

	ie := C.E2setupRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupRequestIEs__value{
			present: vpr,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIE(e2srIe *e2ctypes.E2SetupResponseIEsT) (*C.E2setupResponseIEs_t, error) {

	critC, err := criticalityToC(e2srIe.GetCriticality())
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(e2srIe.GetId())
	if err != nil {
		return nil, err
	}
	var vpr C.E2setupResponseIEs__value_PR

	choiceC := [112]byte{} // The size of the E2setupRequestIEs__value_u
	switch choice := e2srIe.Choice.(type) {
	case *e2ctypes.E2SetupResponseIEsT_GlobalRIC_ID:
		vpr = C.E2setupResponseIEs__value_PR_GlobalRIC_ID
		grIDC, err := newGlobalRicID(choice.GlobalRIC_ID)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Assigning to choice of E2setupResponseIE %v %d %d %d %d %d\n",
		//	grIDC, unsafe.Sizeof(grIDC.pLMN_Identity.buf), unsafe.Sizeof(grIDC.pLMN_Identity.size),
		//	unsafe.Sizeof(grIDC.ric_ID.buf), unsafe.Sizeof(grIDC.ric_ID.size), unsafe.Sizeof(grIDC.ric_ID.bits_unused))
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(grIDC.pLMN_Identity.buf))))
		binary.LittleEndian.PutUint64(choiceC[8:], uint64(grIDC.pLMN_Identity.size))
		binary.LittleEndian.PutUint64(choiceC[40:], uint64(uintptr(unsafe.Pointer(grIDC.ric_ID.buf))))
		binary.LittleEndian.PutUint64(choiceC[48:], uint64(grIDC.ric_ID.size))
		binary.LittleEndian.PutUint32(choiceC[56:], uint32(grIDC.ric_ID.bits_unused))

	default:
		return nil, fmt.Errorf("newE2setupResponseIE() %T not yet implemented", choice)
	}

	ie := C.E2setupResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupResponseIEs__value{
			present: vpr,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func decodeErrorIndicationIE(errIndIeC *C.ErrorIndication_IEs_t) (*e2ctypes.ErrorIndication_IEsT, error) {

	ret := e2ctypes.ErrorIndication_IEsT{
		Id:          e2ctypes.ProtocolIE_IDT(errIndIeC.id),
		Criticality: e2ctypes.CriticalityT(errIndIeC.criticality),
	}

	switch errIndIeC.value.present {
	case C.ErrorIndication_IEs__value_PR_RICrequestID:
		ricReqIDC, err := decodeRicRequestID(errIndIeC.value.choice[:16])
		if err != nil {
			return nil, err
		}
		ret.Choice = &e2ctypes.ErrorIndication_IEsT_RICrequestID{RICrequestID: ricReqIDC}
	case C.ErrorIndication_IEs__value_PR_RANfunctionID:
		ranFunctionIDC, err := decodeRanFunctionID(errIndIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.Choice = &e2ctypes.ErrorIndication_IEsT_RANfunctionID{RANfunctionID: *ranFunctionIDC}
	case C.ErrorIndication_IEs__value_PR_Cause:
		fallthrough
	case C.ErrorIndication_IEs__value_PR_CriticalityDiagnostics:
		return nil, fmt.Errorf("decodeErrorIndicationIE(). %v not yet implemneted", errIndIeC.value.present)
	default:
		return nil, fmt.Errorf("decodeErrorIndicationIE(). unexpected choice %v", errIndIeC.value.present)
	}

	return &ret, nil
}

func decodeE2setupRequestIE(e2srIeC *C.E2setupRequestIEs_t) (*e2ctypes.E2SetupRequestIEsT, error) {
	//fmt.Printf("Handling E2SetupReqIE %+v\n", e2srIeC)
	ret := e2ctypes.E2SetupRequestIEsT{
		Id:          e2ctypes.ProtocolIE_IDT(e2srIeC.id),
		Criticality: e2ctypes.CriticalityT(e2srIeC.criticality),
	}

	switch e2srIeC.value.present {
	case C.E2setupRequestIEs__value_PR_GlobalE2node_ID:
		gE2nID, err := decodeGlobalE2NodeID(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.Choice = &e2ctypes.E2SetupRequestIEsT_GlobalE2Node_ID{GlobalE2Node_ID: gE2nID}
	case C.E2setupRequestIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2setupRequestIE(). %v not yet implemneted", e2srIeC.value.present)
	default:
		return nil, fmt.Errorf("decodeE2setupRequestIE(). unexpected choice %v", e2srIeC.value.present)
	}

	return &ret, nil
}

func decodeRICsubscriptionRequestIE(rsreqIeC *C.RICsubscriptionRequest_IEs_t) (*e2ctypes.RICsubscriptionRequest_IEsT, error) {
	fmt.Printf("Handling RICsubscriptionRequest_IE %+v\n", rsreqIeC)
	ret := e2ctypes.RICsubscriptionRequest_IEsT{
		Id:          e2ctypes.ProtocolIE_IDT(rsreqIeC.id),
		Criticality: e2ctypes.CriticalityT(rsreqIeC.criticality),
	}

	switch rsreqIeC.value.present {
	case C.RICsubscriptionRequest_IEs__value_PR_RICrequestID:
		//gE2nID, err := decodeGlobalE2NodeID(rsreqIeC.value.choice)
		//if err != nil {
		//	return nil, err
		//}
		//ret.Choice = &e2ctypes.RICsubscriptionRequest_IEsT_RICrequestID{RICrequestID: gE2nID}
	default:
		return nil, fmt.Errorf("decodeRICsubscriptionRequestIE(). %v not yet implemneted", rsreqIeC.value.present)
	}

	return &ret, nil
}

func decodeRICsubscriptionResponseIE(rsrespIeC *C.RICsubscriptionResponse_IEs_t) (*e2ctypes.RICsubscriptionResponse_IEsT, error) {
	//fmt.Printf("Handling RICsubscriptionResponse_IE %+v\n", rsrespIeC)
	ret := e2ctypes.RICsubscriptionResponse_IEsT{
		Id:          e2ctypes.ProtocolIE_IDT(rsrespIeC.id),
		Criticality: e2ctypes.CriticalityT(rsrespIeC.criticality),
	}

	switch rsrespIeC.value.present {
	case C.RICsubscriptionResponse_IEs__value_PR_NOTHING:
		//Nothing to do
	case C.RICsubscriptionResponse_IEs__value_PR_RICrequestID:
		rrID, err := decodeRicRequestID(rsrespIeC.value.choice[:16])
		if err != nil {
			return nil, err
		}
		ret.Choice = &e2ctypes.RICsubscriptionResponse_IEsT_RICrequestID{
			RICrequestID: rrID,
		}
	case C.RICsubscriptionResponse_IEs__value_PR_RICaction_Admitted_List:
		raal, err := decodeRicActionAdmittedList(rsrespIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.Choice = &e2ctypes.RICsubscriptionResponse_IEsT_RICaction_Admitted_List{
			RICaction_Admitted_List: raal,
		}
	default:
		return nil, fmt.Errorf("decodeRICsubscriptionResponseIE(). %v not yet implemneted", rsrespIeC.value.present)
	}

	return &ret, nil
}
