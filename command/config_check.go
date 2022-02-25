package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	configCheckCmd = &cobra.Command{
		Use:   "check",
		Short: "Verify service configuration",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			fmt.Printf("Checking path %s\n", path)
		},
	}
)

func init() {
	configCheckCmd.Flags().String("path", "", "Path to a launchbox configuration")
	_ = configCheckCmd.MarkFlagRequired("path")
}
