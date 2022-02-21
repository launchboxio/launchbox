package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	applicationsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an application",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			applicationId, _ := cmd.Flags().GetUint("application-id")
			err := client.Apps().Delete(applicationId)
			if err != nil {
				panic(err)
			}
			ui.Raw("Delete successful")
		},
	}
)
