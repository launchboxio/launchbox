package command

import (
	"github.com/gofrs/uuid"
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
)

var (
	secretsGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a secret",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			objectType, _ := cmd.Flags().GetString("object-type")
			objectId, _ := cmd.Flags().GetString("object-id")
			secretId, _ := cmd.Flags().GetString("id")
			key, _ := cmd.Flags().GetString("key")
			uid, _ := uuid.FromString(secretId)
			secret := &api.Secret{
				ID:         uid,
				ObjectType: objectType,
				ObjectId:   objectType,
				Name:       key,
			}

			err := client.Secrets().Find(secret)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(secret)
		},
	}
)

func init() {
	secretsCreateCmd.Flags().String("key", "", "The key for the secret")
	secretsGetCmd.Flags().String("id", "", "ID of the secret")
}
