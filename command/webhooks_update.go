package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	webhooksUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an webhook",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")

			// TODO: Handle params
			webhook := &api.Webhook{}
			err := client.Webhooks().Update(projectId, webhook)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(webhook)
		},
	}
)

func init() {
	webhooksUpdateCmd.Flags().String("tag-filter", "", "Filter of tags")
	webhooksUpdateCmd.Flags().String("branch-filter", "", "Filter of branches")
}
