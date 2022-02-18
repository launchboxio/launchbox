package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	projectsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Deleting project..")
		},
	}
)
