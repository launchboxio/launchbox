package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	applicationsUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Updating the application")
		},
	}
)
