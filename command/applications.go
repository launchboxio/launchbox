package command

import (
	"github.com/spf13/cobra"
)

var (
	applicationsCmd = &cobra.Command{
		Use:   "applications",
		Short: "Manage applications",
	}
)

func init() {
	applicationsCmd.AddCommand(
		applicationsListCmd,
		applicationsGetCmd,
		applicationsCreateCmd,
		applicationsUpdateCmd,
		applicationsDeleteCmd,
	)

	applicationsCmd.PersistentFlags().UintP("application-id", "", 0, "Application ID")
	rootCmd.AddCommand(applicationsCmd)
}
