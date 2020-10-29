// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import "github.com/spf13/cobra"

func getGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get {connections} [args]",
		Short: "Get E2T resources",
	}
	cmd.AddCommand(getGetConnectionsCommand())
	return cmd
}
