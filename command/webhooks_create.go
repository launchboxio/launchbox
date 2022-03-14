package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	webhooksCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new webhook",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")
			tagFilter, _ := cmd.Flags().GetString("tag-filter")
			branchFilter, _ := cmd.Flags().GetString("branch-filter")

			webhook := &api.Webhook{
				ProjectID:    projectId,
				TagFilter:    tagFilter,
				BranchFilter: branchFilter,
			}

			err := client.Webhooks().Create(projectId, webhook)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(webhook)
		},
	}
)

func init() {
	webhooksCreateCmd.Flags().String("tag-filter", "", "Filter of tags")
	webhooksCreateCmd.Flags().String("branch-filter", "", "Filter of branches")
}
