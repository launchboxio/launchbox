package command

import (
	"github.com/spf13/cobra"
)

var (
	infraCmd = &cobra.Command{
		Use:   "infra",
		Short: "Manage infrastructure runs",
	}
)

func init() {
	rootCmd.AddCommand(infraCmd)
}
