package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	projectsUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Updating a project")
		},
	}
)
