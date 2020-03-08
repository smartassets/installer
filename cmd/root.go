package cmd

import (
	"github.com/smartassets/installer/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "assets",
	Short: "Smart assets CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.ReportErr(err.Error())
	}
}

func init(){
	rootCmd.AddCommand(installCmd)
}