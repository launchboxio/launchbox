package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	webhooksListCmd = &cobra.Command{
		Use:   "list",
		Short: "List generated webhooks",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")
			webhooks, err := client.Webhooks().List(projectId)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(webhooks)
		},
	}
)
