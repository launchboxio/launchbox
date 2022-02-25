package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	secretsDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a secret",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Deleting a secret")
		},
	}
)
