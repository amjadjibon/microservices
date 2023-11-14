package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/microservices/auth/conf"
	"github.com/amjadjibon/microservices/auth/infra"
)

var migrateCMD = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

var migrateUpCMD = &cobra.Command{
	Use:   "up",
	Short: "Run migrations up",
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := conf.GetConfig()
		if err := infra.MigrationUp(cfg.DatabaseDSN); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("No migrations files to run up")
				os.Exit(0)
			}

			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Migration up completed")
	},
}

var migrateDownCMD = &cobra.Command{
	Use:   "down",
	Short: "Run migrations down",
	Run: func(cmd *cobra.Command, _ []string) {
		cfg := conf.GetConfig()
		if err := infra.MigrationDown(cfg.DatabaseDSN); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("No more migrations files to run down")
				os.Exit(0)
			}

			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Migration down completed")
	},
}

func init() {
	migrateCMD.AddCommand(migrateUpCMD)
	migrateCMD.AddCommand(migrateDownCMD)
}
