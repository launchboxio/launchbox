package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	applicationsUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an application",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			applicationName, _ := cmd.Flags().GetString("name")
			applicationId, _ := cmd.Flags().GetUint("application-id")
			app := &api.Application{Name: applicationName, ID: applicationId}
			err := client.Apps().Update(app)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(app)
		},
	}
)

func init() {
	applicationsUpdateCmd.Flags().String("name", "", "Name of the application")
}
