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
	lokiUrl, _ := cmd.Flags().GetString("loki-url")
	prometheusUrl, _ := cmd.Flags().GetString("prometheus-url")
	vaultUrl, _ := cmd.Flags().GetString("vault-url")
	opts := &server2.ServerOpts{
		RedisUrl:      redisUrl,
		LokiUrl:       lokiUrl,
		PrometheusUrl: prometheusUrl,
		VaultUrl:      vaultUrl,
	}
	err := server2.Run(opts)
	if err != nil {
		panic(err)
	}
}

func init() {
	serverCmd.Flags().String("redis-url", "localhost:6379", "The Redis connection for task management")
	serverCmd.Flags().String("loki-url", "http://loki.launchbox.local", "The web address for Loki")
	serverCmd.Flags().String("prometheus-url", "http://prometheus.launchbox.local", "The web address for Prometheus")
	serverCmd.Flags().String("vault-url", "http://vault.launchbox.local", "The web address for Vault")
}
