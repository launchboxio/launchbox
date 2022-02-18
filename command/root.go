package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "launchbox",
	Short: "Launchbox: Democratizing Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Our name is cool...")
	},
}

func init() {
	rootCmd.AddCommand(operatorCmd)
	rootCmd.AddCommand(serverCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
