// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package cli

import (
	"context"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/onosproject/onos-e2t/api/admin/v1"
	"github.com/spf13/cobra"
)

func getGetConnectionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connections",
		Short: "Get SB connections",
		RunE:  runConnectionsCommand,
	}
	return cmd
}

func runConnectionsCommand(cmd *cobra.Command, args []string) error {
	conn, err := cli.GetConnection(cmd)
	if err != nil {
		return err
	}
	defer conn.Close()
	outputWriter := cli.GetOutput()
	writer := new(tabwriter.Writer)
	writer.Init(outputWriter, 0, 0, 3, ' ', tabwriter.FilterHTML)

	request := admin.ListE2NodeConnectionsRequest{}

	fmt.Fprintln(writer, "Connecting to server")

	client := admin.CreateE2TAdminServiceClient(conn)

	stream, err := client.ListE2NodeConnections(context.Background(), &request)
	if err != nil {
		return err
	}

	fmt.Fprintln(writer, "Connected to server")
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			cli.Output("recv error")
			return err
		}

		//_, _ = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", response.Ecgi.GetPlmnid(), response.Ecgi.GetEcid(), response.GetCrnti(), response.GetImsi())
		fmt.Fprintln(writer, "connection %v", response)
	}

	_ = writer.Flush()

	return nil
}
