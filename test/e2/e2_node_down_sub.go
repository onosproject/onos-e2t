// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"fmt"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// TestE2NodeDownSubscription checks that a subscription channel read times out if
// the e2 node is down.
func (s *TestSuite) TestE2NodeDownSubscription() {
	// Create a simulator
	sim := s.CreateRanSimulatorWithNameOrDie("e2node-down-subscription")
	nodeID := utils.GetTestNodeID(s.T())

	// Use one of the cell object IDs for action definition
	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	// Create a subscription request to indication messages from the client
	subName := "TestE2NodeDownSubscription"

	kpmv2Sub := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:   subName,
			NodeID: nodeID,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub.UseDefaultReportAction())

	s.CrashSimulatorPodOrDie("e2node-down-subscription")

	// Cause the simulator to crash
	s.UninstallRanSimulatorOrDie(sim, "e2node-down-subscription")

	for {
		pods, err := s.CoreV1().Pods(s.Namespace()).List(s.Context(), metav1.ListOptions{
			LabelSelector: fmt.Sprintf("name=%s-device-simulator", "e2node-down-subscription"),
		})
		s.NoError(err)
		if len(pods.Items) > 0 {
			time.Sleep(time.Second)
		} else {
			s.T().Log("no ransim pod")
			break
		}
	}

	//  Create the subscription
	_, err := kpmv2Sub.Subscribe(s.Context())

	//  Subscribe should have failed because of a timeout
	s.Error(err)

	// Delete the subscription and ran simulator
	sim = s.CreateRanSimulatorWithNameOrDie("e2node-down-subscription")
	s.T().Logf("Unsubscribing %s", subName)
	kpmv2Sub.Sub.UnsubscribeOrFail(s.Context(), s.T())

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "e2node-down-subscription")
	s.NoError(err)

}
