package cmd

import (
	"github.com/amjadjibon/microservices/title"
	"github.com/spf13/cobra"
)

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run the auth service",
	Run: func(cmd *cobra.Command, _ []string) {
		title.Run()
	},
}
