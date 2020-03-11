package cmd

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/smartassets/installer/ui"
	"github.com/smartassets/installer/util"
	"github.com/spf13/cobra"
	"strings"
)

var listCmd = &cobra.Command{
	Use:     "list-services",
	Aliases: []string{"ls"},
	Short:   "List backend services",
	Example: "assets list-services",
	Run: func(cmd *cobra.Command, args []string) {
		err := util.CheckDocker()
		if err != nil {
			ui.ReportErr("Docker daemon is not working. Docker is required to be up and running.")
		}

		cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not communicate with Docker daemon: %s", err.Error()))
		}

		context := context.TODO()
		containers, err := cli.ContainerList(context, types.ContainerListOptions{All:true})
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not list backend services: %s", err.Error()))
		}

		ui.ReportInfoWithoutArgs("Installed backed services:")
		for _, container := range containers {
			ui.ReportSimpleWithoutArgs(strings.Trim(container.Names[0], "/"))
		}

	},
}