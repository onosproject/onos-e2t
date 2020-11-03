// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

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
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"unsafe"
)

func newE2setupRequestIe3GlobalE2NodeID(esIe *e2appducontents.E2SetupRequestIes_E2SetupRequestIes3) (*C.E2setupRequestIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDGlobalE2nodeID)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2setupRequestIEs__value_u

	globalNodeIDC, err := newGlobalE2nodeID(esIe.GetValue())
	if err != nil {
		return nil, err
	}
	fmt.Printf("Assigning to choice of E2setupRequestIE %v %v %v %v %v\n",
		globalNodeIDC, globalNodeIDC.present, &globalNodeIDC.choice,
		unsafe.Sizeof(globalNodeIDC.present), unsafe.Sizeof(globalNodeIDC.choice))
	binary.LittleEndian.PutUint32(choiceC[0:], uint32(globalNodeIDC.present))
	for i := 0; i < 8; i++ {
		choiceC[i+8] = globalNodeIDC.choice[i]
	}

	ie := C.E2setupRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupRequestIEs__value{
			present: C.E2setupRequestIEs__value_PR_GlobalE2node_ID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIe4GlobalRicID(esIe *e2appducontents.E2SetupResponseIes_E2SetupResponseIes4) (*C.E2setupResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDGlobalRicID)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2setupResponseIEs__value_u

	globalRicIDC, err := newGlobalRicID(esIe.Value)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Assigning to choice of E2setupReponseIE %v \n", globalRicIDC)

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(globalRicIDC.pLMN_Identity.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(globalRicIDC.pLMN_Identity.size))
	binary.LittleEndian.PutUint64(choiceC[40:], uint64(uintptr(unsafe.Pointer(globalRicIDC.ric_ID.buf))))
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(globalRicIDC.ric_ID.size))
	binary.LittleEndian.PutUint32(choiceC[56:], uint32(globalRicIDC.ric_ID.bits_unused))

	ie := C.E2setupResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupResponseIEs__value{
			present: C.E2setupResponseIEs__value_PR_GlobalRIC_ID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionRequestIe5RanFunctionID(rsrRfIe *e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes5) (*C.RICsubscriptionRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionIDC := newRanFunctionID(rsrRfIe.Value)

	fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICsubscriptionRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionRequest_IEs__value{
			present: C.RICsubscriptionRequest_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIe9RanFunctionsAccepted(esIe *e2appducontents.E2SetupResponseIes_E2SetupResponseIes9) (*C.E2setupResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionsAccepted)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionsIDList, err := newRanFunctionsIDList(esIe.Value)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Assigning to choice of E2setupReponseIE %v\n", ranFunctionsIDList)
	//binary.LittleEndian.PutUint32(choiceC[0:], uint32(ranFunctionsIDList.present))
	//for i := 0; i < 8; i++ {
	//	choiceC[i+8] = ranFunctionsIDList.choice[i]
	//}

	ie := C.E2setupResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupResponseIEs__value{
			present: C.E2setupResponseIEs__value_PR_RANfunctionsID_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupRequestIe10RanFunctionList(esIe *e2appducontents.E2SetupRequestIes_E2SetupRequestIes10) (*C.E2setupRequestIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionsAdded)
	if err != nil {
		return nil, err
	}

	listC := [48]byte{} // The size of the E2setupRequestIEs__value_u

	ranFunctionsListC, err := newRanFunctionsList(esIe.GetValue())
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Assigning to choice of E2setupRequestIE %v %v %v %v %v\n",
	//	ranFunctionsListC, ranFunctionsListC.present, &ranFunctionsListC.choice,
	//	unsafe.Sizeof(ranFunctionsListC.present), unsafe.Sizeof(ranFunctionsListC.choice))
	binary.LittleEndian.PutUint64(listC[8:], uint64(ranFunctionsListC.list.size))
	ie := C.E2setupRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupRequestIEs__value{
			present: C.E2setupRequestIEs__value_PR_RANfunctions_List,
			choice:  listC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIe13RanFunctionsRejected(esIe *e2appducontents.E2SetupResponseIes_E2SetupResponseIes13) (*C.E2setupResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionsRejected)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionsIDCauseList, err := newRanFunctionsIDcauseList(esIe.Value)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Assigning to choice of E2setupReponseIE %v\n", ranFunctionsIDCauseList)
	//binary.LittleEndian.PutUint32(choiceC[0:], uint32(ranFunctionsIDCauseList.present))
	//for i := 0; i < 8; i++ {
	//	choiceC[i+8] = ranFunctionsIDCauseList.choice[i]
	//}

	ie := C.E2setupResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupResponseIEs__value{
			present: C.E2setupResponseIEs__value_PR_RANfunctionsIDcause_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionRequestIe29RicRequestID(rsrRrIDIe *e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29) (*C.RICsubscriptionRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2setupResponseIEs__value_u

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICsubscriptionRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionRequest_IEs__value{
			present: C.RICsubscriptionRequest_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func decodeE2setupRequestIE(e2srIeC *C.E2setupRequestIEs_t) (*e2appducontents.E2SetupRequestIes, error) {
	//fmt.Printf("Handling E2SetupReqIE %+v\n", e2srIeC)
	ret := new(e2appducontents.E2SetupRequestIes)

	switch e2srIeC.value.present {
	case C.E2setupRequestIEs__value_PR_GlobalE2node_ID:
		gE2nID, err := decodeGlobalE2NodeID(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes3 = &e2appducontents.E2SetupRequestIes_E2SetupRequestIes3{
			Id:          int32(v1beta1.ProtocolIeIDGlobalE2nodeID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       gE2nID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2setupRequestIEs__value_PR_RANfunctions_List:
		fallthrough // TODO Implement it

	case C.E2setupRequestIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2setupRequestIEOld(). %v not yet implemneted", e2srIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2setupRequestIEOld(). unexpected choice %v", e2srIeC.value.present)
	}

	return ret, nil
}

//func decodeRicSubscriptionRequestIE(rsrIeC *C.RICsubscriptionRequest_IEs_t) (*e2appducontents.RicsubscriptionRequestIes, error) {
//	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsrIeC)
//	//ret := new(e2appducontents.RicsubscriptionRequestIes)
//
//	switch rsrIeC.value.present {
//	case C.RICsubscriptionRequest_IEs__value_PR_RICrequestID:
//		fallthrough // TODO Implement it
//
//	case C.RICsubscriptionRequest_IEs__value_PR_RANfunctionID:
//		fallthrough // TODO Implement it
//
//	case C.RICsubscriptionRequest_IEs__value_PR_RICsubscriptionDetails:
//		fallthrough // TODO Implement it
//
//	case C.E2setupRequestIEs__value_PR_NOTHING:
//		return nil, fmt.Errorf("decodeRicSubscriptionRequestIE(). %v not yet implemneted", rsrIeC.value.present)
//
//	default:
//		return nil, fmt.Errorf("decodeRicSubscriptionRequestIE(). unexpected choice %v", rsrIeC.value.present)
//	}
//
//}

func decodeRicSubscriptionResponseIE(rsrIeC *C.RICsubscriptionResponse_IEs_t) (*e2appducontents.RicsubscriptionResponseIes, error) {
	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsrIeC)
	ret := new(e2appducontents.RicsubscriptionResponseIes)

	switch rsrIeC.value.present {
	case C.RICsubscriptionResponse_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes5{
			Value:       decodeRanFunctionIDBytes(rsrIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionResponse_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsrIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionResponse_IEs__value_PR_RICaction_Admitted_List:
		raal, err := decodeRicActionAdmittedListBytes(rsrIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes17 = &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes17{
			Id:          int32(v1beta1.ProtocolIeIDRicactionsAdmitted),
			Value:       raal,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionResponse_IEs__value_PR_RICaction_NotAdmitted_List:
		fallthrough // TODO Implement it

	case C.RICsubscriptionResponse_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionResponseIE(). %v not yet implemneted", rsrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionResponseIE(). unexpected choice %v", rsrIeC.value.present)
	}

	return ret, nil
}
