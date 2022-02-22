package command

import (
	"github.com/spf13/cobra"
)

var (
	projectsCmd = &cobra.Command{
		Use:   "projects",
		Short: "Manage projects",
	}
)

func init() {
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
		projectsDeleteCmd,
	)

	projectsCmd.PersistentFlags().UintP("application-id", "", 0, "The application ID")
	_ = projectsCmd.MarkFlagRequired("application-id")
	rootCmd.AddCommand(projectsCmd)
}
