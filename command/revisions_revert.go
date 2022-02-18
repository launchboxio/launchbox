package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	revisionsRevertCmd = &cobra.Command{
		Use:   "revert",
		Short: "Rollback the revision for a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Reverting a revision")
		},
	}
)
