// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"github.com/onosproject/onos-e2t/test/e2utils"
	"github.com/onosproject/onos-e2t/test/utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
)

// TestIdenticalSubscriptionMultiApps tests identical subscriptions are absorbed by E2T from different xApps
func (s *TestSuite) TestIdenticalSubscriptionMultiApps() {
	sim := s.CreateRanSimulatorWithNameOrDie("identical-subscription-multi-app")
	s.NotNil(sim)

	ctx, cancel := context.WithTimeout(s.Context(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(s.T())

	subName1 := "identical-sub-app1"
	subName2 := "identical-sub-app2"
	cellObjectID := e2utils.GetFirstCellObjectID(s.T(), nodeID)

	clientApp1 := sdkclient.NewClient(sdkclient.WithE2TAddress(utils.E2TServiceHost, utils.E2TServicePort),
		sdkclient.WithServiceModel(utils.KpmServiceModelName,
			utils.Version2),
		sdkclient.WithEncoding(sdkclient.ProtoEncoding),
		sdkclient.WithAppID("app1"))

	clientApp2 := sdkclient.NewClient(sdkclient.WithE2TAddress(utils.E2TServiceHost, utils.E2TServicePort),
		sdkclient.WithServiceModel(utils.KpmServiceModelName, utils.Version2),
		sdkclient.WithEncoding(sdkclient.ProtoEncoding),
		sdkclient.WithAppID("app2"))

	kpmv2Sub1 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:      subName1,
			NodeID:    nodeID,
			SdkClient: clientApp1,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub1.UseDefaultReportAction())
	channelIDApp1, err := kpmv2Sub1.Subscribe(ctx)
	s.NoError(err)

	kpmv2Sub2 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:      subName2,
			NodeID:    nodeID,
			SdkClient: clientApp2,
		},
		CellObjectID: cellObjectID,
	}
	s.NoError(kpmv2Sub2.UseDefaultReportAction())
	channelIDApp2, err := kpmv2Sub2.Subscribe(ctx)
	s.NoError(err)

	s.True(channelIDApp1 != channelIDApp2)

	indicationReportApp1 := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub1.Sub.Ch)
	s.NotNil(indicationReportApp1)
	indicationReportApp2 := e2utils.CheckIndicationMessage(s.T(), e2utils.DefaultIndicationTimeout, kpmv2Sub2.Sub.Ch)
	s.NotNil(indicationReportApp2)

	subList := e2utils.GetSubscriptionList(s.T())
	s.Equal(1, len(subList))

	kpmv2Sub1.Sub.UnsubscribeOrFail(ctx, s.T())

	subList = e2utils.GetSubscriptionList(s.T())
	s.T().Logf("Subscription List after deleting subscription %s is %v:", subName1, subList)

	kpmv2Sub2.Sub.UnsubscribeOrFail(ctx, s.T())

	subList = e2utils.GetSubscriptionList(s.T())
	s.T().Logf("Subscription List after deleting subscription %s is %v:", subName2, subList)

	e2utils.CheckForEmptySubscriptionList(s.T())
	s.UninstallRanSimulatorOrDie(sim, "identical-subscription-multi-app")
}
