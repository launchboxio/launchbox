package command

import (
	server2 "github.com/robwittman/launchbox/server"
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
	opts := &server2.ServerOpts{}
	server := server2.New(opts)
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
