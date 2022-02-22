package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	projectsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			applicationId, _ := cmd.Flags().GetString("application-id")
			apps, err := client.Projects().List(&api.ProjectListOptions{
				ApplicationId: applicationId,
			})
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(apps)
		},
	}
)
