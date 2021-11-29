// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2ap_pdu_descriptions

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-pdu-contents"
	"reflect"
)

//ToDo - move away from using Reflect package
var E2ApPduChoicemap = map[string]map[int]reflect.Type{
	"e2_ap_pdu": {
		1: reflect.TypeOf(E2ApPdu_InitiatingMessage{}),
		2: reflect.TypeOf(E2ApPdu_SuccessfulOutcome{}),
		3: reflect.TypeOf(E2ApPdu_UnsuccessfulOutcome{}),
	},
	"im_values": {
		1:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicSubscription{}),
		2:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDelete{}),
		3:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicServiceUpdate{}),
		4:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicControl{}),
		5:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_E2Setup{}),
		6:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		7:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_E2ConnectionUpdate{}),
		8:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_Reset_{}),
		9:  reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicIndication{}),
		10: reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicServiceQuery{}),
		11: reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_ErrorIndication{}),
		12: reflect.TypeOf(InitiatingMessageE2ApElementaryProcedures_RicSubscriptionDeleteRequired{}),
	},
	"so_values": {
		1: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{}),
		2: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{}),
		3: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{}),
		4: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_RicControl{}),
		5: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_E2Setup{}),
		6: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		7: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{}),
		8: reflect.TypeOf(SuccessfulOutcomeE2ApElementaryProcedures_Reset_{}),
	},
	"uo_values": {
		1: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscription{}),
		2: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicSubscriptionDelete{}),
		3: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicServiceUpdate{}),
		4: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_RicControl{}),
		5: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_E2Setup{}),
		6: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_E2NodeConfigurationUpdate{}),
		7: reflect.TypeOf(UnsuccessfulOutcomeE2ApElementaryProcedures_E2ConnectionUpdate{}),
	},
	"ricsubscription_request_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionRequestIe_Rsd{}),
	},
	"ricaction_to_be_setup_item": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicactionToBeSetupItemIe_Ratbsi{}),
	},
	"ricsubscription_response_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_Raal{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionResponseIe_Ranal{}),
	},
	"ricaction_admitted_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicactionAdmittedItemIe_Ranai{}),
	},
	"ricaction_not_admitted_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicactionNotAdmittedItemIe_Ranai{}),
	},
	"ricsubscription_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_C{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionFailureIe_Cd{}),
	},
	"ricsubscription_delete_request_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequestIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequestIe_RfId{}),
	},
	"ricsubscription_delete_response_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteResponseIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteResponseIe_RfId{}),
	},
	"ricsubscription_delete_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_C{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteFailureIe_Cd{}),
	},
	"ricsubscription_delete_required_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionDeleteRequiredIe_Rsdr{}),
	},
	"ricsubscription_with_cause_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicsubscriptionWithCauseItemIe_E2Curi{}),
	},
	"ricindication_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RaId{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RiSn{}),
		5: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_Rit{}),
		6: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_Rih{}),
		7: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_Rim{}),
		8: reflect.TypeOf(e2ap_pdu_contents.RicindicationIe_RcpId{}),
	},
	"riccontrol_request_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_RcpId{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_Rch{}),
		5: reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_Rcm{}),
		6: reflect.TypeOf(e2ap_pdu_contents.RiccontrolRequestIe_Rcar{}),
	},
	"riccontrol_acknowledge_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_RcpId{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RiccontrolAcknowledgeIe_Co{}),
	},
	"riccontrol_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RfId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_RcpId{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_C{}),
		5: reflect.TypeOf(e2ap_pdu_contents.RiccontrolFailureIe_Co{}),
	},
	"error_indication_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_Rr{}),
		3: reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_Rfid{}),
		4: reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_C{}),
		5: reflect.TypeOf(e2ap_pdu_contents.ErrorIndicationIe_Cd{}),
	},
	"e2setup_request_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_GE2NId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_Rfl{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2SetupRequestIe_E2Nccal{}),
	},
	"e2setup_response_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_GRicId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_RfIdl{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_RfIdcl{}),
		5: reflect.TypeOf(e2ap_pdu_contents.E2SetupResponseIe_E2Nccaal{}),
	},
	"e2setup_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_C{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_Ttw{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_Cd{}),
		5: reflect.TypeOf(e2ap_pdu_contents.E2SetupFailureIe_TnlInfo{}),
	},
	"e2connection_update_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2Cul{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateIe_E2Curl{}),
	},
	"e2connection_update_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateItemIe_E2Curi{}),
	},
	"e2connection_update_remove_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateRemoveItemIe_E2Curi{}),
	},
	"e2connection_update_ack_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_E2Cul{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateAckIe_E2Csfl{}),
	},
	"e2connection_setup_failed_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionSetupFailedItemIe_E2Csfi{}),
	},
	"e2connection_update_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_C{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_Ttw{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2ConnectionUpdateFailureIe_Cd{}),
	},
	"e2node_configuration_update_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_Ge2NId{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Nccal{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Nccul{}),
		5: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Nccrl{}),
		6: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateIe_E2Ntnlarl{}),
	},
	"e2node_component_config_addition_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigAdditionItemIe_E2Nccui{}),
	},
	"e2node_component_config_update_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigUpdateItemIe_E2Nccui{}),
	},
	"e2node_component_config_removal_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigRemovalItemIe_E2Nccri{}),
	},
	"e2node_tnlassociation_removal_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeTnlassociationRemovalItemIe_E2Ntnlai{}),
	},
	"e2node_configuration_update_acknowledge_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2Nccaal{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2Nccual{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateAcknowledgeIe_E2Nccral{}),
	},
	"e2node_component_config_addition_ack_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigAdditionAckItemIe_E2Nccaai{}),
	},
	"e2node_component_config_update_ack_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigUpdateAckItemIe_E2Nccuai{}),
	},
	"e2node_component_config_removal_ack_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeComponentConfigRemovalAckItemIe_E2Nccrai{}),
	},
	"e2node_configuration_update_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_C{}),
		3: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_Value{}),
		4: reflect.TypeOf(e2ap_pdu_contents.E2NodeConfigurationUpdateFailureIe_Cd{}),
	},
	"reset_request_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.ResetRequestIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.ResetRequestIe_C{}),
	},
	"reset_response_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.ResetResponseIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.ResetResponseIe_Cd{}),
	},
	"ric_service_update_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_Rfl{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicserviceUpdateIe_Rfidl{}),
	},
	"ranfunction_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RanfunctionItemIe_Rfi{}),
	},
	"ranfunction_id_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RanfunctionIdItemIe_RfId{}),
	},
	"ric_service_update_acknowledge_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_RfIdl{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateAcknowledgeIe_RfIdcl{}),
	},
	"ranfunction_idcause_item_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RanfunctionIdcauseItemIe_RfIdci{}),
	},
	"ric_service_update_failure_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_C{}),
		3: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_Ttw{}),
		4: reflect.TypeOf(e2ap_pdu_contents.RicServiceUpdateFailureIe_Cd{}),
	},
	"ric_service_query_ie": {
		1: reflect.TypeOf(e2ap_pdu_contents.RicServiceQueryIe_TrId{}),
		2: reflect.TypeOf(e2ap_pdu_contents.RicServiceQueryIe_RfIdl{}),
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
