// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package choiceOptions

import (
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2ap_pdu_descriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"reflect"
)

// move away from using Reflect package
var E2ApPduChoicemap = map[string]map[int]reflect.Type{
	"e2_ap_pdu": {
		1: reflect.TypeOf(e2ap_pdu_descriptions.E2ApPdu_InitiatingMessage{}),
		2: reflect.TypeOf(e2ap_pdu_descriptions.E2ApPdu_SuccessfulOutcome{}),
		3: reflect.TypeOf(e2ap_pdu_descriptions.E2ApPdu_UnsuccessfulOutcome{}),
	},
	"cause": {
		1: reflect.TypeOf(e2ap_ies.Cause_RicRequest{}),
		2: reflect.TypeOf(e2ap_ies.Cause_RicService{}),
		3: reflect.TypeOf(e2ap_ies.Cause_E2Node{}),
		4: reflect.TypeOf(e2ap_ies.Cause_Transport{}),
		5: reflect.TypeOf(e2ap_ies.Cause_Protocol{}),
		6: reflect.TypeOf(e2ap_ies.Cause_Misc{}),
	},
	"e2node_component_id": {
		1: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeNg{}),
		2: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeXn{}),
		3: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1{}),
		4: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1{}),
		5: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{}),
		6: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeS1{}),
		7: reflect.TypeOf(e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeX2{}),
	},
	"enb_id": {
		1: reflect.TypeOf(e2ap_ies.EnbId_MacroENbId{}),
		2: reflect.TypeOf(e2ap_ies.EnbId_HomeENbId{}),
		3: reflect.TypeOf(e2ap_ies.EnbId_ShortMacroENbId{}),
		4: reflect.TypeOf(e2ap_ies.EnbId_LongMacroENbId{}),
	},
	"enb_id_choice": {
		1: reflect.TypeOf(e2ap_ies.EnbIdChoice_EnbIdMacro{}),
		2: reflect.TypeOf(e2ap_ies.EnbIdChoice_EnbIdShortmacro{}),
		3: reflect.TypeOf(e2ap_ies.EnbIdChoice_EnbIdLongmacro{}),
	},
	"engnb_id": {
		1: reflect.TypeOf(e2ap_ies.EngnbId_GNbId{}),
	},
	"global_e2node_id": {
		1: reflect.TypeOf(e2ap_ies.GlobalE2NodeId_GNb{}),
		2: reflect.TypeOf(e2ap_ies.GlobalE2NodeId_EnGNb{}),
		3: reflect.TypeOf(e2ap_ies.GlobalE2NodeId_NgENb{}),
		4: reflect.TypeOf(e2ap_ies.GlobalE2NodeId_ENb{}),
	},
	"global_ng_rannode_id": {
		1: reflect.TypeOf(e2ap_ies.GlobalNgRannodeId_GNb{}),
		2: reflect.TypeOf(e2ap_ies.GlobalNgRannodeId_NgENb{}),
	},
	"gnb_id_choice": {
		1: reflect.TypeOf(e2ap_ies.GnbIdChoice_GnbId{}),
	},
}

var E2ApPduCanonicalChoicemap = map[string]map[int64]reflect.Type{
	"im_values": {
		int64(v2.ProcedureCodeIDRICsubscription):               reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscription{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDelete):         reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDelete{}),
		int64(v2.ProcedureCodeIDRICserviceUpdate):              reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicServiceUpdate{}),
		int64(v2.ProcedureCodeIDRICcontrol):                    reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicControl{}),
		int64(v2.ProcedureCodeIDE2setup):                       reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_E2Setup{}),
		int64(v2.ProcedureCodeIDE2nodeConfigurationUpdate):     reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		int64(v2.ProcedureCodeIDE2connectionUpdate):            reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_E2ConnectionUpdate{}),
		int64(v2.ProcedureCodeIDReset):                         reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_Reset_{}),
		int64(v2.ProcedureCodeIDRICindication):                 reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicIndication{}),
		int64(v2.ProcedureCodeIDRICserviceQuery):               reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicServiceQuery{}),
		int64(v2.ProcedureCodeIDErrorIndication):               reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_ErrorIndication{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDeleteRequired): reflect.TypeOf(e2ap_pdu_descriptions.InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDeleteRequired{}),
	},
	"so_values": {
		int64(v2.ProcedureCodeIDRICsubscription):           reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDelete):     reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{}),
		int64(v2.ProcedureCodeIDRICserviceUpdate):          reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{}),
		int64(v2.ProcedureCodeIDRICcontrol):                reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_RicControl{}),
		int64(v2.ProcedureCodeIDE2setup):                   reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2Setup{}),
		int64(v2.ProcedureCodeIDE2nodeConfigurationUpdate): reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		int64(v2.ProcedureCodeIDE2connectionUpdate):        reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{}),
		int64(v2.ProcedureCodeIDReset):                     reflect.TypeOf(e2ap_pdu_descriptions.SuccessfulOutcomeE2ApElementaryProcedures_Reset_{}),
	},
	"uo_values": {
		int64(v2.ProcedureCodeIDRICsubscription):           reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDelete):     reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{}),
		int64(v2.ProcedureCodeIDRICserviceUpdate):          reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{}),
		int64(v2.ProcedureCodeIDRICcontrol):                reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_RicControl{}),
		int64(v2.ProcedureCodeIDE2setup):                   reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup{}),
		int64(v2.ProcedureCodeIDE2nodeConfigurationUpdate): reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		int64(v2.ProcedureCodeIDE2connectionUpdate):        reflect.TypeOf(e2ap_pdu_descriptions.UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{}),
	},
	"ricsubscription_request_ie": {
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDRicsubscriptionDetails): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RicsubscriptionDetails{}),
	},
	"ricaction_to_be_setup_item_ie": {
		int64(v2.ProtocolIeIDRicactionToBeSetupItem): reflect.TypeOf(e2ap_pdu_contents.RicactionToBeSetupItemIe_RicactionToBeSetupItem{}),
	},
	"ricsubscription_response_ie": {
		int64(v2.ProtocolIeIDRicrequestID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):         reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDRicactionsAdmitted):    reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RicactionsAdmitted{}),
		int64(v2.ProtocolIeIDRicactionsNotAdmitted): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RicactionsNotAdmitted{}),
	},
	"ricaction_admitted_item_ie": {
		int64(v2.ProtocolIeIDRicactionAdmittedItem): reflect.TypeOf(e2ap_pdu_contents.RicactionAdmittedItemIe_RicactionAdmittedItem{}),
	},
	"ricaction_not_admitted_item_ie": {
		int64(v2.ProtocolIeIDRicactionNotAdmittedItem): reflect.TypeOf(e2ap_pdu_contents.RicactionNotAdmittedItemIe_RicactionNotAdmittedItem{}),
	},
	"ricsubscription_failure_ie": {
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_Cause{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_CriticalityDiagnostics{}),
	},
	"ricsubscription_delete_request_ie": {
		int64(v2.ProtocolIeIDRicrequestID):  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequestIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequestIe_RanfunctionId{}),
	},
	"ricsubscription_delete_response_ie": {
		int64(v2.ProtocolIeIDRicrequestID):  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteResponseIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteResponseIe_RanfunctionId{}),
	},
	"ricsubscription_delete_failure_ie": {
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_Cause{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_CriticalityDiagnostics{}),
	},
	"ricsubscription_delete_required_ie": {
		int64(v2.ProtocolIeIDRICsubscriptionToBeRemoved): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequiredIe_RicsubscriptionToBeRemoved{}),
	},
	"ricsubscription_with_cause_item_ie": {
		int64(v2.ProtocolIeIDRICsubscriptionWithCauseItem): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionWithCauseItemIe_RicsubscriptionWithCauseItem{}),
	},
	"ricindication_ie": {
		int64(v2.ProtocolIeIDRicrequestID):         reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):        reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDRicactionID):          reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RicactionId{}),
		int64(v2.ProtocolIeIDRicindicationSn):      reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RicindicationSn{}),
		int64(v2.ProtocolIeIDRicindicationType):    reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RicindicationType{}),
		int64(v2.ProtocolIeIDRicindicationHeader):  reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RicindicationHeader{}),
		int64(v2.ProtocolIeIDRicindicationMessage): reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RicindicationMessage{}),
		int64(v2.ProtocolIeIDRiccallProcessID):     reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RiccallProcessId{}),
	},
	"riccontrol_request_ie": {
		int64(v2.ProtocolIeIDRicrequestID):         reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):        reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDRiccallProcessID):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RiccallProcessId{}),
		int64(v2.ProtocolIeIDRiccontrolHeader):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RiccontrolHeader{}),
		int64(v2.ProtocolIeIDRiccontrolMessage):    reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RiccontrolMessage{}),
		int64(v2.ProtocolIeIDRiccontrolAckRequest): reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RiccontrolAckRequest{}),
	},
	"riccontrol_acknowledge_ie": {
		int64(v2.ProtocolIeIDRicrequestID):      reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDRiccallProcessID):  reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RiccallProcessId{}),
		int64(v2.ProtocolIeIDRiccontrolOutcome): reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RiccontrolOutcome{}),
	},
	"riccontrol_failure_ie": {
		int64(v2.ProtocolIeIDRicrequestID):      reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDRiccallProcessID):  reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RiccallProcessId{}),
		int64(v2.ProtocolIeIDCause):             reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_Cause{}),
		int64(v2.ProtocolIeIDRiccontrolOutcome): reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RiccontrolOutcome{}),
	},
	"error_indication_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_TransactionId{}),
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_RicrequestId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_RanfunctionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_Cause{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_CriticalityDiagnostics{}),
	},
	"e2setup_request_ie": {
		int64(v2.ProtocolIeIDTransactionID):                 reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_TransactionId{}),
		int64(v2.ProtocolIeIDGlobalE2nodeID):                reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_GlobalE2NodeId{}),
		int64(v2.ProtocolIeIDRanfunctionsAdded):             reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_RanfunctionsAdded{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAddition): reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_E2NodeComponentConfigAddition{}),
	},
	"e2setup_response_ie": {
		int64(v2.ProtocolIeIDTransactionID):                    reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_TransactionId{}),
		int64(v2.ProtocolIeIDGlobalRicID):                      reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_GlobalRicId{}),
		int64(v2.ProtocolIeIDRanfunctionsAccepted):             reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_RanfunctionsAccepted{}),
		int64(v2.ProtocolIeIDRanfunctionsRejected):             reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_RanfunctionsRejected{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck): reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_E2NodeComponentConfigAdditionAck{}),
	},
	"e2setup_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TransactionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_Cause{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TimeToWait{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_CriticalityDiagnostics{}),
		int64(v2.ProtocolIeIDTNLinformation):         reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TnlInformation{}),
	},
	"e2connection_update_ie": {
		int64(v2.ProtocolIeIDTransactionID):            reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_TransactionId{}),
		int64(v2.ProtocolIeIDE2connectionUpdateAdd):    reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2ConnectionUpdateAdd{}),
		int64(v2.ProtocolIeIDE2connectionUpdateRemove): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2ConnectionUpdateRemove{}),
		int64(v2.ProtocolIeIDE2connectionUpdateModify): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2ConnectionUpdateModify{}),
	},
	"e2connection_update_item_ie": {
		int64(v2.ProtocolIeIDE2connectionUpdateItem): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateItemIe_E2ConnectionUpdateItem{}),
	},
	"e2connection_update_remove_item_ie": {
		int64(v2.ProtocolIeIDE2connectionUpdateRemoveItem): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateRemoveItemIe_E2ConnectionUpdateRemoveItem{}),
	},
	"e2connection_update_ack_ie": {
		int64(v2.ProtocolIeIDTransactionID):           reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_TransactionId{}),
		int64(v2.ProtocolIeIDE2connectionSetup):       reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_E2ConnectionSetup{}),
		int64(v2.ProtocolIeIDE2connectionSetupFailed): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_E2ConnectionSetupFailed{}),
	},
	"e2connection_setup_failed_item_ie": {
		int64(v2.ProtocolIeIDE2connectionSetupFailedItem): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionSetupFailedItemIe_E2ConnectionSetupFailedItem{}),
	},
	"e2connection_update_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_TransactionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_Cause{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_TimeToWait{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_CriticalityDiagnostics{}),
	},
	"e2node_configuration_update_ie": {
		int64(v2.ProtocolIeIDTransactionID):                 reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_TransactionId{}),
		int64(v2.ProtocolIeIDGlobalE2nodeID):                reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_GlobalE2NodeId{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAddition): reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2NodeComponentConfigAddition{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdate):   reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2NodeComponentConfigUpdate{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemoval):  reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2NodeComponentConfigRemoval{}),
		int64(v2.ProtocolIeIDE2nodeTNLassociationRemoval):   reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2NodeTnlassociationRemoval{}),
	},
	"e2node_component_config_addition_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigAdditionItemIe_E2NodeComponentConfigAdditionItem{}),
	},
	"e2node_component_config_update_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdateItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIe_E2NodeComponentConfigUpdateItem{}),
	},
	"e2node_component_config_removal_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigRemovalItemIe_E2NodeComponentConfigRemovalItem{}),
	},
	"e2node_tnlassociation_removal_item_ie": {
		int64(v2.ProtocolIeIDE2nodeTNLassociationRemovalItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeTnlassociationRemovalItemIe_E2NodeTnlassociationRemovalItem{}),
	},
	"e2node_configuration_update_acknowledge_ie": {
		int64(v2.ProtocolIeIDTransactionID):                    reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_TransactionId{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck): reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2NodeComponentConfigAdditionAck{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck):   reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2NodeComponentConfigUpdateAck{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemovalAck):  reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2NodeComponentConfigRemovalAck{}),
	},
	"e2node_component_config_addition_ack_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionAckItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItemIe_E2NodeComponentConfigAdditionAckItem{}),
	},
	"e2node_component_config_update_ack_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIe_E2NodeComponentConfigUpdateAckItem{}),
	},
	"e2node_component_config_removal_ack_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemovalAckItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItemIe_E2NodeComponentConfigRemovalAckItem{}),
	},
	"e2node_configuration_update_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_TransactionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_Cause{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_TimeToWait{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_CriticalityDiagnostics{}),
	},
	"reset_request_ie": {
		int64(v2.ProtocolIeIDTransactionID): reflect.TypeOf(e2ap_pdu_contents.ResetRequestIe_TransactionId{}),
		int64(v2.ProtocolIeIDCause):         reflect.TypeOf(e2ap_pdu_contents.ResetRequestIe_Cause{}),
	},
	"reset_response_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.ResetResponseIe_TransactionId{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.ResetResponseIe_CriticalityDiagnostics{}),
	},
	"ric_service_update_ie": {
		int64(v2.ProtocolIeIDTransactionID):        reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_TransactionId{}),
		int64(v2.ProtocolIeIDRanfunctionsAdded):    reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_RanfunctionsAdded{}),
		int64(v2.ProtocolIeIDRanfunctionsModified): reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_RanfunctionsModified{}),
		int64(v2.ProtocolIeIDRanfunctionsDeleted):  reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_RanfunctionsDeleted{}),
	},
	"ranfunction_item_ie": {
		int64(v2.ProtocolIeIDRanfunctionItem): reflect.TypeOf(e2ap_pdu_contents.RanfunctionItemIe_RanfunctionItem{}),
	},
	"ranfunction_id_item_ie": {
		int64(v2.ProtocolIeIDRanfunctionIDItem): reflect.TypeOf(e2ap_pdu_contents.RanfunctionIdItemIe_RanfunctionIdItem{}),
	},
	"ric_service_update_acknowledge_ie": {
		int64(v2.ProtocolIeIDTransactionID):        reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_TransactionId{}),
		int64(v2.ProtocolIeIDRanfunctionsAccepted): reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_RanfunctionsAccepted{}),
		int64(v2.ProtocolIeIDRanfunctionsRejected): reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_RanfunctionsRejected{}),
	},
	"ranfunction_idcause_item_ie": {
		int64(v2.ProtocolIeIDRanfunctionIeCauseItem): reflect.TypeOf(e2ap_pdu_contents.RanfunctionIdcauseItemIe_RanfunctionIecauseItem{}),
	},
	"ric_service_update_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_TransactionId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_Cause{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_TimeToWait{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_CriticalityDiagnostics{}),
	},
	"ric_service_query_ie": {
		int64(v2.ProtocolIeIDTransactionID):        reflect.TypeOf(e2ap_pdu_contents.RicServiceQueryIe_TransactionId{}),
		int64(v2.ProtocolIeIDRanfunctionsAccepted): reflect.TypeOf(e2ap_pdu_contents.RicServiceQueryIe_RanfunctionsAccepted{}),
	},
}
