package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	authLoginCmd = &cobra.Command{
		Use:   "login",
		Short: "Authenticate to Launchbox",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cancelling build")
		},
	}
)
