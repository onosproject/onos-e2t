// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func DecodeE2SetupRequestPdu(e2apPdu *e2appdudescriptions.E2ApPdu) (*types.E2NodeIdentity, *types.RanFunctions, error) {
	var nodeIdentity *types.E2NodeIdentity
	var err error

	if err := e2apPdu.Validate(); err != nil {
		return nil, nil, fmt.Errorf("invalid E2APpdu %s", err.Error())
	}

	e2setup := e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup()
	if e2setup == nil {
		return nil, nil, fmt.Errorf("error E2APpdu does not have E2Setup")
	}

	identifierIe := e2setup.GetInitiatingMessage().GetProtocolIes().GetE2ApProtocolIes3()
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
		binary.LittleEndian.PutUint64(nodeIdentity.NodeIdentifier, choice.GnbId.GetValue())
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
		return nil, nil, fmt.Errorf("getting identifier of eNB not yet implemented")

	}

	ranFunctionsList := make(types.RanFunctions)

	return nodeIdentity, &ranFunctionsList, nil
}
