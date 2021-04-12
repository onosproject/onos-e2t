// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channels

import (
	"context"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/protocols/e2ap101/procedures"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap101/types"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"testing"
	"time"
)

func TestChannels(t *testing.T) {
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

	e2NodeCh := NewE2NodeChannel(clientConn, func(channel E2NodeChannel) procedures.E2NodeProcedures {
		return &testClientProcedures{}
	})
	ricCh := NewRICChannel(serverConn, func(channel RICChannel) procedures.RICProcedures {
		return &testServerProcedures{}
	})

	e2SetupRequest := &e2appducontents.E2SetupRequest{
		ProtocolIes: &e2appducontents.E2SetupRequestIes{
			E2ApProtocolIes3: &e2appducontents.E2SetupRequestIes_E2SetupRequestIes3{
				Id:          int32(v1beta2.ProtocolIeIDGlobalE2nodeID),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.GlobalE2NodeId{
					GlobalE2NodeId: &e2apies.GlobalE2NodeId_GNb{
						GNb: &e2apies.GlobalE2NodeGnbId{
							GlobalGNbId: &e2apies.GlobalgNbId{
								PlmnId: &e2ap_commondatatypes.PlmnIdentity{
									Value: []byte{1, 2, 3},
								},
								GnbId: &e2apies.GnbIdChoice{
									GnbIdChoice: &e2apies.GnbIdChoice_GnbId{
										GnbId: &e2ap_commondatatypes.BitString{
											Value: 0x9bcd4,
											Len:   22,
										}},
								},
							},
						},
					},
				},
			},
			E2ApProtocolIes10: &e2appducontents.E2SetupRequestIes_E2SetupRequestIes10{
				Id:          int32(v1beta2.ProtocolIeIDRanfunctionsAdded),
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2appducontents.RanfunctionsList{
					Value: make([]*e2appducontents.RanfunctionItemIes, 0),
				},
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
				Id:          int32(v1beta2.ProtocolIeIDRanfunctionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RanfunctionId{
					Value: 1,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes29: &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes29{
				Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicrequestId{
					RicRequestorId: 1,
					RicInstanceId:  2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes30: &e2appducontents.RicsubscriptionRequestIes_RicsubscriptionRequestIes30{
				Id:          int32(v1beta2.ProtocolIeIDRicsubscriptionDetails),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2appducontents.RicsubscriptionDetails{
					RicEventTriggerDefinition: &e2ap_commondatatypes.RiceventTriggerDefinition{},
					RicActionToBeSetupList: &e2appducontents.RicactionsToBeSetupList{
						Value: []*e2appducontents.RicactionToBeSetupItemIes{
							{
								Id:          int32(v1beta2.ProtocolIeIDRicactionToBeSetupItem),
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
				Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicrequestId{
					RicRequestorId: 1,
					RicInstanceId:  2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes5: &e2appducontents.RicindicationIes_RicindicationIes5{
				Id:          int32(v1beta2.ProtocolIeIDRanfunctionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RanfunctionId{
					Value: 1,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes15: &e2appducontents.RicindicationIes_RicindicationIes15{
				Id:          int32(v1beta2.ProtocolIeIDRicactionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicactionId{
					Value: 2,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes27: &e2appducontents.RicindicationIes_RicindicationIes27{
				Id:          int32(v1beta2.ProtocolIeIDRicindicationSn),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RicindicationSn{
					Value: 3,
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
			},
			E2ApProtocolIes28: &e2appducontents.RicindicationIes_RicindicationIes28{
				Id:          int32(v1beta2.ProtocolIeIDRicindicationType),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       e2apies.RicindicationType_RICINDICATION_TYPE_REPORT,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes25: &e2appducontents.RicindicationIes_RicindicationIes25{
				Id:          int32(v1beta2.ProtocolIeIDRicindicationHeader),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2ap_commondatatypes.RicindicationHeader{
					Value: []byte("foo"),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes26: &e2appducontents.RicindicationIes_RicindicationIes26{
				Id:          int32(v1beta2.ProtocolIeIDRicindicationMessage),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2ap_commondatatypes.RicindicationMessage{
					Value: []byte("bar"),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes20: &e2appducontents.RicindicationIes_RicindicationIes20{
				Id:          int32(v1beta2.ProtocolIeIDRiccallProcessID),
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
				Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
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
			Id:          int32(v1beta2.ProtocolIeIDRicactionAdmittedItem),
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
				Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       request.ProtocolIes.E2ApProtocolIes29.Value,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes5: &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes5{
				Id:          int32(v1beta2.ProtocolIeIDRanfunctionID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.RanfunctionId{
					Value: int32(1),
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
			E2ApProtocolIes17: &e2appducontents.RicsubscriptionResponseIes_RicsubscriptionResponseIes17{
				Id:          int32(v1beta2.ProtocolIeIDRicactionsAdmitted),
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
				Id:          int32(v1beta2.ProtocolIeIDRicrequestID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value:       request.ProtocolIes.E2ApProtocolIes29.Value,
				Presence:    int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

type testServerProcedures struct{}

func (p *testServerProcedures) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (response *e2appducontents.E2SetupResponse, failure *e2appducontents.E2SetupFailure, err error) {
	return &e2appducontents.E2SetupResponse{
		ProtocolIes: &e2appducontents.E2SetupResponseIes{
			E2ApProtocolIes4: &e2appducontents.E2SetupResponseIes_E2SetupResponseIes4{
				Id:          int32(v1beta2.ProtocolIeIDGlobalRicID),
				Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_REJECT),
				Value: &e2apies.GlobalRicId{
					PLmnIdentity: &e2ap_commondatatypes.PlmnIdentity{
						Value: []byte{0x01, 0x02, 0x03},
					},
					RicId: &e2ap_commondatatypes.BitString{
						Value: 0xABCDE,
						Len:   20,
					},
				},
				Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_MANDATORY),
			},
		},
	}, nil, nil
}

func (p *testServerProcedures) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) (err error) {
	return nil
}
