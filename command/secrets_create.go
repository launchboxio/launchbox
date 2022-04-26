package command

import (
	"github.com/launchboxio/launchbox/api"
	"github.com/spf13/cobra"
	"log"
)

var (
	secretsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new secret",
		Run: func(cmd *cobra.Command, args []string) {
			client, _ := api.New()
			ui := NewUI()

			objectType, _ := cmd.Flags().GetString("object-type")
			objectId, _ := cmd.Flags().GetString("object-id")
			key, _ := cmd.Flags().GetString("key")

			if len(args) < 1 {
				log.Fatalln("Please provide a value to store")
			}

			secret := &api.Secret{
				ObjectType: objectType,
				ObjectId:   objectId,
				Name:       key,
				Value:      args[0],
			}
			err := client.Secrets().Create(secret)
			if err != nil {
				panic(err)
			}
			ui.PrettyPrint(secret)
		},
	}
)

func init() {
	secretsCreateCmd.Flags().String("key", "", "The key for the secret")
}
