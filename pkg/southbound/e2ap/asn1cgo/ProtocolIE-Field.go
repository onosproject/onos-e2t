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
//#include "ResetRequest.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	"unsafe"

	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

func newRicServiceUpdateFailureIes1Cause(rsufIe *e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes1) (*C.RICserviceUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsufIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionFailure_IEs__value

	rsufCauseC, err := newCause(rsufIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rsufCauseC.present))
	copy(choiceC[8:16], rsufCauseC.choice[:8])

	ie := C.RICserviceUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateFailure_IEs__value{
			present: C.RICserviceUpdateFailure_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionFailureIe1Cause(rsfRanaIe *e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes1) (*C.RICsubscriptionFailure_IEs_t, error) {

	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsfRanaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the RICsubscriptionFailure_IEs__value

	rsdfCauseC, err := newCause(rsfRanaIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rsdfCauseC.present))
	copy(choiceC[8:16], rsdfCauseC.choice[:8])

	ie := C.RICsubscriptionFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionFailure_IEs__value{
			present: C.RICsubscriptionFailure_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteFailureIe1Cause(rsdfCauseIe *e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes1) (*C.RICsubscriptionDeleteFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsdfCauseIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
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

func newRicControlFailureIe1Cause(rcfCauseIe *e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes1) (*C.RICcontrolFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcfCauseIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICcontrolFailure_IEs__value

	rsdfCauseC, err := newCause(rcfCauseIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rsdfCauseC.present))
	copy(choiceC[8:16], rsdfCauseC.choice[:8])

	ie := C.RICcontrolFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolFailure_IEs__value{
			present: C.RICcontrolFailure_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newErrorIndicationIe1Cause(eiCauseIe *e2appducontents.ErrorIndicationIes_ErrorIndicationIes1) (*C.ErrorIndication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(eiCauseIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	rsdfCauseC, err := newCause(eiCauseIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rsdfCauseC.present))
	copy(choiceC[8:16], rsdfCauseC.choice[:8])

	ie := C.ErrorIndication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ErrorIndication_IEs__value{
			present: C.ErrorIndication_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupFailureIe1Cause(e2sfCauseIe *e2appducontents.E2SetupFailureIes_E2SetupFailureIes1) (*C.E2setupFailureIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfCauseIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [80]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	e2sfCauseC, err := newCause(e2sfCauseIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(e2sfCauseC.present))
	copy(choiceC[8:16], e2sfCauseC.choice[:8])

	ie := C.E2setupFailureIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupFailureIEs__value{
			present: C.E2setupFailureIEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newResetRequestIes1Cause(rrCauseIe *e2appducontents.ResetRequestIes_ResetRequestIes1) (*C.ResetRequestIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rrCauseIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [40]byte{} // The size of the ResetRequest_IEs__value

	rrCauseIeC, err := newCause(rrCauseIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rrCauseIeC.present))
	copy(choiceC[8:16], rrCauseIeC.choice[:8])

	ie := C.ResetRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ResetRequestIEs__value{
			present: C.ResetRequestIEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateFailureIes1Cause(e2cuaIe *e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes1) (*C.E2nodeConfigurationUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the ResetRequest_IEs__value

	rrCauseIeC, err := newCause(e2cuaIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rrCauseIeC.present))
	copy(choiceC[8:16], rrCauseIeC.choice[:8])

	ie := C.E2nodeConfigurationUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdateFailure_IEs__value{
			present: C.E2nodeConfigurationUpdateFailure_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateFailureIes1Cause(e2cuaIe *e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes1) (*C.E2connectionUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCause)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the ResetRequest_IEs__value

	rrCauseIeC, err := newCause(e2cuaIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rrCauseIeC.present))
	copy(choiceC[8:16], rrCauseIeC.choice[:8])

	ie := C.E2connectionUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateFailure_IEs__value{
			present: C.E2connectionUpdateFailure_IEs__value_PR_Cause,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newResetResponseIes2CriticalityDiagnostics(rrCritDiagsIe *e2appducontents.ResetResponseIes_ResetResponseIes2) (*C.ResetResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rrCritDiagsIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the ResetResponse_IEs__value

	rrCritDiagsIeC, err := newCriticalityDiagnostics(rrCritDiagsIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.ricRequestorID))))
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.iEsCriticalityDiagnostics))))

	ie := C.ResetResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ResetResponseIEs__value{
			present: C.ResetResponseIEs__value_PR_CriticalityDiagnostics,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateFailureIe2CriticalityDiagnostics(rsufCritDiagsIe *e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes2) (*C.RICserviceUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsufCritDiagsIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the ResetResponse_IEs__value

	rrCritDiagsIeC, err := newCriticalityDiagnostics(rsufCritDiagsIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.ricRequestorID))))
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.iEsCriticalityDiagnostics))))

	ie := C.RICserviceUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateFailure_IEs__value{
			present: C.RICserviceUpdateFailure_IEs__value_PR_CriticalityDiagnostics,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
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
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(rsdfCritDiagsC.iEsCriticalityDiagnostics))))

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

func newErrorIndicationIe2CriticalityDiagnostics(eiCritDiagsIe *e2appducontents.ErrorIndicationIes_ErrorIndicationIes2) (*C.ErrorIndication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(eiCritDiagsIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the ErrorIndication_IEs__value

	eiCritDiagsC, err := newCriticalityDiagnostics(eiCritDiagsIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(eiCritDiagsC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(eiCritDiagsC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(eiCritDiagsC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(eiCritDiagsC.ricRequestorID))))
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(eiCritDiagsC.iEsCriticalityDiagnostics))))

	ie := C.ErrorIndication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ErrorIndication_IEs__value{
			present: C.ErrorIndication_IEs__value_PR_CriticalityDiagnostics,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupIe2CriticalityDiagnostics(e2sfCritDiagsIe *e2appducontents.E2SetupFailureIes_E2SetupFailureIes2) (*C.E2setupFailureIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfCritDiagsIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [80]byte{} // The size of the ErrorIndication_IEs__value

	e2sfCritDiagsC, err := newCriticalityDiagnostics(e2sfCritDiagsIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2sfCritDiagsC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(e2sfCritDiagsC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(e2sfCritDiagsC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(e2sfCritDiagsC.ricRequestorID))))
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(e2sfCritDiagsC.iEsCriticalityDiagnostics))))

	ie := C.E2setupFailureIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupFailureIEs__value{
			present: C.E2setupFailureIEs__value_PR_CriticalityDiagnostics,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
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
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(rsfCritDiagsC.iEsCriticalityDiagnostics))))

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

func newE2connectionUpdateFailureIes2CriticalityDiagnostics(e2cufIe *e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes2) (*C.E2connectionUpdateFailure_IEs_t, error) {

	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cufIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the ResetResponse_IEs__value

	rrCritDiagsIeC, err := newCriticalityDiagnostics(e2cufIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.ricRequestorID))))
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.iEsCriticalityDiagnostics))))

	ie := C.E2connectionUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateFailure_IEs__value{
			present: C.E2connectionUpdateFailure_IEs__value_PR_CriticalityDiagnostics,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateFailureIes2CriticalityDiagnostics(e2ncufIe *e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes2) (*C.E2nodeConfigurationUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncufIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDCriticalityDiagnostics)
	if err != nil {
		return nil, err
	}

	choiceC := [64]byte{} // The size of the ResetResponse_IEs__value

	rrCritDiagsIeC, err := newCriticalityDiagnostics(e2ncufIe.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCode))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.triggeringMessage))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.procedureCriticality))))
	binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.ricRequestorID))))
	binary.LittleEndian.PutUint64(choiceC[32:], uint64(uintptr(unsafe.Pointer(rrCritDiagsIeC.iEsCriticalityDiagnostics))))

	ie := C.E2nodeConfigurationUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdateFailure_IEs__value{
			present: C.E2nodeConfigurationUpdateFailure_IEs__value_PR_CriticalityDiagnostics,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateIes3GlobalE2nodeID(esIe *e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes3) (*C.E2nodeConfigurationUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDGlobalE2nodeID)
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

	ie := C.E2nodeConfigurationUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdate_IEs__value{
			present: C.E2nodeConfigurationUpdate_IEs__value_PR_GlobalE2node_ID,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDGlobalE2nodeID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDGlobalRicID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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

func newRicControlRequestIe5RanFunctionID(rcrRfIe *e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes5) (*C.RICcontrolRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionIDC := newRanFunctionID(rcrRfIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICcontrolRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolRequest_IEs__value{
			present: C.RICcontrolRequest_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlFailureIe5RanFunctionID(rcfRfIe *e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes5) (*C.RICcontrolFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcfRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICcontrolFailureIEs__value_u

	ranFunctionIDC := newRanFunctionID(rcfRfIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICcontrolFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolFailure_IEs__value{
			present: C.RICcontrolFailure_IEs__value_PR_RANfunctionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlAcknowledgeIe5RanFunctionID(rcaRfIe *e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes5) (*C.RICcontrolAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcaRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ranFunctionIDC := newRanFunctionID(rcaRfIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.RICcontrolAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolAcknowledge_IEs__value{
			present: C.RICcontrolAcknowledge_IEs__value_PR_RANfunctionID,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
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

func newErrorIndicationIe5RanFunctionID(eiRfIe *e2appducontents.ErrorIndicationIes_ErrorIndicationIes5) (*C.ErrorIndication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(eiRfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the ErrorIndication_IEs__value_u

	ranFunctionIDC := newRanFunctionID(eiRfIe.Value)

	//fmt.Printf("Assigning to choice of ErrorIndicationIE %v \n", ranFunctionIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ranFunctionIDC))

	ie := C.ErrorIndication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ErrorIndication_IEs__value{
			present: C.ErrorIndication_IEs__value_PR_RANfunctionID,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsAccepted)
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

func newRicServiceQueryIes9RanFunctionsAccepted(esIe *e2appducontents.RicserviceQueryIes_RicserviceQueryIes9) (*C.RICserviceQuery_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsAccepted)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the RICserviceQueryIEs__value_u

	ranFunctionsIDListC, err := newRanFunctionsIDList(esIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsIDList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsIDListC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ranFunctionsIDListC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(ranFunctionsIDListC.list.size))

	ie := C.RICserviceQuery_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceQuery_IEs__value{
			present: C.RICserviceQuery_IEs__value_PR_RANfunctionsID_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateAcknowledgeIes9RanFunctionsAccepted(esIe *e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes9) (*C.RICserviceUpdateAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsAccepted)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the RICserviceUpdateAcknowledge_IEs__value_u

	ranFunctionsIDListC, err := newRanFunctionsIDList(esIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsIDList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsIDListC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ranFunctionsIDListC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(ranFunctionsIDListC.list.size))

	ie := C.RICserviceUpdateAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateAcknowledge_IEs__value{
			present: C.RICserviceUpdateAcknowledge_IEs__value_PR_RANfunctionsID_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicSubscriptionDeleteRequiredIesRicSubscriptionListWithCause(esIe *e2appducontents.RicsubscriptionDeleteRequiredIes) (*C.RICsubscriptionDeleteRequired_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRICsubscriptionToBeRemoved)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the RICsubscriptionDeleteRequired__value_u

	ranFunctionsIDListC, err := newRicSubscriptionListWithCause(esIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicSubscriptionListWithCause() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsIDListC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ranFunctionsIDListC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(ranFunctionsIDListC.list.size))

	ie := C.RICsubscriptionDeleteRequired_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionDeleteRequired_IEs__value{
			present: C.RICsubscriptionDeleteRequired_IEs__value_PR_RICsubscription_List_withCause,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsAdded)
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

func newRicServiceUpdateIes10RanFunctionAddedList(esIe *e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes10) (*C.RICserviceUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsAdded)
	if err != nil {
		return nil, err
	}

	listC := [48]byte{} // The size of theRICserviceUpdateIEs__value_u

	ranFunctionsListC, err := newRanFunctionsList(esIe.GetValue())
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ranFunctionsListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ranFunctionsListC.list.size))

	ie := C.RICserviceUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdate_IEs__value{
			present: C.RICserviceUpdate_IEs__value_PR_RANfunctions_List,
			choice:  listC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateIes11RanFunctionDeletedList(esIe *e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes11) (*C.RICserviceUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsDeleted)
	if err != nil {
		return nil, err
	}

	listC := [48]byte{} // The size of theRICserviceUpdateIEs__value_u

	ranFunctionsListC, err := newRanFunctionsIDList(esIe.GetValue())
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsIDList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ranFunctionsListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ranFunctionsListC.list.size))

	ie := C.RICserviceUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdate_IEs__value{
			present: C.RICserviceUpdate_IEs__value_PR_RANfunctionsID_List,
			choice:  listC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateIes12RanFunctionModifiedList(esIe *e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes12) (*C.RICserviceUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsModified)
	if err != nil {
		return nil, err
	}

	listC := [48]byte{} // The size of theRICserviceUpdateIEs__value_u

	ranFunctionsListC, err := newRanFunctionsList(esIe.GetValue())
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ranFunctionsListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ranFunctionsListC.list.size))

	ie := C.RICserviceUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdate_IEs__value{
			present: C.RICserviceUpdate_IEs__value_PR_RANfunctions_List,
			choice:  listC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateAcknowledgeIes13RanFunctionsRejected(rsuaIe *e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes13) (*C.RICserviceUpdateAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsRejected)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the RICserviceUpdateAcknowledge_IEs__value_u

	ranFunctionsIDCauseList, err := newRanFunctionsIDcauseList(rsuaIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRanFunctionsIDcauseList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ranFunctionsIDCauseList.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(ranFunctionsIDCauseList.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(ranFunctionsIDCauseList.list.size))

	ie := C.RICserviceUpdateAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateAcknowledge_IEs__value{
			present: C.RICserviceUpdateAcknowledge_IEs__value_PR_RANfunctionsIDcause_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIe13RanFunctionsRejected(esIe *e2appducontents.E2SetupResponseIes_E2SetupResponseIes13) (*C.E2setupResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionsRejected)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicactionID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicactionsAdmitted)
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

func newRicSubscriptionResponseIe18RicActionNotAdmittedList(rsrRanaIe *e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes18) (*C.RICsubscriptionResponse_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRanaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicactionsNotAdmitted)
	if err != nil {
		return nil, err
	}

	listC := [48]byte{} // The size of the E2setupResponseIEs__value_u

	ricActionNotAdmittedListC, err := newRicActionNotAdmittedList(rsrRanaIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicActionAdmittedList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(listC[0:], uint64(uintptr(unsafe.Pointer(ricActionNotAdmittedListC.list.array))))
	binary.LittleEndian.PutUint32(listC[8:], uint32(ricActionNotAdmittedListC.list.count))
	binary.LittleEndian.PutUint32(listC[12:], uint32(ricActionNotAdmittedListC.list.size))

	ie := C.RICsubscriptionResponse_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscriptionResponse_IEs__value{
			present: C.RICsubscriptionResponse_IEs__value_PR_RICaction_NotAdmitted_List,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccallProcessID)
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

func newRicControlRequestIe20RiccallProcessID(rcrIe20 *e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes20) (*C.RICcontrolRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe20.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccallProcessID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricCallProcessIDC := newRicCallProcessID(rcrIe20.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricCallProcessIDC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricCallProcessIDC.size))

	ie := C.RICcontrolRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolRequest_IEs__value{
			present: C.RICcontrolRequest_IEs__value_PR_RICcallProcessID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlFailureIe20RiccallProcessID(rcfIe20 *e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes20) (*C.RICcontrolFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcfIe20.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccallProcessID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricCallProcessIDC := newRicCallProcessID(rcfIe20.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricCallProcessIDC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricCallProcessIDC.size))

	ie := C.RICcontrolFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolFailure_IEs__value{
			present: C.RICcontrolFailure_IEs__value_PR_RICcallProcessID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlAcknowledgeIe20RiccallProcessID(rcrIe20 *e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes20) (*C.RICcontrolAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe20.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccallProcessID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricCallProcessIDC := newRicCallProcessID(rcrIe20.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricCallProcessIDC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricCallProcessIDC.size))

	ie := C.RICcontrolAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolAcknowledge_IEs__value{
			present: C.RICcontrolAcknowledge_IEs__value_PR_RICcallProcessID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlRequestIe21RiccontrolAckRequest(rcrIe21 *e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes21) (*C.RICcontrolRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe21.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccontrolAckRequest)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricControlAckRequestC, err := newRicControlAckRequest(rcrIe21.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicControlAckRequest() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*ricControlAckRequestC))

	ie := C.RICcontrolRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolRequest_IEs__value{
			present: C.RICcontrolRequest_IEs__value_PR_RICcontrolAckRequest,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlRequestIe22RiccontrolHeader(rcrIe22 *e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes22) (*C.RICcontrolRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe22.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccontrolHeader)
	if err != nil {
		return nil, err
	}

	//ToDo - double-check number of bytes required here
	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricControlHeaderC := newRicControlHeader(rcrIe22.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricControlHeaderC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricControlHeaderC.size))

	ie := C.RICcontrolRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolRequest_IEs__value{
			present: C.RICcontrolRequest_IEs__value_PR_RICcontrolHeader,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlRequestIe23RiccontrolMessage(rcrIe23 *e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes23) (*C.RICcontrolRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe23.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccontrolMessage)
	if err != nil {
		return nil, err
	}

	//ToDo - double-check number of bytes required here
	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricControlMessageC := newRicControlMessage(rcrIe23.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricControlMessageC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricControlMessageC.size))

	ie := C.RICcontrolRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolRequest_IEs__value{
			present: C.RICcontrolRequest_IEs__value_PR_RICcontrolMessage,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlAcknowledgeIe24RiccontrolStatus(rcrIe24 *e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes24) (*C.RICcontrolAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe24.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccontrolStatus)
	if err != nil {
		return nil, err
	}

	//ToDo - double-check number of bytes required here
	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricControlStatusC, err := newRicControlStatus(rcrIe24.Value)
	if err != nil {
		return nil, fmt.Errorf("newRicControlAckRequest() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*ricControlStatusC))

	ie := C.RICcontrolAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolAcknowledge_IEs__value{
			present: C.RICcontrolAcknowledge_IEs__value_PR_RICcontrolStatus,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicindicationHeader)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicindicationMessage)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicindicationSn)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicindicationType)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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

func newRicControlFailureIe29RicRequestID(rcfRrIDIe *e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes29) (*C.RICcontrolFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcfRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICcontrolFaiure_IEs__value_u

	ricRequestIDC := newRicRequestID(rcfRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICcontrolFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolFailure_IEs__value{
			present: C.RICcontrolFailure_IEs__value_PR_RICrequestID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlRequestIe29RicRequestID(rcrRrIDIe *e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes29) (*C.RICcontrolRequest_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{} // The size of the RICindication_IEs__value_u

	ricRequestIDC := newRicRequestID(rcrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionRequestIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICcontrolRequest_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolRequest_IEs__value{
			present: C.RICcontrolRequest_IEs__value_PR_RICrequestID,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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

func newErrorIndicationIe29RicRequestID(eiRrIDIe *e2appducontents.ErrorIndicationIes_ErrorIndicationIes29) (*C.ErrorIndication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(eiRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{}

	ricRequestIDC := newRicRequestID(eiRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.ErrorIndication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ErrorIndication_IEs__value{
			present: C.ErrorIndication_IEs__value_PR_RICrequestID,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
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

func newRicControlAcknowledgeIe29RicRequestID(rsrRrIDIe *e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes29) (*C.RICcontrolAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsrRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicrequestID)
	if err != nil {
		return nil, err
	}

	choiceC := [40]byte{}

	ricRequestIDC := newRicRequestID(rsrRrIDIe.Value)

	//fmt.Printf("Assigning to choice of RicSubscriptionResponseIE %v \n", ricRequestIDC)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ricRequestIDC.ricRequestorID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricRequestIDC.ricInstanceID))

	ie := C.RICcontrolAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolAcknowledge_IEs__value{
			present: C.RICcontrolAcknowledge_IEs__value_PR_RICrequestID,
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicsubscriptionDetails)
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

func newE2setupFailureIe31TimeToWait(e2sfIe *e2appducontents.E2SetupFailureIes_E2SetupFailureIes31) (*C.E2setupFailureIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTimeToWait)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [80]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	e2sfTtwC, err := newTimeToWait(e2sfIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*e2sfTtwC))
	//copy(choiceC[8:16], e2sfCauseC.choice[:8])

	ie := C.E2setupFailureIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupFailureIEs__value{
			present: C.E2setupFailureIEs__value_PR_TimeToWait,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateFailureIe31TimeToWait(e2sfIe *e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes31) (*C.RICserviceUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTimeToWait)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	e2sfTtwC, err := newTimeToWait(e2sfIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*e2sfTtwC))
	//copy(choiceC[8:16], e2sfCauseC.choice[:8])

	ie := C.RICserviceUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateFailure_IEs__value{
			present: C.RICserviceUpdateFailure_IEs__value_PR_TimeToWait,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateFailureIes31TimeToWait(e2cuafIe *e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes31) (*C.E2connectionUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuafIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTimeToWait)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	e2sfTtwC, err := newTimeToWait(e2cuafIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*e2sfTtwC))
	//copy(choiceC[8:16], e2sfCauseC.choice[:8])

	ie := C.E2connectionUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateFailure_IEs__value{
			present: C.E2connectionUpdateFailure_IEs__value_PR_TimeToWait,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateFailureIes31TimeToWait(e2cuaIe *e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes31) (*C.E2nodeConfigurationUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTimeToWait)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{} // The size of the RICsubscriptionDeleteFailure_IEs__value

	e2sfTtwC, err := newTimeToWait(e2cuaIe.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*e2sfTtwC))
	//copy(choiceC[8:16], e2sfCauseC.choice[:8])

	ie := C.E2nodeConfigurationUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdateFailure_IEs__value{
			present: C.E2nodeConfigurationUpdateFailure_IEs__value_PR_TimeToWait,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlFailureIe32RiccontrolOutcome(rcfIe32 *e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes32) (*C.RICcontrolFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcfIe32.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccontrolOutcome)
	if err != nil {
		return nil, err
	}

	//ToDo - double-check number of bytes required here
	choiceC := [40]byte{} // The size of the RICcontrolFailureIEs__value_u

	ricControlOutcomeC := newRicControlOutcome(rcfIe32.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricControlOutcomeC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricControlOutcomeC.size))

	ie := C.RICcontrolFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolFailure_IEs__value{
			present: C.RICcontrolFailure_IEs__value_PR_RICcontrolOutcome,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicControlAcknowledgeIe32RiccontrolOutcome(rcrIe32 *e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes32) (*C.RICcontrolAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rcrIe32.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRiccontrolOutcome)
	if err != nil {
		return nil, err
	}

	//ToDo - double-check number of bytes required here
	choiceC := [40]byte{} // The size of the E2setupResponseIEs__value_u

	ricControlOutcomeC := newRicControlOutcome(rcrIe32.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(ricControlOutcomeC.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ricControlOutcomeC.size))

	ie := C.RICcontrolAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICcontrolAcknowledge_IEs__value{
			present: C.RICcontrolAcknowledge_IEs__value_PR_RICcontrolOutcome,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupRequestIe33E2nodeComponentConfigUpdateList(e2sfIe *e2appducontents.E2SetupRequestIes_E2SetupRequestIes33) (*C.E2setupRequestIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdate)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2nodeComponentConfigUpdate_IEs__value_u

	e2ncuIeC, err := newE2nodeComponentConfigUpdateList(e2sfIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2setupRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupRequestIEs__value{
			present: C.E2setupRequestIEs__value_PR_E2nodeComponentConfigUpdate_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateIes33E2nodeComponentConfigUpdateList(e2ncuIe *e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes33) (*C.E2nodeConfigurationUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdate)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2nodeComponentConfigUpdate_IEs__value_u

	e2ncuIeC, err := newE2nodeComponentConfigUpdateList(e2ncuIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2nodeConfigurationUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdate_IEs__value{
			present: C.E2nodeConfigurationUpdate_IEs__value_PR_E2nodeComponentConfigUpdate_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIe35E2nodeComponentConfigUpdateAckList(e2srIe *e2appducontents.E2SetupResponseIes_E2SetupResponseIes35) (*C.E2setupResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2srIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2nodeComponentConfigUpdate_IEs__value_u

	e2ncuaIeC, err := newE2nodeComponentConfigUpdateAckList(e2srIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuaIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuaIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuaIeC.list.size))

	ie := C.E2setupResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupResponseIEs__value{
			present: C.E2setupResponseIEs__value_PR_E2nodeComponentConfigUpdateAck_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateAcknowledgeIes35E2nodeComponentConfigUpdateAckList(e2ncuaIe *e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes_E2NodeConfigurationUpdateAcknowledgeIes35) (*C.E2nodeConfigurationUpdateAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2nodeComponentConfigUpdate_IEs__value_u

	e2ncuaIeC, err := newE2nodeComponentConfigUpdateAckList(e2ncuaIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateAckList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuaIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuaIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuaIeC.list.size))

	ie := C.E2nodeConfigurationUpdateAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdateAcknowledge_IEs__value{
			present: C.E2nodeConfigurationUpdateAcknowledge_IEs__value_PR_E2nodeComponentConfigUpdateAck_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateAck39E2connectionUpdateList(e2cuaIe *e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes39) (*C.E2connectionUpdateAck_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionSetup)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2connectionUpdate_IEs__value_u

	e2ncuIeC, err := newE2connectionUpdateList(e2cuaIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2connectionUpdateAck_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateAck_IEs__value{
			present: C.E2connectionUpdateAck_IEs__value_PR_E2connectionUpdate_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateAck40E2connectionSetupFailedList(e2cuaIe *e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes40) (*C.E2connectionUpdateAck_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionSetupFailed)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2connectionUpdate_IEs__value_u

	e2ncuIeC, err := newE2connectionSetupFailedList(e2cuaIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2connectionSetupFailedList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2connectionUpdateAck_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateAck_IEs__value{
			present: C.E2connectionUpdateAck_IEs__value_PR_E2connectionSetupFailed_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateIe44E2connectionUpdateList(e2cuIe *e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes44) (*C.E2connectionUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionUpdateAdd)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2connectionUpdate_IEs__value_u

	e2ncuIeC, err := newE2connectionUpdateList(e2cuIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2connectionUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdate_IEs__value{
			present: C.E2connectionUpdate_IEs__value_PR_E2connectionUpdate_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateIe45E2connectionUpdateList(e2cuIe *e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes45) (*C.E2connectionUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionUpdateModify)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2connectionUpdate_IEs__value_u

	e2ncuIeC, err := newE2connectionUpdateList(e2cuIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2connectionUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdate_IEs__value{
			present: C.E2connectionUpdate_IEs__value_PR_E2connectionUpdate_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateIe46E2connectionUpdateRemoveList(e2cuIe *e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes46) (*C.E2connectionUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionUpdateRemove)
	if err != nil {
		return nil, err
	}

	choiceC := [48]byte{} // The size of the E2connectionUpdate_IEs__value_u

	e2ncuIeC, err := newE2connectionUpdateRemoveList(e2cuIe.Value)
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))

	ie := C.E2connectionUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdate_IEs__value{
			present: C.E2connectionUpdate_IEs__value_PR_E2connectionUpdateRemove_List,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupFailureIe48Tnlinformation(e2sfIe *e2appducontents.E2SetupFailureIes_E2SetupFailureIes48) (*C.E2setupFailureIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTNLinformation)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2setupFailureIEs__value_u

	tnlInfoC, err := newTnlinformation(e2sfIe.Value)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Assigning to choice of E2setupFailureIE %v \n", globalRicIDC)

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(tnlInfoC.tnlAddress.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(tnlInfoC.tnlAddress.size))
	binary.LittleEndian.PutUint32(choiceC[16:], uint32(tnlInfoC.tnlAddress.bits_unused))
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(tnlInfoC.tnlPort))))

	ie := C.E2setupFailureIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupFailureIEs__value{
			present: C.E2setupFailureIEs__value_PR_TNLinformation,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newErrorIndicationIe49TransactionID(eiRrIDIe *e2appducontents.ErrorIndicationIes_ErrorIndicationIes49) (*C.ErrorIndication_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(eiRrIDIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{}

	transactionID := newTransactionID(eiRrIDIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.ErrorIndication_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ErrorIndication_IEs__value{
			present: C.ErrorIndication_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupRequestIe49TransactionID(e2srIe *e2appducontents.E2SetupRequestIes_E2SetupRequestIes49) (*C.E2setupRequestIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2srIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(e2srIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2setupRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupRequestIEs__value{
			present: C.E2setupRequestIEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupResponseIe49TransactionID(e2srIe *e2appducontents.E2SetupResponseIes_E2SetupResponseIes49) (*C.E2setupResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2srIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [112]byte{}

	transactionID := newTransactionID(e2srIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2setupResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupResponseIEs__value{
			present: C.E2setupResponseIEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2setupFailureIe49TransactionID(e2sfIe *e2appducontents.E2SetupFailureIes_E2SetupFailureIes49) (*C.E2setupFailureIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2sfIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [80]byte{}

	transactionID := newTransactionID(e2sfIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2setupFailureIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2setupFailureIEs__value{
			present: C.E2setupFailureIEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateIe49TransactionID(e2cuIe *e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes49) (*C.E2connectionUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(e2cuIe.GetValue())
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2connectionUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdate_IEs__value{
			present: C.E2connectionUpdate_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateAck49TransactionID(e2cuaIe *e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes49) (*C.E2connectionUpdateAck_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuaIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(e2cuaIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2connectionUpdateAck_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateAck_IEs__value{
			present: C.E2connectionUpdateAck_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2connectionUpdateFailureIes49TransactionID(e2cufIe *e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes49) (*C.E2connectionUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cufIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{}

	transactionID := newTransactionID(e2cufIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2connectionUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateFailure_IEs__value{
			present: C.E2connectionUpdateFailure_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateIes49TransactionID(e2ncuIe *e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes49) (*C.E2nodeConfigurationUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(e2ncuIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2nodeConfigurationUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdate_IEs__value{
			present: C.E2nodeConfigurationUpdate_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateAcknowledgeIes49TransactionID(e2ncuIe *e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes_E2NodeConfigurationUpdateAcknowledgeIes49) (*C.E2nodeConfigurationUpdateAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(e2ncuIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2nodeConfigurationUpdateAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdateAcknowledge_IEs__value{
			present: C.E2nodeConfigurationUpdateAcknowledge_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newE2nodeConfigurationUpdateFailureIes49TransactionID(e2ncuIe *e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes49) (*C.E2nodeConfigurationUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{}

	transactionID := newTransactionID(e2ncuIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.E2nodeConfigurationUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeConfigurationUpdateFailure_IEs__value{
			present: C.E2nodeConfigurationUpdateFailure_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newResetRequestIes49TransactionID(rrIe *e2appducontents.ResetRequestIes_ResetRequestIes49) (*C.ResetRequestIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rrIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [40]byte{}

	transactionID := newTransactionID(rrIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.ResetRequestIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ResetRequestIEs__value{
			present: C.ResetRequestIEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newResetResponseIes49TransactionID(rrIe *e2appducontents.ResetResponseIes_ResetResponseIes49) (*C.ResetResponseIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rrIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{}

	transactionID := newTransactionID(rrIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.ResetResponseIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_ResetResponseIEs__value{
			present: C.ResetResponseIEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateIes49TransactionID(rsuIe *e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes49) (*C.RICserviceUpdate_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(rsuIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.RICserviceUpdate_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdate_IEs__value{
			present: C.RICserviceUpdate_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateAcknowledgeIes49TransactionID(rsuIe *e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes49) (*C.RICserviceUpdateAcknowledge_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsuIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(rsuIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.RICserviceUpdateAcknowledge_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateAcknowledge_IEs__value{
			present: C.RICserviceUpdateAcknowledge_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceUpdateFailureIes49TransactionID(rsufIe *e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes49) (*C.RICserviceUpdateFailure_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rsufIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [64]byte{}

	transactionID := newTransactionID(rsufIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.RICserviceUpdateFailure_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceUpdateFailure_IEs__value{
			present: C.RICserviceUpdateFailure_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

func newRicServiceQueryIes49TransactionID(esIe *e2appducontents.RicserviceQueryIes_RicserviceQueryIes49) (*C.RICserviceQuery_IEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(esIe.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDTransactionID)
	if err != nil {
		return nil, err
	}

	//TODO: Size should be double-checked
	choiceC := [48]byte{}

	transactionID := newTransactionID(esIe.Value)
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(*transactionID))

	ie := C.RICserviceQuery_IEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICserviceQuery_IEs__value{
			present: C.RICserviceQuery_IEs__value_PR_TransactionID,
			choice:  choiceC,
		},
	}

	return &ie, nil
}

//func newE2nodeConfigurationUpdateIeE2nodeComponentConfigUpdateList(e2ncuIe *e2appducontents.E2NodeConfigurationUpdateIes) (*C.E2nodeConfigurationUpdate_IEs_t, error) {
//	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuIe.GetCriticality()))
//	if err != nil {
//		return nil, err
//	}
//	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdate)
//	if err != nil {
//		return nil, err
//	}
//
//	choiceC := [48]byte{} // The size of the E2nodeComponentConfigUpdate_IEs__value_u
//
//	e2ncuIeC, err := newE2nodeComponentConfigUpdateList(e2ncuIe.Value)
//	if err != nil {
//		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
//	}
//	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuIeC.list.array))))
//	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuIeC.list.count))
//	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuIeC.list.size))
//
//	ie := C.E2nodeConfigurationUpdate_IEs_t{
//		id:          idC,
//		criticality: critC,
//		value: C.struct_E2nodeConfigurationUpdate_IEs__value{
//			present: C.E2nodeConfigurationUpdate_IEs__value_PR_E2nodeComponentConfigUpdate_List,
//			choice:  choiceC,
//		},
//	}
//
//	return &ie, nil
//}

//func newE2nodeConfigurationUpdateAcknowledgeIeE2nodeComponentConfigUpdateList(e2ncuaIe *e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes) (*C.E2nodeConfigurationUpdateAcknowledge_IEs_t, error) {
//	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2ncuaIe.GetCriticality()))
//	if err != nil {
//		return nil, err
//	}
//	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck)
//	if err != nil {
//		return nil, err
//	}
//
//	choiceC := [48]byte{} // The size of the E2nodeComponentConfigUpdate_IEs__value_u
//
//	e2ncuaIeC, err := newE2nodeComponentConfigUpdateAckList(e2ncuaIe.Value)
//	if err != nil {
//		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateList() %s", err.Error())
//	}
//	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2ncuaIeC.list.array))))
//	binary.LittleEndian.PutUint32(choiceC[8:], uint32(e2ncuaIeC.list.count))
//	binary.LittleEndian.PutUint32(choiceC[12:], uint32(e2ncuaIeC.list.size))
//
//	ie := C.E2nodeConfigurationUpdateAcknowledge_IEs_t{
//		id:          idC,
//		criticality: critC,
//		value: C.struct_E2nodeConfigurationUpdateAcknowledge_IEs__value{
//			present: C.E2nodeConfigurationUpdateAcknowledge_IEs__value_PR_E2nodeComponentConfigUpdateAck_List,
//			choice:  choiceC,
//		},
//	}
//
//	return &ie, nil
//}

func newRANfunctionItemIEs(rfItemIes *e2appducontents.RanfunctionItemIes) (*C.RANfunction_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rfItemIes.GetE2ApProtocolIes10().GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionItem)
	if err != nil {
		return nil, err
	}

	choiceC := [120]byte{} // The size of the RANfunction_ItemIEs__value_u
	rfItemC := newRanFunctionItem(rfItemIes.GetE2ApProtocolIes10().GetValue())
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(rfItemC.ranFunctionID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(uintptr(unsafe.Pointer(rfItemC.ranFunctionDefinition.buf))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(rfItemC.ranFunctionDefinition.size))
	// Gap of 24 for the asn_struct_ctx_t belonging to OCTET STRING
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(rfItemC.ranFunctionRevision))
	binary.LittleEndian.PutUint64(choiceC[56:], uint64(uintptr(unsafe.Pointer(rfItemC.ranFunctionOID.buf))))
	binary.LittleEndian.PutUint64(choiceC[64:], uint64(rfItemC.ranFunctionOID.size))
	// Gap of 24 for the asn_struct_ctx_t belonging to PrintableString

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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionIDItem)
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

func newE2nodeComponentConfigUpdateItemIEs(e2nccuItemIes *e2appducontents.E2NodeComponentConfigUpdateItemIes) (*C.E2nodeComponentConfigUpdate_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccuItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdateItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigUpdate_ItemIEs__value_u
	e2nccuItemC, err := newE2nodeComponentConfigUpdateItem(e2nccuItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:8], uint64(e2nccuItemC.e2nodeComponentType))
	//ToDo - something wrong with conversion here
	binary.LittleEndian.PutUint64(choiceC[8:16], uint64(uintptr(unsafe.Pointer(e2nccuItemC.e2nodeComponentID))))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2nccuItemC.e2nodeComponentConfigUpdate.present))
	copy(choiceC[24:32], e2nccuItemC.e2nodeComponentConfigUpdate.choice[:8])

	rfItemIesC := C.E2nodeComponentConfigUpdate_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeComponentConfigUpdate_ItemIEs__value{
			present: C.E2nodeComponentConfigUpdate_ItemIEs__value_PR_E2nodeComponentConfigUpdate_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newE2nodeComponentConfigUpdateAckItemIEs(e2nccuaItemIes *e2appducontents.E2NodeComponentConfigUpdateAckItemIes) (*C.E2nodeComponentConfigUpdateAck_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccuaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigUpdate_ItemIEs__value_u
	e2nccuItemC, err := newE2nodeComponentConfigUpdateAckItem(e2nccuaItemIes.GetValue())
	if err != nil {
		return nil, fmt.Errorf("newE2nodeComponentConfigUpdateAckItemIEs() %s", err.Error())
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:8], uint64(e2nccuItemC.e2nodeComponentType))
	binary.LittleEndian.PutUint64(choiceC[8:16], uint64(uintptr(unsafe.Pointer(e2nccuItemC.e2nodeComponentID))))
	binary.LittleEndian.PutUint64(choiceC[16:24], uint64(e2nccuItemC.e2nodeComponentConfigUpdateAck.updateOutcome))
	binary.LittleEndian.PutUint64(choiceC[24:32], uint64(uintptr(unsafe.Pointer(e2nccuItemC.e2nodeComponentConfigUpdateAck.failureCause))))

	rfItemIesC := C.E2nodeComponentConfigUpdateAck_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeComponentConfigUpdateAck_ItemIEs__value{
			present: C.E2nodeComponentConfigUpdateAck_ItemIEs__value_PR_E2nodeComponentConfigUpdateAck_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newE2connectionUpdateItemIEs(e2cuItemIes *e2appducontents.E2ConnectionUpdateItemIes) (*C.E2connectionUpdate_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2cuItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionUpdateItem)
	if err != nil {
		return nil, err
	}

	choiceC := [112]byte{} // The size of the E2nodeComponentConfigUpdate_ItemIEs__value_u
	e2cuItemC, err := newE2connectionUpdateItem(e2cuItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2cuItemC.tnlInformation.tnlAddress.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2cuItemC.tnlInformation.tnlAddress.size))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2cuItemC.tnlInformation.tnlAddress.bits_unused))
	// Gap of 24 for the asn_struct_ctx_t belonging to BIT STRING
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2cuItemC.tnlInformation.tnlPort))))
	// Gap of 24 for the asn_struct_ctx_t belonging to BIT STRING
	binary.LittleEndian.PutUint64(choiceC[80:], uint64(e2cuItemC.tnlUsage))

	//ToDo - print it out
	//fmt.Printf("Choice is \n%v", hex.Dump(choiceC))

	rfItemIesC := C.E2connectionUpdate_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdate_ItemIEs__value{
			present: C.E2connectionUpdate_ItemIEs__value_PR_E2connectionUpdate_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newE2connectionUpdateRemoveItemIEs(e2curItemIes *e2appducontents.E2ConnectionUpdateRemoveItemIes) (*C.E2connectionUpdateRemove_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2curItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionUpdateRemoveItem)
	if err != nil {
		return nil, err
	}

	choiceC := [104]byte{} // The size of the E2connectionUpdateRemove_ItemIEs__value_u
	e2cuItemC, err := newE2connectionUpdateRemoveItem(e2curItemIes.GetValue())
	if err != nil {
		return nil, err
	}

	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2cuItemC.tnlInformation.tnlAddress.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2cuItemC.tnlInformation.tnlAddress.size))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2cuItemC.tnlInformation.tnlAddress.bits_unused))
	// Gap of 24 for the asn_struct_ctx_t belonging to BIT STRING
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2cuItemC.tnlInformation.tnlPort))))

	rfItemIesC := C.E2connectionUpdateRemove_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionUpdateRemove_ItemIEs__value{
			present: C.E2connectionUpdateRemove_ItemIEs__value_PR_E2connectionUpdateRemove_Item,
			choice:  choiceC,
		},
	}

	return &rfItemIesC, nil
}

func newE2nodeComponentConfigAdditionItemIEs(e2nccaItemIes *e2appducontents.E2NodeComponentConfigAdditionItemIes) (*C.E2nodeComponentConfigAddition_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigAddition_ItemIEs__value_u
	e2nccaItemC, err := newE2nodeComponentConfigAdditionItem(e2nccaItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2nccaItemC.e2nodeComponentInterfaceType))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2nccaItemC.e2nodeComponentID))
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration
	binary.LittleEndian.PutUint64(choiceC[40:], uint64(e2nccaItemC.e2nodeComponentConfiguration.present))
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2nccaItemC.e2nodeComponentConfiguration.choice))))

	e2nccaItemIesC := C.E2nodeComponentConfigAddition_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeComponentConfigAddition_ItemIEs__value{
			present: C.E2nodeComponentConfigaddition_ItemIEs__value_PR_E2nodeComponentConfigAddition_Item,
			choice:  choiceC,
		},
	}

	return &e2nccaItemIesC, nil
}

func newE2nodeComponentConfigAdditionAckItemIEs(e2nccaaItemIes *e2appducontents.E2NodeComponentConfigAdditionAckItemIes) (*C.E2nodeComponentConfigAdditionAck_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccaaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigAdditionAckItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigAddition_ItemIEs__value_u
	e2nccaaItemC, err := newE2nodeComponentConfigAdditionAckItem(e2nccaaItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentInterfaceType))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2nccaaItemC.e2nodeComponentID))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2nccaaItemC.e2nodeComponentConfigurationAck))
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration
	//binary.LittleEndian.PutUint64(choiceC[40:], uint64(e2nccaaItemC.e2nodeComponentConfiguration.present))
	//binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentConfiguration.choice))))

	e2nccaaItemIesC := C.E2nodeComponentConfigAdditionAck_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeComponentConfigAdditionAck_ItemIEs__value{
			present: C.E2nodeComponentConfigAdditionAck_ItemIEs__value_PR_E2nodeComponentConfigAdditionAck_Item,
			choice:  choiceC,
		},
	}

	return &e2nccaaItemIesC, nil
}

func newE2nodeComponentConfigRemovalAckItemIEs(e2nccraItemIes *e2appducontents.E2NodeComponentConfigRemovalAckItemIes) (*C.E2nodeComponentConfigRemovalAck_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccraItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigRemovalAckItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigRemovalAck_ItemIEs__value_u
	e2nccaaItemC, err := newE2nodeComponentConfigRemovalAckItem(e2nccraItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentInterfaceType))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2nccaaItemC.e2nodeComponentID))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2nccaaItemC.e2nodeComponentConfigurationAck))
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration
	//binary.LittleEndian.PutUint64(choiceC[40:], uint64(e2nccaaItemC.e2nodeComponentConfiguration.present))
	//binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentConfiguration.choice))))

	e2nccraItemIesC := C.E2nodeComponentConfigRemovalAck_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeComponentConfigRemovalAck_ItemIEs__value{
			present: C.E2nodeComponentConfigRemovalAck_ItemIEs__value_PR_E2nodeComponentConfigRemovalAck_Item,
			choice:  choiceC,
		},
	}

	return &e2nccraItemIesC, nil
}

func newE2nodeComponentConfigRemovalItemIEs(e2nccaaItemIes *e2appducontents.E2NodeComponentConfigRemovalItemIes) (*C.E2nodeComponentConfigRemoval_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccaaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigRemoval_ItemIEs__value_u
	e2nccaaItemC, err := newE2nodeComponentConfigRemovalItem(e2nccaaItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentInterfaceType))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2nccaaItemC.e2nodeComponentID))
	//binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2nccaaItemC.e2nodeComponentConfigurationAck))
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration
	//binary.LittleEndian.PutUint64(choiceC[40:], uint64(e2nccaaItemC.e2nodeComponentConfiguration.present))
	//binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentConfiguration.choice))))

	e2nccaaItemIesC := C.E2nodeComponentConfigRemoval_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeComponentConfigRemoval_ItemIEs__value{
			present: C.E2nodeComponentConfigRemoval_ItemIEs__value_PR_E2nodeComponentConfigRemoval_Item,
			choice:  choiceC,
		},
	}

	return &e2nccaaItemIesC, nil
}

func newE2nodeTnlAssociationRemovalItemIEs(e2nccaaItemIes *e2appducontents.E2NodeTnlassociationRemovalItemIes) (*C.E2nodeTNLassociationRemoval_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2nccaaItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem)
	if err != nil {
		return nil, err
	}

	choiceC := [80]byte{} // The size of the E2nodeComponentConfigRemoval_ItemIEs__value_u
	e2nccaaItemC, err := newE2nodeTNLassociationRemovalItem(e2nccaaItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.tnlInformation))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2nccaaItemC.tnlInformationRIC))
	//binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2nccaaItemC.e2nodeComponentConfigurationAck))
	// Gap of 24 for the asn_struct_ctx_t belonging to E2nodeComponentConfiguration
	//binary.LittleEndian.PutUint64(choiceC[40:], uint64(e2nccaaItemC.e2nodeComponentConfiguration.present))
	//binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2nccaaItemC.e2nodeComponentConfiguration.choice))))

	e2nccaaItemIesC := C.E2nodeTNLassociationRemoval_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2nodeTNLassociationRemoval_ItemIEs__value{
			present: C.E2nodeTNLassociationRemoval_ItemIEs__value_PR_E2nodeTNLassociationgRemoval_Item,
			choice:  choiceC,
		},
	}

	return &e2nccaaItemIesC, nil
}

func newE2connectionSetupFailedItemIEs(e2csfItemIes *e2appducontents.E2ConnectionSetupFailedItemIes) (*C.E2connectionSetupFailed_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(e2csfItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDE2connectionSetupFailedItem)
	if err != nil {
		return nil, err
	}

	choiceC := [144]byte{} // The size of the E2nodeComponentConfigUpdate_ItemIEs__value_u
	e2csfItemC, err := newE2connectionSetupFailedItem(e2csfItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	//ToDo - verify correctness of passing bytes there..
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(e2csfItemC.tnlInformation.tnlAddress.buf))))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(e2csfItemC.tnlInformation.tnlAddress.size))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(e2csfItemC.tnlInformation.tnlAddress.bits_unused))
	// Gap of 24 for the asn_struct_ctx_t belonging to BIT STRING
	binary.LittleEndian.PutUint64(choiceC[48:], uint64(uintptr(unsafe.Pointer(e2csfItemC.tnlInformation.tnlPort))))
	// Gap of 24 for the asn_struct_ctx_t belonging to BIT STRING
	binary.LittleEndian.PutUint64(choiceC[80:], uint64(e2csfItemC.cause.present))
	copy(choiceC[88:96], e2csfItemC.cause.choice[:8])

	e2csfItemIesC := C.E2connectionSetupFailed_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_E2connectionSetupFailed_ItemIEs__value{
			present: C.E2connectionSetupFailed_ItemIEs__value_PR_E2connectionSetupFailed_Item,
			choice:  choiceC,
		},
	}

	return &e2csfItemIesC, nil
}

func newRANfunctionIDCauseItemIEs(rfIDItemIes *e2appducontents.RanfunctionIdcauseItemIes) (*C.RANfunctionIDcause_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(rfIDItemIes.GetRanFunctionIdcauseItemIes7().GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRanfunctionIeCauseItem)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicactionAdmittedItem)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicactionNotAdmittedItem)
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
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRicactionToBeSetupItem)
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

func newRicSubscriptionWithCauseItemIEs(ratbsItemIes *e2appducontents.RicsubscriptionWithCauseItemIes) (*C.RICsubscription_withCause_ItemIEs_t, error) {
	critC, err := criticalityToC(e2ap_commondatatypes.Criticality(ratbsItemIes.GetCriticality()))
	if err != nil {
		return nil, err
	}
	idC, err := protocolIeIDToC(v2.ProtocolIeIDRICsubscriptionWithCauseItem)
	if err != nil {
		return nil, err
	}

	choiceC := [56]byte{} // The size of the RANfunction_ItemIEs__value_u
	ratbsItemC, err := newRicSubscriptionWithCauseItem(ratbsItemIes.GetValue())
	if err != nil {
		return nil, err
	}
	binary.LittleEndian.PutUint64(choiceC[0:], uint64(ratbsItemC.ricRequestID))
	binary.LittleEndian.PutUint64(choiceC[8:], uint64(ratbsItemC.ranFunctionID))
	binary.LittleEndian.PutUint64(choiceC[16:], uint64(uintptr(unsafe.Pointer(ratbsItemC.cause))))
	//binary.LittleEndian.PutUint64(choiceC[24:], uint64(uintptr(unsafe.Pointer(ratbsItemC.ricSubsequentAction))))

	rfItemIesC := C.RICsubscription_withCause_ItemIEs_t{
		id:          idC,
		criticality: critC,
		value: C.struct_RICsubscription_withCause_ItemIEs__value{
			present: C.RICsubscription_withCause_ItemIEs__value_PR_RICsubscription_withCause_Item,
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
		gE2nID, err := decodeGlobalE2NodeIDBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes3 = &e2appducontents.E2SetupRequestIes_E2SetupRequestIes3{
			Id:          int32(v2.ProtocolIeIDGlobalE2nodeID),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionsAdded),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2setupRequestIEs__value_PR_E2nodeComponentConfigUpdate_List:
		rfl, err := decodeE2nodeComponentConfigUpdateListBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes33 = &e2appducontents.E2SetupRequestIes_E2SetupRequestIes33{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdate),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2setupRequestIEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2SetupRequestIes_E2SetupRequestIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2srIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
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
			Id:          int32(v2.ProtocolIeIDGlobalRicID),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionsRejected),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfRejected,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	case C.E2setupResponseIEs__value_PR_E2nodeComponentConfigUpdateAck_List:
		e2nccual, err := decodeE2nodeComponentConfigUpdateAckListBytes(e2srIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes35 = &e2appducontents.E2SetupResponseIes_E2SetupResponseIes35{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       e2nccual,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}
	case C.E2setupResponseIEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2SetupResponseIes_E2SetupResponseIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2srIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
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
			Id:          int32(v2.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsrIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionRequest_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes5{
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
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
			Id:          int32(v2.ProtocolIeIDRicsubscriptionDetails),
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

func decodeRicSubscriptionDeleteRequiredIE(rsrIeC *C.RICsubscriptionDeleteRequired_IEs_t) (*e2appducontents.RicsubscriptionDeleteRequiredIes, error) {
	//	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsrIeC)
	ret := new(e2appducontents.RicsubscriptionDeleteRequiredIes)
	//
	switch rsrIeC.value.present {
	case C.RICsubscriptionDeleteRequired_IEs__value_PR_RICsubscription_List_withCause:
		list, err := decodeRicSubscriptionListWithCauseBytes(rsrIeC.value.choice[:16])
		if err != nil {
			return nil, err
		}
		ret = &e2appducontents.RicsubscriptionDeleteRequiredIes{
			Id:          int32(v2.ProtocolIeIDRICsubscriptionToBeRemoved),
			Value:       list,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionDeleteRequired_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteRequiredIE(). %v not yet implemneted", rsrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionDeleteRequiredIE(). unexpected choice %v", rsrIeC.value.present)
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
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionResponse_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
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
			Id:          int32(v2.ProtocolIeIDRicactionsAdmitted),
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
			Id:          int32(v2.ProtocolIeIDRicactionsNotAdmitted),
			Value:       ranal,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICsubscriptionResponse_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionResponseIE(). Unexpected value.\n%v", rsrIeC.value.present)

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
				Id:          int32(v2.ProtocolIeIDRanfunctionItem),
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
				Id:          int32(v2.ProtocolIeIDRanfunctionIDItem),
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

func decodeE2nodeComponentConfigUpdateItemIes(e2nccuiIesValC *C.struct_E2nodeComponentConfigUpdate_ItemIEs__value) (*e2appducontents.E2NodeComponentConfigUpdateItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccuiIesValC.present; present {
	case C.E2nodeComponentConfigUpdate_ItemIEs__value_PR_E2nodeComponentConfigUpdate_Item:

		e2nccuiIes := e2appducontents.E2NodeComponentConfigUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		//ToDo - verify whether it's decoded correctly
		e2ccui, err := decodeE2nodeComponentConfigUpdateItemBytes(e2nccuiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateItemBytes() %s", err.Error())
		}
		e2nccuiIes.Value = e2ccui
		return &e2nccuiIes, nil
	default:
		return nil, fmt.Errorf("error decoding E2NodeComponentConfigUpdateItemIE - present %v not supported", present)
	}
}

func decodeE2nodeComponentConfigUpdateAckItemIes(e2nccuaiIesValC *C.struct_E2nodeComponentConfigUpdateAck_ItemIEs__value) (*e2appducontents.E2NodeComponentConfigUpdateAckItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccuaiIesValC.present; present {
	case C.E2nodeComponentConfigUpdateAck_ItemIEs__value_PR_E2nodeComponentConfigUpdateAck_Item:

		e2nccuaiIes := e2appducontents.E2NodeComponentConfigUpdateAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		//ToDo - verify whether it's decoded correctly
		e2ccui, err := decodeE2nodeComponentConfigUpdateAckItemBytes(e2nccuaiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigUpdateAckItemBytes() %s", err.Error())
		}
		e2nccuaiIes.Value = e2ccui
		return &e2nccuaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding E2NodeComponentConfigUpdateItemIE - present %v not supported", present)
	}
}

func decodeE2connectionUpdateItemIes(e2cuiIesValC *C.struct_E2connectionUpdate_ItemIEs__value) (*e2appducontents.E2ConnectionUpdateItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2cuiIesValC.present; present {
	case C.E2nodeComponentConfigUpdate_ItemIEs__value_PR_E2nodeComponentConfigUpdate_Item:

		e2cuiIes := e2appducontents.E2ConnectionUpdateItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		//ToDo - verify whether it's decoded correctly
		e2ccui, err := decodeE2connectionUpdateItemBytes(e2cuiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2connectionUpdateItemBytes() %s", err.Error())
		}
		e2cuiIes.Value = e2ccui
		return &e2cuiIes, nil
	default:
		return nil, fmt.Errorf("error decoding E2NodeComponentConfigUpdateItemIE - present %v not supported", present)
	}
}

func decodeE2connectionUpdateRemoveItemIes(e2cuiIesValC *C.struct_E2connectionUpdateRemove_ItemIEs__value) (*e2appducontents.E2ConnectionUpdateRemoveItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2cuiIesValC.present; present {
	case C.E2connectionUpdateRemove_ItemIEs__value_PR_E2connectionUpdateRemove_Item:

		e2curiIes := e2appducontents.E2ConnectionUpdateRemoveItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemoveItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		e2curi, err := decodeE2connectionUpdateRemoveItemBytes(e2cuiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2connectionUpdateRemoveItemBytes() %s", err.Error())
		}
		e2curiIes.Value = e2curi
		return &e2curiIes, nil
	default:
		return nil, fmt.Errorf("error decoding E2connectionUpdateRemoveItemIE - present %v not supported", present)
	}
}

func decodeE2nodeComponentConfigAdditionItemIes(e2nccaiC *C.struct_E2nodeComponentConfigAddition_ItemIEs__value) (*e2appducontents.E2NodeComponentConfigAdditionItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccaiC.present; present {
	case C.E2nodeComponentConfigAddtion_ItemIEs__value_PR_E2nodeComponentConfigAddition_Item:

		e2nccaiIes := e2appducontents.E2NodeComponentConfigAdditionItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		e2nccai, err := decodeE2nodeComponentConfigAdditionItemBytes(e2nccaiC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigAdditionItemBytes() %s", err.Error())
		}
		e2nccaiIes.Value = e2nccai
		return &e2nccaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding decodeE2nodeComponentConfigAddtionItemIE - present %v not supported", present)
	}
}

func decodeE2nodeComponentConfigAdditionAckItemIes(e2nccaaiC *C.struct_E2nodeComponentConfigAdditionAck_ItemIEs__value) (*e2appducontents.E2NodeComponentConfigAdditionAckItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccaaiC.present; present {
	case C.E2nodeComponentConfigAdditionAck_ItemIEs__value_PR_E2nodeComponentConfigAdditionAck_Item:

		e2nccaaiIes := e2appducontents.E2NodeComponentConfigAdditionAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		e2nccaai, err := decodeE2nodeComponentConfigAdditionAckItemBytes(e2nccaaiC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigAdditionAckItemBytes() %s", err.Error())
		}
		e2nccaaiIes.Value = e2nccaai
		return &e2nccaaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding decodeE2nodeComponentConfigAddtionAckItemIE - present %v not supported", present)
	}
}

func decodeE2nodeComponentConfigRemovalAckItemIes(e2nccraiC *C.struct_E2nodeComponentConfigRemovalAck_ItemIEs__value) (*e2appducontents.E2NodeComponentConfigRemovalAckItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccraiC.present; present {
	case C.E2nodeComponentConfigRemovalAck_ItemIEs__value_PR_E2nodeComponentConfigRemovalAck_Item:

		e2nccraiIes := e2appducontents.E2NodeComponentConfigRemovalAckItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalAckItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		e2nccrai, err := decodeE2nodeComponentConfigRemovalAckItemBytes(e2nccraiC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigRemovalAckItemBytes() %s", err.Error())
		}
		e2nccraiIes.Value = e2nccrai
		return &e2nccraiIes, nil
	default:
		return nil, fmt.Errorf("error decoding decodeE2nodeComponentConfigRemovalAckItemIE - present %v not supported", present)
	}
}

func decodeE2nodeComponentConfigRemovalItemIes(e2nccaaiC *C.struct_E2nodeComponentConfigRemoval_ItemIEs__value) (*e2appducontents.E2NodeComponentConfigRemovalItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccaaiC.present; present {
	case C.E2nodeComponentConfigRemoval_ItemIEs__value_PR_E2nodeComponentConfigRemoval_Item:

		e2nccaaiIes := e2appducontents.E2NodeComponentConfigRemovalItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		e2nccaai, err := decodeE2nodeComponentConfigRemovalItemBytes(e2nccaaiC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeComponentConfigRemovalItemBytes() %s", err.Error())
		}
		e2nccaaiIes.Value = e2nccaai
		return &e2nccaaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding decodeE2nodeComponentConfigRemovalItemIE - present %v not supported", present)
	}
}

func decodeE2nodeTnlAssociationRemovalItemIes(e2nccaaiC *C.struct_E2nodeTNLassociationRemoval_ItemIEs__value) (*e2appducontents.E2NodeTnlassociationRemovalItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2nccaaiC.present; present {
	case C.E2nodeTNLassociationRemoval_ItemIEs__value_PR_E2nodeTNLassociationRemoval_Item:

		e2nccaaiIes := e2appducontents.E2NodeTnlassociationRemovalItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeTNLassociationRemovalItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		//ToDo - verify whether it's decoded correctly
		e2nccaai, err := decodeE2nodeTNLassociationRemovalItemBytes(e2nccaaiC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2nodeTNLassociationRemovalItemBytes() %s", err.Error())
		}
		e2nccaaiIes.Value = e2nccaai
		return &e2nccaaiIes, nil
	default:
		return nil, fmt.Errorf("error decoding decodeE2nodeTNLassociationRemovalItemIE - present %v not supported", present)
	}
}

func decodeE2connectionSetupFailedItemIes(e2cuiIesValC *C.struct_E2connectionSetupFailed_ItemIEs__value) (*e2appducontents.E2ConnectionSetupFailedItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDiIesValC, rfIDiIesValC)

	switch present := e2cuiIesValC.present; present {
	case C.E2connectionSetupFailed_ItemIEs__value_PR_E2connectionSetupFailed_Item:

		e2csfiIes := e2appducontents.E2ConnectionSetupFailedItemIes{
			Id:          int32(v2.ProtocolIeIDE2connectionSetupFailedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		//ToDo - verify whether it's decoded correctly
		e2ccsfi, err := decodeE2connectionSetupFailedItemBytes(e2cuiIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeE2connectionSetupFailedItemBytes() %s", err.Error())
		}
		e2csfiIes.Value = e2ccsfi
		return &e2csfiIes, nil
	default:
		return nil, fmt.Errorf("error decoding E2connectionSetupFailedItemIE - present %v not supported", present)
	}
}

func decodeRANfunctionIDCauseItemIes(rfIDciIesValC *C.struct_RANfunctionIDcause_ItemIEs__value) (*e2appducontents.RanfunctionIdcauseItemIes, error) {
	//fmt.Printf("Value %T %v\n", rfIDciIesValC, rfIDciIesValC)

	switch present := rfIDciIesValC.present; present {
	case C.RANfunctionIDcause_ItemIEs__value_PR_RANfunctionIDcause_Item:

		rfIDiIes := e2appducontents.RanfunctionIdcauseItemIes{
			RanFunctionIdcauseItemIes7: &e2appducontents.RanfunctionIdcauseItemIes_RanfunctionIdcauseItemIes7{
				Id:          int32(v2.ProtocolIeIDRanfunctionIeCauseItem),
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
			Id:          int32(v2.ProtocolIeIDRicactionAdmittedItem),
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
			Id:          int32(v2.ProtocolIeIDRicactionNotAdmittedItem),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICactionID:
		raID := decodeRicActionIDBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes15 = &e2appducontents.RicindicationIes_RicindicationIes15{
			Id:          int32(v2.ProtocolIeIDRicactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       raID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICcallProcessID:
		rcpID := decodeRicCallProcessIDBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes20 = &e2appducontents.RicindicationIes_RicindicationIes20{
			Id:          int32(v2.ProtocolIeIDRiccallProcessID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcpID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICindication_IEs__value_PR_RICindicationHeader:
		rih := decodeRicIndicationHeaderBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes25 = &e2appducontents.RicindicationIes_RicindicationIes25{
			Id:          int32(v2.ProtocolIeIDRicindicationHeader),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rih,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICindicationMessage:
		rim := decodeRicIndicationMessageBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes26 = &e2appducontents.RicindicationIes_RicindicationIes26{
			Id:          int32(v2.ProtocolIeIDRicindicationMessage),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rim,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICindicationSN:
		risn := decodeRicIndicationSnBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes27 = &e2appducontents.RicindicationIes_RicindicationIes27{
			Id:          int32(v2.ProtocolIeIDRicindicationSn),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       risn,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICindication_IEs__value_PR_RICindicationType:
		rit := decodeRicIndicationTypeBytes(riIeC.value.choice[0:8])
		ret.E2ApProtocolIes28 = &e2appducontents.RicindicationIes_RicindicationIes28{
			Id:          int32(v2.ProtocolIeIDRicindicationType),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rit,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICindication_IEs__value_PR_RICrequestID:
		rrID := decodeRicRequestIDBytes(riIeC.value.choice[0:16])
		ret.E2ApProtocolIes29 = &e2appducontents.RicindicationIes_RicindicationIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
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

func decodeRicControlRequestIE(rcRIeC *C.RICcontrolRequest_IEs_t) (*e2appducontents.RiccontrolRequestIes, error) {
	//fmt.Printf("Handling E2SetupReqIE %+v\n", riIeC)
	ret := new(e2appducontents.RiccontrolRequestIes)

	switch rcRIeC.value.present {
	case C.RICcontrolRequest_IEs__value_PR_RANfunctionID:
		rfID := decodeRanFunctionIDBytes(rcRIeC.value.choice[0:8])
		ret.E2ApProtocolIes5 = &e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes5{
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolRequest_IEs__value_PR_RICcallProcessID:
		rcpID := decodeRicCallProcessIDBytes(rcRIeC.value.choice[0:16])
		ret.E2ApProtocolIes20 = &e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes20{
			Id:          int32(v2.ProtocolIeIDRiccallProcessID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcpID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICcontrolRequest_IEs__value_PR_RICrequestID:
		rrID := decodeRicRequestIDBytes(rcRIeC.value.choice[0:16])
		ret.E2ApProtocolIes29 = &e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rrID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolRequest_IEs__value_PR_RICcontrolHeader:
		rch := decodeRicControlHeaderBytes(rcRIeC.value.choice[0:16])
		ret.E2ApProtocolIes22 = &e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes22{
			Id:          int32(v2.ProtocolIeIDRiccontrolHeader),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rch,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolRequest_IEs__value_PR_RICcontrolMessage:
		rcm := decodeRicControlMessageBytes(rcRIeC.value.choice[0:16])
		ret.E2ApProtocolIes23 = &e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes23{
			Id:          int32(v2.ProtocolIeIDRiccontrolMessage),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcm,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolRequest_IEs__value_PR_RICcontrolAckRequest:
		rcar := decodeRicControlAckRequestBytes(rcRIeC.value.choice[0:16])
		ret.E2ApProtocolIes21 = &e2appducontents.RiccontrolRequestIes_RiccontrolRequestIes21{
			Id:          int32(v2.ProtocolIeIDRiccontrolAckRequest),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcar,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICcontrolRequest_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicControlRequestIE(). %v not yet implemneted", rcRIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicControlRequestIE(). unexpected choice %v", rcRIeC.value.present)
	}

	return ret, nil
}

func decodeRicControlFailureIE(rcfIeC *C.RICcontrolFailure_IEs_t) (*e2appducontents.RiccontrolFailureIes, error) {
	//fmt.Printf("Handling RicControlFailureIE %+v\n", riIeC)
	ret := new(e2appducontents.RiccontrolFailureIes)

	switch rcfIeC.value.present {
	case C.RICcontrolFailure_IEs__value_PR_RANfunctionID:
		rfID := decodeRanFunctionIDBytes(rcfIeC.value.choice[0:8])
		ret.E2ApProtocolIes5 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes5{
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolFailure_IEs__value_PR_RICcallProcessID:
		rcpID := decodeRicCallProcessIDBytes(rcfIeC.value.choice[0:16])
		ret.E2ApProtocolIes20 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes20{
			Id:          int32(v2.ProtocolIeIDRiccallProcessID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcpID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICcontrolFailure_IEs__value_PR_RICrequestID:
		rrID := decodeRicRequestIDBytes(rcfIeC.value.choice[0:16])
		ret.E2ApProtocolIes29 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rrID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolFailure_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(rcfIeC.value.choice[0:16])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes1 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cause,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolFailure_IEs__value_PR_RICcontrolOutcome:
		rco := decodeRicControlOutcomeBytes(rcfIeC.value.choice[0:16])
		ret.E2ApProtocolIes32 = &e2appducontents.RiccontrolFailureIes_RiccontrolFailureIes32{
			Id:          int32(v2.ProtocolIeIDRiccontrolOutcome),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rco,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICcontrolFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicControlFailureIE(). %v not yet implemneted", rcfIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicControlFailureIE(). unexpected choice %v", rcfIeC.value.present)
	}

	return ret, nil
}

func decodeRicControlAcknowledgeIE(rcaIeC *C.RICcontrolAcknowledge_IEs_t) (*e2appducontents.RiccontrolAcknowledgeIes, error) {
	//fmt.Printf("Handling E2SetupReqIE %+v\n", riIeC)
	ret := new(e2appducontents.RiccontrolAcknowledgeIes)

	switch rcaIeC.value.present {
	case C.RICcontrolAcknowledge_IEs__value_PR_RANfunctionID:
		rfID := decodeRanFunctionIDBytes(rcaIeC.value.choice[0:8])
		ret.E2ApProtocolIes5 = &e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes5{
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolAcknowledge_IEs__value_PR_RICcallProcessID:
		rcpID := decodeRicCallProcessIDBytes(rcaIeC.value.choice[0:16])
		ret.E2ApProtocolIes20 = &e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes20{
			Id:          int32(v2.ProtocolIeIDRiccallProcessID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcpID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICcontrolAcknowledge_IEs__value_PR_RICcontrolStatus:
		rcs := decodeRicControlStatusBytes(rcaIeC.value.choice[0:16])
		ret.E2ApProtocolIes24 = &e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes24{
			Id:          int32(v2.ProtocolIeIDRiccontrolStatus),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rcs,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolAcknowledge_IEs__value_PR_RICrequestID:
		rrID := decodeRicRequestIDBytes(rcaIeC.value.choice[0:16])
		ret.E2ApProtocolIes29 = &e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rrID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICcontrolAcknowledge_IEs__value_PR_RICcontrolOutcome:
		rco := decodeRicControlOutcomeBytes(rcaIeC.value.choice[0:16])
		ret.E2ApProtocolIes32 = &e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes32{
			Id:          int32(v2.ProtocolIeIDRiccontrolOutcome),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rco,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICcontrolAcknowledge_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicControlAcknowledgeIE(). %v not yet implemneted", rcaIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicControlAcknowledgeIE(). unexpected choice %v", rcaIeC.value.present)
	}

	return ret, nil
}

func decodeRicActionToBeSetupItemIes(ratbsIesValC *C.struct_RICaction_ToBeSetup_ItemIEs__value) (*e2appducontents.RicactionToBeSetupItemIes, error) {
	//fmt.Printf("Value %T %v\n", ratbsIesValC, ratbsIesValC)

	switch present := ratbsIesValC.present; present {
	case C.RICaction_ToBeSetup_ItemIEs__value_PR_RICaction_ToBeSetup_Item:
		ratbsIIes := e2appducontents.RicactionToBeSetupItemIes{
			Id:          int32(v2.ProtocolIeIDRicactionToBeSetupItem),
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

func decodeRicSubscriptionWithCauseItemIes(ratbsIesValC *C.struct_RICsubscription_withCause_ItemIEs__value) (*e2appducontents.RicsubscriptionWithCauseItemIes, error) {
	//fmt.Printf("Value %T %v\n", ratbsIesValC, ratbsIesValC)

	switch present := ratbsIesValC.present; present {
	case C.RICsubscription_withCause_ItemIEs__value_PR_RICsubscription_withCause_Item:
		ratbsIIes := e2appducontents.RicsubscriptionWithCauseItemIes{
			Id:          int32(v2.ProtocolIeIDRICsubscriptionWithCauseItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		ratbsI, err := decodeRicSubscriptionWithCauseItemBytes(ratbsIesValC.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeRicSubscriptionWithCauseItemBytes() %s", err.Error())
		}
		ratbsIIes.Value = ratbsI
		return &ratbsIIes, nil
	default:
		return nil, fmt.Errorf("error decoding RicsubscriptionWithCauseItemIes - present %v not supported", present)
	}
}

func decodeRicSubscriptionDeleteRequestIE(rsrdIeC *C.RICsubscriptionDeleteRequest_IEs_t) (*e2appducontents.RicsubscriptionDeleteRequestIes, error) {
	//	//fmt.Printf("Handling RicSubscriptionResp %+v\n", rsrdIeC)
	ret := new(e2appducontents.RicsubscriptionDeleteRequestIes)
	//
	switch rsrdIeC.value.present {
	case C.RICsubscriptionDeleteRequest_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(rsrdIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
	case C.RICsubscriptionDeleteRequest_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.RicsubscriptionDeleteRequestIes_RicsubscriptionDeleteRequestIes5{
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionDeleteResponse_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionDeleteResponseIes_RicsubscriptionDeleteResponseIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionDeleteFailure_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionDeleteFailureIes_RicsubscriptionDeleteFailureIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
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
			Id:          int32(v2.ProtocolIeIDCause),
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
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
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
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
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
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Value:       cd,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(rsfIeC.value.choice[:16])
		if err != nil {
			return nil, fmt.Errorf("decodeCauseBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes1 = &e2appducontents.RicsubscriptionFailureIes_RicsubscriptionFailureIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Value:       cause,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICsubscriptionFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicSubscriptionFailureIE(). %v not yet implemneted", rsfIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicSubscriptionFailureIE(). unexpected choice %v", rsfIeC.value.present)
	}

	return ret, nil
}

func decodeErrorIndicationIE(eiIeC *C.ErrorIndication_IEs_t) (*e2appducontents.ErrorIndicationIes, error) {
	//fmt.Printf("Handling ErrorIndication %+v\n", rsfIeC)
	ret := new(e2appducontents.ErrorIndicationIes)

	switch eiIeC.value.present {
	case C.ErrorIndication_IEs__value_PR_RANfunctionID:
		ret.E2ApProtocolIes5 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes5{
			Value:       decodeRanFunctionIDBytes(eiIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			Id:          int32(v2.ProtocolIeIDRanfunctionID),
		}

	case C.ErrorIndication_IEs__value_PR_RICrequestID:
		ret.E2ApProtocolIes29 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes29{
			Id:          int32(v2.ProtocolIeIDRicrequestID),
			Value:       decodeRicRequestIDBytes(eiIeC.value.choice[:16]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.ErrorIndication_IEs__value_PR_CriticalityDiagnostics:
		cd, err := decodeCriticalityDiagnosticsBytes(eiIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeCriticalityDiagnosticsBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes2 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes2{
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Value:       cd,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.ErrorIndication_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(eiIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeCauseBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes1 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Value:       cause,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.ErrorIndication_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.ErrorIndicationIes_ErrorIndicationIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Value:       decodeTransactionIDBytes(eiIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.ErrorIndication_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeErrorIndicationIE(). %v not yet implemneted", eiIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeErrorIndicationIE(). unexpected choice %v", eiIeC.value.present)
	}

	return ret, nil
}

func decodeRicServiceQueryIE(rsqIeC *C.RICserviceQuery_IEs_t) (*e2appducontents.RicserviceQueryIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.RicserviceQueryIes)

	switch rsqIeC.value.present {
	case C.RICserviceQuery_IEs__value_PR_RANfunctionsID_List:
		a := [112]byte{}
		copy(a[0:48], rsqIeC.value.choice[:])
		rfAccepted, err := decodeRanFunctionsIDListBytes(a)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes9 = &e2appducontents.RicserviceQueryIes_RicserviceQueryIes9{
			Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfAccepted,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICserviceQuery_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.RicserviceQueryIes_RicserviceQueryIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(rsqIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICserviceQuery_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicServiceQueryIE(). %v not yet implemneted", rsqIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicServiceQueryIE(). unexpected choice %v", rsqIeC.value.present)
	}

	return ret, nil
}

func decodeResetRequestIE(rrIeC *C.ResetRequestIEs_t) (*e2appducontents.ResetRequestIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.ResetRequestIes)

	switch rrIeC.value.present {
	case C.ResetRequestIEs__value_PR_Cause:
		cause, err := decodeCauseBytes(rrIeC.value.choice[:16])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes1 = &e2appducontents.ResetRequestIes_ResetRequestIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cause,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.ResetRequestIEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.ResetRequestIes_ResetRequestIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(rrIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.ResetRequestIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeResetRequestIE(). %v not yet implemneted", rrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeResetRequestIE(). unexpected choice %v", rrIeC.value.present)
	}

	return ret, nil
}

func decodeResetResponseIE(rrIeC *C.ResetResponseIEs_t) (*e2appducontents.ResetResponseIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.ResetResponseIes)

	switch rrIeC.value.present {
	case C.ResetResponseIEs__value_PR_CriticalityDiagnostics:
		//a := []byte{}
		//copy(a, rrIeC.value.choice[:])
		cd, err := decodeCriticalityDiagnosticsBytes(rrIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes2 = &e2appducontents.ResetResponseIes_ResetResponseIes2{
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cd,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.ResetResponseIEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.ResetResponseIes_ResetResponseIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(rrIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.ResetResponseIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeResetResponseIE(). %v not yet implemneted", rrIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeResetResponseIE(). unexpected choice %v", rrIeC.value.present)
	}

	return ret, nil
}

func decodeRicServiceUpdateIE(rsuIeC *C.RICserviceUpdate_IEs_t) (*e2appducontents.RicserviceUpdateIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.RicserviceUpdateIes)

	switch rsuIeC.value.present {
	case C.RICserviceUpdate_IEs__value_PR_RANfunctions_List: //This one is for added
		rfl, err := decodeRanFunctionsListBytes(rsuIeC.value.choice)
		if err != nil {
			return nil, err
		}

		id := decodeProtocolIeID(rsuIeC.id)
		if id.GetValue() == int32(v2.ProtocolIeIDRanfunctionsAdded) {
			ret.E2ApProtocolIes10 = &e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes10{
				Id:          int32(v2.ProtocolIeIDRanfunctionsAdded),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       rfl,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			}
		}

		if id.GetValue() == int32(v2.ProtocolIeIDRanfunctionsModified) {
			ret.E2ApProtocolIes12 = &e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes12{
				Id:          int32(v2.ProtocolIeIDRanfunctionsModified),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       rfl,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			}
		}

	case C.RICserviceUpdate_IEs__value_PR_RANfunctionsID_List:
		a := [112]byte{}
		copy(a[0:48], rsuIeC.value.choice[:])
		rfdl, err := decodeRanFunctionsIDListBytes(a)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes11 = &e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes11{
			Id:          int32(v2.ProtocolIeIDRanfunctionsDeleted),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfdl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICserviceUpdate_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(rsuIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICserviceUpdate_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicServiceUpdateIE(). %v not yet implemneted", rsuIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicServiceUpdateIE(). unexpected choice %v", rsuIeC.value.present)
	}

	return ret, nil
}

func decodeRicServiceUpdateAcknowledgeIE(rsuaIeC *C.RICserviceUpdateAcknowledge_IEs_t) (*e2appducontents.RicserviceUpdateAcknowledgeIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.RicserviceUpdateAcknowledgeIes)

	switch rsuaIeC.value.present {
	case C.RICserviceUpdateAcknowledge_IEs__value_PR_RANfunctionsID_List: //This one is for added
		a := [112]byte{}
		copy(a[0:48], rsuaIeC.value.choice[:])
		rfIDl, err := decodeRanFunctionsIDListBytes(a)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes9 = &e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes9{
			Id:          int32(v2.ProtocolIeIDRanfunctionsAccepted),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfIDl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICserviceUpdateAcknowledge_IEs__value_PR_RANfunctionsIDcause_List:
		a := [112]byte{}
		copy(a[0:48], rsuaIeC.value.choice[:])
		rfIDcl, err := decodeRanFunctionsIDCauseListBytes(a)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes13 = &e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes13{
			Id:          int32(v2.ProtocolIeIDRanfunctionsRejected),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       rfIDcl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICserviceUpdateAcknowledge_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.RicserviceUpdateAcknowledgeIes_RicserviceUpdateAcknowledgeIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(rsuaIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICserviceUpdateAcknowledge_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicServiceUpdateAcknowledgeIE(). %v not yet implemneted", rsuaIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicServiceUpdateAcknowledgeIE(). unexpected choice %v", rsuaIeC.value.present)
	}

	return ret, nil
}

func decodeRicServiceUpdateFailureIE(rsufIeC *C.RICserviceUpdateFailure_IEs_t) (*e2appducontents.RicserviceUpdateFailureIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.RicserviceUpdateFailureIes)

	switch rsufIeC.value.present {
	case C.RICserviceUpdateFailure_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(rsufIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}

		ret.E2ApProtocolIes1 = &e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       cause,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICserviceUpdateFailure_IEs__value_PR_CriticalityDiagnostics:
		cd, err := decodeCriticalityDiagnosticsBytes(rsufIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes2 = &e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes2{
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cd,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICserviceUpdateFailure_IEs__value_PR_TimeToWait:
		ttw := decodeTimeToWaitBytes(rsufIeC.value.choice[:8])

		ret.E2ApProtocolIes31 = &e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes31{
			Id:          int32(v2.ProtocolIeIDTimeToWait),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       ttw,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.RICserviceUpdateFailure_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.RicserviceUpdateFailureIes_RicserviceUpdateFailureIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(rsufIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.RICserviceUpdateFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeRicServiceUpdateFailureIE(). %v not yet implemneted", rsufIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeRicServiceUpdateFailureIE(). unexpected choice %v", rsufIeC.value.present)
	}

	return ret, nil
}

func decodeE2nodeConfigurationUpdateIE(e2ncuIeC *C.E2nodeConfigurationUpdate_IEs_t) (*e2appducontents.E2NodeConfigurationUpdateIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.E2NodeConfigurationUpdateIes)

	switch e2ncuIeC.value.present {
	case C.E2nodeConfigurationUpdate_IEs__value_PR_GlobalE2node_ID:
		gE2nID, err := decodeGlobalE2NodeIDBytes(e2ncuIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes3 = &e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes3{
			Id:          int32(v2.ProtocolIeIDGlobalE2nodeID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       gE2nID,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2nodeConfigurationUpdate_IEs__value_PR_E2nodeComponentConfigUpdate_List:
		e2ncul, err := decodeE2nodeComponentConfigUpdateListBytes(e2ncuIeC.value.choice)
		if err != nil {
			return nil, err
		}

		ret.E2ApProtocolIes33 = &e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes33{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdate),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       e2ncul,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2nodeConfigurationUpdate_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2ncuIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2nodeConfigurationUpdate_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateIE(). %v not yet implemneted", e2ncuIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateIE(). unexpected choice %v", e2ncuIeC.value.present)
	}

	return ret, nil
}

func decodeE2nodeConfigurationUpdateAcknowledgeIE(e2ncuaIeC *C.E2nodeConfigurationUpdateAcknowledge_IEs_t) (*e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes)

	switch e2ncuaIeC.value.present {
	case C.E2nodeConfigurationUpdateAcknowledge_IEs__value_PR_E2nodeComponentConfigUpdateAck_List:
		a := [112]byte{}
		copy(a[0:48], e2ncuaIeC.value.choice[:])
		e2ncual, err := decodeE2nodeComponentConfigUpdateAckListBytes(a)
		if err != nil {
			return nil, err
		}

		ret.E2ApProtocolIes35 = &e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes_E2NodeConfigurationUpdateAcknowledgeIes35{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       e2ncual,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2nodeConfigurationUpdateAcknowledge_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes_E2NodeConfigurationUpdateAcknowledgeIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2ncuaIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2nodeConfigurationUpdateAcknowledge_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateAcknowledgeIE(). %v not yet implemneted", e2ncuaIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateAcknowledgeIE(). unexpected choice %v", e2ncuaIeC.value.present)
	}

	return ret, nil
}

func decodeE2nodeConfigurationUpdateFailureIE(e2ncufIeC *C.E2nodeConfigurationUpdateFailure_IEs_t) (*e2appducontents.E2NodeConfigurationUpdateFailureIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.E2NodeConfigurationUpdateFailureIes)

	switch e2ncufIeC.value.present {
	case C.E2nodeConfigurationUpdateFailure_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(e2ncufIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes1 = &e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cause,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2nodeConfigurationUpdateFailure_IEs__value_PR_CriticalityDiagnostics: //This one is for added
		cd, err := decodeCriticalityDiagnosticsBytes(e2ncufIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes2 = &e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes2{
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cd,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2nodeConfigurationUpdateFailure_IEs__value_PR_TimeToWait: //This one is for added
		ttw := decodeTimeToWaitBytes(e2ncufIeC.value.choice[:8])

		ret.E2ApProtocolIes31 = &e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes31{
			Id:          int32(v2.ProtocolIeIDTimeToWait),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       ttw,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2nodeConfigurationUpdateFailure_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2NodeConfigurationUpdateFailureIes_E2NodeConfigurationUpdateFailureIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2ncufIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2nodeConfigurationUpdateFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateFailureIE(). %v not yet implemneted", e2ncufIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2nodeConfigurationUpdateFailureIE(). unexpected choice %v", e2ncufIeC.value.present)
	}

	return ret, nil
}

func decodeE2connectionUpdateFailureIE(e2cufIeC *C.E2connectionUpdateFailure_IEs_t) (*e2appducontents.E2ConnectionUpdateFailureIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.E2ConnectionUpdateFailureIes)

	switch e2cufIeC.value.present {
	case C.E2connectionUpdateFailure_IEs__value_PR_Cause:
		cause, err := decodeCauseBytes(e2cufIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes1 = &e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       cause,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2connectionUpdateFailure_IEs__value_PR_CriticalityDiagnostics: //This one is for added
		cd, err := decodeCriticalityDiagnosticsBytes(e2cufIeC.value.choice[:48])
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes2 = &e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes2{
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       cd,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2connectionUpdateFailure_IEs__value_PR_TimeToWait: //This one is for added
		ttw := decodeTimeToWaitBytes(e2cufIeC.value.choice[:8])

		ret.E2ApProtocolIes31 = &e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes31{
			Id:          int32(v2.ProtocolIeIDTimeToWait),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value:       ttw,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2connectionUpdateFailure_IEs__value_PR_TransactionID: //This one is for added
		ret.E2ApProtocolIes49 = &e2appducontents.E2ConnectionUpdateFailureIes_E2ConnectionUpdateFailureIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2cufIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2connectionUpdateFailure_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2connectionUpdateFailureIE(). %v not yet implemneted", e2cufIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2connectionUpdateFailureIE(). unexpected choice %v", e2cufIeC.value.present)
	}

	return ret, nil
}

func decodeE2connectionUpdateIE(e2cuIeC *C.E2connectionUpdate_IEs_t) (*e2appducontents.E2ConnectionUpdateIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.E2ConnectionUpdateIes)

	switch e2cuIeC.value.present {
	case C.E2connectionUpdate_IEs__value_PR_E2connectionUpdate_List: //This one is for added
		cul, err := decodeE2connectionUpdateListBytes(e2cuIeC.value.choice)
		if err != nil {
			return nil, err
		}

		id := decodeProtocolIeID(e2cuIeC.id)
		if id.GetValue() == int32(v2.ProtocolIeIDE2connectionUpdateAdd) {
			ret.E2ApProtocolIes44 = &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes44{
				Id:          int32(v2.ProtocolIeIDE2connectionUpdateAdd),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       cul,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			}
		}

		if id.GetValue() == int32(v2.ProtocolIeIDE2connectionUpdateModify) {
			ret.E2ApProtocolIes45 = &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes45{
				Id:          int32(v2.ProtocolIeIDE2connectionUpdateModify),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       cul,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			}
		}

	case C.E2connectionUpdate_IEs__value_PR_E2connectionUpdateRemove_List:
		crl, err := decodeE2connectionUpdateRemoveListBytes(e2cuIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes46 = &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes46{
			Id:          int32(v2.ProtocolIeIDE2connectionUpdateRemove),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       crl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2connectionUpdate_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2cuIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2connectionUpdate_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2connectionUpdateIE(). %v not yet implemneted", e2cuIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2connectionUpdateIE(). unexpected choice %v", e2cuIeC.value.present)
	}

	return ret, nil
}

func decodeE2connectionUpdateAckIE(e2cuaIeC *C.E2connectionUpdateAck_IEs_t) (*e2appducontents.E2ConnectionUpdateAckIes, error) {
	//fmt.Printf("Handling RicserviceQuery %+v\n", rsfIeC)
	ret := new(e2appducontents.E2ConnectionUpdateAckIes)

	switch e2cuaIeC.value.present {
	case C.E2connectionUpdateAck_IEs__value_PR_E2connectionUpdate_List: //This one is for added
		//a := [112]byte{}
		//copy(a[0:48], rrIeC.value.choice[:])
		csl, err := decodeE2connectionUpdateListBytes(e2cuaIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes39 = &e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes39{
			Id:          int32(v2.ProtocolIeIDE2connectionSetup),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       csl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2connectionUpdateAck_IEs__value_PR_E2connectionSetupFailed_List:
		//a := [112]byte{}
		//copy(a[0:48], rrIeC.value.choice[:])
		csfl, err := decodeE2connectionSetupFailedListBytes(e2cuaIeC.value.choice)
		if err != nil {
			return nil, err
		}
		ret.E2ApProtocolIes40 = &e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes40{
			Id:          int32(v2.ProtocolIeIDE2connectionSetupFailed),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       csfl,
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2connectionUpdateAck_IEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value:       decodeTransactionIDBytes(e2cuaIeC.value.choice[:8]),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2connectionUpdate_IEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeE2connectionUpdateAckIE(). %v not yet implemneted", e2cuaIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeE2connectionUpdateAckIE(). unexpected choice %v", e2cuaIeC.value.present)
	}

	return ret, nil
}

func decodeE2setupFailureIE(eiIeC *C.E2setupFailureIEs_t) (*e2appducontents.E2SetupFailureIes, error) {
	//fmt.Printf("Handling ErrorIndication %+v\n", rsfIeC)
	ret := new(e2appducontents.E2SetupFailureIes)

	switch eiIeC.value.present {
	case C.E2setupFailureIEs__value_PR_TimeToWait:
		ret.E2ApProtocolIes31 = &e2appducontents.E2SetupFailureIes_E2SetupFailureIes31{
			Value:       decodeTimeToWaitBytes(eiIeC.value.choice[:8]), //TODO: See RICtimeToWait
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			Id:          int32(v2.ProtocolIeIDTimeToWait),
		}

	case C.E2setupFailureIEs__value_PR_CriticalityDiagnostics:
		cd, err := decodeCriticalityDiagnosticsBytes(eiIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeCriticalityDiagnosticsBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes2 = &e2appducontents.E2SetupFailureIes_E2SetupFailureIes2{
			Id:          int32(v2.ProtocolIeIDCriticalityDiagnostics),
			Value:       cd,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2setupFailureIEs__value_PR_Cause:
		cause, err := decodeCauseBytes(eiIeC.value.choice[:48])
		if err != nil {
			return nil, fmt.Errorf("decodeCauseBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes1 = &e2appducontents.E2SetupFailureIes_E2SetupFailureIes1{
			Id:          int32(v2.ProtocolIeIDCause),
			Value:       cause,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2setupFailureIEs__value_PR_TNLinformation:
		tnl, err := decodeTnlinformationBytes(eiIeC.value.choice)
		if err != nil {
			return nil, fmt.Errorf("decodeTnlinformationBytes() %s", err.Error())
		}
		ret.E2ApProtocolIes48 = &e2appducontents.E2SetupFailureIes_E2SetupFailureIes48{
			Id:          int32(v2.ProtocolIeIDTNLinformation),
			Value:       tnl,
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

	case C.E2setupFailureIEs__value_PR_TransactionID:
		ret.E2ApProtocolIes49 = &e2appducontents.E2SetupFailureIes_E2SetupFailureIes49{
			Id:          int32(v2.ProtocolIeIDTransactionID),
			Value:       decodeTransactionIDBytes(eiIeC.value.choice[:8]),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}

	case C.E2setupFailureIEs__value_PR_NOTHING:
		return nil, fmt.Errorf("decodeErrorIndicationIE(). %v not yet implemneted", eiIeC.value.present)

	default:
		return nil, fmt.Errorf("decodeErrorIndicationIE(). unexpected choice %v", eiIeC.value.present)
	}

	return ret, nil
}
