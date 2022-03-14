package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	webhooksDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a webhook",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")

			webhookId, _ := cmd.Flags().GetUint("webhook-id")
			err := client.Webhooks().Delete(projectId, webhookId)
			if err != nil {
				panic(err)
			}
			ui.Raw("Delete successful")
		},
	}
)
