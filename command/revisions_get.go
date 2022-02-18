package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	revisionsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the current revision for a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting a revision")
		},
	}
)
