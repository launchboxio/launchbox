package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	revisionsLogsCmd = &cobra.Command{
		Use:   "logs",
		Short: "Get logs for a revision",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting logs for the revision")
		},
	}
)
