// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"fmt"
	"math"

	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
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
		//nodeIdentity.NodeIdentifier = make([]byte, 0)
		//ToDo - this approach should be fine
		nodeIdentity.NodeIdentifier = choice.GnbId.GetValue()
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
		//identifierBytes := make([]byte, 0)
		var identifierBytes []byte
		var lenBytes int
		var idLength int
		switch enbt := e2NodeID.ENb.GetGlobalENbId().GetENbId().GetEnbId().(type) {
		case *e2apies.EnbId_MacroENbId:
			identifierBytes = enbt.MacroENbId.GetValue()
			lenBytes = int(math.Ceil(float64(enbt.MacroENbId.Len) / 8.0))
			idLength = int(enbt.MacroENbId.Len)
		case *e2apies.EnbId_HomeENbId:
			identifierBytes = enbt.HomeENbId.GetValue()
			lenBytes = int(math.Ceil(float64(enbt.HomeENbId.Len) / 8.0))
			idLength = int(enbt.HomeENbId.Len)
		case *e2apies.EnbId_ShortMacroENbId:
			identifierBytes = enbt.ShortMacroENbId.GetValue()
			lenBytes = int(math.Ceil(float64(enbt.ShortMacroENbId.Len) / 8.0))
			idLength = int(enbt.ShortMacroENbId.Len)
		case *e2apies.EnbId_LongMacroENbId:
			identifierBytes = enbt.LongMacroENbId.GetValue()
			lenBytes = int(math.Ceil(float64(enbt.LongMacroENbId.Len) / 8.0))
			idLength = int(enbt.LongMacroENbId.Len)
		}
		nodeIdentity.NodeIdentifier = make([]byte, lenBytes)
		copy(nodeIdentity.NodeIdentifier, identifierBytes[:lenBytes])
		nodeIdentity.NodeIDLength = idLength
		//ToDo - couldn't it be just this?
		//nodeIdentity.NodeIdentifier = identifierBytes
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
			OID:         types.RanFunctionOID(string(rfItem.GetRanFunctionOid().GetValue())),
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

func GetE2NodeID(nodeID []byte, length int) string {
	unusedBits := 8 - length%8
	var result uint64 = 0
	for i, b := range nodeID {
		result += uint64(b) << ((len(nodeID) - i - 1) * 8)
	}
	result = result >> unusedBits
	return fmt.Sprintf("%x", result)
}
