package cmd

import (
	"cmsgo/cmd/api"
	"cmsgo/cmd/migrate"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "cmsgo",
	Short:             "cmsgo API server",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `Start cmsgo API server`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.MigrateCmd)

}

//Execute : run commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
