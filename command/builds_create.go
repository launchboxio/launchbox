package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	buildsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Start a build",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting a build")
		},
	}
)
