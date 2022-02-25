package command

import "github.com/spf13/cobra"

var (
	secretsCmd = &cobra.Command{
		Use:   "secrets",
		Short: "Manage secrets",
	}
)

func init() {
	secretsCmd.AddCommand(
		secretsCreateCmd,
		secretsDeleteCmd,
		secretsUpdateCmd,
	)
	rootCmd.AddCommand(secretsCmd)
}
