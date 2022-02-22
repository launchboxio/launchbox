package command

import (
	"github.com/robwittman/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	revisionsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the current revision for a project",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, err := cmd.Flags().GetUint("project-id")
			revisionId, err := cmd.Flags().GetUint("revision-id")
			revisions, err := client.Revisions().Find(projectId, revisionId)

			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(revisions)
		},
	}
)

func init() {
	revisionsGetCmd.Flags().Uint("revision-id", 0, "Revision ID")
}
