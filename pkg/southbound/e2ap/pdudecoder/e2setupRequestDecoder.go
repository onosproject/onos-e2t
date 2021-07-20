// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appducontents"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"math"
)

func DecodeE2SetupRequest(request *e2appducontents.E2SetupRequest) (*types.E2NodeIdentity, *types.RanFunctions, error) {
	var nodeIdentity *types.E2NodeIdentity
	var err error

	identifierIe := request.GetProtocolIes().GetE2ApProtocolIes3()
	if identifierIe == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have id-GlobalE2node-ID")
	}
	switch e2NodeID := identifierIe.GetValue().GetGlobalE2NodeId().(type) {
	case *e2apies.GlobalE2NodeId_GNb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.GNb.GetGlobalGNbId().GetPlmnId().GetValue())
		if err != nil {
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeGNB
		choice, ok := e2NodeID.GNb.GetGlobalGNbId().GetGnbId().GetGnbIdChoice().(*e2apies.GnbIdChoice_GnbId)
		if !ok {
			return nil, nil, fmt.Errorf("expected a gNBId")
		}
		nodeIdentity.NodeIdentifier = make([]byte, 8)
		binary.BigEndian.PutUint64(nodeIdentity.NodeIdentifier, choice.GnbId.GetValue())
		nodeIdentity.NodeIDLength = int(choice.GnbId.Len)
		// TODO: investigate GNB-CU-UP-ID and GNB-DU-ID

	case *e2apies.GlobalE2NodeId_EnGNb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.EnGNb.GetGlobalGNbId().GetPLmnIdentity().GetValue())
		if err != nil {
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeEnGNB
		return nil, nil, fmt.Errorf("getting identifier of EnGnb not yet implemented")

	case *e2apies.GlobalE2NodeId_NgENb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.NgENb.GetGlobalNgENbId().GetPlmnId().GetValue())
		if err != nil {
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeNgENB
		return nil, nil, fmt.Errorf("getting identifier of ngENb not yet implemented")

	case *e2apies.GlobalE2NodeId_ENb:
		nodeIdentity, err = types.NewE2NodeIdentity(e2NodeID.ENb.GetGlobalENbId().GetPLmnIdentity().GetValue())
		if err != nil {
			return nil, nil, fmt.Errorf("error extracting node identifier")
		}
		nodeIdentity.NodeType = types.E2NodeTypeENB
		identifierBytes := make([]byte, 8)
		var lenBytes int
		switch enbt := e2NodeID.ENb.GetGlobalENbId().GetENbId().GetEnbId().(type) {
		case *e2apies.EnbId_MacroENbId:
			binary.LittleEndian.PutUint64(identifierBytes, enbt.MacroENbId.GetValue())
			lenBytes = int(math.Ceil(float64(enbt.MacroENbId.Len) / 8.0))
			nodeIdentity.NodeIDLength = int(enbt.MacroENbId.Len)
		case *e2apies.EnbId_HomeENbId:
			binary.LittleEndian.PutUint64(identifierBytes, enbt.HomeENbId.GetValue())
			lenBytes = int(math.Ceil(float64(enbt.HomeENbId.Len) / 8.0))
			nodeIdentity.NodeIDLength = int(enbt.HomeENbId.Len)
		case *e2apies.EnbId_ShortMacroENbId:
			binary.LittleEndian.PutUint64(identifierBytes, enbt.ShortMacroENbId.GetValue())
			lenBytes = int(math.Ceil(float64(enbt.ShortMacroENbId.Len) / 8.0))
			nodeIdentity.NodeIDLength = int(enbt.ShortMacroENbId.Len)
		case *e2apies.EnbId_LongMacroENbId:
			binary.LittleEndian.PutUint64(identifierBytes, enbt.LongMacroENbId.GetValue())
			lenBytes = int(math.Ceil(float64(enbt.LongMacroENbId.Len) / 8.0))
			nodeIdentity.NodeIDLength = int(enbt.LongMacroENbId.Len)
		}
		nodeIdentity.NodeIdentifier = make([]byte, lenBytes)
		copy(nodeIdentity.NodeIdentifier, identifierBytes[:lenBytes])
	}

	ranFunctionsList := make(types.RanFunctions)
	ranFunctionsIe := request.GetProtocolIes().GetE2ApProtocolIes10()
	if ranFunctionsIe == nil {
		return nodeIdentity, nil, fmt.Errorf("error E2APpdu does not have id-RANfunctionsAdded")
	}

	for _, rfIe := range ranFunctionsIe.GetValue().GetValue() {
		rfItem := rfIe.GetE2ApProtocolIes10().GetValue()
		ranFunctionsList[types.RanFunctionID(rfItem.GetRanFunctionId().GetValue())] = types.RanFunctionItem{
			Description: types.RanFunctionDescription(string(rfItem.GetRanFunctionDefinition().GetValue())),
			Revision:    types.RanFunctionRevision(rfItem.GetRanFunctionRevision().GetValue()),
		}
	}

	return nodeIdentity, &ranFunctionsList, nil
}

func DecodeE2SetupRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*types.E2NodeIdentity, *types.RanFunctions, error) {
	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2setup := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup()
	if e2setup == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}
	return DecodeE2SetupRequest(e2setup.GetInitiatingMessage())
}
