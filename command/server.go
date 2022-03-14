package command

import (
	server2 "github.com/launchboxio/launchbox/server"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run the launchbox API server",
		Run:   RunServer,
	}
)

func RunServer(cmd *cobra.Command, args []string) {
	configFilePath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	err = server2.Run(configFilePath)
	if err != nil {
		panic(err)
	}
}

func init() {
	serverCmd.Flags().StringP("config", "c", "config.yaml", "The location of server configuration file")
}
