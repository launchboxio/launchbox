package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	revisionsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List a revision",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")
			projects, err := client.Revisions().List(projectId)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(projects)
		},
	}
)
