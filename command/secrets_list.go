package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	secretsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List secrets",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			objectType, _ := cmd.Flags().GetString("object-type")
			objectId, _ := cmd.Flags().GetString("object-id")

			secrets, err := client.Secrets().List(objectType, objectId)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(secrets)
		},
	}
)
