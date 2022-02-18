package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	buildsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get build and status",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Fetching build status")
		},
	}
)
