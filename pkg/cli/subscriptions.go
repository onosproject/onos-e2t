// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import (
	"context"
	"fmt"
	"io"
	"text/tabwriter"

	subscription "github.com/onosproject/onos-e2t/api/subscription/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/spf13/cobra"
)

const (
	subscriptionHeaders = "ID\tRevision\tApp ID\tService Model ID"
	subscriptionFormat  = "%s\t%d\t%s\t%s\n"
)

func getListSubscriptionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Get NB subscriptions",
		RunE:  runListSubscriptionsCommand,
	}
	cmd.Flags().Bool("no-headers", false, "disables output headers")
	return cmd
}

func runListSubscriptionsCommand(cmd *cobra.Command, args []string) error {
	noHeaders, _ := cmd.Flags().GetBool("no-headers")
	conn, err := cli.GetConnection(cmd)
	if err != nil {
		return err
	}
	defer conn.Close()
	outputWriter := cli.GetOutput()
	writer := new(tabwriter.Writer)
	writer.Init(outputWriter, 0, 0, 3, ' ', tabwriter.FilterHTML)

	if !noHeaders {
		_, _ = fmt.Fprintln(writer, subscriptionHeaders)
	}

	request := subscription.ListSubscriptionsRequest{}

	client := subscription.NewSubscriptionServiceClient(conn)

	response, err := client.ListSubscriptions(context.Background(), &request)
	if err != nil {
		return err
	}

	for _, sub := range response.Subscriptions {
		_, _ = fmt.Fprintf(writer, subscriptionFormat, sub.ID, sub.Revision, sub.AppID, sub.ServiceModel.ID)
	}

	_ = writer.Flush()

	return nil
}

func getWatchSubscriptionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscriptions",
		Short: "Watch NB subscriptions",
		RunE:  runWatchSubscriptionsCommand,
	}
	cmd.Flags().Bool("no-headers", false, "disables output headers")
	return cmd
}

func runWatchSubscriptionsCommand(cmd *cobra.Command, args []string) error {
	noHeaders, _ := cmd.Flags().GetBool("no-headers")
	conn, err := cli.GetConnection(cmd)
	if err != nil {
		return err
	}
	defer conn.Close()
	outputWriter := cli.GetOutput()
	writer := new(tabwriter.Writer)
	writer.Init(outputWriter, 0, 0, 3, ' ', tabwriter.FilterHTML)

	if !noHeaders {
		_, _ = fmt.Fprintln(writer, subscriptionHeaders)
		_ = writer.Flush()
	}

	request := subscription.WatchSubscriptionsRequest{}

	client := subscription.NewSubscriptionServiceClient(conn)

	stream, err := client.WatchSubscriptions(context.Background(), &request)
	if err != nil {
		return err
	}

	for {
		sub, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			cli.Output("recv error")
			return err
		}

		_, _ = fmt.Fprintf(writer, subscriptionFormat, sub.Subscription.ID, sub.Subscription.Revision, sub.Subscription.AppID, sub.Subscription.ServiceModel.ID)
	}

	_ = writer.Flush()

	return nil
}
