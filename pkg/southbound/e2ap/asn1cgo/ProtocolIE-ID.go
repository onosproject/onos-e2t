// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package asn1cgo

//#include "ProtocolIE-ID.h"
import "C"
import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
)

func newprotocolIeID(id *e2ap_commondatatypes.ProtocolIeId) (C.ProtocolIE_ID_t, error) {
	return protocolIeIDToC(v2beta1.ProtocolIeID(id.GetValue()))
}

func protocolIeIDToC(pcIeID v2beta1.ProtocolIeID) (C.ProtocolIE_ID_t, error) {
	switch pcIeID {
	case v2beta1.ProtocolIeIDCause: // Value from e2ap_constants.proto:86
		return C.ProtocolIE_ID_id_Cause, nil
	case v2beta1.ProtocolIeIDCriticalityDiagnostics:
		return C.ProtocolIE_ID_id_CriticalityDiagnostics, nil
	case v2beta1.ProtocolIeIDGlobalE2nodeID:
		return C.ProtocolIE_ID_id_GlobalE2node_ID, nil
	case v2beta1.ProtocolIeIDGlobalRicID:
		return C.ProtocolIE_ID_id_GlobalRIC_ID, nil
	case v2beta1.ProtocolIeIDRanfunctionID:
		return C.ProtocolIE_ID_id_RANfunctionID, nil
	case v2beta1.ProtocolIeIDRanfunctionIDItem:
		return C.ProtocolIE_ID_id_RANfunctionID_Item, nil
	case v2beta1.ProtocolIeIDRanfunctionIeCauseItem:
		return C.ProtocolIE_ID_id_RANfunctionIEcause_Item, nil
	case v2beta1.ProtocolIeIDRanfunctionItem:
		return C.ProtocolIE_ID_id_RANfunction_Item, nil
	case v2beta1.ProtocolIeIDRanfunctionsAccepted:
		return C.ProtocolIE_ID_id_RANfunctionsAccepted, nil
	case v2beta1.ProtocolIeIDRanfunctionsAdded:
		return C.ProtocolIE_ID_id_RANfunctionsAdded, nil
	case v2beta1.ProtocolIeIDRanfunctionsDeleted:
		return C.ProtocolIE_ID_id_RANfunctionsDeleted, nil
	case v2beta1.ProtocolIeIDRanfunctionsModified:
		return C.ProtocolIE_ID_id_RANfunctionsModified, nil
	case v2beta1.ProtocolIeIDRanfunctionsRejected:
		return C.ProtocolIE_ID_id_RANfunctionsRejected, nil
	case v2beta1.ProtocolIeIDRicactionAdmittedItem:
		return C.ProtocolIE_ID_id_RICaction_Admitted_Item, nil
	case v2beta1.ProtocolIeIDRicactionID:
		return C.ProtocolIE_ID_id_RICactionID, nil
	case v2beta1.ProtocolIeIDRicactionNotAdmittedItem:
		return C.ProtocolIE_ID_id_RICaction_NotAdmitted_Item, nil
	case v2beta1.ProtocolIeIDRicactionsAdmitted:
		return C.ProtocolIE_ID_id_RICactions_Admitted, nil
	case v2beta1.ProtocolIeIDRicactionsNotAdmitted:
		return C.ProtocolIE_ID_id_RICactions_NotAdmitted, nil
	case v2beta1.ProtocolIeIDRicactionToBeSetupItem:
		return C.ProtocolIE_ID_id_RICaction_ToBeSetup_Item, nil
	case v2beta1.ProtocolIeIDRiccallProcessID:
		return C.ProtocolIE_ID_id_RICcallProcessID, nil
	case v2beta1.ProtocolIeIDRiccontrolAckRequest:
		return C.ProtocolIE_ID_id_RICcontrolAckRequest, nil
	case v2beta1.ProtocolIeIDRiccontrolHeader:
		return C.ProtocolIE_ID_id_RICcontrolHeader, nil
	case v2beta1.ProtocolIeIDRiccontrolMessage:
		return C.ProtocolIE_ID_id_RICcontrolMessage, nil
	case v2beta1.ProtocolIeIDRiccontrolStatus:
		return C.ProtocolIE_ID_id_RICcontrolStatus, nil
	case v2beta1.ProtocolIeIDRicindicationHeader:
		return C.ProtocolIE_ID_id_RICindicationHeader, nil
	case v2beta1.ProtocolIeIDRicindicationMessage:
		return C.ProtocolIE_ID_id_RICindicationMessage, nil
	case v2beta1.ProtocolIeIDRicindicationSn:
		return C.ProtocolIE_ID_id_RICindicationSN, nil
	case v2beta1.ProtocolIeIDRicindicationType:
		return C.ProtocolIE_ID_id_RICindicationType, nil
	case v2beta1.ProtocolIeIDRicrequestID:
		return C.ProtocolIE_ID_id_RICrequestID, nil
	case v2beta1.ProtocolIeIDRicsubscriptionDetails:
		return C.ProtocolIE_ID_id_RICsubscriptionDetails, nil
	case v2beta1.ProtocolIeIDTimeToWait:
		return C.ProtocolIE_ID_id_TimeToWait, nil
	case v2beta1.ProtocolIeIDRiccontrolOutcome:
		return C.ProtocolIE_ID_id_RICcontrolOutcome, nil
	case v2beta1.ProtocolIeIDE2nodeComponentConfigUpdate:
		return C.ProtocolIE_ID_id_E2nodeComponentConfigUpdate, nil
	case v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateItem:
		return C.ProtocolIE_ID_id_E2nodeComponentConfigUpdate_Item, nil
	case v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateAck:
		return C.ProtocolIE_ID_id_E2nodeComponentConfigUpdateAck, nil
	case v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateAckItem:
		return C.ProtocolIE_ID_id_E2nodeComponentConfigUpdateAck_Item, nil
	case v2beta1.ProtocolIeIDE2connectionSetup:
		return C.ProtocolIE_ID_id_E2connectionSetup, nil
	case v2beta1.ProtocolIeIDE2connectionSetupFailed:
		return C.ProtocolIE_ID_id_E2connectionSetupFailed, nil
	case v2beta1.ProtocolIeIDE2connectionSetupFailedItem:
		return C.ProtocolIE_ID_id_E2connectionSetupFailed_Item, nil
	case v2beta1.ProtocolIeIDE2connectionFailedItem:
		return C.ProtocolIE_ID_id_E2connectionFailed_Item, nil
	case v2beta1.ProtocolIeIDE2connectionUpdateItem:
		return C.ProtocolIE_ID_id_E2connectionUpdate_Item, nil
	case v2beta1.ProtocolIeIDE2connectionUpdateAdd:
		return C.ProtocolIE_ID_id_E2connectionUpdateAdd, nil
	case v2beta1.ProtocolIeIDE2connectionUpdateModify:
		return C.ProtocolIE_ID_id_E2connectionUpdateModify, nil
	case v2beta1.ProtocolIeIDE2connectionUpdateRemove:
		return C.ProtocolIE_ID_id_E2connectionUpdateRemove, nil
	case v2beta1.ProtocolIeIDE2connectionUpdateRemoveItem:
		return C.ProtocolIE_ID_id_E2connectionUpdateRemove_Item, nil
	case v2beta1.ProtocolIeIDTNLinformation:
		return C.ProtocolIE_ID_id_TNLinformation, nil
	case v2beta1.ProtocolIeIDTransactionID:
		return C.ProtocolIE_ID_id_TransactionID, nil
	default:
		return C.ProtocolIE_ID_t(-1), fmt.Errorf("unexpected value for ProtocolIE_IDT %v", pcIeID)
	}
}

func decodeProtocolIeID(pcIeIDC C.ProtocolIE_ID_t) *e2ap_commondatatypes.ProtocolIeId {
	ret := e2ap_commondatatypes.ProtocolIeId{
		Value: int32(pcIeIDC),
	}
	return &ret
}
