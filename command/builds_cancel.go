package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	buildsCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel a build",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cancelling build")
		},
	}
)
