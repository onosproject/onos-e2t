// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import "github.com/spf13/cobra"

func getListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list {connections | subscriptions} [args]",
		Short: "List E2T resources",
	}
	cmd.AddCommand(getListConnectionsCommand())
	return cmd
}
