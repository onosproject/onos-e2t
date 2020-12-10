// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"strconv"
	"strings"
	"time"

	"github.com/onosproject/onos-e2t/test/utils"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	"text/tabwriter"

	"github.com/spf13/cobra"
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

// createSubscriptionRequest make a proto-encoded request for a subscription to indication data.
// TODO : revisit this when JSON encoding is supported, and make this more general
func createSubscriptionRequest(nodeID string) (subscription.SubscriptionDetails, error) {
	return subscription.SubscriptionDetails{
		E2NodeID: subscription.E2NodeID(nodeID),
		ServiceModel: subscription.ServiceModel{
			ID: subscription.ServiceModelID("test"),
		},
		EventTrigger: subscription.EventTrigger{
			Payload: subscription.Payload{
				Encoding: subscription.Encoding_ENCODING_PROTO,
				Data:     []byte{},
			},
		},
		Actions: []subscription.Action{
			{
				ID:   100,
				Type: subscription.ActionType_ACTION_TYPE_REPORT,
				SubsequentAction: &subscription.SubsequentAction{
					Type:       subscription.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
					TimeToWait: subscription.TimeToWait_TIME_TO_WAIT_ZERO,
				},
			},
		},
	}, nil
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
