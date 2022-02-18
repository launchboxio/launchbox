package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	projectsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting a project")
		},
	}
)
