// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2ap_pdu_descriptions

import (
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap_go/v2/e2ap-ies"
	"reflect"
)

var E2ApPduChoicemap = map[string]map[int]reflect.Type{
	"e2_ap_pdu": {
		1: reflect.TypeOf(E2ApPduRicServiceQuery_InitiatingMessage{}),
		2: reflect.TypeOf(E2ApPduRicServiceQuery_SuccessfulOutcome{}),
		3: reflect.TypeOf(E2ApPduRicServiceQuery_UnsuccessfulOutcome{}),
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
