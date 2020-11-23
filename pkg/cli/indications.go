// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta1/e2apies"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/pdubuilder"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"strconv"
	"strings"
	"time"

	//e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/encoding"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/node"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/subscription"
	"github.com/spf13/cobra"
	"text/tabwriter"
)

func getWatchIndicationsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "indications",
		Short: "watch indications traffic",
		RunE:  runWatchIndicationsCommand,
	}
	cmd.Flags().Bool("no-headers", false, "disables output headers")
	cmd.PersistentFlags().String("service-address", "onos-e2sub:5150", "the gRPC endpoint")
	cmd.Flags().Duration("timeout", time.Hour, "specifies maximum wait time for new indications")
	return cmd
}

func createSubscriptionRequest(nodeID string) (subscription.Subscription, error) {
	ricActionsToBeSetup := make(map[types.RicActionID]types.RicActionDef)
	ricActionsToBeSetup[100] = types.RicActionDef{
		RicActionID:         100,
		RicActionType:       e2apies.RicactionType_RICACTION_TYPE_REPORT,
		RicSubsequentAction: e2apies.RicsubsequentActionType_RICSUBSEQUENT_ACTION_TYPE_CONTINUE,
		Ricttw:              e2apies.RictimeToWait_RICTIME_TO_WAIT_ZERO,
		RicActionDefinition: []byte{0x11, 0x22},
	}

	E2apPdu, err := pdubuilder.CreateRicSubscriptionRequestE2apPdu(types.RicRequest{RequestorID: 0, InstanceID: 0},
		0, nil, ricActionsToBeSetup)

	if err != nil {
		return subscription.Subscription{}, err
	}

	subReq := subscription.Subscription{
		EncodingType: encoding.PROTO,
		NodeID:       node.ID(nodeID),
		Payload: subscription.Payload{
			Value: E2apPdu,
		},
	}

	return subReq, nil

}

func runWatchIndicationsCommand(cmd *cobra.Command, args []string) error {
	outputWriter := cli.GetOutput()
	writer := new(tabwriter.Writer)
	writer.Init(outputWriter, 0, 0, 3, ' ', tabwriter.FilterHTML)

	timeout, _ := cmd.Flags().GetDuration("timeout")
	address, _ := cmd.Flags().GetString("service-address")
	tokens := strings.Split(address, ":")
	if len(tokens) != 2 {
		return errors.New("service-address must be of the form host:port")
	}
	host := tokens[0]

	port, err := strconv.Atoi(tokens[1])
	if err != nil {
		return err
	}

	clientConfig := e2client.Config{
		AppID: "subscription-test",
		SubscriptionService: e2client.ServiceConfig{
			Host: host,
			Port: port,
		},
	}

	client, err := e2client.NewClient(clientConfig)
	if err != nil {
		return err
	}

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	if err != nil {
		return err
	}

	subReq, err := createSubscriptionRequest(nodeIDs[0])
	if err != nil {
		return err
	}

	err = client.Subscribe(ctx, subReq, ch)
	if err != nil {
		return err
	}

	done := false
	for !done {
		select {
		case indicationMsg := <-ch:
			_, _ = fmt.Fprintf(writer, "%v\n\n", indicationMsg)
			_ = writer.Flush()
			break
		case <-time.After(timeout):
			done = true
			break
		}
	}

	return nil
}
