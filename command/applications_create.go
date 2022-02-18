package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	applicationsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating an application")
		},
	}
)
