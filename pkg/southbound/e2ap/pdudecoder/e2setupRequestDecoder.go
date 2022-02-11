// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdudecoder

import (
	"fmt"
	v2 "github.com/onosproject/onos-e2t/api/e2ap/v2"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2SetupRequest(request *e2appducontents.E2SetupRequest) (*int32, *types.E2NodeIdentity, *types.RanFunctions,
	[]*types.E2NodeComponentConfigAdditionItem, error) {

	var err error
	var transactionID int32
	var globalE2NodeID *e2apies.GlobalE2NodeId
	ranFunctionsList := make(types.RanFunctions)
	e2nccul := make([]*types.E2NodeComponentConfigAdditionItem, 0)
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			transactionID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDGlobalE2nodeID) {
			globalE2NodeID = v.GetValue().GetGE2NId()
		}
		if v.Id == int32(v2.ProtocolIeIDRanfunctionsAdded) {
			ranFunctionsIe := v.GetValue()
			if ranFunctionsIe == nil {
				return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionsAdded")
			}
			for _, rfIe := range ranFunctionsIe.GetRfl().GetValue() {
				ranFunctionsList[types.RanFunctionID(rfIe.GetValue().GetRfi().GetRanFunctionId().GetValue())] = types.RanFunctionItem{
					Description: rfIe.GetValue().GetRfi().GetRanFunctionDefinition().GetValue(),
					Revision:    types.RanFunctionRevision(rfIe.GetValue().GetRfi().GetRanFunctionRevision().GetValue()),
					OID:         types.RanFunctionOID(rfIe.GetValue().GetRfi().GetRanFunctionOid().GetValue()),
				}
			}
		}
		if v.Id == int32(v2.ProtocolIeIDE2nodeComponentConfigAddition) {
			list := v.GetValue().GetE2Nccal().GetValue()
			for _, ie := range list {
				e2nccuai := types.E2NodeComponentConfigAdditionItem{}
				e2nccuai.E2NodeComponentType = ie.GetValue().GetE2Nccui().GetE2NodeComponentInterfaceType()
				e2nccuai.E2NodeComponentID = ie.GetValue().GetE2Nccui().GetE2NodeComponentId()
				e2nccuai.E2NodeComponentConfiguration = *ie.GetValue().GetE2Nccui().GetE2NodeComponentConfiguration()

				e2nccul = append(e2nccul, &e2nccuai)
			}
		}
	}

	nodeIdentity, err := ExtractE2NodeIdentity(globalE2NodeID, e2nccul)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return &transactionID, nodeIdentity, &ranFunctionsList, e2nccul, nil
}

func DecodeE2SetupRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*int32, *types.E2NodeIdentity, *types.RanFunctions,
	[]*types.E2NodeComponentConfigAdditionItem, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2setup := e2apPdu.GetInitiatingMessage().GetValue().GetE2Setup()
	if e2setup == nil {
		return nil, nil, nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}
	return DecodeE2SetupRequest(e2setup)
}

func GetE2NodeID(nodeID []byte, length int) string {
	var result uint64
	for i, b := range nodeID {
		result += uint64(b) << ((len(nodeID) - i - 1) * 8)
	}
	if length%8 != 0 {
		result = result >> (8 - length%8)
	}
	return fmt.Sprintf("%x", result)
}

func ExtractE2NodeIdentity(ge2nID *e2apies.GlobalE2NodeId, e2ncu interface{}) (*types.E2NodeIdentity, error) {
	var nodeIdentity *types.E2NodeIdentity
	var err error

	switch e2NodeID := ge2nID.GetGlobalE2NodeId().(type) {
	case *e2apies.GlobalE2NodeId_GNb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.GNb.GetGlobalGNbId().GetPlmnId().GetValue())
		if err != nil {
			return nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeGNB
		choice, ok := e2NodeID.GNb.GetGlobalGNbId().GetGnbId().GetGnbIdChoice().(*e2apies.GnbIdChoice_GnbId)
		if !ok {
			return nil, fmt.Errorf("expected a gNBId")
		}
		nodeIdentity.NodeIdentifier = choice.GnbId.GetValue()
		nodeIdentity.NodeIDLength = int(choice.GnbId.Len)
		if e2NodeID.GNb.GNbCuUpId != nil {
			nodeIdentity.CuID = &e2NodeID.GNb.GNbCuUpId.Value
		}
		if e2NodeID.GNb.GNbDuId != nil {
			nodeIdentity.DuID = &e2NodeID.GNb.GNbDuId.Value
		}

		// ToDo - how to deal with EnbID??
		//if e2NodeID.GNb.GlobalEnGNbId != nil {
		//copy(nodeIdentity.Plmn[:], e2NodeID.GNb.GlobalEnGNbId.PLmnIdentity.Value[:])
		//}

	case *e2apies.GlobalE2NodeId_EnGNb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.EnGNb.GetGlobalEnGNbId().GetPLmnIdentity().GetValue())
		if err != nil {
			return nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeEnGNB
		nodeIdentity.NodeIDLength = int(e2NodeID.EnGNb.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetLen())
		nodeIdentity.NodeIdentifier = e2NodeID.EnGNb.GetGlobalEnGNbId().GetGNbId().GetGNbId().GetValue()
		if e2NodeID.EnGNb.EnGNbCuUpId != nil {
			nodeIdentity.CuID = &e2NodeID.EnGNb.EnGNbCuUpId.Value
		}
		if e2NodeID.EnGNb.EnGNbDuId != nil {
			nodeIdentity.DuID = &e2NodeID.EnGNb.EnGNbDuId.Value
		}
		//return nil, fmt.Errorf("getting identifier of EnGnb not yet implemented")

	case *e2apies.GlobalE2NodeId_NgENb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.NgENb.GetGlobalNgENbId().GetPlmnId().GetValue())
		if err != nil {
			return nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeNgENB
		nodeIdentity.NodeIdentifier, nodeIdentity.NodeIDLength, err = ExtractEnbIDchoice(e2NodeID.NgENb.GetGlobalNgENbId().GetEnbId())
		if err != nil {
			return nil, err
		}
		if e2NodeID.NgENb.NgEnbDuId != nil {
			nodeIdentity.DuID = &e2NodeID.NgENb.NgEnbDuId.Value
		}
		return nil, fmt.Errorf("getting identifier of ngENb not yet implemented")

	case *e2apies.GlobalE2NodeId_ENb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.ENb.GetGlobalENbId().GetPLmnIdentity().GetValue())
		if err != nil {
			return nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeENB
		nodeIdentity.NodeIdentifier, nodeIdentity.NodeIDLength, err = ExtractEnbID(e2NodeID.ENb.GetGlobalENbId().GetENbId())
		if err != nil {
			return nil, err
		}

		// Goal of the below code block assigns the CU and DU IDs if the node type is eNB and the eNB is not the standalone eNB but CU/DU eNB
		// In the E2AP v1.01, the eNB ID field does not have CU and DU ID unlike the gNB ID field
		// If the eNB consists of CU and DU like SD-RAN OAI, there is no way to assign CU/DU IDs by referring to the eNB ID field
		// To avoid this problem, 'E2nodeComponentConfigUpdate` can be used, which includes CU and DU IDs (Current OAI CU and DU report their CU and DU IDs with this field)
		// The below block decodes the 'E2nodeComponentConfigUpdate` field and assign CU and DU ID only for the eNB CU/DU case
		// Since e2AP v2.0 has the CU and DU ID in eNB or NgEnb field, this block is only submitted into e2ap101 branch
		switch list := e2ncu.(type) {
		case []*types.E2NodeComponentConfigAdditionItem:
			if len(list) != 0 {
				for _, c := range list {
					switch id := c.E2NodeComponentID.GetE2NodeComponentId().(type) {
					case *e2apies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1:
						nodeIdentity.CuID = &id.E2NodeComponentInterfaceTypeE1.GNbCuCpId.Value
					case *e2apies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1:
						nodeIdentity.DuID = &id.E2NodeComponentInterfaceTypeF1.GNbDuId.Value
					}
				}
			}
		case []*types.E2NodeComponentConfigUpdateItem:
			if len(list) != 0 {
				for _, c := range list {
					switch id := c.E2NodeComponentID.GetE2NodeComponentId().(type) {
					case *e2apies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1:
						nodeIdentity.CuID = &id.E2NodeComponentInterfaceTypeE1.GNbCuCpId.Value
					case *e2apies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1:
						nodeIdentity.DuID = &id.E2NodeComponentInterfaceTypeF1.GNbDuId.Value
					}
				}
			}
		case []*e2appducontents.E2NodeComponentConfigUpdateItem:
			if len(list) != 0 {
				for _, c := range list {
					switch id := c.E2NodeComponentId.GetE2NodeComponentId().(type) {
					case *e2apies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1:
						nodeIdentity.CuID = &id.E2NodeComponentInterfaceTypeE1.GNbCuCpId.Value
					case *e2apies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1:
						nodeIdentity.DuID = &id.E2NodeComponentInterfaceTypeF1.GNbDuId.Value
					}
				}
			}
		}
	}

	return nodeIdentity, nil
}

func ExtractEnbID(e2NodeID *e2apies.EnbId) ([]byte, int, error) {
	var identifierBytes []byte
	var idLength int
	switch enbt := e2NodeID.EnbId.(type) {
	case *e2apies.EnbId_MacroENbId:
		identifierBytes = enbt.MacroENbId.GetValue()
		idLength = int(enbt.MacroENbId.Len)
	case *e2apies.EnbId_HomeENbId:
		identifierBytes = enbt.HomeENbId.GetValue()
		idLength = int(enbt.HomeENbId.Len)
	case *e2apies.EnbId_ShortMacroENbId:
		identifierBytes = enbt.ShortMacroENbId.GetValue()
		idLength = int(enbt.ShortMacroENbId.Len)
	case *e2apies.EnbId_LongMacroENbId:
		identifierBytes = enbt.LongMacroENbId.GetValue()
		idLength = int(enbt.LongMacroENbId.Len)
	}

	return identifierBytes, idLength, nil
}

func ExtractEnbIDchoice(e2NodeID *e2apies.EnbIdChoice) ([]byte, int, error) {
	var identifierBytes []byte
	var idLength int
	switch enbt := e2NodeID.EnbIdChoice.(type) {
	case *e2apies.EnbIdChoice_EnbIdMacro:
		identifierBytes = enbt.EnbIdMacro.GetValue()
		idLength = int(enbt.EnbIdMacro.Len)
	case *e2apies.EnbIdChoice_EnbIdShortmacro:
		identifierBytes = enbt.EnbIdShortmacro.GetValue()
		idLength = int(enbt.EnbIdShortmacro.Len)
	case *e2apies.EnbIdChoice_EnbIdLongmacro:
		identifierBytes = enbt.EnbIdLongmacro.GetValue()
		idLength = int(enbt.EnbIdLongmacro.Len)
	}

	return identifierBytes, idLength, nil
}
