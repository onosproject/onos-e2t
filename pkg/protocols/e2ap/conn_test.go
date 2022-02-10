// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2ap

import (
	"context"
	"encoding/binary"
	"github.com/onosproject/onos-e2t/api/e2ap/v2"
	e2apies "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"testing"
	"time"
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

	configComponentAdditionItems := []*types.E2NodeComponentConfigAdditionItem{
		{E2NodeComponentType: e2apies.E2NodeComponentInterfaceType_E2NODE_COMPONENT_INTERFACE_TYPE_E1,
			E2NodeComponentID: pdubuilder.CreateE2NodeComponentIDE1(21),
			E2NodeComponentConfiguration: e2apies.E2NodeComponentConfiguration{
				E2NodeComponentRequestPart:  []byte{0x00, 0x01, 0x02},
				E2NodeComponentResponsePart: []byte{0x03, 0x04, 0x05},
			},
		},
	}

	e2SetupRequest := &e2appducontents.E2SetupRequest{
		ProtocolIes: make([]*e2appducontents.E2SetupRequestIes, 0),
	}
	e2SetupRequest.SetGlobalE2nodeID(ge2nID).SetRanFunctionsAdded(ranFunctionList).
		SetE2nodeComponentConfigAddition(configComponentAdditionItems).SetTransactionID(2)

	e2SetupResponse, e2SetupFailure, err := e2NodeCh.E2Setup(context.TODO(), e2SetupRequest)
	assert.NotNil(t, e2SetupResponse)
	assert.Nil(t, e2SetupFailure)
	assert.NoError(t, err)

	ranFuncID := types.RanFunctionID(1)
	ricRequestID := types.RicRequest{
		RequestorID: types.RicRequestorID(1),
		InstanceID:  types.RicInstanceID(2),
	}
	ricActionsTobeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsTobeSetup[types.RicActionID(1)] = types.RicActionDef{
		RicActionID:   1,
		RicActionType: e2apies.RicactionType_RICACTION_TYPE_REPORT,
	}

	ricSubscriptionRequest := &e2appducontents.RicsubscriptionRequest{
		ProtocolIes: make([]*e2appducontents.RicsubscriptionRequestIes, 0),
	}
	ricSubscriptionRequest.SetRanFunctionID(&ranFuncID).SetRicRequestID(&ricRequestID).
		SetRicSubscriptionDetails(nil, ricActionsTobeSetup)

	ricSubscriptionResponse, ricSubscriptionFailure, err := ricCh.RICSubscription(context.TODO(), ricSubscriptionRequest)
	assert.NotNil(t, ricSubscriptionResponse)
	assert.Nil(t, ricSubscriptionFailure)
	assert.NoError(t, err)

	rrID := types.RicRequest{
		RequestorID: types.RicRequestorID(1),
		InstanceID:  types.RicInstanceID(2),
	}
	rfID := types.RanFunctionID(1)

	ricIndication := &e2appducontents.Ricindication{
		ProtocolIes: make([]*e2appducontents.RicindicationIes, 0),
	}
	ricIndication.SetRicRequestID(rrID).SetRanFunctionID(rfID).SetRicActionID(2).
		SetRicIndicationSN(types.RicIndicationSn(3)).SetRicIndicationType(e2apies.RicindicationType_RICINDICATION_TYPE_REPORT).
		SetRicIndicationHeader([]byte("foo")).SetRicIndicationMessage([]byte("bar")).SetRicCallProcessID([]byte("baz"))

	err = e2NodeCh.RICIndication(context.TODO(), ricIndication)
	assert.NoError(t, err)
	time.Sleep(time.Second)

	// Test E2 connection update
	portBytes := make([]byte, 2)
	port := uint16(36421)
	binary.BigEndian.PutUint16(portBytes, port)

	connectionUpdateRequest := &e2appducontents.E2ConnectionUpdate{
		ProtocolIes: make([]*e2appducontents.E2ConnectionUpdateIes, 0),
	}
	connectionUpdateRequest.SetTransactionID(3).SetE2ConnectionUpdateAdd([]*types.E2ConnectionUpdateItem{
		{TnlInformation: types.TnlInformation{
			TnlPort: &asn1.BitString{
				Value: portBytes,
				Len:   16,
			},
			TnlAddress: asn1.BitString{
				Value: testIP,
				Len:   128,
			}},
			TnlUsage: e2apies.Tnlusage_TNLUSAGE_BOTH}})

	ack, failure, err := ricCh.E2ConnectionUpdate(context.TODO(), connectionUpdateRequest)
	assert.NoError(t, err)
	assert.NotNil(t, ack)
	assert.Nil(t, failure)

	var tnlInformation *e2apies.Tnlinformation
	for _, v := range ack.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDE2connectionSetup) {
			tnlInformation = v.GetValue().GetE2Cul().GetValue()[0].GetValue().GetE2Curi().GetTnlInformation()
			break
		}
	}

	assert.Equal(t, testIP, net.IP(tnlInformation.GetTnlAddress().GetValue()))
	assert.Equal(t, port, binary.BigEndian.Uint16(tnlInformation.GetTnlPort().GetValue()))

	// Test E2 configuration update
	configUpdate := &e2appducontents.E2NodeConfigurationUpdate{
		ProtocolIes: make([]*e2appducontents.E2NodeConfigurationUpdateIes, 0),
	}
	configUpdate.SetTransactionID(2).SetGlobalE2nodeID(ge2nID)

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

	var rrID types.RicRequest
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			rrID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			rrID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicRequestorId())
			break
		}
	}

	rca := &e2appducontents.RiccontrolAcknowledge{
		ProtocolIes: make([]*e2appducontents.RiccontrolAcknowledgeIes, 0),
	}
	rca.SetRicRequestID(rrID)

	return rca, nil, nil
}

func (p *testClientProcedures) RICSubscription(ctx context.Context, request *e2appducontents.RicsubscriptionRequest) (response *e2appducontents.RicsubscriptionResponse, failure *e2appducontents.RicsubscriptionFailure, err error) {
	var ricActionAdmitted10 types.RicActionID = 10
	var ricActionAdmitted20 types.RicActionID = 20
	ricActionsAccepted := []*types.RicActionID{&ricActionAdmitted10, &ricActionAdmitted20}

	var rrID types.RicRequest
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			rrID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			rrID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicRequestorId())
			break
		}
	}

	rfID := types.RanFunctionID(1)

	rsr := &e2appducontents.RicsubscriptionResponse{
		ProtocolIes: make([]*e2appducontents.RicsubscriptionResponseIes, 0),
	}
	rsr.SetRicRequestID(&rrID).SetRanFunctionID(&rfID).SetRicActionAdmitted(ricActionsAccepted)

	return rsr, nil, nil
}

func (p *testClientProcedures) RICSubscriptionDelete(ctx context.Context, request *e2appducontents.RicsubscriptionDeleteRequest) (response *e2appducontents.RicsubscriptionDeleteResponse, failure *e2appducontents.RicsubscriptionDeleteFailure, err error) {
	var rrID types.RicRequest
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			rrID.RequestorID = types.RicRequestorID(v.GetValue().GetRrId().GetRicRequestorId())
			rrID.InstanceID = types.RicInstanceID(v.GetValue().GetRrId().GetRicRequestorId())
			break
		}
	}

	rsdr := &e2appducontents.RicsubscriptionDeleteResponse{
		ProtocolIes: make([]*e2appducontents.RicsubscriptionDeleteResponseIes, 0),
	}
	rsdr.SetRicRequestID(&rrID)

	return rsdr, nil, nil
}

func (p *testClientProcedures) E2ConnectionUpdate(ctx context.Context, request *e2appducontents.E2ConnectionUpdate) (response *e2appducontents.E2ConnectionUpdateAcknowledge, failure *e2appducontents.E2ConnectionUpdateFailure, err error) {
	var trID int32
	var tnlInfo *e2appducontents.E2ConnectionUpdateItem
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			trID = v.GetValue().GetTrId().GetValue()
		}
		if v.Id == int32(v2.ProtocolIeIDE2connectionUpdateAdd) {
			tnlInfo = v.GetValue().GetE2Cul().GetValue()[0].GetValue().GetE2Curi()
		}
	}

	response = &e2appducontents.E2ConnectionUpdateAcknowledge{
		ProtocolIes: make([]*e2appducontents.E2ConnectionUpdateAckIes, 0),
	}
	response.SetTransactionID(trID).SetE2ConnectionSetup([]*types.E2ConnectionUpdateItem{{TnlInformation: types.TnlInformation{
		TnlPort: &asn1.BitString{
			Value: tnlInfo.GetTnlInformation().GetTnlPort().GetValue(),
			Len:   tnlInfo.GetTnlInformation().GetTnlPort().GetLen(),
		},
		TnlAddress: asn1.BitString{
			Value: tnlInfo.GetTnlInformation().GetTnlAddress().GetValue(),
			Len:   tnlInfo.GetTnlInformation().GetTnlAddress().GetLen(),
		}},
		TnlUsage: tnlInfo.GetTnlUsage()}})

	return response, nil, nil
}

type testServerProcedures struct{}

func (p *testServerProcedures) E2ConfigurationUpdate(ctx context.Context, request *e2appducontents.E2NodeConfigurationUpdate) (response *e2appducontents.E2NodeConfigurationUpdateAcknowledge, failure *e2appducontents.E2NodeConfigurationUpdateFailure, err error) {
	var trID int32
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			trID = v.GetValue().GetTrId().GetValue()
			break
		}
	}

	e2ncua := &e2appducontents.E2NodeConfigurationUpdateAcknowledge{
		ProtocolIes: make([]*e2appducontents.E2NodeConfigurationUpdateAcknowledgeIes, 0),
	}
	e2ncua.SetTransactionID(trID)

	return e2ncua, nil, nil
}

func (p *testServerProcedures) E2Setup(ctx context.Context, request *e2appducontents.E2SetupRequest) (response *e2appducontents.E2SetupResponse, failure *e2appducontents.E2SetupFailure, err error) {
	var trID int32
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDTransactionID) {
			trID = v.GetValue().GetTrId().GetValue()
			break
		}
	}

	plmnID := [3]byte{0x79, 0x78, 0x70}
	ricID := types.RicIdentifier{
		RicIdentifierValue: []byte{0x4d, 0x20, 0x00},
		RicIdentifierLen:   20,
	}

	e2sr := &e2appducontents.E2SetupResponse{
		ProtocolIes: make([]*e2appducontents.E2SetupResponseIes, 0),
	}
	e2sr.SetTransactionID(trID).SetGlobalRicID(plmnID, ricID)
	return e2sr, nil, nil
}

func (p *testServerProcedures) RICIndication(ctx context.Context, request *e2appducontents.Ricindication) (err error) {
	return nil
}
