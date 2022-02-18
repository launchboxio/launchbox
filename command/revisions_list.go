package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	revisionsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List a revision",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing a revision")
		},
	}
)
