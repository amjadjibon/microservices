package cmd

import (
	"github.com/spf13/cobra"

	"github.com/amjadjibon/microservices/auth"
)

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run the auth service",
	Run: func(cmd *cobra.Command, _ []string) {
		auth.Run()
	},
}
