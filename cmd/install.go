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
		ui.ReportInfoWithoutArgs("Checking if Docker is running...")
		err := util.CheckDocker()
		if err != nil {
			ui.ReportErr("Docker daemon is not working. In order the installer to work, Docker is required to be up and running.")
		}
		installationLocation, err := os.Stat(installationDirectory)
		if os.IsNotExist(err) {
			ui.ReportInfo("Creating directory %s...", installationDirectory)
			err = os.MkdirAll(installationDirectory, os.ModePerm)
			if err != nil {
				ui.ReportErr(fmt.Sprintf("Could not create directory %s: %s", installationDirectory, err.Error()))
			}
		}

		if os.IsExist(err) {
			isEmpty, err := util.IsEmpty(installationDirectory)
			if err != nil {
				ui.ReportErr(fmt.Sprintf("Could not check contents of folder %s", installationDirectory))
			}
			if !isEmpty {
				// TODO: fix this
				// If the directory exists but the previous installation was successful, continue without doing anything new.
				// else print this message.
				ui.ReportErr(fmt.Sprintf("Directory %s already exists and it is not empty. Installation directory must be empty", installationDirectory))
			}
		}

		installationLocation, _ = os.Stat(installationDirectory)
		if !installationLocation.IsDir() {
			ui.ReportErr(fmt.Sprintf("Location %s is not a directory", installationDirectory))
		}

		// Download the latest release of the project
		ui.ReportInfoWithoutArgs("Downloading latest backend release...")
		resp, err := http.Get("https://github.com/smartassets/backend-server/releases/download/0.0.1/backend-server.zip")
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

		ui.ReportInfoWithoutArgs("Extracting backend...")
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
			ui.ReportInfo("Option --start not specified, the backend will not be started.\n"+
				"Use the command 'assets start %s' in order to star the backend", installationDirectory)
		}

	},
}

func init() {
	currentDirectory := getCurrentDirectory()
	installCmd.Flags().StringVar(&installationDirectory, "install-dir", currentDirectory, "installation directory (default: current working directory)")
	installCmd.Flags().BoolVar(&shouldStart, "start", false, "specifies whether server should start after installation (default: false)")
}

func getCurrentDirectory() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		ui.ReportErr(err.Error())
	}

	return currentDirectory
}
