package command

import "github.com/spf13/cobra"

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
	}
)

func init() {
	configCmd.AddCommand(configCheckCmd)
	rootCmd.AddCommand(configCmd)
}
