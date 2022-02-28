package command

import (
	"github.com/spf13/cobra"
)

var (
	logsCmd = &cobra.Command{
		Use:   "logs",
		Short: "Get logs from a project",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func init() {
	rootCmd.AddCommand(logsCmd)
}
