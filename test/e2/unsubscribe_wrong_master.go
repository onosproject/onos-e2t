// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/onos-e2t/test/e2utils"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-e2t/test/utils"
)

// TestUnsubscribeWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestUnsubscribeWrongMaster() {
	if s.E2TReplicaCount == 1 {
		// Test is not applicable - no non-master nodes
		s.T().Skip("Test not applicable to single node")
		return
	}

	sim := s.CreateRanSimulatorWithNameOrDie("unsubscribe-non-master")
	s.NotNil(sim)

	e2NodeID := utils.GetTestNodeID(s.T())
	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), e2NodeID)

	nonMasters := utils.GetE2NodeNonMasterNodes(s.T(), e2NodeID)
	master := utils.GetE2NodeMaster(s.T(), e2NodeID)

	nonMasterClient := utils.GetSubClientForIP(s.T(), nonMasters[0].IP, nonMasters[0].Port)
	s.NotNil(nonMasterClient)
	masterClient := utils.GetSubClientForIP(s.T(), master.Interface.IP, master.Interface.Port)
	s.NotNil(masterClient)

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

	c, err := masterClient.Subscribe(s.Context(), req)
	s.NoError(err)

	msg, err := c.Recv()
	s.NotNil(msg)
	s.NoError(err)

	unsubscribeRequest := &e2api.UnsubscribeRequest{
		Headers:       headers,
		TransactionID: "sub1",
	}
	unsubscribeResponse, err := nonMasterClient.Unsubscribe(s.Context(), unsubscribeRequest)
	s.NoError(err)
	s.NotNil(unsubscribeResponse)

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "unsubscribe-non-master")
}
