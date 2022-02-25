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
	redisUrl, _ := cmd.Flags().GetString("redis-url")
	opts := &server2.ServerOpts{
		RedisUrl: redisUrl,
	}
	err := server2.Run(opts)
	if err != nil {
		panic(err)
	}
}

func init() {
	serverCmd.Flags().String("redis-url", "localhost:6379", "The Redis connection for task management")
}
