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
		secretsListCmd,
		secretsGetCmd,
	)

	secretsCmd.PersistentFlags().String("object-type", "", "The object type of the secret. One of organization, application, secret")
	secretsCmd.PersistentFlags().String("object-id", "", "The ID of the object for this secret")
	secretsCmd.PersistentFlags().String("key", "", "The key to store the secret under")

	rootCmd.AddCommand(secretsCmd)
}
