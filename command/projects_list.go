package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	projectsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List services")
		},
	}
)
