package command

import (
	"github.com/spf13/cobra"
)

var (
	webhooksCmd = &cobra.Command{
		Use:   "webhooks",
		Short: "Manage webhooks",
	}
)

func init() {
	webhooksCmd.AddCommand(
		webhooksListCmd,
		webhooksGetCmd,
		webhooksCreateCmd,
		webhooksUpdateCmd,
		webhooksDeleteCmd,
	)

	webhooksCmd.PersistentFlags().Uint("project-id", 0, "The Project ID")
	webhooksCmd.PersistentFlags().Uint("webhook-id", 0, "Webhook ID")
	rootCmd.AddCommand(webhooksCmd)
}
