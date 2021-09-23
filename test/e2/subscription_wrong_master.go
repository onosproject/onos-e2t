// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"os"
	"testing"

	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/stretchr/testify/assert"

	"github.com/onosproject/onos-e2t/test/utils"
)

func findMasterRelation(t *testing.T, e2Node topoapi.Object, nodeID topoapi.ID) *topoapi.Relation {
	topoSdkClient, err := utils.NewTopoClient()
	assert.NoError(t, err)
	relations, err := topoSdkClient.GetControlRelationsForTarget()
	assert.NoError(t, err)

	// TODO - replace this with a filter when one is available
	var result *topoapi.Relation
	for _, relationObject := range relations {
		relation := relationObject.GetRelation()
		fmt.Fprintf(os.Stderr, "src entity %s node id %s\n", relation.SrcEntityID, e2Node.ID)
		fmt.Fprintf(os.Stderr, "tgt entity %s node id %s\n", relation.TgtEntityID, nodeID)
		if relation.SrcEntityID == e2Node.ID &&
			relation.TgtEntityID == nodeID {
			result = relation
		}
	}
	return result
}

// TestSubscriptionWrongMaster tests e2 subscription to a non-master node
func (s *TestSuite) TestSubscriptionWrongMaster(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, s.c, "subscription-kpm-v2")
	assert.NotNil(t, sim)

	ctx, cancel := context.WithTimeout(context.Background(), subscriptionTimeout)
	defer cancel()

	topoClient, err := utils.NewTopoClient()
	assert.NoError(t, err)

	e2tNodes, err := topoClient.E2TNodes(ctx)
	assert.NoError(t, err)

	e2NodeID := utils.GetTestNodeID(t)

	nonMasterIP := ""
	nonMasterPort := uint32(0)
	for _, node := range e2tNodes {
		e2tIP := ""
		e2tPort := uint32(0)
		e2tInfo := &topoapi.E2TInfo{}
		err = node.GetAspect(e2tInfo)
		assert.NoError(t, err)
		for _, iface := range e2tInfo.Interfaces {
			fmt.Fprintf(os.Stderr, "!!!!!!!    found e2ap200 interface E2T node %v \n", iface)
			if iface.Type == topoapi.Interface_INTERFACE_E2T {
				e2tIP = iface.IP
				e2tPort = iface.Port
				fmt.Fprintf(os.Stderr, "!!!!!!!    found e2ap200 interface E2T node %v:%v \n", e2tIP, e2tPort)
				break
			}
		}
		rel := findMasterRelation(t, node, e2NodeID)
		if rel != nil {
			fmt.Fprintf(os.Stderr, "------    E2T node %v:%v is master\n", e2tIP, e2tPort)
		} else {
			nonMasterIP = e2tIP
			nonMasterPort = e2tPort
			fmt.Fprintf(os.Stderr, "++++++    E2T node %v:%v is not master\n", e2tIP, e2tPort)
		}
	}

	fmt.Fprintf(os.Stderr, "non master ip is %s:%d\n", nonMasterIP, nonMasterPort)
	client := utils.GetSubClientForIP(t, nonMasterIP, nonMasterPort)
	assert.NotNil(t, client)

	spec := utils.CreateKpmV2Sub(t, e2NodeID)

	req := &e2api.SubscribeRequest{
		Headers:            e2api.RequestHeaders{
			AppID:         "app",
			AppInstanceID: "",
			E2NodeID:      e2api.E2NodeID(e2NodeID),
			ServiceModel:  e2api.ServiceModel{
				Name:    utils.KpmServiceModelName,
				Version: utils.Version2,
			},
		},
		TransactionID:      "sub1",
		Subscription:       spec,
	}

	c, err := client.Subscribe(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	response, err := c.Recv()
	assert.NoError(t, err)
	assert.NotNil(t, response)

	fmt.Fprintf(os.Stderr, "Recv of:\n%v\n", response)

	//assert.NoError(t, sim.Uninstall())
	//e2utils.CheckForEmptySubscriptionList(t)
}
