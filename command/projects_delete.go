package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	projectsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")
			err := client.Projects().Delete(projectId)
			if err != nil {
				panic(err)
			}
			ui.Raw("Delete successful")
		},
	}
)

func init() {
	projectsDeleteCmd.Flags().Uint("project-id", 0, "The project ID")
	_ = projectsDeleteCmd.MarkFlagRequired("project-id")
}
