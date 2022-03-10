package command

import (
	"github.com/launchboxio/launchbox/api"
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

			opts := &api.ApplicationFindOptions{}
			deleted, _ := cmd.Flags().GetBool("deleted")
			if deleted == true {
				opts.Deleted = true
			}
			app, err := client.Apps().Find(applicationId, opts)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(app)
		},
	}
)

func init() {
	applicationsGetCmd.Flags().Bool("deleted", false, "Show deleted applications")
}
