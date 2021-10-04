// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package e2ap

import (
	"context"
	"encoding/binary"
	"io"
	"net"
	"testing"
	"time"

	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	"github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
)

func TestConns(t *testing.T) {
	testIP := net.ParseIP("127.0.0.1")
	clientCh := make(chan []byte)
	serverCh := make(chan []byte)
	clientConn := &testConn{
		readCh:  serverCh,
		writeCh: clientCh,
	}
	serverConn := &testConn{
		readCh:  clientCh,
		writeCh: serverCh,
	}

	e2NodeCh := NewClientConn(clientConn, func(conn ClientConn) ClientInterface {
		return &testClientProcedures{}
	})
	ricCh := NewServerConn(serverConn, func(conn ServerConn) ServerInterface {
		return &testServerProcedures{}
	})

	ge2nID, err := pdubuilder.CreateGlobalE2nodeIDGnb([3]byte{0x4F, 0x4E, 0x46}, &asn1.BitString{
		Value: []byte{0x00, 0x00, 0x04},
		Len:   22,
	})
	assert.NoError(t, err)
	gnbIDIe := &e2appducontents.E2SetupRequestIes_E2SetupRequestIes3{
		Id:          int32(v2.ProtocolIeIDGlobalE2nodeID),
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ge2nID,
	}

	ranFunctionList := make(types.RanFunctions)
	ranFunctionList[100] = types.RanFunctionItem{
		Description: []byte("Type 1"),
		Revision:    1,
		OID:         "oid1",
	}

	ranFunctionList[200] = types.RanFunctionItem{
		Description: []byte("Type 2"),
		Revision:    2,
		OID:         "oid2",
	}

	ranFunctions := &e2appducontents.E2SetupRequestIes_E2SetupRequestIes10{
		Id:          int32(v2.ProtocolIeIDRanfunctionsAdded),
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.RanfunctionsList{
			Value: make([]*e2appducontents.RanfunctionItemIes, 0),
		},
	}

	for id, ranFunctionID := range ranFunctionList {
		ranFunction := e2appducontents.RanfunctionItemIes{
			E2ApProtocolIes8: &e2appducontents.RanfunctionItemIes_RanfunctionItemIes8{
				Id:          int32(v2.ProtocolIeIDRanfunctionItem),
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

	configComponentAdditionItems := []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
			E2NodeComponentID: pdubuilder.CreateE2NodeComponentIDE1(21),
			E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
				E2NodeComponentRequestPart:  []byte{0x00, 0x01, 0x02},
				E2NodeComponentResponsePart: []byte{0x03, 0x04, 0x05},
			},
		},
	}

	configUpdateAdditionList := e2appducontents.E2NodeComponentConfigAdditionList{
		Value: make([]*e2appducontents.E2NodeComponentConfigAdditionItemIes, 0),
	}
	for _, configAdditionItem := range configComponentAdditionItems {
		cui := &e2appducontents.E2NodeComponentConfigAdditionItemIes{
			Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAdditionItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
			Value: &e2appducontents.E2NodeComponentConfigAdditionItem{
				E2NodeComponentInterfaceType: configAdditionItem.E2NodeComponentType,
				E2NodeComponentId:            configAdditionItem.E2NodeComponentID,
				E2NodeComponentConfiguration: &configAdditionItem.E2NodeComponentConfiguration,
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		configUpdateAdditionList.Value = append(configUpdateAdditionList.Value, cui)
	}

	e2SetupRequest := &e2appducontents.E2SetupRequest{
		ProtocolIes: &e2appducontents.E2SetupRequestIes{
			E2ApProtocolIes3:  gnbIDIe,
			E2ApProtocolIes10: ranFunctions,
			E2ApProtocolIes50: &e2appducontents.E2SetupRequestIes_E2SetupRequestIes50{
				Id:          int32(v2.ProtocolIeIDE2nodeComponentConfigAddition),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       &configUpdateAdditionList,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes49: &e2appducontents.E2SetupRequestIes_E2SetupRequestIes49{
				Id:          int32(v2.ProtocolIeIDTransactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.TransactionId{
					Value: 2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}

	e2SetupResponse, e2SetupFailure, err := e2NodeCh.E2Setup(context.TODO(), e2SetupRequest)
	assert.NotNil(t, e2SetupResponse)
	assert.Nil(t, e2SetupFailure)
	assert.NoError(t, err)

	ricSubscriptionRequest := &e2appducontents.RicsubscriptionRequest{
		ProtocolIes: &e2appducontents.RicsubscriptionRequestIes{
			E2ApProtocolIes5: &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes5{
				Id:          int32(v2.ProtocolIeIDRanfunctionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RanfunctionId{
					Value: 1,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes29: &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
				Id:          int32(v2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicrequestId{
					RicRequestorId: 1,
					RicInstanceId:  2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes30: &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes30{
				Id:          int32(v2.ProtocolIeIDRicsubscriptionDetails),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2appducontents.RicsubscriptionDetails{
					RicEventTriggerDefinition: &e2ap_commondatatypes.RiceventTriggerDefinition{},
					RicActionToBeSetupList: &e2appducontents.RicactionsToBeSetupList{
						Value: []*e2appducontents.RicactionToBeSetupItemIes{
							{
								Id:          int32(v2.ProtocolIeIDRicactionToBeSetupItem),
								Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
								Value: &e2appducontents.RicactionToBeSetupItem{
									RicActionId: &e2apies.RicactionId{
										Value: int32(1),
									},
									RicActionType:       e2apies.RicactionType_RICACTION_TYPE_REPORT,
									RicActionDefinition: &e2ap_commondatatypes.RicactionDefinition{},
									RicSubsequentAction: &e2apies.RicsubsequentAction{},
								},
								Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
							},
						},
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}
	ricSubscriptionResponse, ricSubscriptionFailure, err := ricCh.RICSubscription(context.TODO(), ricSubscriptionRequest)
	assert.NotNil(t, ricSubscriptionResponse)
	assert.Nil(t, ricSubscriptionFailure)
	assert.NoError(t, err)

	ricIndication := &e2appducontents.Ricindication{
		ProtocolIes: &e2appducontents.RicindicationIes{
			E2ApProtocolIes29: &e2appducontents.RicindicationIes_RicindicationIes29{
				Id:          int32(v2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicrequestId{
					RicRequestorId: 1,
					RicInstanceId:  2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes5: &e2appducontents.RicindicationIes_RicindicationIes5{
				Id:          int32(v2.ProtocolIeIDRanfunctionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RanfunctionId{
					Value: 1,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes15: &e2appducontents.RicindicationIes_RicindicationIes15{
				Id:          int32(v2.ProtocolIeIDRicactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicactionId{
					Value: 2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes27: &e2appducontents.RicindicationIes_RicindicationIes27{
				Id:          int32(v2.ProtocolIeIDRicindicationSn),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicindicationSn{
					Value: 3,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			},
			E2ApProtocolIes28: &e2appducontents.RicindicationIes_RicindicationIes28{
				Id:          int32(v2.ProtocolIeIDRicindicationType),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       e2apies.RicindicationType_RICINDICATION_TYPE_REPORT,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes25: &e2appducontents.RicindicationIes_RicindicationIes25{
				Id:          int32(v2.ProtocolIeIDRicindicationHeader),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2ap_commondatatypes.RicindicationHeader{
					Value: []byte("foo"),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes26: &e2appducontents.RicindicationIes_RicindicationIes26{
				Id:          int32(v2.ProtocolIeIDRicindicationMessage),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2ap_commondatatypes.RicindicationMessage{
					Value: []byte("bar"),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes20: &e2appducontents.RicindicationIes_RicindicationIes20{
				Id:          int32(v2.ProtocolIeIDRiccallProcessID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2ap_commondatatypes.RiccallProcessId{
					Value: []byte("baz"),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			},
		},
	}
	err = e2NodeCh.RICIndication(context.TODO(), ricIndication)
	assert.NoError(t, err)
	time.Sleep(time.Second)

	connectionAddList := &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes44{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateAdd),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2appducontents.E2ConnectionUpdateList{
			Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	// Test E2 connection update
	portBytes := make([]byte, 2)
	port := uint16(36421)
	binary.BigEndian.PutUint16(portBytes, port)
	cai := &e2appducontents.E2ConnectionUpdateItemIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2ConnectionUpdateItem{
			TnlInformation: &e2apies.Tnlinformation{
				TnlPort: &asn1.BitString{
					Value: portBytes,
					Len:   16,
				},
				TnlAddress: &asn1.BitString{
					Value: testIP,
					Len:   128,
				},
			},
			TnlUsage: e2apies.Tnlusage_TNLUSAGE_BOTH,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}

	transactionID := &e2appducontents.E2ConnectionUpdateIes_E2ConnectionUpdateIes49{
		Id:          int32(v2.ProtocolIeIDTransactionID),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value: &e2apies.TransactionId{
			Value: 3,
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}
	connectionAddList.Value.Value = append(connectionAddList.Value.Value, cai)
	connectionUpdateRequest := &e2appducontents.E2ConnectionUpdate{
		ProtocolIes: &e2appducontents.E2ConnectionUpdateIes{
			E2ApProtocolIes44: connectionAddList,
			E2ApProtocolIes49: transactionID,
		},
	}

	ack, failure, err := ricCh.E2ConnectionUpdate(context.TODO(), connectionUpdateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, ack)
	assert.Nil(t, failure)
	tnlInformation := ack.GetProtocolIes().GetE2ApProtocolIes39().Value.GetValue()[0].GetValue()
	assert.Equal(t, testIP, net.IP(tnlInformation.GetTnlInformation().GetTnlAddress().GetValue()))
	assert.Equal(t, port, binary.BigEndian.Uint16(tnlInformation.GetTnlInformation().GetTnlPort().GetValue()))

	configUpdateNodeID := &e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes3{
		Id:          int32(v2.ProtocolIeIDGlobalE2nodeID),
		Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
		Value:       ge2nID,
	}
	// Test E2 configuration update
	configUpdate := &e2appducontents.E2NodeConfigurationUpdate{
		ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateIes{
			E2ApProtocolIes49: &e2appducontents.E2NodeConfigurationUpdateIes_E2NodeConfigurationUpdateIes49{
				Id:          int32(v2.ProtocolIeIDTransactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.TransactionId{
					Value: 2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes3: configUpdateNodeID,
		},
	}

	configUpdateAck, configUpdateFailure, err := e2NodeCh.E2ConfigurationUpdate(context.TODO(), configUpdate)
	assert.NoError(t, err)
	assert.NotNil(t, configUpdateAck)
	assert.Nil(t, configUpdateFailure)

}

type testConn struct {
	net.Conn
	readCh  chan []byte
	writeCh chan []byte
}

func (c *testConn) Read(b []byte) (n int, err error) {
	bytes, ok := <-c.readCh
	if !ok {
		return 0, io.EOF
	}
	copy(b, bytes)
	return len(bytes), nil
}

func (c *testConn) Write(b []byte) (n int, err error) {
	c.writeCh <- b
	return len(b), nil
}

func (c *testConn) Close() error {
	close(c.writeCh)
	return nil
}

type testClientProcedures struct{}

func (p *testClientProcedures) RICControl(ctx context.Context, request *e2appducontents.RiccontrolRequest) (response *e2appducontents.RiccontrolAcknowledge, failure *e2appducontents.RiccontrolFailure, err error) {
	return &e2appducontents.RiccontrolAcknowledge{
		ProtocolIes: &e2appducontents.RiccontrolAcknowledgeIes{
			E2ApProtocolIes29: &e2appducontents.RiccontrolAcknowledgeIes_RiccontrolAcknowledgeIes29{
				Id:          int32(v2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       request.ProtocolIes.E2ApProtocolIes29.Value,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

func (p *testClientProcedures) RICSubscription(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (response *e2appducontents.RicsubscriptionResponse, failure *e2appducontents.RicsubscriptionFailure, err error) {
	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	ricActionsAccepted := []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20}
	res := make([]*e2appducontents.RicactionAdmittedItemIes, 0)
	for _, raa := range ricActionsAccepted {
		raaIe := &e2appducontents.RicactionAdmittedItemIes{
			Id:          int32(v2.ProtocolIeIDRicactionAdmittedItem),
			Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
			Value: &e2appducontents.RicactionAdmittedItem{
				RicActionId: &e2apies.RicactionId{
					Value: int32(*raa),
				},
			},
			Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
		}
		res = append(res, raaIe)
	}
	return &e2appducontents.RicsubscriptionResponse{
		ProtocolIes: &e2appducontents.RicsubscriptionResponseIes{
			E2ApProtocolIes29: &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes29{
				Id:          int32(v2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       request.ProtocolIes.E2ApProtocolIes29.Value,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes5: &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes5{
				Id:          int32(v2.ProtocolIeIDRanfunctionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RanfunctionId{
					Value: int32(1),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes17: &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes17{
				Id:          int32(v2.ProtocolIeIDRicactionsAdmitted),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2appducontents.RicactionAdmittedList{
					Value: res,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

func (p *testClientProcedures) RICSubscriptionDelete(ctx context.Context, request *e2appducontents.RicsubscriptionDeleteRequest) (response *e2appducontents.RicsubscriptionDeleteResponse, failure *e2appducontents.RicsubscriptionDeleteFailure, err error) {
	return &e2appducontents.RicsubscriptionDeleteResponse{
		ProtocolIes: &e2appducontents.RicsubscriptionDeleteResponseIes{
			E2ApProtocolIes29: &e2appducontents.RicsubscriptionDeleteResponseIes_RicsubscriptionDeleteResponseIes29{
				Id:          int32(v2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       request.ProtocolIes.E2ApProtocolIes29.Value,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

func (p *testClientProcedures) E2ConnectionUpdate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (response *e2appducontents.E2ConnectionUpdateAcknowledge, failure *e2appducontents.E2ConnectionUpdateFailure, err error) {
	response = &e2appducontents.E2ConnectionUpdateAcknowledge{
		ProtocolIes: &e2appducontents.E2ConnectionUpdateAckIes{
			E2ApProtocolIes39: &e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes39{
				Id:          int32(v2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2appducontents.E2ConnectionUpdateList{
					Value: make([]*e2appducontents.E2ConnectionUpdateItemIes, 0),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			},
			E2ApProtocolIes49: &e2appducontents.E2ConnectionUpdateAckIes_E2ConnectionUpdateAckIes49{
				Id:          int32(v2.ProtocolIeIDTransactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.TransactionId{
					Value: request.GetProtocolIes().GetE2ApProtocolIes49().GetValue().Value,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}
	tnlInfo := request.GetProtocolIes().GetE2ApProtocolIes44().GetValue().GetValue()[0].GetValue()
	si := &e2appducontents.E2ConnectionUpdateItemIes{
		Id:          int32(v2.ProtocolIeIDE2connectionUpdateItem),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2appducontents.E2ConnectionUpdateItem{
			TnlInformation: tnlInfo.GetTnlInformation(),
			TnlUsage:       tnlInfo.GetTnlUsage(),
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
	}
	response.GetProtocolIes().GetE2ApProtocolIes39().Value.Value = append(response.GetProtocolIes().GetE2ApProtocolIes39().Value.Value, si)

	return response, nil, nil
}

type testServerProcedures struct{}

func (p *testServerProcedures) E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error) {
	return &e2appducontents.E2NodeConfigurationUpdateAcknowledge{
		ProtocolIes: &e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes{
			E2ApProtocolIes49: &e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes_E2NodeConfigurationUpdateAcknowledgeIes49{
				Id:          int32(v2.ProtocolIeIDTransactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.TransactionId{
					Value: request.GetProtocolIes().GetE2ApProtocolIes49().GetValue().Value,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

func (p *testServerProcedures) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (response *e2appducontents.E2SetupResponse, failure *e2appducontents.E2SetupFailure, err error) {
	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0x4d, 0x20, 0x00},
		RicIdentifierLen:   20,
	}
	return &e2appducontents.E2SetupResponse{
		ProtocolIes: &e2appducontents.E2SetupResponseIes{
			E2ApProtocolIes4: &e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
				Id:          int32(v2.ProtocolIeIDGlobalRicID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.GlobalRicId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{plmnID[0], plmnID[1], plmnID[2]},
					},
					RicId: &asn1.BitString{
						Value: ricID.RicIdentifierValue,
						Len:   uint32(ricID.RicIdentifierLen),
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes49: &e2appducontents.E2SetupResponseIes_E2SetupResponseIes49{
				Id:          int32(v2.ProtocolIeIDTransactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.TransactionId{
					Value: request.GetProtocolIes().GetE2ApProtocolIes49().GetValue().Value,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

func (p *testServerProcedures) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) (err error) {
	return nil
}
