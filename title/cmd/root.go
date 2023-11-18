package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCMD.AddCommand(runCMD)
	rootCMD.AddCommand(migrateCMD)
}
