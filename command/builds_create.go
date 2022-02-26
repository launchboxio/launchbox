package command

import (
	"context"
	"github.com/buildpacks/pack/pkg/client"
	"github.com/spf13/cobra"
)

var (
	buildsCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Start a build",
		Run: func(cmd *cobra.Command, args []string) {
			builder, err := client.NewClient()
			if err != nil {
				panic(err)
			}
			err = builder.Build(context.TODO(), client.BuildOptions{
				Image:    "launchbox",
				Builder:  "paketobuildpacks/builder:base",
				Registry: "http://localhost:5000",
			})
			if err != nil {
				panic(err)
			}
		},
	}
)
