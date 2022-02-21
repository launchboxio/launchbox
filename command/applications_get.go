package command

import (
	"fmt"
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
			fmt.Println(applicationId)
			client, _ := api.New()

			app, err := client.Apps().Find(applicationId)
			fmt.Println(app)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(app)
		},
	}
)
