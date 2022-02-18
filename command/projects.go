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
	rootCmd.AddCommand(projectsCmd)
}
