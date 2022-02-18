package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	applicationsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Deleting an application")
		},
	}
)
