// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"testing"

	"github.com/onosproject/onos-e2t/test/e2utils"
	sdkclient "github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

// TestIdenticalSubscriptionMultiApps tests identical subscriptions are absorbed by E2T from different xApps
func (s *TestSuite) TestIdenticalSubscriptionMultiApps(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "identical-subscription-multi-app")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	nodeID := utils.GetTestNodeID(t)

	subName1 := "identical-sub-app1"
	subName2 := "identical-sub-app2"
	cellObjectID := e2utils.GetFirstCellObjectID(t, nodeID)

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
	channelIDApp1, err := kpmv2Sub1.Subscribe(ctx)
	assert.NoError(t, err)

	kpmv2Sub2 := e2utils.KPMV2Sub{
		Sub: e2utils.Sub{
			Name:      subName2,
			NodeID:    nodeID,
			SdkClient: clientApp2,
		},
		CellObjectID: cellObjectID,
	}
	channelIDApp2, err := kpmv2Sub2.Subscribe(ctx)
	assert.NoError(t, err)

	assert.True(t, channelIDApp1 != channelIDApp2)

	indicationReportApp1 := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub1.Sub.Ch)
	assert.NotNil(t, indicationReportApp1)
	indicationReportApp2 := e2utils.CheckIndicationMessage(t, e2utils.DefaultIndicationTimeout, kpmv2Sub2.Sub.Ch)
	assert.NotNil(t, indicationReportApp2)

	subList := e2utils.GetSubscriptionList(t)
	assert.Equal(t, 1, len(subList))

	kpmv2Sub1.Sub.UnsubscribeOrFail(ctx, t)

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName1, subList)

	kpmv2Sub2.Sub.UnsubscribeOrFail(ctx, t)

	subList = e2utils.GetSubscriptionList(t)
	t.Logf("Subscription List after deleting subscription %s is %v:", subName2, subList)

	e2utils.CheckForEmptySubscriptionList(t)
	utils.UninstallRanSimulatorOrDie(t, sim)
}
