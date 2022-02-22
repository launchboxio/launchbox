package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	projectsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			applicationId, _ := cmd.Flags().GetUint("application-id")
			name, _ := cmd.Flags().GetString("name")
			repo, _ := cmd.Flags().GetString("repo")
			branch, _ := cmd.Flags().GetString("branch")
			project := &api.Project{
				Name:          name,
				Repo:          repo,
				Branch:        branch,
				ApplicationID: applicationId,
			}
			err := client.Projects().Create(project)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(project)
		},
	}
)

func init() {
	projectsCreateCmd.Flags().String("name", "", "The name of the project")
	projectsCreateCmd.Flags().String("repo", "", "The URL for the repository")
	projectsCreateCmd.Flags().String("branch", "", "The branch to deploy from")

	_ = projectsCreateCmd.MarkFlagRequired("name")
	_ = projectsCreateCmd.MarkFlagRequired("repo")
	_ = projectsCreateCmd.MarkFlagRequired("branch")
}
