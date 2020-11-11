// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import "github.com/spf13/cobra"

func getGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get {connections | subscriptions} [args]",
		Short: "Get E2T resources",
	}
	cmd.AddCommand(getGetConnectionsCommand())
	return cmd
}

func getWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch {subscriptions} [args]",
		Short: "Monitor E2T resources",
	}
	return cmd
}
