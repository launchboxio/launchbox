package command

import (
	"github.com/spf13/cobra"
)

var (
	buildCmd = &cobra.Command{
		Use:   "builds",
		Short: "Manage builds",
	}
)

func init() {
	buildCmd.AddCommand(
		buildsCreateCmd,
		buildsGetCmd,
		buildsCancelCmd,
		buildsLogsCmd,
	)
	rootCmd.AddCommand(buildCmd)
}
