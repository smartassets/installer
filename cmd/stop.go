package cmd

import (
	"fmt"
	"github.com/smartassets/installer/ui"
	"github.com/smartassets/installer/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stops already installed backend in the current directory",
	Example: "assets stop",
	Run: func(cmd *cobra.Command, args []string) {
		err := util.CheckDocker()
		if err != nil {
			ui.ReportErr("Docker daemon is not working. Docker is required to be up and running.")
		}

		installationDirectory := getCurrentDirectory()
		_, err = os.Stat(installationDirectory)
		if os.IsNotExist(err) {
			ui.ReportErr(fmt.Sprintf("Location %s does not exists. Is the backend correctly installed?", installationDirectory))
		}

		ui.ReportInfoWithoutArgs("Stopping server...")
		dockerComposeCmd := exec.Command("docker-compose",  "-p", "assets", "down")
		dockerComposeCmd.Stdout = os.Stdout
		dockerComposeCmd.Stderr = os.Stderr
		dockerComposeCmd.Stdin = os.Stdin
		err = dockerComposeCmd.Run()
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not start server: %s", err))
		}
	},
}