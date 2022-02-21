package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	applicationsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List generated applications",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			apps, err := client.Apps().List()
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(apps)
		},
	}
)
