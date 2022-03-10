package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	projectsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			applicationId, _ := cmd.Flags().GetUint("application-id")
			apps, err := client.Projects().List(&api.ProjectListOptions{
				ApplicationId: strconv.FormatUint(uint64(applicationId), 10),
			})
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(apps)
		},
	}
)
