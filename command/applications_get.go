package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	applicationsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get an individual application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Getting you applicaiton")
		},
	}
)
