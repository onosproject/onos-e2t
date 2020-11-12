// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/gogo/protobuf/proto"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"

	"github.com/onosproject/onos-e2t/test/utils"

	"gotest.tools/assert"

	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
)

const (
	OnosE2subAddress = "onos-e2sub:5150"
)

func createSubscriptionRequest() (*subapi.Subscription, error) {

	var ricAction = e2apies.RicactionType_RICACTION_TYPE_REPORT
	var ricSubsequentAction = e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE
	var ricttw = e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO
	E2apPdu, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(1, 1,
		33, 255, ricAction, ricSubsequentAction, ricttw, []byte{0xAA}, []byte{0xBB})

	if err != nil {
		return &subapi.Subscription{}, err
	}
	e2apPayload, err := proto.Marshal(E2apPdu)
	if err != nil {
		return &subapi.Subscription{}, err
	}

	id := uuid.New()
	subReq := subapi.Subscription{
		ID: subapi.ID(id.String()),
		Payload: &subapi.Payload{
			Bytes: e2apPayload,
		},
		E2NodeID: 0xcf7a40,
	}

	return &subReq, nil

}

// TestSubscription
func (s *TestSuite) TestSubscription(t *testing.T) {
	utils.CreateE2Simulator(t)
	client, err := subscription.NewClient(context.Background(), subscription.Destination{
		Addrs: []string{OnosE2subAddress},
	})
	assert.NilError(t, err)

	subReq, err := createSubscriptionRequest()
	assert.NilError(t, err)

	err = client.Add(context.Background(), subReq)
	assert.NilError(t, err)
}
