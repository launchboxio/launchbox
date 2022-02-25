package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	applicationsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get an individual application",
		Run: func(cmd *cobra.Command, args []string) {
			ui := NewUI()
			applicationId, _ := cmd.Flags().GetUint("application-id")
			client, _ := api.New()

			app, err := client.Apps().Find(applicationId)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(app)
		},
	}
)
