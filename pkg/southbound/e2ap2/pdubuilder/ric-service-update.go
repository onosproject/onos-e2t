// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
)

func CreateRicServiceUpdateE2apPdu(rfal types.RanFunctions, rfDeleted types.RanFunctionRevisions, rfml types.RanFunctions) (*e2appdudescriptions.E2ApPdu, error) {

	if rfal == nil && rfDeleted == nil && rfml == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						InitiatingMessage: &e2appducontents.RicserviceUpdate{
							ProtocolIes: &e2appducontents.RicserviceUpdateIes{
								//E2ApProtocolIes10: &ranFunctionsAddedList,    //RAN functions (added) List
								//E2ApProtocolIes11: &ranFunctionsDeletedList,  //RAN functions ID (deleted) List
								//E2ApProtocolIes12: &ranFunctionsModifiedList, //RAN functions (modified) List
							},
						},
						ProcedureCode: &e2ap_constants.IdRicserviceUpdate{
							Value: int32(v1beta2.ProcedureCodeIDRICserviceUpdate),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}

	if rfal != nil {
		ranFunctionsAddedList := e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes10{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionsAdded),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			RanFunctionsAddedList: &e2appducontents.RanfunctionsList{
				Value: make([]*e2appducontents.RanfunctionItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for id, ranFunctionID := range rfal {
			ranFunction := e2appducontents.RanfunctionItemIes{
				E2ApProtocolIes10: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
					Id:          int32(v1beta2.ProtocolIeIDRanfunctionItem),
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
							Value: ranFunctionID.OID,
						},
					},
				},
			}
			ranFunctionsAddedList.RanFunctionsAddedList.Value = append(ranFunctionsAddedList.RanFunctionsAddedList.Value, &ranFunction)
		}
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes10 = &ranFunctionsAddedList
	}

	if rfDeleted != nil {
		ranFunctionsDeletedList := e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes11{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionsDeleted),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			RanFunctionsDeletedList: &e2appducontents.RanfunctionsIdList{
				Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for rfID, rfRevision := range rfDeleted {
			rfIDiIe := e2appducontents.RanfunctionIdItemIes{
				RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
					Id:          int32(v1beta2.ProtocolIeIDRanfunctionIDItem),
					Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
					Value: &e2appducontents.RanfunctionIdItem{
						RanFunctionId: &e2apies.RanfunctionId{
							Value: int32(rfID),
						},
						RanFunctionRevision: &e2apies.RanfunctionRevision{
							Value: int32(rfRevision),
						},
					},
					Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
				},
			}
			ranFunctionsDeletedList.RanFunctionsDeletedList.Value = append(ranFunctionsDeletedList.RanFunctionsDeletedList.Value, &rfIDiIe)
		}
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes11 = &ranFunctionsDeletedList
	}

	if rfml != nil {
		ranFunctionsModifiedList := e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes12{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionsModified),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			RanFunctionsModifiedList: &e2appducontents.RanfunctionsList{
				Value: make([]*e2appducontents.RanfunctionItemIes, 0),
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
		}

		for id, ranFunctionID := range rfml {
			ranFunction := e2appducontents.RanfunctionItemIes{
				E2ApProtocolIes10: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
					Id:          int32(v1beta2.ProtocolIeIDRanfunctionItem),
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
							Value: ranFunctionID.OID,
						},
					},
				},
			}
			ranFunctionsModifiedList.RanFunctionsModifiedList.Value = append(ranFunctionsModifiedList.RanFunctionsModifiedList.Value, &ranFunction)
		}
		e2apPdu.GetInitiatingMessage().GetProcedureCode().GetRicServiceUpdate().GetInitiatingMessage().GetProtocolIes().E2ApProtocolIes12 = &ranFunctionsModifiedList
	}

	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
