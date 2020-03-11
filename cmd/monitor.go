package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/smartassets/installer/ui"
	"github.com/smartassets/installer/util"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:     "monitor",
	Short:   "Monitors backend server",
	Example: "assets monitor <container-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := util.CheckDocker()
		if err != nil {
			ui.ReportErr("Docker daemon is not working. Docker is required to be up and running.")
		}

		cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not communicate with Docker: %s", err.Error()))
		}

		context := context.TODO()
		_, err = cli.ContainerLogs(context, args[0], types.ContainerLogsOptions{Follow: true, ShowStdout: true, ShowStderr: true})
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not monitor service %s: %s", args[0], err.Error()))
		}

		//buffer := make([]byte, 16)
		//if _, err := io.CopyBuffer(os.Stdout, logsReader, buffer); err != nil {
		//	logsReader.Close()
		//	ui.ReportErr(fmt.Sprintf("Could not stream backend service logs: %s", err.Error()))
		//}
	},
}
