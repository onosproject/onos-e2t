// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestSubscribeWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestSubscribeWrongMaster() {
	if s.E2TReplicaCount == 1 {
		// Test is not applicable - no non-master nodes
		s.T().Skip("Test not applicable to single node")
		return
	}
	sim := s.CreateRanSimulatorWithNameOrDie("subscription-wrong-master")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()

	e2NodeID := utils.GetTestNodeID(s.T())
	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), e2NodeID)
	nonMasters := utils.GetE2NodeNonMasterNodes(s.T(), e2NodeID)

	client := utils.GetSubClientForIP(s.T(), nonMasters[0].IP, nonMasters[0].Port)
	s.NotNil(client)

	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   "sub1",
			NodeID: e2NodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())
	spec, err := kpmv2Sub.CreateSubscriptionSpec()
	s.NoError(err)

	headers := e2api.RequestHeaders{
		AppID:         "app",
		AppInstanceID: "",
		E2NodeID:      e2api.E2NodeID(e2NodeID),
		ServiceModel: e2api.ServiceModel{
			Name:    utils.KpmServiceModelName,
			Version: utils.Version2,
		},
	}

	req := &e2api.SubscribeRequest{
		Headers:       headers,
		TransactionID: "sub1",
		Subscription:  spec,
	}

	c, err := client.Subscribe(ctx, req)
	s.NoError(err)

	resp, err := c.Recv()
	s.Nil(resp)
	s.Equal(codes.Unavailable, status.Code(err))

	unsubscribeRequest := &e2api.UnsubscribeRequest{
		Headers:       headers,
		TransactionID: "sub1",
	}
	unsubscribeResponse, err := client.Unsubscribe(ctx, unsubscribeRequest)
	s.NoError(err)
	s.NotNil(unsubscribeResponse)

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "subscription-wrong-master")
}
