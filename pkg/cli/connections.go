// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import (
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/onosproject/onos-e2t/api/admin/v1"
	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/spf13/cobra"
)

func getGetConnectionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connections",
		Short: "Get SB connections",
		RunE:  runConnectionsCommand,
	}
	cmd.Flags().Bool("no-headers", false, "disables output headers")
	return cmd
}

func runConnectionsCommand(cmd *cobra.Command, args []string) error {
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
		_, _ = fmt.Fprintln(writer, "Global ID\tPLNM ID\tIP Addr\tPort")
	}

	request := admin.ListE2NodeConnectionsRequest{}

	client := admin.CreateE2TAdminServiceClient(conn)

	stream, err := client.ListE2NodeConnections(context.Background(), &request)
	if err != nil {
		return err
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			cli.Output("recv error")
			return err
		}

		_, _ = fmt.Fprintf(writer, "%d\t%s\t%s\t%d\n", response.Id, response.PlmnId, strings.Join(response.RemoteIp, ","), response.RemotePort)
	}

	_ = writer.Flush()

	return nil
}
