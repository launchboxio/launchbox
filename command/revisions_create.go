package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	revisionsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new revision",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating a revision")
		},
	}
)
