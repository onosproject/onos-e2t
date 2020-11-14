// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"

	"github.com/onosproject/onos-e2t/api/admin/v1"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"gotest.tools/assert"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/node"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
)

const (
	SubscriptionServiceHost = "onos-e2sub"
	SubscriptionServicePort = 5150
	OnosE2TAddress          = "onos-e2t:5150"
)

func createSubscriptionRequest(nodeID string) (subscription.Subscription, error) {
	var ricAction = e2apies.RicactionType_RICACTION_TYPE_REPORT
	var ricSubsequentAction = e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE
	var ricttw = e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO
	E2apPdu, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(0, 0,
		0, 255, ricAction, ricSubsequentAction, ricttw, []byte{0xAA}, []byte{0xBB})

	if err != nil {
		return subscription.Subscription{}, err
	}

	subReq := subscription.Subscription{
		NodeID:  node.ID(nodeID),
		Payload: E2apPdu,
	}

	return subReq, nil

}

func getNodeIDs() ([]string, error) {
	tlsConfig, err := creds.GetClientCredentials()
	var nodeIDs []string
	if err != nil {
		return []string{}, err
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	conn, err := grpc.DialContext(context.Background(), OnosE2TAddress, opts...)
	if err != nil {
		return []string{}, err
	}
	adminClient := admin.NewE2TAdminServiceClient(conn)
	connections, err := adminClient.ListE2NodeConnections(context.Background(), &admin.ListE2NodeConnectionsRequest{})

	if err != nil {
		return []string{}, err
	}

	for {
		connection, err := connections.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []string{}, err
		}
		if connection != nil {
			nodeID := connection.Id
			nodeIDs = append(nodeIDs, nodeID)
		}

	}

	return nodeIDs, nil
}

// TestSubscription
func (s *TestSuite) TestSubscription(t *testing.T) {
	utils.CreateE2Simulator(t)

	clientConfig := e2client.Config{
		AppID: "test-subscription",
		SubscriptionService: e2client.ServiceConfig{
			Host: SubscriptionServiceHost,
			Port: SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NilError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := getNodeIDs()

	subReq, err := createSubscriptionRequest(nodeIDs[0])
	assert.NilError(t, err)

	err = client.Subscribe(ctx, subReq, ch)
	assert.NilError(t, err)

	for indicationMessage := range ch {
		fmt.Println(indicationMessage)
	}

}
