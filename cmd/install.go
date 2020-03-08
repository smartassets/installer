package cmd

import (
	"fmt"
	"github.com/smartassets/installer/ui"
	"github.com/smartassets/installer/util"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var installationDirectory string
var shouldStart bool

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Runs installation procedure",
	Long:  `Downloads and install latest smart assets binaries`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.ReportInfo("Checking if Docker is running...")
		err := util.CheckDocker()
		if err != nil {
			ui.ReportErr("Docker daemon is not working. In order the installer to work, Docker is required to be up and running.")
		}
		installationLocation, err := os.Stat(installationDirectory)
		if os.IsNotExist(err) {
			err = os.MkdirAll(installationDirectory, os.ModePerm)
			if err != nil {
				ui.ReportErr(err.Error())
			}
		}

		if !installationLocation.IsDir() {
			ui.ReportErr(fmt.Sprintf("location %s is not a directory", installationDirectory))
		}

		// Download the latest release of the project
		ui.ReportInfo("Downloading latest backend release...")
		resp, err := http.Get("https://github.com/smartassets/backend/releases/latest")
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not download latest release: %s", err.Error()))
		}
		defer resp.Body.Close()

		downloadFileName := filepath.Join(installationDirectory, "backend.zip")
		output, err := os.Create(downloadFileName)
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not open download file location: %s", err.Error()))
		}
		defer output.Close()

		_, err = io.Copy(output, resp.Body)
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not save download file: %s", err.Error()))
		}

		ui.ReportInfo("Extracting backend...")
		err = util.Unzip(downloadFileName, installationDirectory)
		if err != nil {
			ui.ReportErr(fmt.Sprintf("Could not unzip downloaded file: %s", err.Error()))
		}

		os.Remove(downloadFileName)

		if shouldStart {
			cmd := exec.Command("docker-compose", "-f", filepath.Join(installationDirectory, "docker-compose.yml", "up", "-d"))
			_, err := cmd.Output()
			if err != nil {
				ui.ReportErr(fmt.Sprintf("Could not startup server: %s", err))
			}
		} else {
			ui.ReportWarning("Options --start not specified, the backend will not be started.\n" +
				"Use the command 'assets start %s' in order to star the backend", installationDirectory)
		}

	},
}

func init() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		ui.ReportErr(err.Error())
	}
	installCmd.Flags().StringVar(&installationDirectory, "install-dir", currentDirectory, "installation directory (default: current working directory)")
	installCmd.Flags().BoolVar(&shouldStart, "start", false, "specifies whether server should start after installation (default: false)")
}
