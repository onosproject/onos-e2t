// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package channel

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2appdudescriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/codec"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2/channel/filter"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	conn := NewMockConn(ctrl)

	readCh := make(chan *e2appdudescriptions.E2ApPdu, 1)
	conn.EXPECT().Write(gomock.Any()).DoAndReturn(func(out []byte) (int, error) {
		_, err := codec.XER.Decode(out)
		if err != nil {
			return 0, err
		}
		return len(out), nil
	}).AnyTimes()

	conn.EXPECT().Read(gomock.Any()).DoAndReturn(func(in []byte) (int, error) {
		res := <-readCh
		out, err := codec.XER.Encode(res)
		if err != nil {
			return 0, err
		}
		return copy(in, out), nil
	}).AnyTimes()

	meta := Metadata{
		ID:           1,
		PlmnID:       "onf",
		RANFunctions: map[RANFunctionID]RANFunctionMetadata{},
	}
	channel := newChannel(context.Background(), conn, meta)
	assert.Equal(t, ID(1), channel.ID())
	assert.Equal(t, ID(1), channel.Metadata().ID)
	assert.Equal(t, PlmnID("onf"), channel.Metadata().PlmnID)

	req := newSubscribeRequest()
	err := channel.Send(req, codec.XER)
	assert.NoError(t, err)

	subCh := channel.Recv(filter.RicSubscription(req.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value), codec.XER)

	readCh <- newSubscribeResponse()

	res := <-subCh

	readCh <- newSubscribeResponse()

	res, err = channel.SendRecv(req, filter.RicSubscription(req.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value), codec.XER)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	ch := channel.Recv(filter.RicIndication(req.GetInitiatingMessage().ProcedureCode.RicSubscription.InitiatingMessage.ProtocolIes.E2ApProtocolIes29.Value), codec.XER)

	readCh <- newIndication()
	indication := <-ch
	assert.NotNil(t, indication)

	readCh <- newIndication()
	indication = <-ch
	assert.NotNil(t, indication)
}

func newSubscribeRequest() *e2appdudescriptions.E2ApPdu {
	req, _ := pdubuilder.CreateRicSubscriptionRequestE2apPdu(1, 2, 3, 4, e2apies.RicactionType_RICACTION_TYPE_REPORT, e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE, e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO, []byte{}, []byte{})
	return req
}

func newSubscribeResponse() *e2appdudescriptions.E2ApPdu {
	res, _ := pdubuilder.CreateRicSubscriptionResponseE2apPdu(1, 2, 3)
	return res
}

func newIndication() *e2appdudescriptions.E2ApPdu {
	res, _ := pdubuilder.RicIndicationE2apPdu(1, 2, 3, 4, 1, e2apies.RicindicationType_RICINDICATION_TYPE_REPORT, "foo", "bar", "baz")
	return res
}
