// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2NodeConfigurationUpdateE2apPdu(trID int32) (*e2appdudescriptions.E2ApPdu, error) {

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: int32(v2.ProcedureCodeIDE2nodeConfigurationUpdate),
				Criticality:   e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
				Value: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures{
					ImValues: &e2appdudescriptions.InitiatingMessageE2ApElementaryProcedures_E2NodeConfigurationUpdate{
						E2NodeConfigurationUpdate: &e2appducontents.E2NodeConfigurationUpdate{
							ProtocolIes: make([]*e2appducontents.E2NodeConfigurationUpdateIes, 0),
						},
					},
				},
			},
		},
	}

	e2apPdu.GetInitiatingMessage().GetValue().GetE2NodeConfigurationUpdate().SetTransactionID(trID)

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}

func CreateE2NodeComponentIDNg(value string) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeNg{
			E2NodeComponentInterfaceTypeNg: &e2ap_ies.E2NodeComponentInterfaceNg{
				AmfName: &e2ap_commondatatypes.Amfname{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentIDXn(glNgRanNodeID *e2ap_ies.GlobalNgRannodeId) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeXn{
			E2NodeComponentInterfaceTypeXn: &e2ap_ies.E2NodeComponentInterfaceXn{
				GlobalNgRanNodeId: glNgRanNodeID,
			},
		},
	}
}

func CreateE2NodeComponentIDE1(value int64) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeE1{
			E2NodeComponentInterfaceTypeE1: &e2ap_ies.E2NodeComponentInterfaceE1{
				GNbCuCpId: &e2ap_ies.GnbCuUpId{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentIDF1(value int64) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeF1{
			E2NodeComponentInterfaceTypeF1: &e2ap_ies.E2NodeComponentInterfaceF1{
				GNbDuId: &e2ap_ies.GnbDuId{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentIDW1(value int64) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeW1{
			E2NodeComponentInterfaceTypeW1: &e2ap_ies.E2NodeComponentInterfaceW1{
				NgENbDuId: &e2ap_ies.NgenbDuId{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentIDS1(value string) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeS1{
			E2NodeComponentInterfaceTypeS1: &e2ap_ies.E2NodeComponentInterfaceS1{
				MmeName: &e2ap_commondatatypes.Mmename{
					Value: value,
				},
			},
		},
	}
}

func CreateE2NodeComponentIDX2(glEnbID *e2ap_ies.GlobalEnbId, glEnGnbID *e2ap_ies.GlobalenGnbId) *e2ap_ies.E2NodeComponentId {
	return &e2ap_ies.E2NodeComponentId{
		E2NodeComponentId: &e2ap_ies.E2NodeComponentId_E2NodeComponentInterfaceTypeX2{
			E2NodeComponentInterfaceTypeX2: &e2ap_ies.E2NodeComponentInterfaceX2{
				GlobalENbId:   glEnbID,
				GlobalEnGNbId: glEnGnbID,
			},
		},
	}
}

func CreateGlobalNgRanNodeIDGnb(plmnID []byte, bs *asn1.BitString) (*e2ap_ies.GlobalNgRannodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("CreateGlobalNgRanNodeIDGnb(): PlmnID should be 3 bytes long, have %d", len(plmnID))
	}

	if bs.Len < 22 || bs.Len > 32 {
		return nil, fmt.Errorf("CreateGlobalNgRanNodeIDGnb(): BitString should be of length 22 to 32 bits, have %d", bs.Len)
	}

	if len(bs.Value) < 3 || len(bs.Value) > 4 {
		return nil, fmt.Errorf("CreateGlobalNgRanNodeIDGnb(): BitString should be of length 22 to 32 bits, have %d", 8*len(bs.Value))
	}

	return &e2ap_ies.GlobalNgRannodeId{
		GlobalNgRannodeId: &e2ap_ies.GlobalNgRannodeId_GNb{
			GNb: &e2ap_ies.GlobalgNbId{
				PlmnId: &e2ap_commondatatypes.PlmnIdentity{
					Value: plmnID,
				},
				GnbId: &e2ap_ies.GnbIdChoice{
					GnbIdChoice: &e2ap_ies.GnbIdChoice_GnbId{
						GnbId: bs,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalNgRanNodeIDNgEnb(plmnID []byte, enbID *e2ap_ies.EnbIdChoice) (*e2ap_ies.GlobalNgRannodeId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("CreateGlobalNgRanNodeIDNgEnb(): PlmnID should be 3 bytes long, have %d", len(plmnID))
	}

	return &e2ap_ies.GlobalNgRannodeId{
		GlobalNgRannodeId: &e2ap_ies.GlobalNgRannodeId_NgENb{
			NgENb: &e2ap_ies.GlobalngeNbId{
				PlmnId: &e2ap_commondatatypes.PlmnIdentity{
					Value: plmnID,
				},
				EnbId: enbID,
			},
		},
	}, nil
}

func CreateGlobalEnbID(plmnID []byte, enbID *e2ap_ies.EnbId) (*e2ap_ies.GlobalEnbId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 bytes long, have %d", len(plmnID))
	}

	return &e2ap_ies.GlobalEnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: plmnID,
		},
		ENbId: enbID,
	}, nil
}

func CreateGlobalEnGnbID(plmnID []byte, bs *asn1.BitString) (*e2ap_ies.GlobalenGnbId, error) {

	if len(plmnID) != 3 {
		return nil, fmt.Errorf("PlmnID should be 3 bytes long, have %d", len(plmnID))
	}

	if bs.Len < 22 || bs.Len > 32 {
		return nil, fmt.Errorf("CreateGlobalNgRanNodeIDGnb(): BitString should be of length 22 to 32 bits, have %d", bs.Len)
	}

	if len(bs.Value) < 3 || len(bs.Value) > 4 {
		return nil, fmt.Errorf("CreateGlobalNgRanNodeIDGnb(): BitString should be of length 22 to 32 bits, have %d", 8*len(bs.Value))
	}

	return &e2ap_ies.GlobalenGnbId{
		PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
			Value: plmnID,
		},
		GNbId: &e2ap_ies.EngnbId{
			EngnbId: &e2ap_ies.EngnbId_GNbId{
				GNbId: bs,
			},
		},
	}, nil
}
