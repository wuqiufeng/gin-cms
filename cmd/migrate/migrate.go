// Initialize & migrate departments & users & maybe other stuffs
package migrate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	config     string
	MigrateCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize data before whole system ready(run once)",
		Example: "migrate",
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	MigrateCmd.PersistentFlags().StringVarP(&config, "config", "c", "./config/in-local.yaml", "Start server with provided configuration file")
}
func run() error {
	fmt.Println("run")

	return nil
}
func setup() {
	fmt.Println("setup")
}
