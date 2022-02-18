package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	applicationsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List generated applications",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing applications")
		},
	}
)
