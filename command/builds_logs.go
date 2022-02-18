package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	buildsLogsCmd = &cobra.Command{
		Use:   "logs",
		Short: "Get logs for a build",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Fetching build logs")
		},
	}
)
