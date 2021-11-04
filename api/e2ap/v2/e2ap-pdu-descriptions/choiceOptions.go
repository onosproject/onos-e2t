// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2ap_pdu_descriptions

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	"reflect"
)

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
