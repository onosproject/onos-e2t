// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0
package pdubuilder

import (
	"fmt"

	"github.com/onosproject/onos-e2t/api/e2ap/v2beta1"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v2beta1/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
)

func CreateE2SetupRequestPdu(trID int32, ge2nID *e2apies.GlobalE2NodeId, ranFunctionIds types.RanFunctions,
	e2nccul []*types.E2NodeComponentConfigUpdateItem) (*e2appdudescriptions.E2ApPdu, error) {

	gnbIDIe := e2appducontents.E2SetupRequestIes_E2SetupRequestIes3{
		Id:          int32(v2beta1.ProtocolIeIDGlobalE2nodeID),
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ge2nID,
	}

	if e2nccul == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	configUpdateList := e2appducontents.E2NodeComponentConfigUpdateList{
		Value: make([]*e2appducontents.E2NodeComponentConfigUpdateItemIes, 0),
	}

	for _, e2nccui := range e2nccul {
		cui := &e2appducontents.E2NodeComponentConfigUpdateItemIes{
			Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdateItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2NodeComponentConfigUpdateItem{
				E2NodeComponentType:         e2nccui.E2NodeComponentType,
				E2NodeComponentId:           e2nccui.E2NodeComponentID,
				E2NodeComponentConfigUpdate: &e2nccui.E2NodeComponentConfigUpdate,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		configUpdateList.Value = append(configUpdateList.Value, cui)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					E2Setup: &e2appdudescriptions.E2Setup{
						InitiatingMessage: &e2appducontents.E2SetupRequest{
							ProtocolIes: &e2appducontents.E2SetupRequestIes{
								E2ApProtocolIes3: &gnbIDIe,
								//E2ApProtocolIes10: &ranFunctions,
								E2ApProtocolIes33: &e2appducontents.E2SetupRequestIes_E2SetupRequestIes33{
									Id:          int32(v2beta1.ProtocolIeIDE2nodeComponentConfigUpdate),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value:       &configUpdateList,
									Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
								E2ApProtocolIes49: &e2appducontents.E2SetupRequestIes_E2SetupRequestIes49{
									Id:          int32(v2beta1.ProtocolIeIDTransactionID),
									Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
									Value: &e2apies.TransactionId{
										Value: trID,
									},
									Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
								},
							},
						},
						ProcedureCode: &e2ap_constants.IdE2Setup{
							Value: int32(v2beta1.ProcedureCodeIDE2setup),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	ranFunctions := e2appducontents.E2SetupRequestIes_E2SetupRequestIes10{
		Id:          int32(v2beta1.ProtocolIeIDRanfunctionsAdded),
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsList{
			Value: make([]*e2appducontents.RanfunctionItemIes, 0),
		},
	}

	for id, ranFunctionID := range ranFunctionIds {
		ranFunction := e2appducontents.RanfunctionItemIes{
			E2ApProtocolIes10: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
				Id:          int32(v2beta1.ProtocolIeIDRanfunctionItem),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
				Value: &e2appducontents.RanfunctionItem{
					RanFunctionId: &e2apies.RanfunctionId{
						Value: int32(id),
					},
					RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
						Value: []byte(ranFunctionID.Description),
					},
					RanFunctionRevision: &e2apies.RanfunctionRevision{
						Value: int32(ranFunctionID.Revision),
					},
					RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
						Value: string(ranFunctionID.OID),
					},
				},
			},
		}
		ranFunctions.Value.Value = append(ranFunctions.Value.Value, &ranFunction)
	}
	e2apPdu.GetInitiatingMessage().GetProcedureCode().GetE2Setup().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes10 = &ranFunctions

	//if err := e2apPdu.Validate(); err != nil {
	//	return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	//}
	return &e2apPdu, nil
}

func CreateGnbIDchoice(value []byte, len uint32) (*e2apies.GnbIdChoice, error) {

	if len < 22 || len > 32 {
		return nil, fmt.Errorf("length should be in range 22 to 32 bits")
	}

	return &e2apies.GnbIdChoice{
		GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
			GnbId: &asn1.BitString{
				Value: value,
				Len:   len,
			},
		},
	}, nil
}

func CreateGlobalE2nodeIDGnb(plmnID types.PlmnID, bs *asn1.BitString) (*e2apies.GlobalE2NodeId, error) {

	if bs.GetLen() < 22 && bs.GetLen() > 32 {
		return nil, fmt.Errorf("expecting GNbID length in range 22 to 32 bits, got %d", bs.GetLen())
	}

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_GNb{
			GNb: &e2apies.GlobalE2NodeGnbId{
				GlobalGNbId: &e2apies.GlobalgNbId{
					PlmnId: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
					},
					GnbId: &e2apies.GnbIdChoice{
						GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
							GnbId: bs,
						},
					},
				},
			},
		},
	}, nil
}

func CreateGlobalE2nodeIDEnGnb(plmnID types.PlmnID, enGnbID *asn1.BitString) (*e2apies.GlobalE2NodeId, error) {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_EnGNb{
			EnGNb: &e2apies.GlobalE2NodeEnGnbId{
				GlobalGNbId: &e2apies.GlobalenGnbId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
					},
					GNbId: &e2apies.EngnbId{
						EngnbId: &e2apies.EngnbId_GNbId{
							GNbId: enGnbID,
						},
					},
				},
			},
		},
	}, nil
}

func CreateGlobalE2nodeIDNgEnb(plmnID types.PlmnID, enbIDchoice *e2apies.EnbIdChoice) (*e2apies.GlobalE2NodeId, error) {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_NgENb{
			NgENb: &e2apies.GlobalE2NodeNgEnbId{
				GlobalNgENbId: &e2apies.GlobalngeNbId{
					PlmnId: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
					},
					EnbId: enbIDchoice,
				},
			},
		},
	}, nil
}

func CreateGlobalE2nodeIDEnb(plmnID types.PlmnID, enbID *e2apies.EnbId) (*e2apies.GlobalE2NodeId, error) {

	return &e2apies.GlobalE2NodeId{
		GlobalE2NodeId: &e2apies.GlobalE2NodeId_ENb{
			ENb: &e2apies.GlobalE2NodeEnbId{
				GlobalENbId: &e2apies.GlobalEnbId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
					},
					ENbId: enbID,
				},
			},
		},
	}, nil
}

func CreateEnbIDChoiceMacro(bs *asn1.BitString) (*e2apies.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2apies.EnbIdChoice{
		EnbIdChoice: &e2apies.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: bs,
		},
	}, nil
}

func CreateEnbIDChoiceShortMacro(bs *asn1.BitString) (*e2apies.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x3f > 0 {
		return nil, fmt.Errorf("expected last 6 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2apies.EnbIdChoice{
		EnbIdChoice: &e2apies.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: bs,
		},
	}, nil
}

func CreateEnbIDChoiceLongMacro(bs *asn1.BitString) (*e2apies.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x07 > 0 {
		return nil, fmt.Errorf("expected last 3 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}

	return &e2apies.EnbIdChoice{
		EnbIdChoice: &e2apies.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: bs,
		},
	}, nil
}

func CreateEnbIDMacro(bs *asn1.BitString) (*e2apies.EnbId, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2apies.EnbId{
		EnbId: &e2apies.EnbId_MacroENbId{
			MacroENbId: bs,
		},
	}, nil
}

func CreateEnbIDHome(bs *asn1.BitString) (*e2apies.EnbId, error) {

	if len(bs.GetValue()) != 4 {
		return nil, fmt.Errorf("expecting length to be exactly 4 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[4]&0x0f > 0 {
		return nil, fmt.Errorf("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2apies.EnbId{
		EnbId: &e2apies.EnbId_HomeENbId{
			HomeENbId: bs,
		},
	}, nil
}

func CreateEnbIDShortMacro(bs *asn1.BitString) (*e2apies.EnbId, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x3f > 0 {
		return nil, fmt.Errorf("expected last 6 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2apies.EnbId{
		EnbId: &e2apies.EnbId_ShortMacroENbId{
			ShortMacroENbId: bs,
		},
	}, nil
}

func CreateEnbIDLongMacro(bs *asn1.BitString) (*e2apies.EnbId, error) {

	if len(bs.GetValue()) != 3 {
		return nil, fmt.Errorf("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x07 > 0 {
		return nil, fmt.Errorf("expected last 3 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}

	return &e2apies.EnbId{
		EnbId: &e2apies.EnbId_LongMacroENbId{
			LongMacroENbId: bs,
		},
	}, nil
}
