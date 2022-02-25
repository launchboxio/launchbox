package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	secretsUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a secret",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Updating some secret")
		},
	}
)
