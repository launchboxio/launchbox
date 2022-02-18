package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	projectsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating a project")
		},
	}
)
