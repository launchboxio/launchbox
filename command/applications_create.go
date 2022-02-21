package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	applicationsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			applicationName, _ := cmd.Flags().GetString("name")
			app := &api.Application{Name: applicationName}
			err := client.Apps().Create(app)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(app)
		},
	}
)

func init() {
	applicationsCreateCmd.Flags().String("name", "", "Name of the application")
}
