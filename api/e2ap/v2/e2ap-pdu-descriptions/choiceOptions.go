// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap_pdu_descriptions

import (
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"reflect"
)

//ToDo - move away from using Reflect package
var E2ApPduChoicemap = map[string]map[int]reflect.Type{
	"e2_ap_pdu": {
		1: reflect.TypeOf(E2ApPdu_InitiatingMessage{}),
		2: reflect.TypeOf(E2ApPdu_SuccessfulOutcome{}),
		3: reflect.TypeOf(E2ApPdu_UnsuccessfulOutcome{}),
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
		int64(v2.ProcedureCodeIDRICsubscription):               reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicSubscription{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDelete):         reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDelete{}),
		int64(v2.ProcedureCodeIDRICserviceUpdate):              reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicServiceUpdate{}),
		int64(v2.ProcedureCodeIDRICcontrol):                    reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicControl{}),
		int64(v2.ProcedureCodeIDE2setup):                       reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_E2Setup{}),
		int64(v2.ProcedureCodeIDE2nodeConfigurationUpdate):     reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		int64(v2.ProcedureCodeIDE2connectionUpdate):            reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_E2ConnectionUpdate{}),
		int64(v2.ProcedureCodeIDReset):                         reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_Reset_{}),
		int64(v2.ProcedureCodeIDRICindication):                 reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicIndication{}),
		int64(v2.ProcedureCodeIDRICserviceQuery):               reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicServiceQuery{}),
		int64(v2.ProcedureCodeIDErrorIndication):               reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_ErrorIndication{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDeleteRequired): reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDeleteRequired{}),
	},
	"so_values": {
		int64(v2.ProcedureCodeIDRICsubscription):           reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDelete):     reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{}),
		int64(v2.ProcedureCodeIDRICserviceUpdate):          reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{}),
		int64(v2.ProcedureCodeIDRICcontrol):                reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicControl{}),
		int64(v2.ProcedureCodeIDE2setup):                   reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_E2Setup{}),
		int64(v2.ProcedureCodeIDE2nodeConfigurationUpdate): reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		int64(v2.ProcedureCodeIDE2connectionUpdate):        reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{}),
		int64(v2.ProcedureCodeIDReset):                     reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_Reset_{}),
	},
	"uo_values": {
		int64(v2.ProcedureCodeIDRICsubscription):           reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{}),
		int64(v2.ProcedureCodeIDRICsubscriptionDelete):     reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{}),
		int64(v2.ProcedureCodeIDRICserviceUpdate):          reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{}),
		int64(v2.ProcedureCodeIDRICcontrol):                reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicControl{}),
		int64(v2.ProcedureCodeIDE2setup):                   reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup{}),
		int64(v2.ProcedureCodeIDE2nodeConfigurationUpdate): reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		int64(v2.ProcedureCodeIDE2connectionUpdate):        reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{}),
	},
	"ricsubscription_request_ie": {
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RfId{}),
		int64(v2.ProtocolIeIDRicsubscriptionDetails): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_Rsd{}),
	},
	"ricaction_to_be_setup_item": {
		int64(v2.ProtocolIeIDRicactionToBeSetupItem): reflect.TypeOf(e2ap_pdu_contents.RicactionToBeSetupItemIe_Ratbsi{}),
	},
	"ricsubscription_response_ie": {
		int64(v2.ProtocolIeIDRicrequestID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):         reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RfId{}),
		int64(v2.ProtocolIeIDRicactionsAdmitted):    reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_Raal{}),
		int64(v2.ProtocolIeIDRicactionsNotAdmitted): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_Ranal{}),
	},
	"ricaction_admitted_item_ie": {
		int64(v2.ProtocolIeIDRicactionAdmittedItem): reflect.TypeOf(e2ap_pdu_contents.RicactionAdmittedItemIe_Ranai{}),
	},
	"ricaction_not_admitted_item_ie": {
		int64(v2.ProtocolIeIDRicactionNotAdmittedItem): reflect.TypeOf(e2ap_pdu_contents.RicactionNotAdmittedItemIe_Ranai{}),
	},
	"ricsubscription_failure_ie": {
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_RfId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_C{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_Cd{}),
	},
	"ricsubscription_delete_request_ie": {
		int64(v2.ProtocolIeIDRicrequestID):  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequestIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequestIe_RfId{}),
	},
	"ricsubscription_delete_response_ie": {
		int64(v2.ProtocolIeIDRicrequestID):  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteResponseIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteResponseIe_RfId{}),
	},
	"ricsubscription_delete_failure_ie": {
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_RfId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_C{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_Cd{}),
	},
	"ricsubscription_delete_required_ie": {
		int64(v2.ProtocolIeIDRICsubscriptionToBeRemoved): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequiredIe_Rsdr{}),
	},
	"ricsubscription_with_cause_item_ie": {
		int64(v2.ProtocolIeIDRICsubscriptionWithCauseItem): reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionWithCauseItemIe_E2Curi{}),
	},
	"ricindication_ie": {
		int64(v2.ProtocolIeIDRicrequestID):         reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):        reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RfId{}),
		int64(v2.ProtocolIeIDRicactionID):          reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RaId{}),
		int64(v2.ProtocolIeIDRicindicationSn):      reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RiSn{}),
		int64(v2.ProtocolIeIDRicindicationType):    reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_Rit{}),
		int64(v2.ProtocolIeIDRicindicationHeader):  reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_Rih{}),
		int64(v2.ProtocolIeIDRicindicationMessage): reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_Rim{}),
		int64(v2.ProtocolIeIDRiccallProcessID):     reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RcpId{}),
	},
	"riccontrol_request_ie": {
		int64(v2.ProtocolIeIDRicrequestID):         reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):        reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RfId{}),
		int64(v2.ProtocolIeIDRiccallProcessID):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RcpId{}),
		int64(v2.ProtocolIeIDRiccontrolHeader):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_Rch{}),
		int64(v2.ProtocolIeIDRiccontrolMessage):    reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_Rcm{}),
		int64(v2.ProtocolIeIDRiccontrolAckRequest): reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_Rcar{}),
	},
	"riccontrol_acknowledge_ie": {
		int64(v2.ProtocolIeIDRicrequestID):      reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RfId{}),
		int64(v2.ProtocolIeIDRiccallProcessID):  reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RcpId{}),
		int64(v2.ProtocolIeIDRiccontrolOutcome): reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_Co{}),
	},
	"riccontrol_failure_ie": {
		int64(v2.ProtocolIeIDRicrequestID):      reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RrId{}),
		int64(v2.ProtocolIeIDRanfunctionID):     reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RfId{}),
		int64(v2.ProtocolIeIDRiccallProcessID):  reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RcpId{}),
		int64(v2.ProtocolIeIDCause):             reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_C{}),
		int64(v2.ProtocolIeIDRiccontrolOutcome): reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_Co{}),
	},
	"error_indication_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_TrId{}),
		int64(v2.ProtocolIeIDRicrequestID):           reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_Rr{}),
		int64(v2.ProtocolIeIDRanfunctionID):          reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_RfId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_C{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_Cd{}),
	},
	"e2setup_request_ie": {
		int64(v2.ProtocolIeIDTransactionID):                 reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_TrId{}),
		int64(v2.ProtocolIeIDGlobalE2nodeID):                reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_GE2NId{}),
		int64(v2.ProtocolIeIDRanfunctionsAdded):             reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_Rfl{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAddition): reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_E2Nccal{}),
	},
	"e2setup_response_ie": {
		int64(v2.ProtocolIeIDTransactionID):                    reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_TrId{}),
		int64(v2.ProtocolIeIDGlobalRicID):                      reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_GRicId{}),
		int64(v2.ProtocolIeIDRanfunctionsAccepted):             reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_RfIdl{}),
		int64(v2.ProtocolIeIDRanfunctionsRejected):             reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_RfIdcl{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck): reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_E2Nccaal{}),
	},
	"e2setup_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TrId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_C{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_Ttw{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_Cd{}),
		int64(v2.ProtocolIeIDTNLinformation):         reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TnlInfo{}),
	},
	"e2connection_update_ie": {
		int64(v2.ProtocolIeIDTransactionID):            reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_TrId{}),
		int64(v2.ProtocolIeIDE2connectionUpdateAdd):    reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2Cul{}),
		int64(v2.ProtocolIeIDE2connectionUpdateRemove): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2Curl{}),
		int64(v2.ProtocolIeIDE2connectionUpdateModify): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2Cul{}),
	},
	"e2connection_update_item_ie": {
		int64(v2.ProtocolIeIDE2connectionUpdateItem): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateItemIe_E2Curi{}),
	},
	"e2connection_update_remove_item_ie": {
		int64(v2.ProtocolIeIDE2connectionUpdateRemoveItem): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateRemoveItemIe_E2Curi{}),
	},
	"e2connection_update_ack_ie": {
		int64(v2.ProtocolIeIDTransactionID):           reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_TrId{}),
		int64(v2.ProtocolIeIDE2connectionSetup):       reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_E2Cul{}),
		int64(v2.ProtocolIeIDE2connectionSetupFailed): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_E2Csfl{}),
	},
	"e2connection_setup_failed_item_ie": {
		int64(v2.ProtocolIeIDE2connectionSetupFailedItem): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionSetupFailedItemIe_E2Csfi{}),
	},
	"e2connection_update_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_TrId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_C{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_Ttw{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_Cd{}),
	},
	"e2node_configuration_update_ie": {
		int64(v2.ProtocolIeIDTransactionID):                 reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_TrId{}),
		int64(v2.ProtocolIeIDGlobalE2nodeID):                reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_Ge2NId{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAddition): reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Nccal{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdate):   reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Nccul{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemoval):  reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Nccrl{}),
		int64(v2.ProtocolIeIDE2nodeTNLassociationRemoval):   reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Ntnlarl{}),
	},
	"e2node_component_config_addition_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigAdditionItemIe_E2Nccui{}),
	},
	"e2node_component_config_update_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdateItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIe_E2Nccui{}),
	},
	"e2node_component_config_removal_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemovalItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigRemovalItemIe_E2Nccri{}),
	},
	"e2node_tnlassociation_removal_item_ie": {
		int64(v2.ProtocolIeIDE2nodeTNLassociationRemovalItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeTnlassociationRemovalItemIe_E2Ntnlai{}),
	},
	"e2node_configuration_update_acknowledge_ie": {
		int64(v2.ProtocolIeIDTransactionID):                    reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_TrId{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionAck): reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2Nccaal{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdateAck):   reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2Nccual{}),
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemovalAck):  reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2Nccral{}),
	},
	"e2node_component_config_addition_ack_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigAdditionAckItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItemIe_E2Nccaai{}),
	},
	"e2node_component_config_update_ack_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigUpdateAckItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIe_E2Nccuai{}),
	},
	"e2node_component_config_removal_ack_item_ie": {
		int64(v2.ProtocolIeIDE2nodeComponentConfigRemovalAckItem): reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItemIe_E2Nccrai{}),
	},
	"e2node_configuration_update_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_TrId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_C{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_Ttw{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_Cd{}),
	},
	"reset_request_ie": {
		int64(v2.ProtocolIeIDTransactionID): reflect.TypeOf(e2ap_pdu_contents.ResetRequestIe_TrId{}),
		int64(v2.ProtocolIeIDCause):         reflect.TypeOf(e2ap_pdu_contents.ResetRequestIe_C{}),
	},
	"reset_response_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.ResetResponseIe_TrId{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.ResetResponseIe_Cd{}),
	},
	"ric_service_update_ie": {
		int64(v2.ProtocolIeIDTransactionID):        reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_TrId{}),
		int64(v2.ProtocolIeIDRanfunctionsAdded):    reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_Rfl{}),
		int64(v2.ProtocolIeIDRanfunctionsModified): reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_Rfl{}),
		int64(v2.ProtocolIeIDRanfunctionsDeleted):  reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_Rfidl{}),
	},
	"ranfunction_item_ie": {
		int64(v2.ProtocolIeIDRanfunctionItem): reflect.TypeOf(e2ap_pdu_contents.RanfunctionItemIe_Rfi{}),
	},
	"ranfunction_id_item_ie": {
		int64(v2.ProtocolIeIDRanfunctionIDItem): reflect.TypeOf(e2ap_pdu_contents.RanfunctionIdItemIe_RfId{}),
	},
	"ric_service_update_acknowledge_ie": {
		int64(v2.ProtocolIeIDTransactionID):        reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_TrId{}),
		int64(v2.ProtocolIeIDRanfunctionsAccepted): reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_RfIdl{}),
		int64(v2.ProtocolIeIDRanfunctionsRejected): reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_RfIdcl{}),
	},
	"ranfunction_idcause_item_ie": {
		int64(v2.ProtocolIeIDRanfunctionIeCauseItem): reflect.TypeOf(e2ap_pdu_contents.RanfunctionIdcauseItemIe_RfIdci{}),
	},
	"ric_service_update_failure_ie": {
		int64(v2.ProtocolIeIDTransactionID):          reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_TrId{}),
		int64(v2.ProtocolIeIDCause):                  reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_C{}),
		int64(v2.ProtocolIeIDTimeToWait):             reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_Ttw{}),
		int64(v2.ProtocolIeIDCriticalityDiagnostics): reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_Cd{}),
	},
	"ric_service_query_ie": {
		int64(v2.ProtocolIeIDTransactionID):        reflect.TypeOf(e2ap_pdu_contents.RicServiceQueryIe_TrId{}),
		int64(v2.ProtocolIeIDRanfunctionsAccepted): reflect.TypeOf(e2ap_pdu_contents.RicServiceQueryIe_RfIdl{}),
	},
}
