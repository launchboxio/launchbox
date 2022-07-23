package command

import (
	"github.com/launchboxio/launchbox/internal/server"
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

	serv, err := server.New(configFilePath)
	if err != nil {
		panic(err)
	}

	err = serv.Run()
	if err != nil {
		panic(err)
	}
}

func init() {
	serverCmd.Flags().StringP("config", "c", "config.yaml", "The location of server configuration file")
}
