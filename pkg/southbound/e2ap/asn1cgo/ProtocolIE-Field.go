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
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2ap-commondatatypes"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"unsafe"
)

func newRicSubscriptionDeleteFailureIe1Cause(rsdfCauseIe *e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes1) (*C.RICsubscriptionDeleteFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsdfCauseIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	rsdfCauseC, err := newCause(rsdfCauseIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rsdfCauseC.present))
	copy(choiceC[8:16], rsdfCauseC.choice[:8])

	ie := C.RICsubscriptionDeleteFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteFailure_IEs__value{
			present: C.RICsubscriptionDeleteFailure_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteFailureIe2CriticalityDiagnostics(rsdfCritDiagsIe *e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes2) (*C.RICsubscriptionDeleteFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsdfCritDiagsIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	rsdfCritDiagsC, err := newCriticalityDiagnostics(rsdfCritDiagsIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsdfCritDiagsC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rsdfCritDiagsC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(rsdfCritDiagsC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(rsdfCritDiagsC.ricRequestorID))))
	//binary.LittleEndian.PutUint64(choiceC[40:], uint64(uintptr(unsafe.Pointer(rsdfCritDiagsC.iEsCriticalityDiagnostics))))

	ie := C.RICsubscriptionDeleteFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteFailure_IEs__value{
			present: C.RICsubscriptionDeleteFailure_IEs__value_PR_CriticalityDiagnostics,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionFailureIe2CriticalityDiagnostics(rsfCritDiagsIe *e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes2) (*C.RICsubscriptionFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsfCritDiagsIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionFailure_IEs__value

	rsfCritDiagsC, err := newCriticalityDiagnostics(rsfCritDiagsIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsfCritDiagsC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rsfCritDiagsC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(rsfCritDiagsC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(rsfCritDiagsC.ricRequestorID))))
	//binary.LittleEndian.PutUint64(choiceC[40:], uint64(uintptr(unsafe.Pointer(rsfCritDiagsC.iEsCriticalityDiagnostics))))

	ie := C.RICsubscriptionFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionFailure_IEs__value{
			present: C.RICsubscriptionFailure_IEs__value_PR_CriticalityDiagnostics,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

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
	//fmt.Printf("Assigning to choice of E2setupRequestIE %v %v %v %v %v\n",
	//	globalNodeIDC, globalNodeIDC.present, &globalNodeIDC.choice,
	//	unsafe.Sizeof(globalNodeIDC.present), unsafe.Sizeof(globalNodeIDC.choice))
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
	//fmt.Printf("Assigning to choice of E2setupReponseIE %v \n", globalRicIDC)

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

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
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

func newRicSubscriptionResponseIe5RanFunctionID(rsrRfIe *e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes5) (*C.RICsubscriptionResponse_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionIDC := newRanFunctionID(rsrRfIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICsubscriptionResponse_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionResponse_IEs__value{
			present: C.RICsubscriptionResponse_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe5RanFunctionID(rsrRfIe *e2appducontents.RicindicationIes_RicindicationIes5) (*C.RICsubscriptionRequest_IEs_t, error) {
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

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
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

func newRicSubscriptionDeleteRequestIe5RanFunctionID(rsdrRfIe *e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes5) (*C.RICsubscriptionDeleteRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsdrRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionIDC := newRanFunctionID(rsdrRfIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICsubscriptionDeleteRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteRequest_IEs__value{
			present: C.RICsubscriptionDeleteRequest_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteResponseIe5RanFunctionID(rsdrRfIe *e2appducontents.RicsubscriptionDeleteResponseIes_RicsubscriptionDeleteResponseIes5) (*C.RICsubscriptionDeleteResponse_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsdrRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionIDC := newRanFunctionID(rsdrRfIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICsubscriptionDeleteResponse_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteResponse_IEs__value{
			present: C.RICsubscriptionDeleteResponse_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteFailureIe5RanFunctionID(rsdfRfIe *e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes5) (*C.RICsubscriptionDeleteFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsdfRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value_u

	ranFunctionIDC := newRanFunctionID(rsdfRfIe.Value)

	//fmt.Printf("Assigning to choice of RICsubscriptionDeleteFailureIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICsubscriptionDeleteFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteFailure_IEs__value{
			present: C.RICsubscriptionDeleteFailure_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionFailureIe5RanFunctionID(rsfRfIe *e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes5) (*C.RICsubscriptionFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsfRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value_u

	ranFunctionIDC := newRanFunctionID(rsfRfIe.Value)

	//fmt.Printf("Assigning to choice of RICsubscriptionDeleteFailureIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICsubscriptionFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionFailure_IEs__value{
			present: C.RICsubscriptionFailure_IEs__value_PR_RANfunctionID,
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

	ranFunctionsIDListC, err := newRanFunctionsIDList(esIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsIDList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsIDListC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ranFunctionsIDListC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(ranFunctionsIDListC.list.size))

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
		return nil, fmt.Errorf("newRanFunctionsList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ranFunctionsListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ranFunctionsListC.list.size))

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
		return nil, fmt.Errorf("newRanFunctionsIDcauseList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsIDCauseList.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ranFunctionsIDCauseList.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(ranFunctionsIDCauseList.list.size))

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

func newRicIndicationIe15RicActionID(riIe *e2appducontents.RicindicationIes_RicindicationIes15) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(riIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicactionID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricActionID := newRicActionID(riIe.Value)

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*ricActionID))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionResponseIe17RactionAdmittedList(rsrRrIe *e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes17) (*C.RICsubscriptionResponse_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicactionsAdmitted)
	if err != nil {
		return nil, err
	}

	listC := [48]byte{} // The size of the E2setupResponseIEs__value_u

	ricActionAdmittedListC, err := newRicActionAdmittedList(rsrRrIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicActionAdmittedList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ricActionAdmittedListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ricActionAdmittedListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ricActionAdmittedListC.list.size))

	ie := C.RICsubscriptionResponse_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionResponse_IEs__value{
			present: C.RICsubscriptionResponse_IEs__value_PR_RICaction_Admitted_List,
			choice:  listC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionFailureIe18RicActionNotAdmittedList(rsfRanaIe *e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes18) (*C.RICsubscriptionFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsfRanaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicactionsNotAdmitted)
	if err != nil {
		return nil, err
	}

	listC := [64]byte{} // The size of the E2setupResponseIEs__value_u

	ricActionNotAdmittedListC, err := newRicActionNotAdmittedList(rsfRanaIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicActionAdmittedList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ricActionNotAdmittedListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ricActionNotAdmittedListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ricActionNotAdmittedListC.list.size))

	ie := C.RICsubscriptionFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionFailure_IEs__value{
			present: C.RICsubscriptionFailure_IEs__value_PR_RICaction_NotAdmitted_List,
			choice:  listC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe20RiccallProcessID(riIe20 *e2appducontents.RicindicationIes_RicindicationIes20) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(riIe20.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRiccallProcessID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricCallProcessID := newRicCallProcessID(riIe20.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricCallProcessID.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricCallProcessID.size))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICcallProcessID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe25RicIndicationHeader(rihIe *e2appducontents.RicindicationIes_RicindicationIes25) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rihIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicindicationHeader)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricIndicationHeader := newRicIndicationHeader(rihIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricIndicationHeader.buf))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ricIndicationHeader.size))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICindicationHeader,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe26RicIndicationMessage(rimIe *e2appducontents.RicindicationIes_RicindicationIes26) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rimIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicindicationMessage)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricIndicationMessage := newRicIndicationMessage(rimIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricIndicationMessage.buf))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ricIndicationMessage.size))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICindicationMessage,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe27RicIndicationSn(risnIe *e2appducontents.RicindicationIes_RicindicationIes27) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(risnIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicindicationSn)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricIndicationSequenceNumber := newRicIndicationSn(risnIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*ricIndicationSequenceNumber))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICindicationSN,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe28RicIndicationType(ritIe *e2appducontents.RicindicationIes_RicindicationIes28) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(ritIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicindicationType)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricIndicationTypeC, err := newRicIndicationType(&ritIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicIndicationType() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*ricIndicationTypeC))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICindicationType,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicIndicationIe29RicRequestID(rsrRrIDIe *e2appducontents.RicindicationIes_RicindicationIes29) (*C.RICindication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICindication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICindication_IEs__value{
			present: C.RICindication_IEs__value_PR_RICrequestID,
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

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ricRequestIDC)
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

func newRicSubscriptionResponseIe29RicRequestID(rsrRrIDIe *e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes29) (*C.RICsubscriptionResponse_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{}

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICsubscriptionResponse_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionResponse_IEs__value{
			present: C.RICsubscriptionResponse_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteRequestIe29RicRequestID(rsrdRrIDIe *e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes29) (*C.RICsubscriptionDeleteRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrdRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricRequestIDC := newRicRequestID(rsrdRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICsubscriptionDeleteRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteRequest_IEs__value{
			present: C.RICsubscriptionDeleteRequest_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteResponseIe29RicRequestID(rsrRrIDIe *e2appducontents.RicsubscriptionDeleteResponseIes_RicsubscriptionDeleteResponseIes29) (*C.RICsubscriptionDeleteResponse_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{}

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICsubscriptionDeleteResponse_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteResponse_IEs__value{
			present: C.RICsubscriptionDeleteResponse_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteFailureIe29RicRequestID(rsrRrIDIe *e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes29) (*C.RICsubscriptionDeleteFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{}

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICsubscriptionDeleteFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteFailure_IEs__value{
			present: C.RICsubscriptionDeleteFailure_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionFailureIe29RicRequestID(rsrRrIDIe *e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes29) (*C.RICsubscriptionFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{}

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICsubscriptionFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionFailure_IEs__value{
			present: C.RICsubscriptionFailure_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionRequestIe30RicSubscriptionDetails(rsrDetIe *e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes30) (*C.RICsubscriptionRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrDetIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicsubscriptionDetails)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2setupResponseIEs__value_u

	rsrDetC, err := newRicSubscriptionDetails(rsrDetIe.GetValue())
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", rsrDetC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rsrDetC.ricEventTriggerDefinition.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(rsrDetC.ricEventTriggerDefinition.size))
	binary.LittleEndian.PutUint64(choiceC[40:], uint64(uintptr(unsafe.Pointer(rsrDetC.ricAction_ToBeSetup_List.list.array))))
	binary.LittleEndian.PutUint32(choiceC[48:], uint32(rsrDetC.ricAction_ToBeSetup_List.list.count))
	binary.LittleEndian.PutUint32(choiceC[52:], uint32(rsrDetC.ricAction_ToBeSetup_List.list.size))

	ie := C.RICsubscriptionRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionRequest_IEs__value{
			present: C.RICsubscriptionRequest_IEs__value_PR_RICsubscriptionDetails,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRANfunctionItemIEs(rfItemIes *e2appducontents.RanfunctionItemIes) (*C.RANfunction_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rfItemIes.GetE2ApProtocolIes10().GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the RANfunction_ItemIEs__value_u
	rfItemC := newRanFunctionItem(rfItemIes.GetE2ApProtocolIes10().GetValue())
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rfItemC.ranFunctionID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rfItemC.ranFunctionDefinition.buf))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(rfItemC.ranFunctionDefinition.size))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(rfItemC.ranFunctionRevision))

	rfItemIesC := C.RANfunction_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RANfunction_ItemIEs__value{
			present: C.RANfunction_ItemIEs__value_PR_RANfunction_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newRANfunctionIDItemIEs(rfIDItemIes *e2appducontents.RanfunctionIdItemIes) (*C.RANfunctionID_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rfIDItemIes.GetRanFunctionIdItemIes6().GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionIDItem)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RANfunction_ItemIEs__value_u
	rfIDItemC := newRanFunctionIDItem(rfIDItemIes.GetRanFunctionIdItemIes6().GetValue())
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rfIDItemC.ranFunctionID))
	binary.LittleEndian.PutUint64(choiceC[8:16], uint64(rfIDItemC.ranFunctionRevision))

	rfItemIesC := C.RANfunctionID_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RANfunctionID_ItemIEs__value{
			present: C.RANfunctionID_ItemIEs__value_PR_RANfunctionID_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newRANfunctionIDCauseItemIEs(rfIDItemIes *e2appducontents.RanfunctionIdcauseItemIes) (*C.RANfunctionIDcause_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rfIDItemIes.GetRanFunctionIdcauseItemIes7().GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRanfunctionIeCauseItem)
	if err != nil {
		return nil, err
	}

	choiceC := [72]byte{} // The size of the RANfunction_ItemIEs__value_u
	rfIDItemC, err := newRanFunctionIDCauseItem(rfIDItemIes.GetRanFunctionIdcauseItemIes7().GetValue())
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionIDCauseItem() error %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rfIDItemC.ranFunctionID))
	binary.LittleEndian.PutUint64(choiceC[8:16], uint64(rfIDItemC.cause.present))
	copy(choiceC[16:24], rfIDItemC.cause.choice[:])

	rfItemIesC := C.RANfunctionIDcause_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RANfunctionIDcause_ItemIEs__value{
			present: C.RANfunctionIDcause_ItemIEs__value_PR_RANfunctionIDcause_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newRicActionAdmittedItemIEs(raaItemIes *e2appducontents.RicactionAdmittedItemIes) (*C.RICaction_Admitted_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(raaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicactionAdmittedItem)
	if err != nil {
		return nil, err
	}

	choiceC := [32]byte{} // The size of the RANfunction_ItemIEs__value_u
	rfItemC := newRicActionAdmittedItem(raaItemIes.GetValue())
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rfItemC.ricActionID))

	rfItemIesC := C.RICaction_Admitted_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICaction_Admitted_ItemIEs__value{
			present: C.RICaction_Admitted_ItemIEs__value_PR_RICaction_Admitted_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newRicActionNotAdmittedItemIEs(ranaItemIes *e2appducontents.RicactionNotAdmittedItemIes) (*C.RICaction_NotAdmitted_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(ranaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicactionNotAdmittedItem)
	if err != nil {
		return nil, err
	}

	choiceC := [72]byte{} // The size of the RANfunction_ItemIEs__value_u
	rfItemC, err := newRicActionNotAdmittedItem(ranaItemIes.GetValue())
	if err != nil {
		return nil, fmt.Errorf("newRicActionNotAdmittedItem() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rfItemC.ricActionID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(rfItemC.cause.present))
	copy(choiceC[16:24], rfItemC.cause.choice[:])

	rfItemIesC := C.RICaction_NotAdmitted_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICaction_NotAdmitted_ItemIEs__value{
			present: C.RICaction_NotAdmitted_ItemIEs__value_PR_RICaction_NotAdmitted_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newRicActionToBeSetupItemIEs(ratbsItemIes *e2appducontents.RicactionToBeSetupItemIes) (*C.RICaction_ToBeSetup_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(ratbsItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v1beta1.ProtocolIeIDRicactionToBeSetupItem)
	if err != nil {
		return nil, err
	}

	choiceC := [56]byte{} // The size of the RANfunction_ItemIEs__value_u
	ratbsItemC, err := newRicActionToBeSetupItem(ratbsItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ratbsItemC.ricActionID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ratbsItemC.ricActionType))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(ratbsItemC.ricActionDefinition))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(ratbsItemC.ricSubsequentAction))))

	rfItemIesC := C.RICaction_ToBeSetup_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICaction_ToBeSetup_ItemIEs__value{
			present: C.RICaction_ToBeSetup_ItemIEs__value_PR_RICaction_ToBeSetup_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
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
		rfl, err := decodeRanFunctionsListBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes10 = &e2appducontents.E2SetupRequestIes_E2SetupRequestIes10{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionsAdded),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	case C.E2setupRequestIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2setupRequestIE(). %v not yet implemneted", e2srIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2setupRequestIE(). unexpected choice %v", e2srIeC.value.present)
	}

	return ret, nil
}

func decodeE2setupResponseIE(e2srIeC *C.E2setupResponseIEs_t) (*e2appducontents.E2SetupResponseIes, error) {
	//fmt.Printf("Handling E2SetupReqIE %+v\n", e2srIeC)
	ret := new(e2appducontents.E2SetupResponseIes)

	switch e2srIeC.value.present {
	case C.E2setupResponseIEs__value_PR_GlobalRIC_ID:
		gE2nID, err := decodeGlobalRicIDBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes4 = &e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
			Id:          int32(v1beta1.ProtocolIeIDGlobalRicID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       gE2nID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.E2setupResponseIEs__value_PR_RANfunctionsID_List:
		rfAccepted, err := decodeRanFunctionsIDListBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes9 = &e2appducontents.E2SetupResponseIes_E2SetupResponseIes9{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionsAccepted),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfAccepted,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	case C.E2setupResponseIEs__value_PR_RANfunctionsIDcause_List:
		rfRejected, err := decodeRanFunctionsIDCauseListBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes13 = &e2appducontents.E2SetupResponseIes_E2SetupResponseIes13{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionsRejected),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfRejected,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	case C.E2setupResponseIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2setupResponseIE(). %v not yet implemneted", e2srIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2setupResponseIE(). unexpected choice %v", e2srIeC.value.present)
	}

	return ret, nil
}

func decodeRicSubscriptionRequestIE(rsrIeC *C.RICsubscriptionRequest_IEs_t) (*e2appducontents.RicsubscriptionRequestIes, error) {
	//	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsrIeC)
	ret := new(e2appducontents.RicsubscriptionRequestIes)
	//
	switch rsrIeC.value.present {
	case C.RICsubscriptionRequest_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsrIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionRequest_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes5{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeRanFunctionIDBytes(rsrIeC.value.choice[0:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionRequest_IEs__value_PR_RICsubscriptionDetails:
		rsDet, err := decodeRicSubscriptionDetailsBytes(rsrIeC.value.choice[0:64])
		if err != nil {
			return nil, fmt.Errorf("decodeRicSubscriptionDetailsBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes30 = &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes30{
			Id:          int32(v1beta1.ProtocolIeIDRicsubscriptionDetails),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rsDet,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionRequest_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionRequestIE(). %v not yet implemneted", rsrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionRequestIE(). unexpected choice %v", rsrIeC.value.present)
	}

	return ret, nil
}

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
		ranal, err := decodeRicActionNotAdmittedListBytes(rsrIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes18 = &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes18{
			Id:          int32(v1beta1.ProtocolIeIDRicactionsNotAdmitted),
			Value:       ranal,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICsubscriptionResponse_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionResponseIE(). No components present.\n%v", rsrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionResponseIE(). unexpected choice %v", rsrIeC.value.present)
	}

	return ret, nil
}

func decodeRANfunctionItemIes(rfiIesValC *C.struct_RANfunction_ItemIEs__value) (*e2appducontents.RanfunctionItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfiIesValC, rfiIesValC)

	switch present := rfiIesValC.present; present {
	case C.RANfunction_ItemIEs__value_PR_RANfunction_Item:

		rfiIes := e2appducontents.RanfunctionItemIes{
			E2ApProtocolIes10: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
				Id:          int32(v1beta1.ProtocolIeIDRanfunctionItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		rfi, err := decodeRanFunctionItemBytes(rfiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeRANfunctionItemIes() %s", err.Error())
		}
		rfiIes.GetE2ApProtocolIes10().Value = rfi
		return &rfiIes, nil
	default:
		return nil, fmt.Errorf("error decoding RanFunctionItemIE - present %v not supported", present)
	}
}

func decodeRANfunctionIDItemIes(rfIDiIesValC *C.struct_RANfunctionID_ItemIEs__value) (*e2appducontents.RanfunctionIdItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := rfIDiIesValC.present; present {
	case C.RANfunctionID_ItemIEs__value_PR_RANfunctionID_Item:

		rfIDiIes := e2appducontents.RanfunctionIdItemIes{
			RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
				Id:          int32(v1beta1.ProtocolIeIDRanfunctionIDItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		rfi, err := decodeRanFunctionIDItemBytes(rfIDiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeRANfunctionIdItemIes() %s", err.Error())
		}
		rfIDiIes.GetRanFunctionIdItemIes6().Value = rfi
		return &rfIDiIes, nil
	default:
		return nil, fmt.Errorf("error decoding RanFunctionIDItemIE - present %v not supported", present)
	}
}

func decodeRANfunctionIDCauseItemIes(rfIDciIesValC *C.struct_RANfunctionIDcause_ItemIEs__value) (*e2appducontents.RanfunctionIdcauseItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDciIesValC, rfIDciIesValC)

	switch present := rfIDciIesValC.present; present {
	case C.RANfunctionIDcause_ItemIEs__value_PR_RANfunctionIDcause_Item:

		rfIDiIes := e2appducontents.RanfunctionIdcauseItemIes{
			RanFunctionIdcauseItemIes7: &e2appducontents.RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
				Id:          int32(v1beta1.ProtocolIeIDRanfunctionIeCauseItem),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		}
		rfi, err := decodeRanFunctionIDcauseItemBytes(rfIDciIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeRANfunctionIdcauseItemIes() %s", err.Error())
		}
		rfIDiIes.GetRanFunctionIdcauseItemIes7().Value = rfi
		return &rfIDiIes, nil
	default:
		return nil, fmt.Errorf("error decoding RanFunctionIDCauseItemIE - present %v not supported", present)
	}
}

func decodeRicActionAdmittedIDItemIes(raaiIesValC *C.struct_RICaction_Admitted_ItemIEs__value) (*e2appducontents.RicactionAdmittedItemIes, error) {
	//fmt.Printf("Value %T %v\n", raaiIesValC, raaiIesValC)

	switch present := raaiIesValC.present; present {
	case C.RICaction_Admitted_ItemIEs__value_PR_RICaction_Admitted_Item:

		raaiIes := e2appducontents.RicactionAdmittedItemIes{
			Id:          int32(v1beta1.ProtocolIeIDRicactionAdmittedItem),
			Value:       decodeRicActionAdmittedItemBytes(raaiIesValC.choice),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		return &raaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding RicactionAdmittedItemIes - present %v. not supported", present)
	}
}

func decodeRicActionNotAdmittedIDItemIes(ranaiIesValC *C.struct_RICaction_NotAdmitted_ItemIEs__value) (*e2appducontents.RicactionNotAdmittedItemIes, error) {
	//fmt.Printf("Value %T %v\n", ranaiIesValC, ranaiIesValC)

	switch present := ranaiIesValC.present; present {
	case C.RICaction_NotAdmitted_ItemIEs__value_PR_RICaction_NotAdmitted_Item:
		rana, err := decodeRicActionNotAdmittedItemBytes(ranaiIesValC.choice[:24])
		if err != nil {
			return nil, fmt.Errorf("decodeRicActionNotAdmittedItemBytes() %s", err.Error())
		}
		ranaiIes := e2appducontents.RicactionNotAdmittedItemIes{
			Id:          int32(v1beta1.ProtocolIeIDRicactionNotAdmittedItem),
			Value:       rana,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		return &ranaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding RicactionNotAdmittedItemIes - present %v. not supported", present)
	}
}

func decodeRicIndicationIE(riIeC *C.RICindication_IEs_t) (*e2appducontents.RicindicationIes, error) {
	//fmt.Printf("Handling E2SetupReqIE %+v\n", riIeC)
	ret := new(e2appducontents.RicindicationIes)

	switch riIeC.value.present {
	case C.RICindication_IEs__value_PR_RANfunctionID:
		rfID := decodeRanFunctionIDBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes5 = &e2appducontents.RicindicationIes_RicindicationIes5{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICactionID:
		raID := decodeRicActionIDBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes15 = &e2appducontents.RicindicationIes_RicindicationIes15{
			Id:          int32(v1beta1.ProtocolIeIDRicactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       raID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICcallProcessID:
		rcpID := decodeRicCallProcessIDBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes20 = &e2appducontents.RicindicationIes_RicindicationIes20{
			Id:          int32(v1beta1.ProtocolIeIDRiccallProcessID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcpID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICindication_IEs__value_PR_RICindicationHeader:
		rih := decodeRicIndicationHeaderBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes25 = &e2appducontents.RicindicationIes_RicindicationIes25{
			Id:          int32(v1beta1.ProtocolIeIDRicindicationHeader),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rih,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICindicationMessage:
		rim := decodeRicIndicationMessageBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes26 = &e2appducontents.RicindicationIes_RicindicationIes26{
			Id:          int32(v1beta1.ProtocolIeIDRicindicationMessage),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rim,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICindicationSN:
		risn := decodeRicIndicationSnBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes27 = &e2appducontents.RicindicationIes_RicindicationIes27{
			Id:          int32(v1beta1.ProtocolIeIDRicindicationSn),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       risn,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICindication_IEs__value_PR_RICindicationType:
		rit := decodeRicIndicationTypeBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes28 = &e2appducontents.RicindicationIes_RicindicationIes28{
			Id:          int32(v1beta1.ProtocolIeIDRicindicationType),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rit,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICrequestID:
		rrID := decodeRicRequestIDBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes29 = &e2appducontents.RicindicationIes_RicindicationIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rrID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicIndicationIE(). %v not yet implemneted", riIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicIndicationIE(). unexpected choice %v", riIeC.value.present)
	}

	return ret, nil
}

func decodeRicActionToBeSetupItemIes(ratbsIesValC *C.struct_RICaction_ToBeSetup_ItemIEs__value) (*e2appducontents.RicactionToBeSetupItemIes, error) {
	//fmt.Printf("Value %T %v\n", ratbsIesValC, ratbsIesValC)

	switch present := ratbsIesValC.present; present {
	case C.RICaction_ToBeSetup_ItemIEs__value_PR_RICaction_ToBeSetup_Item:
		ratbsIIes := e2appducontents.RicactionToBeSetupItemIes{
			Id:          int32(v1beta1.ProtocolIeIDRicactionToBeSetupItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		ratbsI, err := decodeRicActionToBeSetupItemBytes(ratbsIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeRicActionToBeSetupItemBytes() %s", err.Error())
		}
		ratbsIIes.Value = ratbsI
		return &ratbsIIes, nil
	default:
		return nil, fmt.Errorf("error decoding RicactionToBeSetupItemIes - present %v not supported", present)
	}
}

func decodeRicSubscriptionDeleteRequestIE(rsrdIeC *C.RICsubscriptionDeleteRequest_IEs_t) (*e2appducontents.RicsubscriptionDeleteRequestIes, error) {
	//	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsrdIeC)
	ret := new(e2appducontents.RicsubscriptionDeleteRequestIes)
	//
	switch rsrdIeC.value.present {
	case C.RICsubscriptionDeleteRequest_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsrdIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionDeleteRequest_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes5{
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeRanFunctionIDBytes(rsrdIeC.value.choice[0:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionDeleteRequest_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteRequestIE(). %v not yet implemneted", rsrdIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteRequestIE(). unexpected choice %v", rsrdIeC.value.present)
	}

	return ret, nil
}

func decodeRicSubscriptionDeleteResponseIE(rsdrIeC *C.RICsubscriptionDeleteResponse_IEs_t) (*e2appducontents.RicsubscriptionDeleteResponseIes, error) {
	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsdrIeC)
	ret := new(e2appducontents.RicsubscriptionDeleteResponseIes)

	switch rsdrIeC.value.present {
	case C.RICsubscriptionDeleteResponse_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionDeleteResponseIes_RicsubscriptionDeleteResponseIes5{
			Value:       decodeRanFunctionIDBytes(rsdrIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionDeleteResponse_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionDeleteResponseIes_RicsubscriptionDeleteResponseIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsdrIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionDeleteResponse_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteResponseIE(). %v not yet implemneted", rsdrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteResponseIE(). unexpected choice %v", rsdrIeC.value.present)
	}

	return ret, nil
}

func decodeRicSubscriptionDeleteFailureIE(rsdfIeC *C.RICsubscriptionDeleteFailure_IEs_t) (*e2appducontents.RicsubscriptionDeleteFailureIes, error) {
	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsdfIeC)
	ret := new(e2appducontents.RicsubscriptionDeleteFailureIes)

	switch rsdfIeC.value.present {
	case C.RICsubscriptionDeleteFailure_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes5{
			Value:       decodeRanFunctionIDBytes(rsdfIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionDeleteFailure_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsdfIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionDeleteFailure_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(rsdfIeC.value.choice[:16])
		if err != nil {
			return nil, fmt.Errorf("decodeCauseBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes1 = &e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes1{
			Id:          int32(v1beta1.ProtocolIeIDCause),
			Value:       cause,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionDeleteFailure_IEs__value_PR_CriticalityDiagnostics:
		cd, err := decodeCriticalityDiagnosticsBytes(rsdfIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeCriticalityDiagnosticsBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes2 = &e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes2{
			Id:          int32(v1beta1.ProtocolIeIDCriticalityDiagnostics),
			Value:       cd,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICsubscriptionDeleteFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteFailureIE(). %v not yet implemneted", rsdfIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteFailureIE(). unexpected choice %v", rsdfIeC.value.present)
	}

	return ret, nil
}

func decodeRicSubscriptionFailureIE(rsfIeC *C.RICsubscriptionFailure_IEs_t) (*e2appducontents.RicsubscriptionFailureIes, error) {
	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsfIeC)
	ret := new(e2appducontents.RicsubscriptionFailureIes)

	switch rsfIeC.value.present {
	case C.RICsubscriptionFailure_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes5{
			Value:       decodeRanFunctionIDBytes(rsfIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			Id:          int32(v1beta1.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes29{
			Id:          int32(v1beta1.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsfIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_CriticalityDiagnostics:
		cd, err := decodeCriticalityDiagnosticsBytes(rsfIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeCriticalityDiagnosticsBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes2 = &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes2{
			Id:          int32(v1beta1.ProtocolIeIDCriticalityDiagnostics),
			Value:       cd,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_RICaction_NotAdmitted_List:
		ranaL, err := decodeRicActionNotAdmittedListBytes(rsfIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeRicActionNotAdmittedListBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes18 = &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes18{
			Id:          int32(v1beta1.ProtocolIeIDRicactionsNotAdmitted),
			Value:       ranaL,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionFailureIE(). %v not yet implemneted", rsfIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionFailureIE(). unexpected choice %v", rsfIeC.value.present)
	}

	return ret, nil
}
