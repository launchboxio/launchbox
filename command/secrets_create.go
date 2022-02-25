package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	secretsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new secret",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating a secret")
		},
	}
)
