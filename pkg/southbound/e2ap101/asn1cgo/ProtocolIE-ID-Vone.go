// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package asn1cgo

//#include "ProtocolIE-ID-Vone.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
)

func newprotocolIeID(id *e2ap_commondatatypes.ProtocolIeId) (C.ProtocolIE_ID_Vone_t, error) {
	return protocolIeIDToC(v1beta2.ProtocolIeID(id.GetValue()))
}

func protocolIeIDToC(pcIeID v1beta2.ProtocolIeID) (C.ProtocolIE_ID_Vone_t, error) {
	switch pcIeID {
	case v1beta2.ProtocolIeIDCause: // Value from e2ap_constants.proto:86
		return C.ProtocolIE_ID_Vone_id_CauseVone, nil
	case v1beta2.ProtocolIeIDCriticalityDiagnostics:
		return C.ProtocolIE_ID_Vone_id_CriticalityDiagnosticsVone, nil
	case v1beta2.ProtocolIeIDGlobalE2nodeID:
		return C.ProtocolIE_ID_Vone_id_GlobalE2node_IDVone, nil
	case v1beta2.ProtocolIeIDGlobalRicID:
		return C.ProtocolIE_ID_Vone_id_GlobalRIC_IDVone, nil
	case v1beta2.ProtocolIeIDRanfunctionID:
		return C.ProtocolIE_ID_Vone_id_RANfunctionIDVone, nil
	case v1beta2.ProtocolIeIDRanfunctionIDItem:
		return C.ProtocolIE_ID_Vone_id_RANfunctionIDVone_Item, nil
	case v1beta2.ProtocolIeIDRanfunctionIeCauseItem:
		return C.ProtocolIE_ID_Vone_id_RANfunctionIEcause_Item, nil
	case v1beta2.ProtocolIeIDRanfunctionItem:
		return C.ProtocolIE_ID_Vone_id_RANfunction_Item_Vone, nil
	case v1beta2.ProtocolIeIDRanfunctionsAccepted:
		return C.ProtocolIE_ID_Vone_id_RANfunctionsAccepted, nil
	case v1beta2.ProtocolIeIDRanfunctionsAdded:
		return C.ProtocolIE_ID_Vone_id_RANfunctionsAdded, nil
	case v1beta2.ProtocolIeIDRanfunctionsDeleted:
		return C.ProtocolIE_ID_Vone_id_RANfunctionsDeleted, nil
	case v1beta2.ProtocolIeIDRanfunctionsModified:
		return C.ProtocolIE_ID_Vone_id_RANfunctionsModified, nil
	case v1beta2.ProtocolIeIDRanfunctionsRejected:
		return C.ProtocolIE_ID_Vone_id_RANfunctionsRejected, nil
	case v1beta2.ProtocolIeIDRicactionAdmittedItem:
		return C.ProtocolIE_ID_Vone_id_RICaction_Admitted_Item_Vone, nil
	case v1beta2.ProtocolIeIDRicactionID:
		return C.ProtocolIE_ID_Vone_id_RICactionIDVone, nil
	case v1beta2.ProtocolIeIDRicactionNotAdmittedItem:
		return C.ProtocolIE_ID_Vone_id_RICaction_NotAdmitted_Item_Vone, nil
	case v1beta2.ProtocolIeIDRicactionsAdmitted:
		return C.ProtocolIE_ID_Vone_id_RICactions_Admitted, nil
	case v1beta2.ProtocolIeIDRicactionsNotAdmitted:
		return C.ProtocolIE_ID_Vone_id_RICactions_NotAdmitted, nil
	case v1beta2.ProtocolIeIDRicactionToBeSetupItem:
		return C.ProtocolIE_ID_Vone_id_RICaction_ToBeSetup_Item_Vone, nil
	case v1beta2.ProtocolIeIDRiccallProcessID:
		return C.ProtocolIE_ID_Vone_id_RICcallProcessIDVone, nil
	case v1beta2.ProtocolIeIDRiccontrolAckRequest:
		return C.ProtocolIE_ID_Vone_id_RICcontrolAckRequestVone, nil
	case v1beta2.ProtocolIeIDRiccontrolHeader:
		return C.ProtocolIE_ID_Vone_id_RICcontrolHeaderVone, nil
	case v1beta2.ProtocolIeIDRiccontrolMessage:
		return C.ProtocolIE_ID_Vone_id_RICcontrolMessageVone, nil
	case v1beta2.ProtocolIeIDRiccontrolStatus:
		return C.ProtocolIE_ID_Vone_id_RICcontrolStatusVone, nil
	case v1beta2.ProtocolIeIDRicindicationHeader:
		return C.ProtocolIE_ID_Vone_id_RICindicationHeaderVone, nil
	case v1beta2.ProtocolIeIDRicindicationMessage:
		return C.ProtocolIE_ID_Vone_id_RICindicationMessageVone, nil
	case v1beta2.ProtocolIeIDRicindicationSn:
		return C.ProtocolIE_ID_Vone_id_RICindicationSNVone, nil
	case v1beta2.ProtocolIeIDRicindicationType:
		return C.ProtocolIE_ID_Vone_id_RICindicationTypeVone, nil
	case v1beta2.ProtocolIeIDRicrequestID:
		return C.ProtocolIE_ID_Vone_id_RICrequestIDVone, nil
	case v1beta2.ProtocolIeIDRicsubscriptionDetails:
		return C.ProtocolIE_ID_Vone_id_RICsubscriptionDetailsVone, nil
	case v1beta2.ProtocolIeIDTimeToWait:
		return C.ProtocolIE_ID_Vone_id_TimeToWaitVone, nil
	case v1beta2.ProtocolIeIDRiccontrolOutcome:
		return C.ProtocolIE_ID_Vone_id_RICcontrolOutcomeVone, nil
	case v1beta2.ProtocolIeIDE2nodeComponentConfigUpdate:
		return C.ProtocolIE_ID_Vone_id_E2nodeComponentConfigUpdateVone, nil
	case v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateItem:
		return C.ProtocolIE_ID_Vone_id_E2nodeComponentConfigUpdateVone_Item, nil
	case v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAck:
		return C.ProtocolIE_ID_Vone_id_E2nodeComponentConfigUpdateAckVone, nil
	case v1beta2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem:
		return C.ProtocolIE_ID_Vone_id_E2nodeComponentConfigUpdateAckVone_Item, nil
	case v1beta2.ProtocolIeIDE2connectionSetup:
		return C.ProtocolIE_ID_Vone_id_E2connectionSetup, nil
	case v1beta2.ProtocolIeIDE2connectionSetupFailed:
		return C.ProtocolIE_ID_Vone_id_E2connectionSetupFailed, nil
	case v1beta2.ProtocolIeIDE2connectionSetupFailedItem:
		return C.ProtocolIE_ID_Vone_id_E2connectionSetupFailed_Item_Vone, nil
	case v1beta2.ProtocolIeIDE2connectionFailedItem:
		return C.ProtocolIE_ID_Vone_id_E2connectionFailed_Item, nil
	case v1beta2.ProtocolIeIDE2connectionUpdateItem:
		return C.ProtocolIE_ID_Vone_id_E2connectionUpdateVone_Item, nil
	case v1beta2.ProtocolIeIDE2connectionUpdateAdd:
		return C.ProtocolIE_ID_Vone_id_E2connectionUpdateAdd, nil
	case v1beta2.ProtocolIeIDE2connectionUpdateModify:
		return C.ProtocolIE_ID_Vone_id_E2connectionUpdateModify, nil
	case v1beta2.ProtocolIeIDE2connectionUpdateRemove:
		return C.ProtocolIE_ID_Vone_id_E2connectionUpdateRemove, nil
	case v1beta2.ProtocolIeIDE2connectionUpdateRemoveItem:
		return C.ProtocolIE_ID_Vone_id_E2connectionUpdateRemove_Item_Vone, nil
	case v1beta2.ProtocolIeIDTNLinformation:
		return C.ProtocolIE_ID_Vone_id_TNLinformationVone, nil
	default:
		return C.ProtocolIE_ID_Vone_t(-1), fmt.Errorf("unexpected value for ProtocolIE_IDT %v", pcIeID)
	}
}

func decodeProtocolIeID(pcIeIDC C.ProtocolIE_ID_Vone_t) *e2ap_commondatatypes.ProtocolIeId {
	ret := e2ap_commondatatypes.ProtocolIeId{
		Value: int32(pcIeIDC),
	}
	return &ret
}
