// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
)

func CreateRicServiceUpdateE2apPdu() (*e2appdudescriptions.E2ApPdu, error) {

	ranFunctionsAddedList := e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes10{
		Id:          int32(v1beta2.ProcedureCodeIDE2nodeConfigurationUpdate),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		RanFunctionsAddedList: &e2appducontents.RanfunctionsList{
			Value: make([]*e2appducontents.RanfunctionItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfi := &e2appducontents.RanfunctionItemIes{
		E2ApProtocolIes10: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RanfunctionItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 123,
				},
				RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
					Value: []byte{0x01, 0x02, 0x03},
				},
				RanFunctionRevision: &e2apies.RanfunctionRevision{
					Value: 1,
				},
				RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
					Value: []byte{0x01, 0x02, 0x03},
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	ranFunctionsAddedList.RanFunctionsAddedList.Value = append(ranFunctionsAddedList.RanFunctionsAddedList.Value, rfi)

	ranFunctionsDeletedList := e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes11{
		Id:          int32(v1beta2.ProcedureCodeIDE2connectionUpdate),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		RanFunctionsDeletedList: &e2appducontents.RanfunctionsIdList{
			Value: make([]*e2appducontents.RanfunctionIdItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfdi := &e2appducontents.RanfunctionIdItemIes{
		RanFunctionIdItemIes6: &e2appducontents.RanfunctionIdItemIes_RanfunctionIdItemIes6{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionIDItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RanfunctionIdItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 123,
				},
				RanFunctionRevision: &e2apies.RanfunctionRevision{
					Value: 1,
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	ranFunctionsDeletedList.RanFunctionsDeletedList.Value = append(ranFunctionsDeletedList.RanFunctionsDeletedList.Value, rfdi)

	ranFunctionsModifiedList := e2appducontents.RicserviceUpdateIes_RicserviceUpdateIes12{
		Id:          int32(v1beta2.ProtocolIeIDRanfunctionsModified),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		RanFunctionsModifiedList: &e2appducontents.RanfunctionsList{
			Value: make([]*e2appducontents.RanfunctionItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	rfmi := &e2appducontents.RanfunctionItemIes{
		E2ApProtocolIes10: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
			Id:          int32(v1beta2.ProtocolIeIDRanfunctionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RanfunctionItem{
				RanFunctionId: &e2apies.RanfunctionId{
					Value: 123,
				},
				RanFunctionDefinition: &e2ap_commondatatypes.RanfunctionDefinition{
					Value: []byte{0x01, 0x02, 0x03},
				},
				RanFunctionRevision: &e2apies.RanfunctionRevision{
					Value: 1,
				},
				RanFunctionOid: &e2ap_commondatatypes.RanfunctionOid{
					Value: []byte{0x01, 0x02, 0x03},
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		},
	}
	ranFunctionsModifiedList.RanFunctionsModifiedList.Value = append(ranFunctionsModifiedList.RanFunctionsModifiedList.Value, rfmi)

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_InitiatingMessage{
			InitiatingMessage: &e2appdudescriptions.InitiatingMessage{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					RicServiceUpdate: &e2appdudescriptions.RicServiceUpdate{
						InitiatingMessage: &e2appducontents.RicserviceUpdate{
							ProtocolIes: &e2appducontents.RicserviceUpdateIes{
								E2ApProtocolIes10: &ranFunctionsAddedList,    //RAN functions (added) List
								E2ApProtocolIes11: &ranFunctionsDeletedList,  //RAN functions ID (deleted) List
								E2ApProtocolIes12: &ranFunctionsModifiedList, //RAN functions (modified) List
							},
						},
					},
				},
			},
		},
	}
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}
