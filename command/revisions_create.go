package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	revisionsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new revision",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			projectId, _ := cmd.Flags().GetUint("project-id")
			commitSha, _ := cmd.Flags().GetString("commit-sha")
			revision := &api.Revision{
				CommitSha: commitSha,
			}

			err := client.Revisions().Create(projectId, revision)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(revision)
		},
	}
)

func init() {
	revisionsCreateCmd.Flags().String("commit-sha", "", "The commit for this revision")
}
