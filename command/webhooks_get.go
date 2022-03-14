package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	webhooksGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get an individual webhook",
		Run: func(cmd *cobra.Command, args []string) {
			ui := NewUI()
			webhookId, _ := cmd.Flags().GetUint("webhook-id")
			client, _ := api.New()

			projectId, _ := cmd.Flags().GetUint("project-id")

			webhook, err := client.Webhooks().Find(projectId, webhookId)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(webhook)
		},
	}
)
