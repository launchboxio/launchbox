package command

import (
	"github.com/spf13/cobra"
)

var (
	revisionsCmd = &cobra.Command{
		Use:   "revisions",
		Short: "Manager the revisions for a project",
	}
)

func init() {
	revisionsCmd.AddCommand(
		revisionsListCmd,
		revisionsGetCmd,
		revisionsCreateCmd,
		revisionsRevertCmd,
		revisionsLogsCmd,
	)

	revisionsCmd.PersistentFlags().Uint("project-id", 0, "Project ID")
	_ = revisionsCmd.MarkFlagRequired("project-id")
	rootCmd.AddCommand(revisionsCmd)
}
