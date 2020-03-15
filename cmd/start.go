package cmd

import (
	"fmt"
	"github.com/smartassets/installer/ui"
	"github.com/smartassets/installer/util"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Starts already installed backend in the current directory",
	Long:    "Starts the installed backend. Upon command execution, the current directory must have the installed location of the project",
	Example: "assets start",
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
		ui.ReportInfoWithoutArgs("Starting server...")
		dockerComposeCmd := exec.Command("docker-compose",  "-p", "assets", "up", "-d")
		dockerComposeCmd.Stdout = os.Stdout
		dockerComposeCmd.Stderr = os.Stderr
		dockerComposeCmd.Stdin = os.Stdin
		err = dockerComposeCmd.Run()
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not start server: %s", err))
		}
	},
}