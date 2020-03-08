package cmd

import (
	"fmt"
	"github.com/smartassets/installer/ui"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Starts already installed backend",
	Example: "assets start <install-dir>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		installationDirectory := args[0]
		_, err := os.Stat(installationDirectory)
		if os.IsNotExist(err) {
			ui.ReportErr(fmt.Sprint("Location %s does not exists. Is the backend correctly installed?", installationDirectory))
		}

		dockerComposeCmd := exec.Command("docker-compose", "-f", filepath.Join(installationDirectory, "docker-compose.yml", "up", "-d"))
		_, err = dockerComposeCmd.Output()
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not start server: %s", err))
		}
	},
}
